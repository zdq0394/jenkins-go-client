package template

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

type JobConfig struct {
	ProjectURL     string
	BranchName     string
	ProjectTagLike string
	Registry       string
	RepoNamespace  string
	RepoName       string
	ImageTagPrefix string
	Override       string
	CredentialsId  string
}

func GetJobBranchConfig() string {
	t := template.Must(template.New("escape").Parse(GetBranchTemplateFromFile()))
	var data JobConfig
	data.ProjectURL = "https://github.com/zdq0394/docker_example"
	data.BranchName = "develop"
	data.Registry = "reg.qiniu.com"
	data.RepoNamespace = "zhangdongqi"
	data.RepoName = "docker_example"
	data.ImageTagPrefix = "develop"
	data.Override = "True"
	data.CredentialsId = "zdq0394"
	config := bytes.NewBufferString("")
	if err := t.Execute(config, data); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func GetJobTagConfig() string {
	t := template.Must(template.New("escape").Parse(GetJobTemplateFromFile()))
	var data JobConfig
	data.ProjectURL = "https://github.com/zdq0394/docker_example"
	data.Registry = "reg.qiniu.com"
	data.RepoNamespace = "zhangdongqi"
	data.RepoName = "docker_example"
	data.Override = "True"
	data.CredentialsId = "zdq0394"
	data.ProjectTagLike = "release*"
	config := bytes.NewBufferString("")
	if err := t.Execute(config, data); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func GetBranchTemplateFromFile() string {
	return GetTemplate("../template/job.branch.tpl")
}

func GetJobTemplateFromFile() string {
	return GetTemplate("../template/job.tag.tpl")
}

func GetTemplate(tmplPath string) string {
	data, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
