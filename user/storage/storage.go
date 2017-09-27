package storage

import (
	"github.com/xuebing1110/fortify/pkg/wechat"
	"github.com/xuebing1110/fortify/user/models"
)

type Storage interface {
	SaveSession(sess_3rd string, sessInfo *wechat.SessionResp) error
	QuerySession(sess_3rd string) (*wechat.SessionResp, error)

	UpsertUser(user models.User) error
	AddUser(user models.User) error
	Exist(uid string) bool
}
