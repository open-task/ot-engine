package engine

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/open-task/ot-engine/types"
	"log"
	"net/http"
	"strconv"
)

// curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func AddUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("user")
	skill_ := c.PostForm("skill")
	log.Printf("user: %s, skill: %s\n", address, skill_)

	// TODO: check the input format

	skill := types.Skill{Skill: skill_}
	db.FirstOrCreate(&skill, skill)
	user := types.User{Address: address}
	db.FirstOrCreate(&user, user)
	user.Skills = append(user.Skills, skill)
	db.Save(&user)
	c.JSON(http.StatusOK, skill)
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func FetchUserSkills(c *gin.Context, db *gorm.DB) {
	address := c.Param("user")
	log.Printf("user: %s\n", address)

	user := types.User{Address: address}
	db.FirstOrCreate(&user, user)

	var skills []types.Skill
	db.Model(&user).Association("Skills").Find(&skills)
	c.JSON(http.StatusOK, skills)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .
func FetchUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("user")
	idStr := c.Param("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	skill := types.Skill{Id: id}
	db.First(&skill, skill) // don't create

	user := types.User{Address: address}
	db.First(&user, user) // don't create
	var skills []types.Skill
	db.Model(&user).Association("Skills").Find(&skills)
	// make sure user has this skill
	for _, s := range skills {
		if s.Id == skill.Id {
			c.JSON(http.StatusOK, s)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"msg": "no this skill."})
}

// curl -s -X DELETE http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .
func DeleteUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("user")
	idStr := c.Param("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	log.Printf("user: %s, skillId: %s\n", address, idStr)

	skill := types.Skill{Id: id}
	user := types.User{Address: address}
	db.First(&user, user)                                // don't create
	db.Model(&user).Association("Skills").Delete(&skill) //don't check for exist
	// always ok
	c.JSON(http.StatusOK, skill)
}

// curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .
func UpdateUserSkill(c *gin.Context, db *gorm.DB) {
	//user := c.Param("user")
	//idStr := c.Param("skill")
	//id, err := checkId(idStr)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	//	return
	//}
	//skill := c.PostForm("skill")
	//status := c.PostForm("status")
	//submitNum := c.PostForm("submit_num")
	//confirmNum := c.PostForm("confirm_num")
	//filter := c.PostForm("filter")
	//log.Printf("user: %s, skill: %s\n", user, skill)
	//
	//query := "UPDATE skill SET ";
	//var s = Skill{
	//	User: user,
	//}
	//var values []interface{}
	//
	//if skill != "" {
	//	query += "skill = ?, "
	//	values = append(values, skill)
	//	s.Skill = skill
	//}
	//
	//if status != "" {
	//	i, err := strconv.Atoi(status)
	//	if err != nil {
	//		query += "status = ?, "
	//		values = append(values, i)
	//		s.Status = i
	//	}
	//}
	//
	//if submitNum != "" {
	//	i, err := strconv.Atoi(submitNum)
	//	if err != nil {
	//		query += "submit_num = ?, "
	//		values = append(values, i)
	//		s.Submit = i
	//	}
	//}
	//
	//if confirmNum != "" {
	//	i, err := strconv.Atoi(confirmNum)
	//	if err != nil {
	//		query += "confirm_num = ?, "
	//		values = append(values, i)
	//		s.Confirm = i
	//	}
	//}
	//
	//if filter != "" {
	//	i, err := strconv.Atoi(filter)
	//	if err != nil {
	//		query += "filter = ?, "
	//		values = append(values, filter)
	//		s.Filter = i
	//	}
	//}
	//
	//if len(values) == 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "no valid data"})
	//	return
	//}
	//query = strings.Trim(query, ", ")
	//query += " WHERE id = ?"
	//values = append(values, id)
	//log.Printf("query = \"%s\", values = %v", query, values)
	//
	//result, err := db.Exec(query, values...)
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
	//	return
	//}
	//n, err := result.RowsAffected()
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
	//	return
	//}
	//if n <= 1 {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "no this data."})
	//	return
	//}
	//id2, err := result.LastInsertId()
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
	//	return
	//}
	//s.Id = id2
	//
	//c.JSON(http.StatusOK, s)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/skill/top?limit=30 | jq .
func TopSkills(c *gin.Context, db *gorm.DB) {
	limit := c.DefaultQuery("limit", "30")
	log.Printf("limit = %s\n", limit)

	query := `
SELECT id,
       skill,
       providers
FROM
  (SELECT skill_id,
          count(skill_id) AS providers
   FROM statements
   WHERE filter=0
   GROUP BY skill_id) AS st
INNER JOIN skills AS sk ON st.skill_id = sk.id
ORDER BY providers DESC
LIMIT ?;
	`
	rows, err := db.Raw(query, limit).Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}
	type Stat struct {
		Id        int64  `json:"id"`
		Skill     string `json:"skill"`
		Providers int    `json:"providers"`
	}
	var stats []Stat

	for rows.Next() {
		var s Stat
		err = rows.Scan(&s.Id, &s.Skill, &s.Providers)
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

func FetchSkillProviders(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	skill := types.Skill{Id: id}
	var users []types.User
	db.Model(&skill).Association("Users").Find(&users)
	c.JSON(http.StatusOK, users)
}

func checkId(idStr string) (int64, error) {
	if idStr == "" {
		return 0, errors.New("empty ID.")
	}
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return 0, errors.New("Invalid ID.")
	}
	return id, nil
}
