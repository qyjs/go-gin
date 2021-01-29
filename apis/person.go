package apis

import (
	"fmt"
	. "gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexAPI ...
func IndexAPI(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

// AddPersonAPI godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param body body Person true "Main product,policyholder and insured information"
// @Success 200 {object} Person
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {string} 404
// @Failure 500 {string} 500
// @Failure default {string} 500
// @Router /person [post]
func AddPersonAPI(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// GetPersonsAPI ...
func GetPersonsAPI(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	id := c.Param("id")
	person := Person{Id: id}
	// log.Fatal(person)
	person.GetPerson()
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

// GetPersonAPI ...
func GetPersonAPI(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}

// ModPersonAPI ...
func ModPersonAPI(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id := c.Param("id")
	p := Person{Id: id}

	p.GetPerson()
	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.ModPerson()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

// DelPersonAPI ...
func DelPersonAPI(c *gin.Context) {
	p := Person{Id: c.Param("id")}
	ra, err := p.DelPerson()

	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("delete successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}
