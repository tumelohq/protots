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