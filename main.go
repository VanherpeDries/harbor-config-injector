package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func parseYaml(fileName string) (*Config, error) {
	fmt.Println("Parsing YAML files")

	// Opening file from path
	yamlfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// Parsing YAML file
	yamlConfig := new(Config)
	err = yaml.Unmarshal(yamlfile, &yamlConfig)
	if err != nil {
		return nil, err
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
	fmt.Printf("Result: %v\n", yamlConfig)

	fmt.Println(yamlConfig["harbor"])
}
