# Kube-Secrets
### Installation

Clone the repo in your $GOPATH
This project uses Dep (Golang vendoring tool) https://github.com/golang/dep

```sh
$ cd $GOPATH/src/github.com/ishansd94/kube-secrets
$ dep ensure
$ go run cmd/secret/main.go
```

Defualt port is ```:8000```

### Build

Buidld script is available under ```hack``` folder.
Change ```USERNAME``` and ```IMAGE``` fields in ```hack\build.sh``` with your docker hub username and desired image name.

```sh
$ sh hack/build.sh
```

### Usage
In order kube-secrets to work the ```ServiceAccount``` within the pod where it's running should have the necessary rbac permissions.

```
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
    name: kube-secrets
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kube-secrets
subjects:
  - kind: Group
    name: system:serviceaccounts
    apiGroup: ""
roleRef:
  kind: ClusterRole
  name: kube-secrets
  apiGroup: rbac.authorization.k8s.io
```
*NOTE: For local clusters this is not needed* 

##### Deploy to Kubernetes

create a seperate ns for kube-secrets ex: app

```
$ kubectl create ns app
$ kubectl create deployment kube-secrets --image=emzian7/kube-secrets -n app
$ kubectl get all -n app
```

##### Using kube-secrets web service

create a nginx or other pod inside the same ns as kube-secrets
```
$ kubectl create deployment nginx --image=nginx -n app
```
Get pod ips
```
$ kubectl get pods -n app -o wide
```
log into nginx or any other pod and use the pod ip of kube-secrets to call the web service.

##### Payloads

Expected payload as a ```POST``` request.
```
{
    "name": <name of the secret, string>,
    "namespace": <kubernetes namespace, string>
    "content": <content of the secret, json obj, optional>
}
```
*NOTE: content is mapped to map[string]string, json obj expected is something like {"foo": "bar"}. If content field is not specified default uuid will be created*. 

```
$ kubectl exec -it -n <any other ns> <any other pod> -- bash
$ curl -d '{"name":"foo", "namespace":"app"}' -H "Content-Type: application/json" -X POST <kube-secrets pod ip>:8000
```

```
$ kubectl get secret -n app foo -o yaml
apiVersion: v1
data:
  uuid: ZThmODhmZDgtNmMzMy00ODM5LThhMzItYzMzMDcxNWYyMzdk
kind: Secret
metadata:
  creationTimestamp: "2019-02-15T06:42:58Z"
  name: foo
  namespace: app
  resourceVersion: "17299"
  selfLink: /api/v1/namespaces/app/secrets/foo
  uid: ee596159-30ec-11e9-a4a2-00155d8a3211
type: Opaque
```

### Todo
- [ ] Jinja templating for Dockerfile to fill in placeholders and create a app specific Dockerfile.
