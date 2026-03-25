#!/bin/bash

# 1. Load variables from .env file if it exists
if [ -f .env ]; then
  export $(echo $(cat .env | sed 's/#.*//g' | xargs) | envsubst)
fi

# 2. Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
  echo "Error: DATABASE_URL is not set. Please check your .env file."
  exit 1
fi

# 3. Execute the Atlas command
echo "Applying schema to database..."

atlas schema apply \
  --url "$DATABASE_URL" \
  --to "file://schema.sql" \
  --dev-url "docker://postgres/15/dev?search_path=public"