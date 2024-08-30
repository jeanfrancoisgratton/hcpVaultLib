package kv

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// ReadSecret reads a secret from the KV engine at the given path
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
