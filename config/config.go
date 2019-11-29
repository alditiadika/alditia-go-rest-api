package config

//Conf struct
type Conf struct {
	URL    string
	DBName string
}

//GetConf return configuration of database
func GetConf() Conf {
	return Conf{URL: "mongodb://localhost:27017", DBName: "test"}
}
