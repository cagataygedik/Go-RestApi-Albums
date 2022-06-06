package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Quantity int    `json:"quantity"`
}

var albums = []album{
	{ID: "1", Title: "Dawn FM", Artist: "The Weeknd", Quantity: 100},
	{ID: "2", Title: "Donda", Artist: "Kanye West", Quantity: 100},
	{ID: "3", Title: "If I Can't Have Love,I Have Power", Artist: "Halsey", Quantity: 100},
	{ID: "4", Title: "Certified Lover Boy", Artist: "Drake", Quantity: 100},
	{ID: "5", Title: "Random Access Memories", Artist: "Daft Punk", Quantity: 100},
	{ID: "6", Title: "SYRE", Artist: "Jaden", Quantity: 100},
	{ID: "7", Title: "ANTI", Artist: "Rihanna", Quantity: 100},
	{ID: "8", Title: "Goodbye&Good Riddance", Artist: "Juice WRLD", Quantity: 100},
	{ID: "9", Title: "?", Artist: "XXXTENTACION", Quantity: 100},
	{ID: "10", Title: "Astroworld", Artist: "Travis Scott", Quantity: 100},
	{ID: "11", Title: "Legendary", Artist: "Tyga", Quantity: 100},
	{ID: "12", Title: "Lights Out", Artist: "Exhel&UFO 361", Quantity: 100},
	{ID: "13", Title: "Harry's House", Artist: "Harry Styles", Quantity: 100},
	{ID: "14", Title: "Un Verano Sin Ti", Artist: "Bad Bunny", Quantity: 100},
	{ID: "15", Title: "The Eminem Show", Artist: "Eminem", Quantity: 100},
	{ID: "16", Title: "4:44", Artist: "Jay-Z", Quantity: 100},
	{ID: "17", Title: "Heaven or Hell", Artist: "Don Toliver", Quantity: 100},
	{ID: "18", Title: "JACKBOYS", Artist: "Jackboys", Quantity: 100},
	{ID: "19", Title: "Drip Harder", Artist: "Lil Baby", Quantity: 100},
	{ID: "20", Title: "Thriller", Artist: "Michael Jackson", Quantity: 100},
	{ID: "21", Title: "Memories...Do Not Open", Artist: "The Chainsmokers", Quantity: 100},
	{ID: "22", Title: "I NEVER LIKED YOU", Artist: "Future", Quantity: 100},
	{ID: "23", Title: "SR3MM", Artist: "Swae Lee", Quantity: 100},
	{ID: "24", Title: "Back to Black", Artist: "Amy Winehouse", Quantity: 100},
	{ID: "25", Title: "Sunset in the Blue", Artist: "Melody Gardot", Quantity: 100},
	{ID: "26", Title: "Future Nostalgia", Artist: "Dua Lipa", Quantity: 100},
	{ID: "27", Title: "=", Artist: "Ed Sheeran", Quantity: 100},
	{ID: "28", Title: "CKay The First", Artist: "CKay", Quantity: 100},
	{ID: "29", Title: "Planet Her", Artist: "Doja Cat", Quantity: 100},
	{ID: "30", Title: "MONTERO", Artist: "Lil Nas X", Quantity: 100},
	{ID: "31", Title: "American Teen", Artist: "Khalid", Quantity: 100},
	{ID: "32", Title: "30", Artist: "Adele", Quantity: 100},
	{ID: "33", Title: "Justice", Artist: "Justin Bieber", Quantity: 100},
	{ID: "34", Title: "Based On A Feeling", Artist: "Sabrina Claudio", Quantity: 100},
	{ID: "35", Title: "Familia", Artist: "Camilla Cabello", Quantity: 100},
	{ID: "36", Title: "World Of Walker", Artist: "Alan Walker", Quantity: 100},
	{ID: "37", Title: "It'll All Make Sense In The End", Artist: "James Arthur", Quantity: 100},
	{ID: "38", Title: "Mr. Percocet", Artist: "Noah Cyrus", Quantity: 100},
	{ID: "39", Title: "Rare", Artist: "Selena Gomez", Quantity: 100},
	{ID: "40", Title: "Shockwave", Artist: "Marshmello", Quantity: 100},
	{ID: "41", Title: "DS4EVER", Artist: "Gunna", Quantity: 100},
	{ID: "42", Title: "NOT ALL HEROES WEAR CAPES", Artist: "Metro Boomin", Quantity: 100},
	{ID: "43", Title: "i am < i was", Artist: "21 Savage", Quantity: 100},
	{ID: "44", Title: "Punk", Artist: "Young Thug", Quantity: 100},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func albumByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutAlbum(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Album not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnAlbum(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getAlbumByID(id string) (*album, error) {
	for i, b := range albums {
		if b.ID == id {
			return &albums[i], nil
		}
	}

	return nil, errors.New("album not found")
}

func createAlbum(c *gin.Context) {
	var newBook album

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	albums = append(albums, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", albumByID)
	router.POST("/albums", createAlbum)
	router.PATCH("/checkout", checkoutAlbum)
	router.PATCH("/return", returnAlbum)
	router.Run("localhost:8080")
}
