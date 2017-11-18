package template

import (
	"bytes"
	"html/template"
	"log"
)

type CredentialsConfig struct {
	CredID      string
	UserName    string
	Password    string
	Description string
}

func (c CredentialsConfig) ToConfigString() string {
	t := template.Must(template.New("escape").Parse(GetCredentialsTemplate()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, c); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func GetCredentialsTemplate() string {
	return `{
		"": "0",
		"credentials": {
		"scope": "GLOBAL",
		"id": "{{.CredID}}",
		"username": "{{.UserName}}",
		"password": "{{.Password}}",
		"description": "{{.Description}}",
		"$class": "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl"
		}
	}
	`
}
