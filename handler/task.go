package handler

import (
	"intikom-test-be/helper"
	"intikom-test-be/model"
	"intikom-test-be/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	ReadAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	ReadById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

func (h *taskHandler) ReadAll(ctx *gin.Context) {
	var err error
	data, err := h.taskUsecase.ReadAll()
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	helper.HandleSuccess(ctx, data)
}

func (h *taskHandler) Create(ctx *gin.Context) {
	var input model.TaskInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}

	newTask, err := h.taskUsecase.Create(input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HandleSuccess(ctx, newTask)

}

func (h *taskHandler) ReadById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	u, err := h.taskUsecase.ReadById(id)
	if err != nil {
		helper.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}
	helper.HandleSuccess(ctx, u)

}

func (h *taskHandler) Update(ctx *gin.Context) {
	var err error
	var input model.TaskInput
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

	u, err := h.taskUsecase.Update(id, &input)
	if err != nil {
		helper.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helper.HandleSuccess(ctx, u)

}

func (h *taskHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}
	err = h.taskUsecase.Delete(id)
	if err != nil {
		helper.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}
	helper.HandleSuccess(ctx, "success delete data")

}
