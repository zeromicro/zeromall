

################################################################################################
# update requirements:
################################################################################################

# update package:
.PHONY:
go.mod.tidy:
	go mod tidy -v


############################################################
# quick run:
############################################################

# internal api:
run.svc:
	#go run cmd/main.go -f=configs/configs.yaml --serverType=inner
	go run cmd/main.go --configFile=configs/configs.yaml --serverType=inner

# interface api:
run.api:
	go run cmd/main.go -f=configs/configs.yaml -r=outer

############################################################
# air run: auto reload, 自动watch 代码变更, 热加载
############################################################

# 服务类型:
SERVER_TYPE = ${serverType}
# 配置文件路径:
CONFIG_FILE = ${configFile}

# 编译:
build:
	go build -o ./tmp/main ./cmd

# 编译+执行: 传入类型参数e
build.run:
	./tmp/main --configFile=${CONFIG_FILE} --serverType=${SERVER_TYPE}

############################################################

# 编译+执行: 传入类型参数
build.run.way1: build
	./tmp/main --configFile=configs/configs.yaml --serverType=inner

build.run.way2: build
	SERVER_TYPE=inner CONFIG_FILE=configs/configs.yaml ./tmp/main

############################################################
# run 不同子服务:
############################################################

# 自动编译+执行 SVC: inner
run.auto.svc:
	export serverType=inner; export configFile=configs/configs.yaml; air -d    # 开发环境

# 自动编译+执行 API: outer/interface
run.auto.api:
	export serverType=outer; export configFile=configs/configs.yaml; air -d    # 开发环境

# 自动编译+执行 Job:
run.auto.job:
	export serverType=job; export configFile=configs/configs.yaml; air -d      # 开发环境

# 自动编译+执行 Admin:
run.auto.admin:
	export serverType=admin; export configFile=configs/configs.yaml; air -d      # 开发环境


############################################################
# API 测试:
############################################################

# echo test:
curl.api:
	curl -i http://localhost:8000/from/me

curl.api2:
	curl -i http://localhost:8000/hello/tom

	curl -i http://localhost:8000/hello/jim

curl.pub.msg:
	curl --request POST \
		--url http://localhost:8080/queue/publish/ \
		--header 'content-type: application/json' \
		--data '{"queue_name": "test-queue",	"type": "type1",	"data": "test messagexxxx6",	"raw_data": 1,	"timestamp": 2	}'

curl.graphql:
	curl -i http://localhost:8000/graphql/do

curl.admin.api:
	curl --request GET \
		--url 'http://localhost:8001/admin/api/rpc/test?from=jim' \
		--header 'Content-Type: application/json' \
		--data '{    	"from": "xxx"    }'


############################################################
# kill 进程:
############################################################

# under mac osx:
pid.kill:
	# rpc server
	lsof -ti :6000 | xargs kill
	# api server
	lsof -ti :8000 | xargs kill
	# admin api
	lsof -ti :8001 | xargs kill
	# job server
	lsof -ti :8002 | xargs kill
