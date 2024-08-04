# Setup producer for user side
#

import pika
import os

def send_rabbit_msg(rabbit_url, input_message):
    # Set connection parameters
    rmq_url = os.getenv('RABBITMQ_URL')
    connection = pika.BlockingConnection(pika.URLParameters(rmq_url))
    channel = connection.channel()

    # Declare the queue from user to RabbitMQ
    channel.queue_declare(queue='poem_requests')
    # Declare the queue from RabbitMQ to user
    channel.queue_declare(queue='poem_responses')

    # Send message from user to RabbitMQ
    channel.basic_publish(
        exchange='', 
        routing_key='poem_requests', 
        body=input_message
    )

    print(f"Your prompt: {input_message}")

    # Close connection for cleanliness
    connection.close()