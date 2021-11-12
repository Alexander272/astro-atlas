package models

type PlanetShort struct {
	Id         string  `json:"id" bson:"_id,omitempty"`
	SystemId   string  `json:"systemId" bson:"systemId,omitempty"`
	Name       string  `json:"name" bson:"name,omitempty"`
	DateOpened int64   `json:"dateOpened" bson:"dateOpened"`
	Class      string  `json:"class" bson:"class,omitempty"`
	Age        float32 `json:"age" bson:"age,omitempty"`
}

type Planet struct {
	Id              string   `json:"id" bson:"_id,omitempty"`
	SystemId        string   `json:"systemId" bson:"systemId,omitempty"`
	Name            string   `json:"name" bson:"name,omitempty"`
	Description     string   `json:"description" bson:"description,omitempty"`
	DateOpened      int64    `json:"dateOpened" bson:"dateOpened,omitempty"`
	Discoverer      []string `json:"discoverer" bson:"discoverer,omitempty"`
	DetectionMethod string   `json:"detectionMethod" bson:"detectionMethod,omitempty"`
	Class           string   `json:"class" bson:"class,omitempty"`
	Weight          float32  `json:"weight" bson:"weight,omitempty"`
	Radius          float32  `json:"radius" bson:"radius,omitempty"`
	Temperature     int32    `json:"temperature" bson:"temperature,omitempty"`
	Age             float32  `json:"age" bson:"age,omitempty"`
	Period          float32  `json:"period" bson:"period,omitempty"`
	Speed           float32  `json:"speed" bson:"speed,omitempty"`
}

func NewPlanet(dto CreatePlanetDTO) Planet {
	return Planet{
		SystemId:        dto.SystemId,
		Name:            dto.Name,
		Description:     dto.Description,
		DateOpened:      dto.DateOpened,
		Discoverer:      dto.Discoverer,
		DetectionMethod: dto.DetectionMethod,
		Class:           dto.Class,
		Weight:          dto.Weight,
		Radius:          dto.Radius,
		Temperature:     dto.Temperature,
		Age:             dto.Age,
		Period:          dto.Period,
		Speed:           dto.Speed,
	}
}

type CreatePlanetDTO struct {
	SystemId        string   `json:"systemId" binding:"required"`
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description"`
	DateOpened      int64    `json:"dateOpened"`
	Discoverer      []string `json:"discoverer"`
	DetectionMethod string   `json:"detectionMethod"`
	Class           string   `json:"class"`
	Weight          float32  `json:"weight"`
	Radius          float32  `json:"radius"`
	Temperature     int32    `json:"temperature"`
	Age             float32  `json:"age"`
	Period          float32  `json:"period"`
	Speed           float32  `json:"speed"`
}

type UpdatePlanetDTO struct {
	Id              string   `json:"id"`
	SystemId        string   `json:"systemId"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	DateOpened      int64    `json:"dateOpened"`
	Discoverer      []string `json:"discoverer"`
	DetectionMethod string   `json:"detectionMethod"`
	Class           string   `json:"class"`
	Weight          float32  `json:"weight"`
	Radius          float32  `json:"radius"`
	Temperature     int32    `json:"temperature"`
	Age             float32  `json:"age"`
	Period          float32  `json:"period"`
	Speed           float32  `json:"speed"`
}
