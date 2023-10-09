package main

import (
	"fmt"
	"ht-serverless-authorizer/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// main initializes aws session and dynamodb client
func main() {
	// region := os.Getenv("AWS_REGION")
	// awsSession, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(region)})

	// if err != nil {
	// 	return
	// }
	// dynaClient = dynamodb.New(awsSession)
	fmt.Print("print!")
	lambda.Start(handler)
}

const tableName = "serverless-api-golang-dev"

// handler takes care of our http method routing
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("print!")
	handlers.Authorize(req)
	fmt.Print("print!")
	return nil, nil
}
