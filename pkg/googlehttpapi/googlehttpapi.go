package googlehttpapi

import "github.com/emicklei/proto"

func DoesRPCContainGoogleHTTPAPI(r *proto.RPC) (bool bool, position int) {
	for i, option := range r.Options {
		if option.Name == Name {
			return true, i
		}
	}
	return false, 0
}

func DoesServiceContainGoogleHTTPAPIRPCs(r *proto.Service) bool {
	for _, e := range r.Elements {
		switch e.(type) {
		case *proto.RPC:
			if doesContain, _ := DoesRPCContainGoogleHTTPAPI(e.(*proto.RPC)); doesContain {
				return true
			}
		}
	}
	return false
}
