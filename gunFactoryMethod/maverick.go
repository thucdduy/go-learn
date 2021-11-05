package main

type maverick struct {
	gun
}

func newMaverick() *maverick {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
