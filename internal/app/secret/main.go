package secret

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation"
	gouuid "github.com/satori/go.uuid"

	"github.com/ishansd94/kube-secrets/internal/pkg/k8s"
	"github.com/ishansd94/kube-secrets/internal/pkg/response"
	"github.com/ishansd94/kube-secrets/pkg/log"
)

type SecretRequest struct {
	Namespace string            `json:"namespace"`
	Name      string            `json:"name"`
	Content   map[string]string `json:"content"`
}

func (r SecretRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Namespace, validation.Required),
	)
}

func Create(c *gin.Context) {

	var req SecretRequest
	if err := c.BindJSON(&req); err != nil {
		log.Error("secret.Create", "error while binding request", err)
		response.Default(c, http.StatusInternalServerError)
		return
	}

	if err := req.Validate(); err != nil {
		response.Custom(c, http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	if req.Content == nil {
		req.Content = map[string]string{
			"uuid": uuid(),
		}
	}

	if err := k8s.CreateSecret(req.Name, req.Namespace, req.Content); err != nil{
		response.Custom(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Default(c, http.StatusCreated)
}

func uuid() string {
	return gouuid.NewV4().String()
}
