package config

import "github.com/zeromicro/go-zero/rest"

// load config:
type Config struct {
	rest.RestConf

	//
	Meta *MetaUnit // core meta

	//
	// common:
	//
	DB    *DBUnit    // db
	Cache *CacheUnit // cache
	MQ    *MQUnit    // mq
	HTTP  *HttpUnit  // http
	RPC   *RpcUnit   // rpc

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
	//Demo *DemoMQ
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
