package engine

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Skill struct {
	User  string `json:"user"`
	Skill string `json:"skill"`
}

// curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func AddUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.PostForm("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	s1 := Skill{
		User:  user,
		Skill: skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func FetchUserSkills(c *gin.Context, db *sql.DB) {
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
func FetchUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)
	s1 := Skill{
		User:  user,
		Skill: skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X DELETE http://127.0.0.1:8080/backend/v1/user/u1/skill/s1 | jq .
func DeleteUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)
	s1 := Skill{
		User:  user,
		Skill: skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .
func UpdateUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	s1 := Skill{
		User:  user,
		Skill: skill,
	}
	c.JSON(http.StatusOK, s1)
}

// curl -X GET http://127.0.0.1:8080/backend/v1/skills/top?limit=30
func TopSkills(c *gin.Context, db *sql.DB) {
	limit := c.Query("limit")
	log.Printf("limit = %s\n", limit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "It works.",
	})
}
