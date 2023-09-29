#!/bin/bash

# Step 1: Build the Docker image
docker build -t hungddoo/telebot:latest .

# Step 2: Push the Docker image to Docker Hub
docker push hungddoo/telebot:latest

# Step 3: SSH into the host server and restart the Docker Compose orchestration
# ssh user@your-host-server "cd /path/to/your/docker-compose/directory && docker-compose down && docker-compose pull && docker-compose up -d"
