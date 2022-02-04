package offer

import (
	"context"
	"time"

	"smtp2api/pkg/yarus_platform/data/domain/entity"
)

const (
	EntityType                = "Offer"
	PropertySysnameCreatedAt  = "Offer.CreatedAt"
	PropertySysnameStartedAt  = "Offer.StartedAt"
	PropertySysnameFinishedAt = "Offer.FinishedAt"
)

// Offer is the user entity
type Offer struct {
	*entity.Entity
	ID         uint
	CreatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time
}

var _ entity.Searchable = (*Offer)(nil)

func (e Offer) EntityType() string {
	return EntityType
}

func (e Offer) Validate() error {
	return nil
}

func (e *Offer) GetMapNameSysname() map[string]string {
	return map[string]string{
		"CreatedAt":  PropertySysnameCreatedAt,
		"StartedAt":  PropertySysnameStartedAt,
		"FinishedAt": PropertySysnameFinishedAt,
	}
}

func (e *Offer) SetCreatedAt(ctx context.Context, value time.Time) error {
	prop, err := e.PropertyFinder.GetBySysname(ctx, PropertySysnameCreatedAt, 0)
	if err != nil {
		return err
	}

	if err = e.Entity.SetValueForProperty(prop, value, 0); err != nil {
		return err
	}

	e.CreatedAt = value
	return nil
}

func (e *Offer) SetStartedAt(ctx context.Context, value time.Time) error {
	prop, err := e.PropertyFinder.GetBySysname(ctx, PropertySysnameStartedAt, 0)
	if err != nil {
		return err
	}

	if err = e.Entity.SetValueForProperty(prop, value, 0); err != nil {
		return err
	}

	e.StartedAt = value
	return nil
}

func (e *Offer) SetFinishedAt(ctx context.Context, value time.Time) error {
	prop, err := e.PropertyFinder.GetBySysname(ctx, PropertySysnameFinishedAt, 0)
	if err != nil {
		return err
	}

	if err = e.Entity.SetValueForProperty(prop, value, 0); err != nil {
		return err
	}

	e.FinishedAt = value
	return nil
}
