package controller

import (
	"fmt"
	"strconv"
	"time"

	"smtp2api/pkg/yarus_platform/reference/domain/text_value"

	"smtp2api/pkg/yarus_platform/reference/domain/property_type"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/minipkg/log"
	"github.com/minipkg/selection_condition"
	"smtp2api/pkg/yarus_platform"
)

type referenceTestController struct {
	Logger        log.ILogger
	yaruzPlatform yarus_platform.IPlatform
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/models/ - список всех моделей
//	GET /api/model/{ID} - детали модели
func RegisterReferenceTestHandlers(r *routing.RouteGroup, yaruzPlatform yarus_platform.IPlatform, logger log.ILogger, authHandler routing.Handler) {
	c := referenceTestController{
		Logger:        logger,
		yaruzPlatform: yaruzPlatform,
	}

	r.Get("/text-source", c.textSource)
	r.Get("/text-value", c.textValue)
	r.Get("/property-unit", c.propertyUnit)
	r.Get("/property-group", c.propertyGroup)
	r.Get("/property-view-type", c.propertyViewType)
	r.Get("/property-type", c.propertyType)
	r.Get("/property-type_with-text", c.propertyTypeWithText)
	r.Get("/property", c.property)
	r.Get("/relation", c.relation)
	r.Get("/property-options-validation", c.propertyOptionsValidation)
	r.Get("/entity-type", c.entityType)
	r.Get("/entity-type-binding", c.entityTypeBinding)
}

func (c referenceTestController) textSource(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "text-source"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().TextSource.NewEntity()

	err := c.yaruzPlatform.ReferenceSubsystem().TextSource.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"1. errCreate": err.Error()})
	}

	val := c.yaruzPlatform.ReferenceSubsystem().TextValue.NewEntity()
	val.TextSourceID = entity.ID
	val.TextLangID = 1
	val.Value = "Тестовое значение текстового поля"

	err = c.yaruzPlatform.ReferenceSubsystem().TextValue.Create(cntx, val)
	if err != nil {
		res = append(res, map[string]interface{}{"1. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().TextSource.TGet(cntx, entity.ID, val.TextLangID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": e})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().TextSource.Delete(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) textValue(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "text-value"})
	cntx := ctx.Request.Context()

	source := c.yaruzPlatform.ReferenceSubsystem().TextSource.NewEntity()
	err := c.yaruzPlatform.ReferenceSubsystem().TextSource.Create(cntx, source)
	if err != nil {
		res = append(res, map[string]interface{}{"1. errCreate": err.Error()})
	}

	entity := c.yaruzPlatform.ReferenceSubsystem().TextValue.NewEntity()
	entity.TextSourceID = source.ID
	entity.TextLangID = 1
	entity.Value = "text value"
	err = c.yaruzPlatform.ReferenceSubsystem().TextValue.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. Create err": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().TextValue.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"3. Get err": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"3. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Value = "updated text value - " + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().TextValue.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"4. Update err": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().TextValue.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"5. Get err": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"5. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().TextValue.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Value",
			Condition: "eq",
			Value:     entity.Value,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().TextValue.Delete(cntx, &text_value.TextValue{ID: entity.ID})
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) propertyUnit(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property-unit"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.NewEntity()
	entity.Sysname = "WrongName"

	err := entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. errValidate": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	entity.Sysname = "property_unit"
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "property_unit_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyUnit.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) propertyGroup(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property-group"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.NewEntity()
	entity.Sysname = "WrongName"

	err := entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. errValidate": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	entity.Sysname = "property_group"
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "property_group_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) propertyViewType(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property-view-type"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.NewEntity()
	entity.Sysname = "WrongName"

	err := entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. errValidate": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	entity.Sysname = "property_view_type"
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "property_view_type_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) propertyType(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property-type"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().PropertyType.NewEntity()
	entity.Sysname = "WrongName"

	err := entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. errValidate": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	entity.Sysname = "property_type"
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().PropertyType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "property_type_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().PropertyType.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	// References
	viewType := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.NewEntity()
	viewType.Sysname = "view_type_" + strconv.Itoa(int(time.Now().Unix()))

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Create(cntx, viewType)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errCreate ViewType": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.BindPropertyViewType(cntx, entity, viewType.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"8. BindPropertyViewType err": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.InitPropertyViewTypes(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"9. GetPropertyViewTypes err": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"9. GetPropertyViewTypes list": fmt.Sprintf("%#v", entity.PropertyViewTypes)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.UnbindPropertyViewType(cntx, entity, viewType.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"10. UnbindPropertyViewType err": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"11. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) propertyTypeWithText(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property-type-with-text"})
	cntx := ctx.Request.Context()
	langID := uint(1)

	entity := c.yaruzPlatform.ReferenceSubsystem().PropertyType.NewEntity()
	entity.Sysname = "property_type_" + strconv.Itoa(int(time.Now().Unix()))
	name := "name1_" + strconv.Itoa(int(time.Now().Unix()))
	entity.Name = &name

	err := c.yaruzPlatform.ReferenceSubsystem().PropertyType.TCreate(cntx, entity, langID)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().PropertyType.TGet(cntx, entity.ID, langID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	desc := "desc_" + strconv.Itoa(int(time.Now().Unix()))
	entity.Description = &desc
	*entity.Name = "name_" + strconv.Itoa(int(time.Now().Unix()))
	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.TUpdate(cntx, entity, langID)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.TGet(cntx, entity.ID, langID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().PropertyType.TQuery(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	}, langID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyType.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"11. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) property(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property"})
	cntx := ctx.Request.Context()
	// ViewType
	viewType := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.NewEntity()
	viewType.Sysname = "view_type_" + strconv.Itoa(int(time.Now().Unix()))

	err := c.yaruzPlatform.ReferenceSubsystem().PropertyViewType.Create(cntx, viewType)
	if err != nil {
		res = append(res, map[string]interface{}{"0. Create ViewType err": err.Error()})
	}
	// Group
	group := c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.NewEntity()
	group.Sysname = "group_" + strconv.Itoa(int(time.Now().Unix()))

	err = c.yaruzPlatform.ReferenceSubsystem().PropertyGroup.Create(cntx, group)
	if err != nil {
		res = append(res, map[string]interface{}{"0. Create Group err": err.Error()})
	}
	var one uint = 1
	entity := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	entity.Sysname = "WrongName"
	entity.PropertyTypeID = property_type.IDInt
	entity.PropertyUnitID = &one
	entity.PropertyViewTypeID = &one
	entity.PropertyGroupID = &one

	err = entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. errValidate": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().Property.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. errCreate1": err.Error()})
	}

	entity.Sysname = "property" + strconv.Itoa(int(time.Now().Unix()))
	err = c.yaruzPlatform.ReferenceSubsystem().Property.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. errCreate": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().Property.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. errCreate1": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. entity1": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "property_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().Property.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. errCreate": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().Property.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().Property.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"6. errGet": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().Property.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"7. errDelete": err.Error()})
	}

	return ctx.Write(res)
}

type PropertyOptionsTestCases [][][]map[string]interface{}

func (c referenceTestController) propertyOptionsValidation(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "property option validation"})

	entity := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	entity.Sysname = "options_validation_test_" + strconv.Itoa(int(time.Now().Unix()))
	var one uint = 1
	entity.PropertyUnitID = &one

	var typeTestCases = [][][]map[string]interface{}{
		{ // boolean
			{
				{"1": 1},
				{"2": 2},
			},
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": true},
				{"2": false},
			},
		},
		{ // int
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": true},
				{"2": false},
			},
			{
				{"1": 1},
				{"2": 2},
			},
		},
		{ // bigint
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": true},
				{"2": false},
			},
			{
				{"1": 1},
				{"2": 2},
			},
		},
		{ // float
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": true},
				{"2": false},
			},
			{
				{"1": 1},
				{"2": 2},
			},
			{
				{"1": 1.1},
				{"2": 2.1},
			},
		},
		{ // date
			{
				{"1": 1},
				{"2": 2},
			},
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": "2021-05-27"},
				{"2": "2021-05-30"},
			},
		},
		{ // timestamp
			{
				{"1": "one"},
				{"2": "two"},
			},
			{
				{"1": true},
				{"2": false},
			},
			{
				{"1": 1405544146},
				{"2": 1405545146},
			},
		},
		{ // text
			{
				{"1": 1},
				{"2": 2},
			},
			{
				{"1": true},
				{"2": false},
			},
			{
				{"1": "one"},
				{"2": "two"},
			},
		},
	}

	for propertyType, testCases := range typeTestCases {
		propertyType++
		entity.PropertyTypeID = uint(propertyType)

		for caseNum, options := range testCases {
			entity.Options = options

			err := entity.Validate()
			if err != nil {
				res = append(res, map[string]interface{}{"PropertyType:" + strconv.Itoa(propertyType) + " test case #" + strconv.Itoa(caseNum) + ". error": err.Error()})
			} else {
				res = append(res, map[string]interface{}{"PropertyType:" + strconv.Itoa(propertyType) + " test case #" + strconv.Itoa(caseNum) + ".": "success"})
			}
		}
	}

	return ctx.Write(res)
}

func (c referenceTestController) entityType(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity-type"})
	cntx := ctx.Request.Context()

	entity := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entity.Sysname = "WrongName"

	err := entity.Validate()
	if err != nil {
		res = append(res, map[string]interface{}{"1. EntityType Validation error: ": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"2. EntityType Creation error: ": err.Error()})
	}

	entity.Sysname = "entity_type"
	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. EntityType Creation error: ": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"3. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. EntityType: ": fmt.Sprintf("%#v", e)})
	}

	entity.Sysname = "entity_type_" + strconv.Itoa(int(entity.ID))
	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Update(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"5. EntityType Updating error: ": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entity2": fmt.Sprintf("%#v", e)})
	}

	list, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.Query(cntx, &selection_condition.SelectionCondition{
		Where: selection_condition.WhereCondition{
			Field:     "Sysname",
			Condition: "eq",
			Value:     entity.Sysname,
		},
	})
	if err != nil {
		res = append(res, map[string]interface{}{"7. EntityType Getting a list error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"7. list": fmt.Sprintf("%#v", list)})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Delete(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"8. EntityType Deleting error: ": err.Error()})
	}

	return ctx.Write(res)
}

func (c referenceTestController) entityTypeBinding(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity-type"})
	cntx := ctx.Request.Context()

	var one uint = 1
	property1 := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	property1.Sysname = "property_" + strconv.Itoa(int(time.Now().Unix()))
	property1.PropertyTypeID = property_type.IDInt
	property1.PropertyUnitID = &one
	property2 := c.yaruzPlatform.ReferenceSubsystem().Property.NewEntity()
	property2.Sysname = "property_" + strconv.Itoa(int(time.Now().Unix()))
	property2.PropertyTypeID = property_type.IDFloat
	property2.PropertyUnitID = &one
	err := c.yaruzPlatform.ReferenceSubsystem().Property.Create(cntx, property1)
	if err != nil {
		res = append(res, map[string]interface{}{"0. Property1 Create error: ": err.Error()})
	}
	err = c.yaruzPlatform.ReferenceSubsystem().Property.Create(cntx, property2)
	if err != nil {
		res = append(res, map[string]interface{}{"0. Property2 Create error: ": err.Error()})
	}

	entity := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entity.Sysname = "entity_type_" + strconv.Itoa(int(time.Now().Unix()))
	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"1. EntityType Creation error: ": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.BindProperty(cntx, entity.ID, property1.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"2. EntityType Updating error: ": err.Error()})
	}
	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.BindProperty(cntx, entity.ID, property2.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"2. EntityType Updating error: ": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"3. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"3. EntityType: ": fmt.Sprintf("%#v", e)})
	}

	props, rels, err := c.yaruzPlatform.ReferenceSubsystem().Relation.GetPropertiesAndRelationsByEntityTypeID(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. EntityType props: ": props})
		res = append(res, map[string]interface{}{"4. EntityType rels: ": rels})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.UnbindProperty(cntx, entity.ID, property1.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"5. EntityType Unbind error: ": err.Error()})
	}

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.UnbindAllProperties(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. EntityType UnbindAll error: ": err.Error()})
	}

	e, err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"7. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"7. entity2": fmt.Sprintf("%#v", e)})
	}

	props, rels, err = c.yaruzPlatform.ReferenceSubsystem().Relation.GetPropertiesAndRelationsByEntityTypeID(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. EntityType props: ": props})
		res = append(res, map[string]interface{}{"4. EntityType rels: ": rels})
	}

	return ctx.Write(res)
}

func (c referenceTestController) relation(ctx *routing.Context) error {
	res := make([]map[string]interface{}, 0, 10)
	res = append(res, map[string]interface{}{"test": "entity-type"})
	cntx := ctx.Request.Context()

	entityType1 := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entityType1.Sysname = "entity_type1_" + strconv.Itoa(int(time.Now().Unix()))

	err := c.yaruzPlatform.ReferenceSubsystem().EntityType.Create(cntx, entityType1)
	if err != nil {
		res = append(res, map[string]interface{}{"1. Create EntityType1 error": err.Error()})
	}

	entityType2 := c.yaruzPlatform.ReferenceSubsystem().EntityType.NewEntity()
	entityType2.Sysname = "entity_type2_" + strconv.Itoa(int(time.Now().Unix()))

	err = c.yaruzPlatform.ReferenceSubsystem().EntityType.Create(cntx, entityType2)
	if err != nil {
		res = append(res, map[string]interface{}{"2. Create EntityType2 error": err.Error()})
	}

	entity := c.yaruzPlatform.ReferenceSubsystem().Relation.NewEntity()
	entity.Sysname = "relation_" + strconv.Itoa(int(time.Now().Unix()))
	entity.PropertyTypeID = property_type.IDRelation
	entity.UndependedEntityType = entityType2
	entity.DependedEntityType = entityType1

	err = c.yaruzPlatform.ReferenceSubsystem().Relation.Create(cntx, entity)
	if err != nil {
		res = append(res, map[string]interface{}{"3. Create Relation error": err.Error()})
	}

	e, err := c.yaruzPlatform.ReferenceSubsystem().Relation.Get(cntx, entity.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"4. Get new Relation error": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"4. New Relation": fmt.Sprintf("%#v", e)})
	}

	props, rels, err := c.yaruzPlatform.ReferenceSubsystem().Relation.GetPropertiesAndRelationsByEntityTypeID(cntx, entityType1.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"5. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"5. EntityType props: ": props})
		res = append(res, map[string]interface{}{"5. EntityType rels: ": rels})
	}

	entityType, err := c.yaruzPlatform.ReferenceSubsystem().EntityType.Get(cntx, entityType1.ID)
	if err != nil {
		res = append(res, map[string]interface{}{"6. EntityType Getting error: ": err.Error()})
	} else {
		res = append(res, map[string]interface{}{"6. entityType": entityType})
	}

	return ctx.Write(res)
}
