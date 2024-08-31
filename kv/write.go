package kv

import cerr "github.com/jeanfrancoisgratton/customError"

// WriteSecret writes a secret to the KV engine at the given path
func (km *KVManager) WriteSecret(path, field string, value interface{}) *cerr.CustomError {
	var nErr error
	engineVersion, err := km.GetEngineVersion(path)
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
