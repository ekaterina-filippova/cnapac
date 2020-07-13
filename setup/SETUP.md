# SETUP
## Virtual Machine in VirtualBox
Ubuntu LTS (tested with 19.10) https://ubuntu.com/#download
Disk Space - 10 GB, RAM - 5 GB

TIP: Create a shared folder using Guest Additions. Enable bidirectional clipboard and Drag'n'Drop.

## IntelliJ IDEA
```asp
sudo snap install intellij-idea-ultimate --classic
```
Install plugins: Go, Python, Robot Plugin 

## Git
```
sudo apt install git
git --version
```

## Go
```asp
sudo snap install go --classic
go version
```

## Docker
```bash
sudo apt -y update
sudo apt -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
sudo apt remove docker docker-engine docker.io containerd runc
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update
sudo apt -y install docker-ce docker-ce-cli containerd.io
sudo usermod -aG docker $USER
newgrp docker
docker version
```


## Kind & Insecure Docker Registry
Add a local registry 
```bash
cat <<EOF > /etc/docker/daemon.json
{
    "insecure-registries": ["localhost:5000"]
}
EOF
```

Run `./setup/kind-with-registry.sh`

```bash
curl localhost:5000/v2/_catalog
kind version
```
Ref: https://docs.docker.com/registry/deploying/, 
https://docs.docker.com/registry/insecure/,
https://kind.sigs.k8s.io/docs/user/local-registry/

## kubectl
```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.16.8/bin/linux/amd64/kubectl
sudo chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl version
```
Ref: https://v1-16.docs.kubernetes.io/docs/tasks/tools/install-kubectl/

## Helm
```asp
sudo snap install helm --classic
helm version
```

## Jenkins
```asp
docker run -p 8080:8080 -p 50000:50000 jenkins/jenkins:lts
```
Ref: https://github.com/jenkinsci/docker/blob/master/README.md


## Robot
```bash
pip install robotframework
robot --version
```
