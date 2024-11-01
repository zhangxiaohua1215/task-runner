package router

import (
	"github.com/gin-gonic/gin"
)

// func GinLogger(logger *zap.Logger) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		path := c.Request.URL.Path
// 		c.Next()

// 		logger.Info(path,
// 			zap.String("method", c.Request.Method),
// 			zap.String("path", path),
// 			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
// 		)
// 	}
// }

func Init() *gin.Engine {
	// r := gin.New()
	// l, err := zap.NewDevelopment()
	// if err!= nil {
	// 	panic(err)
	// }
	// zap.ReplaceGlobals(l)
	// r.Use(GinLogger(zap.L()), gin.Recovery())

	r := gin.Default()
	AppRouterGroup.RegisterScriptRouter(r)
	AppRouterGroup.RegisterTaskRouter(r)
	return r

}
