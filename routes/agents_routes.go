package routes

import (
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/middleware"

	"github.com/gofiber/fiber/v2"
)

type AgentsRoutes struct {
	agentHandler handlers.AgentHandlerInterface
	// tokenHandler handlers.TokenHandlerInterface
}

func NewAgentRoutes(agentHandler handlers.AgentHandlerInterface) *AgentsRoutes {
	return &AgentsRoutes{agentHandler: agentHandler,
		// tokenHandler: tokenHandler,
	}
}

func (ar *AgentsRoutes) InitialAgentRoutes(app *fiber.App) {
	app.Post(POST_AGENTS, middleware.JWTProtected(), ar.agentHandler.Save)
	app.Get(GETALL_AGENTS, middleware.JWTProtected(), ar.agentHandler.GetAll)
	app.Get(GETBYID_AGENTS, middleware.JWTProtected(), ar.agentHandler.GetById)
	app.Post(POST_SEARCH_AGENT, middleware.JWTProtected(), ar.agentHandler.SearchAgent)
}
