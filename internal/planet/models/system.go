package models

type SystemShort struct {
	Id            string  `json:"id" bson:"_id,omitempty"`
	Constellation string  `json:"constellation" bson:"constellation,omitempty"`
	Magnitude     float32 `json:"magnitude" bson:"magnitude,omitempty"`
	Distance      float32 `json:"distance" bson:"distance,omitempty"`
	Class         string  `json:"class" bson:"class,omitempty"`
	Age           float32 `json:"age" bson:"age,omitempty"`
	PlanetCount   int32   `json:"planetCount" bson:"planetCount"`
}

type System struct {
	Id            string  `json:"id" bson:"_id,omitempty"`
	Constellation string  `json:"constellation" bson:"constellation,omitempty"`
	Magnitude     float32 `json:"magnitude" bson:"magnitude,omitempty"`
	Distance      float32 `json:"distance" bson:"distance,omitempty"`
	Class         string  `json:"class" bson:"class,omitempty"`
	Weight        float32 `json:"weight" bson:"weight,omitempty"`
	Radius        float32 `json:"radius" bson:"radius,omitempty"`
	Temperature   int32   `json:"temperature" bson:"temperature,omitempty"`
	Metallicity   float32 `json:"metallicity" bson:"metallicity,omitempty"`
	Age           float32 `json:"age" bson:"age,omitempty"`
	PlanetCount   int32   `json:"planetCount" bson:"planetCount"`
}

func NewSystem(dto CreateSystemDTO) System {
	return System{
		Constellation: dto.Constellation,
		Magnitude:     dto.Magnitude,
		Distance:      dto.Distance,
		Class:         dto.Class,
		Weight:        dto.Weight,
		Radius:        dto.Radius,
		Temperature:   dto.Temperature,
		Metallicity:   dto.Metallicity,
		Age:           dto.Age,
		PlanetCount:   dto.PlanetCount,
	}
}

type CreateSystemDTO struct {
	Constellation string  `json:"constellation"`
	Magnitude     float32 `json:"magnitude"`
	Distance      float32 `json:"distance"`
	Class         string  `json:"class"`
	Weight        float32 `json:"weight"`
	Radius        float32 `json:"radius"`
	Temperature   int32   `json:"temperature"`
	Metallicity   float32 `json:"metallicity"`
	Age           float32 `json:"age"`
	PlanetCount   int32   `json:"planetCount"`
}

// func UpdateSystem(dto UpdateSystemDTO) System {
// 	return System{
// 		Id:            dto.Id,
// 		Constellation: dto.Constellation,
// 		Magnitude:     dto.Magnitude,
// 		Distance:      dto.Distance,
// 		Class:         dto.Class,
// 		Weight:        dto.Weight,
// 		Radius:        dto.Radius,
// 		Temperature:   dto.Temperature,
// 		Metallicity:   dto.Metallicity,
// 		Age:           dto.Age,
// 		PlanetCount:   dto.PlanetCount,
// 	}
// }

type UpdateSystemDTO struct {
	Id            string  `json:"id"`
	Constellation string  `json:"constellation"`
	Magnitude     float32 `json:"magnitude"`
	Distance      float32 `json:"distance"`
	Class         string  `json:"class"`
	Weight        float32 `json:"weight"`
	Radius        float32 `json:"radius"`
	Temperature   int32   `json:"temperature"`
	Metallicity   float32 `json:"metallicity"`
	Age           float32 `json:"age"`
	PlanetCount   int32   `json:"planetCount"`
}
