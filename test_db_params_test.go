package goleafcore_test

import (
	"log"
	"testing"

	"github.com/frediansah/goleafcore/gldb"
)

func TestTestQParams(t *testing.T) {
	p := gldb.Qparams{}
	var addedKeyword bool
	var keyword string = "aku adalah"

	query := `SELECT table_id, name, COALESCE(status_payment, ` + p.New("DRAFT") + `) AS status_payment ` +
		` FROM t_apa ` +
		` WHERE tenant_id = ` + p.New("tenantId") +
		`  	AND active = ` + p.New("active") +
		` 	AND status_payment = ` + p.New("DRAFT") +
		gldb.QAddIfNotEmpty(&addedKeyword, keyword, ` AND A.remark ILIKE  `+p.New("keyword"))

	log.Println("Query : ", query)

	p.Set("active", "Y")
	p.Set("tenantId", 10)
	p.Set("DRAFT", "D")
	err := p.Set("raono", "Y")
	if err != nil {
		log.Println("error set params : ", err)
	}
	if addedKeyword {
		p.Set("keyword", gldb.QWrapLikeBoth(keyword))
	}

	log.Println("Param values ", p.GetValues())
}
