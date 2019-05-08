package engine

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Skill struct {
	Id         int64  `json:id`
	User       string `json:"user"`
	Skill      string `json:"skill"`
	Status     int    `json:"status"`
	UpdateTime string `json:"update_time"`
}

// curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func AddUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.PostForm("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	// TODO: check the input format

	stmtIns, err := db.Prepare(`INSERT INTO skill (addr, skill) VALUES(?, ?)`)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(user, skill)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}

	c.JSON(http.StatusOK, Skill{
		Id:     id,
		User:   user,
		Skill:  skill,
		Status: 0,
	})
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func FetchUserSkills(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	log.Printf("user: %s\n", user)
	stmtOuts, err := db.Prepare(`
SELECT id,
       addr,
       skill,
       status,
       updatetime
FROM skill
WHERE filter=0
  AND addr = ?
`)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	defer stmtOuts.Close()
	rows, err := stmtOuts.Query(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	defer rows.Close()

	var skills []Skill

	for rows.Next() {
		var s Skill
		err = rows.Scan(&s.Id, &s.User, &s.Skill, &s.Status, &s.UpdateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		skills = append(skills, s)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	c.JSON(http.StatusOK, skills)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .
func FetchUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	skill := c.Param("skill")
	log.Printf("user: %s, skill: %s\n", user, skill)

	var s Skill
	err := db.QueryRow(`
SELECT id,
       addr,
       skill,
       status,
       updatetime
FROM skill
WHERE id=?
  AND addr = ?
`, skill, user).Scan(&s.Id, &s.User, &s.Skill, &s.Status, &s.UpdateTime)
	if err != nil {
		// TODO: empty set
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	c.JSON(http.StatusOK, s)
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
