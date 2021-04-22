package goleafcore

type BaseEntity struct {
	CreateDatetime string `json:"createDatetime"`
	UpdateDatetime string `json:"updateDatetime"`
	CreateUserId   int64  `json:"createUserId"`
	UpdateUserId   int64  `json:"updateUserId"`
	Version        int64  `json:"version"`
}
