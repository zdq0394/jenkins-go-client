package main

import (
	"fmt"
	"strings"

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
	cred.CredID = "jack@github"
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
	branch.BranchName = "test"
	branch.Registry = "reg.qiniu.com"
	branch.RepoNamespace = "zhangdongqi"
	branch.RepoName = "docker_example"
	branch.ImageTagPrefix = "test"
	branch.Override = "True"
	branch.CredentialsId = "zdq0394"
	config := branch.ToConfigString()
	jenkins.CreateJobInFolder(config, "test", "hub", "zdq0394", "docker_example")
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
	//jenkins.RemoveCredentials("auto-test-888")
	//deleteJob(jenkins)

}
