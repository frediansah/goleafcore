package glutil

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/frediansah/goleafcore"
	"github.com/frediansah/goleafcore/glapi"
	"github.com/frediansah/goleafcore/glconstant"
	"github.com/frediansah/goleafcore/gldata"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ResponseError(err error, prefix ...string) *glapi.OutgoingApi {
	prefixFirst := ""
	if len(prefix) > 0 {
		prefixFirst = prefix[0]
	}

	output := glapi.OutgoingApi{
		Result: glapi.OutgoingApiResult{
			Status:    glconstant.STATUS_FAIL,
			ErrorCode: err.Error(),
			ErrorMsg:  prefixFirst + err.Error(),
			ErrorArgs: []interface{}{},
		},
	}

	if errCore, okCasting := err.(*goleafcore.CoreError); okCasting {
		output.Result.ErrorCode = errCore.ErrorCode
		output.Result.ErrorMsg = prefixFirst + errCore.ErrorMessage
		output.Result.ErrorFullMsg = prefixFirst + errCore.ErrorFullMessage
		output.Result.ErrorArgs = errCore.ErrorArgs
	}

	return &output
}

func ResponseSuccess(payload interface{}) *glapi.OutgoingApi {
	output := glapi.OutgoingApi{
		Result: glapi.OutgoingApiResult{
			Status:    glconstant.STATUS_OK,
			ErrorCode: glconstant.EMPTY_VALUE,
			ErrorMsg:  glconstant.EMPTY_VALUE,
			ErrorArgs: []interface{}{},
			Payload:   payload,
		},
	}

	return &output
}

func FetchAuditData(c *fiber.Ctx) (*gldata.AuditData, error) {
	glclaim, err := FetchGlClaim(c)
	if err != nil {
		return nil, err
	}

	roleLoginId, _ := strconv.Atoi(c.Get(glconstant.HEADER_ROLE_ID, ToString(glclaim.Data.RoleId)))
	tenantLoginId := glclaim.Data.TenantId
	if tenantLoginId == glconstant.TENANT_SUPERADMIN {
		if tenantSuperAdmin, errTenantSuperadmin := strconv.Atoi(c.Get(glconstant.HEADER_TENANT_ID)); errTenantSuperadmin == nil {
			tenantLoginId = int64(tenantSuperAdmin)
		}
	}

	timestamp, err := time.Parse(time.RFC3339, c.Get(glconstant.HEADER_DATETIME))
	if err != nil {
		timestamp = time.Now()
	}

	auditData := gldata.AuditData{
		UserLoginId:   glclaim.Data.UserId,
		RoleLoginId:   int64(roleLoginId),
		TenantLoginId: tenantLoginId,
		Timestamp:     timestamp,
	}

	return &auditData, nil
}

func FetchJwtBody(authorization []byte) goleafcore.Dto {
	raw := string(authorization)
	arrs := strings.Split(raw, " ")
	if len(arrs) > 1 {
		token := arrs[1]
		arrsToken := strings.Split(token, ".")
		if len(arrsToken) > 1 {
			tokenBody := arrsToken[1]
			tokenJsonByte, err := base64.RawURLEncoding.DecodeString(tokenBody)
			if err != nil {
				logrus.Error("Error decode token json :", err)
			} else {
				jwtDto, err := goleafcore.NewDto(tokenJsonByte)
				if err == nil {
					return jwtDto
				}
			}
		}
	}

	return goleafcore.Dto{}
}

func FetchGlClaim(c *fiber.Ctx) (*glapi.GoleafClaim, error) {
	raw := c.Get("Authorization")

	arrs := strings.Split(raw, " ")
	if len(arrs) > 1 {
		token := arrs[1]
		arrsToken := strings.Split(token, ".")
		if len(arrsToken) > 1 {
			tokenBody := arrsToken[1]
			tokenJsonByte, err := base64.RawURLEncoding.DecodeString(tokenBody)
			if err != nil {
				return nil, errors.New("error decode base64 :" + err.Error())
			}

			var glclaim glapi.GoleafClaim
			err = json.Unmarshal(tokenJsonByte, &glclaim)
			if err != nil {
				return nil, errors.New("error parse token as glclaim :" + err.Error())
			}

			return &glclaim, nil
		}
	}

	return nil, errors.New("header Authorization not valid")
}
