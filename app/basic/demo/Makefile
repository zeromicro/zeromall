

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

# echo test:
curl.api:
	curl -i http://localhost:8888/from/you
