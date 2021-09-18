package order

type Geometry struct {
	Type        string
	Coordinates [2]float32
}
type Properties struct {
	Label       string
	Score       float64
	HouseNumber string
	Id          string
	Type        string
	Name        string
	PostCode    string
	CityCode    string
	x           float32
	y           float32
	City        string
	Context     string
	Importance  float64
	Street      string
}

type Feature struct {
	Type       string
	Geometry   Geometry
	Properties Properties
}

type ModelAdr struct {
	Type        string
	Version     string
	Features    []Feature
	Attribution string
	License     string
	Query       string
	Limit       int32
}
