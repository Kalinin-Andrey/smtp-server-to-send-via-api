package controller

import (
	"context"
	"strconv"
	"time"

	"smtp2api/internal/domain/offer"

	"smtp2api/internal/domain/advertiser"
	"smtp2api/internal/domain/advertising_campaign"

	"smtp2api/internal/domain/email"
	"smtp2api/internal/pkg/config"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/minipkg/log"
	"smtp2api/pkg/yarus_platform"
)

type dataTestController struct {
	Logger              log.ILogger
	user                email.IService
	advertiser          advertiser.IService
	advertisingCampaign advertising_campaign.IService
	offer               offer.IService
	yaruzPlatform       yarus_platform.IPlatform
}

var langEngID = uint(1)
var langRusID = uint(2)
var propertyUnitMMSysname = "mm"
var entityTypeBloggerSysname = "blogger"
var entityTypePostSysname = "post"
var entityTypeStorySysname = "story"
var relationBlogger2StorySysname = "blogger2story"
var relationBlogger2PostSysname = "blogger2post"
var propertyNumberSysname = "num"
var propertyLenSysname = "len"
var propertyBoolSysname = "bool"
var propertyDateSysname = "date"
var propertyTSSysname = "time"
var propertyNameSysname = "name"
var propertyOptSysname = "opt"
var propertyOpt0ID = 0
var propertyOpt0Key = "one"
var propertyOpt0Val = 1
var propertyOpt1ID = 1
var propertyOpt1Key = "two"
var propertyOpt1Val = 2
var propertyOpt2ID = 2
var propertyOpt2Key = "three"
var propertyOpt2Val = 3

var entityTypeID uint

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/models/ - список всех моделей
//	GET /api/model/{ID} - детали модели
func RegisterDataTestHandlers(r *routing.RouteGroup, yaruzPlatform yarus_platform.IPlatform, user email.IService, advertiser advertiser.IService, advertisingCampaign advertising_campaign.IService, offer offer.IService, logger log.ILogger, authHandler routing.Handler) {
	c := dataTestController{
		Logger:              logger,
		yaruzPlatform:       yaruzPlatform,
		user:                user,
		advertiser:          advertiser,
		advertisingCampaign: advertisingCampaign,
		offer:               offer,
	}
	ctx := context.Background()
	c.init(ctx)

	//r.Get("/entity", c.entity)
	//r.Get("/entity-text", c.entityText)
	//r.Get("/entity-relation", c.entityRelation)
	//r.Get("/user", c.userInst)
	//r.Get("/user-search", c.userSearch)
	r.Get("/sharding", c.sharding)
	//r.Get("/entity-properties-search", c.entityPropertiesSearch)
}

func (c dataTestController) init(ctx context.Context) (err error) {
	entityTypeID, err = c.yaruzPlatform.ReferenceSubsystem().EntityType.GetIDBySysname(ctx, email.EntityType)
	if err != nil {
		return err
	}

	return nil
}

/*
func (c dataTestController) entity(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity"})
	cntx := ctx.Request.Context()

	if err := c.propertyUnitsInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertyUnitsInit": err.Error()})
	}

	if err := c.propertiesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertiesInit": err.Error()})
	}

	if err := c.entityTypesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"entityTypesInit": err.Error()})
	}

	propertyNumber, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyNumberSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyNumber": err.Error()})
	}

	propertyLen, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyLenSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyLen": err.Error()})
	}

	propertyOpt, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyOptSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyOpt": err.Error()})
	}

	propertyBool, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyBoolSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyBool": err.Error()})
	}

	propertyDate, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyDateSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyDate": err.Error()})
	}

	propertyTime, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyTSSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyTS": err.Error()})
	}

	propertyName, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyNameSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyName": err.Error()})
	}

	entityType, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TFirst(cntx, &entity_type.EntityType{Sysname: entityTypeBloggerSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for entityType": err.Error()})
	}

	entity := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
	entity.EntityTypeID = entityType.ID
	//entity.PropertiesValuesMap = map[uint]interface{}{
	//	propertyNumber.ID: int(158),
	//	propertyLen.ID:    float64(32.543),
	//	propertyOpt.ID:    propertyOpt1Val,
	//	propertyBool.ID:   true,
	//	propertyDate.ID:   "2021-09-22",
	//	propertyTime.ID:     time.Now().Unix(),
	//}

	err = entity.SetValueForProperty(propertyNumber, int(158), langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyNumber": err.Error()})
	}

	err = entity.SetValueForProperty(propertyLen, float64(32.543), langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyLen": err.Error()})
	}

	err = entity.SetValueForProperty(propertyOpt, propertyOpt1Val, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyOpt": err.Error()})
	}

	err = entity.SetValueForProperty(propertyBool, true, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	err = entity.SetValueForProperty(propertyDate, time.Now(), langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyDate": err.Error()})
	}

	err = entity.SetValueForProperty(propertyTime, time.Now(), langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyTime": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, entity, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Create()": err.Error()})
	}

	entity2, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, entity.ID, entityTypeID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Get()": err.Error()})
	}
	res = append(res, map[string]interface{}{"entity2": entity2})

	err = entity2.SetValueForProperty(propertyLen, float64(64.789), langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyLen": err.Error()})
	}

	err = entity2.SetValueForProperty(propertyName, "Андрей", langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyName": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, entity2, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity2.Create()": err.Error()})
	}

	entity3, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, entity.ID, entityTypeID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Get()": err.Error()})
	}
	res = append(res, map[string]interface{}{"entity3": entity3})

	return ctx.Write(res)
}

func (c dataTestController) entityText(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity-text"})
	cntx := ctx.Request.Context()

	if err := c.propertyUnitsInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertyUnitsInit": err.Error()})
	}

	if err := c.propertiesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertiesInit": err.Error()})
	}

	if err := c.entityTypesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"entityTypesInit": err.Error()})
	}

	propertyName, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyNameSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyName": err.Error()})
	}

	entityType, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TFirst(cntx, &entity_type.EntityType{Sysname: entityTypeBloggerSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for entityType": err.Error()})
	}

	entity := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
	entity.EntityTypeID = entityType.ID

	err = entity.SetValueForProperty(propertyName, "Андрей", langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, entity, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Create()": err.Error()})
	}

	err = entity.SetValueForProperty(propertyName, "Андрей Калинин", langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, entity, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Update()": err.Error()})
	}

	entity1, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, entity.ID, entityTypeID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity1.Get()": err.Error()})
	}
	res = append(res, map[string]interface{}{"entity1": entity1})

	err = entity1.SetValueForProperty(propertyName, "Андрей Калинин 1", langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, entity1, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity.Update()": err.Error()})
	}

	entity2, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, entity.ID, entityTypeID, langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity2.Get()": err.Error()})
	}
	res = append(res, map[string]interface{}{"entity2": entity2})

	err = entity2.SetValueForProperty(propertyName, "Andrey", langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	entity2.DeletePropertyValues(propertyName.ID)

	err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, entity2, langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity2.Update()": err.Error()})
	}

	err = entity2.SetValueForProperty(propertyName, "Andrey Kalinin", langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"Set propertyBool": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, entity2, langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity2.Update()": err.Error()})
	}

	entity3, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, entity.ID, entityTypeID, langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"entity3.Get()": err.Error()})
	}
	res = append(res, map[string]interface{}{"entity3": entity3})

	return ctx.Write(res)
}

func (c dataTestController) entityRelation(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity-relation"})
	cntx := ctx.Request.Context()

	if err := c.propertyUnitsInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertyUnitsInit": err.Error()})
	}

	if err := c.propertiesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"propertiesInit": err.Error()})
	}

	if err := c.entityTypesInit(cntx); err != nil {
		res = append(res, map[string]interface{}{"entityTypesInit": err.Error()})
	}

	rel2Post, err := c.yaruzPlatform.ReferenceSubsystem().Relation.TFirst(cntx, &entity_type.Relation{Property: property.Property{Sysname: relationBlogger2PostSysname}}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for propertyName": err.Error()})
	}

	entityTypeBlogger, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TFirst(cntx, &entity_type.EntityType{Sysname: entityTypeBloggerSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for entityTypeBlogger": err.Error()})
	}

	entityTypePost, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TFirst(cntx, &entity_type.EntityType{Sysname: entityTypePostSysname}, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"finding for entityTypePost": err.Error()})
	}

	blogger := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
	blogger.EntityTypeID = entityTypeBlogger.ID

	post1 := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
	post1.EntityTypeID = entityTypePost.ID

	post2 := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
	post2.EntityTypeID = entityTypePost.ID

	err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, blogger, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Create blogger": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, post1, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Create post1": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, post2, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Create post2": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.BindRelatedEntities(cntx, rel2Post, blogger, []*entity.Entity{post1, post2}, langRusID, true)
	if err != nil {
		res = append(res, map[string]interface{}{"BindRelatedEntities": err.Error()})
	}

	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, blogger, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update blogger": err.Error()})
	//}
	//
	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, post1, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update post1": err.Error()})
	//}
	//
	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, post2, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update post2": err.Error()})
	//}

	blogger2, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, blogger.ID, entityTypeID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Get": err.Error()})
	}

	err = c.yaruzPlatform.DataSubsystem().Entity.UnbindRelatedEntities(cntx, rel2Post, blogger2, []*entity.Entity{post1}, langRusID, true)
	if err != nil {
		res = append(res, map[string]interface{}{"UnbindRelatedEntities": err.Error()})
	}

	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, blogger2, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update blogger2": err.Error()})
	//}
	//
	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, post1, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update post1": err.Error()})
	//}
	//
	//err = c.yaruzPlatform.DataSubsystem().Entity.Update(cntx, post2, langRusID)
	//if err != nil {
	//	res = append(res, map[string]interface{}{"Update post2": err.Error()})
	//}

	blogger3, err := c.yaruzPlatform.DataSubsystem().Entity.Get(cntx, blogger.ID, entityTypeID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"Get": err.Error()})
	}
	res = append(res, map[string]interface{}{"blogger3": blogger3})

	return ctx.Write(res)
}

func (c dataTestController) userInst(cntx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity"})
	ctx := cntx.Request.Context()

	var langRusID uint
	var langEngID uint
	var err error

	if langRusID, err = c.yaruzPlatform.ReferenceSubsystem().TextLang.GetIDByCode(ctx, config.LangRus); err != nil {
		res = append(res, map[string]interface{}{"TextLang.GetIDByCode(rus): ": err.Error()})
	}

	if langEngID, err = c.yaruzPlatform.ReferenceSubsystem().TextLang.GetIDByCode(ctx, config.LangEng); err != nil {
		res = append(res, map[string]interface{}{"TextLang.GetIDByCode(eng): ": err.Error()})
	}

	user, err := c.user.New(ctx)
	if err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = user.SetName(ctx, "Андрей", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.SetName: ": err.Error()})
	}

	if err = user.SetAge(ctx, 41); err != nil {
		res = append(res, map[string]interface{}{"user.SetAge: ": err.Error()})
	}

	if err = user.SetHeight(ctx, 75); err != nil {
		res = append(res, map[string]interface{}{"user.SetHeight: ": err.Error()})
	}

	if err := c.user.Create(ctx, user, langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	user1, err := c.user.Get(ctx, user.ID, langEngID)
	if err != nil {
		res = append(res, map[string]interface{}{"user.Get: ": err.Error()})
	}

	if err = user1.SetName(ctx, "Andrey", langEngID); err != nil {
		res = append(res, map[string]interface{}{"user.SetName: ": err.Error()})
	}

	if err = user1.SetAge(ctx, 42); err != nil {
		res = append(res, map[string]interface{}{"user.SetAge: ": err.Error()})
	}

	if err = user1.SetWeight(ctx, 88); err != nil {
		res = append(res, map[string]interface{}{"user.SetWeight: ": err.Error()})
	}

	if err := c.user.Update(ctx, user1, langEngID); err != nil {
		res = append(res, map[string]interface{}{"user.Update: ": err.Error()})
	}

	user2, err := c.user.Get(ctx, user.ID, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"user.Get: ": err.Error()})
	}

	res = append(res, map[string]interface{}{"user2": user2})

	if err := c.user.Delete(ctx, user2.ID); err != nil {
		res = append(res, map[string]interface{}{"user.Delete: ": err.Error()})
	}

	return cntx.Write(res)
}

func (c dataTestController) userSearch(cntx *routing.Context) error {
	ctx := cntx.Request.Context()

	struc, err := c.user.New(ctx)
	if err != nil {
		errors.Wrapf(err, "Can not get a new user instant.")
	}

	cond, err := ozzo_routing.ParseQueryParams(cntx, struc)
	if err != nil {
		errors.Wrapf(apperror.ErrBadRequest, err.Error())
	}

	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity"})

	var langRusID uint
	//var langEngID uint

	if langRusID, err = c.yaruzPlatform.ReferenceSubsystem().TextLang.GetIDByCode(ctx, config.LangRus); err != nil {
		res = append(res, map[string]interface{}{"TextLang.GetIDByCode(rus): ": err.Error()})
	}

	//if langEngID, err = c.yaruzPlatform.ReferenceSubsystem().TextLang.GetIDByCode(ctx, config.LangEng); err != nil {
	//	res = append(res, map[string]interface{}{"TextLang.GetIDByCode(eng): ": err.Error()})
	//}

	users := make([]*user.User, 20)

	for i := range users {
		if users[i], err = c.user.New(ctx); err != nil {
			res = append(res, map[string]interface{}{"user.New: ": err.Error()})
			break
		}
		if err = users[i].SetName(ctx, "Имя-"+strconv.Itoa(i), langRusID); err != nil {
			res = append(res, map[string]interface{}{"user.New: ": err.Error()})
			break
		}
		if err = users[i].SetAge(ctx, uint(i*2)); err != nil {
			res = append(res, map[string]interface{}{"user.New: ": err.Error()})
			break
		}
		if err = users[i].SetHeight(ctx, float64(i*2+100)); err != nil {
			res = append(res, map[string]interface{}{"user.New: ": err.Error()})
			break
		}
		if err = users[i].SetWeight(ctx, float64(i*2+50)); err != nil {
			res = append(res, map[string]interface{}{"user.New: ": err.Error()})
			break
		}

		//if err = c.user.Create(ctx, users[i], langRusID); err != nil {
		//	res = append(res, map[string]interface{}{"user.Create: ": err.Error()})
		//	break
		//}
	}

	if err = users[0].SetName(ctx, users[0].Name+" У лукоморья дуб зеленый", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[1].SetName(ctx, users[1].Name+" Златая цепь на дубе том", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[2].SetName(ctx, users[2].Name+" И днем и ночью кот ученый", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[3].SetName(ctx, users[3].Name+" Всё ходит по цепи кругом", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[4].SetName(ctx, users[4].Name+" Идет направо — песнь заводит", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[5].SetName(ctx, users[5].Name+" Идет налево — сказку говорит", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[6].SetName(ctx, users[6].Name+" Там чудеса: там леший бродит", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[7].SetName(ctx, users[7].Name+" Русалка на ветвях идет", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[8].SetName(ctx, users[8].Name+" Там на неведомых дорожках", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[9].SetName(ctx, users[9].Name+" Следы невиданных зверей идут", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[10].SetName(ctx, users[10].Name+" Избушка там на курьих ножках", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[11].SetName(ctx, users[11].Name+" Стоит без окон, без дверей", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[12].SetName(ctx, users[12].Name+" Там лес и дол видений полны", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[13].SetName(ctx, users[13].Name+" Там о заре прихлынут волны", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[14].SetName(ctx, users[14].Name+" На брег песчаный и пустой", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[15].SetName(ctx, users[15].Name+" И тридцать витязей прекрасных", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[16].SetName(ctx, users[16].Name+" Чредой из вод выходят ясных и идут", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[17].SetName(ctx, users[17].Name+" И с ними дядька их морской идет", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[18].SetName(ctx, users[18].Name+" Там королевич мимоходом", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	if err = users[19].SetName(ctx, users[19].Name+" Пленяет грозного царя", langRusID); err != nil {
		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	}

	//for i := range users {
	//
	//	if err = users[i].SetHeight(ctx, users[i].Height+10); err != nil {
	//		res = append(res, map[string]interface{}{"user.New: ": err.Error()})
	//		break
	//	}
	//
	//	if err = c.user.Update(ctx, users[i], langRusID); err != nil {
	//		res = append(res, map[string]interface{}{"user.Update: ": err.Error()})
	//		break
	//	}
	//}

	count, err := c.user.Count(ctx, cond, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"user.Count: ": err.Error()})
	}

	items, err := c.user.Query(ctx, cond, langRusID)

	if err != nil {
		res = append(res, map[string]interface{}{"user.Query: ": err.Error()})
	}
	res = append(res, map[string]interface{}{"users: ": items})

	//for i := range users {
	//	if err := c.user.Delete(ctx, users[i].ID); err != nil {
	//		res = append(res, map[string]interface{}{"user.Delete: ": err.Error()})
	//		break
	//	}
	//}

	res = append(res, map[string]interface{}{"count: ": count})
	res = append(res, map[string]interface{}{"items: ": items})
	//return cntx.Write(pages)
	return cntx.Write(res)
}

//func (c dataTestController) entityPropertiesSearch(ctx *routing.Context) error {
//	res := make([]map[string]interface{}, 0, 10)
//	res = append(res, map[string]interface{}{"test": "entity-relation"})
//	cntx := ctx.Request.Context()
//
//	if err := c.propertyUnitsInit(cntx); err != nil {
//		res = append(res, map[string]interface{}{"propertyUnitsInit": err.Error()})
//	}
//
//	if err := c.propertiesInit(cntx); err != nil {
//		res = append(res, map[string]interface{}{"propertiesInit": err.Error()})
//	}
//
//	if err := c.entityTypesInit(cntx); err != nil {
//		res = append(res, map[string]interface{}{"entityTypesInit": err.Error()})
//	}
//
//	entityTypePost, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TFirst(cntx, &entity_type.EntityType{Sysname: entityTypePostSysname}, langRusID)
//	if err != nil {
//		res = append(res, map[string]interface{}{"finding for entityTypePost": err.Error()})
//	}
//
//	propertyLen, err := c.yaruzPlatform.ReferenceSubsystem().Property.TFirst(cntx, &property.Property{Sysname: propertyLenSysname}, langRusID)
//	if err != nil {
//		res = append(res, map[string]interface{}{"finding for propertyLen": err.Error()})
//	}
//
//	for i := 0; i < 1000; i++ {
//		post := c.yaruzPlatform.DataSubsystem().Entity.NewEntity()
//		post.EntityTypeID = entityTypePost.ID
//		err = c.yaruzPlatform.DataSubsystem().Entity.Service.EntitySetValueForProperty(post, propertyLen, float64(float64(i)/100), langRusID)
//		if err != nil {
//			res = append(res, map[string]interface{}{"Set propertyLen": err.Error()})
//			break
//		}
//		err = c.yaruzPlatform.DataSubsystem().Entity.Create(cntx, post, langRusID)
//
//		if err != nil {
//			res = append(res, map[string]interface{}{"Create post": err.Error()})
//			break
//		}
//	}
//
//	return ctx.Write(res)
//}

func (c dataTestController) propertyUnitsInit(ctx context.Context) error {

	propertyUnitMM := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.NewEntity()
	propertyUnitMM.Sysname = propertyUnitMMSysname
	propertyUnitLenName := "мм"
	propertyUnitLenDesc := "миллиметры"
	propertyUnitMM.Name = &propertyUnitLenName
	propertyUnitMM.Description = &propertyUnitLenDesc

	err := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.TCreate(ctx, propertyUnitMM, langRusID)
	if err != nil {
		return err
	}
	return nil
}

func (c dataTestController) propertiesInit(ctx context.Context) error {

	propertyUnitItem, err := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.TFirst(ctx, &property_unit.PropertyUnit{Sysname: property_unit.SysnameItem}, langRusID)
	if err != nil {
		return err
	}

	propertyNumber := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyNumber.Sysname = propertyNumberSysname
	propertyNumber.PropertyTypeID = property_type.IDInt
	propertyNumber.PropertyUnitID = &propertyUnitItem.ID
	propertyNumberName := "количество"
	propertyNumberDesc := "количество (шт)"
	propertyNumber.Name = &propertyNumberName
	propertyNumber.Description = &propertyNumberDesc

	if err := c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyNumber, langRusID); err != nil {
		return err
	}

	propertyUnitMM, err := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.TFirst(ctx, &property_unit.PropertyUnit{Sysname: "mm"}, langRusID)
	if err != nil {
		return err
	}

	propertyLen := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyLen.Sysname = propertyLenSysname
	propertyLen.PropertyTypeID = property_type.IDFloat
	propertyLen.PropertyUnitID = &propertyUnitMM.ID
	propertyLenName := "длина"
	propertyLenDesc := "длина (мм)"
	propertyLen.Name = &propertyLenName
	propertyLen.Description = &propertyLenDesc

	if err = c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyLen, langRusID); err != nil {
		return err
	}

	propertyBool := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyBool.Sysname = propertyBoolSysname
	propertyBool.PropertyTypeID = property_type.IDBoolean
	propertyBool.PropertyUnitID = &propertyUnitMM.ID
	propertyBoolName := "логическое"
	propertyBoolDesc := "логическое свойство (true/false)"
	propertyBool.Name = &propertyBoolName
	propertyBool.Description = &propertyBoolDesc

	if err = c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyBool, langRusID); err != nil {
		return err
	}

	propertyDate := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyDate.Sysname = propertyDateSysname
	propertyDate.PropertyTypeID = property_type.IDDate
	propertyDate.PropertyUnitID = &propertyUnitMM.ID
	propertyDateName := "дата"
	propertyDateDesc := "дата события"
	propertyDate.Name = &propertyDateName
	propertyDate.Description = &propertyDateDesc

	if err = c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyDate, langRusID); err != nil {
		return err
	}

	propertyTS := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyTS.Sysname = propertyTSSysname
	propertyTS.PropertyTypeID = property_type.IDTime
	propertyTS.PropertyUnitID = &propertyUnitMM.ID
	propertyTSName := "дата"
	propertyTSDesc := "дата события"
	propertyTS.Name = &propertyTSName
	propertyTS.Description = &propertyTSDesc

	if err = c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyTS, langRusID); err != nil {
		return err
	}

	propertyName := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyName.Sysname = propertyNameSysname
	propertyName.PropertyTypeID = property_type.IDText
	propertyName.PropertyUnitID = &propertyUnitMM.ID
	propertyNameName := "имя"
	propertyNameDesc := "ну, чьё-то имя, например"
	propertyName.Name = &propertyNameName
	propertyName.Description = &propertyNameDesc

	if err = c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyName, langRusID); err != nil {
		return err
	}

	propertyOpts := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	propertyOpts.Sysname = propertyOptSysname
	propertyOpts.PropertyTypeID = property_type.IDInt
	propertyOpts.PropertyUnitID = &propertyUnitItem.ID
	propertyOpts.Options = []map[string]interface{}{
		{propertyOpt0Key: propertyOpt0Val},
		{propertyOpt1Key: propertyOpt1Val},
		{propertyOpt2Key: propertyOpt2Val},
	}
	propertyOptsName := "опции"
	propertyOptsDesc := "опции (шт)"
	propertyOpts.Name = &propertyOptsName
	propertyOpts.Description = &propertyOptsDesc

	if err := c.yaruzPlatform.ReferenceSubsystem().Property.TCreate(ctx, propertyOpts, langRusID); err != nil {
		return err
	}

	return nil
}

func (c dataTestController) entityTypesInit(ctx context.Context) error {
	entityTypeBlogger := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entityTypeBlogger.Sysname = entityTypeBloggerSysname
	entityTypeBloggerName := "блогер"
	entityTypeBloggerDesc := "блогер"
	entityTypeBlogger.Name = &entityTypeBloggerName
	entityTypeBlogger.Description = &entityTypeBloggerDesc

	if err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TCreate(ctx, entityTypeBlogger, langRusID); err != nil {
		return err
	}

	entityTypePost := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entityTypePost.Sysname = entityTypePostSysname
	entityTypePostName := "пост"
	entityTypePostDesc := "публикация типа пост"
	entityTypePost.Name = &entityTypePostName
	entityTypePost.Description = &entityTypePostDesc

	if err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TCreate(ctx, entityTypePost, langRusID); err != nil {
		return err
	}

	entityTypeStory := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entityTypeStory.Sysname = entityTypeStorySysname
	entityTypeStoryName := "стори"
	entityTypeStoryDesc := "публикация типа стори"
	entityTypeStory.Name = &entityTypeStoryName
	entityTypeStory.Description = &entityTypeStoryDesc

	if err := c.yaruzPlatform.ReferenceSubsystem().EntityType.TCreate(ctx, entityTypeStory, langRusID); err != nil {
		return err
	}

	relBlogger2Post := c.yaruzPlatform.ReferenceSubsystem().Relation.NewEntity()
	relBlogger2Post.Sysname = relationBlogger2PostSysname
	relBlogger2PostName := "блогер-пост"
	relBlogger2PostDesc := "блогер-пост"
	relBlogger2Post.Name = &relBlogger2PostName
	relBlogger2Post.Description = &relBlogger2PostDesc
	relBlogger2Post.SetDependedEntityType(entityTypeBlogger)
	relBlogger2Post.SetUndependedEntityType(entityTypePost)

	if err := c.yaruzPlatform.ReferenceSubsystem().Relation.TCreate(ctx, relBlogger2Post, langRusID); err != nil {
		return err
	}

	relBlogger2Story := c.yaruzPlatform.ReferenceSubsystem().Relation.NewEntity()
	relBlogger2Story.Sysname = relationBlogger2StorySysname
	relBlogger2StoryName := "блогер-стори"
	relBlogger2StoryDesc := "блогер-стори"
	relBlogger2Story.Name = &relBlogger2StoryName
	relBlogger2Story.Description = &relBlogger2StoryDesc
	relBlogger2Story.SetDependedEntityType(entityTypeBlogger)
	relBlogger2Story.SetUndependedEntityType(entityTypeStory)

	if err := c.yaruzPlatform.ReferenceSubsystem().Relation.TCreate(ctx, relBlogger2Story, langRusID); err != nil {
		return err
	}

	return nil
}*/

func (c dataTestController) sharding(cntx *routing.Context) (err error) {
	ctx := cntx.Request.Context()

	//cond, err := ozzo_routing.ParseQueryParams(cntx, struc)
	//if err != nil {
	//	errors.Wrapf(apperror.ErrBadRequest, err.Error())
	//}

	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "sharding"})

	var langRusID uint
	//var langEngID uint

	if langRusID, err = c.yaruzPlatform.ReferenceSubsystem().TextLang.GetIDByCode(ctx, config.LangRus); err != nil {
		res = append(res, map[string]interface{}{"TextLang.GetIDByCode(rus): ": err.Error()})
	}

	users, advertisers, advertisingCompains, offers, err := c.dataCreate(ctx, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"dataCreate err": err.Error()})
	}

	if err = c.dataUpdate(ctx, users, advertisers, advertisingCompains, offers, langRusID); err != nil {
		res = append(res, map[string]interface{}{"dataUpdate err": err.Error()})
	}

	resG, err := c.dataGet(ctx, users, advertisers, advertisingCompains, offers, langRusID)
	if err != nil {
		res = append(res, map[string]interface{}{"dataUpdate err": err.Error()})
	}
	res = append(res, resG...)

	if err = c.dataDelete(ctx, users, advertisers, advertisingCompains, offers, langRusID); err != nil {
		res = append(res, map[string]interface{}{"dataUpdate err": err.Error()})
	}
	res = append(res, map[string]interface{}{"test": "all done!"})

	return cntx.Write(res)
}

func (c dataTestController) dataCreate(ctx context.Context, langId uint) ([]email.User, []advertiser.Advertiser, []advertising_campaign.AdvertisingCampaign, []offer.Offer, error) {

	users := []email.User{
		{
			Email: "mail1@mail.ru",
			Phone: 79998886655,
		},
		{
			Email: "mail2@mail.ru",
			Phone: 79998886644,
		},
		{
			Email: "mail3@mail.ru",
			Phone: 79998885544,
		},
		{
			Email: "mail4@mail.ru",
			Phone: 79998885533,
		},
		{
			Email: "mail5@mail.ru",
			Phone: 79998884433,
		},
	}

	advertisers := []advertiser.Advertiser{
		{
			Name: "Рекламодатель1",
		},
		{
			Name: "Рекламодатель2",
		},
		{
			Name: "Рекламодатель3",
		},
		{
			Name: "Рекламодатель4",
		},
	}

	compaings := []advertising_campaign.AdvertisingCampaign{
		{
			Name: "Рекламная компания 1",
		},
		{
			Name: "Рекламная компания 2",
		},
		{
			Name: "Рекламная компания 3",
		},
		{
			Name: "Рекламная компания 4",
		},
		{
			Name: "Рекламная компания 5",
		},
	}

	offers := []offer.Offer{
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 1),
			FinishedAt: time.Now().AddDate(0, 1, 1),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 2),
			FinishedAt: time.Now().AddDate(0, 1, 2),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 3),
			FinishedAt: time.Now().AddDate(0, 1, 3),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 4),
			FinishedAt: time.Now().AddDate(0, 1, 4),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 5),
			FinishedAt: time.Now().AddDate(0, 1, 5),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 6),
			FinishedAt: time.Now().AddDate(0, 1, 6),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 7),
			FinishedAt: time.Now().AddDate(0, 1, 7),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 8),
			FinishedAt: time.Now().AddDate(0, 1, 8),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 9),
			FinishedAt: time.Now().AddDate(0, 1, 9),
		},
		{
			CreatedAt:  time.Now(),
			StartedAt:  time.Now().AddDate(0, 0, 10),
			FinishedAt: time.Now().AddDate(0, 1, 10),
		},
	}

	for i := range users {
		user, err := c.user.New(ctx)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if err = user.SetEmail(ctx, users[i].Email); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = user.SetPhone(ctx, users[i].Phone); err != nil {
			return nil, nil, nil, nil, err
		}

		if err := c.user.Create(ctx, user, langId); err != nil {
			return nil, nil, nil, nil, err
		}
		users[i] = *user
	}

	for i := range advertisers {
		advertiser, err := c.advertiser.New(ctx)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if err = advertiser.SetName(ctx, advertisers[i].Name, langId); err != nil {
			return nil, nil, nil, nil, err
		}

		if err := c.advertiser.Create(ctx, advertiser, langId); err != nil {
			return nil, nil, nil, nil, err
		}
		advertisers[i] = *advertiser
	}

	for i := range compaings {
		compaing, err := c.advertisingCampaign.New(ctx)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if err = compaing.SetName(ctx, compaings[i].Name, langId); err != nil {
			return nil, nil, nil, nil, err
		}

		if err := c.advertisingCampaign.Create(ctx, compaing, langId); err != nil {
			return nil, nil, nil, nil, err
		}
		compaings[i] = *compaing
	}

	for i := range offers {
		offer, err := c.offer.New(ctx)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if err = offer.SetCreatedAt(ctx, offers[i].CreatedAt); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = offer.SetStartedAt(ctx, offers[i].StartedAt); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = offer.SetFinishedAt(ctx, offers[i].FinishedAt); err != nil {
			return nil, nil, nil, nil, err
		}

		if err := c.offer.Create(ctx, offer, langId); err != nil {
			return nil, nil, nil, nil, err
		}
		offers[i] = *offer
	}

	return users, advertisers, compaings, offers, nil
}

func (c dataTestController) dataUpdate(ctx context.Context, users []email.User, advertisers []advertiser.Advertiser, compaings []advertising_campaign.AdvertisingCampaign, offers []offer.Offer, langId uint) error {

	for i := range users {
		if err := users[i].SetEmail(ctx, "u_"+users[i].Email); err != nil {
			return err
		}

		if err := c.user.Update(ctx, &users[i], langId); err != nil {
			return err
		}
	}

	for i := range advertisers {
		if err := advertisers[i].SetName(ctx, "u_"+advertisers[i].Name, langId); err != nil {
			return err
		}

		if err := c.advertiser.Update(ctx, &advertisers[i], langId); err != nil {
			return err
		}
	}

	for i := range compaings {
		if err := compaings[i].SetName(ctx, "u_"+compaings[i].Name, langId); err != nil {
			return err
		}

		if err := c.advertisingCampaign.Update(ctx, &compaings[i], langId); err != nil {
			return err
		}
	}

	for i := range offers {
		if err := offers[i].SetFinishedAt(ctx, offers[i].FinishedAt.AddDate(0, 1, 0)); err != nil {
			return err
		}

		if err := c.offer.Update(ctx, &offers[i], langId); err != nil {
			return err
		}
	}

	return nil
}

func (c dataTestController) dataGet(ctx context.Context, users []email.User, advertisers []advertiser.Advertiser, compaings []advertising_campaign.AdvertisingCampaign, offers []offer.Offer, langId uint) ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0, 10)

	for i := range users {
		user, err := c.user.Get(ctx, users[i].ID, langId)
		if err != nil {
			return res, err
		}

		res = append(res, map[string]interface{}{"user #" + strconv.Itoa(int(users[i].ID)): user})
	}

	for i := range advertisers {
		advertiser, err := c.advertiser.Get(ctx, advertisers[i].ID, langId)
		if err != nil {
			return res, err
		}

		res = append(res, map[string]interface{}{"advertiser #" + strconv.Itoa(int(advertisers[i].ID)): advertiser})
	}

	for i := range compaings {
		compaing, err := c.advertisingCampaign.Get(ctx, compaings[i].ID, langId)
		if err != nil {
			return res, err
		}

		res = append(res, map[string]interface{}{"compaing #" + strconv.Itoa(int(compaings[i].ID)): compaing})
	}

	for i := range offers {
		offer, err := c.offer.Get(ctx, offers[i].ID, langId)
		if err != nil {
			return res, err
		}

		res = append(res, map[string]interface{}{"offer #" + strconv.Itoa(int(offers[i].ID)): offer})
	}

	return res, nil
}

func (c dataTestController) dataDelete(ctx context.Context, users []email.User, advertisers []advertiser.Advertiser, compaings []advertising_campaign.AdvertisingCampaign, offers []offer.Offer, langId uint) error {

	for i := range users {
		if err := c.user.Delete(ctx, users[i].ID); err != nil {
			return err
		}
	}

	for i := range advertisers {
		if err := c.advertiser.Delete(ctx, advertisers[i].ID); err != nil {
			return err
		}
	}

	for i := range compaings {
		if err := c.advertisingCampaign.Delete(ctx, compaings[i].ID); err != nil {
			return err
		}
	}

	for i := range offers {
		if err := c.offer.Delete(ctx, offers[i].ID); err != nil {
			return err
		}
	}

	return nil
}
