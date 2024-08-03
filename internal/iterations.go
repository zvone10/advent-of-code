package internal

func GenerateMoveDiffs() []Coordinates {
	return []Coordinates{
		{
			X: -1,
			Y: 0,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 0,
			Y: -1,
		},
		{
			X: 0,
			Y: 1,
		},
	}
}

type Coordinates struct {
	X int
	Y int
}
