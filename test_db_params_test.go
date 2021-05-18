package goleafcore_test

import (
	"log"
	"testing"

	"github.com/frediansah/goleafcore/gldb"
)

func TestTestQParams(t *testing.T) {
	p := gldb.Qparams{}

	query := `SELECT table_id, name, COALESCE(status_payment, ` + p.New("DRAFT") + `) AS status_payment ` +
		` FROM t_apa ` +
		` WHERE tenant_id = ` + p.New("tenantId") +
		`  AND active = ` + p.New("active")

	log.Println("Query : ", query)

	p.Set("active", "Y")
	p.Set("tenantId", 10)
	p.Set("DRAFT", "D")
	err := p.Set("raono", "Y")
	if err != nil {
		log.Println("error set params : ", err)
	}

	log.Println("Param values ", p.GetValues())
}
