// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/gogits/go-gogs-client"

	"github.com/gitbus/gitbus/pkg/context"
	"github.com/gitbus/gitbus/routes/api/v1/user"
	"github.com/gitbus/gitbus/routes/api/v1/org"
)

// https://github.com/gogits/go-gogs-client/wiki/Administration-Organizations#create-a-new-organization
func CreateOrg(c *context.APIContext, form api.CreateOrgOption) {
	org.CreateOrgForUser(c, form, user.GetUserByParams(c))
}
