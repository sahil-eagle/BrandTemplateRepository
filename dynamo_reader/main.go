package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func dynamodbHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Starting dynamo reader lambda...")
	return events.APIGatewayProxyResponse{Body: string("Response"), StatusCode: 200}, nil

}

func main() {
	lambda.Start(dynamodbHandler)
}
