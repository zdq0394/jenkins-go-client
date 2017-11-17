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
	t := template.Must(template.New("escape").Parse(GetCredentialsTemplateFromFile()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, c); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func GetCredentialsTemplateFromFile() string {
	return GetTemplate("../template/credentials.tpl")
}
