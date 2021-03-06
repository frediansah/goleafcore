package glentity

import (
	"time"

	"github.com/mattn/go-nulltype"
)

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

type BaseEntityMigrateTs struct {
	CreateDatetime  string    `json:"createDatetime"`
	UpdateDatetime  string    `json:"updateDatetime"`
	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

type MasterEntityTs struct {
	Active             string            `json:"active"`
	ActiveTimestamp    nulltype.NullTime `json:"activeTimestamp"`
	NonActiveTimestamp nulltype.NullTime `json:"nonActiveTimestamp"`

	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

type MasterEntityMigrateTs struct {
	ActiveDatetime    string `json:"activeDatetime"`
	NonActiveDatetime string `json:"nonActiveDatetime"`
	CreateDatetime    string `json:"createDatetime"`
	UpdateDatetime    string `json:"updateDatetime"`

	Active             string            `json:"active"`
	ActiveTimestamp    nulltype.NullTime `json:"activeTimestamp"`
	NonActiveTimestamp nulltype.NullTime `json:"nonActiveTimestamp"`

	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

type Entity interface {
	TableName() string
}
