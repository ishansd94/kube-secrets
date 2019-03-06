package k8s

import (
    "github.com/ericchiang/k8s"
)

type Handler struct {
    Client *k8s.Client
}

func NewHandler(client *k8s.Client) Handler {
    return Handler{
        Client: client,
    }
}

