curl --request POST \
  --url http://localhost:3001/conversations \
  --header 'Content-Type: application/json' \
  --data '{
  "user_ids": [
    "[UUID]",
    "[UUID]"
  ],
  "name": "Private chat"
}'

curl --request POST \
  --url http://localhost:3001/conversations \
  --header 'Content-Type: application/json' \
  --data '{
  "user_ids": [
    "[UUID]",
    "[UUID]",
    "[UUID]"
  ],
  "name": "Group chat"
}'
