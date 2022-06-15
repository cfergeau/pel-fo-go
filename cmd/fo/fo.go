package main

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	fodb "github.com/cfergeau/fo-go/pkg/db"
	"io/ioutil"
	"strings"
)

func main() {
	//fmt.Printf("downloading latest version of the database\n")
	filename, err := downloadLatest()
	if err != nil {
		fmt.Printf("Error downloading database: %v\n", err)
		return
	}
	//fmt.Printf("parsing %s\n", filename)
	db, err := fodb.ParseFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing database: %v\n", err)
		return
	}
	printConnaissances(db)
	//printMyConnaissances(db, 141)
}

func downloadLatest() (string, error) {
	rawUrl, err := ioutil.ReadFile("url.txt")
	if err != nil {
		return "", fmt.Errorf("missing url.txt with URL to the database")
	}
	url := strings.TrimSpace(string(rawUrl))
	resp, err := grab.Get(".", string(url))
	if err != nil {
		return "", err
	}

	return resp.Filename, nil
}

func printConnaissances(db *fodb.Db) {
	fmt.Printf("Game Name\tOk\tBof\tNiet\n")
	for _, game := range db.Games {
		if len(game.Boxes) == 0 {
			//fmt.Printf("skipping %s (game removed from library)\n", game.Titre)
			continue
		}
		fmt.Printf("%s\t%d\t%d\t%d\n", game.Titre, len(game.Ok), len(game.Bof), len(game.Niet))
	}
}

func printMyConnaissances(db *fodb.Db, userId int) {
	fmt.Printf("Game Name\tOk\tBof\tNiet\n")
	user, knownUser := db.VolunteersById[userId]
	if !knownUser {
		panic("Unknown user ID")
	}
	gameIds := []int{}
	gameIds = append(gameIds, user.Ok...)
	gameIds = append(gameIds, user.Bof...)

	for _, gameId := range gameIds {
		game, hasGame := db.GamesById[gameId]
		if !hasGame {
			panic("Unknown game ID")
		}
		if len(game.Boxes) == 0 {
			//fmt.Printf("skipping %s (game removed from library)\n", game.Titre)
			continue
		}
		fmt.Printf("%s\t%d\t%d\t%d\n", game.Titre, len(game.Ok), len(game.Bof), len(game.Niet))
	}
}
