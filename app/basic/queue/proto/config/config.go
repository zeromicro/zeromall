package config

import "github.com/tal-tech/go-zero/rest"

// load config:
type Config struct {
	rest.RestConf

	//
	Meta *MetaUnit // core meta

	//
	// common:
	//
	DB      *DBUnit      // db
	Cache   *CacheUnit   // cache
	MQ      *MQUnit      // mq
	HTTP    *HttpUnit    // http
	RPC     *RpcUnit     // rpc
	GraphQL *GraphQLUnit // graphql
	Job     *JobUnit     // 异步 Job

	//
	// biz:
	//
	//Biz *BizUnit //
}

/////////////////////////////////////////////////////////////////////////////////////
// app  biz config unit:
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
	Demo *DemoBiz
}

type DemoBiz struct {
}

/////////////////////////////////////////////////////////////////////////////////////
//  middleware config unit:
/////////////////////////////////////////////////////////////////////////////////////

// db:
type DBUnit struct {
	//DB1 *orm.Options
	//DB2 *orm.Options
}

/////////////////////////////////////////////////////////////////////////////////////

// cache:
type CacheUnit struct {
	//R1 *redis.Options
	//R2 *redis.Options
}

/////////////////////////////////////////////////////////////////////////////////////

// mq:
type MQUnit struct {
	Demo *DemoMQ
}

type DemoMQ struct {
	Url     string
	Timeout string
}

/////////////////////////////////////////////////////////////////////////////////////
//  http config unit:
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
}

/////////////////////////////////////////////////////////////////////////////////////
//  rpc config unit:
/////////////////////////////////////////////////////////////////////////////////////

type RpcUnit struct {
	//Demo *DemoRpc
}

type DemoRpc struct {
	AppKey    string
	AppSecret string
}
