package router

import (
	"api/src/controllers"
	// "net/http"
	// "api/src/router/rotas"
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

//Gerar Rotas
// func Gerar() *mux.Router {
// 	r := mux.NewRouter()

// 	return rotas.Configurar(r)
// }

func InitRoutes(r *gin.RouterGroup) {
	// http.ResponseWriter, *http.Request
	// resWrite := r.ResponseWriter
	// var resWrite = r.ResponseWriter
	r.GET("/getAuditId/:auditId")
	r.POST("/audit", controllers.CreateAudit)
	r.PUT("/getAuditId/:auditId")
	r.DELETE("/getAuditId/:auditId")
	r.GET("/getAudit")
}
