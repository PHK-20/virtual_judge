package models

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"crypto/md5"

	"github.com/astaxie/beego/orm"
)

type User_info struct {
	Username     string    `orm:"column(username);pk"`
	Password     string    `orm:"column(password)"`
	Nickname     string    `orm:"column(nickname)"`
	RegisterTime time.Time `orm:"column(register_time)"`
}

func init() {
	orm.RegisterModel(new(User_info))
}

func (item *User_info) itemName() string {
	return "user_info"
}

func (item *User_info) Register() error {
	o := orm.NewOrm()
	item.Password = md5V(item.Password)
	_, err := o.Insert(item)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("username is existed")
	}
	return nil
}

func (item *User_info) GetUser(username string) (*User_info, error) {
	o := orm.NewOrm()
	item.Username = username
	err := o.Read(item)
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return nil, errors.New("user is not existed")
	}
	return item, nil
}
func (item *User_info) Check(username, password string) (*User_info, error) {
	item, err := item.GetUser(username)
	if err != nil {
		return nil, err
	}
	if md5V(password) != item.Password {
		return nil, errors.New("username or password wrong")
	}
	return item, nil
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
