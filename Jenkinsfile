pipeline {
    agent any
    environment {
        AUTH_SECRET=credentials('auth-secret')
        BOOKS_REPOSITORY='SQL'
        BOOKS_DB_HOST='books-db-1'
        BOOKS_DB_PASSWORD=credentials('books-db-password')
        BOOKS_DB_PORT=3306
        BOOKS_API_PORT=8081
    }
    stages {
        stage("build") {
            steps {
                echo 'build'
                dir('app') {
                    git branch: 'main', url: 'https://github.com/lkcsi/bookstore.git'
                }
                sh 'cd app'
                script {
                    try {
                        sh 'docker stop $(docker ps -a -q)'
                        sh 'docker rm $(docker ps -a -q)'
                    } catch (err) {
                        echo 'No probs'
                    }
                    echo "${AUTH_ENABLED}"
                    sh 'docker compose build'
                    sh 'docker compose up -d'
                }
            }
        }
        stage("test") {
            environment {
                BOOK_USER='test'
                PASSWORD='password'
                ENDPOINT='http:localhost:8081'
            }
            steps {
                dir('tests') {
                    git branch: 'main', url: 'https://github.com/lkcsi/bookstore-api-test.git'
                }
                sh 'cd tests'
                sh 'pip install requirements.txt'
                sh 'pytest -k ""'
            }
        }
    }
}
