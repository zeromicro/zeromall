
###########################################################################
# start docker:
###########################################################################

run.mq.rabbitmq:
	docker-compose -f ./infra-mq-rabbitmq.yml up -d


run.kv.etcd:
	docker-compose -f ./infra-kv-etcd.yml up -d


run.mysql:
	docker-compose -f ./infra-db-mysql.yml up -d
	docker-compose -f ./infra-db-mysql.yml ps

run.redis:
	docker-compose -f ./infra-cache-redis.yml -p redis up -d
	docker-compose -f ./infra-cache-redis.yml ps


restart.kv.etcd:
	docker-compose -f ./infra-kv-etcd.yml restart


###########################################################################
# stop docker:
###########################################################################


stop.mq.rabbitmq:
	docker-compose -f ./infra-mq-rabbitmq.yml down

stop.kv.etcd:
	docker-compose -f ./infra-kv-etcd.yml down


# ETCDCTL_API =2 或者 3, 切换 API 版本
run.etcd:
	docker run -d -it --rm \
	--name etcd_test \
	-e ETCDCTL_API=3 \
	-p 2379:2379 \
	-p 2380:2380 \
	quay.io/coreos/etcd:v3.3.9 \
	etcd \
	--advertise-client-urls http://0.0.0.0:2379 \
	--listen-client-urls http://0.0.0.0:2379

exec.etcd:
	docker exec -it etcd_test sh
	# etcdctl ls

query.etcd.v2.keys:
	curl http://127.0.0.1:2379/v2/keys/ | python -m json.tool


# https://github.com/etcd-io/etcd/issues/6904
query.etcd.v3.keys:
	# ETCDCTL_API=3 ./etcdctl get / --prefix --keys-only
	etcdctl get "" --prefix --keys-only | sed '/^\s*$/d'


# 提取本机 IP:
#	- https://cloud.tencent.com/developer/article/1015510
ip.get.mac:
	#ifconfig | grep inet | grep -v inet6 | grep -v 127
	ifconfig | grep inet | grep -v inet6 | grep -v 127 | cut -d ' ' -f2

ip.get.linux:
	ip a | grep inet | grep -v inet6 | grep -v 127 | sed 's/^[ \t]*//g' | cut -d ' ' -f2
