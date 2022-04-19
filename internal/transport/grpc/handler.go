package grpc

import (
	"context"
	"log"
	"net"

	"github.com/yogparra/ejemplo-grpc/internal/juego"
	pbjg "github.com/yogparra/ejemplo-grpc/internal/pb"

	"google.golang.org/grpc"
)

type JuegoService interface {
	GetJuegoById(id string) (juego.Juego, error)
	AddJuego(ctx context.Context, rjg juego.Juego) (juego.Juego, error)
}

type Handler struct {
	JuegoService JuegoService
}

func New(jgService JuegoService) Handler {
	return Handler{
		JuegoService: jgService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pbjg.RegisterJuegoServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}

func (h Handler) GetJuegoById(ctx context.Context, req *pbjg.GetJuegoRequest) (*pbjg.GetJuegoResponse, error) {
	log.Print("Get Juego gRPC endpoint")
	juego, err := h.JuegoService.GetJuegoById(req.Id)
	if err != nil {
		return &pbjg.GetJuegoResponse{}, err
	}
	return &pbjg.GetJuegoResponse{
		Juego: &pbjg.Juego{
			Id:     juego.Id,
			Tipo:   juego.Tipo,
			Nombre: juego.Nombre,
		},
	}, nil
}

func (h Handler) AddJuego(ctx context.Context, req *pbjg.AddJuegoRequest) (*pbjg.AddJuegoResponse, error) {
	log.Print("Add Juego gRPC endpoint")

	newRjg, err := h.JuegoService.AddJuego(ctx, juego.Juego{
		Id:     req.Juego.Id,
		Tipo:   req.Juego.Tipo,
		Nombre: req.Juego.Nombre,
	})
	if err != nil {
		log.Print("Error al insertar un juego en la base de datos")
		return &pbjg.AddJuegoResponse{}, err
	}
	return &pbjg.AddJuegoResponse{
		Juego: &pbjg.Juego{
			Id:     newRjg.Id,
			Tipo:   newRjg.Tipo,
			Nombre: newRjg.Nombre,
		},
	}, nil
}
