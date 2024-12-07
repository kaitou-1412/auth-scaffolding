package routes

import (
	"auth/jwt"
	"auth/models"
	"auth/types"
	"auth/utils"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func signup(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user *models.User

		// Request JSON Validation
		err := context.ShouldBindJSON(&user)
		if err != nil {
			slog.Error(ErrAuthSignup["Request"])
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": ErrAuthSignup["Request"],
				})
			return
		}

		// Username: 5-50 characters, alphanumeric
		if len(user.Username) < 5 || len(user.Username) > 50 || !utils.IsAlphanumeric(user.Username) {
			slog.Error(ErrAuthSignup["Username"])
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": ErrAuthSignup["Username"],
				})
			return
		}

		// Email: Must be valid email format
		if !utils.IsEmail(user.Email) {
			slog.Error(ErrAuthSignup["Email"])
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": ErrAuthSignup["Email"],
				})
			return
		}

		// Username or email already exists.
		if user.UsernameExists(db) || user.EmailExists(db) {
			slog.Error(ErrAuthSignup["Conflict"])
			context.JSON(
				http.StatusConflict,
				gin.H{
					"status":  http.StatusConflict,
					"message": ErrAuthSignup["Conflict"],
				})
			return
		}

		// Password: Minimum 8 characters, must contain at least one uppercase letter, one lowercase letter, one number and one special character
		if !utils.IsValidPassword(user.Password) {
			slog.Error(ErrAuthSignup["Password"])
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": ErrAuthSignup["Password"],
				})
			return
		}

		// Hash the password and store it
		if err := user.HashPassword(); err != nil {
			slog.Error(http.StatusText(http.StatusInternalServerError))
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": http.StatusText(http.StatusInternalServerError),
				})
			return
		}

		// Create user entry in database
		tx := db.Create(&user)
		if tx.Error != nil {
			slog.Error(ErrAuthSignup["Create"])
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": ErrAuthSignup["Create"],
				})
			return
		}

		// User successfully created
		slog.Info(SuccessAuthSignup)
		context.JSON(
			http.StatusCreated,
			gin.H{
				"status":  http.StatusCreated,
				"message": SuccessAuthSignup,
				"data": types.SignupResponse{
					UserId: user.ID,
					Role:   user.Role,
				},
			})
	}
}

func login(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req *types.LoginRequest

		// Request JSON Validation
		err := context.ShouldBindJSON(&req)
		if err != nil {
			slog.Error(ErrAuthLogin["Request"])
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": ErrAuthLogin["Request"],
				})
			return
		}

		// Invalid credentials
		var user = models.User{Username: req.Username}
		if !user.UsernameExists(db) {
			slog.Error(ErrAuthLogin["Credentials"], "Invalid username")
			context.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": ErrAuthLogin["Credentials"],
				})
			return
		}
		user.GetUserByUsername(db)
		if !user.CheckPasswordHash(req.Password) {
			slog.Error(ErrAuthLogin["Credentials"], "Invalid password")
			context.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": ErrAuthLogin["Credentials"],
				})
			return
		}

		// Account locked or disabled
		if !user.Active {
			slog.Error(ErrAuthLogin["Inactive"])
			context.JSON(
				http.StatusForbidden,
				gin.H{
					"status":  http.StatusForbidden,
					"message": ErrAuthLogin["Inactive"],
				})
			return
		}

		// Generate access and refresh tokens
		tokens, err := jwt.GenerateTokenPair(user.ID, user.Role)
		if err != nil {
			slog.Error(ErrAuthLogin["GenerateToken"])
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": ErrAuthLogin["GenerateToken"],
				})
			return
		}

		// Successful login
		slog.Info(SuccessAuthLogin)
		context.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": SuccessAuthLogin,
				"data": types.LoginResponse{
					AccessToken:  tokens.AccessToken,
					RefreshToken: tokens.RefreshToken,
					ExpiresIn:    3600,
					User: types.SignupResponse{
						UserId: user.ID,
						Role:   user.Role,
					},
				},
			})
	}
}

func refresh(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Verify Refresh Token
		refreshToken := context.GetHeader("Authorization")
		claims, ok := jwt.VerifyTokenAndGetClaims(refreshToken, jwt.RefreshToken)
		if !ok {
			slog.Error(ErrAuthRefresh["Token"])
			context.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": ErrAuthRefresh["Token"],
				})
			return
		}

		// Rotating refresh tokens are generally recommended for improved security
		tokens, err := jwt.GenerateTokenPair(claims.UserId, claims.Role)
		if err != nil {
			slog.Error(http.StatusText(http.StatusInternalServerError))
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": http.StatusText(http.StatusInternalServerError),
				})
			return
		}

		// Successful Refresh
		slog.Info(SuccessAuthRefresh)
		context.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": SuccessAuthRefresh,
				"data": types.RefreshResponse{
					AccessToken:  tokens.AccessToken,
					RefreshToken: tokens.RefreshToken,
					ExpiresIn:    3600,
				},
			})
	}
}
