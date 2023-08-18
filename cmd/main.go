package main

import (
	"log"
	"net"
	"os"

	"github.com/aburtasov/books/api"
	"github.com/aburtasov/books/pkg/repository"
	"github.com/aburtasov/books/pkg/service"
	transport "github.com/aburtasov/books/pkg/transport/grpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
)

func main() {
	////////////////////////CONFIG INIT//////////////////////////////////////////////////

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s ", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//////////////////////////MYSQL DB INIT////////////////////////////////////////////////

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialized db: %s", err.Error())
	}

	//////////////////////////Iinit Service and Repository///////////////////////////////

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	//////////////////////////GRPC Server. Init and Run//////////////////////////////////

	s := grpc.NewServer()
	srv := transport.NewGRPCServer(*services)

	api.RegisterBookerServer(s, srv)

	l, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
