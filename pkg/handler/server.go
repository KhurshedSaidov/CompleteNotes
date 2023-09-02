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
	//r.HandleFunc("/notes/{id}", handler.UpdateNoteHandler).Methods("PUT")
	//r.HandleFunc("/notes/{id}", handler.DeleteNoteHandler).Methods("DELETE")
	//r.HandleFunc("/notes/{id}", handler.GetNoteByIDHandler).Methods("GET")
	return r
}
