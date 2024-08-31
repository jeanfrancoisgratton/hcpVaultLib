// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : kv.go
// Original timestamp : 2024/08/01 18:48

package kv

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	"strconv"
	"strings"
)

// GetEngineVersion determines the KV engine version for the given path
func (km *KVManager) GetEngineVersion(path string) (int, *cerr.CustomError) {
	mounts, err := km.client.Sys().ListMounts()
	if err != nil {
		return 0, &cerr.CustomError{Title: "Unable to list mounts", Message: err.Error()}
	}

	for mountPath, mount := range mounts {
		if strings.HasPrefix(path, mountPath) {
			if versionStr, ok := mount.Options["version"]; ok {
				version, err := strconv.Atoi(versionStr)
				if err != nil {
					return 0, &cerr.CustomError{Title: fmt.Sprintf("Invalid version format for path: %s", path),
						Message: err.Error()}
				}
				return version, nil
			}
			return 1, nil // Default to version 1 if no version is specified
		}
	}

	return 0, &cerr.CustomError{Title: fmt.Sprintf("KV engine not found on path %s", path)}
}

// DeleteSecret deletes a secret from the KV engine at the given path
func (km *KVManager) DeleteSecret(path string) *cerr.CustomError {
	version, ce := km.GetEngineVersion(path)
	if ce != nil {
		return ce
	}

	var err error
	if version == 2 {
		_, err = km.client.Logical().Delete("secret/data/" + path)
	} else {
		_, err = km.client.Logical().Delete("secret/" + path)
	}
	if err != nil {
		return &cerr.CustomError{Title: "Unable to delete secret", Message: err.Error()}
	}
	return nil
}

// ListSecrets lists the secrets at the given path
func (km *KVManager) ListSecrets(path string) ([]string, *cerr.CustomError) {
	version, ce := km.GetEngineVersion(path)
	if ce != nil {
		return nil, ce
	}

	var secret *api.Secret
	var err error
	if version == 2 {
		secret, err = km.client.Logical().List("secret/metadata/" + path)
	} else {
		secret, err = km.client.Logical().List("secret/" + path)
	}
	if err != nil {
		return nil, &cerr.CustomError{Title: "Unable to list secrets", Message: err.Error()}
	}
	if secret == nil || secret.Data == nil {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
	}

	keys, ok := secret.Data["keys"].([]interface{})
	if !ok {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("invalid data at path: %s", path)}
		//return nil, fmt.Errorf("invalid data at path: %s", path)
	}

	var keyList []string
	for _, key := range keys {
		keyList = append(keyList, key.(string))
	}
	return keyList, nil
}

func (km *KVManager) GetLatestVersion(path string) (int, *cerr.CustomError) {
	engineVersion, err := km.GetEngineVersion(path)
	if err != nil {
		return 0, err
	}

	if engineVersion != 2 {
		return 0, &cerr.CustomError{Title: "Unable to get secret", Message: "versioning is only supported in KV version 2"}
	}

	secret, nE := km.client.Logical().Read("secret/metadata/" + path)
	if nE != nil {
		return 0, &cerr.CustomError{Title: "Unable to read secret", Message: nE.Error()}
	}

	if secret == nil || secret.Data == nil {
		return 0, &cerr.CustomError{Title: "Unable to read secret", Message: fmt.Sprintf("metadata for secret at path %s not found", path)}
	}

	if version, ok := secret.Data["current_version"].(int); ok {
		return version, nil
	}

	return 0, &cerr.CustomError{Title: "Unable to read secret", Message: fmt.Sprintf("unable to determine the latest version for secret %s", path)}
}
