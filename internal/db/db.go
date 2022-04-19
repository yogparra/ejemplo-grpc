package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/yogparra/ejemplo-grpc/internal/juego"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {

	dbHost := "localhost"
	dbPort := "5432"
	dbTable := "GRPC"
	dbUsername := "gmo"
	dbPassword := "1234"
	dbSSLMode := "disable"

	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbTable,
		dbUsername,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

func (s Store) GetJuegoById(id string) (juego.Juego, error) {
	var rjg juego.Juego
	row := s.db.QueryRow(
		`SELECT id, tipo, nombre FROM public.juegos where id=$1;`,
		id,
	)
	err := row.Scan(&rjg.Id, &rjg.Tipo, &rjg.Nombre)
	if err != nil {
		log.Print(err.Error())
		return juego.Juego{}, err
	}
	return rjg, nil
}

func (s Store) AddJuego(rjg juego.Juego) (juego.Juego, error) {
	_, err := s.db.NamedQuery(`INSERT INTO juegos (id, tipo, nombre) VALUES (:id, :tipo, :nombre)`, rjg)
	if err != nil {
		return juego.Juego{}, errors.New("failed to insert into database")
	}
	return juego.Juego{
		Id:     rjg.Id,
		Tipo:   rjg.Tipo,
		Nombre: rjg.Nombre,
	}, nil
}
