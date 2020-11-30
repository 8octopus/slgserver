package model

type CityFacility struct {
	DB         dbSync `xorm:"-"`
	Id         int    `xorm:"id pk autoincr"`
	RId        int    `xorm:"rid"`
	CityId     int    `xorm:"cityId"`
	Facilities string `xorm:"facilities"`
}

func (this *CityFacility) TableName() string {
	return "city_facility"
}