package models

import "time"

type RefreshToken struct{
	TokenID int `db:"id" json:"token_id"`
	UserID string `db:"user_id" json:"user_id"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	ExpiresAT time.Time `db:"expires_at" json:"expires_at"`
}