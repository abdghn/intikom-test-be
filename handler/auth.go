package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"intikom-test-be/helper"
	"intikom-test-be/usecase"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler interface {
	Googlelogin(ctx *gin.Context)
	GoogleCallback(ctx *gin.Context)
}

type authHandler struct {
	googleConfig oauth2.Config
	userUsecase  usecase.UserUsecase
}

func NewAuthHandler(googleConfig oauth2.Config, userUsecase usecase.UserUsecase) AuthHandler {
	return &authHandler{googleConfig: googleConfig, userUsecase: userUsecase}
}

func (h *authHandler) Googlelogin(ctx *gin.Context) {
	url := h.googleConfig.AuthCodeURL("randomstate")
	ctx.Redirect(http.StatusSeeOther, url)
}

func (h *authHandler) GoogleCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	if state != "randomstate" {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	code := ctx.Query("code")
	token, err := h.googleConfig.Exchange(context.Background(), code)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Code-Token Exchange Failed")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + token.AccessToken)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "User Data Fetch Failed")
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var j map[string]interface{}

	if err := json.Unmarshal(body, &j); err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.userUsecase.ReadByEmail(fmt.Sprintf("%s", j["email"]))
	if err != nil {
		helper.HandleError(ctx, http.StatusUnauthorized, err.Error())
	}

	tokenString := helper.GenerateToken(user)

	helper.HandleSuccess(ctx, map[string]interface{}{
		"token": tokenString,
	})

}
