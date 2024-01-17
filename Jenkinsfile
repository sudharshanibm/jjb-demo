pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                // Checkout the source code from your version control system (e.g., Git)
                checkout scm
            }
        }
        
        stage('Build') {
            steps {
                // Assuming you have Go installed on the Jenkins agent
                script {
                    sh 'go build -o myapp main.go'
                }
            }
        }
        
        stage('Run') {
            steps {
                // Run the generated executable
                script {
                    sh './myapp'
                }
            }
        }
    }
    
    post {
        always {
            // Clean up any artifacts or perform cleanup steps if needed
            cleanWs()
        }
    }
}
