package main

import (
	"log"
	"net/http"
	"time"

	"git.01.alem.school/quazar/forum-authentication/config"
	httpdelivery "git.01.alem.school/quazar/forum-authentication/delivery/http"
	"git.01.alem.school/quazar/forum-authentication/repository"
	"git.01.alem.school/quazar/forum-authentication/repository/sqlite"
	"git.01.alem.school/quazar/forum-authentication/usecase"
)

func main() {
	// read configs
	config, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatal("config:", err)
	}
	// fmt.Println(config)

	// Create pool of connection for DB
	db, err := sqlite.OpenDB(config.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Setup sql database with predefined categories ...
	if err := sqlite.Setup(db, config.Database.Path); err != nil {
		log.Fatal(err)
	}

	// repositories ...
	UserRepository := repository.NewSqliteUserRepository(db)

	// contextTimeout setup
	timeoutContext := time.Duration(config.Context.Timeout) * time.Second

	// usecases ...
	userUsecase := usecase.NewUserUsecase(UserRepository, timeoutContext)

	// delivery ...
	handler := &httpdelivery.Handler{
		UserUsecase: userUsecase,
	}

	// router init ...
	router := httpdelivery.NewHandler(handler)

	// server configs ...
	srv := &http.Server{
		ReadTimeout:       1 * time.Second, // the maximum duration for reading the entire request, including the body
		WriteTimeout:      1 * time.Second, // the maximum duration before timing out writes of the response
		IdleTimeout:       30 * time.Second, // the maximum amount of time to wait for the next request when keep-alive is enabled
		ReadHeaderTimeout: 2 * time.Second, // the amount of time allowed to read request headers
		Addr:              config.Server.Address,
		Handler:           router,
	}

	log.Printf("Starting server on %s", config.Server.Address)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
