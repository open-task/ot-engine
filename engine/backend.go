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

// curl -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/s1
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

// curl -X GET http://127.0.0.1:8080/backend/v1/skills/top?limit=30
func topSkills(c *gin.Context) {
	limit := c.Query("limit")
	log.Printf("limit = %s\n", limit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}
