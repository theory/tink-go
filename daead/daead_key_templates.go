// Copyright 2019 Google LLC
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

package daead

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"github.com/tink-crypto/tink-go/v2/internal/tinkerror"
	aspb "github.com/tink-crypto/tink-go/v2/proto/aes_siv_go_proto"
	tinkpb "github.com/tink-crypto/tink-go/v2/proto/tink_go_proto"
)

// AESSIVKeyTemplate is a KeyTemplate that generates a AES-SIV key.
func AESSIVKeyTemplate() *tinkpb.KeyTemplate {
	format := &aspb.AesSivKeyFormat{
		KeySize: 64,
	}
	serializedFormat, err := proto.Marshal(format)
	if err != nil {
		tinkerror.Fail(fmt.Sprintf("failed to marshal key format: %s", err))
	}
	return &tinkpb.KeyTemplate{
		TypeUrl:          "type.googleapis.com/google.crypto.tink.AesSivKey",
		OutputPrefixType: tinkpb.OutputPrefixType_TINK,
		Value:            serializedFormat,
	}
}
