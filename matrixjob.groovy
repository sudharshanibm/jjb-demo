import hudson.model.Hudson;


// Define parameters
def parameters = [
    string(name: 'PARAM1', defaultValue: 'DefaultValue1', description: 'Description for Parameter 1'),
    choice(name: 'PARAM2', choices: ['Option1', 'Option2', 'Option3'], description: 'Description for Parameter 2')
]

// Job DSL script to create a matrix job with parameters
pipelineJob('MatrixJob') {
    description('Matrix Job with Parameters')

    parameters {
        // Add parameters to the job
        parameters.each { param ->
            param()
        }
    }

    definition {
        matrix {
            axes {
                axis {
                    name 'OS'
                    values 'Linux', 'Windows'
                }
                axis {
                    name 'JDK'
                    values 'OpenJDK8', 'OpenJDK11'
                }
            }

            // Define the build steps
            configurations {
                matrixValue('Linux', 'OpenJDK8') {
                    steps {
                        echo "Building on Linux with JDK OpenJDK8"
                        // Add your build steps here
                    }
                }

                matrixValue('Linux', 'OpenJDK11') {
                    steps {
                        echo "Building on Linux with JDK OpenJDK11"
                        // Add your build steps here
                    }
                }

                matrixValue('Windows', 'OpenJDK8') {
                    steps {
                        echo "Building on Windows with JDK OpenJDK8"
                        // Add your build steps here
                    }
                }

                matrixValue('Windows', 'OpenJDK11') {
                    steps {
                        echo "Building on Windows with JDK OpenJDK11"
                        // Add your build steps here
                    }
                }
            }
        }
    }
}