package glentity

import "time"

type BaseEntity struct {
	CreateDatetime string `json:"createDatetime"`
	UpdateDatetime string `json:"updateDatetime"`
	CreateUserId   int64  `json:"createUserId"`
	UpdateUserId   int64  `json:"updateUserId"`
	Version        int64  `json:"version"`
}

type MasterEntity struct {
	Active            string `json:"active"`
	ActiveDatetime    string `json:"activeDatetime"`
	NonActiveDatetime string `json:"nonActiveDatetime"`
	CreateDatetime    string `json:"createDatetime"`
	UpdateDatetime    string `json:"updateDatetime"`
	CreateUserId      int64  `json:"createUserId"`
	UpdateUserId      int64  `json:"updateUserId"`
	Version           int64  `json:"version"`
}

type BaseEntityTs struct {
	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

type MasterEntityTs struct {
	Active             string `json:"active"`
	ActiveTimestamp    string `json:"activeTimestamp"`
	NonActiveTimestamp string `json:"nonActiveTimestamp"`

	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}
