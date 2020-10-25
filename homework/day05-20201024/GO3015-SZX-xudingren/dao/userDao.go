package dao

import "GO3015-SZX-xudingren/model"

type userManagerDao interface {
	Create()
	Get(userId int) (model.User, bool)

}