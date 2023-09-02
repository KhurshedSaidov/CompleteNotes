package main

import (
	"awesomeNotes/config"
	"awesomeNotes/pkg/handler"
	"awesomeNotes/pkg/middleware"
	"awesomeNotes/pkg/repository"
	"awesomeNotes/pkg/service"
	"log"
	"net/http"
)

func main() {
	Run()
}

func Run() error {
	err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(config.DB)
	noteRepository := repository.NewNoteRepository(config.DB)
	userService := service.UserService{Repository: userRepository}
	noteService := service.NoteService{Repository: noteRepository}
	authMiddleware := middleware.NewAuthMiddleware(&service.UserService{})
	userHandler := handler.UserHandler{UserService: &userService}
	noteHandler := handler.NoteHandler{NoteService: &noteService}
	router := handler.InitRouters(&userHandler, &noteHandler, authMiddleware)
	config, err := config.InitConfigs()
	if err != nil {
		return err
	}
	address := config.Ip + config.Port
	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Println("Listen and server error", err)
		return err
	}
	return nil
}
