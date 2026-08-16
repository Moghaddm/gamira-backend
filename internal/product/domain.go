package product

type product struct {
	ID      int64  `bson:"_id,omitempty"`
	Title   string `bson:"title,omitempty"`
	Address string `bson:"address,string"`
	CPU     string `bson:"cpu,omitempty"`
	Memory  string `bson:"memory,omitempty"`
	GPU     string `bson:"gpu,omitempty"`
	Storage string `bson:"storage,omitempty"`
}
