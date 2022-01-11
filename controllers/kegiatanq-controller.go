package controllers

import (
	"fmt"
	"golang_api_kegiatanQ/dto"
	"golang_api_kegiatanQ/entity"
	"golang_api_kegiatanQ/helper"
	"golang_api_kegiatanQ/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//BookController is a ...
type KegiatanQController interface {
	All(context *gin.Context)
	AllUser(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type kegiatanqController struct {
	kegiatanqService service.KegiatanQService
	jwtService       service.JWTService
}

//NewBookController create a new instances of BoookController
func NewKegiatanQController(kegiatanqServ service.KegiatanQService, jwtServ service.JWTService) KegiatanQController {
	return &kegiatanqController{
		kegiatanqService: kegiatanqServ,
		jwtService:       jwtServ,
	}
}

func (c *kegiatanqController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	fmt.Println(id)
	var kegiatanqs []entity.KegiatanQ = c.kegiatanqService.All(id)
	res := helper.BuildResponse(true, "OK", kegiatanqs)
	context.JSON(http.StatusOK, res)

}

func (c *kegiatanqController) AllUser(context *gin.Context) {
	// authHeader := context.GetHeader("Authorization")
	// token, err := c.jwtService.ValidateToken(authHeader)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// id := fmt.Sprintf("%v", claims["user_id"])
	// fmt.Println(id)
	var kegiatanqs []entity.User = c.kegiatanqService.AllUser()
	res := helper.BuildResponse(true, "OK", kegiatanqs)
	context.JSON(http.StatusOK, res)

}
func (c *kegiatanqController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var kegiatan entity.KegiatanQ = c.kegiatanqService.FindByID(id)
	if (kegiatan == entity.KegiatanQ{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", kegiatan)
		context.JSON(http.StatusOK, res)
	}
}

func (c *kegiatanqController) Insert(context *gin.Context) {
	var kegiatanqCreateDTO dto.KegiatanQCreateDTO
	errDTO := context.ShouldBind(&kegiatanqCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			kegiatanqCreateDTO.UserID = convertedUserID
		}
		result := c.kegiatanqService.Insert(kegiatanqCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *kegiatanqController) Update(context *gin.Context) {
	var kegiatanqUpdateDTO dto.KegiatanQUpdateDTO
	errDTO := context.ShouldBind(&kegiatanqUpdateDTO)
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
	if c.kegiatanqService.IsAllowedToEdit(userID, kegiatanqUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			kegiatanqUpdateDTO.UserID = id
		}
		result := c.kegiatanqService.Update(kegiatanqUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *kegiatanqController) Delete(context *gin.Context) {
	var hutang entity.KegiatanQ
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	hutang.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.kegiatanqService.IsAllowedToEdit(userID, hutang.ID) {
		c.kegiatanqService.Delete(hutang)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *kegiatanqController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
