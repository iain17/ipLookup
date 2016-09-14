package environment

import (
	//"labix.org/v2/mgo"
	//"gopkg.in/redis.v3"
	"github.com/iain17/ipLookup/src/config"
	"github.com/oschwald/geoip2-golang"
)

type Environment struct {
	Config		*config.Config
	Maxmind		*geoip2.Reader
	//MongoDB *mgo.Session
	//Redis    *redis.Client
}

var Env *Environment

func SetEnvironment(env *Environment) {
	Env = env
}