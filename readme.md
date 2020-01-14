# Rock-Paper-Scissors game (rps game)

### Requirements
  - GNU make utility (http://ftp.gnu.org/gnu/make)
  - docker (with sudo privileges)

### How to start the game
Run the following commands to play the game:
```sh
$ make build 
$ make play
```

-- The "make build" command build a docker  image of the program
-- The "make play" run the docker image build by the "make build" command

### Folder structure
```
rock-paper-scisors
│   Dockerfile
│   go.mod    
│   main.go
|   Makefile
|   readme.md
└─── players
│   │   rps-player.go 
└─── rps
    │   rps-game.go
    │   rps-game-imp.go
    │   rps-game_test.go
```

* rps-player.go is a player for rps games player, it has functionalities like start, stop, pause etc. It accepts any implementation of an rps game
* rps-game.go defines what a rps game is. It's an interface containing all the properties and function a rps game should have. As a direct consequence, an rps game that implements this interface can be played by our "rps-player".
* rps-game-imp.go is an implementation of the rps game defined in rps-game, this implmentation is provided to the "rps-player" as the game engine.