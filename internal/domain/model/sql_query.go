package model

type DbConfig struct {
	DSN      string `json:"dsn"`
	DbType   string `json:"dbType"`
	DbName   string `json:"dbName"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	PassWord string `json:"password"`
	Charset  string `json:"charset"`
}

type SqlQuery struct {
	Query    string                   `json:"query"`
	DbConfig DbConfig                 `json:"dbConfig"`
	Results  []map[string]interface{} `json:"results"`
}
