package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"views-counter/src/badge"
	"views-counter/src/db"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	mongoDatabase := db.NewMongoDatabase()
	defer mongoDatabase.Close()

	createBadge := badge.MakeCreate()
	message := mongoDatabase.GetCurrentCount()
	response := createBadge(message)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "image/svg+xml"},
		Body:       response,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
