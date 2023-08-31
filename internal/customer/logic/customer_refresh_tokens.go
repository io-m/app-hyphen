package customer_logic

import (
	"context"
	"fmt"
)

func (cl *customerLogic) SaveRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return cl.customerOutgoing.SaveRefreshToken(ctx, customerId, refreshToken)
}

func (cl *customerLogic) VerifyRefreshToken(ctx context.Context, customerId, refreshToken string) (bool, error) {
	savedToken, err := cl.customerOutgoing.RetrieveRefreshToken(ctx, customerId, refreshToken)
	if err != nil || savedToken == "" {
		return false, fmt.Errorf("refresh token for customer %s is not found", customerId)
	}
	return true, nil
}

func (cl *customerLogic) DeleteRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return cl.customerOutgoing.DeleteRefreshToken(ctx, customerId, refreshToken)
}
