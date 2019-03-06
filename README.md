# Kube-Secrets
Web service to create kubernetes secrets.

### Pre-requisites

```
1. https://github.com/go-task/task (Makefile alternative.)
2. GO
3. https://github.com/golang/dep (Go dependency manager)
```

### Installation

Clone the repo in your $GOPATH.
This project uses Dep (Golang vendoring tool) https://github.com/golang/dep

```sh
$ cd $GOPATH/src/github.com/ishansd94/kube-secrets
$ task install
$ task run
```
*NOTE: Default port is ```:8000```. Port can be changed by setting ```PORT``` environment variable*

### Build

Build parameters are available in the ```Taskfile.yml```
Change ```USERNAME``` and ```IMAGE``` parameters with your docker hub username and desired image name.

```sh
$ task build
```
If you use separate key for gitlab, change the location of the private key file.
```
SSH_PRIVATE_KEY: $(cat ~/.ssh/id_rsa)
```

### Usage
In order for kube-quotas to work the ```ServiceAccount``` within the pod where it's running should have the necessary RBAC permissions.

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
*NOTE: For local clusters this is not needed.* 

##### Deploy to Kubernetes

Create a seperate ```namespace``` for kube-quotas ex: ```app``` and create a ```deployment```.

```
$ kubectl create ns app
$ kubectl create deployment kube-secrets --image=emzian7/kube-secrets -n app
```

##### Using kube-secrets web service

Get the pod ip using,
```
$ kubectl get pods -n app -o wide
```

##### Payloads
---
##### 1. Creating Secrets  

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
##### 2. Listing Secrets
- **Listing a particular secret**
```
$ curl '<kube-secrets pod ip>:8000/?namespace=<ns>&name=<name>'
```

- **Listing secrets in a particular namespace**
```
$ curl '<kube-secrets pod ip>:8000/?namespace=<ns>'
```
- **Listing all secrets**
```
$ curl '<kube-secrets pod ip>:8000'
```

### Testing

```
$ task test
```