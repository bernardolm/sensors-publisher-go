#!/bin/bash

set -e

echo "sh: starting..."

source ./dev/installer/clear.sh

NOW=$(date +"%Y%m%d_%H%M%S")
export RUN_FILE="sensors-publisher-go-installer_${NOW}.run"
echo "> run file is $RUN_FILE"

source ./dev/installer/build.sh

source ./dev/installer/send.sh

source ./dev/installer/run.sh

echo "sh: starting done"
