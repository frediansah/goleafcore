package glapi

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type GoleafClaim struct {
	jwt.StandardClaims

	Data TokenClaimData `json:"data,omitempty"`
}

type TokenClaimData struct {
	UserId   int64                        `json:"userId,omitempty"`
	RoleId   int64                        `json:"roleId,omitempty"`
	TenantId int64                        `json:"tenantId,omitempty"`
	Username string                       `json:"username,omitempty"`
	Fullname string                       `json:"fullname,omitempty"`
	RoleName string                       `json:"roleName,omitempty"`
	RoleList []TokenClaimDataRoleListItem `json:"roleList,omitempty"`
}

type TokenClaimDataRoleListItem struct {
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	RoleType string `json:"roleType"`
}
