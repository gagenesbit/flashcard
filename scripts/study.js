var iter = 0;
var cards = [];
var cardText;
window.onload = function () {
    var cardContainer = document.getElementById('cardContainer');
    cardText = document.getElementById('cardText');
    //Get the DeckID from the URl
    var urlPath = window.location.pathname;
    var deckID = urlPath.split('/').pop();
    //Call the api to get cards for this deck
    fetch("/api/cards/".concat(deckID))
        .then(function (response) { return response.json(); })
        .then(function (data) {
        cards = data;
        cardText.innerHTML = "".concat(data[iter].Question);
    })
        .catch(function (error) {
        console.error('Error fetching cards:', error);
    });
    cardContainer.addEventListener('click', function () {
        if (cardText.innerHTML == "".concat(cards[iter].Question)) {
            cardText.innerHTML = "".concat(cards[iter].Answer);
        }
        else {
            cardText.innerHTML = "".concat(cards[iter].Question);
        }
    });
};
//This is an implementation of the Fisher-Yates shuffle algorithm 
//https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_modern_algorithm
function shuffleCards() {
    console.log("Shuffling cards");
    for (var i = cards.length - 1; i > 0; i--) {
        var j = Math.floor(Math.random() * (i + 1));
        var temp = cards[i];
        cards[i] = cards[j];
        cards[j] = temp;
    }
    renderCard();
    var notyf = new Notyf();
    notyf.success('Cards Shuffled');
}
function nextCard() {
    if (iter + 1 < cards.length) {
        iter++;
    }
    else {
        iter = 0;
    }
    renderCard();
}
function previousCard() {
    if (iter > 0) {
        iter--;
    }
    else {
        iter = cards.length - 1;
    }
    renderCard();
}
function renderCard() {
    cardText.innerHTML = "".concat(cards[iter].Question);
}
