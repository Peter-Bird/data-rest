# scripts/trunc.sh
#!/bin/bash

URL="http://localhost:8083/workflows"
curl -X DELETE "$URL" -H "Content-Type: application/json"

