#!groovy

node {
    def TIDB_OPERATOR_BRANCH = "master"

    fileLoader.withGit('git@github.com:onlymellb/jenkins.git', 'master', null, '') {
        fileLoader.load('pingcap_tidb_operator_build.groovy').call(TIDB_OPERATOR_BRANCH)
    }
}
