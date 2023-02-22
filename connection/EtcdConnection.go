package connection

import (
	"context"
	"github.com/Unit-testing/common"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func EtcdConnection() (*clientv3.Client, error) {
	etcdconfig, err := common.ConfigEtcd()
	if err != nil {
		panic("get etcdconfig failed")
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdconfig.Endpoints,
		DialTimeout: etcdconfig.DialTimeout * time.Second,
	})
	if err != nil {
		panic("etcd connection failed")
	}
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return cli, err
}
