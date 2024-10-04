package usecase

import (
	"FitnessCenter_GoBackEnd/auth"
	"FitnessCenter_GoBackEnd/constants"
	"FitnessCenter_GoBackEnd/dtos"
	"FitnessCenter_GoBackEnd/models"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthUseCase struct {
	userRepo           auth.UserRepository
	refreshSessionRepo auth.RefreshSessionRepository
}

func NewAuthUseCase(userRepo auth.UserRepository, refreshSessionRepo auth.RefreshSessionRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo:           userRepo,
		refreshSessionRepo: refreshSessionRepo,
	}
}

type payload struct {
	clientId uuid.UUID
	email    string
}

func (a *AuthUseCase) SignUp(ctx context.Context, signUpDTO dtos.SignUpDTO) (string, string, error) {

	user := new(models.User)

	_, err := a.userRepo.GetUserByEmail(signUpDTO.Email)
	if err == nil {
		return "", "", auth.ErrUserAlreadyExists
	}

	user.ID = uuid.New()
	hashedPassword, err := HashPassword(signUpDTO.Password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return "", "", err
	}
	user.PasswordHash = hashedPassword
	user.Role = constants.ROLES.Client
	user.FirstName = signUpDTO.FirstName
	user.LastName = signUpDTO.LastName
	user.Email = signUpDTO.Email

	client, err := a.userRepo.CreateUser(*user)
	if err != nil {
		return "", "", err
	}

	payload := payload{client.ID, client.Email}

	var jwtSecret = []byte(viper.GetString("jwtSecret"))
	accessToken, err := GenerateAccessToken(payload, jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(payload, jwtSecret)
	if err != nil {
		return "", "", err
	}

	var refreshSession models.RefreshSession
	refreshSession.ID = uuid.New()
	refreshSession.UserID = client.ID
	refreshSession.RefreshToken = refreshToken
	refreshSession.FingerPrint = signUpDTO.FingerPrint

	_, err = a.refreshSessionRepo.CreateRefreshSession(refreshSession)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func (a *AuthUseCase) SignIn(ctx context.Context, signInDTO dtos.SignInDTO) (*models.User, string, string, error) {
	user, err := a.userRepo.GetUserByEmail(signInDTO.Email)
	if err != nil {
		return nil, "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(signInDTO.Password)); err != nil {
		return nil, "", "", auth.InvalidPassword
	}

	payload := payload{user.ID, user.Email}
	var jwtSecret = []byte(viper.GetString("jwtSecret"))
	accessToken, err := GenerateAccessToken(payload, jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := GenerateRefreshToken(payload, jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	var refreshSession models.RefreshSession
	refreshSession.ID = uuid.New()
	refreshSession.UserID = user.ID
	refreshSession.RefreshToken = refreshToken
	refreshSession.FingerPrint = signInDTO.FingerPrint

	_, err = a.refreshSessionRepo.CreateRefreshSession(refreshSession)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (a *AuthUseCase) LogOut(refreshToken string) error {
	return a.refreshSessionRepo.DeleteRefreshSession(refreshToken)
}

func (a *AuthUseCase) Refresh(fingerprint string, refreshToken string) (*models.User, string, string, error) {
	refreshSession, err := a.refreshSessionRepo.GetRefreshSession(refreshToken)
	if err != nil {
		return nil, "", "", err
	}

	if fingerprint != refreshSession.FingerPrint {
		return nil, "", "", errors.New("bad fingerprint")
	}

	if err := a.refreshSessionRepo.DeleteRefreshSession(refreshToken); err != nil {
		return nil, "", "", err
	}

	token, err := verifyRefreshToken(refreshToken)
	if err != nil {
		log.Fatalf("Token verification failed: %v", err)
		return nil, "", "", err
	}

	var payload *payload = new(payload)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Access specific claim data
		if id, ok := claims["id"].(string); ok {
			payload.clientId, _ = uuid.Parse(id)
		}

		if email, ok := claims["email"].(string); ok {
			payload.email = email
		}

		// You can access other claims here
	} else {
		log.Println("Invalid JWT token")
	}

	var jwtSecret = []byte(viper.GetString("jwtSecret"))

	var accessTokenNew, refreshTokenNew string

	accessTokenNew, err = GenerateAccessToken(*payload, jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	refreshTokenNew, err = GenerateRefreshToken(*payload, jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	var refreshSessionNew models.RefreshSession
	refreshSessionNew.ID = uuid.New()
	refreshSessionNew.UserID = payload.clientId
	refreshSessionNew.RefreshToken = refreshTokenNew
	refreshSessionNew.FingerPrint = fingerprint

	_, err = a.refreshSessionRepo.CreateRefreshSession(refreshSessionNew)
	if err != nil {
		return nil, "", "", err
	}

	user, err := a.userRepo.GetUserByEmail(payload.email)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessTokenNew, refreshTokenNew, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateAccessToken(payload payload, jwtSecret []byte) (string, error) {
	// Create the JWT claims
	claims := jwt.MapClaims{
		"id":    payload.clientId,
		"email": payload.email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(), // Access token expires in 15 minutes
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(payload payload, jwtSecret []byte) (string, error) {
	// Create the JWT claims
	claims := jwt.MapClaims{
		"id":    payload.clientId,
		"email": payload.email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	return token.SignedString(jwtSecret)
}

func verifyRefreshToken(refreshToken string) (*jwt.Token, error) {
	// Secret key
	secret := []byte(viper.GetString("jwtSecret"))

	// Verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
