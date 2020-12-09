package dao

import (
	"Go000/common"
	"database/sql"

	"github.com/pkg/errors"
)

type userInfo struct {
	name string
}

// User info
type User struct {
	ID   uint64
	Info userInfo
}

// GetUserByID query user info
func GetUserByID(id uint64) (*User, error) {
	//access DB...
	sqlCmd := "select ** from **"
	data, err := queryData(sqlCmd)
	if err == sql.ErrNoRows {
		return nil, errors.Wrapf(common.ErrNotFound, "sql: %s error: %v", sqlCmd, err)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "sql: %s error: %v", sqlCmd, err)
	}
	return data, nil
}

func queryData(sqlCmd string) (*User, error) {
	// user := User{123243, userInfo{"test"}}
	// return &user, nil
	return nil, sql.ErrNoRows
}
