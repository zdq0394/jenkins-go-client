package template

import (
	"bytes"
	"html/template"
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
	t := template.Must(template.New("escape").Parse(getJobBranchTemplate()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, job); err != nil {
		log.Fatal(err)
	}
	return config.String()
}

func (job JobTagConfig) ToConfigString() string {
	t := template.Must(template.New("escape").Parse(getJobTagTemplate()))
	config := bytes.NewBufferString("")
	if err := t.Execute(config, job); err != nil {
		log.Fatal(err)
	}
	return config.String()
}
