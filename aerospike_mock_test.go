package aerospike_test

import (
	plugin "github.com/G-Research/vault-plugin-database-aerospike"
	"github.com/aerospike/aerospike-client-go"
)

type MockClient struct {
	OnCreateUser     func(user string, password string, roles []string)
	OnChangePassword func(user string, password string)
	OnDropUser       func(user string)
}

type MockClientFactory struct {
	OnCreateUser     func(user string, password string, roles []string)
	OnChangePassword func(user string, password string)
	OnDropUser       func(user string)
}

func (f *MockClientFactory) NewClientWithPolicyAndHost(clientPolicy *aerospike.ClientPolicy, hosts ...*aerospike.Host) (plugin.Client, error) {
	client := &MockClient{
		OnCreateUser:     f.OnCreateUser,
		OnChangePassword: f.OnChangePassword,
		OnDropUser:       f.OnDropUser,
	}
	return client, nil
}

func (*MockClient) IsConnected() bool {
	return true
}

func (*MockClient) Close() {}

func (c *MockClient) CreateUser(policy *aerospike.AdminPolicy, user string, password string, roles []string) error {
	if c.OnCreateUser != nil {
		c.OnCreateUser(user, password, roles)
	}
	return nil
}

func (c *MockClient) DropUser(policy *aerospike.AdminPolicy, user string) error {
	if c.OnDropUser != nil {
		c.OnDropUser(user)
	}
	return nil
}

func (c *MockClient) ChangePassword(policy *aerospike.AdminPolicy, user string, password string) error {
	if c.OnChangePassword != nil {
		c.OnChangePassword(user, password)
	}
	return nil
}
