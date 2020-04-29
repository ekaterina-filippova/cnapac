#How to use it

##### 1. Modify code

```/app/cmd/http-server.go```

##### 2. Build application

`make build`

##### 3. Create and Push Docker image

 `make image` and `make push`

##### 4. Deploy to Kubernetes

kubectl:
`
kubectl create deployment hello-server-3 --image=localhost:5000/http-server:0.0.1-834b731
`

Helm chart:
`helm install release ./helm-chart/ --set tag="0.0.1-834b731"`

##### 5. Test all the things

##### 6. Use Jenkins Pipelines to automate the process
`/.ci`

###Run application 
* **locally** 
`go run http-server.go`

* **in docker** `docker run --name test-5 -p 8090:8090 -it http-server`

* **in kuberenetes** `kubectl create deployment hello-server-3 --image=localhost:5000/http-server:0.0.1-834b731`

**Test access** `curl localhost:8090/hit`