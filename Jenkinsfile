pipeline {
    agent any

    environment {
        GO_VERSION = '1.17'
        GOPATH = "${env.WORKSPACE}/go"
        GOBIN = "${GOPATH}/bin"
        PATH = "${GOBIN};${env.PATH}"
    }

    stages {
        stage('Install Go') {
            steps {
                script {
                    // Install Go
                    bat "curl -O https://dl.google.com/go/go${GO_VERSION}.windows-amd64.zip"
                    bat "tar -xf go${GO_VERSION}.windows-amd64.zip -C C:\\"

                    // Set Go environment variables
                    bat "setx GOROOT C:\\go"
                    bat "setx GOPATH ${GOPATH}"
                    bat "setx GOBIN ${GOBIN}"
                    bat "setx PATH \"${GOBIN};${env.PATH}\""

                    // Verify Go installation
                    bat "go version"
                }
            }
        }

        stage('Build and Run') {
            steps {
                script {
                    // Navigate to the project directory
                    dir("${env.WORKSPACE}") {
                        // Build the Go executable
                        bat "go build -o myapp.exe main.go"

                        // Run the executable
                        bat "start myapp.exe"
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
