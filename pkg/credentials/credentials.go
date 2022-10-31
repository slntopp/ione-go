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
package credentials

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	accountspb "github.com/slntopp/nocloud/pkg/registry/proto/accounts"
	sc "github.com/slntopp/nocloud/pkg/settings/client"
	pb "github.com/slntopp/nocloud/pkg/settings/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func SetupSettingsClient(logger *zap.Logger, settingsClient pb.SettingsServiceClient, internal_token string) {
	sc.Setup(
		logger, metadata.AppendToOutgoingContext(
			context.Background(), "authorization", "bearer "+internal_token,
		), &settingsClient,
	)

	_GetWHMCSConfig(&WHMCSConfig{
		Api: "", User: "api", Pass: "",
	})
}

type Link struct {
	From driver.DocumentID `json:"_from"`
	To   driver.DocumentID `json:"_to"`
	Type string            `json:"type"`
	Role string            `json:"role"`

	driver.DocumentMeta
}

type Credentials interface {
	// Check if given authorization data are mapped
	// to existent Credentials
	Authorize(...string) bool
	// Return Credentials type
	Type() string

	// Find Credentials in database by authorisation data and Unmarshall it's data into struct
	Find(context.Context, driver.Database) bool
	// Find Credentials in database by document key and Unmarshall it's data into struct
	FindByKey(context.Context, driver.Collection, string) error

	// Set Logger for Credentials methods
	SetLogger(*zap.Logger)
}

func Determine(auth_type string) (cred Credentials, ok bool) {
	switch auth_type {
	case "standard":
		return &StandardCredentials{}, true
	case "whmcs":
		return &WHMCSCredentials{}, true
	default:
		return nil, false
	}
}

func Find(ctx context.Context, db driver.Database, log *zap.Logger, auth_type string, args ...string) (cred Credentials, err error) {
	var ok bool
	switch auth_type {
	case "standard":
		cred = &StandardCredentials{Username: args[0]}
	case "whmcs":
		cred = &WHMCSCredentials{Email: args[0]}
	default:
		return nil, errors.New("unknown auth type")
	}

	cred.SetLogger(log)

	ok = cred.Find(ctx, db)
	if !ok {
		return nil, errors.New("couldn't find credentials")
	}

	if cred.Authorize(args...) {
		return cred, nil
	}

	return nil, errors.New("couldn't authorize")
}

func MakeCredentials(credentials *accountspb.Credentials, log *zap.Logger) (Credentials, error) {
	var cred Credentials
	var err error
	switch credentials.Type {
	case "standard":
		cred, err = NewStandardCredentials(credentials.Data)
	case "whmcs":
		cred, err = NewWHMCSCredentials(credentials.Data)
	default:
		return nil, errors.New("auth type is wrong")
	}

	cred.SetLogger(log)

	return cred, err
}
