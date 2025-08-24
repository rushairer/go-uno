package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/go-uno/gouno"
)

func RegisterWebRouter(server *gin.Engine) {
	registerWebTestRouter(server)
}

func registerWebTestRouter(server *gin.Engine) {
	testGroup := server.Group("/test")
	{
		testGroup.GET(
			"/alive",
			func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gouno.NewSuccessResponse("pong"))
			},
		)
	}
}
