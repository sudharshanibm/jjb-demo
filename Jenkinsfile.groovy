import hudson.model.Hudson;

pipeline {
    agent any

    environment {
        GO_VERSION = '1.18'
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
                  
                        // Build the Go executable
                        def port = params.PORT
                        // bat "${GOROOT}\\bin\\go run main.go ${port}"
                        bat "git clone https://github.com/sudharshan3/GO-lms.git && cd GO-lms"
                        bat "dir"
                        bat "cd ./GO-lms"
                        bat "dir"
                        bat "${GOROOT}\\bin\\go mod tidy"
                        bat "${GOROOT}\\bin\\go build -o myapp.exe main.go"

                        // Run the executable
                        bat "start myapp.exe ${port}"
                    
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
