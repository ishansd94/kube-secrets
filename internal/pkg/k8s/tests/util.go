package tests

import (
    "encoding/json"
    "os/exec"
    "testing"

    k8slib "github.com/ericchiang/k8s"
)

func NewOuterClusterTestClient(t *testing.T) *k8slib.Client {

	cmd := exec.Command("kubectl", "config", "view", "--raw", "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("'kubectl config view -o json': %v %s", err, out)
	}

	config := new(k8slib.Config)
	if err := json.Unmarshal(out, config); err != nil {
		t.Fatalf("parse kubeconfig: %v '%s'", err, out)
	}
	client, err := k8slib.NewClient(config)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	return client
}

func NewInClusterTestClient(t *testing.T) *k8slib.Client{

	client, err := k8slib.NewInClusterClient()
	if err != nil {
		t.Fatalf("error creating client %v", err)
	}

	return client
}
