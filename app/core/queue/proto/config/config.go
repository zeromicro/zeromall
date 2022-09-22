package config

import (
	"github.com/better-go/pkg/cache/redis"
	"github.com/better-go/pkg/database/orm"
	"github.com/better-go/pkg/x/go-zero/option"
)

// load option:
type Config struct {
	//
	// meta:
	//
	Server ServerUnit // 服务定义
	Client ClientUnit // 依赖服务 client(包含本服务的 rpc client)

	//
	// common required:
	//
	DB      DBUnit      // db
	Cache   CacheUnit   // cache
	MQ      MQUnit      // mq
	HTTP    HttpUnit    // http
	RPC     RpcUnit     // rpc
	GraphQL GraphQLUnit // graphql
	Job     JobUnit     // 异步 Job

	//
	// biz:
	//
	//Biz *BizUnit //
}

// 本服务定义: internal total
type ServerUnit struct {
	option.ServerUnit
}

// 依赖服务定义:
type ClientUnit struct {
	option.ClientUnit

	// 其他依赖服务:
	// TODO: required other services, add here
	Grace option.RpcClientConf // other rpc client
}

/////////////////////////////////////////////////////////////////////////////////////
// app  biz option unit:
/////////////////////////////////////////////////////////////////////////////////////

type MetaUnit struct {
	Name    string // app name
	Version string // version
}

// 业务本身配置参数:
type BizUnit struct {
	//ServiceName    string // this service name
	//ServiceToken   string // this service token
	//AuthGatewayUrl string // auth gateway url

	// biz:
	Demo DemoBiz
}

type DemoBiz struct {
}

/////////////////////////////////////////////////////////////////////////////////////
//  middleware option unit:
/////////////////////////////////////////////////////////////////////////////////////

// db:
type DBUnit struct {
	Demo *orm.Options // 示例 db
	//DB2 *orm.Options
}

/////////////////////////////////////////////////////////////////////////////////////

// cache:
type CacheUnit struct {
	Demo *redis.Options // 示例 redis
	//R2 *redis.Options
}

/////////////////////////////////////////////////////////////////////////////////////

// mq:
type MQUnit struct {
	Demo DemoMQ
}

type DemoMQ struct {
	Url     string
	Timeout string
}

/////////////////////////////////////////////////////////////////////////////////////
//  http option unit:
/////////////////////////////////////////////////////////////////////////////////////

type HttpUnit struct {
	//HomeServer *MatrixHttp
	//Email      *EmailHttp // SendCloud
}

/*


# Currently unused
SUBGRAPH_URL_1={mainnet-subgraph}
SUBGRAPH_URL_42={kovan-subgraph}
SUBGRAPH_URL_LOCAL={local-subgraph}
REACT_APP_SUPPORTED_NETWORK_ID=42
FROM_BLOCK
QUERY_INTERVAL=12

*/
type GraphQLUnit struct {
	Url           string // url
	Url42         string // SubGraph URL 42
	NetworkID     int32  // 区块链主链 ID
	BlockStart    int64  // 起始块高度
	QueryInterval int32  // 查询间隔周期
}

// Job 触发调度条件:
type JobUnit struct {
	IntervalMinute int64 // 间隔:
	TickerSecond   int32 //
	SlotNum        int   // 时间轮
	SleepDuration  int32 // 休眠

	// Job 计划表:
	Schedule JobSchedule
}

// Job 计划表:
type JobSchedule struct {
	QueryBlock string // 块分析
}

/////////////////////////////////////////////////////////////////////////////////////
//  rpc option unit:
/////////////////////////////////////////////////////////////////////////////////////

type RpcUnit struct {
	//DBCached *DemoRpc
}

type DemoRpc struct {
	AppKey    string
	AppSecret string
}
