//Deletes the specified card from the DOM
function deleteCard(button) {
    //Prompts the user to confirm they want to delete the card
    var userConfirmed = window.confirm("Are you sure you want to delete this card?");
    if (userConfirmed) {
        //Locates closes li and removes it
        var card = button.closest('li');
        if (card) {
            card.remove();
        }
        //Notifies the user that the card has been removed.
        var deleteNotyf = new Notyf();
        deleteNotyf.success('Card Removed');
    }
}
