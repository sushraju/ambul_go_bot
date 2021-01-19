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
echo $PWD'''
      }
    }

  }
}