curl --request POST \
  --url http://localhost:3001/users \
  --header 'Content-Type: application/json' \
  --data '{
  "username": "1",
  "email": "test1@abc.com"
}'
curl --request POST \
  --url http://localhost:3001/users \
  --header 'Content-Type: application/json' \
  --data '{
  "username": "2",
  "email": "test2@abc.com"
}'
curl --request POST \
  --url http://localhost:3001/users \
  --header 'Content-Type: application/json' \
  --data '{
  "username": "3",
  "email": "test3@abc.com"
}'
