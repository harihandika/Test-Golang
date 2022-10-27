package repositories

import (
	"golang-test/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindTodo() ([]models.Todo, error)
	GetTodo(ID int) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(todo models.Todo) (models.Todo, error)
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTodo() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error

	return todos, err
}

func (r *repository) GetTodo(ID int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, ID).Error

	return todo, err
}

func (r *repository) CreateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error

	return todo, err
}

func (r *repository) UpdateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error

	return todo, err
}

func (r *repository) DeleteTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Delete(&todo).Error

	return todo, err
}
