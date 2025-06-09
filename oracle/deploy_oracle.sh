#!/bin/bash

echo "Starting Esport Oracle deployment..."

cd oracle

echo "Compiling project with Forge..."
forge build

if [ $? -ne 0 ]; then
    echo "Error during Forge compilation"
    exit 1
fi

echo "Compilation successful"

echo "Starting services with Docker Compose..."
docker-compose up --build

if [ $? -ne 0 ]; then
    echo "Error during Docker Compose execution"
    exit 1
fi

echo "Deployment completed"