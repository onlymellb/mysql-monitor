podTemplate(
		label: 'mypod-e2e-operator',
		containers: [
			containerTemplate(
				name: 'k8s-operator-e2e',
				image: 'localhost:5000/pingcap/tidb-operator-e2e:1df176f-2017-07-05_06-20-16',
				ttyEnabled: true,
				command: '/usr/local/bin/e2e.test',
				args: '-ginkgo.v --operator-image=localhost:5000/pingcap/tidb-operator:latest --operator-service-ip=10.233.2.135')]){
		node('mypod-e2e-operator') {
			stage('start stage one') {
				container('k8s-operator-e2e') {
					stage('first') {
						sh 'echo do nothing one '
					}
				}
			}
			stage('start stage two') {
				container('k8s-operator-e2e') {
					stage('second') {
						sh 'echo do nothing two '
					}
				}
			}
		}
}
