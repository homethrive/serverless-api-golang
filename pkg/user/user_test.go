package user

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/google/uuid"
)

type mockDynamodbClient struct {
	dynamodbiface.DynamoDBAPI
}

var id = "4a228a82-cbf6-497d-a07a-d94179e1a3a0"
var email = "test@mail.com"
var firstName = "Test"
var lastName = "Tester"
var dob = "1/1/2001"

func (m *mockDynamodbClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	result := dynamodb.GetItemOutput{}

	if input.TableName == nil || *input.TableName == "" {
		return &result, errors.New("You must supply a table name")
	}

	item := User{
		Id:          id,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dob,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return &result, err
	}

	result = dynamodb.GetItemOutput{
		Item: av,
	}

	return &result, nil
}

func TestFetchUser(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	// mock resources
	id := uuid.NewString()
	table := "test-table-" + id

	mockSvc := &mockDynamodbClient{}

	item, err := FetchUser(id, table, mockSvc)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Retrieved user email '" + item.Email + "' from table ")
}
