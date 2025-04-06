package main

import (
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"

	docs "Users/alanmcnaney/adoption-webservice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/v1/dogs", getDogs)
	router.GET("/v1/dogs/:id", getDogByID)
	router.POST("/v1/dogs/:id", updateDogById)
	router.DELETE("v1/dogs/:id", adoptADog)
	router.POST("/v1/dogs", postDog)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}

// getDogs returns all available adoptable dogs
//
//	@Summary		Get all dogs available for adoption
//	@Description	Get all dogs available for adoption
//	@Tags			dogs
//	@Produce		json
//	@Success		200		{object}	dog
//	@Failure		500		{object}	string
//	@Router			/v1/dogs [get]
func getDogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dogs)
}

// postDog adds a new dog from JSON received in the request body.
//
//	@Summary		Add a dog for adoption
//	@Description	Add a dog for adoption
//	@Tags			dogs
//	@Accept			json
//	@Produce		json
//	@Param			dog		body		dog			true	"Dog info for updating"
//	@Success		201		{object}	dog
//	@Failure		500		{object}	string
//	@Router			/v1/dogs [post]
func postDog(c *gin.Context) {
	var newDog dog

	// Call BindJSON to bind the received JSON to
	// newDog.
	if err := c.BindJSON(&newDog); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to bind new dog"})
		return
	}

	// Add the new dog to the in-memory slice.
	dogs = append(dogs, newDog)
	c.IndentedJSON(http.StatusCreated, newDog)
}

// getDogByID locates the dog whose ID value matches the id
// parameter sent by the client, then returns that dog as a response.
//
//	@Summary		Gets a dog by it's id
//	@Description	Gets a dog by it's id
//	@Tags			dogs
//	@Produce		json
//	@Param			id		path		int		true	"Dog ID"
//	@Success		200		{object}	dog
//	@Failure		400		{object}	string	"invalid dog id"
//	@Failure		404		{object}	string	"dog not found"
//	@Failure		500		{object}	string	"internal server error"
//	@Router			/v1/dogs/{id} [get]
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

// updateDogById updates the information for a dog using the JSON sent in the body of the POST
// only updates if the dog already exists
//
//	@Summary		Update a dog's information, if it exists
//	@Description	Update a dog's information, if it exists
//	@Tags			dogs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			dog		body		dog			true	"Dog info for updating"
//	@Success		201		{object}	dog
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Router			/v1/dogs/{id} [post]
func updateDogById(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid dog id type"})
		return
	}

	// Loop over the list of dogs, looking for
	// a dog whose ID value matches the parameter.
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

// adoptADog removes a dog from the system, because it has been adopted!
//
//	@Summary		Remove an adopted dog from the system
//	@Description	Remove an adopted dog from the system
//	@Tags			dogs
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Success		202		{object}	dog
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Router			/v1/dogs/{id} [delete]
func adoptADog(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid dog id type"})
	}

	// Loop over the list of dogs, looking for
	// an dog whose ID value matches the parameter.
	for i, a := range dogs {
		if a.ID == id {
			// delete here
			dogs = slices.Delete(dogs, i, i+1)
			c.IndentedJSON(http.StatusAccepted, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "dog not found"})
}
