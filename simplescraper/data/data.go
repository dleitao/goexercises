package data

//Entry is the data type of this project
type Entry struct {
	//asdf
	// ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string  `bson:"title,omitempty" json:"title,omitempty"`
	Buy   float64 `bson:"buy,omitempty" json:"buy,omitempty"`
	Sell  float64 `bson:"sell,omitempty" json:"sell,omitempty"`
}
