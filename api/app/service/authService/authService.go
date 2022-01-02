package authService

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	passwordHelper "dou-survey/app/helpers/password"
	"dou-survey/app/repository/userRepository"
	"dou-survey/internal/logger"

	"github.com/golang-jwt/jwt"
)

//UserServiceInterface define the user service interface methods
type AuthServiceInterface interface {
	TokenBuildAccess(id uint) (tokenString string, err error)
	tokenBuildRefresh(id uint) (tokenString string, err error)
	TokenValidate(tokenString string, secret string) (id uint, err error)
	TokenValidateRefresh(tokenString string) (id uint, err error)
	Logout(id uint, refreshToken string) error
	Login(id string, password string) (refreshToken string, accessToken string, err error)
}

// billingService handles communication with the user repository
type AuthService struct {
	userRepo userRepository.UserRepositoryInterface
	logger   logger.Logger
}

// NewUserService implements the user service interface.
func NewAuthService(userRepo userRepository.UserRepositoryInterface, logger logger.Logger) AuthServiceInterface {
	return &AuthService{
		userRepo,
		logger,
	}
}

// FindByID implements the method to find a user model by primary key
func (s *AuthService) TokenBuildAccess(id uint) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	var idStr string = strconv.FormatUint(uint64(id), 10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    idStr,
		ExpiresAt: time.Now().Add(time.Duration(1) * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := []byte(os.Getenv("ACCESS_SECRET"))
	tokenString, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// FindByID implements the method to find a user model by primary key
func (s *AuthService) tokenBuildRefresh(id uint) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	var idStr string = strconv.FormatUint(uint64(id), 10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    idStr,
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := []byte(os.Getenv("REFRESH_SECRET"))
	tokenString, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	err = s.userRepo.AddSession(id, tokenString)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) TokenValidate(tokenString string, secret string) (id uint, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})
	if err != nil {
		return 0, err
	}

	s.logger.Info("TESTTTTT")

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {

		s.logger.Info("claims: %s", claims)

		u64, err := strconv.ParseUint(claims.Issuer, 10, 32)

		if err != nil {
			return 0, err
		}

		return uint(u64), nil
	} else {
		s.logger.Info("claims: %s", claims)
		return 0, errors.New("token is not valid")
	}
}

func (s *AuthService) TokenValidateRefresh(tokenString string) (id uint, err error) {
	id, err = s.TokenValidate(tokenString, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		return 0, err
	}

	sessions, err := s.userRepo.GetSessions(id)
	if err != nil {
		return 0, err
	}

	split := strings.Split(sessions, "/")

	contains := false
	for _, value := range split {
		if value == tokenString {
			contains = true
			break
		}
	}

	if !contains {
		return 0, errors.New("session is not active")
	}

	return id, nil
}

func (s *AuthService) Logout(id uint, refreshToken string) error {
	err := s.userRepo.RemoveSession(id, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(IDNumber string, password string) (refreshToken string, accessToken string, err error) {
	// FIND USER IF EXISTS
	user, err := s.userRepo.FindByIdNumber(IDNumber)

	if errors.Is(err, sql.ErrNoRows) {
		// USER DOES NOT EXIST
		s.logger.Error(err.Error())
		return "", "", err
	} else if err != nil {
		s.logger.Error(err.Error())
		return "", "", err
	}

	s.logger.Info("USER: ", user)

	matchPassword := passwordHelper.CheckPasswordHash(password, user.Password)

	if !matchPassword {
		err = errors.New("wrong password")
		s.logger.Error(err.Error())
		return "", "", err
	}

	if err != nil {
		s.logger.Error(err.Error())
		return "", "", err
	}

	refreshToken, err = s.tokenBuildRefresh(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return "", "", err
	}

	accessToken, err = s.TokenBuildAccess(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return "", "", err
	}

	return refreshToken, accessToken, nil
}
