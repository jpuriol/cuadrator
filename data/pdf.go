package data

import (
	"fmt"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

const pdfFileName = "quadrant.pdf"

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

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	for _, shiftN := range shiftNums {

		m.Row(5, func() {
			m.Text(schema.ShiftName(shiftN), shiftsStyle)
		})

		m.Line(1.0, props.Line{
			Style: consts.Dashed,
			Width: 0.3,
		})

		m.Row(2, func() {})

		for _, occupationN := range ocuppationNums {

			var teams strings.Builder
			for _, team := range quadrant[shiftN][occupationN] {
				teamStr := strings.Join(team, "-")
				teams.WriteString(fmt.Sprintf("%v, ", teamStr))
			}

			m.Row(11, func() {
				m.Col(2, func() {
					m.Text(schema.OcupationName(occupationN), occupationStyle)
				})
				m.ColSpace(1)
				m.Col(9, func() {
					m.Text(strings.TrimSuffix(teams.String(), ", "), namesStyle)
				})
			})
		}
	}

	err = m.OutputFileAndClose(pdfFileName)
	if err != nil {
		return err
	}

	return nil
}

var shiftsStyle = props.Text{
	Family: consts.Helvetica,
	Style:  consts.BoldItalic,
	Align:  consts.Left,
	Size:   12,
}

var occupationStyle = props.Text{
	Family: consts.Helvetica,
	Style:  consts.Italic,
	Align:  consts.Right,
	Size:   10,
}

var namesStyle = props.Text{
	Family: consts.Helvetica,
	Style:  consts.Normal,
	Align:  consts.Left,
	Size:   10,
}
