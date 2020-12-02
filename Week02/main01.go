package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

func main() {
	err := Biz()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("error: %v", errors.Cause(err))
		fmt.Printf("stack: \n%+v\n", err)
	}
}

// User info
type User struct {
	ID   uint64
	Info struct{}
}

// Biz handle rpc request
func Biz() error {
	id := uint64(9008000000048942)
	user, err := Dao(id)
	if err != nil {
		return err
	}

	fmt.Printf("Info: login complete.userID=%d,info=%+v\n", id, user)
	return nil
}

// Dao query user info
func Dao(id uint64) (*User, error) {
	//access DB...
	err := sql.ErrNoRows
	return nil, errors.Wrapf(err, "query user info has error.userID=%d", id)
}
