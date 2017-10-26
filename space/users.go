package space

import (
	"fmt"
	"strings"

	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	. "github.com/pivotalservices/cf-mgmt/space/constants"
	"github.com/pivotalservices/cf-mgmt/ldap"
	"github.com/pivotalservices/cf-mgmt/uaac"
	"github.com/xchapter7x/lo"
)

// UserMgr - interface type encapsulating Update space users behavior
type UserMgr interface {
	UpdateSpaceUsers(config *ldap.Config, uaacUsers map[string]string, updateUsersInput UpdateUsersInput) error
}

// NewUserManager -
func NewUserManager(
	cloudController cloudcontroller.Manager,
	ldapMgr ldap.Manager,
	uaacMgr uaac.Manager) UserMgr {
	return &UserManager{
		cloudController: cloudController,
		LdapMgr:         ldapMgr,
		UAACMgr:         uaacMgr,
	}
}

type UserManager struct {
	cloudController cloudcontroller.Manager
	LdapMgr         ldap.Manager
	UAACMgr         uaac.Manager
}

// UpdateSpaceUserInput
type UpdateUsersInput struct {
	SpaceGUID                                   string
	OrgGUID                                     string
	Role                                        string
	LdapUsers, Users, LdapGroupNames, SamlUsers []string
	SpaceName                                   string
	OrgName                                     string
	RemoveUsers                                 bool
}

//UpdateSpaceUsers Update space users
func (m *UserManager) UpdateSpaceUsers(config *ldap.Config, uaacUsers map[string]string, updateUsersInput UpdateUsersInput) error {
	spaceUsers, err := m.cloudController.GetCFUsers(updateUsersInput.SpaceGUID, SPACES, updateUsersInput.Role)
	if err != nil {
		return err
	}

	lo.G.Debugf("SpaceUsers before: %v", spaceUsers)
	if config.Enabled {
		var ldapUsers []ldap.User
		ldapUsers, err = m.getLdapUsers(config, updateUsersInput.LdapGroupNames, updateUsersInput.LdapUsers)
		if err != nil {
			return err
		}
		lo.G.Debugf("LdapUsers: %v", ldapUsers)
		for _, user := range ldapUsers {
			err = m.updateLdapUser(config, updateUsersInput.SpaceGUID, updateUsersInput.OrgGUID, updateUsersInput.Role, updateUsersInput.OrgName, updateUsersInput.SpaceName, uaacUsers, user, spaceUsers)
			if err != nil {
				return err
			}
		}
	} else {
		lo.G.Debug("Skipping LDAP sync as LDAP is disabled (enable by updating config/ldap.yml)")
	}
	for _, userID := range updateUsersInput.Users {
		lowerUserID := strings.ToLower(userID)
		if _, userExists := uaacUsers[lowerUserID]; !userExists {
			return fmt.Errorf("user %s doesn't exist in cloud foundry, so must add internal user first", lowerUserID)
		}
		if _, ok := spaceUsers[lowerUserID]; !ok {
			if err = m.addUserToOrgAndRole(userID, updateUsersInput.OrgGUID, updateUsersInput.SpaceGUID, updateUsersInput.Role, updateUsersInput.OrgName, updateUsersInput.SpaceName); err != nil {
				lo.G.Error(err)
				return err
			}
		} else {
			delete(spaceUsers, lowerUserID)
		}
	}

	for _, userEmail := range updateUsersInput.SamlUsers {
		lowerUserEmail := strings.ToLower(userEmail)
		if _, userExists := uaacUsers[lowerUserEmail]; !userExists {
			lo.G.Info("User", userEmail, "doesn't exist in cloud foundry, so creating user")
			if err = m.UAACMgr.CreateExternalUser(userEmail, userEmail, userEmail, config.Origin); err != nil {
				lo.G.Info("Unable to create user", userEmail)
				return err
			} else {
				uaacUsers[userEmail] = userEmail
			}
		}
		if _, ok := spaceUsers[lowerUserEmail]; !ok {
			if err = m.addUserToOrgAndRole(userEmail, updateUsersInput.OrgGUID, updateUsersInput.SpaceGUID, updateUsersInput.Role, updateUsersInput.OrgName, updateUsersInput.SpaceName); err != nil {
				lo.G.Error(err)
				return err
			}
		} else {
			delete(spaceUsers, lowerUserEmail)
		}
	}
	if updateUsersInput.RemoveUsers {
		lo.G.Debugf("Deleting users for org/space: %s/%s", updateUsersInput.OrgName, updateUsersInput.SpaceName)
		for spaceUser, spaceUserGUID := range spaceUsers {
			lo.G.Infof("removing user: %s from space: %s and role: %s", spaceUser, updateUsersInput.SpaceName, updateUsersInput.Role)
			err = m.cloudController.RemoveCFUser(updateUsersInput.SpaceGUID, SPACES, spaceUserGUID, updateUsersInput.Role)
			if err != nil {
				lo.G.Errorf("Unable to remove user : %s from space %s role in space : %s", spaceUser, updateUsersInput.Role, updateUsersInput.SpaceName)
				lo.G.Errorf("Cloud controller API error: %s", err)
				return err
			}
		}
	} else {
		lo.G.Infof("Not removing users. Set enable-remove-users: true to spaceConfig for org/space: %s/%s", updateUsersInput.OrgName, updateUsersInput.SpaceName)
	}

	lo.G.Debugf("SpaceUsers after: %v", spaceUsers)
	return nil
}

func (m *UserManager) updateLdapUser(config *ldap.Config, spaceGUID, orgGUID string,
	role string, orgName, spaceName string, uaacUsers map[string]string,
	user ldap.User, spaceUsers map[string]string) error {

	userID := user.UserID
	externalID := user.UserDN
	if config.Origin != "ldap" {
		userID = user.Email
		externalID = user.Email
	} else {
		if user.Email == "" {
			user.Email = fmt.Sprintf("%s@user.from.ldap.cf", userID)
		}
	}
	userID = strings.ToLower(userID)

	if _, ok := spaceUsers[userID]; !ok {
		lo.G.Debugf("User[%s] not found in: %v", userID, spaceUsers)
		if _, userExists := uaacUsers[userID]; !userExists {
			lo.G.Info("User", userID, "doesn't exist in cloud foundry, so creating user")
			if err := m.UAACMgr.CreateExternalUser(userID, user.Email, externalID, config.Origin); err != nil {
				lo.G.Info("Unable to create user", userID)
			} else {
				uaacUsers[userID] = userID
				if err := m.addUserToOrgAndRole(userID, orgGUID, spaceGUID, role, orgName, spaceName); err != nil {
					return err
				}
			}
		} else {
			if err := m.addUserToOrgAndRole(userID, orgGUID, spaceGUID, role, orgName, spaceName); err != nil {
				return err
			}
		}
	} else {
		delete(spaceUsers, userID)
	}
	return nil
}

func (m *UserManager) getLdapUsers(config *ldap.Config, groupNames []string, userList []string) ([]ldap.User, error) {
	users := []ldap.User{}
	for _, groupName := range groupNames {
		if groupName != "" {
			lo.G.Info("Finding LDAP user for group:", groupName)
			if groupUsers, err := m.LdapMgr.GetUserIDs(config, groupName); err == nil {
				users = append(users, groupUsers...)
			} else {
				lo.G.Error(err)
				return nil, err
			}
		}
	}
	for _, user := range userList {
		if ldapUser, err := m.LdapMgr.GetUser(config, user); err == nil {
			if ldapUser != nil {
				users = append(users, *ldapUser)
			}
		} else {
			lo.G.Error(err)
			return nil, err
		}
	}
	return users, nil
}

func (m *UserManager) addUserToOrgAndRole(userID, orgGUID, spaceGUID, role, orgName, spaceName string) error {
	if err := m.cloudController.AddUserToOrg(userID, orgGUID); err != nil {
		lo.G.Error(err)
		return err
	}
	lo.G.Infof("Adding user: %s to org/space: %s/%s with role: %s", userID, orgName, spaceName, role)
	if err := m.cloudController.AddUserToSpaceRole(userID, role, spaceGUID); err != nil {
		lo.G.Error(err)
		return err
	}
	return nil
}
