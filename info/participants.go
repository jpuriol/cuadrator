package info

import (
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type Occupation struct {
	ShifID         int
	OccupationName string
}

func GetPartipantOccupation(participantName string) ([]Occupation, error) {
	partcipants, err := getParticipants()
	if err != nil {
		return nil, err
	}

	if !partipantsExists(participantName, partcipants) {
		return nil, fmt.Errorf("Participant %q does not exist in %s", participantName, participantsFile)
	}

	quadrantData, err := getQuadrantData()
	if err != nil {
		return nil, err
	}

	schema, err := getSchema()
	if err != nil {
		return nil, err
	}

	var res []Occupation

	var shiftIDs []int
	for k := range quadrantData {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	for _, shiftID := range shiftIDs {
		for occupationID, occupation := range quadrantData[shiftID] {
			for _, team := range occupation {
				for _, person := range team {
					if person == participantName {
						occup := Occupation{
							ShifID:         shiftID,
							OccupationName: schema.ocupationName(occupationID),
						}
						res = append(res, occup)
					}
				}
			}
		}
	}

	return res, nil
}

func getParticipants() ([]string, error) {
	data, err := os.ReadFile(participantsFile)
	if err != nil {
		return []string{}, err
	}

	var participants []string
	err = yaml.Unmarshal(data, &participants)
	if err != nil {
		return []string{}, err
	}

	return participants, nil
}
