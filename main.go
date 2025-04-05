package main

import (
	"net/http"
	"strconv"

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
	router.GET("/v1/dogs", getDogs)
	router.GET("/v1/dogs/:id", getDogByID)
	router.POST("/v1/dogs/:id", updateDogById)
	router.POST("/v1/dogs", postDog)

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

// getDogByID locates the dog whose ID value matches the id
// parameter sent by the client, then returns that dog as a response.
func getDogByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid dog id type"})
	}

	// Loop over the list of dogs, looking for
	// an dog whose ID value matches the parameter.
	for _, a := range dogs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "dog not found"})
}

// updateDogById updates the information for a dog using the JSON presented
// only updates if the dog already exists
func updateDogById(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid dog id type"})
		return
	}

	// Loop over the list of dogs, looking for
	// an dog whose ID value matches the parameter.
	for i, a := range dogs {
		if a.ID == id {
			// update the dog
			var updatedDog dog
			if err := c.BindJSON(&updatedDog); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"messsage": "error binding JSON"})
				return
			}
			dogs[i] = updatedDog
			c.IndentedJSON(http.StatusCreated, updatedDog)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cannot update, dog not found"})
}
