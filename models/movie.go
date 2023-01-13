package models

type Cast struct {
	FirstName   string `json:"firstname" bson:"firstname"`
	LastName   string `json:"lastname" bson:"lastname"`
}

type Movie struct {
	Id         int        `json:"id" bson:"id"` 
	Name       string     `json:"name" bson:"name"`
	Banner     string     `json:"banner" bson:"banner"`
	year       int        `json:"year" bson:"year"`
	Cast       Cast       `json:"Cast" bson:"Cast"`
	ratings    int        `json:"ratings" bson:"ratings"`
}
