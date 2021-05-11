package goleafcore

import (
	"database/sql"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/glconstant"
	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glentity"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type VaBillTrxBca struct {
	BillTrxBcaId        int64           `json:"billTrxBcaId" gleaf:"pk,seq"`
	VirtualAccountBcaId int64           `json:"virtualAccountBcaId"`
	TenantId            int64           `json:"tenantId"`
	AccountNo           string          `json:"accountNo"`
	CompanyCode         string          `json:"companyCode"`
	SubcompanyCode      string          `json:"subcompanyCode"`
	AccountName         string          `json:"accountName"`
	BillNumber          string          `json:"billNumber"`
	BillTimestamp       time.Time       `json:"billTimestamp"`
	BillDate            string          `json:"billDate"`
	BillAmount          decimal.Decimal `json:"billAmount"`
	CurrCode            string          `json:"currCode"`
	BillDescriptionEn   string          `json:"billDescriptionEn"`
	BillDescriptionId   string          `json:"billDescriptionId"`
	BillFreeTextsEn     string          `json:"billFreeTextsEn"`
	BillFreeTextsId     string          `json:"billFreeTextsId"`
	RefDocNo            string          `json:"refDocNo"`
	RefDocDate          string          `json:"refDocDate"`
	RefDocId            int64           `json:"refDocId"`
	RefDocTypeId        int64           `json:"refDocTypeId"`
	PaidChannelType     string          `json:"paidChannelType"`
	Remark              string          `json:"remark"`
	StatusPayment       string          `json:"statusPayment"`
	PaymentFlgStatusBca string          `json:"paymentFlgStatusBca"`
	PaidTimestamp       sql.NullTime    `json:"paidTimestamp"`
	CancelTimestamp     sql.NullTime    `json:"cancelTimetamp"`
	BillReferenceBca    string          `json:"billReference"`
	FlgExpired          string          `json:"flgExpired"`
	ExpiredTimestamp    sql.NullTime    `json:"expiredTimestamp"`

	glentity.BaseEntityTs
}

func TestEntityInsertError(t *testing.T) {
	glinit.InitLog()
	glinit.InitDb()

	err := gldb.BeginTrx(func(trx pgx.Tx) error {
		billTrx := VaBillTrxBca{
			TenantId:            10,
			VirtualAccountBcaId: 10,
			AccountNo:           "0823572323",
			CompanyCode:         "14333",
			SubcompanyCode:      "00000",
			AccountName:         "Frediansah",
			BillNumber:          "072364732657",
			BillTimestamp:       time.Now(),
			BillDate:            "20160102",
			BillAmount:          decimal.NewFromInt(100000),
			CurrCode:            "IDR",
			BillDescriptionEn:   "English",
			BillDescriptionId:   "Indonesia",
			BillFreeTextsEn:     "Free En",
			BillFreeTextsId:     "Free Id",
			RefDocNo:            "Ref doc no",
			RefDocDate:          "",
			RefDocId:            232,
			RefDocTypeId:        32532,
			PaidChannelType:     glconstant.EMPTY_VALUE,
			Remark:              "Remark",
			StatusPayment:       "D",
			PaymentFlgStatusBca: glconstant.EMPTY_VALUE,
			PaidTimestamp: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
			CancelTimestamp: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
			BillReferenceBca: glconstant.EMPTY_VALUE,
			FlgExpired:       glconstant.NO,
			ExpiredTimestamp: sql.NullTime{
				Valid: true,
				Time:  time.Now().Add(time.Second * time.Duration(20000)),
			},
			BaseEntityTs: glentity.BaseEntityTs{
				CreateTimestamp: time.Now(),
				UpdateTimestamp: time.Now(),
				CreateUserId:    -1,
				UpdateUserId:    -1,
				Version:         0,
			},
		}

		tableName := initTableTestEntityError(trx)
		errTrx := gldb.InsertTx(trx, &billTrx, tableName)
		logrus.Debug("Error : ", errTrx)

		logrus.Debug("TABLE NAME : ", tableName)
		return nil
	})

	logrus.Debug("Error trx : ", err)
}

func initTableTestEntityError(trx pgx.Tx) string {
	tableName := `t_test_table_entity_` + glutil.DateTimeNow()

	sql := `DROP TABLE IF EXISTS ` + tableName + `; ` +
		`CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
		`	bill_trx_bca_id			bigserial,` +
		`	virtual_account_bca_id	bigint,` +
		`	tenant_id 				bigint,` +
		`	account_no				character varying(200),		` +
		`	company_code			character varying(100),		` +
		`	subcompany_code			character varying(100),		` +
		`	account_name			character varying(500),		` +
		`	bill_number				character varying(250),		` +
		`	bill_timestamp			timestamptz,` +
		`	bill_date				character varying(8),` +
		`	bill_amount				numeric(13,2),` +
		`	curr_code				character varying(3),` +
		`	bill_description_en		character varying(18),` +
		`	bill_description_id		character varying(18),` +
		`	bill_free_texts_en		text,  						` +
		`	bill_free_texts_id		text,  						` +
		`	ref_doc_no				character varying(50) default '',	` +
		`	ref_doc_date			character varying(8) default '',	` +
		`	ref_doc_id				bigint default -99,  				` +
		`	ref_doc_type_id			bigint default -99,  				` +
		`	paid_channel_type		character varying(4),  				` +
		`	remark					text,  								` +
		`	status_payment			character varying(50) default 'D',	` +
		`	payment_flg_status_bca	character varying(2),` +
		`	paid_timestamp			timestamptz,` +
		`	cancel_timestamp		timestamptz,` +
		`	bill_reference_bca		character varying (30),` +
		`	flg_expired				character varying (1) DEFAULT 'N',` +
		`	expired_timestamp		timestamptz,` +
		`	create_timestamp	timestamptz,` +
		`	update_timestamp	timestamptz,` +
		`	create_user_id		bigint,` +
		`	update_user_id		bigint,` +
		`	version				bigint,` +
		`	CONSTRAINT ` + tableName + `_pk PRIMARY KEY (bill_trx_bca_id)` +
		`); `

	err := gldb.ExecTx(trx, sql)
	if err != nil {
		logrus.Error("Error create table :", err)
	}
	return tableName
}
