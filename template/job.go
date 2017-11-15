package template

import (
	"bytes"
	"html/template"
	"log"
)

type JobConfig struct {
	ProjectURL     string
	BranchName     string
	Registry       string
	RepoNamespace  string
	RepoName       string
	ImageTagPrefix string
	Override       string
}

func GetJobConfig() string {
	t := template.Must(template.New("escape").Parse(GetTemplate()))
	var data JobConfig
	data.ProjectURL = "https://github.com/zdq0394/docker_example"
	data.BranchName = "test"
	data.Registry = "reg.qiniu.com"
	data.RepoNamespace = "zhangdongqi"
	data.RepoName = "docker_example"
	data.ImageTagPrefix = "test"
	data.Override = "True"
	config := bytes.NewBufferString("")
	if err := t.Execute(config, data); err != nil {
		log.Fatal(err)
	}
	return config.String()
}
