package models

type PlanetShort struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	SystemId string `json:"systemId" bson:"systemId,omitempty"`
}

type Planet struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	SystemId string `json:"systemId" bson:"systemId,omitempty"`
}

type CreatePlanetDTO struct {
	SystemId string `json:"systemId"`
}

type UpdatePlanetDTO struct {
	Id       string `json:"id"`
	SystemId string `json:"systemId"`
}
