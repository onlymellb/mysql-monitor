#!groovy

node {
	def TIDB_OPERATOR_BRANCH = "tennix/multi-controllers"

	fileLoader.withGit('git@github.com:onlymellb/jenkins.git', 'master', 'k8s', '') {
		fileLoader.load('pingcap_tidb_operator_branch.groovy').call(TIDB_OPERATOR_BRANCH)
	}
}
