package customer

import (
	"angrymiao-ai/app/handler"
	"angrymiao-ai/app/handler/customer"
	"angrymiao-ai/app/middleware"
	"angrymiao-ai/app/router"
	"angrymiao-ai/config"
	_ "angrymiao-ai/docs/customer"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	framework *router.Framework
}

func NewRouter(f *router.Framework) *Router {
	return &Router{
		framework: f,
	}
}

func (r *Router) Init(engine *gin.Engine, mode string) *gin.Engine {
	engine.Use(middleware.MethodOverrideMiddleware(engine))
	group := r.initRouter(engine, mode)
	engine = r.loadRouter(engine, group)
	engine = r.loadMiddleWare(engine)

	// gzip
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	return engine
}

func (*Router) loadMiddleWare(engine *gin.Engine) *gin.Engine {
	return engine
}

func (r *Router) loadRouter(engine *gin.Engine, group *gin.RouterGroup) *gin.Engine {
	// handler
	h := &handler.Handler{
		Customer: &customer.Handler{},
	}

	// routers
	loadRouter(group, h)

	return engine
}

func loadRouter(group *gin.RouterGroup, handler *handler.Handler) {
	ChatRouter(group, handler)
}

func (*Router) initRouter(engine *gin.Engine, mode string) *gin.RouterGroup {

	// cors
	if mode == config.ReleaseMode {
		engine.Use(middleware.CORSMiddleware())
	}

	baseGroup := engine.Group("api")

	// swagger
	if mode == config.DevMode {
		baseGroup.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return baseGroup
}
