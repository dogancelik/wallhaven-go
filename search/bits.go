package search

type Purity struct {
	Sfw     bool
	Sketchy bool
	Nsfw    bool
}

type Categories struct {
	General bool
	Anime   bool
	People  bool
}

func getBit(bit bool) string {
	if bit == true {
		return "1"
	} else {
		return "0"
	}
}

func (cat Categories) ToBits() string {
	bits := ""
	bits += getBit(cat.General)
	bits += getBit(cat.Anime)
	bits += getBit(cat.People)
	return bits
}

func (pur Purity) ToBits() string {
	bits := ""
	bits += getBit(pur.Sfw)
	bits += getBit(pur.Sketchy)
	bits += getBit(pur.Nsfw)
	return bits
}
