pipeline {
    agent any
    parameters {
        choice(name: 'TARGETOS', choices: ['linux', 'darwin', 'windows'], description: 'Select target OS')
        choice(name: 'TARGETARCH', choices: ['amd64', 'arm64'], description: 'Select target architecture')
    }
    stages {
        stage('Clone') {
            steps {
                echo "Cloning repository..."
                git branch: 'develop', url: 'https://github.com/Iamfittz/telegram-bot.git'
            }
        }
        stage('Build') {
            steps {
                echo "Building for ${params.TARGETOS}/${params.TARGETARCH}..."
                sh "make image OS=${params.TARGETOS} ARCH=${params.TARGETARCH}"
            }
        }
        stage('Push') {
            steps {
                echo "Pushing image..."
                sh "make push OS=${params.TARGETOS} ARCH=${params.TARGETARCH}"
            }
        }
    }
}