package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TwiN/go-color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/FadhilAF/perbakin-be/common/env"
	"github.com/FadhilAF/perbakin-be/common/session"
	"github.com/FadhilAF/perbakin-be/common/validation"

	"github.com/FadhilAF/perbakin-be/internal/middleware"
)

type App struct {
	Config   env.Config
	delivery deliveries
	usecase  usecases
	dbPool   *pgxpool.Pool
}

func New(config env.Config, dbPool *pgxpool.Pool) App {
	var app App
	app.Config = config
	app.dbPool = dbPool
	app.initUsecase()
	app.initDelivery()
	return app
}

func (app *App) StartServer() {
	if app.Config.Env == env.EnvProd {
		fmt.Println(
			color.Ize(color.Yellow, color.InBold("\nAPP RUN IN PRODUCTION MODE\n")),
		)
	} else {
		fmt.Println(
			color.Ize(color.Red, color.InBold("\nAPP RUN IN DEVELOPMENT MODE\n")),
		)
	}

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.InitValidation(validator)
	}
	handler := app.createHandlers()
	address := fmt.Sprintf("%s:%s", app.Config.AppHost, app.Config.AppPort)
	log.Printf("Server listening on %v\n", address)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Cannot start server %v\n", err)
		}
	}()

	<-osSignalChan
	err := server.Close()
	if err != nil {
		log.Fatalf("cannot shutdown server %v", err)
	}

	fmt.Println()
	log.Println("Server exiting bye bye :D")
}

func (app *App) createHandlers() http.Handler {
	// // Bagian Handler HTTP
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Accept")

	if app.Config.Env == env.EnvProd {
		corsCfg.AllowAllOrigins = false
		corsCfg.AllowOrigins = app.Config.AllowedOrigins
	} else {
		corsCfg.AllowAllOrigins = true
	}

	router.Use(cors.New(corsCfg))
	router.Use(middleware.LoadSessionMiddleware())

	v1 := router.Group("/api/v1")
	app.handlerV1(v1)

	routes := router.Routes()
	if gin.Mode() == gin.DebugMode {
		fmt.Println()
		for _, v := range routes {
			path := color.InBold(v.Path)
			method := color.InYellow(fmt.Sprintf("%-6s", v.Method))
			fmt.Println(method, path)
		}
		fmt.Println()
	}

	// // Bagian Handler Session
	handler := session.SessionHandler(router, app.dbPool, app.Config)

	return handler
}
