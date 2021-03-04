

################################################################################################
# update requirements:
################################################################################################

# update package:
.PHONY:
go.mod.tidy:
	go mod tidy -v


# api server:
run.demo:
	go run cmd/main.go -f configs/configs.yaml

# 自动watch 代码变更, 热加载:
run.auto.reload:
	air       # 开发环境

# echo test:
curl.api:
	curl -i http://localhost:8080/from/me

curl.api2:
	curl -i http://localhost:8080/hello/tom

	curl -i http://localhost:8080/hello/jim

