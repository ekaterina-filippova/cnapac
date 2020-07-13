.PHONY: all docker-and-k8s check-versions go-ver

all: go docker-and-k8s

check-versions:
    go version
    git --version
    docker version
    kubectl version
    kind version
    helm version

docker-and-k8s: docker kind-and-docker-registry kubectl kubhelm

go:
    sudo snap install intellij-idea-ultimate --classic
    #GOROOT
    #GOPATH

git:
    sudo apt install git

docker:
    sudo apt -y update
    sudo apt -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
    sudo apt remove docker docker-engine docker.io containerd runc
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    sudo apt update
    sudo apt -y install docker-ce docker-ce-cli containerd.io
    sudo usermod -aG docker $USER
    newgrp docker

kind-and-docker-registry:
    ./kind-with-registry.sh

helm:
    sudo snap install helm --classic