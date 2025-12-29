package exporter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jpuriol/cuadrator/data"
)

func PrintPDF() error {
	d, err := data.LoadAll()
	if err != nil {
		return err
	}

	err = d.Quadrant.ValidateNames(d.Participants)
	if err != nil {
		return err
	}

	err = d.Quadrant.ValidateShifts()
	if err != nil {
		return err
	}

	var shiftNums []int
	for k := range d.Schema.Shifts {
		shiftNums = append(shiftNums, k)
	}
	sort.Ints(shiftNums)

	var occupationNums []int
	for k := range d.Schema.Occupations {
		occupationNums = append(occupationNums, k)
	}
	sort.Ints(occupationNums)

	m := maroto.New()

	m.AddRow(10, text.NewCol(12, d.Schema.Name))

	for _, shiftN := range shiftNums {

		m.AddRow(5, text.NewCol(12, d.Schema.ShiftName(shiftN)))

		m.AddRow(2,
			line.NewCol(12, props.Line{
				Style: linestyle.Dashed,
			}))

		for _, occupationN := range occupationNums {

			var teams strings.Builder
			for _, team := range d.Quadrant[shiftN][occupationN] {
				teamStr := strings.Join(team, "-")
				teams.WriteString(fmt.Sprintf("%v, ", teamStr))
			}

			teamsText := strings.TrimSuffix(teams.String(), ", ")
			if teamsText == "" {
				continue
			}

			m.AddRow(11,
				text.NewCol(3, d.Schema.OccupationName(occupationN)),
				text.NewCol(9, teamsText),
			)
		}
	}

	document, err := m.Generate()
	if err != nil {
		return err
	}

	pdfFileName := fmt.Sprintf("%s.pdf", d.Schema.Name)
	err = document.Save(pdfFileName)
	if err != nil {
		return err
	}

	return nil
}
