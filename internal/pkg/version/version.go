package version

import (
	"github.com/gin-gonic/gin"

	"github.com/ishansd94/kube-secrets/internal/pkg/response"
)

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)

func Get(c *gin.Context) {

	response.Custom(c, 200, gin.H{"buildTime": BuildTime, "commit": Commit, "release": Release})
	return
}
