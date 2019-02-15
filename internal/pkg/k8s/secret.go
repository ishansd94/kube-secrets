package k8s

import (
	"context"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"

	"github.com/ishansd94/kube-secrets/pkg/log"
)


func CreateSecret(name , ns string, content map[string]string) error {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		log.Error("k8s.CreateSecret", "error creating k8 client", err)
		return  err
	}

	secret := &corev1.Secret{
		Metadata: &metav1.ObjectMeta{
			Name:      k8s.String(name),
			Namespace: k8s.String(ns),
		},
		StringData: content,
	}

	err = client.Create(context.Background(), secret)
	if err != nil {
		log.Error("k8s.CreateSecret", "error creating secret", err)
		return err
	}

	return nil
}
