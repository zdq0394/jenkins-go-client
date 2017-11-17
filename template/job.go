package template

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

type JobConfig struct {
	Registry      string
	RepoNamespace string
	RepoName      string
	Override      string
	CredentialsId string
	ProjectURL    string
}

type JobBranchConfig struct {
	JobConfig
	BranchName     string
	ImageTagPrefix string
}

type JobTagConfig struct {
	JobConfig
	ProjectTagLike string
}

func (job JobBranchConfig) ToConfigString() string {
	t := template.Must(template.New("escape").Parse(GetBranchTemplateFromFile()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, job); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func (job JobTagConfig) ToConfigString() string {
	t := template.Must(template.New("escape").Parse(GetTagTemplateFromFile()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, job); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func GetBranchTemplateFromFile() string {
	return GetTemplate("../template/job.branch.tpl")
}

func GetTagTemplateFromFile() string {
	return GetTemplate("../template/job.tag.tpl")
}

func GetTemplate(tmplPath string) string {
	data, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
