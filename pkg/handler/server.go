package handler

import (
	"awesomeNotes/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitRouters(handler *UserHandler, noteHandler *NoteHandler, authMiddleware *middleware.AuthMiddleware) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	authRoute := r.PathPrefix("/auth").Subrouter()
	authRoute.Use(authMiddleware.Middleware)
	authRoute.HandleFunc("/notes", noteHandler.CreateNote).Methods("POST")
	authRoute.HandleFunc("/notes/{id}", noteHandler.UpdateNoteHandler).Methods("PUT")
	authRoute.HandleFunc("/notes/{id}", noteHandler.DeleteNoteHandler).Methods("DELETE")
	authRoute.HandleFunc("/notes/{id}", noteHandler.GetNoteByIDHandler).Methods("GET")
	authRoute.HandleFunc("/notes", noteHandler.GetAllUserNotes).Methods("GET")
	return r
}
