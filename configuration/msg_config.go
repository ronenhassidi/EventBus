package configuration

type MsgData struct {
	MsgId   string `yaml:"msgId" json:"msgId"`
	Station string `yaml:"station" json:"station"`
	MsgType string `yaml:"msgType" json:"msgType"`
	ApiMsg  ApiMsg `yaml:"apiMsg" json:"apiMsg"`
	DbMsg   DbMsg  `yaml:"dbMsg" json:"dbMsg"`
	FsMsg   FsMsg  `yaml:"fsMsg" json:"fsMsg"`
}

type ApiMsg struct {
	Params        []KeyValue             `yaml:"params" json:"params"`
	Authorization Authorization          `yaml:"authorization" json:"authorization"`
	Headers       []KeyValue             `yaml:"headers" json:"headers"`
	Body          map[string]interface{} `yaml:"body" json:"body"`
	MethodType    string                 `yaml:"methodType" json:"methodType"`
	Url           string                 `yaml:"url" json:"url"`
	//ApiResponse   ApiResponse            `yaml:"apiResponse" json:"apiResponse"`
	ApiResponse map[string]interface{} `yaml:"apiResponse" json:"apiResponse"`
}

type KeyValue struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

type KeyValueType struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
	Type  string `yaml:"type" json:"type"`
}

type Authorization struct {
	Type  string `yaml:"type" json:"type"`
	Token string `yaml:"token" json:"token"`
}

type ApiResponse struct {
	Status        string     `yaml:"status" json:"status"`
	Message       string     `yaml:"message" json:"message"`
	KeyValueArray []KeyValue `yaml:"keyValueArray" json:"keyValueArray"`
}

type DbMsg struct {
	Sql          string       `yaml:"body" json:"sql"`
	DbConnection DbConnection `yaml:"dbConnection" json:"dbConnection"`
}

type DbConnection struct {
	User       string     `yaml:"user" json:"user"`
	Password   string     `yaml:"password" json:"password"`
	Server     string     `yaml:"server" json:"server"`
	Db         string     `yaml:"db" json:"db"`
	DbResponse DbResponse `yaml:"dbResponse" json:"dbResponse"`
}

type FsMsg struct {
	Script       string       `yaml:"script" json:"script"`
	DbConnection DbConnection `yaml:"dbConnection" json:"dbConnection"`
	FsResponse   FsResponse   `yaml:"fsResponse" json:"fsResponse"`
}

type DbResponse struct {
	Status  string     `yaml:"status" json:"status"`
	Message string     `yaml:"message" json:"message"`
	Header  []KeyValue `yaml:"headers" json:"headers"`
}

type FsResponse struct {
	Status  string     `yaml:"status" json:"status"`
	Message string     `yaml:"message" json:"message"`
	Header  []KeyValue `yaml:"headers" json:"headers"`
}
