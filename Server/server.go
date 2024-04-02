package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name  *string `json:"name"`
	Age   *uint8  `json:"age"`
	Email *string `json:"email"`
}

type UserInfo struct {
	age   uint8
	email string
}

type Server struct {
	m map[string]UserInfo
}

func New() *Server {
	return &Server{
		m: make(map[string]UserInfo),
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			log.Printf("CreateUser : Could Not Read the Request Body!")
			c.IndentedJSON(http.StatusUnsupportedMediaType, err)
			return
		}

		if u.Name == nil || *u.Name == "" || u.Age == nil {
			log.Printf("CreateUser : Does not have name field!")
			c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Primary Key Empty",
			})
			return
		}

		if u.Email == nil {
			*u.Email = ""
		}

		s.m[*u.Name] = UserInfo{
			email: *u.Email,
			age:   *u.Age,
		}
		log.Printf("Welcome %s", *u.Name)
	}
}
