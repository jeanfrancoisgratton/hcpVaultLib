// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// policies/policies.go :
// Original file timestamp: 2024.08.02 22:06:36

package policies

import "fmt"

func (pm *PolicyManager) CreatePolicy(name, policy string) error {
	data := map[string]interface{}{
		"policy": policy,
	}
	_, err := pm.client.Logical().Write("sys/policies/acl/"+name, data)
	return err
}

func (pm *PolicyManager) DeletePolicy(name string) error {
	_, err := pm.client.Logical().Delete("sys/policies/acl/" + name)
	return err
}

func (pm *PolicyManager) ReadPolicy(name string) (string, error) {
	secret, err := pm.client.Logical().Read("sys/policies/acl/" + name)
	if err != nil {
		return "", err
	}
	return secret.Data["policy"].(string), nil
}

func (pm *PolicyManager) PolicyExists(name string) (bool, error) {
	secret, err := pm.client.Logical().Read("sys/policies/acl/" + name)
	if err != nil {
		return false, err
	}
	return secret != nil, nil
}

func (pm *PolicyManager) AssignPolicy(authMethod, username, policy string) error {
	exists, err := pm.PolicyExists(policy)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("policy %s does not exist", policy)
	}

	switch authMethod {
	case "userpass":
		userData := map[string]interface{}{
			"policies": []string{policy},
		}
		_, err = pm.client.Logical().Write("auth/userpass/users/"+username, userData)
		return err
	case "approle":
		roleData := map[string]interface{}{
			"token_policies": []string{policy},
		}
		_, err = pm.client.Logical().Write("auth/approle/role/"+username, roleData)
		return err
	default:
		return fmt.Errorf("unsupported auth method: %s", authMethod)
	}
}
