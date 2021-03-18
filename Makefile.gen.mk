
# 项目全局 + 批量生成:
proto.gen.std:
	find ./app -iname "*.proto" -exec \
		protoc -I=. \
			-I=${GOPATH}/src \
			-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		 	--proto_path=. \
				--gofast_out=plugins=grpc:. \
		{} \;



# 项目全局 + 批量生成:
proto.gen.micro:
	cd ..; find ./creator-mate/app -iname "*.proto" -exec \
		protoc --proto_path=. \
				--micro_out=${MODIFY}:. \
				--go_out=paths=source_relative:. \
		{} \;


# 项目全局 + 批量生成:
proto.gen.data.only:
	find ./app -iname "*.proto" -exec \
		protoc --proto_path=. \
				--go_out=paths=source_relative:. \
		{} \;


app.gen.basic:
	cd app/basic; cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"

app.gen.biz:
	cd app/biz; cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"

