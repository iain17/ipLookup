package lookup

import "github.com/ant0ine/go-json-rest/rest"

type lookup struct {

}

type Lookup interface {
	GetIp(w rest.ResponseWriter, r *rest.Request)
}