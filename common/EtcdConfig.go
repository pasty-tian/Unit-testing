package common

import (
	"github.com/Unit-testing/modules"
	"time"
)

type EtcdConfig struct {
	Endpoints   []string
	DialTimeout time.Duration
}

// get etcd config file
func ConfigEtcd() (*EtcdConfig, error) {
	etcdconfig, err := modules.GetViperConfig("Etcd")
	if err != nil {
		panic("read etcd config error")
		return nil, err
	}
	etcdcfg := &EtcdConfig{
		Endpoints:   etcdconfig.GetStringSlice("host"),
		DialTimeout: etcdconfig.GetDuration("timeout"),
	}
	return etcdcfg, nil

}
