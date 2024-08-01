// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : operations.go
// Original timestamp : 2024/08/01 14:06

package policies

func (pm *Policy) CreatePolicy(name, policy string) error {
	data := map[string]interface{}{
		"policy": policy,
	}
	_, err := pm.client.Logical().Write("sys/policies/acl/"+name, data)
	return err
}

func (pm *Policy) DeletePolicy(name string) error {
	_, err := pm.client.Logical().Delete("sys/policies/acl/" + name)
	return err
}

func (pm *Policy) ReadPolicy(name string) (string, error) {
	secret, err := pm.client.Logical().Read("sys/policies/acl/" + name)
	if err != nil {
		return "", err
	}
	return secret.Data["policy"].(string), nil
}
