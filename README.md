# KUBESCALE - Containerized Web App Scaling on AKS

## Introduction
KUBESCALE is a project focused on containerizing and scaling a web application using Azure Kubernetes Service (AKS). This project utilizes Docker, Jenkins, Kubernetes, Helm, CRI-O, and Azure services to achieve a scalable, high-availability solution capable of handling up to 10,000 concurrent users.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [CI/CD Pipeline](#cicd-pipeline)
- [Kubernetes Cluster Setup](#kubernetes-cluster-setup)
- [Scaling and High Availability](#scaling-and-high-availability)
- [Contributing](#contributing)
- [Demonstration Video](#demonstration-video)

## Features
- Containerization using Docker and Docker-Compose
- Versioning and image storage with DockerHub
- CI/CD pipeline automation with Jenkins and GitHub webhooks
- Automated Kubernetes cluster setup using Kubeadm and Azure AKS CLI
- Integration of CRI-O as the CRI runtime for faster container start-up
- Deployment of Kubernetes manifest files using Helm
- High availability with multi-node cluster setup and Load Balancer
- Improved scalability and reduced downtime using Azure Managed AKS

## Technology Stack
- **Docker**: Containerization
- **Docker-Compose**: Managing multi-container applications
- **Jenkins**: CI/CD pipelines
- **Kubernetes**: Container orchestration
- **Helm**: Kubernetes package management
- **CRI-O**: Container runtime interface
- **Azure Kubernetes Service (AKS)**: Managed Kubernetes service

## Setup and Installation

### Prerequisites
- [Docker](https://www.docker.com/get-started)
- [Docker-Compose](https://docs.docker.com/compose/install/)
- [Jenkins](https://www.jenkins.io/download/)
- [Kubernetes CLI (kubectl)](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)
- [Helm](https://helm.sh/docs/intro/install/)
- [Kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/)

### Installation
1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/your-repo.git
    cd your-repo
    ```

2. **Build and push Docker images**:
    ```sh
    docker-compose build
    docker-compose push
    ```

3. **Setup Kubernetes cluster**:
    ```sh
    # Instructions to set up your Kubernetes cluster using Azure AKS or other providers
    ```

4. **Deploy Helm charts**:
    ```sh
    helm install your-release-name ./helm-chart-directory
    ```

## Usage
- Start the application:
    ```sh
    docker-compose up
    ```
- Access the application as per your deployment setup.

## CI/CD Pipeline
The CI/CD pipeline is automated using Jenkins and GitHub webhooks. The pipeline is triggered by updates in the source code.

### Jenkins Setup
1. **Install Jenkins**: Follow [Jenkins installation guide](https://www.jenkins.io/doc/book/installing/).
2. **Configure GitHub Webhook**:
    - Go to your GitHub repository settings.
    - Add a webhook with the Jenkins URL and `/github-webhook/` endpoint.

3. **Create Jenkins Pipeline**:
    - Create a new pipeline in Jenkins.
    - Configure the pipeline script with the following stages:
        - Build Docker images
        - Push Docker images to DockerHub
        - Deploy to Kubernetes using Helm

## Kubernetes Cluster Setup
- **Automated with Kubeadm and Azure AKS CLI**:
    ```sh
    # Instructions to set up your Kubernetes cluster using Kubeadm and Azure AKS CLI
    ```

## Scaling and High Availability
- **Multi-node cluster setup for high availability**
- **Load Balancer for traffic distribution**
- **Azure Managed AKS for improved scalability**:
    ```sh
    # Instructions for scaling your Azure AKS cluster
    ```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Demonstration Video
Watch a demonstration of this project [here](https://drive.google.com/drive/u/0/folders/1BwiClpP-vdeAmNBedWCPOTsxASKrgGOb).
