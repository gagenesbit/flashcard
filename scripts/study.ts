let iter= 0;
let cards: any[] = [];
let cardText: HTMLParagraphElement;
window.onload = () => {

    const cardContainer = document.getElementById('cardContainer') as HTMLDivElement;
    cardText = document.getElementById('cardText') as HTMLParagraphElement;

    //Get the DeckID from the URl
    const urlPath = window.location.pathname;
    const deckID = urlPath.split('/').pop();

    //Call the api to get cards for this deck
    fetch(`/api/cards/${deckID}`)
        .then(response => response.json())
        .then(data => {
            cards=data;
            cardText.innerHTML=`${data[iter].Question}`;
        })
        .catch(error => {
            console.error('Error fetching cards:', error);
        });
   

    

    cardContainer.addEventListener('click', () => {
        if(cardText.innerHTML==`${cards[iter].Question}`){
            cardText.innerHTML=`${cards[iter].Answer}`;
        }
        else{
            cardText.innerHTML=`${cards[iter].Question}`;
        }
    
    });
};

function nextCard(): void{
    if(iter+1<cards.length){
        iter++;
    }
    else{
        iter=0;
    }
    renderCard()
}
function previousCard(): void{
    if(iter>0){
        iter--;
    }
    else{
        iter=cards.length-1;
    }
    renderCard()
}

function renderCard(): void {
    cardText.innerHTML=`${cards[iter].Question}`;
}
