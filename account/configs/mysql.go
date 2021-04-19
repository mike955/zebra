package account

type Mysql struct {
	MysqlAddr     string `yaml:"mysql_addr"`
	MysqlUsername string `yaml:"mysql_username"`
	MysqlPassword string `yaml:"mysql_password"`
	MysqlDatabase string `yaml:"mysql_database"`
}
