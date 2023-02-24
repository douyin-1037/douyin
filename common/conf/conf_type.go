package conf

var (
	Server   *ServerConfig
	COS      *COSConfig
	Database *DatabaseConfig
	JWT      *JWTConfig
	Redis    *RedisConfig
	Pulsar   *PulsarConfig
	MongoDB  *MongoDBConfig
)

type ServerConfig struct {
	RunMode            string
	HttpPort           string
	UserServiceAddr    string
	VideoServiceAddr   string
	CommentServiceAddr string
	MessageServiceAddr string
	Timeout            int
	EtcdAddress        string
	FeedCount          int64
}

type COSConfig struct {
	VideoBucket string
	CoverBucket string
	SecretID    string
	SecretKey   string
}

type DatabaseConfig struct {
	DBType    string
	UserName  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime string
}

type RedisConfig struct {
	Address        string
	MaxIdle        int
	MaxActive      int
	ExpireTime     int
	MaxRandAddTime int
	BloomOpen      bool
}

type PulsarConfig struct {
	URL               string
	OperationTimeout  int
	ConnectionTimeout int
}

type MongoDBConfig struct {
	URL      string
	DataBase string
}
