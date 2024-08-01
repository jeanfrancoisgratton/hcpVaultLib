// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : structs.go
// Original timestamp : 2024/08/01 14:16

package auth

import "github.com/hashicorp/vault/api"

type AuthManager struct {
	client *api.Client
}

func NewAuthManager(client *api.Client) *AuthManager {
	return &AuthManager{client: client}
}
