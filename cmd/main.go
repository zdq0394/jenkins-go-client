package main

import (
	"fmt"

	"github.com/zdq0394/jenkins-go-client/gojenkins"
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

func main() {
	jenkinsURL := "http://123.59.204.155:8080/"
	username := "admin"
	password := "a8ccd5481d6342b992543321928e1861"
	jobName := "auto_build"
	jenkins := gojenkins.CreateJenkins(nil, jenkinsURL, username, password)
	_, err := jenkins.Init()
	if err != nil {
		panic("Something Went Wrong")
	}
	fmt.Println(jenkins.Version)
	fmt.Println(jenkins.Server)

	getJob(jenkins, jobName)
}
