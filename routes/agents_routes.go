package routes
import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

type AgentsRoutes struct {
	agentHandler handlers.AgentHandlerInterface
}

func NewAgentRoutes(agentHandler handlers.AgentHandlerInterface) *AgentsRoutes {
	return &AgentsRoutes{agentHandler: agentHandler}
}

func (ar *AgentsRoutes) InitialAgentRoutes(app *fiber.App){
	app.Post("/agents",ar.agentHandler.Save)
	app.Get("/agents",ar.agentHandler.GetAll)
	app.Get("/agents/:id",ar.agentHandler.GetById)
}