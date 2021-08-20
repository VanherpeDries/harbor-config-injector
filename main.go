package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"

	config "github.com/VanherpeDries/harbor-config-injector/config"
	project "github.com/VanherpeDries/harbor-config-injector/projects"

	"gopkg.in/yaml.v2"
)

type YamlFile struct {
	Config   config.Config
	Projects []project.Project
}
type PutTest struct {
	Email_host string `json:"email_host"`
	Email_from string `json:"email_from"`
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
func projectLoop(projects []project.Project, host string, user string, password string) (int, int, int) {
	countChanged, countExists, countError := 0, 0, 0

	for p := range projects {
		txt, status := project.PutProject(projects[p], host, user, password)
		fmt.Printf("\n %+v\n", txt)
		if status == 200 {
			countChanged++
		} else if status == 409 {
			countExists++
		} else {
			countError++
		}
	}
	return countChanged, countExists, countError

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

	fmt.Printf("Get config output: %+v\n", config.PutConfig(yamlConfig.Config, hostname, username, password))

	changed, exists, projecterr := projectLoop(yamlConfig.Projects, hostname, username, password)
	fmt.Printf("\n Inject project: \n Changed: %+v \n Exists: %+v \n Error: %+v \n", changed, exists, projecterr)
	fmt.Printf("Result : %+v\n", yamlConfig)

}
