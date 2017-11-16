package main

import (
	"fmt"

	"github.com/zdq0394/jenkins-go-client/gojenkins"
	"github.com/zdq0394/jenkins-go-client/template"
)

func getAllJobs(jenkins *gojenkins.Jenkins) {
	jobs, err := jenkins.GetAllJobs()
	if err != nil {
		fmt.Println(err)
	}
	for _, job := range jobs {
		fmt.Println(job.GetName())
	}
}

func trigerBuild(job *gojenkins.Job) {
	params := map[string]string{
		"REPO_NAME": "ubuntu",
		"TAG_NAME":  "b1",
	}
	buildNumber, err := job.InvokeSimple(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", buildNumber)
}

func getJob(jenkins *gojenkins.Jenkins, jobName string) {
	job, err := jenkins.GetJob(jobName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(job.GetName())
	config, err := job.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}

func getBuild(jenkins *gojenkins.Jenkins, jobName string, buildNumber int64) {
	build, err := jenkins.GetBuild(jobName, buildNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(build.GetConsoleOutput())
}

func createJob(jenkins *gojenkins.Jenkins) {
	jobConfig := template.GetJobTagConfig()
	fmt.Println(jobConfig)
	_, err := jenkins.CreateJobInFolder(jobConfig, "tags_release", "hub", "zdq0394", "docker_example")
	if err != nil {
		fmt.Println(err)
	}
}

func createCredentials(jenkins *gojenkins.Jenkins) {
	data := `json={
		"": "0",
		"credentials": {
		  "scope": "GLOBAL",
		  "id": "auto-test2",
		  "username": "zdq0395",
		  "password": "c17551ff8604d4f1eff12accf27559ea5d9823a2",
		  "description": "accesstoken for zdq0395",
		  "$class": "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl"
		}
	  }
	`
	jenkins.CreateCredentials(data)
}
func main() {
	jenkinsURL := "http://123.59.204.155:8080/"
	username := "admin"
	password := "a8ccd5481d6342b992543321928e1861"
	//jobName := "auto_build"
	jenkins := gojenkins.CreateJenkins(nil, jenkinsURL, username, password)
	_, err := jenkins.Init()
	if err != nil {
		panic("Something Went Wrong")
	}
	//getJob(jenkins, "build_private")
	//createJob(jenkins)
	//createCredentials(jenkins)
	creds, err := jenkins.GetAllCredentials()
	for _, cred := range creds {
		fmt.Println(cred.ID)
		fmt.Println(cred.FullName)
	}
}
