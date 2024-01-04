package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/quwan-sre/observability-go-contrib/test/e2e/metrics/grpc_server"

	metrics "github.com/quwan-sre/observability-go-contrib/metrics/gin"
)

func main() {
	go runGinServer()
	go grpc_server.RunGRPCServer()
}

func runGinServer() {
	r := gin.Default()

	r.Use(metrics.NewMetricsMiddleware())

	r.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
		return
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
		return
	})

	r.GET("/exist", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
		return
	})

	r.Run()
}
