package handler

import (
	"awesomeNotes/models"
	"awesomeNotes/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *service.UserService
}

type NoteHandler struct {
	NoteService *service.NoteService
}

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	password := r.Header.Get("Password")

	if err := uh.UserService.RegisterUser(username, password); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Пользователь успешно зарегистрирован"))
	} else if err == service.ErrUsernameTaken {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Пользователь уже существует"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка сервера"))
	}
}
func (uh *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неправильный формат JSON"))
		return
	}
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Нет доступа"))
		return
	}
	err := uh.NoteService.CreateNote(user.ID, note.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка создания заметки"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Заметка успешно создана"))
}

func (uh *NoteHandler) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	note, err := uh.NoteService.GetNoteByID(uint(noteId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Заметка не найдена"))
		return
	}

	if note.UserId != user.ID {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Доступ запрещен"))
		return
	}

	var updatedNote models.Note
	err = json.NewDecoder(r.Body).Decode(&updatedNote)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	err = uh.NoteService.UpdateNote(uint(noteId), updatedNote.Content)
	if err != nil {
		http.Error(w, "Ошибка изменения заметки", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Заметка успешно изменена"))
}

func (uh *NoteHandler) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	note, err := uh.NoteService.GetNoteByID(uint(noteId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Заметка не найдена"))
		return
	}

	if note.UserId != user.ID {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Доступ запрещен"))
		return
	}

	err = uh.NoteService.DeleteNote(uint(noteId))
	if err != nil {
		http.Error(w, "Ошибка удаления заметки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Заметка успешно удалена"))
}

func (uh *NoteHandler) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	note, err := uh.NoteService.GetNoteByID(uint(noteId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Заметка не найдена"))
		return
	}

	if note.UserId != user.ID {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Доступ запрещен"))
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (uh *NoteHandler) GetAllUserNotes(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	notes, err := uh.NoteService.GetAllNotes(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при получении заметок пользователя"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
