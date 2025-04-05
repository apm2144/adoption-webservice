package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// dog represents data about an adoptable dog
type dog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Breed       string `json:"breed"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

// dogs slice to seed adoptable dog data
var dogs = []dog{
	{ID: 1, Name: "Ginger", Gender: "Female", Breed: "Golden Retriever", Age: 2, Description: "Such a good girl!"},
	{ID: 2, Name: "Fido", Gender: "Male", Breed: "Mixed", Age: 1, Description: "Rescued from an abandoned building."},
	{ID: 3, Name: "Sweaters", Gender: "Female", Breed: "French Bulldog", Age: 3, Description: "Adorable meatball."},
	{ID: 4, Name: "Othello", Gender: "Male", Breed: "Standard Poodle", Age: 6, Description: "A little gray in the muzzle."},
}

func main() {
	router := gin.Default()
	router.GET("/dogs", getDogs)
	router.POST("/dogs", postDog)

	router.Run("localhost:8080")
}

// getDogs responds with the list of all the adoptable dogs as JSON.
func getDogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dogs)
}

// postDog adds a dog from JSON received in the request body.
func postDog(c *gin.Context) {
	var newDog dog

	// Call BindJSON to bind the received JSON to
	// newDog.
	if err := c.BindJSON(&newDog); err != nil {
		return
	}

	// Add the new dog to the in-memory slice.
	dogs = append(dogs, newDog)
	c.IndentedJSON(http.StatusCreated, newDog)
}
