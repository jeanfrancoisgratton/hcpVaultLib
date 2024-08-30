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

// getEngineVersion determines the KV engine version for the given path
func (km *KVManager) getEngineVersion(path string) (int, *cerr.CustomError) {
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

// ReadSecret reads a secret from the KV engine at the given path
//
//	func (km *KVManager) ReadSecret(path string) (map[string]interface{}, *cerr.CustomError) {
//		version, ce := km.getEngineVersion(path)
//		if ce != nil {
//			return nil, ce
//		}
//
//		var secret *api.Secret
//		var err error
//		if version == 2 {
//			secret, err = km.client.Logical().Read("secret/data/" + path)
//		} else {
//			secret, err = km.client.Logical().Read("secret/" + path)
//		}
//		if err != nil {
//			return nil, &cerr.CustomError{Title: "Unable to read secret", Message: err.Error()}
//		}
//		if secret == nil {
//			return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
//		}
//
//		if version == 2 {
//			return secret.Data["data"].(map[string]interface{}), nil
//		}
//		return secret.Data, nil
//	}
func (km *KVManager) ReadSecret(path, field string) (interface{}, *cerr.CustomError) {
	engineVersion, err := km.getEngineVersion(path)
	var nErr error
	if err != nil {
		return nil, err
	}

	var secret *api.Secret
	if engineVersion == 2 {
		secret, nErr = km.client.Logical().Read("secret/data/" + path)
	} else {
		secret, nErr = km.client.Logical().Read("secret/" + path)
	}

	if nErr != nil {
		return nil, &cerr.CustomError{Title: "Unable to read secret", Message: nErr.Error()}
	}
	if secret == nil || secret.Data == nil {
		return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
	}

	var data map[string]interface{}
	if engineVersion == 2 {
		data = secret.Data["data"].(map[string]interface{})
	} else {
		data = secret.Data
	}

	if value, ok := data[field]; ok {
		return value, nil
	}

	return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
}

// WriteSecret writes a secret to the KV engine at the given path
//func (km *KVManager) WriteSecret(path string, data map[string]interface{}) *cerr.CustomError {
//	version, ce := km.getEngineVersion(path)
//	if ce != nil {
//		return ce
//	}
//
//	var writeData map[string]interface{}
//	var err error
//	if version == 2 {
//		writeData = map[string]interface{}{
//			"data": data,
//		}
//		_, err = km.client.Logical().Write("secret/data/"+path, writeData)
//	} else {
//		_, err = km.client.Logical().Write("secret/"+path, data)
//	}
//	if err != nil {
//		return &cerr.CustomError{Title: "Unable to write secret", Message: err.Error()}
//	}
//	return nil
//}

func (km *KVManager) WriteSecret(path, field string, value interface{}) *cerr.CustomError {
	var nErr error
	engineVersion, err := km.getEngineVersion(path)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		field: value,
	}

	var writeData map[string]interface{}
	if engineVersion == 2 {
		writeData = map[string]interface{}{
			"data": data,
		}
	} else {
		writeData = data
	}

	var writePath string
	if engineVersion == 2 {
		writePath = "secret/data/" + path
	} else {
		writePath = "secret/" + path
	}

	_, nErr = km.client.Logical().Write(writePath, writeData)
	if nErr != nil {
		return &cerr.CustomError{Title: "Unable to write secret", Message: nErr.Error()}
	}
	return nil
}

// DeleteSecret deletes a secret from the KV engine at the given path
func (km *KVManager) DeleteSecret(path string) *cerr.CustomError {
	version, ce := km.getEngineVersion(path)
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
	version, ce := km.getEngineVersion(path)
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
