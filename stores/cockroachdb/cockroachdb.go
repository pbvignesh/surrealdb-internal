// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cockroachdb

import (
	"github.com/abcum/surreal/cnf"
	"github.com/abcum/surreal/stores"
	"github.com/cockroachdb/cockroach/client"
	"github.com/cockroachdb/cockroach/util/stop"
)

func init() {
	stores.Register("cockroachdb", New)
}

type Store struct {
	ctx cnf.Context
	db  client.DB
}

func New(ctx cnf.Context) (stores.Store, error) {

	stopper := stop.NewStopper()

	defer stopper.Stop()

	db, err := client.Open(stopper, ctx.DbPath)

	if err != nil {
		stopper.Stop()
		return nil, err
	}

	store := Store{ctx: ctx, db: *db}

	return &store, nil

}

func (store *Store) Get(key interface{}) stores.KeyValue {
	return stores.KeyValue{}
}

func (store *Store) Put(key, val interface{}) error {
	return nil
}

func (store *Store) Del(key interface{}) error {
	return nil
}

func (store *Store) Scan(beg, end interface{}, max int64) []stores.KeyValue {
	return []stores.KeyValue{}
}
