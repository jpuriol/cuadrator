package adapters

import (
	"fmt"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jpuriol/cuadrator/core"
)

func WritePDF(q core.Quadrant, p core.Participants, s core.Schema) error {
	err := q.ValidateNames(p)
	if err != nil {
		return err
	}

	err = q.ValidateShifts()
	if err != nil {
		return err
	}

	var shiftNums []int
	for k := range s.Shifts {
		shiftNums = append(shiftNums, k)
	}
	sort.Ints(shiftNums)

	var occupationNums []int
	for k := range s.Occupations {
		occupationNums = append(occupationNums, k)
	}
	sort.Ints(occupationNums)

	m := maroto.New()

	m.AddRow(10, text.NewCol(12, s.Name))

	for _, shiftN := range shiftNums {

		m.AddRow(5, text.NewCol(12, s.ShiftName(shiftN)))

		m.AddRow(2,
			line.NewCol(12, props.Line{
				Style: linestyle.Dashed,
			}))

		for _, occupationN := range occupationNums {

			var teams strings.Builder
			for _, team := range q[shiftN][occupationN] {
				teamStr := strings.Join(team, "-")
				teams.WriteString(fmt.Sprintf("%v, ", teamStr))
			}

			teamsText := strings.TrimSuffix(teams.String(), ", ")
			if teamsText == "" {
				continue
			}

			m.AddRow(11,
				text.NewCol(3, s.OccupationName(occupationN)),
				text.NewCol(9, teamsText),
			)
		}
	}

	document, err := m.Generate()
	if err != nil {
		return err
	}

	pdfFileName := fmt.Sprintf("%s.pdf", s.Name)
	err = document.Save(pdfFileName)
	if err != nil {
		return err
	}

	return nil
}
