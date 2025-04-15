package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Deck struct {
	ID          int
	Name        string
	Description string
}

type Card struct {
	ID       int
	Question string
	Answer   string
	DeckID   int
}

var db *sql.DB

func initDB() {
	var err error
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get the values from the environment
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	//Connect to DB
	db, err = sql.Open("mysql", dsn)
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
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating decks table:", err)
	}

	fmt.Println("Database and table are ready!")
}

func deleteAllCardsInDeck(deckID int) {
	removeDeckSQL := "DELETE FROM cards WHERE deckID = ?;"
	_, err := db.Exec(removeDeckSQL, deckID)
	if err != nil {
		log.Fatal("Error Removing all cards from deck:", err)
	}
}

func createNewDeck(deckName string, description string, questions []string, answers []string) {
	var err error
	insertDeckSQL := "INSERT INTO decks (name, description) VALUES (?, ?);"

	result, err := db.Exec(insertDeckSQL, deckName, description)
	if err != nil {
		log.Fatal("Error creating new deck:", err)
	}
	deckID, err := result.LastInsertId()
	addCardsToDeck(int(deckID), questions, answers)
}

func addCardsToDeck(deckID int, questions []string, answers []string) {
	var err error
	insertCardSQL := "INSERT INTO cards (question, answer,deckid) VALUES (?, ?, ?);"
	for i := 0; i < len(questions); i++ {

		_, err = db.Exec(insertCardSQL, questions[i], answers[i], deckID)
		if err != nil {
			log.Fatal("Error creating new card:", err)
		}
	}

}

func getAllDecks() []Deck {
	rows, err := db.Query("SELECT * FROM decks")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var decks []Deck

	for rows.Next() {
		var d Deck
		rows.Scan(&d.ID, &d.Name, &d.Description)
		decks = append(decks, d)
	}

	return decks
}

func getDeckById(ID int) (Deck, []Card, error) {
	row := db.QueryRow("SELECT id, name, description FROM decks WHERE id = ?", ID)
	var deck Deck
	row.Scan(&deck.ID, &deck.Name, &deck.Description)

	var cards []Card
	rows, err := db.Query("SELECT * FROM cards Where deckid= ?", ID)
	if err != nil {
		return deck, cards, nil
	}
	for rows.Next() {
		var c Card
		rows.Scan(&c.ID, &c.Question, &c.Answer, &c.DeckID)
		cards = append(cards, c)
	}

	return deck, cards, nil
}

func getCards(deckID int) []Card {
	var cards []Card
	rows, err := db.Query("SELECT * FROM cards Where deckid= ?", deckID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var c Card
		rows.Scan(&c.ID, &c.Question, &c.Answer, &c.DeckID)
		cards = append(cards, c)
	}
	return cards
}
