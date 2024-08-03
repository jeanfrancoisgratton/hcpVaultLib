// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// users/init.go :
// Original file timestamp: 2024.08.02 19:47:27

package users

import (
	"github.com/hashicorp/vault/api"
)

type UserManager struct {
	client *api.Client
}

func NewUserManager(client *api.Client) *UserManager {
	return &UserManager{client: client}
}
