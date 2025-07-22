package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
)

type AuthController struct {
	Repo repositories.UserRepository
}

func NewAuthController(repo repositories.UserRepository) *AuthController {
	return &AuthController{Repo: repo}
}

func (a *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		helpers.Error(ctx, http.StatusBadRequest, err.Error(), "Invalid request")
		return
	}
	
	user, err := a.Repo.GetByEmail(req.Email)
	if err != nil {
		helpers.Error(ctx, http.StatusUnauthorized, "Unauthorized", "Email tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		helpers.Error(ctx, http.StatusUnauthorized, "Unauthorized", "Password salah")
		return
	}

	accessToken, accessExp := helpers.GenerateJWT(user.ID.String(), "access-token", 15*time.Minute)
	refreshToken, refreshExp := helpers.GenerateJWT(user.ID.String(), "refresh-token", 1*24*time.Hour)

	ipAddr := ctx.ClientIP()
	// Simpan token ke DB
	err = a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		Nama:      ipAddr,
		IDUser:    user.ID,
		Token:     accessToken,
		Type:      "access-token",
		ExpiredAt: accessExp,
	})

	if err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to store access token")
		return
	}

	// Simpan refresh token ke DB
	err =a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		Nama:      ipAddr,
		IDUser:    user.ID,
		Token:     refreshToken,
		Type:      "refresh-token",
		ExpiredAt: refreshExp,
	})

	if err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to store refresh token")
		return
	}

	// Set cookie
	ctx.SetCookie(
		"refresh_token",      						// cookie name
		refreshToken,         						// value
		int(time.Until(refreshExp).Seconds()),  	// maxAge dalam detik (1 hari)
		"/",                  						// path
		"",          								// domain
		false,                 						// secure (true jika pakai HTTPS)
		true,                 						// httpOnly (tidak bisa diakses dari JS)
	)

	helpers.Success(ctx, http.StatusOK, map[string]any{
		"user":          map[string]any{"id": user.ID, 
										"email": user.Email, 
										"nama": user.Nama, 
										"role": user.Role, 
										"data_pelanggan": user.DataPelanggan},
		"access_token":  accessToken,
	}, "Success")
}

func (a *AuthController) Register(ctx *gin.Context) {
	var req struct {
		Nama     string `form:"nama" validate:"required,min=2,max=100"`
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required,min=6,max=255"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		helpers.Error(ctx, http.StatusBadRequest, err.Error(), "Invalid request")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
        helpers.Error(ctx, http.StatusUnprocessableEntity, errors, "Validasi gagal")
        return
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to hash password")
		return
	}

	user := models.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := a.Repo.Create(&user); err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to create user")
		return
	}

	accessToken, accessExp := helpers.GenerateJWT(user.ID.String(), "access-token", 15*time.Minute)
	refreshToken, refreshExp := helpers.GenerateJWT(user.ID.String(), "refresh-token", 1*24*time.Hour)

	ipAddr := ctx.ClientIP()
	
	// Simpan token ke DB
	err = a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		Nama:      ipAddr,
		IDUser:    user.ID,
		Token:     accessToken,
		Type:      "access-token",
		ExpiredAt: accessExp,
	})

	if err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to store access token")
		return
	}

	// Simpan refresh token ke DB
	err =a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		Nama:      ipAddr,
		IDUser:    user.ID,
		Token:     refreshToken,
		Type:      "refresh-token",
		ExpiredAt: refreshExp,
	})

	if err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to store refresh token")
		return
	}

	// Set cookie
	ctx.SetCookie(
		"refresh_token",      						// cookie name
		refreshToken,         						// value
		int(time.Until(refreshExp).Seconds()),  	// maxAge dalam detik (1 hari)
		"/",                  						// path
		"",          								// domain
		false,                 						// secure (true jika pakai HTTPS)
		true,                 						// httpOnly (tidak bisa diakses dari JS)
	)

	helpers.Success(ctx, http.StatusOK, map[string]any{
		"user":          map[string]any{"id": user.ID, 
										"email": user.Email, 
										"nama": user.Nama, 
										"role": user.Role, 
										"data_pelanggan": user.DataPelanggan},
		"access_token"	: accessToken,
	}, "Success")
}

func (a *AuthController) Refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		helpers.Error(ctx, http.StatusUnauthorized, err.Error(), "Refresh token not found")
		return
	}

	claims, err := helpers.VerifyJWT(refreshToken)
	if err != nil || claims.Type != "refresh-token" {
		helpers.Error(ctx, http.StatusUnauthorized, err.Error(), "Refresh token invalid")
		return
	}

	_, err = a.Repo.FindValidRefreshToken(refreshToken)
	if err != nil {
		helpers.Error(ctx, http.StatusUnauthorized, "Not found", "Refresh token tidak valid")
		return
	}

	if err := a.Repo.RevokeAllUserTokens(claims.UserID); err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to revoke all user tokens")
		return
	}

	newAccessToken, newAccessExp := helpers.GenerateJWT(claims.UserID, "access-token", 15*time.Minute)
	newRefreshToken, newRefreshExp := helpers.GenerateJWT(claims.UserID, "refresh-token", 1*24*time.Hour)
	
	ipAddr := ctx.ClientIP()

	// Simpan token ke DB
	a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		IDUser:    uuid.MustParse(claims.UserID),
		Nama:      ipAddr,
		Token:     newAccessToken,
		Type:      "access-token",
		ExpiredAt: newAccessExp,
	})

	// Simpan refresh token ke DB
	a.Repo.StoreAccessToken(&models.PersonalAccessToken{
		IDUser:    uuid.MustParse(claims.UserID),
		Nama:      ipAddr,
		Token:     newRefreshToken,
		Type:      "refresh-token",
		ExpiredAt: newRefreshExp,
	})

	// Set cookie
	ctx.SetCookie(
		"refresh_token",      						// cookie name
		newRefreshToken,         					// value
		int(time.Until(newRefreshExp).Seconds()),  	// maxAge dalam detik (1 hari)
		"/",                  						// path
		"",          								// domain
		false,                 						// secure (true jika pakai HTTPS)
		true,                 						// httpOnly (tidak bisa diakses dari JS)
	)

	helpers.Success(ctx, http.StatusOK, map[string]string{
		"access_token":  newAccessToken,
	}, "Success Generate New Token")
}

func (a *AuthController) Logout(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		helpers.Error(ctx, http.StatusUnauthorized, err.Error(), "Token not found")
		return
	}
	claims, err := helpers.VerifyJWT(refreshToken)
	if err != nil || claims.Type != "refresh-token" {
		helpers.Error(ctx, http.StatusUnauthorized, err.Error(), "Token invalid")
		return
	}

	// Set revoke_at ke semua token milik user
	if err := a.Repo.RevokeAllUserTokens(claims.UserID); err != nil {
		helpers.Error(ctx, http.StatusInternalServerError, err.Error(), "Failed to revoke all user tokens")
		return
	}

	// Hapus refresh token cookie
	ctx.SetCookie(
		"refresh_token",
		"",             // kosongkan value
		-1,             // hapus
		"/",
		"",    			// domain
		false,          // secure
		true,           // httpOnly
	)

	helpers.Success(ctx, http.StatusOK, nil, "Success Logout")
}
