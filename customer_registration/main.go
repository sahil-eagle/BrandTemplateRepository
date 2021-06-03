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

type SignUpObject struct {
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

	// Starting a DynamoDB Session
	svc := dynamodb.New(sess)

	// Checking if method is signup
	if request.Body.methods == "signup" {
		var item SignUpObject
		err := json.Unmarshal([]byte(request.Body), &item) // converting the request body to json type(byte data) and then into struct object, assigning to item
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}

		UserPrefix := "User#"
		newUser := ItemResponse{
			pk:             UserPrefix + item.CustomerEmail,
			sk:             item.sk,
			CustomerEmail:  item.CustomerEmail,
			CustomerMobile: item.CustomerMobile,
			CustomerName:   item.CustomerName,
			Password:       item.Password,
			IsVerified:     item.IsVerified,
		}

		_, err = json.Marshal(&newUser) // Convert struct data into JSON
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil //Error if not found
		}
		tableName := "BrandTemplateTable"

		data, err2 := dynamodbattribute.MarshalMap(newUser)
		if err2 != nil {
			log.Fatalf("Got error marshalling new item: %s", err2)
		}

		// DynamoDB getItem method for checking user existance
		result, err := svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"pk": {
					S: aws.String(newUser.pk),
				},
			},
		})

		// Checking if User with given email exist
		if result != nil {
			return events.APIGatewayProxyResponse{Body: "User with this email already exist", StatusCode: 400}, nil
		}

		input := &dynamodb.PutItemInput{
			Item:      data,
			TableName: aws.String(tableName),
		}
		_, err = svc.PutItem(input)
		if err != nil {
			log.Fatalf("Got error calling PutItem: %s", err)
		}

		// Returning the response object that has been written to the Database
	}
	return events.APIGatewayProxyResponse{Body: string("Data written"), StatusCode: 200}, nil

}

func main() {
	lambda.Start(postHandler)
}
