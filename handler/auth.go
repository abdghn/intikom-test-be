package handler

import (
	"context"
	"intikom-test-be/helper"
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
}

func NewAuthHandler(googleConfig oauth2.Config) AuthHandler {
	return &authHandler{googleConfig: googleConfig}
}

func (h *authHandler) Googlelogin(ctx *gin.Context) {
	url := h.googleConfig.AuthCodeURL("randomstate")

	ctx.Redirect(0, url)
	ctx.JSON(http.StatusSeeOther, map[string]interface{}{url: url})
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

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "User Data Fetch Failed")
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "JSON Parsing Failed")
		return
	}

	helper.HandleSuccess(ctx, string(userData))

}
