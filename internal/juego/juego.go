package juego

type Juego struct {
	Id     string
	Tipo   string
	Nombre string
}

type Store interface {
	GetJuegoById(id string) (Juego, error)
	//InsertJuego(rjg Juego) (Juego, error)
	//DeleteJuego(id string) error
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

/*
func (s Service) InsertJuego(rjg Juego) (Juego, error) {
	rjg, err := s.Store.InsertJuego(rjg)
	if err != nil {
		return Juego{}, err
	}
	return rjg, nil
}
*/

/*
func (s Service) DeleteJuego(id string) error {
	err := s.Store.DeleteJuego(id)
	if err != nil {
		return err
	}
	return nil
}
*/
