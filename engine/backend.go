package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	fmt.Println(user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func fetchAllSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	fmt.Println(user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func fetchSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	fmt.Println(user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func deleteSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	fmt.Println(user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func updateSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	fmt.Println(user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}
