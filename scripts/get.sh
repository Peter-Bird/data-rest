# scripts/get.sh
#!/bin/bash

WORKFLOW_ID="$1"
if [ -z "$WORKFLOW_ID" ]; then
  echo "Usage: $0 <workflow_id>"
  exit 1
fi

URL="http://localhost:8083/workflows/$WORKFLOW_ID"
curl -X GET "$URL"
