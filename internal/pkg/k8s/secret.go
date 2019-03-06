package k8s

import (
	"context"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"

	"github.com/ishansd94/kube-secrets/pkg/log"
)


func (h *Handler) CreateSecret(name , ns string, content map[string]string) error {

	secret := &corev1.Secret{
		Metadata: &metav1.ObjectMeta{
			Name:      k8s.String(name),
			Namespace: k8s.String(ns),
		},
		StringData: content,
	}

	err := h.Client.Create(context.Background(), secret)
	if err != nil {
		log.Error("k8s.CreateSecret", "error creating secret", err)
		return err
	}

	return nil
}


func (h *Handler) AllSecrets(ns string) (*corev1.SecretList, error) {

	var secrets corev1.SecretList

	err := h.Client.List(context.Background(), ns, &secrets)
	if err != nil {
		log.Error("k8s.CreateSecret", "error listing secrets", err)
		return nil, err
	}

	return &secrets, nil
}


func (h *Handler) GetSecret(name, ns string) (*corev1.Secret, error){

	var secret corev1.Secret

	err := h.Client.Get(context.Background(), ns, name, &secret)
	if err != nil {
		log.Error("k8s.CreateSecret", "error listing secret", err)
		return nil, err
	}

	return &secret, nil
}
