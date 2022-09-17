package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormConfig = &gorm.Config{
	DisableAutomaticPing:   false,
	PrepareStmt:            true,
	SkipDefaultTransaction: false,
}

func GetDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		panic(errors.New("DSN EMPTY"))
	}

	// connect to cockroach
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), gormConfig)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to Cockroach DB")
	}
	return db
}

func InitDB() {
	db := GetDB()
	db.AutoMigrate(StreetEntry{})
}

func DBAddEntry(newEntry *StreetEntry) StreetEntry {
	res := GetDB().Create(newEntry)
	if res.Error != nil {
		panic(res.Error)
	}
	return *newEntry
}

func DBUpdateEntry(entry *StreetEntry) {
	GetDB().Save(entry)
}

func DBGetEntry(id string) (StreetEntry, bool) {
	var ret StreetEntry
	res := GetDB().First(ret, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return ret, false
	} else if res.Error != nil {
		panic(res.Error)
	}
	return ret, true
}
