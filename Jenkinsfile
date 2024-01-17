pipeline {
    agent any

    environment {
        GO_VERSION = '1.17'
        GOPATH = "${env.WORKSPACE}/go"
        GOBIN = "${GOPATH}/bin"
        GOROOT = "C:\\go"
        PATH = "C:\\go\\bin;${GOPATH}/bin;${env.PATH}"
    }

   parameters {
        string(name: 'PORT', defaultValue: '4001', description: 'Port number for the Go program')
    }

    stages {
        stage('Install Go') {
            steps {
                script {
                    // Install Go
                    bat "curl -O https://dl.google.com/go/go${GO_VERSION}.windows-amd64.zip"
                    bat "tar -xf go${GO_VERSION}.windows-amd64.zip -C C:\\"

                    // Verify Go installation
                    bat "${GOROOT}\\bin\\go version"
                }
            }
        }


        stage('Build and Run') {
            steps {
                script {
                    // Navigate to the project directory
                    dir("${env.WORKSPACE}") {
                        // Build the Go executable
                        def port = params.PORT
                        bat "${GOROOT}\\bin\\go build -o myapp.exe main.go ${port}"

                        // Run the executable
                        bat "start myapp.exe"
                    }
                }
            }
        }
            stage('Run Tests') {
            steps {
                script {
                   bat "${GOROOT}\\bin\\go test -v ./..."
                }
            }
        }
    }

    post {
        success {
            echo 'Build and tests successful!'
        }
        failure {
            echo 'Build or tests failed!'
        }
    }
}
