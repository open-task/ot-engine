package engine

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Skill struct {
	User  string `json:"user"`
	Skill string `json:"skill"`
}

// curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func addUserSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.PostForm("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	s1 := Skill{
		User:user,
		Skill:skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func fetchUserSkills(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	dummy := []Skill{
		{
			User:  "u1",
			Skill: "s1",
		},
	}
	c.JSON(http.StatusOK, dummy)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/s1 | jq .
func fetchUserSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)
	s1 := Skill{
		User:user,
		Skill:skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X DELETE http://127.0.0.1:8080/backend/v1/user/u1/skill/s1 | jq .
func deleteUserSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)
	s1 := Skill{
		User:user,
		Skill:skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .
func updateUserSkill(c *gin.Context) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	s1 := Skill{
		User:user,
		Skill:skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -X GET http://127.0.0.1:8080/backend/v1/skills/top?limit=30
func topSkills(c *gin.Context) {
	limit := c.Query("limit")
	log.Printf("limit = %s\n", limit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}
