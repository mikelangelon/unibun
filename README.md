## UniBun

Game for ebitenengine gamejam 2025 https://itch.io/jam/ebitengine-game-jam-2025

The thema is UNITE

## Gameplay

Main idea is the player controls 2 buns and have to get united to the patty burger.

The plan:
- [x] Grid map
- [x] Move around something
- [x] Turn base
  - [x] Make more obvious the turn (flashing)
- [ ] Sokoban style of moving around and pushing stuff
- [x] Allow to push patty
  - [ ] Allow to push other objects
- [ ] Goal: 
  - [x] Merge as one
  - [x] Bring it to a winning position
  - [x] Properly draw some restaurant client
- [x] Different ingredients? Cheese, Lettuce, Bacon...
  - [x] Cheese --> make bun faster
  - [x] Lettuce --> cross walls
  - [ ] Bacon --> ??
- [ ] Enemies?
  - [x] Pidgeon --> random move
  - [x] Snake --> follow target + dash --> a bit buggy
  - [ ] Rat
  - [ ] Animate characters
- [x] Pass to next level
- [x] Sounds
- [x] Music
- [ ] Endless mode
- [ ] Show all levels individually

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