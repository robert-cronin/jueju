# Setup consumer for RabbitMQ side
#

import pika
import os

def on_message_received(ch, method, properties, body):
    print(f"received new prompt: {body}")

# Set connection parameters
rmq_url = os.getenv('RABBITMQ_URL')
connection = pika.BlockingConnection(pika.URLParameters(rmq_url))
channel = connection.channel()

# Note(Lani): it's ok that we declare the queue twice (ie in here and producer)
# Declare the queue from user to RabbitMQ
channel.queue_declare(queue='poem_requests')
# Declare the queue from RabbitMQ to user
channel.queue_declare(queue='poem_responses')

# Setup consume parameters
channel.basic_consume(queue='poem_requests', auto_ack=True,
    on_message_callback=on_message_received)

print("Starting Consuming")
# Start consumption
channel.start_consuming()