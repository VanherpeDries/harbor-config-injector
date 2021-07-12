package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"

	config "github.com/VanherpeDries/harbor-config-injector/config"
	"gopkg.in/yaml.v2"
)

type YamlFile struct {
	Config config.Config
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

	flag.StringVar(&filename, "f", "./config.yaml", "path of the config file")
	flag.StringVar(&username, "u", "", "username to authenticate to harbor")
	flag.StringVar(&password, "p", "", "password to authenticate to harbor")
	flag.StringVar(&password, "h", "", "hostname of the harbor server")
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
	basicAuthToken := generateAuthToken(username, password)

	fmt.Println("Token: ", basicAuthToken)
	fmt.Printf("Result : %+v\n", yamlConfig)

}
