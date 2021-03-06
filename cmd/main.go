package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
	if build.IsRunning() {
		fmt.Println("IS RUNNING")
	} else {
		fmt.Println(build.GetActions())
		fmt.Println(build.GetDuration())
		fmt.Println(build.GetRevision())
		fmt.Println(build.GetRevisionBranch())
		fmt.Println(build.GetTimestamp())
		fmt.Println(build.GetBuildNumber())
		fmt.Println(build.GetResult())
	}

}

func createCredentials(jenkins *gojenkins.Jenkins) {
	var cred template.CredentialsConfig
	cred.CredID = "zdq0394@github"
	cred.UserName = "zdq0394"
	cred.Password = "243e082c37a843418ff5841c3753d20ea3e92402"
	cred.Description = "zdq0394 access token of github"
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
	branch.BranchName = "test"
	branch.Registry = "reg.qiniu.com"
	branch.RepoNamespace = "zhangdongqi"
	branch.RepoName = "docker_example"
	branch.ImageTagPrefix = "test"
	branch.Override = "True"
	branch.CredentialsId = "zdq0394@github"
	config := branch.ToConfigString()
	var i int
	for i = 101; i <= 180; i++ {
		t_i := i
		go func() {
			jobName := fmt.Sprintf("loop%d", t_i)
			fmt.Println(jobName)
			_, err := jenkins.CreateJobInFolder(config, jobName, "hub", "zdq0394", "docker_example")
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	select {}

}

func updateBranchJob(jenkins *gojenkins.Jenkins) {
	var branch template.JobBranchConfig
	branch.ProjectURL = "https://github.com/zdq0394/docker_example"
	branch.BranchName = "test"
	branch.Registry = "reg.qiniu.com"
	branch.RepoNamespace = "zhangdongqi"
	branch.RepoName = "docker_example"
	branch.ImageTagPrefix = "Hahaha"
	branch.Override = "True"
	branch.CredentialsId = "zdq0394@github"
	config := branch.ToConfigString()
	job, err := jenkins.GetJob("test", "hub", "zdq0394", "docker_example")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(job.Base)
	job.UpdateConfig(config)
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
	jenkins.CreateJobInFolder(config, "tag_prefixed_release", "hub", "zdq0394", "docker_example")
}

func deleteJob(jenkins *gojenkins.Jenkins) {
	job, err := jenkins.GetJob("test", "hub", "zdq0394", "docker_example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.GetName())

	ok, err := job.Delete()
	if ok {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK")
	}
	if err != nil {
		fmt.Println(err)
	}
}

func getDockerTagFromConsoleOutput(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fmt.Println(line)
		if strings.HasPrefix(line, "+ DOCKER_TAG=") {
			return strings.TrimRight(line[13:], "\n")
		}
	}
	return ""
}

func getJobBuildLog(jenkins *gojenkins.Jenkins) {
	job, err := jenkins.GetJob("5a265aa93b5a640001000001", "hub", "wzjgo", "2048")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.Base)
	fmt.Println(job.GetBuildConsoleOutputWithTimestamp(1))
}

func getJobAllBuildIDS(jenkins *gojenkins.Jenkins) {
	job, err := jenkins.GetJob("5a265aa93b5a640001000001", "hub", "wzjgo", "2048")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.Base)
	fmt.Println(job.GetBuildConsoleOutputWithTimestamp(1))

	b, err := job.GetBuild(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%v", b)
	bIDs, _ := job.GetAllBuildInfos()
	for _, bid := range bIDs {
		fmt.Println(bid.Number)
		fmt.Println(bid.URL)
		fmt.Println(bid.Duration)
		fmt.Println(bid.Timestamp)
		fmt.Println(bid.Result)
	}
}

func getDockerTag(jenkins *gojenkins.Jenkins) {
	job, err := jenkins.GetJob("master", "hub", "zdq0394", "docker_example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.Base)
	b, err := job.GetBuild(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%v", b)
	output := b.GetConsoleOutput()
	dockerTag := getDockerTagFromConsoleOutput(output)
	fmt.Println(dockerTag)
}

func main() {
	jenkinsURL := "http://123.59.204.159"
	username := "admin"
	password := "Qwer1234"
	//jobName := "auto_build"
	client := http.DefaultClient
	client.Timeout = 600 * time.Second
	jenkins := gojenkins.CreateJenkins(client, jenkinsURL, username, password)
	_, err := jenkins.Init()
	if err != nil {
		fmt.Println(err)
	}

	createBranchJob(jenkins)
	//updateBranchJob(jenkins)
	//createTagJob(jenkins)
	//createCredentials(jenkins)
	//jenkins.RemoveCredentials("auto-test-888")
	//deleteJob(jenkins)

	//getJob(jenkins, "auto_build")
	//getAllCredentials(jenkins)

	//getJobAllBuildIDS(jenkins)
	//getJobBuildLog(jenkins)

}
