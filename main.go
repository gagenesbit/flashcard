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

	//Loading Static content to be referenced
	app := gin.Default()
	app.Static("/static", "./static")
	app.Static("/scripts", "./scripts")

	app.LoadHTMLGlob("templates/*")

	//Redirects to home page
	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/home")
	})

	//Renders index.html to the home page
	app.GET("/home", func(c *gin.Context) {
		decks := getAllDecks()
		c.HTML(200, "index.html", gin.H{
			"Decks": decks,
		})

	})
	//Renders createdeck.html to the create deck page
	app.GET("/createdeck", func(c *gin.Context) {
		c.HTML(200, "createdeck.html", nil)
	})

	//Implements Post-Redirect-Get to prevent form resubmission
	app.POST("/createdeck", func(c *gin.Context) {
		//Gets data from form
		deck_name := c.PostForm("deck_name")
		description := c.PostForm("description")
		questions := c.PostFormArray("question[]")
		answers := c.PostFormArray("answer[]")
		//Sends deck data and cards to DB
		createNewDeck(deck_name, description, questions, answers)
		//Redirects Home
		c.Redirect(http.StatusSeeOther, "/home")
	})

	//Dynamically renders the page for each deck
	app.GET("/deck/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		//Retrieves deck from DB
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}
		//Sends deck and card data to template
		c.HTML(200, "deck.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	//Used by HTMX GET to retrieve a form with all the card data. Sends it to /deck/:id
	app.GET("/deck/edit/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		//Retrieves deck from DB
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}
		//Sends deck and card data to template
		c.HTML(200, "editDeck.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	//Handles form submisson of edits to deck
	//Implements Post-Redirect-Get to prevent form resubmission
	app.POST("/deck/edit/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		//Retrieves card data from form
		questions := c.PostFormArray("question[]")
		answers := c.PostFormArray("answer[]")

		//Recreates deck with current card data
		addCardsToDeck(deckID, questions, answers)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}

		c.Redirect(http.StatusSeeOther, "/deck/"+c.Param("id"))
	})

	//Renders the study page
	app.GET("/deck/study/:id", func(c *gin.Context) {
		deckID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error strconv")
		}
		//Retrieves deck and its cards from the DB
		deck, cards, err := getDeckById(deckID)
		if err != nil {
			fmt.Println("error getting deck and cards")
		}
		//Renders template with deck and card data
		c.HTML(200, "study.html", gin.H{
			"Deck":  deck,
			"Cards": cards,
		})
	})

	//Used by HTMX GET to retrieve a card form to hold all the card data. Sends it to /createdeck and /deck/:id when editing
	app.GET("/createcardtemplate", func(c *gin.Context) {
		c.HTML(200, "createcardtemplate.html", nil)
	})

	//Gets all cards in a deck for the study.ts scripting to handle the study functionality
	app.GET("/api/cards/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("error getting deck id from front end")
		}
		//gets all cards for a deck
		cards := getCards(id)
		//Sends the cards formatted as JSON
		c.JSON(200, cards)
	})

	//Runs app on port 8080
	app.Run(":8080")
}
