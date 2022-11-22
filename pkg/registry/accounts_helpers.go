/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package registry

import (
	"context"

	accountspb "github.com/slntopp/nocloud-proto/registry/accounts"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	sc "github.com/slntopp/nocloud/pkg/settings/client"
	"go.uber.org/zap"
)

func MakeAccountMessage(acc graph.Account) *accountspb.Account {
	return &accountspb.Account{Uuid: acc.Key, Title: acc.Title}
}

func MergeMaps[K comparable](old map[K]interface{}, new map[K]interface{}) map[K]interface{} {
	result := make(map[K]interface{})
	for key, ov := range old {
		result[key] = ov
	}

	for key, nv := range new {
		if nv == nil {
			delete(result, key)
			continue
		}

		result[key] = nv
	}

	return result
}

const accountPostCreateSettingsKey = "post-create-account"

type AccountPostCreateSettings struct {
	CreateNamespace bool `json:"create-ns"`
}

var defaultSettings = &sc.Setting[AccountPostCreateSettings]{
	Value:       AccountPostCreateSettings{CreateNamespace: true},
	Description: "Post Account Creation Actions",
	Public:      false,
}

func (s *AccountsServiceServer) PostCreateActions(ctx context.Context, account graph.Account) {
	log := s.log.Named("PostCreateActions")
	var settings AccountPostCreateSettings
	if scErr := sc.Fetch(accountPostCreateSettingsKey, &settings, defaultSettings); scErr != nil {
		log.Warn("Cannot fetch settings", zap.Error(scErr))
	}

	if settings.CreateNamespace {
		_CreatePersonalNamespace(ctx, log, s.ns_ctrl, account)
	}
}

func _CreatePersonalNamespace(ctx context.Context, log *zap.Logger, ns_ctrl graph.NamespacesController, account graph.Account) {
	personal_ctx := context.WithValue(ctx, nocloud.NoCloudAccount, account.GetUuid())

	if _, err := ns_ctrl.Create(personal_ctx, account.GetTitle()); err != nil {
		log.Warn("Cannot create a namespace for new Account", zap.String("account", account.GetUuid()), zap.Error(err))
		return
	}
}
