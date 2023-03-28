package entity

import (
	"errors"
	"time"

	"github.com/hpaes/api-project-golang/pkg/entity"
)

var (
	ErrInvalidID      = errors.New("invalid id")
	ErrIDIsRequired   = errors.New("id is required")
	ErrNameIsRequired = errors.New("name is required")
	ErrInvalidPrice   = errors.New("invalid price")
)

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewProduct(name, description string, price float64) (*Product, error) {
	product := &Product{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   entity.GetTime(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}
