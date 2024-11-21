package router

import (
	"angrymiao-ai/config"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Framework struct {
	engine     *gin.Engine
	httpServer *http.Server
}

var (
	instance *Framework
	once     sync.Once
)

func GetInstance() *Framework {
	mode := config.Conf.Mode
	once.Do(func() {
		var engine *gin.Engine
		switch mode {
		case config.DevMode:
			gin.SetMode(gin.DebugMode)
			engine = gin.Default()
		case config.TestMode:
			gin.SetMode(gin.TestMode)
			engine = gin.Default()
		case config.ReleaseMode:
			gin.SetMode(gin.ReleaseMode)
			engine = gin.New()
			engine.Use(gin.Recovery())
		default:
			gin.SetMode(gin.DebugMode)
			engine = gin.Default()
		}
		instance = &Framework{
			engine: engine,
		}
	})
	return instance
}

func (f *Framework) Init() error {
	f.httpServer = &http.Server{
		Addr:    config.Conf.App.Host + ":" + config.Conf.App.Port,
		Handler: f.engine,
	}
	return nil
}

func (f *Framework) GetEngine() *gin.Engine {
	return f.engine
}

func (f *Framework) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := f.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	return f.Shutdown()
}

func (f *Framework) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := f.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown error: %v\n", err)
		return err
	}

	log.Println("Server exiting")
	return nil
}
