pipeline {
    agent {
        label 'docker'
    }
    environment {
        GCP_SERVICE_ACCOUNT = credentials('service_account_jenkins')
        KUBE_CONFIG = credentials('config_kube')
        GCP_SSH_KEY = credentials('id_rsa')
    }
    stages {
        stage('Build') {
            steps {
                echo 'build image golang'
                sh 'docker build -t gcr.io/ferrous-module-395010/golang-apps:${BUILD_NUMBER} .'
            }
        }
        stage('Push') {
            steps {
                echo 'push to gcr'
                sh 'cat "$GCP_SERVICE_ACCOUNT" | docker login -u _json_key --password-stdin https://gcr.io'  
                sh 'docker push gcr.io/ferrous-module-395010/golang-apps:${BUILD_NUMBER}'
            }
        }
        stage('Active GCP Account'){
            steps {
                echo 'Active Account'
                sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" dickysetiadi64@34.101.98.183 "whoami"'
                sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" dickysetiadi64@34.101.98.183 ""'
            }
        }
        stage('Deploy') {
            steps {
                echo 'deploy image'
                sh 'helm repo add dicky-charts https://adhithia21.github.io/helm-charts/charts' 
                sh 'helm upgrade --install goapp dicky-charts/application'
            }
        }
    }
}


