// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/gleke/hexya/src/models/security"
	"github.com/gleke/pool/h"
)

func init() {
	h.Filter().Methods().AllowAllToGroup(security.GroupEveryone)
}
