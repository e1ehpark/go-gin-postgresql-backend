package models

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Name string `json:"name" bson:"name"`
	CEO string `json:"ceo" bson:"ceo"`
	ValueProposition string `json:"value_proposition" bson:"value_proposition"`
	Industry string `json:"industry" bson:"industry"`
	Founded Founder `json:"founder" bson:"founder" gorm:"embeded"`
}

type Founder struct {
	gorm.Model
	Name string `json:"name" bson:"name"`
	Profession string `json:"profession" bson:"profession"`
	Age   int `json:"age" bson:"age"`
}

// saves a startup to the database

func (startup *Startup) Save() (*Startup, error) {
	err := Database.Model(&Startup{}).Create(&startup).Error
	if err != nil {
		return &Startup{}, err
	}
	return startup, nil
}

// Fetches all startups from the database
func FetchAllStartups() ([]*Startup, error) {
    var startups []*Startup
    err := Database.Find(&startups).Error
    if err != nil {
        return []*Startup{}, err
    }
    return startups, nil
}

//Fetches a startup from the database
func (startup *Startup) UpdateStartup(id string) (*Startup, error) {
	err := Database.Model(&Startup{}).Where("id = ?", id).Updates(&startup).Error
	if err != nil {
		return &Startup{}, err
	}
	return startup, nil	
	
}

// Deletes a startup from the database
func DeleteStartup(id string) error {
	err := Database.Model(&Startup{}).Where("id = ?", id).Delete(&Startup{}).Error
	if err != nil {
		return err
	}
	return nil
}