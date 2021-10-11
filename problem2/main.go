package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-bit/callers"
	"go-bit/entities/proto"
	grpcHandler "go-bit/handlers/grpc"
	"go-bit/handlers/rest"
	"go-bit/repositories"
	"go-bit/services"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	logCallerRepository := repositories.NewLogCallerRepository()
	imdbCaller := callers.NewIMDBCaller(logCallerRepository)
	imdService := services.NewIMDBService(imdbCaller)
	imdbRestHandler := rest.NewIMDBRestHandler(imdService)

	router := mux.NewRouter()
	router.Handle("/v1/imdb/movies", imdbRestHandler.SearchIMDBHandler()).Methods(http.MethodGet)
	router.Handle("/v1/imdb/movies/detail/{id}", imdbRestHandler.GetIMDBDetailHandler()).Methods(http.MethodGet)

	go func() {
		lis, err := net.Listen("tcp", "localhost:4000")
		if err != nil {
			log.Fatal(err)
		}
		// create new grpc server
		srv := grpc.NewServer()

		// register struct MMathService to implement interface MathServiceServer
		proto.RegisterIMDBServiceServer(srv, grpcHandler.NewIMDBGrpcHandler(imdService))

		log.Println("grpc server will started at port 4000")
		// running grpc server over tcp at localhost:4000
		if err := srv.Serve(lis); err != nil {
			log.Println("error when running server because", err.Error())
			panic(err)
		}
	}()

	// Setup http server
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8070",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("We received an interrupt signal, shut down.")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
		log.Println("Bye.")
	}()
	log.Println("Listening on port 8070")
	log.Fatal(srv.ListenAndServe())
	<-idleConnsClosed
}
