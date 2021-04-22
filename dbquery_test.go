package goleafcore

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/stsgoleaf/goleafcore/glconstant"
	"gitlab.com/stsgoleaf/goleafcore/gldb"
	"gitlab.com/stsgoleaf/goleafcore/glinit"
	"gitlab.com/stsgoleaf/goleafcore/glutil"
)

type User struct {
	BaseEntity

	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTwo struct {
	UserId            int64  `json:"userId" gleaf:"pk,seq"`
	TenantId          int64  `json:"tenantId"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Fullname          string `json:"fullname"`
	Password          string `json:"password"`
	Phone             string `json:"phone"`
	RoleDefaultId     int64  `json:"roleDefaultId"`
	PrivateKey        string `json:"privateKey"`
	CreateDatetime    string `json:"createDatetime"`
	CreateUserId      int64  `json:"createUserId"`
	UpdateDatetime    string `json:"updateDatetime"`
	UpdateUserId      int64  `json:"updateUserId"`
	Version           int64  `json:"version"`
	Active            string `json:"active"`
	ActiveDatetime    string `json:"activeDatetime"`
	NonActiveDatetime string `json:"nonActiveDatetime"`
	OuDefaultId       int64  `json:"ouDefaultId"`
	PolicyDefaultId   int64  `json:"policyDefaultId"`
}

func testEntity(t *testing.T) {
	log.Println("Jalankah?")
	glinit.InitLog()
	glinit.InitDb()

	user := User{
		BaseEntity: BaseEntity{
			CreateDatetime: "20210102021010",
			UpdateDatetime: "20210102021010",
			CreateUserId:   -1,
			UpdateUserId:   -1,
			Version:        0,
		},
		Username: uuid.New().String(),
		Password: "sts123",
	}

	gldb.Insert(user, "t_user")

	for i := 0; i < 1; i++ {
		user2 := genUser()
		err := gldb.Insert(user2, "t_user")
		logrus.Debug("Apakah error user2 ? ", err)
	}
}

func genUser() UserTwo {
	return UserTwo{
		TenantId:          10,
		Username:          uuid.NewString(),
		Email:             uuid.NewString() + "@gmail.com",
		Fullname:          uuid.NewString(),
		Password:          uuid.NewString(),
		Phone:             "0823423423523",
		RoleDefaultId:     -1,
		PrivateKey:        uuid.NewString(),
		CreateDatetime:    glutil.DateTimeNow(),
		CreateUserId:      glconstant.ROLE_SUPERADMIN,
		UpdateDatetime:    glutil.DateTimeNow(),
		UpdateUserId:      glconstant.USER_SUPERADMIN,
		Version:           0,
		Active:            glconstant.YES,
		ActiveDatetime:    glutil.DateTimeNow(),
		NonActiveDatetime: glconstant.EMPTY_VALUE,
		OuDefaultId:       glconstant.NULL_REF_VALUE_FOR_LONG,
		PolicyDefaultId:   glconstant.NULL_REF_VALUE_FOR_LONG,
	}
}
