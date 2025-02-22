#!/bin/sh
set -e  # Exit on error

cd /app

if [ "$ENVIRONMENT" = "development" ]; then
    echo "Starting in development mode..."

    echo "Starting dev server..."
    air

    wait
else
    echo "Starting in production mode..."
    make build
    ./bin/app
fi
