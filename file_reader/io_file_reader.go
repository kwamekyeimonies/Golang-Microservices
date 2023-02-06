package filereader

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kwamekyeimonies/Golang-Microservices/models"
)

func IO_File_Reader_Test() {

	jsonData := `
	{
		"commonname":"Crocodile",
		"scientificname":"Crocodylidae",
		"heighinches":34,
		"favoritefoods":["Humans","Hen","Chocolate"],
		"canswim":true
	}
	`

	//Read Json data
	rdr := strings.NewReader(jsonData)

	var sloth models.AnimalsFacts
	if err := json.NewDecoder(rdr).Decode(&sloth); err != nil {
		fmt.Println("Error deserializing JSON ", err)
		return
	}

	fmt.Printf(sloth.CommonName, sloth.ScientificName, sloth.HeightInches, sloth.FavoriteFoods, sloth.CanSwim)
}
