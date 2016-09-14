package lookup

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/iain17/ipLookup/src/utils/logger"
	"net"
	"github.com/iain17/ipLookup/src/environment"
	"strings"
)

func New() Lookup {
	return &lookup{}
}

type lookupResult struct {
	Success bool `json:"success"`
	Result interface{} `json:"result"`
}

func (u *lookup) GetIp(w rest.ResponseWriter, r *rest.Request) {
	preProcessedIp := r.PathParam("ip")
	preProcessedIp = strings.Replace(preProcessedIp, "_", ".", 3)


	ip := net.ParseIP( preProcessedIp )

	if ip == nil {
		w.WriteJson(lookupResult{
			Success: false,
			Result: "The submitted ip is invalid.",
		})
		return
	}

	record, err := environment.Env.Maxmind.City(ip)
	if err != nil {
		logger.Fatal(err)
		w.WriteJson(lookupResult{
			Success: false,
			Result: "IP could not be found. Sorry!",
		})
		return
	}

	//Spit back that lovely data we enjoy so much!
	w.WriteJson(lookupResult{
		Success: true,
		Result: record,
	})
}