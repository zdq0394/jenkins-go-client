package template

import "fmt"

const templ_prefix = `
<project>
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties>
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.28.1">
      <projectUrl>{{.ProjectURL}}</projectUrl>
      <displayName></displayName>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
    <hudson.model.ParametersDefinitionProperty>
      <parameterDefinitions>
        <hudson.model.StringParameterDefinition>
          <name>REGISTRY</name>
          <description></description>
          <defaultValue>{{.Registry}}</defaultValue>
        </hudson.model.StringParameterDefinition>
        <hudson.model.StringParameterDefinition>
          <name>REPO_NAMESPACE</name>
          <description></description>
          <defaultValue>{{.RepoNamespace}}</defaultValue>
        </hudson.model.StringParameterDefinition>
        <hudson.model.StringParameterDefinition>
          <name>REPO_NAME</name>
          <description></description>
          <defaultValue>{{.RepoName}}</defaultValue>
        </hudson.model.StringParameterDefinition>
        <hudson.model.StringParameterDefinition>
          <name>IMAGE_TAG_PREFIX</name>
          <description></description>
          <defaultValue>{{.ImageTagPrefix}}</defaultValue>
        </hudson.model.StringParameterDefinition>
        <hudson.model.StringParameterDefinition>
          <name>OVERRIDE</name>
          <description></description>
          <defaultValue>{{.Override}}</defaultValue>
        </hudson.model.StringParameterDefinition>
      </parameterDefinitions>
    </hudson.model.ParametersDefinitionProperty>
  </properties>
  <scm class="hudson.plugins.git.GitSCM" plugin="git@3.6.4">
    <configVersion>2</configVersion>
    <userRemoteConfigs>
      <hudson.plugins.git.UserRemoteConfig>
        <refspec>+refs/heads/{{.BranchName}}:refs/remotes/origin/{{.BranchName}}</refspec>
        <url>{{.ProjectURL}}</url>
      </hudson.plugins.git.UserRemoteConfig>
    </userRemoteConfigs>
    <branches>
      <hudson.plugins.git.BranchSpec>
        <name>*/{{.BranchName}}</name>
      </hudson.plugins.git.BranchSpec>
    </branches>
    <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
    <submoduleCfg class="list"/>
    <extensions/>
  </scm>
  <canRoam>true</canRoam>
  <disabled>false</disabled>
  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
  <triggers>
    <com.cloudbees.jenkins.GitHubPushTrigger plugin="github@1.28.1">
      <spec></spec>
    </com.cloudbees.jenkins.GitHubPushTrigger>
  </triggers>
  <concurrentBuild>true</concurrentBuild>
  <builders>
    <hudson.tasks.Shell>
	  <command>COMMIT_ID=`

const templ_mid = "`git rev-parse --short HEAD`"
const templ_last = `
if [ $OVERRIDE = &quot;True&quot; ] ; then
    DOCKER_TAG=$IMAGE_TAG_PREFIX
else
    DOCKER_TAG=$IMAGE_TAG_PREFIX-$COMMIT_ID
fi

docker build -t $REGISTRY/$REPO_NAMESPACE/$REPO_NAME:$DOCKER_TAG -f my_ubuntu/Dockerfile .

</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers/>
  <buildWrappers>
    <hudson.plugins.ws__cleanup.PreBuildCleanup plugin="ws-cleanup@0.34">
      <deleteDirs>false</deleteDirs>
      <cleanupParameter></cleanupParameter>
      <externalDelete></externalDelete>
    </hudson.plugins.ws__cleanup.PreBuildCleanup>
    <hudson.plugins.timestamper.TimestamperBuildWrapper plugin="timestamper@1.8.8"/>
  </buildWrappers>
</project> `

func GetTemplate() string {
	return fmt.Sprintf("%s%s%s", templ_prefix, templ_mid, templ_last)
}
