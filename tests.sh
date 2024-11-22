#!/bin/bash

# Directory containing the test scripts
SCRIPTS_DIR="./scripts"
LOG_FILE="./intense.log"

# Number of iterations for intensive testing
ITERATIONS=100
CONCURRENT_JOBS=5

# Clean up the log file
> $LOG_FILE

echo "Starting intensive test..." | tee -a $LOG_FILE

# Function to run all test scripts in one iteration
run_iteration() {
  local iteration=$1
  echo "Starting iteration $iteration..." | tee -a $LOG_FILE
  for script in "$SCRIPTS_DIR"/*.sh; do
    if [[ -f "$script" && -x "$script" ]]; then
      echo "Executing $script in iteration $iteration..." | tee -a $LOG_FILE
      bash "$script" >> $LOG_FILE 2>&1
      if [ $? -ne 0 ]; then
        echo "Error executing $script in iteration $iteration" | tee -a $LOG_FILE
      fi
    else
      echo "Skipping $script: not a file or not executable" | tee -a $LOG_FILE
    fi
  done
  echo "Completed iteration $iteration" | tee -a $LOG_FILE
}

# Export the function for use in parallel
export -f run_iteration
export SCRIPTS_DIR
export LOG_FILE

# Run iterations in parallel
seq 1 $ITERATIONS | xargs -n1 -P$CONCURRENT_JOBS -I{} bash -c 'run_iteration "$@"' _ {}

echo "Intensive test completed. Results are logged in $LOG_FILE" | tee -a $LOG_FILE
