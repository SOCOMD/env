package env

import "os"

// Env holds all environment variables use by socomd
type Env struct {
	Staff staff
	Tsbot tsbot
}

type staff struct {
	WebHost   string
	WebPort   string
	GrpcPort  string
	Database  database
	JWTSecret string
}

type tsbot struct {
	TSHost   string
	TSPort   string
	TSUser   string
	TSPass   string
	LogFile  string
	GrpcHost string
	GrpcPort string
}

type database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

// Get returns a list of all environment variables used
func Get() (e Env) {

	e.Staff.WebHost = os.Getenv("STAFF_WEB_HOST")
	e.Staff.WebPort = os.Getenv("STAFF_WEB_PORT")
	e.Staff.GrpcPort = os.Getenv("STAFF_GRPC_PORT")
	e.Staff.JWTSecret = os.Getenv("STAFF_JWT_SECRET")
	e.Staff.Database.Host = os.Getenv("STAFF_DATABASE_HOST")
	e.Staff.Database.Port = os.Getenv("STAFF_DATABASE_PORT")
	e.Staff.Database.Name = os.Getenv("STAFF_DATABASE_NAME")
	e.Staff.Database.User = os.Getenv("STAFF_DATABASE_USER")
	e.Staff.Database.Pass = os.Getenv("STAFF_DATABASE_PASS")

	e.Tsbot.TSHost = os.Getenv("TSBOT_TS_HOST")
	e.Tsbot.TSPort = os.Getenv("TSBOT_TS_PORT")
	e.Tsbot.TSUser = os.Getenv("TSBOT_TS_USER")
	e.Tsbot.TSPass = os.Getenv("TSBOT_TS_PASS")
	e.Tsbot.GrpcHost = os.Getenv("TSBOT_GRPC_HOST")
	e.Tsbot.GrpcPort = os.Getenv("TSBOT_GRPC_PORT")
	e.Tsbot.LogFile = os.Getenv("TSBOT_LOGFILE")

	return
}

/*
STAFF_WEB_HOST=localhost
STAFF_WEB_PORT=8080
STAFF_GRPC_PORT=8081
STAFF_DATABASE_HOST=0.0.0.0
STAFF_DATABASE_PORT=3306
STAFF_DATABASE_NAME=socomd
STAFF_DATABASE_USER=root
STAFF_DATABASE_PASS=
STAFF_JWT_SECRET=myshlong

TSBOT_TS_HOST=ts.socomd.com \
TSBOT_TS_PORT=10099 \
TSBOT_TS_USER=zanven \
TSBOT_TS_PASS=RsGElP6o \
TSBOT_LOGFILE= \
TSBOT_GRPC_PORT=9200 \
TSBOT_GRPC_HOST=localhost \
*/
