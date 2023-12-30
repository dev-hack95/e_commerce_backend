package models

import (
	"time"

	"github.com/beego/beego/orm"
	"github.com/spf13/cast"
)

type Users struct {
	Id        int    `orm:"column(id);auto"`
	FirstName string `orm:"column(first_name);"`
	LastName  string `orm:"column(last_name);"`
	Email     string `orm:"column(email);"`
	Password  string `orm:"column(password);"`
	UserToken string `orm:"column(token);"`
	//IsSeller  bool      `orm:"column(IsSeller);"`
	CreatedAt time.Time `orm:"column(created_at);"`
	UpdateAt  time.Time `orm:"column(updated_at)"`
}

func (t *Users) TableName() string {
	return "user_details"
}

func init() {
	orm.RegisterModel(new(Users))
}

func GetUserByEmail(email string) (u *Users, err error) {
	o := orm.NewOrm()

	user := &Users{Email: email}
	err = o.Read(user, "Email")

	switch {
	case err == orm.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return user, nil
}

func AddUserDetails(u *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(u)
	return
}

func GetUserByEmailAndPassWord(email string, password string) (u *Users, err error) {
	o := orm.NewOrm()
	u = &Users{Email: cast.ToString(email), Password: cast.ToString(password)}
	if err := o.Read(u); err == nil {
		return u, nil
	}
	return nil, err
}
