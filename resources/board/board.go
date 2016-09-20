package board

import "fmt"

type Board struct {
	Zones []BoardZone
}

type BoardZone struct {
	ZoneName   string
	Multiplier int
	Values     []BoardArea
}

type BoardArea struct {
	Name  string
	Value int
}

func getBoardValues(prefix string) []BoardArea {
	boardAreas := []BoardArea{}

	for i := 1; i <= 20; i++ {
		var ba BoardArea
		ba.Name = fmt.Sprint(prefix, i)
		ba.Value = i
		boardAreas = append(boardAreas, ba)
	}
	return boardAreas
}

func newBoard() *Board {
	var b Board

	//Area
	var z BoardZone
	z.ZoneName = "Area"
	z.Multiplier = 1
	z.Values = getBoardValues("")
	b.Zones = append(b.Zones, z)

	//Singles
	z.ZoneName = "Single"
	z.Multiplier = 1
	z.Values = getBoardValues("S")
	b.Zones = append(b.Zones, z)

	//Doubles
	z.ZoneName = "Double"
	z.Multiplier = 2
	z.Values = getBoardValues("D")
	b.Zones = append(b.Zones, z)

	//Trebles
	z.ZoneName = "Treble"
	z.Multiplier = 3
	z.Values = getBoardValues("T")
	b.Zones = append(b.Zones, z)

	//Center
	z = BoardZone{}
	z.ZoneName = "Bull"
	z.Multiplier = 1
	var bull BoardArea
	bull.Name = "BullsEye"
	bull.Value = 50
	z.Values = append(z.Values, bull)
	var outerBull BoardArea
	outerBull.Name = "OuterBull"
	outerBull.Value = 25
	z.Values = append(z.Values, outerBull)

	b.Zones = append(b.Zones, z)

	return &b
}
