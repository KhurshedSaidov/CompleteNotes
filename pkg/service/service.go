package service

import (
	"awesomeNotes/models"
	"awesomeNotes/pkg/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

var ErrUsernameTaken = errors.New("Пользователь уже существует")

type UserService struct {
	Repository *repository.UserRepository
}

type NoteService struct {
	Repository *repository.NotesRepository
}

func (us *UserService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = us.Repository.GetUserByUsername(username)
	if err == nil {
		return ErrUsernameTaken
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if err := us.Repository.CreateUser(username, string(hashedPassword)); err != nil {
		return err
	}
	return nil
}

func (us *UserService) LoginUser(username, password string) (*models.User, error) {
	user, err := us.Repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *NoteService) CreateNote(userId uint, content string) error {
	note := &models.Note{
		Content: content,
		UserId:  userId,
	}
	log.Println(note)
	return s.Repository.Create(note)
}

func (s *NoteService) UpdateNote(id uint, content string) error {
	note, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}
	note.Content = content
	return s.Repository.Update(note)
}

func (s *NoteService) DeleteNote(id uint) error {
	return s.Repository.Delete(id)
}

func (s *NoteService) GetNoteByID(id uint) (*models.Note, error) {
	return s.Repository.GetByID(id)
}

func (s *NoteService) GetAllNotes(userId uint) (*models.Note, error) {
	return s.Repository.GetAllNotes(userId)
}
