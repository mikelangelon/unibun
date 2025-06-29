## UniBun

Game for ebitenengine gamejam 2025 https://itch.io/jam/ebitengine-game-jam-2025

The theme is UNITE

The game is a puzzle about uniting ingredients to make the definitive burger.

## Gameplay

You have to control the ingredients that are in the game, except the patty meat. 
Cheese & Lettuce are optional to make the burger, but they can add powers to the buns if they unite as one:
- Bun + Lettuce: Can cross walls
- Bun + Cheese: Press Z(o X) and direction to prepare a dash (10 steps)

In any case, goal is:
- Put Bun + Patty meat + Bun in a row (except being downsided vertically, don't be a bad chef!)
- Go with your united burger to the client

Because the ingredients move always 2 steps, it has some strategy (and frustration)

To don't make it easy for you there are some animals ready to beat you. Each has a different behaviour and some, even some meal preference.

15 levels ready to try to challenge you a bit!

### To execute

You would need Go installed! Hopefully if you are in this repo is because you already have it :D

Get all the dependencies first:

    go mod tidy
And run it

    make run

### Controls

    Arrows/WASD --> Move around
    Z/X + Arrows/WASD --> Dash (when united with Lettuce)
    ENTER --> Select menu options
    SPACE --> Pause game

### Modes

The game has 2 modes.

    Normal --> 15 levels
    Endless --> random level. Not too validated or tested.... 

The plan:
- [x] Grid map
- [x] Move around something
- [x] Turn base
  - [x] Make more obvious the turn (flashing)
  - [x] Make obvious goal (glowing & growing)
- [x] Allow to push patty
  - [ ] Sokoban style of moving around and pushing stuff
- [ ] Goal: 
  - [x] Merge as one
  - [x] Bring it to a winning position
  - [x] Properly draw some restaurant client
- [x] Different ingredients? Cheese, Lettuce, Bacon...
  - [x] Cheese --> make bun faster
  - [x] Lettuce --> cross walls
    - [ ] Avoid going outside the screen 
  - [ ] Bacon --> ??
- [x] Enemies
  - [x] Pidgeon --> random move
  - [x] Mouse --> follows a path
  - [x] Duck --> follows target
  - [x] Snake --> follow target + dash --> a bit buggy
  - [x] Fly --> follows patty
  - [ ] Cat --> stands until you are too close
  - [x] Animate characters
    - [x] Fly
    - [x] Others
- [x] Pass to next level
- [x] Sounds
- [x] Music
- [x] Endless mode
  - [ ] Guarantee every puzzle can be solved
- [X] Show all levels individually
- [ ] Replace all DebugPrintAt by a proper Text
- [ ] Dash should not pass above patty.

## Packages
- Game --> Main game with all the logic
- Entities --> Players, enemies... (*Maybe this should just be in Game package?)
- Level --> About the grid, it should not have dependencies with the game at all
- Config --> Params to change the game

## Resources
Music: 
- By iemusic 
- https://www.gamedevmarket.net/asset/8-bit-chiptune-puzzle-pack-2

Eat sound: 
- By Naht 
- https://pixabay.com/sound-effects/eat-323883/

Tiles:
- By ElvGames
- https://elvgames.itch.io/free-fantasy-dreamland-dungeon

Characters:
- By ChatGPT 