node('master') {
  stage('Unit Tests') {
    git url: "https://github.com/bobbydeveaux/micro-sub.git"
  }
  stage('Build Bin') {
    sh "go get -v -d ./..."
    sh "CGO_ENABLED=0 GOOS=linux go build -o micro-sub ."
  }
  stage('Build Image') {
    sh "oc start-build micro-sub --from-file=. --follow"
  }
  stage('Deploy') {
    openshiftDeploy depCfg: 'micro-sub', namespace: 'fbac'
    openshiftVerifyDeployment depCfg: 'micro-sub', replicaCount: 1, verifyReplicaCount: true
  }
}