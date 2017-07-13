#!groovy

node {
	def TIDB_CLOUD_MANAGE_BRANCH = "master"

	fileLoader.withGit('git@github.com:onlymellb/jenkins.git', 'master', 'k8s', '') {
		fileLoader.load('pingcap_tidb_cloud_manage_build.groovy').call(TIDB_CLOUD_MANAGE_BRANCH)
	}
}
