package controllers

import (
	"fmt"
	"golang_api_kegiatanQ/dto"
	"golang_api_kegiatanQ/helper"
	"golang_api_kegiatanQ/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//BookController is a ...
type FinishController interface {
	UpdateIsFinishHutang(context *gin.Context)
}

type isfinishController struct {
	isfinishService service.KegiatanQFinishService
	jwtService      service.JWTService
}

//NewBookController create a new instances of BoookController
func NewIsFinishController(isfinishServ service.KegiatanQFinishService, jwtServ service.JWTService) FinishController {
	return &isfinishController{
		isfinishService: isfinishServ,
		jwtService:      jwtServ,
	}
}

func (c *isfinishController) UpdateIsFinishHutang(context *gin.Context) {
	var isfinishUpdateDTO dto.IsFinishUpdateDTO
	// fmt.Println("con", isFinishUpdateDTO.Islunas)

	errDTO := context.ShouldBind(&isfinishUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.isfinishService.LunasIsAllowedToEdit(userID, isfinishUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			isfinishUpdateDTO.UserID = id
		}
		result := c.isfinishService.UpdateLunas(isfinishUpdateDTO)
		// fmt.Println("result", result)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
