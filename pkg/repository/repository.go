package repository

import (
	"awesomeNotes/config"
	"awesomeNotes/models"
	_ "database/sql"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type NotesRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func NewNoteRepository(db *gorm.DB) *NotesRepository {
	return &NotesRepository{DB: db}
}

func (ur *UserRepository) CreateUser(username, hashedPassword string) error {
	newUser := &models.User{Username: username, Password: hashedPassword}
	if err := config.DB.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *NotesRepository) GetByID(id uint) (*models.Note, error) {
	note := &models.Note{}
	if err := r.DB.First(note, id).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (r *NotesRepository) GetAllNotes(userId uint) (*models.Note, error) {
	var notes models.Note
	if err := r.DB.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return nil, err
	}
	return &notes, nil
}

func (r *NotesRepository) Create(note *models.Note) error {
	return config.DB.Create(note).Error
}

func (r *NotesRepository) Update(note *models.Note) error {
	return config.DB.Save(note).Error
}

func (r *NotesRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Note{}, id).Error
}
