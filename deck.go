package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// A struct to represent a Deck
type Deck struct {
	ID          int
	Name        string
	Description string
}

// A struct to represent a Card
type Card struct {
	ID       int
	Question string
	Answer   string
	DeckID   int
}

// Sets up a variable to reference the DB
var db *sql.DB

// Initializes DB tables if they are not created
func initDB() {
	var err error
	//Loads environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get the environment variable values to access the DB
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//Constructs the credentials to access the DB
	credentials := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	//Connect to the DB
	db, err = sql.Open("mysql", credentials)
	if err != nil {
		log.Fatal("Error connecting to flashcards_db:", err)
	}

	//Create table for cards if does not exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cards (
		id INT AUTO_INCREMENT PRIMARY KEY,
		question TEXT NOT NULL,
		answer TEXT NOT NULL,
		deckid INT NOT NULL
	);`

	//Executes create cards table SQL
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error cards creating cards table:", err)
	}

	//Create table for decks if does not exist
	createTableSQL = `
	CREATE TABLE IF NOT EXISTS decks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL
	);`
	//Executes create decks table SQL
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating decks table:", err)
	}
	//Output to confirm the database is ready
	fmt.Println("Database and tables are ready!")
}

// Deletes all the cards in a deck based on ID
func deleteAllCardsInDeck(deckID int) {
	removeDeckSQL := "DELETE FROM cards WHERE deckID = ?;"
	_, err := db.Exec(removeDeckSQL, deckID)
	if err != nil {
		log.Fatal("Error Removing all cards from deck:", err)
	}
}

// Creates a new entry in the deck table and creates cards for all submitted cards
func createNewDeck(deckName string, description string, questions []string, answers []string) {
	var err error
	insertDeckSQL := "INSERT INTO decks (name, description) VALUES (?, ?);"

	//Inserts deck into decks table
	result, err := db.Exec(insertDeckSQL, deckName, description)
	if err != nil {
		log.Fatal("Error creating new deck:", err)
	}
	//Retrieves the deck ID from the DB
	deckID, err := result.LastInsertId()
	//Adds all cards to the cards table based on deck ID
	addCardsToDeck(int(deckID), questions, answers)
}

// Adds all given cards to the cards table based on deck ID
func addCardsToDeck(deckID int, questions []string, answers []string) {
	//Removes any old versions of existing cards in deck
	deleteAllCardsInDeck(deckID)
	var err error
	insertCardSQL := "INSERT INTO cards (question, answer,deckid) VALUES (?, ?, ?);"
	//Iterativly inserts each card into the card table
	for i := 0; i < len(questions); i++ {
		_, err = db.Exec(insertCardSQL, questions[i], answers[i], deckID)
		if err != nil {
			log.Fatal("Error creating new card:", err)
		}
	}

}

// Retrieves all decks from DB
func getAllDecks() []Deck {
	rows, err := db.Query("SELECT * FROM decks")
	if err != nil {
		return nil
	}
	defer rows.Close()
	//Initalizes a slice to hold all decks
	var decks []Deck

	//Initializes a struct of type Deck to represent the deck data from the DB scan
	for rows.Next() {
		var d Deck
		rows.Scan(&d.ID, &d.Name, &d.Description)
		decks = append(decks, d)
	}
	//Returns the slice of all Deck structs
	return decks
}

// Retrieves a Deck and its cards from the DB based on deck ID
func getDeckById(ID int) (Deck, []Card, error) {
	//Retrives deck from decks table
	row := db.QueryRow("SELECT id, name, description FROM decks WHERE id = ?", ID)
	var deck Deck
	row.Scan(&deck.ID, &deck.Name, &deck.Description)

	var cards []Card = getCards(ID)

	//Returns the deck and its respective cards
	return deck, cards, nil
}

// r
func getCards(deckID int) []Card {
	//Initalizes a slice to hold all cards
	var cards []Card
	//Retrieves card data from DB
	rows, err := db.Query("SELECT * FROM cards Where deckid= ?", deckID)
	if err != nil {
		return cards
	}
	//Initializes a struct of type Card to represent the card data from the DB scan
	for rows.Next() {
		var c Card
		rows.Scan(&c.ID, &c.Question, &c.Answer, &c.DeckID)
		cards = append(cards, c)
	}
	//Returns the slice of cards
	return cards
}
