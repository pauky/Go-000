package service

import (
	"Go000/dao"
	"fmt"
)

// GetUser handle rpc request
func GetUser(id uint64) (*dao.User, error) {
	user, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Info: login complete.userID=%d,info=%+v\n", id, user)
	return user, nil
}
