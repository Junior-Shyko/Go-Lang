package routes

import (
	"github.com/Junior-Shyko/Go-Lang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup) {

	// r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	// r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	// r.POST("/createUser", userController.CreateUser)
	// r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	// r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	// r.POST("/login", userController.LoginUser)

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

	//Rota de auditoria
	r.POST("/audit", controller.CreateAudit)
	r.GET("getAuditUser/:userId", controller.FindAudit)
}
