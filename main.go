package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	initDB()

	app := gin.Default()
	app.Static("/static", "./static")
	app.Static("/scripts", "./scripts")

	app.LoadHTMLGlob("templates/*")

	// Route to serve the HTML page
	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/home")
	})
	app.GET("/home", func(c *gin.Context) {
		decks := getAllDecks()
		c.HTML(200, "index.html", gin.H{
			"Decks": decks,
		})

	})

	app.GET("/createdeck", func(c *gin.Context) {
		c.HTML(200, "createdeck.html", nil)
	})
	app.POST("/createdeck", func(c *gin.Context) {
		deck_name := c.PostForm("deck_name")
		description := c.PostForm("description")
		questions := c.PostFormArray("question[]")
		answers := c.PostFormArray("answer[]")
		createNewDeck(deck_name, description, questions, answers)
		c.Redirect(http.StatusSeeOther, "/home")
	})

	app.GET("/deck/0", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/home")
	})

	app.GET("/deck/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}

		c.HTML(200, "deck.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	app.POST("/deck/edit/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		questions := c.PostFormArray("question[]")
		answers := c.PostFormArray("answer[]")
		deleteAllCardsInDeck(deckID)
		addCardsToDeck(deckID, questions, answers)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}

		c.Redirect(http.StatusSeeOther, "/deck/:id")
	})

	app.GET("/deck/study/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}
		c.HTML(200, "study.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	app.GET("/deck/edit/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}

		c.HTML(200, "editDeck.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	app.GET("/createcardtemplate", func(c *gin.Context) {
		c.HTML(200, "createcardtemplate.html", nil)
	})

	app.GET("/api/cards/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error getting deck id from front end")
		}
		cards := getCards(id)
		c.JSON(200, cards)
	})
	app.Run(":8080")
}
