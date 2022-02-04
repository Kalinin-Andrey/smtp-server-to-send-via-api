package controller

import (
	"smtp2api/internal/domain/email"
	"smtp2api/internal/pkg/apperror"

	"github.com/minipkg/log"
	ozzo_routing "github.com/minipkg/ozzo_routing"
	"github.com/minipkg/ozzo_routing/errorshandler"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type userController struct {
	Logger  log.ILogger
	Service email.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterUserHandlers(r *routing.RouteGroup, service email.IService, logger log.ILogger, authHandler routing.Handler) {
	c := userController{
		Logger:  logger,
		Service: service,
	}

	r.Get(`/user/<id:\d+>`, c.get)
	//r.Get("/users", c.list)

}

// get method is for a getting a one enmtity by ID
func (c userController) get(ctx *routing.Context) error {
	id, err := ozzo_routing.ParseUintParam(ctx, "id")
	if err != nil {
		return errorshandler.BadRequest("ID is required to be uint")
	}

	entity, err := c.Service.Get(ctx.Request.Context(), uint(id), 1)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
/*func (c userController) list(ctx *routing.Context) error {
	rctx := ctx.Request.Context()
	items, err := c.Service.List(rctx)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(items)
}*/
