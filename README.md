## UniBun

Game for ebitenengine gamejam 2025 https://itch.io/jam/ebitengine-game-jam-2025

The thema is UNITE

## Gameplay

Main idea is the player controls 2 buns and have to get united to the patty burger.

The plan:
- [x] Grid map
- [x] Move around something
- [x] Turn base
- [ ] Sokoban style of moving around and pushing stuff
- [x] Allow to push patty
  - [ ] Allow to push other objects
- [ ] Goal: 
  - [x] Merge as one
  - [x] Bring it to a winning position
  - [ ] Properly draw some restaurant client
- [ ] Different ingredients? Cheese, Lettuce, Bacon...
- [ ] Enemies? Rats, Pidgeons
- [ ] Pass to next level

## Packages
- Game --> Main game with all the logic
- Entities --> Players, enemies... (*Maybe this should just be in Game package?)
- Level --> About the grid, it should not have dependencies with the game at all
- Config --> Params to change the game