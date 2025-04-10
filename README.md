# Gage Nesbit's Flash Card App

This is my flash card web app, built using a variety of tools. I came up with the idea when I started learning Japanese recently and found that existing flashcard tools didn’t meet my needs. They either had too many ads or lacked useful features. The initial release includes basic functionality, but I believe that with more time, this app will grow into a robust and highly functional tool.

## Technologies I Used

The backend is written in Go, utilizing the Gin framework. As someone familiar with Flask in Python, I found Gin to be similar and easy to learn. I used Go's built-in HTML templating module to handle the templates, and it has worked perfectly for my needs. 

The data for my flash card decks is stored in a MySQL database. I appreciate how lightweight MySQL is and how well it suits smaller-scale projects like this. The setup was quick and straightforward. Currently, my database consists of two tables—one for decks and one for cards. As the app grows, I expect the database to evolve with additional tables and more complex relationships.

On the frontend, I used the Pico CSS library to get started, customizing it from there. I also used HTMX for several small tasks, such as dynamic content updates. Having used HTMX in the past, I found it ideal for some of the design patterns I wanted to implement. 

For the scripting, I used TypeScript to handle the functionality of flipping each flashcard and cycling through the deck. I chose TypeScript for its robust features, offering all the functionality of JavaScript with the added benefit of static typing. I plan to use TypeScript more as I expand the functionality of the study page.

## A Brief Tour of My Application

### Home Page
Here, you can navigate to the "Create a Deck" page or select which deck you want to access.
![Home Page Screenshot](./mdImages/flashcard_1.0_02.png)

### Create A Deck Page
On this page, you can create a deck and add cards to it.
![Create A Deck Page Screenshot](./mdImages/flashcard_1.0_01.png)

### Study Page
Here, you can study the flashcards in your deck. Simply click a card to flip it and view the other side. You can also use the buttons to cycle through the cards.
![Study Page Screenshot](./mdImages/flashcard_1.0_04.png)

## The Future
This is only the first release of the project, and I plan to expand the app's functionality in the future.

Some of my plans for future features include:
1. Being able to add and delete cards to existing decks
2. Shuffling the flashcard deck for studying
3. Adding a flipping animation for the cards
4. Adding other study question styles, such as multiple choice
5. Implementing a spaced repetition algorithm for more efficient learning sessions
