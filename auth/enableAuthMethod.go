// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// auth/enableAuthMethod.go :
// Original file timestamp: 2024.07.31 06:03:36

package auth

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// EnableAuthMethod enables the specified auth method in Vault.
func EnableAuthMethod(client *api.Client, authType, defaultPath string) *cerr.CustomError {
	authMethods, ce := GetAuthMethodMountPaths(client)
	if ce != nil {
		return ce
	}

	mountPath, ok := authMethods[authType]
	if !ok {
		mountPath = defaultPath
	}

	authPath := fmt.Sprintf("sys/auth/%s", mountPath)
	data := map[string]interface{}{
		"type": authType,
	}
	if _, err := client.Logical().Write(authPath, data); err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Failed to enable auth method: %s", authType),
			Message: err.Error()}
	}
	return nil
}

// DisableAuthMethod disables the specified auth method in Vault.
func DisableAuthMethod(client *api.Client, authType, defaultPath string) *cerr.CustomError {
	authMethods, ce := GetAuthMethodMountPaths(client)
	if ce != nil {
		return ce
	}

	mountPath, ok := authMethods[authType]
	if !ok {
		mountPath = defaultPath
	}

	authPath := fmt.Sprintf("sys/auth/%s", mountPath)
	if _, err := client.Logical().Delete(authPath); err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Failed to disable auth method: %s", authType),
			Message: err.Error()}
	}

	return nil
}
