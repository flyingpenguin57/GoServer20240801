package service

import (
	"errors"
	"fmt"
	"goserver/controller/request"
	"goserver/dao/model"
	"time"
)

var db_users = make([]model.User, 0, 10)

func Register(request *request.RegisterRequest) {
	//query by user name
	//query by email
	for _, db_user  := range db_users {
		if (db_user.Username == request.Username) {
			panic("username exist")
		}
		if (db_user.Email == request.Email) {
			panic("email exist")
		}
	}

	var id uint = 0;
	if (len(db_users) > 0) {
		id = db_users[len(db_users)-1].Id + 1;
	}

	db_users = append(db_users, model.User{
		Id: id,
		Username: request.Username,
		Email: request.Email,
		Phone: request.Phone,
		Password: request.Password,
		CreatedAt: time.Now(),
	})


    for _, db_user := range db_users {
        fmt.Println(db_user)
    }
}

func Login(request *request.LoginRequest) (map[string]string, error) {

	for _, db_user  := range db_users {
		if (db_user.Email == request.Email) {
			if (request.Password != db_user.Password) {
				return nil, errors.New("password error");
			} else {
				//return token
				return map[string]string{
					"token":  "123456",
				}, nil
			}
		}
	}

	//golang中错误只能一层一层向外传递，这可真抽象...
	return nil, errors.New("username not exist");

}