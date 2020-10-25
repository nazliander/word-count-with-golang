#!/bin/bash

MEMORY=$1
CPU_NR=$2
VERSION=$3
IMAGE_NAME=$4

function log() {
    echo "$(date +'%Y-%m-%d %H:%M:%S') $1: $2"
}

echo $(log "INFO" "Starting Minikube...")
echo $(log "INFO" "Memory: ${MEMORY} MB")
echo $(log "INFO" "# of CPUs: ${CPU_NR}")
echo $(log "INFO" "K8s version: ${VERSION}")
echo $(log "INFO" "Let's hope that your Docker Desktop is running  ¯\_(ツ)_/¯")

{
    minikube start --memory $MEMORY --cpus $CPU_NR --kubernetes-version $VERSION
    } || {
    echo $(log "ERROR" "Cannot start the minikube with the given memory and cpu settings.")
    exit 1
}

echo $(log "INFO" "Enabling Ingress addon...")

{
    minikube addons enable ingress
    } || {
    echo $(log "ERROR" "Cannot enable Ingress. Something went wrong.")
    exit 1
}

echo $(log "INFO" "Enabling local Docker images for testing.")
{
    eval $(minikube docker-env)
    } || {
    echo $(log "ERROR" "Cannot enable local images. Something went wrong.")
    exit 1
}

echo $(log "INFO" "Installing mandatory resources for Ingress")

{
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.28.0/deploy/static/mandatory.yaml
    } || {
    echo $(log "ERROR" "Cannot apply the mandatory resources. Perhaps check your connection.")
    exit 1
}

echo $(log "INFO" "Building you favorite image locally")

{
    docker build . -t $IMAGE_NAME
    } || {
    echo $(log "ERROR" "Cannot build your favorite image, I have mixed feelings, perhaps check the Dockerfile in your working directory.")
    exit 1
}

echo $(log "INFO" "Setting up a hostname to the Minkube external IP address")

{
    echo "$(minikube ip) dev-env" | sudo tee -a /etc/hosts
    } || {
    echo $(log "ERROR" "Cannot set up a name for your Minikube external IP address.")
    exit 1
}

echo $(log "INFO" "Last but not least, here I am printing your Minikube IP address in any case ;)")

{
    echo $(minikube ip)
    } || {
    echo $(log "ERROR" "Are you really sure your minkube is running?")
    exit 1
}
