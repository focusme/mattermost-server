// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestUserAccessTokenSearchJson(t *testing.T) {
	userAccessTokenSearch := UserAccessTokenSearch{Term: NewId()}
	json := userAccessTokenSearch.ToJson()
	ruserAccessTokenSearch := UserAccessTokenSearchFromJson(strings.NewReader(json))
	require.Equal(t, userAccessTokenSearch.Term, ruserAccessTokenSearch.Term, "Terms do not match")
}