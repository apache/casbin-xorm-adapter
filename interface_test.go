// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xormadapter

import (
	"testing"

	"github.com/casbin/casbin/v3/persist"
)

// TestAdapterInterface verifies that Adapter implements persist.Adapter interface for Casbin v3
func TestAdapterInterface(t *testing.T) {
	var _ persist.Adapter = (*Adapter)(nil)
}

// TestCasbinV3Compatibility verifies that the adapter type is compatible with Casbin v3
func TestCasbinV3Compatibility(t *testing.T) {
	// Create a dummy adapter (won't connect to DB in this test)
	a := &Adapter{
		driverName:     "mysql",
		dataSourceName: "root:@tcp(127.0.0.1:3306)/",
		dbSpecified:    false,
		tableName:      "casbin_rule",
	}

	// This verifies the adapter implements persist.Adapter interface correctly
	// If LoadPolicy method was missing, this would fail to compile
	var adapter persist.Adapter = a

	// Verify the type assertion works
	if adapter == nil {
		t.Fatal("Adapter should not be nil")
	}

	t.Log("Successfully verified Adapter implements persist.Adapter for Casbin v3")
}
