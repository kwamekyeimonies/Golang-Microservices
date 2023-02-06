package models

type AnimalsFacts struct{
	CommonName string `json:"commonname"`
	ScientificName string `json:"scientificname"`
	HeightInches int `json:"heighinches"`
	FavoriteFoods []string `json:"favoritefoods"`
	CanSwim bool `json:"canswim"`
}