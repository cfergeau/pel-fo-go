package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ParseFromFile(filename string) (*Db, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var myRawDb rawDb
	err = json.Unmarshal(content, &myRawDb)
	if err != nil {
		return nil, err
	}

	db := &myRawDb.States.Db
	err = db.index()
	if err != nil {
		return nil, err
	}

	err = db.fillGameBoxes()
	if err != nil {
		return nil, err
	}

	err = db.fillGameKnowledge()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Db) index() error {
	db.GamesById = make(map[int]*Game)
	for i, game := range db.Games {
		if _, idExists := db.GamesById[game.Id]; idExists {
			return fmt.Errorf("duplicate GameID")
		}
		db.GamesById[game.Id] = &db.Games[i]
	}

	db.VolunteersById = make(map[int]*Volunteer)
	for i, volunteer := range db.Volunteers {
		if _, idExists := db.VolunteersById[volunteer.Id]; idExists {
			return fmt.Errorf("duplicate VolunteerID")
		}
		db.VolunteersById[volunteer.Id] = &db.Volunteers[i]
	}

	//fmt.Printf("boxes: %+v", db.Boxes)
	db.BoxesByName = make(map[string][]*Box)
	for _, box := range db.Boxes {
		db.BoxesByName[box.Container] = append(db.BoxesByName[box.Container], &box)
	}

	db.TeamsById = make(map[int]*Team)
	for i, team := range db.Teams {
		if _, idExists := db.TeamsById[team.Id]; idExists {
			return fmt.Errorf("duplicate TeamID")
		}
		db.TeamsById[team.Id] = &db.Teams[i]
	}
	return nil
}

func (db *Db) fillGameBoxes() error {
	for _, box := range db.Boxes {
		game, gameIdExists := db.GamesById[box.GameId]
		if !gameIdExists {
			return fmt.Errorf("unknown GameID %d in box %d", box.GameId, box.Id)
		}
		game.Boxes = append(game.Boxes, box.Id)
	}

	return nil
}

func (db *Db) fillGameKnowledge() error {
	var javTeamId int
	for _, team := range db.Teams {
		if team.Name == "Jeux À Volonté" {
			javTeamId = team.Id
		}
	}

	for _, volunteer := range db.Volunteers {
		if volunteer.Equipe != javTeamId {
			continue
		}
		for _, gameId := range volunteer.Ok {
			game, hasGame := db.GamesById[gameId]
			if !hasGame {
				return fmt.Errorf("unknown GameID %d in Ok list for volunteer %d", gameId, volunteer.Id)
			}
			game.Ok = append(game.Ok, volunteer.Id)
		}
		for _, gameId := range volunteer.Bof {
			game, hasGame := db.GamesById[gameId]
			if !hasGame {
				return fmt.Errorf("unknown GameID %d in Bof list for volunteer %d", gameId, volunteer.Id)
			}
			game.Bof = append(game.Bof, volunteer.Id)
		}
		for _, gameId := range volunteer.Niet {
			game, hasGame := db.GamesById[gameId]
			if !hasGame {
				return fmt.Errorf("unknown GameID %d in Niet list for volunteer %d", gameId, volunteer.Id)
			}
			game.Niet = append(game.Niet, volunteer.Id)
		}
	}

	return nil
}
