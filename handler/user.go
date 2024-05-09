package handler

import (
	"intikom-test-be/helper"
	"intikom-test-be/model"
	"intikom-test-be/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	ReadAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	ReadById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (h *userHandler) ReadAll(ctx *gin.Context) {
	var err error
	data, err := h.userUsecase.ReadAll()
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	helper.HandleSuccess(ctx, data)

}

func (h *userHandler) Create(ctx *gin.Context) {
	var input model.UserInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	newUser, err := h.userUsecase.Create(input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HandleSuccess(ctx, newUser)
}

func (h *userHandler) ReadById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	u, err := h.userUsecase.ReadById(id)
	if err != nil {
		helper.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}
	helper.HandleSuccess(ctx, u)

}

func (h *userHandler) Update(ctx *gin.Context) {
	var err error
	var input model.UserInput
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	u, err := h.userUsecase.Update(id, &input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helper.HandleSuccess(ctx, u)

}

func (h *userHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}
	err = h.userUsecase.Delete(id)
	if err != nil {
		helper.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}
	helper.HandleSuccess(ctx, "success delete data")

}
