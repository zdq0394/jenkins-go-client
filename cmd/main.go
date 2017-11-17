package main

import (
	"fmt"

	"github.com/zdq0394/jenkins-go-client/template"
	gojenkins "github.com/zdq0394/z/jenkins"
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

func createCredentials(jenkins *gojenkins.Jenkins) {
	// data := `json={
	// 	"": "0",
	// 	"credentials": {
	// 	  "scope": "GLOBAL",
	// 	  "id": "auto-test2",
	// 	  "username": "zdq0395",
	// 	  "password": "c17551ff8604d4f1eff12accf27559ea5d9823a2",
	// 	  "description": "accesstoken for zdq0395",
	// 	  "$class": "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl"
	// 	}
	//   }
	// `
	var cred template.CredentialsConfig
	cred.CredID = "auto-test-888"
	cred.UserName = "auto-test-Tom"
	cred.Password = "auto-test-TomPassword"
	cred.Description = "hello Tom"
	config := cred.ToConfigString()
	fmt.Println(config)
	err := jenkins.CreateCredentials(fmt.Sprintf("json=%s", config))
	if err != nil {
		fmt.Println(err)
	}
}

func getAllCredentials(jenkins *gojenkins.Jenkins) {
	creds, err := jenkins.GetAllCredentials()
	if err != nil {
		fmt.Println(err)
	}
	for _, cred := range creds {
		fmt.Println(cred.ID)
		fmt.Println(cred.FullName)
	}
}

func createBranchJob(jenkins *gojenkins.Jenkins) {
	var branch template.JobBranchConfig
	branch.ProjectURL = "https://github.com/zdq0394/docker_example"
	branch.BranchName = "develop"
	branch.Registry = "reg.qiniu.com"
	branch.RepoNamespace = "zhangdongqi"
	branch.RepoName = "docker_example"
	branch.ImageTagPrefix = "develop"
	branch.Override = "True"
	branch.CredentialsId = "zdq0394"
	config := branch.ToConfigString()
	jenkins.CreateJobInFolder(config, "AA", "hub", "zdq0394", "docker_example")
}

func createTagJob(jenkins *gojenkins.Jenkins) {
	var tag template.JobTagConfig
	tag.ProjectURL = "https://github.com/zdq0394/docker_example"
	tag.Registry = "reg.qiniu.com"
	tag.RepoNamespace = "zhangdongqi"
	tag.RepoName = "docker_example"
	tag.Override = "True"
	tag.CredentialsId = "zdq0394"
	tag.ProjectTagLike = "release*"
	config := tag.ToConfigString()
	jenkins.CreateJobInFolder(config, "BB", "hub", "zdq0394", "docker_example")
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

	//createBranchJob(jenkins)
	//createTagJob(jenkins)
	//createCredentials(jenkins)
	jenkins.RemoveCredentials("auto-test-888")
}
