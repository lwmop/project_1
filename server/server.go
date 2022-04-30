package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_1/models"
	"project_1/utils"
)

func CountFrequency(c *gin.Context) {
	args := new(models.TextInput)
	c.ShouldBind(args)
	mostOccurrenceWordsMap := utils.CountTestBase(args.Text)
	c.JSON(http.StatusOK, gin.H{"dat": mostOccurrenceWordsMap})
}

word
