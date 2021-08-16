package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const projectApiPath string = "/api/v2.0/projects"

type Project struct {
	Project_name  string            `json:"project_name"`
	Cve_allowlist Cve_allowlist     `json:"cve_allowlist"`
	Count_limit   int               `json:"count_limit"`
	Registry_id   *int              `json:"registry_id, omitempty"`
	Storage_limit int               `json:"storage_limit"`
	Metadata      map[string]string `json:"metadata"`
	Public        bool              `json:"Public"`
}
type Cve_allowlist struct {
	Items         []map[string]string `json:"items"`
	Project_id    int                 `json:"project_id"`
	Id            int                 `json:"id"`
	Expires_at    int                 `json:"expires_at"`
	Update_time   string              `json:"update_time"`
	Creation_time string              `json:"creation_time"`
}

func CheckProject(x Project, host string, user string, password string) string {
	client := &http.Client{}

	url := host + projectApiPath + "?project_name=" + x.Project_name
	req, err := http.NewRequest("HEAD", url, nil)
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.StatusCode
	return string(status)
}
func PutProject(x Project, host string, user string, password string) (string, int) {
	client := &http.Client{}

	jsonReq, err := json.Marshal(x)

	fmt.Println("json object: ", bytes.NewBuffer(jsonReq))

	url := host + projectApiPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bodyText, _ := ioutil.ReadAll(resp.Body)
	return string(bodyText), resp.StatusCode
}
