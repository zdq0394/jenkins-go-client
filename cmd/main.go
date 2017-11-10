package main

import (
	"fmt"

	"github.com/zdq0394/jenkins-go-client/gojenkins"
)

func main() {

	jenkins := gojenkins.CreateJenkins(nil, "http://123.59.204.155:8080/", "admin", "a8ccd5481d6342b992543321928e1861")
	_, err := jenkins.Init()
	if err != nil {
		panic("Something Went Wrong")
	}

	jobs, err := jenkins.GetAllJobs()

	if err != nil {
		fmt.Println(err)
	}

	if false {
		for _, job := range jobs {
			fmt.Println(job.GetName())
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
	}

	build, err := jenkins.GetBuild("build_docker_image_from_github", 16)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(build.GetConsoleOutput())

}
