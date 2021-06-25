package main

// snippet-start:[dynamodb.go.create_item.imports]
import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// type items struct {
// }

func userDetailsQueryHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("user details lambda started")
	fmt.Println(request.Body)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Starting a DynamoDB Session
	svc := dynamodb.New(sess)

	headers := make(map[string]string)
	var userMap map[string]interface{}
	// var queryCustomerDetailsMap map[string]interface{}
	var userDetails []map[string]interface{}
	// var userArray map[string]interface{}
	json.Unmarshal([]byte(request.Body), &userMap)

	headers["Access-Control-Allow-Headers"] = "Content-Type"
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "OPTIONS,POST,GET"
	// userMethod := userMap["method"].(string) //Using type assertion to concat "User#" prefix
	// fmt.Println(userMethod)
	// fmt.Println(userMap)

	UserPrefix := "User#"
	tableName := "BrandDataTable"

	userMap["user_email"] = UserPrefix + userMap["user_email"].(string) //Using type assertion to concat "User#" prefix

	// DynamoDB getItem method for checking user existance
	var result, err = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userMap["user_email"].(string)),
					},
				},
			},
		},
	})

	jsonData := make(map[string]interface{})

	dynamodbattribute.UnmarshalListOfMaps(result.Items, &userDetails)
	for _, val := range userDetails {
		for k, _ := range val {
			if k == "sk" {
				if val[k] == "CustomerDetails" {
					jsonData["basic_info"] = val
				} else if strings.HasPrefix(val[k].(string), "ElectricityDetails#") {
					jsonData["electricity_details"] = val
				} else if strings.HasPrefix(val[k].(string), "GasDetails#") {
					jsonData["gas_details"] = val
				}
			}
		}
	}

	data, errJson := json.Marshal(jsonData)
	fmt.Println(jsonData)
	//temporary
	// for idx, _ := range result.Items {
	// 	dynamodbattribute.UnmarshalMap(result.Items[idx], &queryCustomerDetailsMap)
	// }

	if errJson != nil {
		return events.APIGatewayProxyResponse{Body: "Cannot convert to json string", Headers: headers, StatusCode: 400}, nil
	}

	// Checking if User with given email exist
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "User not found", Headers: headers, StatusCode: 404}, nil
	}
	// Returning the response object that has been written to the Database
	return events.APIGatewayProxyResponse{Body: string(data), Headers: headers, StatusCode: 200}, nil

}

func main() {
	lambda.Start(userDetailsQueryHandler)
}
