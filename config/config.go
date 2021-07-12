package config

type Configer interface {
}

type Auth struct {
	Auth_mode string
	// oidc parmeters
	Oidc_verify_cert   bool
	Oidc_scope         string
	Oidc_name          string
	Oidc_client_id     string
	Oidc_endpoint      string
	Oidc_client_secret string
	// ldap parameters
	Ldap_group_search_filter  string
	Ldap_search_dn            string
	Ldap_base_dn              string
	Ldap_filter               string
	Ldap_url                  string
	Ldap_uid                  string
	Ldap_ldap_group_base_dn   string
	Ldap_group_attribute_name string
	Ldap_group_admin_dn       string
	Ldap_scope                int
	Ldap_group_search_scope   int
	// other
	Self_registration bool
}

type Email struct {
	Emial_identity     string
	Verify_remote_cert bool
	Email_ssl          bool
	Email_username     string
	Email_insecure     bool
	Email_port         int
	Email_host         string
	Email_from         string
}
type System struct {
	Storage_per_project          string
	Quota_per_project_enable     bool
	Project_creation_restriction string
	ScanPolicy
	Read_only         bool
	Token_expiration  int
	Count_per_project string
}
type ScanPolicy struct {
	ScanType  string
	Parameter struct {
		daily_time int
	}
}
type Config struct {
	Auth   `yaml:",inline"`
	Email  `yaml:",inline"`
	System `yaml:",inline"`
}
