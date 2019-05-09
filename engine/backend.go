package engine

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Skill struct {
	Id         int64  `json:id`
	User       string `json:"user,omitempty"`
	Skill      string `json:"skill,omitempty"`
	Status     int    `json:"status,omitempty"`
	Submit     int    `json:"submit,omitempty"`
	Confirm    int    `json:"confirm,omitempty"`
	Filter     int    `json:"filter,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
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
	skillId := c.Param("skill")
	log.Printf("user: %s, skillId: %s\n", user, skillId)

	stmtIns, err := db.Prepare(`
DELETE
FROM skill
WHERE id=?
  AND addr=?;
`)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	defer stmtIns.Close()

	res, err := stmtIns.Exec(skillId, user)
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
		Id:   id,
		User: user,
	})
}

// curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .
func UpdateUserSkill(c *gin.Context, db *sql.DB) {
	user := c.Param("user")
	id, err := checkId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	skill := c.PostForm("skill")
	status := c.PostForm("status")
	submitNum := c.PostForm("submit_num")
	confirmNum := c.PostForm("confirm_num")
	filter := c.PostForm("filter")
	log.Printf("user: %s, skill: %s\n", user, skill)

	query := "UPDATE skill SET ";
	var s = Skill{
		User: user,
	}
	var values []interface{}

	if skill != "" {
		query += "skill = ?, "
		values = append(values, skill)
		s.Skill = skill
	}

	if status != "" {
		i, err := strconv.Atoi(status)
		if err != nil {
			query += "status = ?, "
			values = append(values, i)
			s.Status = i
		}
	}

	if submitNum != "" {
		i, err := strconv.Atoi(submitNum)
		if err != nil {
			query += "submit_num = ?, "
			values = append(values, i)
			s.Submit = i
		}
	}

	if confirmNum != "" {
		i, err := strconv.Atoi(confirmNum)
		if err != nil {
			query += "confirm_num = ?, "
			values = append(values, i)
			s.Confirm = i
		}
	}

	if filter != "" {
		i, err := strconv.Atoi(filter)
		if err != nil {
			query += "filter = ?, "
			values = append(values, filter)
			s.Filter = i
		}
	}

	if len(values) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no valid data"})
		return
	}
	query = strings.Trim(query, ", ")
	query += " WHERE id = ?"
	values = append(values, id)
	log.Printf("query = \"%s\", values = %v", query, values)

	result, err := db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	if n <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no this data."})
		return
	}
	id2, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	s.Id = id2

	c.JSON(http.StatusOK, s)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/skill/top?limit=30 | jq .
func TopSkills(c *gin.Context, db *sql.DB) {
	limit := c.Query("limit")
	log.Printf("limit = %s\n", limit)

	query := `
SELECT skill,
       count(skill) AS providers
FROM skill
WHERE filter=0
GROUP BY skill
ORDER BY providers DESC
LIMIT ?;
`
	rows, err := db.Query(query, limit)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	type Stat struct {
		Skill     string `json:"skill"`
		Providers int    `json:"providers"`
	}
	var stats []Stat

	for rows.Next() {
		var s Stat
		err = rows.Scan(&s.Skill, &s.Providers)
		if err != nil {
			log.Println(err)
			continue
		}
		stats = append(stats, s)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func checkId(c *gin.Context) (int64, error) {
	idStr := c.Param("skill")
	if idStr == "" {
		return 0, errors.New("empty ID.")
	}
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return 0, errors.New("Invalid ID.")
	}
	return id, nil
}
