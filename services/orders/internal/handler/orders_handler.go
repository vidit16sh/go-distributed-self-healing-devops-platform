package handler 
import( 
	"net/http"

	"github.com/gin-gonic/gin"
) 

func HealthHandler(c *gin.Context) { // c is standard variable and *gin.Context refer to the request/response context 
	c.JSON(http.StatusOK, gin.H{
		"status": "order service running",
	})
}
