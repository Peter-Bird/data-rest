# scripts/submit.sh
#!/bin/bash

WORKFLOW_ID="$1"
if [ -z "$WORKFLOW_ID" ]; then
  echo "Usage: $0 <workflow_id>"
  exit 1
fi

URL="http://localhost:8083/workflows"
WORKFLOW_JSON=$(cat <<EOF
{
  "id": "$WORKFLOW_ID",
  "name": "Sample Workflow",
  "steps": [
    {
      "endpoint": "/step1",
      "method": "POST",
      "parameters":     {
        "endpoint": "/step2",
        "method": "GET",
        "parameters": {},
        "dependencies": ["step1"]
    },
      "dependencies": []
    },
    {
      "endpoint": "/step2",
      "method": "GET",
      "parameters": {},
      "dependencies": ["step1"]
    }
  ]
}
EOF
)

curl -X POST -H "Content-Type: application/json" -d "$WORKFLOW_JSON" "$URL"

