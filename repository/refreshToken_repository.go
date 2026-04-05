package repository

import(	
	"grades-management/models"
	"github.com/jmoiron/sqlx"

	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type RTokenRepository interface {
	SaveRefreshToken(userID string, token string, expiresAt time.Time) error
	FindRefreshToken(oldToken string) (*models.RefreshToken,error)
}

type rTokenRepo struct {
	db *sqlx.DB
}

func NewRTokenRepository(db *sqlx.DB) RTokenRepository {
	return &rTokenRepo{db: db}
}

func (r *rTokenRepo)SaveRefreshToken(userID string, token string, expiresAt time.Time) error {
	_ , _ = r.db.Exec("DELETE FROM refresh_token WHERE user_id =$1 OR expires_at < $2",userID,time.Now().UTC())

	_,err:= r.db.Exec("INSERT INTO refresh_token (user_id,refresh_token,expires_at,created_at) VALUES ($1,$2,$3,$4)",userID,token,expiresAt,time.Now().UTC())

	return err
}

func (r *rTokenRepo)FindRefreshToken(oldToken string)( *models.RefreshToken,error)  {
	var rf models.RefreshToken

	err := r.db.Get(&rf,"SELECT user_id, refresh_token, expires_at FROM refresh_token WHERE refresh_token = $1 AND expires_at > $2 LIMIT 1", oldToken,time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return &rf, nil
}