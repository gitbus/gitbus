// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gitbus/gitbus/models"
	"github.com/gitbus/gitbus/pkg/context"
	"github.com/gitbus/gitbus/pkg/setting"
	"github.com/gitbus/gitbus/routes"
)

const (
	ORGS = "admin/org/list"
)

func Organizations(c *context.Context) {
	c.Data["Title"] = c.Tr("admin.organizations")
	c.Data["PageIsAdmin"] = true
	c.Data["PageIsAdminOrganizations"] = true

	routes.RenderUserSearch(c, &routes.UserSearchOptions{
		Type:     models.USER_TYPE_ORGANIZATION,
		Counter:  models.CountOrganizations,
		Ranger:   models.Organizations,
		PageSize: setting.UI.Admin.OrgPagingNum,
		OrderBy:  "id ASC",
		TplName:  ORGS,
	})
}
