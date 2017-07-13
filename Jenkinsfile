podTemplate(
		label: 'jenkins-slave',
		volumes: [
			hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
		],
		containers: [
			containerTemplate(
				name: 'buildEnv',
				image: 'localhost:5000/pingcap/build_env:latest',
				ttyEnabled: true,
				command: 'cat')]){
		catchError {
			node('jenkins-slave') {
				def GITHASH
				def BUILD_URL = "git@github.com:pingcap/tidb-cloud-manager.git"
				env.GOROOT = "/usr/local/go"
				env.GOPATH = "/go"
				env.PATH = "${env.GOROOT}/bin:/bin:${env.PATH}"
				def ROOT = pwd()
				stage('build process') {
					dir("${ROOT}/go/src/github.com/pingcap/tidb-cloud-manager"){
						def WORKSPACE = pwd()
						container('buildEnv') {
							stage('build tidb-cloud-manager binary'){
									git credentialsId: 'k8s', url: "${BUILD_URL}", branch: "master"
									GITHASH = sh(returnStdout: true, script: "cd ${WORKSPACE} && git rev-parse HEAD").trim()
									sh """
									pwd
									cd ${WORKSPACE}
									export GOPATH=${ROOT}/go:$GOPATH
									make
									mkdir -p docker/bin
									cp bin/tidb-cloud-manager docker/bin/tidb-cloud-manager
									"""
							}
							stage('push tidb-cloud-manager images'){
									def tag = "localhost:5000/pingcap/tidb-cloud-manager_k8s:${GITHASH.take(7)}"
									sh """
									pwd
									cd ${WORKSPACE}/docker
									docker build -t ${tag} .
									docker push ${tag}
									"""
							}
						}
					}
				}
			}
			currentBuild.result = "SUCCESS"
		}
	  	stage('Summary') {
			echo("echo summary info #########")
			slackmsg = "[${env.JOB_NAME.replaceAll('%2F','/')}-${env.BUILD_NUMBER}] `${currentBuild.result}`"
			if(currentBuild.result != "SUCCESS"){
				echo(slackmsg + "currentBuild.result")
				slackSend channel: '#iamgroot', color: 'danger', teamDomain: 'pingcap', tokenCredentialId: 'slack-pingcap-token', message: "${slackmsg}"
			} else {
				echo(slackmsg + "currentBuild.result")
				slackSend channel: '#iamgroot', color: 'good', teamDomain: 'pingcap', tokenCredentialId: 'slack-pingcap-token', message: "${slackmsg}"
			}
		}
	}
