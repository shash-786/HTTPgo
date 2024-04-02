package server

import (
	"encoding/json"
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
	server_database map[string]UserInfo
}

func New() *Server {
	return &Server{
		server_database: make(map[string]UserInfo),
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

		// TODO Implement The case where POST Comes to a user already PRESENT
		// if _, ok := s.server_database[*u.Name]; ok {
		//
		// }

		s.server_database[*u.Name] = UserInfo{
			email: *u.Email,
			age:   *u.Age,
		}
		log.Printf("Welcome %s", *u.Name)
	}
}

func (s *Server) SearchUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input_name string
		if input_name = c.Param("name"); input_name == "" {
			log.Printf("SearchUser : Could not read name")
			c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Name Query Not Found!",
			})
			return
		}

		u, ok := s.server_database[input_name]
		if !ok {
			log.Printf("SerachUser : Couldn't find name in db")
			c.IndentedJSON(http.StatusNotFound, ok)
			return
		}

		result_email := u.email
		result_age := u.age

		ret := user{
			Name:  &input_name,
			Email: &result_email,
			Age:   &result_age,
		}

		json_ret, err := json.Marshal(ret)
		if err != nil {
			log.Printf("SearchUser : Name Found but couldn't Marshall Response! ")
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		log.Printf("Query Found and Returned for %s : %s %d", input_name, result_email, result_age)
		c.IndentedJSON(http.StatusOK, json_ret)
	}
}

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input_name string
		if input_name = c.Param("name"); input_name == "" {
			log.Printf("DeleteUser : Could not read name")
			c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Name Query Not Found!",
			})
			return
		}

		_, ok := s.server_database[input_name]
		if !ok {
			log.Printf("DeleteUser : Couldn't find name in db")
			c.IndentedJSON(http.StatusNotFound, ok)
			return
		}

		log.Printf("Deleted User %s", input_name)
		delete(s.server_database, input_name)
		c.Status(http.StatusOK)
	}
}

func (s *Server) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input_name string
		if input_name = c.Param("name"); input_name == "" {
			log.Printf("UpdateUser : Could not read name")
			c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Name Query Not Found!",
			})
			return
		}

		usr_info, ok := s.server_database[input_name]
		if !ok {
			log.Printf("UpdateUser : Couldn't find name in db")
			c.IndentedJSON(http.StatusNotFound, ok)
			return
		}

		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			log.Printf("CreateUser : Could Not Read the Request Body!")
			c.IndentedJSON(http.StatusUnsupportedMediaType, err)
			return
		}

		if u.Age != nil {
			usr_info.age = *u.Age
		}

		if u.Email != nil {
			usr_info.email = *u.Email
		}

		s.server_database[input_name] = usr_info
		log.Printf("Updated user %s", input_name)
		c.IndentedJSON(http.StatusOK, gin.H{
			"Message": "Updated User",
		})
	}
}
