package main

import (
	"engine"
	"fmt"
	"net"

	"github.com/google/uuid"
)

type client struct {
	addr *net.UDPAddr
}

type game struct {
	entities map[string]*engine.Entity
	clients  []*client
}

var g *game

func initGame() {
	fmt.Println("Intializing game state...")
	g = new(game)
	g.entities = make(map[string]*engine.Entity)

	fmt.Println("Game intialized")
}

func addEntity(new *engine.Entity) string {
	identifier := uuid.New().String()
	g.entities[string(identifier)] = new
	return identifier
}
