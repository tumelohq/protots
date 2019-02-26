package googlehttpapi

/*
import (
	"fmt"
	"github.com/emicklei/proto"
	"net/http"
)

type GoogleHTTPMessage struct {
	httpType string

}

func ParsingGoogleHTTPMessage(r *proto.RPC) (*GoogleHTTPMessage, error) {
	does, position := DoesRPCContainGoogleHTTPAPI(r)
	if !does {
		return nil, fmt.Errorf("The following proto does not contain an %s option\n%+v\n", Name, r)
	}

	httpOption, ok := r.Options[position]
	if !ok


}

*/