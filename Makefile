
################################################################################################
# import other Makefiles:
################################################################################################

# include ./app/basic/demo/Makefile


################################################################################################
# update requirements:
################################################################################################

# update package:
.PHONY:
go.mod.tidy:
	go mod tidy -v

go.require.add:
	go get -u github.com/tal-tech/go-zero
	GO111MODULE=on go get -u github.com/tal-tech/go-zero/tools/goctl


################################################################################################
# gen demo:
################################################################################################

# gen api sever:
new.service.api:
	goctl api new ./tmp/demo

# has path bug:
new.service.rpc:
	goctl rpc new ./tmp/demo

# gen proto file:
new.service.proto:
	goctl rpc template -o ./tmp/demo.proto

################################################################################################
# run demo:
################################################################################################

# api server:
run.demo:
	go run tmp/demo/demo.go -f tmp/demo/etc/demo-api.yaml

# echo test:
curl.api:
	curl -i http://localhost:8888/from/you
