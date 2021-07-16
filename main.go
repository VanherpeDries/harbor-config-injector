package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/VanherpeDries/harbor-config-injector/config"

	"gopkg.in/yaml.v2"
)

type YamlFile struct {
	Config config.Config
}
type PutTest struct {
	Email_host string `json:"email_host"`
	Email_from string `json:"email_from"`
}

func TestGetHttp(hostname string, user string, password string) {
	client := &http.Client{}
	url := hostname + "/api/v2.0/statistics"
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(bodyText))

}

func TestPutHttp(hostname string, user string, password string, config config.Config) {
	client := &http.Client{}
	jsonReq, err := json.Marshal(config)
	fmt.Println("json: ", bytes.NewBuffer(jsonReq))
	url := hostname + "/api/v2.0/configurations"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response : ", string(bodyText))

}

func parseYaml(fileName string) (YamlFile, error) {
	fmt.Println("Parsing YAML files")

	fmt.Println(fileName)
	var yamlConfig YamlFile
	// Opening file from path
	yamlfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return yamlConfig, err

	}
	// Parsing YAML file
	err = yaml.Unmarshal(yamlfile, &yamlConfig)
	if err != nil {
		return yamlConfig, err
	}

	return yamlConfig, nil
}

func generateAuthToken(username string, password string) string {
	token := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return token
}

func main() {
	// parse command option (filename, username, password)
	var filename string
	var username string
	var password string
	var hostname string

	flag.StringVar(&filename, "f", "./config.yaml", "path of the config file")
	flag.StringVar(&username, "u", "", "username to authenticate to harbor")
	flag.StringVar(&password, "p", "", "password to authenticate to harbor")
	flag.StringVar(&hostname, "h", "", "hostname of the harbor server")
	flag.Parse()

	// Check if filename is provided
	if filename == "" {
		fmt.Println("Filename is empty, please provide another by using -p option")
		return
	}
	// Parsing yaml file from filename
	yamlConfig, err := parseYaml(filename)
	if err != nil {
		fmt.Println("Error reading YAML file %s\n", err)
		return
	}

	// generate token from user & password

	fmt.Println("Get config output: ", config.GetConfig(yamlConfig.Config, hostname, username, password))
	fmt.Printf("Result : %+v\n", yamlConfig)

}
