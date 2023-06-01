package router

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

func waitForTableActivated(db *dynamodb.Client, tbn string) error {
	w := dynamodb.NewTableExistsWaiter(db)
	err := w.Wait(context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tbn),
		},
		1*time.Minute,
	)
	if err != nil {
		return errors.Wrap(err, "Time out while waiting for table to become active")
	}
	return nil
}
