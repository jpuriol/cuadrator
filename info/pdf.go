package info

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

const pdfFileName = "quadrant.pdf"

func PrintPDF() error {

	err := Check()
	if err != nil {
		return fmt.Errorf("Inconsistent data: %v", err)
	}

	schema, err := getSchema()
	if err != nil {
		return err
	}

	quadrant, err := getQuadrantData()
	if err != nil {
		return err
	}

	m := pdf.NewMaroto(consts.Landscape, consts.A4)
	m.SetBorder(true)

	m.Row(15, func() {
		m.Col(2, func() {})

		m.Col(3, func() {
			m.Text(schema.Occupations[1], headerStyle)
		})
		m.Col(2, func() {
			m.Text(schema.Occupations[2], headerStyle)
		})
		m.Col(2, func() {
			m.Text(schema.Occupations[3], headerStyle)
		})
		m.Col(2, func() {
			m.Text(schema.Occupations[4], headerStyle)
		})
	})

	for shifID := 1; shifID <= 4; shifID++ {
		m.Row(40, func() {
			m.Col(2, func() {
				m.Text(schema.Shifts[shifID], headerStyle)
			})

			m.Col(3, func() {
				for i, team := range quadrant[shifID][1] {
					str := strings.Join(team, " - ")
					m.Text(str, bodyStyle(i))
				}
			})

			m.Col(2, func() {
				for i, team := range quadrant[shifID][2] {
					str := strings.Join(team, " - ")
					m.Text(str, bodyStyle(i))
				}
			})

			m.Col(2, func() {
				for i, team := range quadrant[shifID][3] {
					str := strings.Join(team, " - ")
					m.Text(str, bodyStyle(i))
				}
			})

			m.Col(2, func() {
				for i, team := range quadrant[shifID][4] {
					str := strings.Join(team, " - ")
					m.Text(str, bodyStyle(i))
				}
			})
		})
	}

	err = m.OutputFileAndClose(pdfFileName)
	if err != nil {
		return err
	}

	return nil
}

var headerStyle = props.Text{
	Top:    5,
	Family: consts.Helvetica,
	Style:  consts.Bold,
	Align:  consts.Center,
	Size:   12,
}

func bodyStyle(i int) props.Text {
	return props.Text{
		Top:    4.0 + float64(i*5),
		Family: consts.Helvetica,
		Style:  consts.Normal,
		Align:  consts.Center,
		Size:   10,
	}
}
