package termui

import (
	"fmt"
	"strings"

	"github.com/jpuriol/cuadrator/core"
)

func FormatStats(schedule core.Schedule, schema core.Schema) string {
	var sb strings.Builder

	maxOccupationLen := 0
	for _, name := range schema.Occupations {
		if len(name) > maxOccupationLen {
			maxOccupationLen = len(name)
		}
	}

	for _, shiftID := range schedule.OrderedShiftIDs() {
		shiftName := schema.ShiftName(shiftID)
		header := fmt.Sprintf(" %s ", shiftName)
		separator := ""
		for i := 0; i < len(header); i++ {
			separator += "-"
		}
		sb.WriteString(fmt.Sprintf("%s\n", separator))
		sb.WriteString(fmt.Sprintf("%s\n", header))
		sb.WriteString(fmt.Sprintf("%s\n", separator))
		shift := schedule[shiftID]
		occupationIDs := shift.OrderedOccupationIDs()
		for i := 0; i < len(occupationIDs); i += 2 {
			line := ""
			for j := 0; j < 2 && i+j < len(occupationIDs); j++ {
				occupationID := occupationIDs[i+j]
				groups := shift[occupationID]
				count := len(groups)
				bar := ""
				for k := 0; k < count; k++ {
					bar += "█"
				}
				line += fmt.Sprintf(" %*s | %-12s (%d)  ", maxOccupationLen, schema.OccupationName(occupationID), bar, count)
			}
			sb.WriteString(fmt.Sprintf("%s\n", line))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
