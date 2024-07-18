# Copyright 2024 Robert Cronin
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pika
import json
import os
import random
import base64
from datetime import datetime, timezone
from dotenv import load_dotenv
from openai import OpenAI
from ratelimit import limits, sleep_and_retry

load_dotenv()

# RabbitMQ connection
rmq_url = os.getenv('RABBITMQ_URL')
connection = pika.BlockingConnection(pika.URLParameters(rmq_url))
channel = connection.channel()

# Declare the queues
channel.queue_declare(queue='poem_requests')
channel.queue_declare(queue='poem_responses')

# OpenAI setup
client = OpenAI(api_key=os.getenv('OPENAI_API_KEY'))

# Rate limiting: 50 messages per hour


@sleep_and_retry
@limits(calls=50, period=3600)
def generate_poem(prompt):
    response = client.completions.create(
        model="gpt-3.5-turbo-instruct",
        prompt=f"Create a Chinese 绝句 (jué jù) poem based on the following prompt: {prompt}",
        max_tokens=100,
        n=1,
        stop=None,
        temperature=0.7
    )
    return response.choices[0].text.strip()


def callback(ch, method, properties, body):
    # Decode the Base64 encoded message
    decoded_body = base64.b64decode(body).decode('utf-8')
    message = json.loads(decoded_body)

    poem_request_id = message['id']
    user_id = message['user_id']
    prompt = message['prompt']
    created_at = message['created_at']

    try:
        poem = generate_poem(prompt)

        response = {
            'id': poem_request_id,
            'user_id': user_id,
            'prompt': prompt,
            'poem': poem,
            'status': 'completed',
            'created_at': created_at,
            'updated_at': datetime.now(timezone.utc).isoformat()
        }

        # Publish the response
        channel.basic_publish(
            exchange='',
            routing_key='poem_responses',
            body=json.dumps(response)
        )

        print(f"Generated poem for request {poem_request_id}: {poem}")

        # Acknowledge the message
        ch.basic_ack(delivery_tag=method.delivery_tag)
    except Exception as e:
        print(f"Error processing request {poem_request_id}: {str(e)}")

        response = {
            'id': poem_request_id,
            'user_id': user_id,
            'prompt': prompt,
            'poem': None,
            'status': 'failed',
            'error': str(e),
            'created_at': created_at,
            'updated_at': datetime.now(timezone.utc).isoformat()
        }

        # Publish the error response
        channel.basic_publish(
            exchange='',
            routing_key='poem_responses',
            body=json.dumps(response)
        )

        # Negative acknowledgement
        ch.basic_nack(delivery_tag=method.delivery_tag, requeue=False)


# Set up the consumer
channel.basic_consume(queue='poem_requests', on_message_callback=callback)

print('AI Service is waiting for messages. To exit press CTRL+C')
channel.start_consuming()
