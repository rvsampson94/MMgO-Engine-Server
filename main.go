package main

import (
	"engine"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var conn *net.UDPConn

func main() {
	fmt.Println("**************************************")
	fmt.Println("* Welcome to MMgO Engine Server v1.1 *")
	fmt.Println("**************************************")

	// Setup UDP socket
	addr, err := net.ResolveUDPAddr("udp", ":8000")
	if err != nil {
		panic(err)
	}
	conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go startUDPHandler(conn)

	// Instantiate a game
	initGame()

	var delta float64 = 0
	t := time.Now().UnixNano()
	for range time.Tick(50 * time.Millisecond) {
		delta = float64(time.Now().UnixNano()-t) / 1000000000 // calculate time elapsed since last frame
		t = time.Now().UnixNano()                             // set time this frame started

		// RUN GAME LOGIC
		//fmt.Println(delta)
		// Update entities
		//fmt.Println(g.entities["1"].getComponent(&remoteEventQueue{}).(*remoteEventQueue))
		for _, ent := range g.entities {
			for _, comp := range ent.Components {
				comp.OnUpdate(delta)
			}
		}

		// Send game state updates to clients
		for identifier, ent := range g.entities {
			if ent.Update {
				for _, client := range g.clients {
					controller := ent.GetComponent(&engine.RemotePlayerControler{}).(*engine.RemotePlayerControler)
					data := []byte(fmt.Sprintf("%s;1;%f,%f,%f,%f", identifier, ent.Position.X, ent.Position.Y, controller.Direction.X, controller.Direction.Y))
					conn.WriteToUDP(data, client.addr)
				}
				ent.Update = false
			}
		}

	}
}

func startUDPHandler(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		go handleUDP(n, addr, buffer)
	}
}

func handleUDP(n int, addr *net.UDPAddr, buffer []byte) {
	buffer = buffer[:n]
	arr := strings.Split(string(buffer), ";")
	identifier := arr[0]
	opcode := arr[1]
	data := arr[2]
	switch opcode {
	case "0": // new player connect
		player := engine.NewEntity(0, 0)
		rpc := engine.NewRemotePlayerController(player, 400)
		player.AddComponent(rpc)
		clientID := addEntity(player)

		g.clients = append(g.clients, &client{
			addr: addr,
		})

		data := []byte(fmt.Sprintf("0;0;%s", clientID))
		conn.WriteToUDP(data, addr)
		sendAll(addr)
	case "1": // player input
		params := strings.Split(data, ",")
		posX, err := strconv.ParseFloat(params[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		posY, err := strconv.ParseFloat(params[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		dirX, _ := strconv.ParseFloat(params[2], 64)
		dirY, _ := strconv.ParseFloat(params[3], 64)
		event := engine.NewInputEvent(engine.NewVector(posX, posY), engine.NewVector(dirX, dirY))
		g.entities[string(identifier)].GetComponent(&engine.RemotePlayerControler{}).(*engine.RemotePlayerControler).AddEvent(event)
	}

}

func sendAll(addr *net.UDPAddr) {
	data := "0;3;"
	i := 0
	for key, ent := range g.entities {
		controller := ent.GetComponent(&engine.RemotePlayerControler{}).(*engine.RemotePlayerControler)
		playerData := fmt.Sprintf("%s:%f:%f:%f:%f", key, ent.Position.X, ent.Position.Y, controller.Direction.X, controller.Direction.Y)
		data += playerData
		if i != len(g.entities)-1 {
			data += ","
		}
	}
	conn.WriteToUDP([]byte(data), addr)
}
