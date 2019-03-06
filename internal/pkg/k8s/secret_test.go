package k8s

import (
    "testing"

    "github.com/ishansd94/kube-secrets/internal/pkg/k8s/tests"
)

func TestCreateSecret(t *testing.T) {

    client := tests.NewInClusterTestClient(t)

    h := NewHandler(client)

    err := h.CreateSecret("test", "secret-test", map[string]string{
        "pods": "100",
    })

    if err != nil {
        t.Fatalf("error getting secrets: %v", err)
    }

    err = h.CreateSecret("test", "secret-test", map[string]string{
        "clusters": "100",
    })

    if err == nil {
        t.Fatalf("expected error, got none")
    }
}

func TestAllSecrets(t *testing.T) {

    client := tests.NewInClusterTestClient(t)

    h := NewHandler(client)

    list, err := h.AllSecrets("secret-test")

    if err != nil {
        t.Fatalf("error getting secrets: %v", err)
    }

    if len(list.Items) == 0 {
        t.Errorf("expected more than 1, got %d", len(list.Items))
    }

}

func TestGetSecret(t *testing.T) {

    client := tests.NewInClusterTestClient(t)

    h := NewHandler(client)

    q, err := h.GetSecret("test", "secret-test")
    if err != nil {
        t.Fatalf("error getting secrets: %v", err)
    }

    if q.GetMetadata().GetName() != "test" {
        t.Errorf("expected secret name test : got %v", q.GetMetadata().GetName())
    }

    _, err = h.GetSecret("test", "bar")
    if err == nil {
        t.Fatalf("expected error, got none")
    }

}
