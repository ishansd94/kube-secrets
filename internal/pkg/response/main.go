package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

var (
	InternalServerErrorMsg = gin.H{
		"status":  http.StatusInternalServerError,
		"message": "internal server error",
	}

	BadRequestMsg = gin.H{
		"status":  http.StatusBadRequest,
		"message": "bad request",
	}

	CreateMsg = gin.H{
		"status":  http.StatusCreated,
		"message": "created",
	}
)

func Default(c *gin.Context, statusCode int) {
	switch statusCode {

	case http.StatusInternalServerError:
		c.JSON(http.StatusInternalServerError, InternalServerErrorMsg)

	case http.StatusBadRequest:
		c.JSON(http.StatusBadRequest, BadRequestMsg)

	case http.StatusCreated:
		c.JSON(http.StatusCreated, CreateMsg)
	}
}

func Custom(c *gin.Context, statusCode int, m gin.H) {

	switch statusCode {

	case http.StatusInternalServerError:
		_ = mergo.Map(&m, InternalServerErrorMsg)
		c.JSON(http.StatusInternalServerError, m)

	case http.StatusBadRequest:
		_ = mergo.Merge(&m, BadRequestMsg)
		c.JSON(http.StatusBadRequest, m)

	case http.StatusCreated:
		_ = mergo.Merge(&m, CreateMsg)
		c.JSON(http.StatusCreated, CreateMsg)

	}
}
