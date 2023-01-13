package models

type Cast struct {
	FirstName   string `json:"firstname" bson:"firstname"`
	LastName   string `json:"lastname" bson:"lastname"`
}

type Movie struct {
	Id         int        `json:"id" bson:"id"` 
	userId     int        `json:"userid" bson:"userid"` 
	Name       string     `json:"name" bson:"name"`
	Cover      string     `json:"cover" bson:"cover"`
	Url      string     `json:"url" bson:"url"`
	year       int        `json:"year" bson:"year"`
	Likes       int        `json:"likes" bson:"likes"`
	Views       int        `json:"views" bson:"views"`
	Cast       Cast       `json:"Cast" bson:"Cast"`
	Comments       string     `json:"comments" bson:"comments"`
	Duration       int        `json:"duration" bson:"duration"	
	Desc       string     `json:"desc" bson:"desc"`
	
}