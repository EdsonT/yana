package company

import (
	"fmt"
	"log"
	"time"
	"yana/model"

	"github.com/rs/xid"
)

//CreateNewPost initializes primary parameters of a Post, and validate data
func CreateNewCompany(params *model.Company) *model.Company {
	np := new(model.Company)
	np.Code = xid.New().String()
	np.Name = params.Name
	np.City = params.City
	np.Country = params.Country
	np.Description = params.Description
	np.Address = params.Address
	np.Status = "Active"
	np.CreatedAt = time.Now()
	np.UpdatedAt = time.Now()
	result, err := Add(np)
	fmt.Println(result, err)
	return np
}

func UpdateCompany(code string, params model.Company) model.Company {
	var up model.Company
	Update(code, params)
	up = Find(code)
	return up
}
func GetCompany(params model.Company) []*model.Company {
	return Get(params)
}

func Search(params *model.Company) []model.Company {
	var result []model.Company
	return result
}
func Delete(code string) (string, error) {
	result, err := DeletLogical(code)
	if err == nil {
		log.Println("Object Deleted:", result)
		return "success", nil
	} else {
		log.Println(err)
		return "", err
	}
}
