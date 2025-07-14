# TEST_API

### To run the application
> source configuration/env.sh && go run cmd/rest/main.go 

### Create Campaign
> curl --location 'localhost:8080/v1/campaign/create' \
--header 'Content-Type: application/json' \
--data '{
    "cid" : "duolingo",
    "name" : "Duolingo: Best way to learn ", 
    "img" : "https://somelink2", 
    "cta" : "Install",
    "status" : "ACTIVE"
}'

### Create Rule
> curl --location 'localhost:8080/v1/rule/create' \
--header 'Content-Type: application/json' \
--data '{
  "cid": "duolingo",
  "rules": {
    "exclude_country": ["US"],
    "include_os": ["Android"]
  }
}'

### Test Delivery service
> curl --location 'localhost:8080/v1/delivery?app=com.abc.xyz&country=germany&os=android' \
--header 'Content-Type: application/json'