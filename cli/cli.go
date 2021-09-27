package cli

import (
	"AntarJemput-Be-C/app"
	"AntarJemput-Be-C/config"
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/repositories"
	"AntarJemput-Be-C/config/database"
	route "AntarJemput-Be-C/routes"
	"AntarJemput-Be-C/services"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(application *app.Application){
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	// set up connection
	connDB:= database.InitDb()

	//role services
	roleRepository := repositories.NewRoleRepository(connDB)
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	//REGISTER HANDLER TO Routes
	roleRoute := route.NewRoleRoutes(roleHandler)
	roleRoute.InitialRoleRoutes(app)

	//not found Routes
	route.NotFoundRoute(app)
	StartServerWithGracefulShutdown(app, application.Config.AppPort)

}

func StartServerWithGracefulShutdown(app *fiber.App,port string) {
	appPort:= fmt.Sprintf(`:%s`,port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal,1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := app.Listen(appPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}