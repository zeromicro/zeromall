
#############################################################################################
# API测试:
#############################################################################################

# 本地运行:
run.local:
	go run ./cmd/main.go

# 自动watch 代码变更, 热加载:
run.auto.reload:
	# gin 运行模式: 默认 debug 模式
	# export GIN_MODE=release; air   # 生产环境
	export GIN_MODE=debug; air       # 开发环境

################################################################################################

# 生成项目 API 文档:
gen.swagger.docs:
	# swagger docs:
	swag init -g ./cmd/main.go

# 浏览器查看 API 文档:
open.swagger.docs:
	open http://127.0.0.1:8080/swagger/index.html
