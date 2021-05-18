package goleafcore_test

import (
	"log"
	"testing"

	"github.com/frediansah/goleafcore"
	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glinit"
)

type UserTestQueryBuilderAllFeature struct {
	UserId   int64
	Username string
	TenantId int64
	Version  int64
}

func TestQueryBuilderAllFeature(t *testing.T) {
	db := glinit.InitDb()
	defer db.Close()

	var keyword string = "sup"
	var active string = "Y"

	q := gldb.QBuilder{}
	q.Add(" SELECT user_id, tenant_id, username, version ").
		Add(" FROM t_user ").
		Add(" WHERE tenant_id = :tenantId ").
		AddIfNotEmpty(keyword, " AND username ILIKE :username ").
		AddIfNotEmpty(active, " AND active = :active ")

	q.SetParam("tenantId", int64(-1))
	q.SetParam("username", gldb.QWrapLikeBoth(keyword))
	q.SetParam("active", active)

	log.Println("Query akhir :", q.Query())
	params, errParams := q.GetParamValues()
	var result []*UserTestQueryBuilderAllFeature

	if errParams != nil {
		log.Println("error params :", errParams)
		t.Fail()
	}

	log.Println("Execute query...")
	err := gldb.Select(&gldb.ReturnSelect{
		Result: &result,
	}, q.Query(), params...)
	if err != nil {
		log.Println("error : ", err)
	}

	log.Println("Isis : ", goleafcore.Dto{}.Put("dataList", result).ToJsonString())

}

func TestQueryBuilderNoParams(t *testing.T) {
	q := gldb.QBuilder{}
	q.Add("SELECT * FROM t_user")

	log.Println("Query akhir :", q.Query())
}
func TestQueryBuilderParamNotSet(t *testing.T) {
	q := gldb.QBuilder{}
	q.Add("SELECT * FROM t_user WHERE tenant_id = :tenantId ")

	log.Println("Query akhir :", q.Query())
	params, err := q.GetParamValues()
	if err != nil {
		log.Println("error values  :", err)
	}
	log.Println("params value  :", params)
}
