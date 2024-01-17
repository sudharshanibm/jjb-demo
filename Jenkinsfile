pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                script {
                    // Clone the repository
                    git 'https://github.com/sudharshan3/GO-lms.git'
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    // Build the main.go file
                    sh 'go build main.go'
                }
            }
        }

        stage('Run') {
            steps {
                script {
                    // Run the executable
                    sh './main'
                }
            }
        }
    }

    post {
        success {
            echo 'Build and Run successful!'
        }
    }
}
