package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/XeniaBgd/CleanArch/internal/adapters/repository/mongodb"
	"github.com/XeniaBgd/CleanArch/internal/application"
	"github.com/XeniaBgd/CleanArch/internal/config"
	"github.com/XeniaBgd/CleanArch/internal/domain/service"
	book_usecase "github.com/XeniaBgd/CleanArch/internal/domain/usecase/book"
	http_v1 "github.com/XeniaBgd/CleanArch/internal/transport/http"
	"github.com/XeniaBgd/CleanArch/internal/transport/http/server"
	"github.com/gin-gonic/gin"
)

func main() {
	confPath, err := mustFlag()
	if err != nil {
		log.Fatal(err)
	}

	// confData := config.MustReadFile(confPath)
	// conf := params.MustGetConfData(confData)
	conf := config.MustGetConfData(confPath)
	fmt.Println(conf)

	bookStorage := mongodb.NewBookStorage(nil)
	bookService := service.NewBookService(bookStorage)

	authorStorage := mongodb.NewAuthorStorage(nil)
	authorService := service.NewAuthorService(authorStorage)

	bookUsecase := book_usecase.NewBookUsecase(bookService, authorService)
	handler := http_v1.NewBookHandler(bookUsecase)

	serverHandler := gin.Default()
	handler.Register(serverHandler)

	srv := server.NewServer(serverHandler, conf.Srv)
	app := application.Application{
		MainFunc: srv.Start,
		Conf:     conf.App,
	}

	fmt.Println(app.Run())
}

var ErrConfNotFound = errors.New("Flag 'conf' not found")

func mustFlag() (string, error) {
	conf := flag.String("conf", "", "Config file")

	flag.Parse()

	c := *conf
	if c == "" {
		log.Fatal(ErrConfNotFound)
	}

	return c, nil
}
