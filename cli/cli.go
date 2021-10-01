package cli

import (
	"AntarJemput-Be-C/app"
	"AntarJemput-Be-C/config"
	"AntarJemput-Be-C/config/database"
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/repositories"
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

func (cli *Cli) Run(application *app.Application) {
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	// set up connection
	connDB := database.InitDb()

	//role services
	roleRepository := repositories.NewRoleRepository(connDB)
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	//agent services
	agentRepository := repositories.NewAgentRepository(connDB)
	agentService := services.NewAgentService(agentRepository)
	agentHandler := handlers.NewAgentHandler(agentService)

	//customerService
	customerRepository := repositories.NewCustomerRepository(connDB)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	//ServiceTransaction service
	serviceTransactionRepo := repositories.NewServiceTransRepo(connDB)
	serviceTransactionService := services.NewSTransactionService(serviceTransactionRepo)
	serviceTransacHandler := handlers.NewSTHandler(serviceTransactionService)

	//typeServiceTransaction
	typeServiceRepo := repositories.NewTypeServiceRepo(connDB)
	typeSService := services.NewSTService(typeServiceRepo)
	typeSHandler := handlers.NewTSHandler(typeSService)

	//users services
	authRepository := repositories.NewAuthRepository(connDB)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	//REGISTER HANDLER TO Routes
	roleRoute := route.NewRoleRoutes(roleHandler)
	agentRoute := route.NewAgentRoutes(agentHandler)
	customerRoute := route.NewCustomerRoutes(customerHandler)
	sTRoute := route.NewSTRoutes(serviceTransacHandler)
	tsRoute := route.NewTSRoutes(typeSHandler)
	authRoute := route.NewAuthRoutes(authHandler)

	roleRoute.InitialRoleRoutes(app)
	agentRoute.InitialAgentRoutes(app)
	customerRoute.InitialCustomerRoute(app)
	sTRoute.InitialSTRoute(app)
	tsRoute.InitialTSRoute(app)
	authRoute.InitialAuthRoutes(app)

	//not found Routes
	route.NotFoundRoute(app)
	StartServerWithGracefulShutdown(app, application.Config.AppPort)

}

func StartServerWithGracefulShutdown(app *fiber.App, port string) {
	appPort := fmt.Sprintf(`:%s`, port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
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
