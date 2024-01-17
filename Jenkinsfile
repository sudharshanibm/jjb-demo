pipeline {
    agent any

    environment {
        GO_VERSION = '1.17' // Set the desired Go version
        GOPATH = "${env.WORKSPACE}/go"
        GOBIN = "${GOPATH}/bin"
        PATH = "${GOBIN}:${env.PATH}"
    }

    stages {
        stage('Install Go') {
            steps {
                script {
                    // Install Go
                    sh "curl -O https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
                    sh "tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz"

                    // Set Go environment variables
                    sh "export GOROOT=/usr/local/go"
                    sh "export GOPATH=${GOPATH}"
                    sh "export GOBIN=${GOBIN}"
                    sh "export PATH=${GOBIN}:${PATH}"

                    // Verify Go installation
                    sh "go version"
                }
            }
        }

        stage('Build and Run') {
            steps {
                script {
                    // Navigate to the project directory
                    dir("${env.WORKSPACE}") {
                        // Build the Go executable
                        sh "go build -o myapp main.go"

                        // Run the executable
                        sh "./myapp"
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Build and run successful!'
        }
        failure {
            echo 'Build or run failed!'
        }
    }
}
