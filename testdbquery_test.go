package goleafcore

import (
	"database/sql"
	"log"
	"testing"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glentity"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/sirupsen/logrus"
)

type User struct {
	glentity.BaseEntity

	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTwo struct {
	UserId            int64         `json:"userId" gleaf:"pk,seq"`
	TenantId          int64         `json:"tenantId"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	Fullname          string        `json:"fullname"`
	Password          string        `json:"password"`
	Phone             string        `json:"phone"`
	RoleDefaultId     int64         `json:"roleDefaultId"`
	PrivateKey        string        `json:"privateKey"`
	CreateDatetime    string        `json:"createDatetime"`
	CreateUserId      int64         `json:"createUserId"`
	UpdateDatetime    string        `json:"updateDatetime"`
	UpdateUserId      int64         `json:"updateUserId"`
	Version           int64         `json:"version"`
	Active            string        `json:"active"`
	ActiveDatetime    string        `json:"activeDatetime"`
	NonActiveDatetime string        `json:"nonActiveDatetime"`
	OuDefaultId       sql.NullInt64 `json:"ouDefaultId"`
	PolicyDefaultId   sql.NullInt64 `json:"policyDefaultId"`
}

// func testEntity(t *testing.T) {
// 	log.Println("Jalankah?")
// 	glinit.InitLog()
// 	glinit.InitDb()

// 	for i := 0; i < 1; i++ {
// 		user2 := genUser()
// 		err := gldb.Insert(user2, "t_user")
// 		logrus.Debug("Apakah error user2 ? ", err)
// 	}
// }

func TestSelect(t *testing.T) {
	glinit.InitLog()
	glinit.InitDb()
	logrus.Debug("START TEST SELECT")

	var resultList []*UserTwo

	err := gldb.Select(&gldb.ReturnSelect{
		Result: &resultList,
	}, `select * from t_user WHERE user_id <> $1`, -1)
	if err != nil {
		log.Panicln("error", err)
	}

	if len(resultList) > 0 {
		log.Println("Usernya? :", resultList[0])
	}

}

// func testExec(t *testing.T) {
// 	glinit.InitLog()
// 	glinit.InitDb()

// 	logrus.Debug("Test EXECUTE UPDATE")

// 	err := gldb.Exec(`UPDATE t_user SET username = 'hahaha' WHERE user_id = $1`, 0)
// 	if err != nil {
// 		log.Panicln("error", err)
// 	}

// }

// func genUser() UserTwo {
// 	return UserTwo{
// 		TenantId:          10,
// 		Username:          uuid.NewString(),
// 		Email:             uuid.NewString() + "@gmail.com",
// 		Fullname:          uuid.NewString(),
// 		Password:          uuid.NewString(),
// 		Phone:             "0823423423523",
// 		RoleDefaultId:     -1,
// 		PrivateKey:        uuid.NewString(),
// 		CreateDatetime:    glutil.DateTimeNow(),
// 		CreateUserId:      glconstant.ROLE_SUPERADMIN,
// 		UpdateDatetime:    glutil.DateTimeNow(),
// 		UpdateUserId:      glconstant.USER_SUPERADMIN,
// 		Version:           0,
// 		Active:            glconstant.YES,
// 		ActiveDatetime:    glutil.DateTimeNow(),
// 		NonActiveDatetime: glconstant.EMPTY_VALUE,
// 		OuDefaultId:       sql.NullInt64{Int64: glconstant.NULL_REF_VALUE_FOR_LONG},
// 		PolicyDefaultId:   sql.NullInt64{Int64: glconstant.NULL_REF_VALUE_FOR_LONG},
// 	}
// }
