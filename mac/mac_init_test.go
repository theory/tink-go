// Copyright 2023 Google LLC
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
//
////////////////////////////////////////////////////////////////////////////////

package mac_test

import (
	"testing"

	"github.com/tink-crypto/tink-go/v2/core/registry"
	"github.com/tink-crypto/tink-go/v2/testutil"
)

func TestMacInit(t *testing.T) {
	// Check that the HMAC key manager is in the global registry.
	_, err := registry.GetKeyManager(testutil.HMACTypeURL)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	// Check that the AES CMAC key manager is in the global registry.
	_, err = registry.GetKeyManager(testutil.AESCMACTypeURL)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
