# serverless-api-golang

### Building and Deploying App to AWS
run `npm install`

run `npm run deploy:dev`

### CURL Commands Test

#### CREATE User:
run `curl --header "Content-Type: application/json" --request POST --data '{ "email": "edrocha@gmail.com", "firstName": "ed", "lastName": "rocha", "dob": "3/7/2001" }' https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev/users`

#### GET All Users:
run `curl -X GET https://ribdhvkk9f.execute-api.us-east-1.amazonaws.com/dev/users`

#### GET User by UUID:
run `curl -X GET https://ribdhvkk9f.execute-api.us-east-1.amazonaws.com/dev/users/{user-uuid}`

#### UPDATE User:
run `curl --header "Content-Type: application/json" --request PUT --data '{ "email": "edrocha@gmail.com", "firstName": "lalala", "lastName": "lalala", "dob": "3/7/2001" }' https://ribdhvkk9f.execute-api.us-east-1.amazonaws.com/dev/users`

#### DELETE User:
run `curl -X DELETE https://ribdhvkk9f.execute-api.us-east-1.amazonaws.com/dev/users/{user-uuid}`

### Unit Tests
run ` go test ./... -coverprofile=coverage.out`# serverless-api-golang
