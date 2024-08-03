// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// users/policies.go :
// Original file timestamp: 2024.08.02 22:01:16

package users

import "fmt"

// CreateUser creates a new user with the specified auth method
func (um *UserManager) CreateUser(authMethod, username, password string) error {
	switch authMethod {
	case "userpass":
		userData := map[string]interface{}{
			"password": password,
		}
		_, err := um.client.Logical().Write("auth/userpass/users/"+username, userData)
		return err
	case "approle":
		roleData := map[string]interface{}{
			"role_id": username,
		}
		_, err := um.client.Logical().Write("auth/approle/role/"+username, roleData)
		return err
	default:
		return fmt.Errorf("unsupported auth method: %s", authMethod)
	}
}

// DeleteUser deletes a user with the specified auth method and username
func (um *UserManager) DeleteUser(authMethod, username string) error {
	switch authMethod {
	case "userpass":
		_, err := um.client.Logical().Delete("auth/userpass/users/" + username)
		return err
	case "approle":
		_, err := um.client.Logical().Delete("auth/approle/role/" + username)
		return err
	default:
		return fmt.Errorf("unsupported auth method: %s", authMethod)
	}
}

// ListUsers lists all users for the specified auth method
func (um *UserManager) ListUsers(authMethod string) ([]string, error) {
	switch authMethod {
	case "userpass":
		users, err := um.client.Logical().List("auth/userpass/users")
		if err != nil {
			return nil, err
		}

		userList := []string{}
		if users != nil {
			for _, user := range users.Data["keys"].([]interface{}) {
				userList = append(userList, user.(string))
			}
		}
		return userList, nil
	case "approle":
		roles, err := um.client.Logical().List("auth/approle/role")
		if err != nil {
			return nil, err
		}

		roleList := []string{}
		if roles != nil {
			for _, role := range roles.Data["keys"].([]interface{}) {
				roleList = append(roleList, role.(string))
			}
		}
		return roleList, nil
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}
}
