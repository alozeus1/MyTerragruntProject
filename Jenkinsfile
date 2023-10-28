pipeline {
    agent any 

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test Scripts') {
            steps {
                script {
                    def files = sh(script: 'ls -l | grep ^- | awk \'{print $9}\'', returnStdout: true).trim().split("\n")

                    for (file in files) {
                        // Checking if the file is an executable script
                        if (sh(script: "file ${file} | grep -q script", returnStatus: true) == 0) {
                            sh "./${file}"
                        }
                    }
                }
            }
        }
    }
}
