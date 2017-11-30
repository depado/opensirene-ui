package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Depado/opensirene-ui/conf"
)

// GetUI returns the UI
func GetUI(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"root": conf.C.OpenSireneAPI})
}

// GetUIData returns UI data
func GetUIData(c *gin.Context) {
	data := c.Param("data")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"data": data, "root": conf.C.OpenSireneAPI})
}
