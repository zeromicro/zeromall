
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


