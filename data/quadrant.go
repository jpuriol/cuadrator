package data

import (
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type Quadrant map[int]Shift

func ReadQuadrant() (Quadrant, error) {
	rawData, err := os.ReadFile(quadrantFile)
	if err != nil {
		return nil, err
	}

	var data Quadrant
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q Quadrant) Check() error {

	// Check that names are in participants file
	participants, _ := readParticipants()
	for shiftID, shift := range q {
		nameFreq := shift.nameFrequency()
		for name := range nameFreq {
			if !participants.Exists(name) {
				return fmt.Errorf("Participant %q on shift ID %d does not exist in %s", name, shiftID, participantsFile)
			}
		}
	}

	// Check that the same person is not twice on the same shift
	for shiftID, shift := range q {
		nameFreq := shift.nameFrequency()
		for name, freq := range nameFreq {
			if freq > 1 {
				return fmt.Errorf("Participant %q has %d ocuppations on shift ID %d", name, freq, shiftID)
			}
		}

	}

	return nil
}

type Occupation struct {
	ShifID         int
	OccupationName string
}

func (q Quadrant) GetOcupation(participantName string) ([]Occupation, error) {
	participants, err := readParticipants()
	if err != nil {
		return nil, err
	}

	schema, err := readSchema()
	if err != nil {
		return nil, err
	}

	if !participants.Exists(participantName) {
		return nil, fmt.Errorf("Participant %q does not exist in %s", participantName, participantsFile)
	}

	var res []Occupation

	var shiftIDs []int
	for k := range q {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	for _, shiftID := range shiftIDs {
		for occupationID, occupation := range q[shiftID] {
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
