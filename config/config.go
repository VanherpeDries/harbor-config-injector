package config

type Configer interface {
}

type Auth struct {
	Auth_mode string
	// oidc parmeters
	oidc_verify_cert   bool
	oidc_scope         string
	oidc_name          string
	oidc_client_id     string
	oidc_endpoint      string
	oidc_client_secret string
	// ldap parameters
	ldap_group_search_filter  string
	ldap_search_dn            string
	ldap_base_dn              string
	ldap_filter               string
	ldap_url                  string
	ldap_uid                  string
	ldap_ldap_group_base_dn   string
	ldap_group_attribute_name string
	ldap_group_admin_dn       string
	ldap_scope                int
	ldap_group_search_scope   int
	// other
	self_registration bool
}

type Email struct {
	emial_identity     string
	verify_remote_cert bool
	email_ssl          bool
	email_username     string
	email_insecure     bool
	email_port         int
	email_host         string
	email_from         string
}
type System struct {
	storage_per_project          string
	quota_per_project_enable     bool
	project_creation_restriction string
	ScanPolicy
	read_only         bool
	token_expiration  int
	count_per_project string
}
type ScanPolicy struct {
	scanType  string
	parameter struct {
		daily_time int
	}
}
type Config struct {
	Auth   `yaml:",inline"`
	Email  `yaml:",inline"`
	System `yaml:",inline"`
}
