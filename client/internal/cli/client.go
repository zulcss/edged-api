package cli

import (
	"fmt"
	"github.com/dghubble/sling"

	"github.com/zulcss/edged/shared/constants"
)

type Client struct {
	sling	*sling.Sling
}

func NewClient(host string) {
	var EdgedAPI = fmt.Sprintf("%s:%s%s",host, constants.Port, constants.EndPoint)
	fmt.Println(EdgedAPI)
}
