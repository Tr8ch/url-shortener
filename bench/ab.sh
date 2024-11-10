#!/bin/bash

# Endpoint you want to test
URL="http://localhost:8888/shortener"

# File containing the POST data
POST_DATA_FILE="postdata.json"

# Content-Type header
CONTENT_TYPE="application/json"

# Number of parallel workers
CONCURRENCY=10

# Total number of requests (adjust based on expected throughput to simulate 30 seconds)
TOTAL_REQUESTS=1000 # Example value, adjust based on your expectations

# Execute ApacheBench
ab -p $POST_DATA_FILE -T $CONTENT_TYPE -c $CONCURRENCY -n $TOTAL_REQUESTS $URL
