#!/bin/bash

# run scraper
make

TOKEN=$(cat ../secrets/token)
ENDPOINT=$(cat ../secrets/endpoint)

JOB_ID=$(uuidgen)

curl --location --request GET '$ENDPOINT/run_pipeline' \
--data-raw '{
    "job_id":"$JOB_ID"
}'












