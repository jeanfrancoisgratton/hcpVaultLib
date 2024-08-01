// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : login.go
// Original timestamp : 2024/08/01 19:26

package vault

import (
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// Login with Token
func (vm *VaultManager) LoginWithToken(client *api.Client, token string) {
	client.SetToken(token)
}

// Login with AppRole
func (vm *VaultManager) LoginWithAppRole(client *api.Client, roleID, secretID string) (string, *cerr.CustomError) {
	data := map[string]interface{}{
		"role_id":   roleID,
		"secret_id": secretID,
	}
	secret, err := client.Logical().Write("auth/approle/login", data)
	if err != nil {
		return "", &cerr.CustomError{Title: "Unable to login", Message: err.Error()}
	}
	return secret.Auth.ClientToken, nil
}

// Login with UserPass
func (vm *VaultManager) LoginWithUserPass(client *api.Client, username, password string) (string, *cerr.CustomError) {
	data := map[string]interface{}{
		"password": password,
	}
	secret, err := client.Logical().Write("auth/userpass/login/"+username, data)
	if err != nil {
		return "", &cerr.CustomError{Title: "Unable to login", Message: err.Error()}
	}
	return secret.Auth.ClientToken, nil
}
