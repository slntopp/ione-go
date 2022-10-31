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
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type StandardCredentials struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`

	log *zap.Logger
	driver.DocumentMeta
}

func NewStandardCredentials(data []string) (Credentials, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("some credentials data is missing, expected data length to be 2, got: %d", len(data))
	}

	username, password := data[0], data[1]
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return &StandardCredentials{
		Username:     username,
		PasswordHash: string(bytes),
	}, err
}

// Authorize method for StandardCredentials assumes that args consist of username and password stored at 0 and 1 accordingly
func (c *StandardCredentials) Authorize(args ...string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.PasswordHash), []byte(args[1]))
	return err == nil
}

func (*StandardCredentials) Type() string {
	return "standard"
}

func (c *StandardCredentials) SetLogger(log *zap.Logger) {
	c.log = log.Named("Standard")
	c.log.Debug("Logger is now set")
}

func (cred *StandardCredentials) Find(ctx context.Context, db driver.Database) bool {
	query := `FOR cred IN @@credentials FILTER cred.username == @username RETURN cred`
	c, err := db.Query(ctx, query, map[string]interface{}{
		"username":     cred.Username,
		"@credentials": schema.CREDENTIALS_COL,
	})
	if err != nil {
		return false
	}
	defer c.Close()

	_, err = c.ReadDocument(ctx, &cred)
	return err == nil
}

func (cred *StandardCredentials) FindByKey(ctx context.Context, col driver.Collection, key string) error {
	_, err := col.ReadDocument(ctx, key, cred)
	return err
}
