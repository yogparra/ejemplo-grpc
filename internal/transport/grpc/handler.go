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
	log.Print("Get Juego gRPC Endpoint Hit")
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
