package customer_logic

import "context"

func (cl *customerLogic) SaveRefreshToken(ctx context.Context, refreshToken string) error {
	return cl.customerOutgoing.SaveRefreshToken(ctx, refreshToken)
}
