package controller

import (
	"github.com/minipkg/ozzo_routing"
	"github.com/pkg/errors"

	"smtp2api/internal/pkg/apperror"

	"github.com/minipkg/log"
	"github.com/minipkg/ozzo_routing/errorshandler"
	"github.com/minipkg/pagination"

	"smtp2api/internal/domain/task"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type modelController struct {
	Logger  log.ILogger
	Service task.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/tasks/ - список всех моделей
//	GET /api/task/{ID} - детали модели
func RegisterTaskHandlers(r *routing.RouteGroup, service task.IService, logger log.ILogger, authHandler routing.Handler) {
	c := modelController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/tasks", c.list)
	r.Get("/tasksp", c.listp) //	try with pagination
	r.Get(`/task/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c modelController) get(ctx *routing.Context) error {
	id, err := ozzo_routing.ParseUintParam(ctx, "id")
	if err != nil {
		return errorshandler.BadRequest("ID is required to be uint")
	}

	entity, err := c.Service.Get(ctx.Request.Context(), id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("not found")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c modelController) list(ctx *routing.Context) error {

	/*where := c.Service.NewEntity()
	err := ozzo_routing.ParseQueryParamsIntoStruct(ctx, where)
	if err != nil {
		errors.Wrapf(apperror.ErrBadRequest, err.Error())
	}
	cond := selection_condition.SelectionCondition{
		Where: where,
	}*/

	st := c.Service.NewEntity()
	cond, err := ozzo_routing.ParseQueryParams(ctx, st)
	if err != nil {
		errors.Wrapf(apperror.ErrBadRequest, err.Error())
	}

	items, err := c.Service.Query(ctx.Request.Context(), cond)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(items)
}

// list method is for a getting a list of all entities
func (c modelController) listp(ctx *routing.Context) error {

	struc := c.Service.NewEntity()
	cond, err := ozzo_routing.ParseQueryParams(ctx, struc)
	if err != nil {
		errors.Wrapf(apperror.ErrBadRequest, err.Error())
	}

	count, err := c.Service.Count(ctx.Request.Context(), cond)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	pages := pagination.NewFromRequest(ctx.Request, int(count))

	cond.Limit = uint(pages.Limit())
	cond.Offset = uint(pages.Offset())

	ctx.Response.Header().Add("pages", pages.BuildLinkHeader("/modelsp", pages.PerPage))

	items, err := c.Service.Query(ctx.Request.Context(), cond)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	pages.Items = items
	return ctx.Write(pages)
}
