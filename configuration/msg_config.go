package configuration

type MsgData struct {
	MsgId   string `yaml:"msgId" json:"msgId"`
	Station string `yaml:"station" json:"station"`
	MsgType string `yaml:"msgType" json:"msgType"`
	ApiMsg  ApiMsg `yaml:"apiMsg,mapstructure" json:"apiMsg,mapstructure"`
	DbMsg   DbMsg  `yaml:"dbMsg,mapstructure" json:"dbMsg,mapstructure"`
	FsMsg   FsMsg  `yaml:"fsMsg,mapstructure" json:"fsMsg,mapstructure"`
}

type ApiMsg struct {
	Params        []KeyValue             `yaml:"params,mapstructure" json:"params,mapstructure"`
	Authorization Authorization          `yaml:"authorization" json:"authorization,mapstructure"`
	Headers       []KeyValue             `yaml:"headers,mapstructure" json:"headers,mapstructure"`
	Body          map[string]interface{} `yaml:"body,mapstructure" json:"body,mapstructure"`
	MethodType    string                 `yaml:"methodType" json:"methodType"`
	Url           string                 `yaml:"url" json:"url"`
	//ApiResponse   ApiResponse            `yaml:"apiResponse,mapstructure" json:"apiResponse,mapstructure"`
	ApiResponse map[string]interface{} `yaml:"apiResponse,mapstructure" json:"apiResponse,mapstructure"`
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
	KeyValueArray []KeyValue `yaml:"keyValueArray,mapstructure" json:"keyValueArray,mapstructure"`
}

type DbMsg struct {
	Sql          string       `yaml:"body" json:"sql"`
	DbConnection DbConnection `yaml:"dbConnection,mapstructure" json:"dbConnection,mapstructure"`
}

type DbConnection struct {
	User       string     `yaml:"user" json:"user"`
	Password   string     `yaml:"password" json:"password"`
	Server     string     `yaml:"server" json:"server"`
	Db         string     `yaml:"db" json:"db"`
	DbResponse DbResponse `yaml:"dbResponse,mapstructure" json:"dbResponse,mapstructure"`
}

type FsMsg struct {
	Script       string       `yaml:"script" json:"script"`
	DbConnection DbConnection `yaml:"dbConnection,mapstructure" json:"dbConnection,mapstructure"`
	FsResponse   FsResponse   `yaml:"fsResponse,mapstructure" json:"fsResponse,mapstructure"`
}

type DbResponse struct {
	Status  string     `yaml:"status" json:"status"`
	Message string     `yaml:"message" json:"message"`
	Header  []KeyValue `yaml:"headers,mapstructure" json:"headers,mapstructure"`
}

type FsResponse struct {
	Status  string     `yaml:"status" json:"status"`
	Message string     `yaml:"message" json:"message"`
	Header  []KeyValue `yaml:"headers,mapstructure" json:"headers,mapstructure"`
}
