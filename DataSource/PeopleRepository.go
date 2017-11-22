package DataSource

import (
	model "ExampleRESTApi/Models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func GetPeople() []model.Person {
	var result []model.Person
	var err = GetInstance().dbContext.C("people").Find(nil).All(&result)
	if err != nil {
		return nil
	} else {
		return result
	}
}
func GetPerson(id string) model.Person {
	var result model.Person
	var err = GetInstance().dbContext.C("people").Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return result
	} else {
		return result
	}
}

func AddPerson(person *model.Person) {
	var err = GetInstance().dbContext.C("people").Insert(person)
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePerson(id string) {
	var err = GetInstance().dbContext.C("people").Remove(bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePerson(id string, person *model.Person) {
	var err = GetInstance().dbContext.C("people").Update(bson.M{"id": id}, person)
	if err != nil {
		log.Fatal(err)
	}
}
