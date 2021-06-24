// package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// )

// type Item struct {
// 	customerAddress  string
// 	customerCity     string
// 	customerEmail    string
// 	customerId       string
// 	customerMobile   string
// 	customerName     string
// 	customerPostcode string
// 	customerState    string
// }

// func queryHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// STarting dynamoDB session
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	// Create DynamoDB client
// 	svc := dynamodb.New(sess)
// 	// Here TableName is people, partiton key is PersonSSN and index is User_mail
// 	table_name := "BrandDataTable"
// 	index_name := "gsi1Table"

// 	//query using GSI(indexing)
// 	var queryInput, err2 = svc.Query(&dynamodb.QueryInput{
// 		TableName: aws.String(table_name),
// 		IndexName: aws.String(index_name),
// 		KeyConditions: map[string]*dynamodb.Condition{
// 			"gsi_1_pk": {
// 				ComparisonOperator: aws.String("EQ"),
// 				AttributeValueList: []*dynamodb.AttributeValue{
// 					{
// 						S: aws.String("1234"),
// 					},
// 				},
// 			},
// 		},
// 	})

// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	fmt.Println(queryInput)
// 	var item Item
// 	dynamodbattribute.UnmarshalMap(queryInput.Items[0], &item)
// 	jsonData, _ := json.Marshal(item)
// 	fmt.Println(string(jsonData))

// 	return events.APIGatewayProxyResponse{
// 		Body:       string(jsonData),
// 		StatusCode: 200,
// 	}, nil

// }

// func main() {
// 	lambda.Start(queryHandler)
// }
// ***************************************************************************

// GSI_2 table

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func queryHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// STarting dynamoDB session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// Here TableName is people, partiton key is PersonSSN and index is User_mail
	table_name := "BrandDataTable"
	index_name := "gsi2Table"

	//query using GSI(indexing)
	var queryInputPreRegistration, err = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("preRegistration"),
					},
				},
			},
		},
	})

	var queryInputRegistrationSubmitted, err2 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("registrationSubmitted"),
					},
				},
			},
		},
	})

	var queryInputRegistrationRejected, err3 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("registrationRejected"),
					},
				},
			},
		},
	})

	var queryInputAwaitingStart, err4 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("awaitingStartOfService"),
					},
				},
			},
		},
	})
	var queryInputFirstMonth, err5 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("firstMonth"),
					},
				},
			},
		},
	})
	var queryInputFirstBillingWindow, err6 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("firstBillingWindow"),
					},
				},
			},
		},
	})
	var queryInputInService, err7 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("inService"),
					},
				},
			},
		},
	})
	var queryInputRenewalWindow, err8 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("renewalWindow"),
					},
				},
			},
		},
	})
	var queryInputRetainedUsers, err9 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("retainedUsers"),
					},
				},
			},
		},
	})
	var queryInputChurnedUsers, err10 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(table_name),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"gsi_2_pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("churnedUsers"),
					},
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if err3 != nil {
		fmt.Println(err3)
	}
	if err4 != nil {
		fmt.Println(err4)
	}
	if err5 != nil {
		fmt.Println(err5)
	}
	if err6 != nil {
		fmt.Println(err6)
	}
	if err7 != nil {
		fmt.Println(err7)
	}
	if err8 != nil {
		fmt.Println(err8)
	}
	if err9 != nil {
		fmt.Println(err9)
	}
	if err10 != nil {
		fmt.Println(err10)
	}

	fmt.Println(queryInputRegistrationRejected)
	fmt.Println(queryInputRegistrationSubmitted)
	fmt.Println(queryInputAwaitingStart)
	fmt.Println(queryInputPreRegistration)
	fmt.Println(queryInputFirstMonth)
	fmt.Println(queryInputFirstBillingWindow)
	fmt.Println(queryInputInService)
	fmt.Println(queryInputRenewalWindow)
	fmt.Println(queryInputRetainedUsers)
	fmt.Println(queryInputChurnedUsers)
	// fmt.Println(queryInput.Items)
	// fmt.Println("***********************")
	// fmt.Println(queryInput.Items[0]["registrationStatus"])
	// fmt.Println("***********************")

	// userJourneyStages := make(map[string]int)
	// for i := 0; i < len(queryInput.Items); i++ {
	// 	val := queryInput.Items[i]["registrationStatus"].String()
	// 	userJourneyStages[val] += 1
	// }
	// val := queryInput.Items[0]["registrationStatus"].String()

	// fmt.Println(val)
	// fmt.Println("***********************")
	// fmt.Println(userJourneyStages)
	// var item Item
	// dynamodbattribute.UnmarshalMap(queryInput.Items[0], &item)
	// jsonData, _ := json.Marshal(item)
	// fmt.Println(string(jsonData))

	return events.APIGatewayProxyResponse{
		Body:       string(""),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(queryHandler)
}
