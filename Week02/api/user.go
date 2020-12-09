package api

import (
	"Go000/common"
	"Go000/dao"
	"Go000/service"
	"errors"
	"log"
)

// HandleQueryUser handle user query
func HandleQueryUser() {
	defer recoverFunc()
	id := uint64(9008000000048942)
	user, err := service.GetUser(id)
	if errors.Is(err, common.ErrNotFound) {
		sendErrMsg(err)
		return
	}
	sendSuccessMsg(user)
}

func recoverFunc() {
	if err := recover(); err != nil {
		sendErrMsg(err.(error))
	}
}

func sendErrMsg(err error) {
	if errors.Is(err, common.ErrNotFound) {
		log.Printf("stack: \n%+v\n", err)
		log.Printf("404 not found, %s", err)
		return
	}
	log.Printf("server error: %+v", err)
}

func sendSuccessMsg(user *dao.User) {
	log.Printf("user found %v", user)
	return
}
