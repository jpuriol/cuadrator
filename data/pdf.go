package data

import (
	"fmt"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func PrintPDF() error {

	quadrant, err := ReadQuadrant()
	if err != nil {
		return err
	}

	participants, err := ReadParticipants()
	if err != nil {
		return err
	}

	err = quadrant.CheckNames(participants)
	if err != nil {
		return err
	}

	err = quadrant.CheckShifts()
	if err != nil {
		return err
	}

	schema, err := ReadSchema()
	if err != nil {
		return err
	}

	var shiftNums []int
	for k := range schema.Shifts {
		shiftNums = append(shiftNums, k)
	}
	sort.Ints(shiftNums)

	var ocuppationNums []int
	for k := range schema.Occupations {
		ocuppationNums = append(ocuppationNums, k)
	}
	sort.Ints(ocuppationNums)

	m := maroto.New()

	m.AddRow(10, text.NewCol(12, schema.Name))

	for _, shiftN := range shiftNums {

		m.AddRow(5, text.NewCol(12, schema.ShiftName(shiftN)))

		m.AddRow(2,
			line.NewCol(12, props.Line{
				Style:         linestyle.Dashed,
				OffsetPercent: 0,
			}))

		for _, occupationN := range ocuppationNums {

			var teams strings.Builder
			for _, team := range quadrant[shiftN][occupationN] {
				teamStr := strings.Join(team, "-")
				teams.WriteString(fmt.Sprintf("%v, ", teamStr))
			}

			teamsText := strings.TrimSuffix(teams.String(), ", ")
			if teamsText == "" {
				continue
			}

			m.AddRow(11,
				text.NewCol(3, schema.OcupationName(occupationN)),
				text.NewCol(9, teamsText),
			)
		}
	}

	document, err := m.Generate()
	if err != nil {
		return err
	}

	pdfFileName := fmt.Sprintf("%s.pdf", schema.Name)
	err = document.Save(pdfFileName)
	if err != nil {
		return err
	}

	return nil
}

var headerStyle = props.Text{
	Size: 16,
}
