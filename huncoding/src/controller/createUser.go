package controller

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {}

// func (uc *userControllerInterface) CreateUser(c *gin.Context) {
// 	logger.Info("Init CreateUser controller",
// 		zap.String("journey", "createUser"),
// 	)
// 	var userRequest request.UserRequest

// 	if err := c.ShouldBindJSON(&userRequest); err != nil {
// 		logger.Error("Error trying to validate user info", err,
// 			zap.String("journey", "createUser"))
// 		errRest := validation.ValidateUserError(err)

// 		c.JSON(errRest.Code, errRest)
// 		return
// 	}

// 	domain := model.NewUserDomain(
// 		userRequest.Email,
// 		userRequest.Password,
// 		userRequest.Name,
// 		userRequest.Age,
// 	)
// 	domainResult, err := uc.service.CreateUserServices(domain)
// 	if err != nil {
// 		logger.Error(
// 			"Error trying to call CreateUser service",
// 			err,
// 			zap.String("journey", "createUser"))
// 		c.JSON(err.Code, err)
// 		return
// 	}

// 	logger.Info(
// 		"CreateUser controller executed successfully",
// 		zap.String("userId", domainResult.GetID()),
// 		zap.String("journey", "createUser"))

// 	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
// 		domainResult,
// 	))
// }
