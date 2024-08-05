// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : init.go
// Original timestamp : 2024/08/01 19:36

package vault

import (
	"github.com/hashicorp/vault/api"
	"os"
)

type VaultOperatorInfo struct {
	Keys []string `json:"keys"`
}

type VaultManager struct {
	vaultServerAddr string
	client          *api.Client
}

// Initialize the Vault Manager. It returns a nil pointer if we cannot figure where is the VAULT server
func NewVaultManager(client *api.Client, addr string) *VaultManager {
	if addr == "" {
		addr = os.Getenv("VAULT_ADDR")
		if addr == "" {
			return nil
		}
	}
	return &VaultManager{client: client, vaultServerAddr: addr}
}

//
//// Initialize Vault client
//func InitVaultClient(addr string) (*api.Client, error) {
//	config := &api.Config{
//		Address: addr,
//	}
//	return api.NewClient(config)
//}
