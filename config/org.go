package config

import (
	"fmt"
	"path/filepath"
)

// OrgConfig describes configuration for an org.
type OrgConfig struct {
	Org                     string   `yaml:"org"`
	BillingManagerGroup     string   `yaml:"org-billingmanager-group,omitempty"`
	ManagerGroup            string   `yaml:"org-manager-group,omitempty"`
	AuditorGroup            string   `yaml:"org-auditor-group,omitempty"`
	BillingManager          UserMgmt `yaml:"org-billingmanager"`
	Manager                 UserMgmt `yaml:"org-manager"`
	Auditor                 UserMgmt `yaml:"org-auditor"`
	PrivateDomains          []string `yaml:"private-domains"`
	RemovePrivateDomains    bool     `yaml:"enable-remove-private-domains"`
	EnableOrgQuota          bool     `yaml:"enable-org-quota"`
	MemoryLimit             int      `yaml:"memory-limit"`
	InstanceMemoryLimit     int      `yaml:"instance-memory-limit"`
	TotalRoutes             int      `yaml:"total-routes"`
	TotalServices           int      `yaml:"total-services"`
	PaidServicePlansAllowed bool     `yaml:"paid-service-plans-allowed"`
	RemoveUsers             bool     `yaml:"enable-remove-users"`
	TotalPrivateDomains     int      `yaml:"total_private_domains"`
	TotalReservedRoutePorts int      `yaml:"total_reserved_route_ports"`
	TotalServiceKeys        int      `yaml:"total_service_keys"`
	AppInstanceLimit        int      `yaml:"app_instance_limit"`
	DefaultIsoSegment       string   `yaml:"default_isolation_segment"`
}

// Orgs contains cf-mgmt configuration for all orgs.
type Orgs struct {
	Orgs             []string `yaml:"orgs"`
	EnableDeleteOrgs bool     `yaml:"enable-delete-orgs"`
	ProtectedOrgs    []string `yaml:"protected_orgs"`
}

// Contains determines whether an org is present in a list of orgs.
func (o *Orgs) Contains(orgName string) bool {
	for _, org := range o.Orgs {
		if org == orgName {
			return true
		}
	}
	return false
}

func (o *Orgs) GetOrgListFilenameAndPath(configDir string) string {
	return filepath.Join(configDir, "orgs.yml")
}

func (i *OrgConfig) GetOrgConfigFilenameAndPath(configDir, orgName string) string {
	return fmt.Sprintf("%s/%s/orgConfig.yml", configDir, orgName)
}

func (i *OrgConfig) GetOrgConfigFilePath(configDir, orgName string) string {
	return fmt.Sprintf("%s/%s", configDir, orgName)
}

func (o *OrgConfig) GetBillingManagerGroups() []string {
	return o.BillingManager.groups(o.BillingManagerGroup)
}

func (o *OrgConfig) GetManagerGroups() []string {
	return o.Manager.groups(o.ManagerGroup)
}

func (o *OrgConfig) GetAuditorGroups() []string {
	return o.Auditor.groups(o.AuditorGroup)
}
