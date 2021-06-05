package main

// snippet-start:[dynamodb.go.create_item.imports]
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// type SignUpObject struct {
// 	pk             string
// 	sk             string
// 	CustomerEmail  string
// 	CustomerMobile string
// 	CustomerName   string
// 	Password       string
// 	IsVerified     bool
// 	method         string
// }
// type ItemResponse struct {
// 	pk             string
// 	sk             string
// 	CustomerEmail  string
// 	CustomerMobile string
// 	CustomerName   string
// 	Password       string
// 	IsVerified     bool
// 	method         string
// }

func postHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("postHandler lambda started")
	fmt.Println(request.Body)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Starting a DynamoDB Session
	svc := dynamodb.New(sess)

	var userMap map[string]interface{}
	json.Unmarshal([]byte(request.Body), &userMap)
	// userMethod := userMap["method"].(string) //Using type assertion to concat "User#" prefix
	// fmt.Println(userMethod)
	// fmt.Println(userMap)

	// Checking if method is signup
	if userMap["method"] == "signup" {

		UserPrefix := "User#"

		fmt.Println("Inside Signup Method")

		tableName := "BrandTemplateTable"

		userMap["pk"] = UserPrefix + userMap["pk"].(string) //Using type assertion to concat "User#" prefix
		data, err2 := dynamodbattribute.MarshalMap(userMap)
		if err2 != nil {
			log.Fatalf("Got error marshalling new item: %s", err2)
		}

		fmt.Println(data)
		// DynamoDB getItem method for checking user existance
		result, err := svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"pk": {
					S: aws.String(userMap["CustomerEmail"].(string)),
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
		return events.APIGatewayProxyResponse{Body: string("User created"), StatusCode: 200}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string("Different method detected, cannot create User"), StatusCode: 400}, nil
	}
}

func main() {
	lambda.Start(postHandler)
}
