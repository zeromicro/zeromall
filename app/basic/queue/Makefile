


############################################################
# init db:
#	- https://segmentfault.com/a/1190000020211225
#	dump:
#		- docker exec -it mysql mysqldump -u用户名 -p密码 数据库 > /mnt/vdb/data/mysql/test_db.sql
############################################################

db.init:
	docker exec -i mysql mysql --default-character-set=utf8mb4 -uroot -proot dev < ./proto/model/schema.sql

db.dump:
	docker exec -it mysql mysqldump --no-data --default-character-set=utf8mb4 -R -E --hex-blob -uroot -proot dev > ./tmp/backup_schema.sql

db.dump.with.data:
	docker exec -it mysql mysqldump --default-character-set=utf8mb4 -R -E --hex-blob -uroot -proot dev > ./tmp/backup.sql


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


