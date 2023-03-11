package logic

import (
	"context"
	"fmt"
	"github.com/better-go/pkg/random"
	"testing"
	"user/authn/rpc/pb"
)

// register test
func TestRegisterLogic_Register(t *testing.T) {
	ctx := context.Background()

	// new logic:
	logic := NewRegisterLogic(ctx, testSvcCtx)

	// rand email
	email := fmt.Sprintf("u%s@e%s.com", random.GenDigit(4), random.GenDigit(4))

	// req:
	req := &pb.UserRegisterReq{
		Username:     "user1",
		Password:     "pwd1234",
		Email:        email,
		RegisterType: pb.AuthType_Email,
	}

	ret, err := logic.Register(req)
	t.Logf("ret=%+v, err=%v", ret, err)

}

// login test:
func TestLoginLogic_Query(t *testing.T) {
	ctx := context.Background()

	// new logic:
	logic := NewRegisterLogic(ctx, testSvcCtx)

	ret, err := logic.svcCtx.MUser.FindOne(ctx, 1)
	t.Logf("ret=%+v, err=%v", ret, err)
}
