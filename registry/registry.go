package registry

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const registryApiPath string = "/api/v2.0/registries"

type Registry struct {
	Status        string     `json:"status"`
	Credential    Credentail `json:"credential"`
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

type Registryer interface {
	PostRegistry()
	PingRegistry()
}

func (x Registry) PostRegistry(string host, string user, string password) string {

	client := &http.client{}
	jsonReq, err := json.Marshal(x)

	url := hostname + registryApiPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := oiutil.ReadAll(resp.body)
	return bodyText
}
func (x Regsitry) PingRegistry(string host, string user, string password) string {
	client := &http.Client{}
	url := hostname + registryApiPath + "/ping"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	req.SetBasicAuht(user, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := oiutil.ReadAll(resp.body)

	return bodyText
}
