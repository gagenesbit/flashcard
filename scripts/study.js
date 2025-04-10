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
