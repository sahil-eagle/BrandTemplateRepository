package main

// snippet-start:[dynamodb.go.create_item.imports]
import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func searchQueryHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("search query lambda started")
	fmt.Println(request.Body)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Starting a DynamoDB Session
	svc := dynamodb.New(sess)

	var searchQueryMap map[string]interface{}
	var queryOutputCustomerDetailsMap map[string]interface{}
	var queryCustomerDetailsMpanMap map[string]interface{}
	json.Unmarshal([]byte(request.Body), &searchQueryMap)
	// userMethod := searchQueryMap["method"].(string) //Using type assertion to concat "User#" prefix
	// fmt.Println(userMethod)
	// fmt.Println(searchQueryMap)

	// Checking if method is signup
	if searchQueryMap["query_by"] == "email" {

		fmt.Println("Inside search by email Method")

		tableName := "BrandDataTable"
		userPrefix := "User#"

		searchQueryMap["query"] = userPrefix + searchQueryMap["query"].(string) //Using type assertion to concat "User#" prefix

		var result, err = svc.Query(&dynamodb.QueryInput{
			TableName: aws.String(tableName),
			KeyConditions: map[string]*dynamodb.Condition{
				"pk": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(searchQueryMap["query"].(string)),
						},
					},
				},
			},
		})

		fmt.Println(result.Items)
		dynamodbattribute.UnmarshalMap(result.Items[0], &queryOutputCustomerDetailsMap)

		resp, errJson := json.Marshal(queryOutputCustomerDetailsMap)

		filters := searchQueryMap["filters"]
		fmt.Println(filters)

		if errJson != nil {
			return events.APIGatewayProxyResponse{Body: "Cannot convert to json string", StatusCode: 400}, nil
		}

		// Checking if User with given email exist
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "User doesn't exists", StatusCode: 404}, nil
		}

		// Returning the response object that has been queries to the Database
		return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
	} else if searchQueryMap["query_by"] == "mpan" {
		fmt.Println("Inside search by mpan Method")

		tableName := "BrandDataTable"
		indexName := "gsi1Table"
		// userPrefix := "User#"
		// searchQueryMap["query"] = userPrefix + searchQueryMap["query"].(string) //Using type assertion to concat "User#" prefix

		fmt.Println(searchQueryMap["query"].(string))
		fmt.Println("*************************")
		var result, err = svc.Query(&dynamodb.QueryInput{
			TableName: aws.String(tableName),
			IndexName: aws.String(indexName),
			KeyConditions: map[string]*dynamodb.Condition{
				"gsi_1_pk": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(searchQueryMap["query"].(string)),
						},
					},
				},
			},
		})

		if err != nil {
			return events.APIGatewayProxyResponse{Body: "MPAN id not found", StatusCode: 404}, nil
		}

		dynamodbattribute.UnmarshalMap(result.Items[0], &queryCustomerDetailsMpanMap)
		resp, errJson := json.Marshal(queryCustomerDetailsMpanMap)
		if errJson != nil {
			return events.APIGatewayProxyResponse{Body: "Cannot convert to json string", StatusCode: 400}, nil
		}
		// Returning the response object that has been queries to the Database
		return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
	} else if searchQueryMap["query_by"] == "phone" {
		tableName := "BrandDataTable"
		indexName := "gsi1Table"
		var result, err = svc.Query(&dynamodb.QueryInput{
			TableName: aws.String(tableName),
			IndexName: aws.String(indexName),
			KeyConditions: map[string]*dynamodb.Condition{
				"gsi_1_pk": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(searchQueryMap["query"].(string)),
						},
					},
				},
			},
		})

		dynamodbattribute.UnmarshalMap(result.Items[0], &queryOutputCustomerDetailsMap)
		resp, errJson := json.Marshal(queryOutputCustomerDetailsMap)

		if errJson != nil {
			return events.APIGatewayProxyResponse{Body: "Cannot convert to json string", StatusCode: 400}, nil
		}

		// Checking if User with given email exist
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "User doesn't exists", StatusCode: 404}, nil
		}

		// Returning the response object that has been queries to the Database
		return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string("Invalid method"), StatusCode: 404}, nil
	}
}

func main() {
	lambda.Start(searchQueryHandler)
}
