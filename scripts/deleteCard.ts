//Imports the Notyf library
declare var Notyf: any;

//Deletes the specified card from the DOM
function deleteCard(button: HTMLElement) {
    //Prompts the user to confirm they want to delete the card
    const userConfirmed = window.confirm("Are you sure you want to delete this card?");
    if(userConfirmed){
        //Locates closes li and removes it
        const card = button.closest('li');
        if (card) {
        card.remove();
        }
        //Notifies the user that the card has been removed.
        var deleteNotyf = new Notyf();
        deleteNotyf.success('Card Removed');
    }
  }