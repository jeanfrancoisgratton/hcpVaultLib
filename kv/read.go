package kv

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// ReadSecret reads a secret from the KV engine at the given path
//func (km *KVManager) ReadSecret(path, field string) (interface{}, *cerr.CustomError) {
//	engineVersion, err := km.GetEngineVersion(path)
//	var nErr error
//	if err != nil {
//		return nil, err
//	}
//
//	var secret *api.Secret
//	if engineVersion == 2 {
//		secret, nErr = km.client.Logical().Read("secret/data/" + path)
//	} else {
//		secret, nErr = km.client.Logical().Read("secret/" + path)
//	}
//
//	if nErr != nil {
//		return nil, &cerr.CustomError{Title: "Unable to read secret", Message: nErr.Error()}
//	}
//	if secret == nil || secret.Data == nil {
//		return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
//	}
//
//	var data map[string]interface{}
//	if engineVersion == 2 {
//		data = secret.Data["data"].(map[string]interface{})
//	} else {
//		data = secret.Data
//	}
//
//	if value, ok := data[field]; ok {
//		return value, nil
//	}
//
//	return nil, &cerr.CustomError{Title: fmt.Sprintf("no data found at path: %s", path)}
//}

func (km *KVManager) ReadSecret(path, field string, version int) (interface{}, *cerr.CustomError) {
	//var cErr *cerr.CustomError
	var secret *api.Secret
	var err error

	engineVersion, cErr := km.GetEngineVersion(path)
	if cErr != nil {
		return nil, cErr
	}

	if engineVersion == 2 {
		secretPath := fmt.Sprintf("secret/data/%s", path)
		if version > 0 {
			secretPath = fmt.Sprintf("%s?version=%d", secretPath, version)
		}
		secret, err = km.client.Logical().Read(secretPath)
	} else {
		secret, err = km.client.Logical().Read("secret/" + path)
	}

	if err != nil {
		return nil, &cerr.CustomError{Title: "Unable to read secret", Message: err.Error()}
	}
	if secret == nil || secret.Data == nil {
		return nil, &cerr.CustomError{Title: "Unable to read secret", Message: fmt.Sprintf("Secret at path %s not found", path)}
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
	return nil, &cerr.CustomError{Title: "Unable to read secret", Message: fmt.Sprintf("field %s not found in secret %s", field, path)}
}
