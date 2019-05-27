package engine

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/open-task/ot-engine/types"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'email=u1@a.com&skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
func AddUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("address")
	if address == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address."})
		return
	}
	user := types.User{Address: address}
	if err := db.FirstOrCreate(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var skill types.Skill
	if err := c.ShouldBind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		skill.UpdateTime = nil //clear fake update time
	}

	db.FirstOrCreate(&skill, skill)
	user.Skills = append(user.Skills, skill)
	db.Save(&user) // for statements
	c.JSON(http.StatusOK, skill)
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b/skill' | jq .
func FetchUserSkills(c *gin.Context, db *gorm.DB) {
	address := c.Param("address")
	user := types.User{Address: address}
	var skills []types.Skill
	if err := db.First(&user, user).Error; err != nil {
		c.JSON(http.StatusOK, [0]types.Skill{}) // empty list
		return
	}

	db.Model(&user).Association("Skills").Find(&skills)
	c.JSON(http.StatusOK, skills)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .
func FetchUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("address")
	user := types.User{Address: address}
	if err := db.First(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	skill := types.Skill{Id: id}
	db.First(&skill, skill) // don't create

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
	address := c.Param("address")
	user := types.User{Address: address}
	if err := db.First(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	log.Printf("user: %s, skillId: %s\n", address, idStr)

	skill := types.Skill{Id: id}
	db.Model(&user).Association("Skills").Delete(&skill) //don't check for exist
	// always ok
	c.JSON(http.StatusOK, skill)
}

// curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'email=user1@bountinet.com&skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .
func UpdateUserSkill(c *gin.Context, db *gorm.DB) {
	address := c.Param("user")
	idStr := c.Param("skill")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	email := c.PostForm("email")
	user := types.User{Address: address}
	if email != "" {
		user.Email = email
		// update user if email changes
		db.Model(&user).Update(user)
	}
	// can't change skill desc ?
	tag := c.PostForm("skill")
	var skill = types.Skill{Id: id}
	if tag != "" {
		skill.Tag = tag
		// update skill if desc changes
		db.Model(&skill).Update(skill)
		log.Printf("update skill, id: %s, new skill is '%s'\n", id, tag)
	}
	status := c.PostForm("status")
	submitNum := c.PostForm("submit_num")
	confirmNum := c.PostForm("confirm_num")
	filter := c.PostForm("filter")

	query := "UPDATE skill SET ";
	var values []interface{}
	var state types.Statement
	if status != "" {
		i, err := strconv.Atoi(status)
		if err != nil {
			query += "status = ?, "
			values = append(values, i)
			state.Status = i
		}
	}

	if submitNum != "" {
		i, err := strconv.Atoi(submitNum)
		if err != nil {
			query += "submit_num = ?, "
			values = append(values, i)
			state.Submit = i
		}
	}

	if confirmNum != "" {
		i, err := strconv.Atoi(confirmNum)
		if err != nil {
			query += "confirm_num = ?, "
			values = append(values, i)
			state.Confirm = i
		}
	}

	if filter != "" {
		i, err := strconv.Atoi(filter)
		if err != nil {
			query += "filter = ?, "
			values = append(values, filter)
			state.Filter = i
		}
	}

	if len(values) == 0 {
		c.JSON(http.StatusOK, skill)
		return
	}
	query = strings.Trim(query, ", ")
	query += " WHERE user_id = ? AND skill_id = ?"
	values = append(values, user.Id, skill.Id)
	log.Printf("query = \"%s\", values = %v", query, values)

	if err := db.Exec(query, values...).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
		return
	}

	c.JSON(http.StatusOK, skill)
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

func checkLimit(limitStr string) (int64, error) {
	if limitStr == "" {
		return 0, errors.New("empty ID.")
	}
	limit, err := strconv.ParseInt(limitStr, 0, 64)
	if err != nil {
		return 0, errors.New("Invalid ID.")
	}
	return limit, nil
}

// curl -s -X GET '127.0.0.1:8080/backend/v1/user/0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b/info' | jq .
func FetchUserInfo(c *gin.Context, db *gorm.DB, engineDB *sql.DB) {
	address := c.Param("address")
	user := types.User{Address: address}
	if err := db.First(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Association("Skills").Find(&user.Skills)
	getMissionSummaryForOne(engineDB, &user)
	c.JSON(http.StatusOK, user)
}

//curl -v -X POST \
//  http://127.0.0.1:8080/backend/v1/user/111/info \
//  -H 'content-type: application/x-www-form-urlencoded' \
//  -d 'email=user111@bountinet.com'
//curl -v -X POST \
//  http://127.0.0.1:8080/backend/v1/user/111/info \
//  -H 'content-type: application/json' \
//  -d '{ "email": "user111@bountinet.com" }'
func UpdateUserInfo(c *gin.Context, db *gorm.DB) {
	address := c.Param("address")
	user := types.User{Address: address}
	if err := db.FirstOrCreate(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user2 types.User
	if err := c.ShouldBind(&user2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user2.UpdateTime = nil
	db.First(&user, user)
	db.Model(&user).Update(user2)
	c.JSON(http.StatusOK, user)
}

func FetchUserMissions(c *gin.Context, db *gorm.DB) {

}

func FetchSkills(c *gin.Context, db *gorm.DB) {
	var skill types.Skill
	if err := c.ShouldBind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	skill.UpdateTime = nil

	limitStr := c.DefaultQuery("limit", "30")
	limit, err := checkLimit(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var skills []types.Skill
	db.Where(skill).Order("confirm desc").Limit(limit).Find(&skills)
	c.JSON(http.StatusOK, skills)
}

//func (c *gin.Context, db *gorm.DB) {
//
//}

func UpdateSkills(c *gin.Context, db *gorm.DB) {

}

func GetSkills(c *gin.Context, db *gorm.DB) {

}

func DeleteSkills(c *gin.Context, db *gorm.DB) {

}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/list_skills | jq .
// curl -s -X GET http://127.0.0.1:8080/backend/v1/list_skills?limit=2 | jq .
func ListSkills(c *gin.Context, backendDB *gorm.DB, engineDB *sql.DB) {
	limitStr := c.DefaultQuery("limit", "30")
	limit, err := checkLimit(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var skills []types.Skill
	backendDB.Order("confirm desc").Limit(limit).Find(&skills)
	for i, s := range skills {
		skills[i].UserNumber = backendDB.Model(s).Association("users").Count()
	}
	c.JSON(http.StatusOK, skills)
}

// curl -s -X GET http://127.0.0.1:8080/backend/v1/list_users?skill_id=2 | jq .
func GetUsers(c *gin.Context, backendDB *gorm.DB, engineDB *sql.DB) {
	idStr := c.Query("id")
	id, err := checkId(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	skill := types.Skill{Id: id}
	var users []types.User
	backendDB.Model(&skill).Association("Users").Find(&users)
	getMissionSummary(engineDB, users)
	c.JSON(http.StatusOK, users)
}

func AddSkill(c *gin.Context, backendDB *gorm.DB, engineDB *sql.DB) {
	address := c.PostForm("address")
	if address == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address."})
		return
	}
	user := types.User{Address: address}
	if err := backendDB.FirstOrCreate(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var skill types.Skill
	if err := c.ShouldBind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		skill.UpdateTime = nil //clear fake update time
	}

	if err := backendDB.FirstOrCreate(&skill, skill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Skills = append(user.Skills, skill)
	backendDB.Save(&user) // for statements
	c.JSON(http.StatusOK, skill)
}

func getMissionSummary(db *sql.DB, users []types.User) {
	if len(users) == 0 {
		return
	}
	pos := make(map[string]int, len(users))
	var addresses []string
	for i, u := range users {
		pos[u.Address] = i
		addresses = append(addresses, u.Address)
	}
	if len(addresses) == 0 {
		return
	}

	addr_list := strings.Join(addresses, "','")
	addr_list = "'" + addr_list + "'"
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	query := fmt.Sprintf(`
SELECT publisher,
       count(publisher),
       MAX(txtime)
FROM mission
WHERE publisher IN (%s)
GROUP BY publisher;
`, addr_list)
	rows1, err := db.Query(query)
	if err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	defer rows1.Close()

	for rows1.Next() {
		var address string
		var count int64
		var txTimeStr string

		if err1 := rows1.Scan(&address, &count, &txTimeStr); err1 != nil {
			log.Println(err1)
			continue
		}
		users[pos[address]].MissionSummary.Publish = count

		txTime, err2 := time.Parse(layout, txTimeStr)
		if err2 != nil {
			log.Println(err2)
		}
		if users[pos[address]].MissionSummary.LastActive == nil || users[pos[address]].MissionSummary.LastActive.Before(txTime) {
			users[pos[address]].MissionSummary.LastActive = &txTime
		}
	}
	if err = rows1.Err(); err != nil {
		log.Fatal(err)
	}

	query = fmt.Sprintf(`
SELECT solver,
       count(solver),
       MAX(txtime)
FROM solution
WHERE solver IN (%s)
GROUP BY solver;
`, addr_list)
	rows2, err := db.Query(query)
	if err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	defer rows2.Close()

	for rows2.Next() {
		var address string
		var count int64
		var txTimeStr string

		if err1 := rows2.Scan(&address, &count, &txTimeStr); err1 != nil {
			log.Println(err1)
			continue
		}
		users[pos[address]].MissionSummary.Submit = count

		txTime, err2 := time.Parse(layout, txTimeStr)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		if users[pos[address]].MissionSummary.LastActive == nil || users[pos[address]].MissionSummary.LastActive.Before(txTime) {
			users[pos[address]].MissionSummary.LastActive = &txTime
		}
	}
	if err = rows2.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetUserInfo(c *gin.Context, backendDB *gorm.DB, engineDB *sql.DB) {
	address := c.Query("address")
	user := types.User{Address: address}
	if err := backendDB.First(&user, user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	backendDB.Model(&user).Association("Skills").Find(&user.Skills)
	getMissionSummaryForOne(engineDB, &user)
	c.JSON(http.StatusOK, user)
}

func getMissionSummaryForOne(db *sql.DB, user *types.User) {
	if user.Address == "" {
		return
	}
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"

	query := `
SELECT count(publisher),
       MAX(txtime)
FROM mission
WHERE publisher = ?;
`
	rows1, err := db.Query(query, user.Address)
	if err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	defer rows1.Close()

	for rows1.Next() {
		var count int64
		var txTimeStr string

		if err1 := rows1.Scan(&count, &txTimeStr); err1 != nil {
			log.Println(err1)
			continue
		}
		user.MissionSummary.Publish = count

		txTime, err2 := time.Parse(layout, txTimeStr)
		if err2 != nil {
			log.Println(err2)
		}
		if user.MissionSummary.LastActive == nil || user.MissionSummary.LastActive.Before(txTime) {
			user.MissionSummary.LastActive = &txTime
		}
	}
	if err = rows1.Err(); err != nil {
		log.Fatal(err)
	}

	query = `
SELECT count(solver),
       MAX(txtime)
FROM solution
WHERE solver = ?;
`
	rows2, err := db.Query(query, user.Address)
	if err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	defer rows2.Close()

	for rows2.Next() {
		var count int64
		var txTimeStr string

		if err1 := rows2.Scan(&count, &txTimeStr); err1 != nil {
			log.Println(err1)
			continue
		}
		user.MissionSummary.Submit = count

		txTime, err2 := time.Parse(layout, txTimeStr)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		if user.MissionSummary.LastActive == nil || user.MissionSummary.LastActive.Before(txTime) {
			user.MissionSummary.LastActive = &txTime
		}
	}
	if err = rows2.Err(); err != nil {
		log.Fatal(err)
	}

	query = `
SELECT sum(reward)
FROM mission
WHERE publisher = ?
  AND solved=1;
`
	var rewardStr sql.NullString
	if err := db.QueryRow(query, user.Address).Scan(&rewardStr); err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	if rewardInDET, success := big.NewFloat(0).SetString(rewardStr.String); success {
		rewardInDET.Quo(rewardInDET, types.Decimals)
		user.MissionSummary.PaidRewardDET = rewardInDET
	}
	//
	//	query = `
	//SELECT mission_id
	//FROM solution
	//WHERE solver = ?;
	//`

}
