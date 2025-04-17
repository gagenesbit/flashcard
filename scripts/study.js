//Initalize global variables
var iter = 0;
var cards = [];
var card;
var cardContainer;
var frontText;
var backText;
var flipped = false;
//Waits for DOM to load
window.onload = function () {
    //Locates relevant elements from DOM
    cardContainer = document.getElementById('cardContainer');
    card = document.getElementById('card');
    frontText = document.getElementById('frontText');
    backText = document.getElementById('backText');
    //Isolates Deck ID from URL
    var urlPath = window.location.pathname;
    var deckID = urlPath.split('/').pop();
    //Fetches all cards from the specified deck from the api endpoint
    fetch("/api/cards/".concat(deckID))
        //Translates the response to JSON
        .then(function (response) { return response.json(); })
        .then(function (data) {
        //Assigns the JSON data to the cards variable
        cards = data;
        //Renders the card text
        renderCard();
    })
        .catch(function (error) {
        console.error('Error fetching cards:', error);
    });
    //Adds an event listener for keyboard shortcuts
    document.addEventListener("keyup", function (event) {
        //Will flip card on space bar press
        if (event.key === " ") {
            flipCard(cardContainer);
        }
        //Switches the visible card to the previous card on left arrow press
        if (event.key === "ArrowLeft") {
            previousCard();
        }
        //Switches the visible card to the next card on right arrow press
        if (event.key === "ArrowRight") {
            nextCard();
        }
        //Shuffles the cards on s key press
        if (event.key === "s") {
            shuffleCards();
        }
    });
};
//This is an implementation of the Fisher-Yates shuffle algorithm 
//https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_modern_algorithm
function shuffleCards() {
    //Makes the front of the card visible before the shuffle takes place
    initalizeCardState();
    //Iterates through card array
    for (var i = cards.length - 1; i > 0; i--) {
        //Finds a random index between 0 and i
        var j = Math.floor(Math.random() * (i + 1));
        //Swaps the location of the cards at index i and index j
        var temp = cards[i];
        cards[i] = cards[j];
        cards[j] = temp;
    }
    //Rerenders card text
    renderCard();
    //Notifies user that cards were shuffled
    var notyf = new Notyf();
    notyf.success('Cards Shuffled');
}
//Progresses the visible card to the next card in cards
function nextCard() {
    initalizeCardState();
    //Loops back to index 0 if iter goes out of the bounds of cards
    iter = (iter + 1) % cards.length;
    renderCard();
}
//Progresses the visible card to the next card in cards
function previousCard() {
    initalizeCardState();
    //Loops back to index 0 if iter goes out of the bounds of cards
    if (iter > 0) {
        iter--;
    }
    else {
        iter = cards.length - 1;
    }
    renderCard();
}
//Displays the current cards text on the DOM
function renderCard() {
    frontText.innerHTML = "".concat(cards[iter].Question);
    backText.innerHTML = "".concat(cards[iter].Answer);
}
//Switches the card text and triggers the flipping animation 
function flipCard(container) {
    flipped = !flipped;
    if (card) {
        //Toggles flipped class on the card
        card.classList.toggle('flipped');
    }
}
//Sets the front of the card to be visible regarless of current state
function initalizeCardState() {
    if (flipped) {
        flipCard(cardContainer);
    }
}
