package controller

import (
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]entities.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"123123456":{
		Id:			1,
		Name:		"123",
		FollowCount:	0,
		FollowerCount:	0,
		IsFollow: false,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	entities.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	entities.Response
	User entities.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entities.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := entities.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entities.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entities.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entities.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: entities.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: entities.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
