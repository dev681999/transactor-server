# sleep for sometime and create the default signoz account
sleep 5
curl -X 'POST' \
    'http://localhost:3301/api/v1/register' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "name": "Admin",
    "orgName": "Admin",
    "email": "admin@admin.com",
    "password": "12345678"
}'