package response

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/imdario/mergo"
)

var (
    StatusOK = gin.H{
        "status": http.StatusOK,
    }

    StatusCreated = gin.H{
        "status":  http.StatusCreated,
        "message": "created",
    }

    StatusBadRequest = gin.H{
        "status":  http.StatusBadRequest,
        "message": "bad request",
    }

    StatusNotFound = gin.H{
        "status":  http.StatusNotFound,
        "message": "item not found",
    }

    StatusConflict = gin.H{
        "status":  http.StatusConflict,
        "message": "a resource with same name exists",
    }

    StatusInternalServerError = gin.H{
        "status":  http.StatusInternalServerError,
        "message": "internal server error",
    }
)

func Default(c *gin.Context, statusCode int) {
    switch statusCode {

    case http.StatusOK:
        c.JSON(http.StatusInternalServerError, StatusOK)

    case http.StatusCreated:
        c.JSON(http.StatusCreated, StatusCreated)

    case http.StatusBadRequest:
        c.JSON(http.StatusBadRequest, StatusBadRequest)

    case http.StatusConflict:
        c.JSON(http.StatusOK, StatusConflict)

    case http.StatusNotFound:
        c.JSON(http.StatusOK, StatusNotFound)

    case http.StatusInternalServerError:
        c.JSON(http.StatusInternalServerError, StatusInternalServerError)

    }
}

func Custom(c *gin.Context, statusCode int, m gin.H) {

    switch statusCode {

    case http.StatusOK:
        _ = mergo.Merge(&m, StatusOK)
        c.JSON(http.StatusOK, m)

    case http.StatusCreated:
        _ = mergo.Merge(&m, StatusCreated)
        c.JSON(http.StatusCreated, m)

    case http.StatusBadRequest:
        _ = mergo.Merge(&m, StatusBadRequest)
        c.JSON(http.StatusBadRequest, m)

    case http.StatusConflict:
        _ = mergo.Merge(&m, StatusConflict)
        c.JSON(http.StatusConflict, m)

    case http.StatusNotFound:
        _ = mergo.Merge(&m, StatusNotFound)
        c.JSON(http.StatusNotFound, m)

    case http.StatusInternalServerError:
        _ = mergo.Map(&m, StatusInternalServerError)
        c.JSON(http.StatusInternalServerError, m)

    }
}
