package http

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/lsendoya/cognitoDalas/internal/user/application"
	"github.com/lsendoya/cognitoDalas/pkg/aws"
	"github.com/lsendoya/cognitoDalas/pkg/db"
)

func Handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	cfg, err := aws.Config()
	if err != nil {
		return event, err
	}

	gormDB, errDB := db.ConnectDB(cfg)
	if errDB != nil {
		return event, errDB
	}

	return application.Service(gormDB, event)
}
