// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : init.go
// Original timestamp : 2024/08/01 14:16

package auth

import (
	"github.com/hashicorp/vault/api"
	"os"
)

type AuthManager struct {
	vaultServerAddr string
	client          *api.Client
}

func NewAuthManager(client *api.Client, addr string) *AuthManager {
	if addr == "" {
		addr = os.Getenv("VAULT_ADDR")
		if addr == "" {
			return nil
		}
	}
	return &AuthManager{client: client, vaultServerAddr: addr}
}
