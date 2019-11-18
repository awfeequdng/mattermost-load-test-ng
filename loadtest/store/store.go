// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import (
	"github.com/mattermost/mattermost-server/model"
)

type UserStore interface {
	User() *model.User
}

type MutableUserStore interface {
	UserStore
	SetUser(user *model.User) error
}