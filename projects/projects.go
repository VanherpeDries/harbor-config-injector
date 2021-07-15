package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const projectApiPath string = "/api/v2.0/projects"

type Project struct {
	Project_name  string            `json:"project_name"`
	Cve_allowlist Cve_allowlist     `json:"cve_allowlist"`
	Count_limit   int               `json:"count_limit"`
	Registry_id   int               `json:"registry_id"`
	Storage_limit int               `json:"storage_limit"`
	Metadata      map[string]string `json:"metadata"`
	Public        bool              `json:"bool"`
}
type Cve_allowlist struct {
	Items         []map[string]string `json:"items"`
	Project_id    int                 `json:"project_id"`
	Id            int                 `json:"id"`
	Expires_at    int                 `json:"expires_at"`
	Update_time   string              `json:"update_time"`
	Creation_time string              `json:"creation_time"`
}

type Projecter interface {
	CheckProject()
	PutProject()
}

func (x Project) CheckProject(string host, string user, string password) string {
	client := &http.client{}

	url := hostname + projectApiPath + "?project_name=" + x.Project_name
	req, err := http.NewRequest("HEAD", url, nil)
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := oiutil.ReadAll(resp.body)
	return bodyText
}
func (x Project) PutProject(string host, string user, string password) string {
	client := &http.client{}

	jsonReq, err := json.Marshal(x)

	fmt.Println("json object: ", bytes.NewBuffer(jsonReq))

	url := hostname + projectApiPath
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))

	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := oiutil.ReadAll(resp.body)

	return bodyText
}
