package location

type Country struct {
	ID				string				`json:"id"`
	Name			string				`json:"name"`
	TimeZone		string				`json:"time_zone"`
	GeoInformation	GeoInformation		`json:"geo_informaion"`
	States			[]State 			`json:"states"`
}

type GeoInformation struct {
	Location GeoLocation		`json:"location"`
}

type GeoLocation struct {
	Latitude 	float64 		`json:"latitude"`
	Longtitude	float64			`json:"longtitude"`
}

type State struct {
	ID		string			`json:"id"`
	Name	string			`json:"name`
}