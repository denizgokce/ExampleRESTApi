package DataSource

import (
	model "ExampleRESTApi/Models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"gopkg.in/mgo.v2"
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
	person.ID = strconv.Itoa(getValueForNextSequence())
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

func getValueForNextSequence() int {
	var doc model.Counter
	var err = GetInstance().dbContext.C("counter").Find(bson.M{"_id": "person_id"}).One(&doc)
	if err != nil {
		log.Fatal(err)
	}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"sequence_value": 1}},
		ReturnNew: true,
	}
	GetInstance().dbContext.C("counter").Find(bson.M{"_id": "person_id"}).Apply(change, &doc)
	return int(doc.Sequence_Value)
}
