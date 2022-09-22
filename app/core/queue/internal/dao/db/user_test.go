package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/better-go/pkg/database/orm"
	"github.com/better-go/pkg/random"

	"mall/app/core/queue/proto/model"
)

func TestUserStorage_UserMeta(t *testing.T) {
	ctx := context.Background()

	in := []*model.UserMeta{
		{
			StatusModel: orm.StatusModel{},
			UserID:      random.SnowFlakeID(),
			Register:    fmt.Sprintf("register-%v", random.SnowFlakeID()),
			UniqueName:  "username",
			NickName:    "jim",
			NickNo:      2233,
			MobileNo:    "17188661234",
			MobileCode:  "",
			Email:       "dd@aa.com",
			Password:    "aabbcc",
		},
		{
			StatusModel: orm.StatusModel{},
			UserID:      random.SnowFlakeID(),
			Register:    fmt.Sprintf("register-%v", random.SnowFlakeID()),
			UniqueName:  "user2",
			NickName:    "jim",
			NickNo:      2233,
			MobileNo:    "17188661200",
			MobileCode:  "",
			Email:       "mm@aa.com",
			Password:    "aabbcc",
		},
	}

	// db insert:
	for _, item := range in {
		result, err := testDao.User.CreateUserMeta(ctx, item)
		t.Logf("create: result=%v, err=%v", result, err)

	}

	// db query:
	for _, item := range in {
		// get:
		resp, err := testDao.User.UserMeta(ctx, item.UserID)
		t.Logf("get: user meta: %+v, err=%v", resp, err)
	}

}
