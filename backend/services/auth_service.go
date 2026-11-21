package services

import (
	"fmt"
	"errors"
	"grades-management/models"
	"grades-management/repository"
	"grades-management/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	userRepo repository.UserRepository
	rRepo repository.RTokenRepository
}

func NewAuthService(repo repository.UserRepository,rRepo repository.RTokenRepository) *AuthService  {
	return &AuthService{
		userRepo: repo,
		rRepo: rRepo,
	}
}

func (s *AuthService) Register(user models.User) error {
	
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.Role = "parent"

	s.userRepo.Save(user)

	return nil
}
func (s *AuthService) Login(username string, password string) (string,string,error)  {
	
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "","", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return "","", err
	}

	
	accessToken, _ := utils.GenerateToken(user.UserID,user.Role)
	rfTokenStr, _ := utils.GenerateRandomString(32)

	expiresAt := time.Now().UTC().Add(24*time.Hour)

	err = s.rRepo.SaveRefreshToken(user.UserID,rfTokenStr,expiresAt)
	if err != nil {
		return "","", err
	}

	return accessToken, rfTokenStr, nil
}

func (s *AuthService)RefreshToken(oldToken string)(string, error)  {
	// 1. Cari token di DB
    storedToken, err := s.rRepo.FindRefreshToken(oldToken)
    if err != nil {
        fmt.Printf("Database Error: %v\n", err) 
        return "", errors.New("refresh token not found or expired")
    }

    // 2. Cari user berdasarkan UserID dari token tersebut
    // storedToken.UserID sekarang adalah string ULID
    user, err := s.userRepo.FindByID(storedToken.UserID)
    if err != nil {
        return "", errors.New("user owner of token not found")
    }

    // 3. Generate Access Token Baru
    newAccessToken, err := utils.GenerateToken(user.UserID, user.Role)
    if err != nil {
        return "", err
    }
    
    return newAccessToken, nil
}
