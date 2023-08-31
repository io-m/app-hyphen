package customer_logic

import "context"

func (cl *customerLogic) RetrieveAndVerifyRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return cl.customerOutgoing.RetrieveAndVerifyRefreshToken(ctx, refreshToken)
}
