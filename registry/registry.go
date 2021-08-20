package registry

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const registryApiPath string = "/api/v2.0/registries"

type Registry struct {
	Status        string     `json:"status"`
	Credential    Credential `json:"credential"`
	Update_time   string     `json:"update_time"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	Insecure      string     `json:"insecure"`
	Creation_time string     `json:"creation_time"`
	Type          string     `json:"type"`
	Id            int        `json:"id"`
	Description   string     `json:"description"`
}

type Credential struct {
	Access_key    string `json:"access_key"`
	Access_secret string `json:"access_secret"`
	Type          string `json:"type"`
}

func PostRegistry(x Registry, host string, user string, password string) string {

	client := &http.Client{}
	jsonReq, err := json.Marshal(x)

	url := host + registryApiPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}
func PingRegistry(x Registry, host string, user string, password string) string {
	client := &http.Client{}
	jsonReq, err := json.Marshal(x)
	url := host + registryApiPath + "/ping"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	return string(bodyText)
}
