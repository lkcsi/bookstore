pipeline {
    agent any
    environment {
        AUTH_SECRET=credentials('bookstore.auth-secret')
        BOOKS_REPOSITORY='SQL'
        BOOKS_DB_HOST='books-db-1'
        BOOKS_DB_PASSWORD=credentials('bookstore.books-db-password')
        BOOKS_DB_PORT=3306
        BOOKS_API_PORT=8081
    }
    stages {
        stage("build") {
            steps {
                echo 'build'
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
    }
}
