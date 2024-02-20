package depd

import (
	"github.com/maxmzkr/downgrade/depa"
	"github.com/maxmzkr/downgrade/depb"
)

func Version() {
	depa.Version()
	depb.Version()
}
