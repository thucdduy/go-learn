package main

type ak47 struct {
	gun
	ak47HiHi string
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
		ak47HiHi: "ak47 hihi",
	}
}
