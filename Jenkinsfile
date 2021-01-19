pipeline {
  agent {
    docker {
      image 'golang:1.15.0'
    }

  }
  stages {
    stage('build') {
      steps {
        sh '''
echo $PWD
go get github.com/dghubble/oauth1
go get github.com/dghubble/go-twitter/twitter
go build *.go
'''
      }
    }

  }
}
