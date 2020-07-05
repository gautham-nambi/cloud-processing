package resources

type Configuration struct {
	Postgres *PostgresConf `yaml:"postgres"`
}

type PostgresConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (c *Configuration) GetPostgresConf() *PostgresConf {
	return &PostgresConf{
		Host:     c.Postgres.Host,
		Port:     c.Postgres.Port,
		Database: c.Postgres.Database,
		User:     c.Postgres.User,
		Password: c.Postgres.Password,
	}
}

func (c *PostgresConf) GetPostgresConnectionString() string {
	return "user=" + c.User +
		" password=" + c.Password +
		" host=" + c.Host +
		" dbname=" + c.Database +
		" port=" + c.Port +
		" sslmode=disable"
}
