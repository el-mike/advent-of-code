package day18

type Droplet struct {
	Coords Coords
}

func NewDroplet(coords Coords) *Droplet {
	return &Droplet{
		Coords: coords,
	}
}
