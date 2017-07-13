podTemplate(
		label: 'mypod-first-test',
		volumes: [
			hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
		],
		containers: [
			containerTemplate(
				name: 'mycontainer',
				image: 'localhost:5000/pingcap/build_env:latest',
				ttyEnabled: true,
				command: 'cat')]){
		catchError {
			node('mypod-first-test') {
				def githash_centos7 = "test"
				def BUILD_URL = "git@github.com:pingcap/tidb-cloud-manager.git"
				env.GOROOT = "/usr/local/go"
				env.GOPATH = "/go"
				env.PATH = "${env.GOROOT}/bin:/bin:${env.PATH}"
				def ws = pwd()
				sh "echo the root path: ${ws}"
				stage('build process') {
					container('mycontainer') {
						stage('build tidb-cloud-manager binary'){
							dir("${ws}/go/src/github.com/pingcap/tidb-cloud-manager"){
								def current = pwd()
								sh "echo container current path is: ${current}"
								git credentialsId: 'k8s', url: "${BUILD_URL}", branch: "master"
								//githash_centos7 = sh(returnStdout: true, script: "git rev-parse HEAD").trim()
								sh "cd ${current} && export GOPATH=${ws}/go:$GOPATH && pwd && make || sleep 600"
								sh "cd ${current} && pwd && mkdir -p docker/bin && cp bin/tidb-cloud-manager docker/bin/tidb-cloud-manager"
							}
						}
						stage('push tidb-cloud-manager images'){
							dir("${ws}/go/src/github.com/pingcap/tidb-cloud-manager/docker"){
								docker.build("localhost:5000/pingcap/tidb-cloud-manager_k8s:${githash_centos7}", "docker").push()
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
