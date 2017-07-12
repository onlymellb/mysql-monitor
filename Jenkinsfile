podTemplate(
		label: 'mypod-first-test',
		containers: [
			containerTemplate(
				name: 'mycontainer',
				image: 'busybox:1.26.2',
				ttyEnabled: true,
				command: 'cat')]){
		node('mypod-first-test') {
			stage('start stage one') {
				container('mycontainer') {
					stage('first') {
						sh 'sleep 300 && echo do nothing one '
					}
				}
			}
			stage('start stage two') {
				container('mycontainer') {
					stage('second') {
						sh 'echo do nothing two '
					}
				}
			}
		}
}
