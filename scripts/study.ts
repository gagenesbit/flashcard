declare var Notyf: any;




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

//This is an implementation of the Fisher-Yates shuffle algorithm 
//https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_modern_algorithm
function shuffleCards(): void{
    console.log("Shuffling cards")
    for(var i=cards.length-1;i>0;i--){
        let j=Math.floor(Math.random()*(i+1))
        let temp=cards[i]
        cards[i]=cards[j]
        cards[j]=temp
    }
    renderCard()
    var notyf = new Notyf();
    notyf.success('Cards Shuffled');
}


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
