package controller

import (
	"awesomeProject/infrastructure/utils"
	"awesomeProject/services/userChatService"
	"awesomeProject/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type userChatController struct {
	userChatUseCase usecase.UserChat
}

type UserChat interface {
	ServeWS(ctx *gin.Context)
	GetUserMessagesHistory(ctx *gin.Context)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewUserChatController(us *usecase.UserChat) UserChat {
	return &userChatController{userChatUseCase: *us}
}

func (u userChatController) GetUserMessagesHistory(ctx *gin.Context) {
	userIdQuery := ctx.Query("user_id")
	recipientIdQuery := ctx.Query("recipient_id")
	if len(userIdQuery) == 0 || len(recipientIdQuery) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, "no required user_id or recipient_id")
		return
	}
	var err error
	var (
		userId      int
		recipientId int
	)

	userId, err = strconv.Atoi(userIdQuery)
	recipientId, err = strconv.Atoi(recipientIdQuery)

	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	history, err := u.userChatUseCase.GetUserMessagesHistory(userId, recipientId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func (u userChatController) ServeWS(ctx *gin.Context) {
	roomID := ctx.Param("roomID")

	hub := u.userChatUseCase.UserChatHub()

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := userChatService.NewClient(roomID, conn, hub)
	hub.Register <- client
	go client.WritePump()
	go client.ReadPump()
}
