# scripts/list.sh
#!/bin/bash

URL="http://localhost:8083/workflows"
curl -X GET "$URL" -H "Content-Type: application/json"
