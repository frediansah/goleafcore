package gldata

import (
	"time"

	"github.com/frediansah/goleafcore/glconstant"
)

type AuditData struct {
	UserLoginId   int64
	TenantLoginId int64
	RoleLoginId   int64
	Timestamp     time.Time
}

func (auditData AuditData) Datetime() string {
	return auditData.Timestamp.Format(glconstant.DATETIME_FORMAT)
}
