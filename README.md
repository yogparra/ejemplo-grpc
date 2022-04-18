# ejemplo-grpc
    go mod init github.com/yogparra/ejemplo-grpc

# bd    
    docker run --name ejemplo-gppc -p 5432:5432 -e POSTGRES_USER=gmo -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=GRPC -d postgres

# Script squema public
    create table JUEGOS(id serial not null primary key, tipo varchar (50), nombre varchar (100));
    insert into public.JUEGOS (id, tipo, nombre) values(1, 'peleas', 'mortal kombat');
    insert into public.JUEGOS (id, tipo, nombre) values(2, 'peleas', 'kof');
    insert into public.JUEGOS (id, tipo, nombre) values(3, 'peleas', 'street fighting');
    insert into public.JUEGOS (id, tipo, nombre) values(4, 'estrategia', 'factorio');
    insert into public.JUEGOS (id, tipo, nombre) values(5, 'supervivencia', 'dont starve together');
    select * from public.JUEGOS j

# export
        Variables de bd
            export DB_USERNAME=gmo
            export DB_PASSWORD=1234
            export DB_HOST=localhost
            export DB_TABLE=GRPC
            export DB_PORT=5432
            export DB_SSL_MODE=disable       
                
# protobuf
protoc --version
apt install -y protobuf-compiler
sudo apt autoremove

# paquetes
    google.golang.org/grpc
    github.com/jmoiron/sqlx    

# verciones 
go get github.com/xxxx/xxxx@v0.13.0

# gRPC clients
    Evans (https://github.com/ktr0731/evans)
        ejemplo-grpc/internal/proto$ evans juego.proto
        show package
        call GetJuegoById

    BloomRPC

