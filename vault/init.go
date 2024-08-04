// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : init.go
// Original timestamp : 2024/08/01 19:36

package vault

import (
	"github.com/hashicorp/vault/api"
)

type VaultOperatorInfo struct {
	Keys []string `json:"keys"`
}

type VaultManager struct {
	client *api.Client
}

func NewVaultManager(client *api.Client, addr string) *VaultManager {
	return &VaultManager{client: client}
}

//
//// Initialize Vault client
//func InitVaultClient(addr string) (*api.Client, error) {
//	config := &api.Config{
//		Address: addr,
//	}
//	return api.NewClient(config)
//}
