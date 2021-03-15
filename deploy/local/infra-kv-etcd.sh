#!/bin/bash


# MAC LOCAL IP:
LOCAL_IP=`ifconfig | grep inet | grep -v inet6 | grep -v 127 | cut -d ' ' -f2`

# LINUX LOCAL IP:
#LOCAL_IP=`ip a | grep inet | grep -v inet6 | grep -v 127 | sed 's/^[ \t]*//g' | cut -d ' ' -f2`


# set my host ip:
#export LOCAL_IP=192.168.0.188
export LOCAL_IP=${LOCAL_IP}


export DATA_DIR="etcd-data"

REGISTRY=quay.io/coreos/etcd
# available from v3.2.5
REGISTRY=gcr.io/etcd-development/etcd

docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  -e ETCDCTL_API=3 \
  --rm \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd ${REGISTRY}:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name LOCAL_IP \
  --initial-advertise-peer-urls http://${LOCAL_IP}:2380 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${LOCAL_IP}:2379 \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${LOCAL_IP}:2380
