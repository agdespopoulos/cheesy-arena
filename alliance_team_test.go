// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNonexistentAlliance(t *testing.T) {
	clearDb()
	defer clearDb()
	db, err := OpenDatabase(testDbPath)
	assert.Nil(t, err)
	defer db.Close()

	allianceTeams, err := db.GetTeamsByAlliance(1114)
	assert.Nil(t, err)
	assert.Empty(t, allianceTeams)
}

func TestAllianceTeamCrud(t *testing.T) {
	clearDb()
	defer clearDb()
	db, err := OpenDatabase(testDbPath)
	assert.Nil(t, err)
	defer db.Close()

	allianceTeam := AllianceTeam{0, 1, 0, 254}
	db.CreateAllianceTeam(&allianceTeam)
	allianceTeams, err := db.GetTeamsByAlliance(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(allianceTeams))
	assert.Equal(t, allianceTeam, allianceTeams[0])

	allianceTeam.TeamId = 1114
	db.SaveAllianceTeam(&allianceTeam)
	allianceTeams, err = db.GetTeamsByAlliance(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(allianceTeams))
	assert.Equal(t, allianceTeam.TeamId, allianceTeams[0].TeamId)

	db.DeleteAllianceTeam(&allianceTeam)
	allianceTeams, err = db.GetTeamsByAlliance(1)
	assert.Nil(t, err)
	assert.Empty(t, allianceTeams)
}

func TestGetTeamsByAlliance(t *testing.T) {
	clearDb()
	defer clearDb()
	db, err := OpenDatabase(testDbPath)
	assert.Nil(t, err)
	defer db.Close()

	buildTestAlliances(db)
	allianceTeams, err := db.GetTeamsByAlliance(1)
	assert.Nil(t, err)
	if assert.Equal(t, 4, len(allianceTeams)) {
		assert.Equal(t, 254, allianceTeams[0].TeamId)
		assert.Equal(t, 469, allianceTeams[1].TeamId)
		assert.Equal(t, 2848, allianceTeams[2].TeamId)
		assert.Equal(t, 74, allianceTeams[3].TeamId)
	}
	allianceTeams, err = db.GetTeamsByAlliance(2)
	assert.Nil(t, err)
	if assert.Equal(t, 2, len(allianceTeams)) {
		assert.Equal(t, 1718, allianceTeams[0].TeamId)
		assert.Equal(t, 2451, allianceTeams[1].TeamId)
	}
}

func TestTruncateAllianceTeams(t *testing.T) {
	clearDb()
	defer clearDb()
	db, err := OpenDatabase(testDbPath)
	assert.Nil(t, err)
	defer db.Close()

	allianceTeam := AllianceTeam{0, 1, 0, 254}
	db.CreateAllianceTeam(&allianceTeam)
	db.TruncateAllianceTeams()
	allianceTeams, err := db.GetTeamsByAlliance(1)
	assert.Nil(t, err)
	assert.Empty(t, allianceTeams)
}

func TestGetAllAlliances(t *testing.T) {
	clearDb()
	defer clearDb()
	db, err := OpenDatabase(testDbPath)
	assert.Nil(t, err)
	defer db.Close()

	alliances, err := db.GetAllAlliances()
	assert.Nil(t, err)
	assert.Empty(t, alliances)

	buildTestAlliances(db)
	alliances, err = db.GetAllAlliances()
	assert.Nil(t, err)
	if assert.Equal(t, 2, len(alliances)) {
		if assert.Equal(t, 4, len(alliances[0])) {
			assert.Equal(t, 254, alliances[0][0].TeamId)
			assert.Equal(t, 469, alliances[0][1].TeamId)
			assert.Equal(t, 2848, alliances[0][2].TeamId)
			assert.Equal(t, 74, alliances[0][3].TeamId)
		}
		if assert.Equal(t, 2, len(alliances[1])) {
			assert.Equal(t, 1718, alliances[1][0].TeamId)
			assert.Equal(t, 2451, alliances[1][1].TeamId)
		}
	}
}

func buildTestAlliances(db *Database) {
	db.CreateAllianceTeam(&AllianceTeam{0, 2, 0, 1718})
	db.CreateAllianceTeam(&AllianceTeam{0, 1, 3, 74})
	db.CreateAllianceTeam(&AllianceTeam{0, 1, 1, 469})
	db.CreateAllianceTeam(&AllianceTeam{0, 1, 0, 254})
	db.CreateAllianceTeam(&AllianceTeam{0, 1, 2, 2848})
	db.CreateAllianceTeam(&AllianceTeam{0, 2, 1, 2451})
}
