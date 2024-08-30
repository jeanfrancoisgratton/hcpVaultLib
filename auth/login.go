// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// auth/login.go :
// Original file timestamp: 2024.07.30 22:15:37

package auth

import cerr "github.com/jeanfrancoisgratton/customError"

// GetAuthMethodMountPaths retrieves the mount paths for all enabled auth methods in Vault.
func (am *AuthManager) LoginWithAppRole(roleID, secretID string) (string, *cerr.CustomError) {
	data := map[string]interface{}{
		"role_id":   roleID,
		"secret_id": secretID,
	}
	secret, err := am.client.Logical().Write("auth/approle/login", data)
	if err != nil {
		return "", &cerr.CustomError{Title: "Unable to login", Message: err.Error()}
	}
	return secret.Auth.ClientToken, nil
}

func (am *AuthManager) LoginWithToken(token string) {
	am.client.SetToken(token)
}

func (am *AuthManager) LoginWithUserPass(username, password string) (string, *cerr.CustomError) {
	data := map[string]interface{}{
		"password": password,
	}
	secret, err := am.client.Logical().Write("auth/userpass/login/"+username, data)
	if err != nil {
		return "", &cerr.CustomError{Title: "Unable to login", Message: err.Error()}
	}
	return secret.Auth.ClientToken, nil
}
