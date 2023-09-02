package handler

import (
	"awesomeNotes/models"
	"awesomeNotes/pkg/service"
	"encoding/json"
	"net/http"
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
		w.Write([]byte("Unauthorized"))
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

//
//func (h *NoteHandler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
//	var note models.Note
//	err := json.NewDecoder(r.Body).Decode(&note)
//	if err != nil {
//		http.Error(w, "Не правильный формат JSON", http.StatusBadRequest)
//		return
//	}
//
//	err = h.Service.CreateNote(note.Content)
//	if err != nil {
//		http.Error(w, "Ошибка создания заметки", http.StatusInternalServerError)
//		return
//	}
//	json.NewEncoder(w).Encode(note)
//}
//
//func (h *NoteHandler) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
//		return
//	}
//
//	var note models.Note
//	err = json.NewDecoder(r.Body).Decode(&note)
//	if err != nil {
//		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
//		return
//	}
//
//	err = h.Service.UpdateNote(uint(noteId), note.Content)
//	if err != nil {
//		http.Error(w, "Ошибка изменения заметки", http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(note)
//}
//
//func (h *NoteHandler) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
//		return
//	}
//
//	err = h.Service.DeleteNote(uint(noteId))
//	if err != nil {
//		http.Error(w, "Ошибка удаления заметки", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
//
//func (h *NoteHandler) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	noteId, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		http.Error(w, "Не правильно указан ID", http.StatusBadRequest)
//		return
//	}
//
//	note, err := h.Service.GetNoteByID(uint(noteId))
//	if err != nil {
//		http.Error(w, "Заметка не найдена!", http.StatusNotFound)
//		return
//	}
//
//	json.NewEncoder(w).Encode(note)
//}
