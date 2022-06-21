/*
env.go file is intended to hold all environment variables in one place for
ease of setup on another machine
*/
package main

import (
	"os"
)

var (
	servLoc       = os.Getenv("server location") // local or remote
	domainName    = os.Getenv("domain name")     // set on remote, will use localhost:8080 on local machine
	EntityName    = os.Getenv("entity name")     // ex: Goggle not https://google.com
	useHTTPS      = os.Getenv("use HTTPS")       // for email setup for http or https, use "true" or "false"
	MySQLusername = os.Getenv("mysqlUsername")
	MySQLpassword = os.Getenv("mysqlPassword")
	MySQLaddress  = os.Getenv("mysqlAddress")
	MySQLport     = os.Getenv("mysqlPort")
	dbName        = os.Getenv("mysqlDBname")
	fromEmail     = os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	SMTPpassword  = os.Getenv("SMTPpwd")       // ex: "ieiemcjdkejspqz"
)
