package engine

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func addSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func fetchAllSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func fetchSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func deleteSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func updateSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}

func topSkills(c *gin.Context) {
	limit := c.Param("limit")
	log.Printf("limit = %d\n", limit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}