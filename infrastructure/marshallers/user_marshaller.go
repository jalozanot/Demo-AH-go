package marshallers

import (
	"encoding/json"
	"fmt"

	"github.com/jalozanot/demoCeiba/domain/model"
)

type PublicUser struct {
	Id           int64  `json:"id"`
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	CodigoBarras string `json:"CodigoBarras"`
}
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func Marshall(isPublic bool, user model.MovieDto) interface{} {
	if isPublic {
		return PublicUser{
			Id:           user.Id,
			Nombre:       user.Nombre,
			Categoria:    user.Categoria,
			CodigoBarras: user.CodigoBarras,
		}
	}
	userJson, errUn := json.Marshal(user)
	fmt.Println(errUn)
	fmt.Println(user)
	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func MarshallArray(isPublic bool, users []model.MovieDto) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = Marshall(isPublic, user)
	}
	return result
}
