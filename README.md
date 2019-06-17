# Aembassador

**Aembassador now has an MVP which will be a sidestep of its original plan but will help get it there faster and better! Check it out here: [MVP](/planning/MVP.md)**

An AI that plays Keyforge Decks X amount of times to see who has the highest chance of winning according to an AI.

I would like to have it so people upload a deck, the deck gets analyzed and then played 1000 times against registered decks and gets a win ratio based on it being played this 1000 times. However, since the deck is in the system its win ratio and plays will change over time. When the system is not in use it will just start randomly pitting decks against eachother to have an ever evolving ratio and data model.

Every day have the top 100 winning decks of the day and also have a page of the highest performing decks.

Eventually... very far and if eventually.... Create an AI arena battle using Crucible Online technology with the match system and create something like Salty Bets.

The Application will consists of four major parts as follows

## The Analyzer
Analyzer assesses the deck and sees what kind of playstyle this deck would favor based on the deck type. It goes over all cards that are in the deck and gives an estimation of what the decks best play style is as well as what it's biggest weakness is. As of Age of Ascension I find that decks fall into these catagories of playstyle:

### Deck Types

- *Racing Decks*: These decks are about gaining Aember fast and getting to 3 keys faster then their opponent an example of a deck like this that I own would be [Formulaically Singleminded Saydee](https://keyforge-compendium.com/decks/de0d7ac3-de52-499f-9c4e-96d1209ca6a7)

- *Flood Decks*: Are about board presence. They can keep their opponents boards down with fighting while reaping in for big aember gain. these decks seem to favor a wall of creatures. An example of a deck that I have that is like this is [The Animal that Walks for Dance](https://keyforge-compendium.com/decks/26f0c80c-f59a-4db2-a14a-aa6a12db0af4)

- *Control Decks*: These are an interesting breed and was wondering if I should make two categories for this 1 being board control and another being aember control. However I decided on a broad Control Deck favoring both aember and creature control with more of an emphasis on aember control. These decks usually have normal to smaller amount of creatures, normal to small amount of aember gain. But usually are able to keep their opponents board/aember at bay while it creeps to 3 keys. An example of a deck I have like this is [Halfstein, He who Hacks the Library](https://keyforge-compendium.com/decks/7d69a2da-2a28-4576-873d-81633b26a5d2)

- *Combo Decks*: Combo decks are more prominent in AoA and are decks that tend to focus on hand size and archive size. These decks enjoy starting slow but can cause huge bursts that can close in fast mid to late game. I could be wrong that this is a deck type that has sprung out of the new set and do not actually have one my self to link here but I have played against a few.

More on the planning of the Anlyzer can be seen here, [Analyzer Planning Notes](/planning/ANALYZER.md)

## The Arena

The Arena will be the part of the application that handles the matches and will most likely be the more grueling part of this application and might be the part that wil be my undoing. I have a rough guess of how this will go as follows.

### How the AI plays
1. Check to see if it can forge a key
2. Check if it won.
3. Assess opponents situation
    - How many keys do they have?
    - What is their current Key cost?
    - How much Aember do they have?
    - Are they in check?
    - Which key is it?
    - What is the threat level of the board?
4. Assess your situation same as opponents to come up with an aggression focus.
    - Aggression Focus is the rating that you will give to the following areas
    - Aember Gaining - What is the maximum amount of Aember I can gain this turn?
    - Aember Control - What is the most Aember I can take from my opponent? Or the highest I can make their keys cost?
    - Board Gaining - Should I try to dominate the board?
    - Board Control - Do I need to dwindle their board?
    - Assassin mode - Is there an artifact or creature that is such a high threat that it needs to be dealt with more then anything but letting your opponent forge the third key? Is their an answer for it on the board? In your hand? Even in your deck? If there is an answer in your deck would it be best if the deck was cycled through?
5. Go over your hand (and archive) and asses which house would be the best option to choose based on Aggression Focus.
    - Look over known combos in hand and board
    - Look at creatures fight/reap/action abilities to see what their max value is.
    - In some instances it might be worth throwing a good card into an enemy card at a loss based on the situation.

### Situation Flags
Do to the nature of card games, cards can affect the rules of the game. We will need to make flags that alter this games vacuumed reality. Such things as Miasma making a player unable to forge a key regardless of their amount of aember or Archimedes making its neighbors archive when destroyed. Each card that has the ability to change the rules of the game will add its situation flag to a pool that each action will check. Some cards affects will add situation flags to your opponents pool. Such as the cards that make it so if your opponents creature reaps it gets stunned.

I will have a link here to a documentation area that has a more indepth look at Arena

## The Stats
Stats will be the part that takes the data from the database and crunches the numbers into awesome data like win percentage, etc. It will be the part that people will want to see

## The Face
This will be the UI that brings it all together!