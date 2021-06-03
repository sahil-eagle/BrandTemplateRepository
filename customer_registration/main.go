package main

// snippet-start:[dynamodb.go.create_item.imports]
import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	pk             string
	sk             string
	CustomerEmail  string
	CustomerMobile string
	CustomerName   string
	Password       string
	IsVerified     bool
}
type ItemResponse struct {
	pk             string
	sk             string
	CustomerEmail  string
	CustomerMobile string
	CustomerName   string
	Password       string
	IsVerified     bool
}

func postHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	var item Item
	err := json.Unmarshal([]byte(request.Body), &item)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	UserPrefix := "User#"
	newItem := ItemResponse{
		pk:             UserPrefix + item.CustomerEmail,
		sk:             item.sk,
		CustomerEmail:  item.CustomerEmail,
		CustomerMobile: item.CustomerMobile,
		CustomerName:   item.CustomerName,
		Password:       item.Password,
		IsVerified:     item.IsVerified,
	}

	response, err := json.Marshal(&newItem)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	tableName := "BrandTemplateTable"

	av, err4 := dynamodbattribute.MarshalMap(newItem)
	if err4 != nil {
		log.Fatalf("Got error marshalling new item: %s", err4)
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(newItem.pk),
			},
		},
	})

	if result != nil {
		return events.APIGatewayProxyResponse{Body: "User with this email already exist", StatusCode: 400}, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

}

func main() {
	lambda.Start(postHandler)
}
