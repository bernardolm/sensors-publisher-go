#!/bin/bash

echo "sh: sending..."

rsync -r "tmp/${RUN_FILE}" "${SYSTEM_USER}@${SYSTEM_HOST}:/tmp"

echo "sh: sending done"
