package info

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Check() error {
	quadrantData, err := getQuadrantData()
	if err != nil {
		return err
	}

	// Check that names are in participants file
	participants, _ := getParticipants()
	for shiftID, shift := range quadrantData {
		nameFreq := nameFrequency(shift)
		for name := range nameFreq {
			if !partipantsExists(name, participants) {
				return fmt.Errorf("Participant %q on shift ID %d does not exist in %s", name, shiftID, participantsFile)
			}
		}
	}

	// Check that the same person is not twice on the same shift
	for shiftID, shift := range quadrantData {
		nameFreq := nameFrequency(shift)
		for name, freq := range nameFreq {
			if freq > 1 {
				return fmt.Errorf("Participant %q has %d ocuppations on shift ID %d", name, freq, shiftID)
			}
		}

	}

	return nil
}

func getQuadrantData() (quadrant, error) {
	rawData, err := os.ReadFile(quadrantFile)
	if err != nil {
		return nil, err
	}

	var data quadrant
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func partipantsExists(name string, participants []string) bool {
	for _, pName := range participants {
		if name == pName {
			return true
		}
	}

	return false
}

func nameFrequency(s shift) map[string]int {
	frequency := make(map[string]int)

	for _, occupation := range s {
		for _, team := range occupation {
			for _, name := range team {
				frequency[name]++
			}
		}
	}

	return frequency
}
