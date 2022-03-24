package info

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Occupation struct {
	ShiftName      string
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

	var res []Occupation
	for shiftID, shift := range quadrantData {
		for occupationID, occupation := range shift {
			for _, team := range occupation {
				for _, person := range team {
					if person == participantName {
						o := Occupation{
							ShiftName:      getShiftName(shiftID),
							OccupationName: getOcupationName(occupationID),
						}
						res = append(res, o)
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
