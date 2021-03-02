
################################################################################################
# import other Makefiles:
################################################################################################

# include ./app/basic/demo/Makefile


################################################################################################
# update requirements:
################################################################################################

dev.require:
	# https://pre-commit.com/
	brew install pre-commit
	# 不建议使用 pip 安装, 如果对 Python 很熟悉, 可以用
	# pip install pre-commit


# update package:
.PHONY:
go.mod.tidy:
	go mod tidy -v

go.require.add:
	go get -u github.com/tal-tech/go-zero
	GO111MODULE=on go get -u github.com/tal-tech/go-zero/tools/goctl
	go get -u -v github.com/cosmtrek/air

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
