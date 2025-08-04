package utils

import (
	"context"
	"errors"

	"github.com/MananLed/upKeepz-cli/internal/model"
)

func GetUserFromContext(ctx context.Context) (*model.User, error){
	user, ok := ctx.Value(UserContextKey).(*model.User)

	if !ok || user == nil{
		return nil, errors.New("user not found in context")
	}

	return user, nil
}