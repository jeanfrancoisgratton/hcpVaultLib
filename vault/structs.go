// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : structs.go
// Original timestamp : 2024/08/01 19:36

package vault

import "github.com/hashicorp/vault/api"

type VaultOperatorInfo struct {
	Keys []string `json:"keys"`
}

type VaultManager struct {
	client *api.Client
}

// Initialize Vault client
func InitVaultClient(addr string) (*api.Client, error) {
	config := &api.Config{
		Address: addr,
	}
	return api.NewClient(config)
}