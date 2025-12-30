package pdf

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	marotoCore "github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jpuriol/cuadrator/core"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(ctx context.Context, schedule core.Schedule, participants core.Participants, schema core.Schema, w io.Writer) error {
	if err := core.ValidateSchedule(schedule, participants); err != nil {
		return err
	}

	var shiftNums []int
	for k := range schema.Shifts {
		shiftNums = append(shiftNums, k)
	}
	sort.Ints(shiftNums)

	var occupationNums []int
	for k := range schema.Occupations {
		occupationNums = append(occupationNums, k)
	}
	sort.Ints(occupationNums)

	m := maroto.New()

	// Header
	m.AddRow(15,
		text.NewCol(12, schema.Title, props.Text{
			Top:    5,
			Size:   20,
			Style:  fontstyle.Bold,
			Align:  align.Center,
			Family: fontfamily.Courier,
		}),
	)

	if schema.Subtitle != "" {
		m.AddRow(10,
			text.NewCol(12, schema.Subtitle, props.Text{
				Size:   14,
				Style:  fontstyle.Italic,
				Align:  align.Center,
				Family: fontfamily.Courier,
			}),
		)
	}

	m.AddRow(5) // Spacer after title

	for _, shiftN := range shiftNums {
		var shiftRows []marotoCore.Row

		shiftRows = append(shiftRows,
			row.New(2).Add(
				line.NewCol(12, props.Line{
					Style:     linestyle.Solid,
					Thickness: 0.3,
				}),
			),
		)

		// Shift Title Row
		shiftRows = append(shiftRows,
			row.New(6).Add(
				text.NewCol(12, schema.ShiftName(shiftN), props.Text{
					Size:   10,
					Style:  fontstyle.Bold,
					Align:  align.Center,
					Family: fontfamily.Courier,
				}),
			),
		)

		shiftRows = append(shiftRows,
			row.New(2).Add(
				line.NewCol(12, props.Line{
					Style:     linestyle.Solid,
					Thickness: 0.3,
				}),
			),
		)

		for _, occupationN := range occupationNums {
			var teams strings.Builder
			for _, group := range schedule[shiftN][occupationN] {
				teamStr := strings.Join(group, "-")
				teams.WriteString(fmt.Sprintf("%v / ", teamStr))
			}

			teamsText := strings.TrimSuffix(teams.String(), " / ")
			if teamsText == "" {
				continue
			}

			shiftRows = append(shiftRows,
				row.New(8).Add(
					text.NewCol(4, schema.OccupationName(occupationN), props.Text{
						Right:  5,
						Top:    1,
						Style:  fontstyle.Bold,
						Size:   9,
						Family: fontfamily.Courier,
						Align:  align.Right,
					}),
					text.NewCol(8, teamsText, props.Text{
						Top:    1,
						Size:   9,
						Family: fontfamily.Courier,
					}),
				),
			)
		}

		if schema.NoOccupation != "" {
			var free []string
			occupied := schedule[shiftN].NameFrequency()
			for pName := range participants {
				if _, ok := occupied[pName]; !ok {
					free = append(free, pName)
				}
			}
			sort.Strings(free)
			if len(free) > 0 {
				// Spacer
				shiftRows = append(shiftRows, row.New(2))

				shiftRows = append(shiftRows,
					row.New(2).Add(
						line.NewCol(12, props.Line{
							Style:     linestyle.Dashed,
							Thickness: 0.1,
						}),
					),
				)
				shiftRows = append(shiftRows,
					row.New(8).Add(
						text.NewCol(4, schema.NoOccupation, props.Text{
							Right:  5,
							Top:    1,
							Style:  fontstyle.Italic,
							Size:   7,
							Family: fontfamily.Courier,
							Align:  align.Right,
						}),
						text.NewCol(8, strings.Join(free, ", "), props.Text{
							Top:    1,
							Style:  fontstyle.Italic,
							Size:   7,
							Family: fontfamily.Courier,
						}),
					),
				)
			}
		}

		// Spacer
		shiftRows = append(shiftRows, row.New(5))

		m.AddRows(shiftRows...)
	}

	document, err := m.Generate()
	if err != nil {
		return err
	}

	_, err = w.Write(document.GetBytes())
	return err
}
