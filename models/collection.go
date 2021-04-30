package models

type Collection struct {
	IDCollection   int    `form:"IDCollection" json:"IDKoleks"`
	NameCollection string `form:"NameCollection" json:"NameCollection"`
}

type CollectionResponse struct {
	Response
	Data []Collection `form:"Data" json:"Data"`
}

type DetailCollection struct {
	IDDetailCollection int     `form:"IDDetailCollection" json:"IDDetailCollection"`
	IDCollection       Collection `form:"IDCollection" json:"IDCollection"`
	IDFilm          Film    `form:"IDFilm" json:"IDFilm"`
}

type DetailCollectionResponse struct {
	Response
	Data []DetailCollection `form:"data" json:"data"`
}
