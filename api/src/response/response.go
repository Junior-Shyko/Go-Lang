package response

import (
	"log"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//FORMATANDO JSON PARA RESPOSTA
func JSON(c *gin.Context, statusCode int, dados interface{}) {
	c.ShouldBindHeader(statusCode)

	if err := json.NewEncoder(c.Writer).Encode(dados); err != nil {
		log.Fatal(err)
	}

}