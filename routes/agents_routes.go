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

func (ar *AgentsRoutes) InitialAgentRoutes(app *fiber.App) {
	app.Post(POST_AGENTS, ar.agentHandler.Save)
	app.Get(GETALL_AGENTS, ar.agentHandler.GetAll)
	app.Get(GETBYID_AGENTS, ar.agentHandler.GetById)
	app.Post(POST_SEARCH_AGENT, ar.agentHandler.SearchAgent)
}
