package repository

import (
	"grades-management/models"
	"github.com/jmoiron/sqlx"
	"log"

	"math/rand"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/oklog/ulid/v2"
)

type UserRepository interface {
	FindAll() []models.User
	FindByUsername(username string)(*models.User,error)
	FindByID(userID string)(*models.User,error)
	Save(users models.User)
	Update(id string, user models.User) error
	Delete(id string) error
	
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}
func (r *userRepo) FindAll() []models.User {
	var users []models.User
	err := r.db.Select(&users,"SELECT id,username, name FROM users")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return users
}

func (r *userRepo) FindByUsername(username string) (*models.User, error) {
    var user models.User

    err := r.db.Get(&user, "SELECT id,username, name FROM users WHERE username = $1", username)

    if err != nil {
        return nil, err 
    }
    return &user, nil
}

func (r *userRepo) FindByID(userID string) (*models.User, error) {
    var user models.User

    err := r.db.Get(&user, "SELECT id,username, name FROM users WHERE id = $1", userID)

    if err != nil {
        return nil, err 
    }
    return &user, nil
}

func (r *userRepo) Save(user models.User) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	user.UserID = id //masukin id ke sql

	_, err := r.db.NamedExec("INSERT INTO users(id, username, name, password, role) VALUES (:id,:username,:name,:password,:role)", 
	user,
)
	if err != nil {
		log.Println("GAGAL nambah database: ", err)
	}
}

func (r *userRepo) Update(id string, user models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, username = $2, name = $3, password = $4 WHERE id= $5", user.Name, user.Username, user.Name,user.Password, id)
	return err
}

func (r *userRepo) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

