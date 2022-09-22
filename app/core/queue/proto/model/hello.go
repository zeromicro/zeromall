package model

import "github.com/better-go/pkg/database/orm"

type HelloEntity struct {
}

type HelloReq struct {
}

type HelloResp struct {
}

// 用户 Meta 信息
type UserMeta struct {
	orm.StatusModel

	UserID     uint64 `gorm:"UNIQUE; type:BIGINT(20) unsigned; NOT NULL;     DEFAULT:'0'; COMMENT:'用户 ID'; "`      // 用户 ID
	UniqueName string `gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案1:唯一昵称'; "`  // 二选一方案: 唯一昵称, 不允许重名
	NickName   string `gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案2:昵称+编号'; "` // 二选一方案: 昵称 + 昵称编号
	NickNo     int64  `gorm:"index;  type:varchar(128) CHARACTER SET utf8mb4; DEFAULT:'';  COMMENT:'方案2:昵称+编号'; "` // 二选一方案: 昵称编号
	// auth:
	Register   string
	MobileNo   string // 手机号
	MobileCode string // 手机号国家码
	Email      string // 邮箱
	Password   string // 密码
}
