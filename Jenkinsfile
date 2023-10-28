pipeline {
    agent any
    triggers {
        githubPush() // Enable GitHub webhook trigger for GitSCM polling
    }
    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
        // skipDefaultCheckout(true)
        disableConcurrentBuilds()
        timeout (time: 7, unit: 'MINUTES')
        timestamps()
    }
    environment {
        DB_USERNAME="admin"
        DB_PASSWORD="put-your-passwor-here"
    }
    parameters {
        string(name: 'BRANCH_NAME', defaultValue: 'dev', description: '')
    }
    
    stages {
        stage('Checkout') {
            steps {
                sh """
                    echo "Cloning the repo. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Building') {
            steps {
                sh """
                    echo "Building the code. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Scanning') {
            steps {
                sh """
                    echo "Scanning the code. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Packaging') {
            steps {
                sh """
                    echo "Packaging the code. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Building the image') {
            steps {
                sh """
                    echo "Building the image. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Pushing into Dockerhub') {
            steps {
                sh """
                    echo "Pushing into Dockerhub. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Deploying To Swarm') {
            steps {
                sh """
                    echo "Deploying To Swarm. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('Clean up') {
            steps {
                sh """
                    echo "Clean up. Please wait ....................."
                    sleep 1
                """
            }
        }
        stage('List contents') {
            steps {
                sh """
                    echo "List contents. Please wait ....................."
                    ls -l
                    pwd
                """
            }
        }
        stage('End Job') {
            steps {
                sh """
                    echo "List contents. Please wait ....................."
                    ls -l
                    pwd
                """
            }
        }
    }
    post {
        success {
            echo 'Pipeline succeeded!'
            // You can add more actions for success here
        }
        failure {
            echo 'Pipeline failed!'
            // You can add more actions for failure here
        }
    }
}

