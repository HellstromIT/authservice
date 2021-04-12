package database

type Database interface {
	Connect(Host string, Port string, DBName string, Password string, User string, SslMode string, TimeZone string)
}
