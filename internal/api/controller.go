package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/learn-video/streaming-platform/internal/model"
	"github.com/learn-video/streaming-platform/internal/service"
)

type InputController struct {
	inputHandler service.InputHandler
}

type NotificationController struct {
	notificationHandler service.NotificationHandler
}

func NewInputController(inputHandler service.InputHandler) *InputController {
	return &InputController{inputHandler: inputHandler}
}

func (c *InputController) CreateInput(ctx *gin.Context) {
	var inputData model.Input
	if err := ctx.BindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := c.inputHandler.CreateInput(ctx, &inputData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, input)
}

func (c *InputController) GetInput(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	input, err := c.inputHandler.GetInput(ctx, uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func NewNotificationController(nh service.NotificationHandler) *NotificationController {
	return &NotificationController{
		notificationHandler: nh,
	}
}

func (n *NotificationController) EnqueuePackaging(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := n.notificationHandler.PackageStream(id); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
