#!/bin/bash

# This shell script exports the environmental variables used in tests.
export DB_URL="postgresql://localhost:5432/test?user=postgres&password=password"
export LINE_CHANNEL_SECRET="dummy_channel_secret"
export LINE_CHANNEL_TOKEN="dummy_channel_token"
