package conf

var Conn  = map[string]map[string]map[string]string{
	"mysql": {
		"gin_study" : {
			"host": "127.0.0.1",
			"port": "3306",
			"dbname": "gin_study",
			"user": "root",
			"pwd": "root",
		},
	},
	"redis": {
		"default": {
			"host": "127.0.0.1",
			"port": "6379",
			"pwd": "",
		},
	},
}