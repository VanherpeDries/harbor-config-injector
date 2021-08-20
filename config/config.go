package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const configApiPath string = "/api/v2.0/configurations"

type Auth struct {
	Auth_mode string `json:"auth_mode"`
	// oidc parmeters
	Oidc_verify_cert          bool   `json:"oidc_verify_cert"`
	Oidc_scope                string `json:"oidc_scope"`
	Oidc_name                 string `json:"oidc_name"`
	Oidc_client_id            string `json:"oidc_client_id"`
	Oidc_endpoint             string `json:"oidc_endpoint"`
	Oidc_client_secret        string `json:"oidc_client_secret"`
	Oidc_extra_redirect_parms string `json:"oidc_extra_redirect_parms"`
	Oidc_auto_onboard         string `json:"oidc_auto_onboard"`
	Oidc_admin_group          string `json:"oidc_admin_group"`
	Oidc_user_claim           string `json:"oidc_user_claim"`
	Oidc_groups_claim         string `json:"oidc_groups_claim"`
	// ldap parameters
	Ldap_group_search_filter  string `json:"ldap_group_search_filter"`
	Ldap_search_password      string `json:"ldap_search_password"`
	Ldap_search_dn            string `json:"ldap_search_dn"`
	Ldap_base_dn              string `json:"ldap_base_dn"`
	Ldap_filter               string `json:"ldap_filter"`
	Ldap_url                  string `json:"ldap_url"`
	Ldap_uid                  string `json:"ldap_uid"`
	Ldap_group_base_dn        string `json:"ldap_group_base_dn"`
	Ldap_group_attribute_name string `json:"ldap_group_attribute_name"`
	Ldap_group_admin_dn       string `json:"ldap_group_admin_dn"`
	Ldap_scope                int    `json:"ldap_scope"`
	Ldap_group_search_scope   int    `json:"ldap_group_search_scope"`
	// other
	Self_registration          bool   `json:"self_registration"`
	Http_authproxy_verify_cert bool   `json:"http_authproxy_verify_cert"`
	Uaa_client_id              string `json:"uaa_client_id"`
}

type Email struct {
	Email_identity     string `json:"email_identity"`
	Verify_remote_cert bool   `json:"verify_remote_cert"`
	Email_ssl          bool   `json:"email_ssl"`
	Email_username     string `json:"email_username"`
	Email_insecure     bool   `json:"email_insecure"`
	Email_password     string `json:"email_password"`
	Email_port         int    `json:"email_port"`
	Email_host         string `json:"email_host"`
	Email_from         string `json:"email_from"`
}
type System struct {
	Storage_per_project          string `json:"storage_per_project"`
	Quota_per_project_enable     bool   `json:"quota_per_project_enable"`
	Project_creation_restriction string `json:"project_creation_restriction"`
	Read_only                    bool   `json:"read_only"`
	Token_expiration             int    `json:"token_expirtation"`
	Count_per_project            string `json:"count_per_project"`
	Robot_name_prefix            string `json:"robot_name_prefix"`
}
type Config struct {
	Auth   `yaml:",inline"`
	Email  `yaml:",inline"`
	System `yaml:",inline"`
}

func PutConfig(x Config, host string, user string, password string) string {

	fmt.Println("Parsing config to json ...")
	// Creating http client
	client := &http.Client{}
	// config => json
	jsonReq, err := json.Marshal(x)
	fmt.Println("json object: ", bytes.NewBuffer(jsonReq))
	// setting Hostname
	url := host + configApiPath

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}

func GetConfig(x Config, host string, user string, password string) string {
	// Creating http client
	client := &http.Client{}
	// setting Hostname
	url := host + configApiPath

	fmt.Println("url: " + url)
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println("req: " + string(req.Method))

	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}
