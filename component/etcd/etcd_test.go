package etcd

import "testing"

func TestRegister(t *testing.T) {
	GetEtcd().RegisterService("test_service_name", "test_host")
}
