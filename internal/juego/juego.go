package juego

import (
	"context"
)

type Juego struct {
	Id     string
	Tipo   string
	Nombre string
}

type Store interface {
	GetJuegoById(id string) (Juego, error)
	AddJuego(rjg Juego) (Juego, error)
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetJuegoById(id string) (Juego, error) {
	rjg, err := s.Store.GetJuegoById(id)
	if err != nil {
		return Juego{}, err
	}
	return rjg, nil
}

func (s Service) AddJuego(ctx context.Context, rjg Juego) (Juego, error) {
	rjg, err := s.Store.AddJuego(rjg)
	if err != nil {
		return Juego{}, err
	}
	return rjg, nil
}
