package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Nykenik24/enkrypted/internal/crypto"
	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Create(client models.User) (*models.User, error)
	Delete(id string) (int, error)
	Count() (int, error)
	Clear() error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo() UserRepository {
	return &userRepository{db: db.GetInstance().Database}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("error retrieving users: %v", err)
	}

	slog.Info("retrieved all users from users table", "count", result.RowsAffected)
	return users, nil
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	ctx := context.Background()

	user, err := gorm.G[models.User](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting user by id (%s): %v", id, err)
	}

	slog.Info("retrieved user by id from users table", "id", id)
	return &user, nil
}

func (r *userRepository) GetByName(name string) (*models.User, error) {
	ctx := context.Background()

	user, err := gorm.G[models.User](r.db).Where("username = ?", name).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting user by name (%s): %v", name, err)
	}

	slog.Info("retrieved user by id from users table", "id", name)
	return &user, nil
}

func (r *userRepository) Create(user models.User) (*models.User, error) {
	ctx := context.Background()

	hashedPassword, err := crypto.HashPasswordSecure(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing user password: %v", err)
	}
	user.Password = hashedPassword

	err = gorm.G[models.User](r.db).Create(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("error inserting user into database: %v", err)
	}

	slog.Info("created new user", "user", user)

	return &user, nil
}

func (r *userRepository) Delete(id string) (int, error) {
	ctx := context.Background()

	rowsAffected, err := gorm.G[models.User](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return -1, fmt.Errorf("error deleting user from database by ID (%s): %v", id, err)
	}

	slog.Info("deleted user from database by ID", "id", id)

	return rowsAffected, nil
}

func (r *userRepository) Count() (int, error) {
	var users []models.User
	result := r.db.Find(&users)

	if err := result.Error; err != nil {
		return -1, fmt.Errorf("error retrieving users: %v", err)
	}

	return len(users), nil
}

func (r *userRepository) Clear() error {
	ctx := context.Background()

	_, err := gorm.G[models.User](r.db).Delete(ctx)
	if err != nil {
		return fmt.Errorf("error clearing users from database: %v", err)
	}

	slog.Info("cleared users from database")

	return nil
}
