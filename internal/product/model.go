package product

type product struct {
	Id    int64  `json:"id"`
	Title string `bson:"title,omitempty"`

}
