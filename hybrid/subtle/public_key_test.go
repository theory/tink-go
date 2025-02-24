// Copyright 2022 Google LLC
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

package subtle_test

import (
	"bytes"
	"testing"

	"google.golang.org/protobuf/proto"
	"github.com/tink-crypto/tink-go/v2/hybrid"
	"github.com/tink-crypto/tink-go/v2/hybrid/subtle"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"github.com/tink-crypto/tink-go/v2/subtle/random"
	"github.com/tink-crypto/tink-go/v2/testutil"
	hpkepb "github.com/tink-crypto/tink-go/v2/proto/hpke_go_proto"
	tinkpb "github.com/tink-crypto/tink-go/v2/proto/tink_go_proto"
)

func TestHPKEPublicKeySerialization(t *testing.T) {
	// Obtain private and public keyset handles via key template.
	keyTemplate := hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Raw_Key_Template()
	privHandle, err := keyset.NewHandle(keyTemplate)
	if err != nil {
		t.Fatalf("NewHandle(%v) err = %v, want nil", keyTemplate, err)
	}
	pubHandle, err := privHandle.Public()
	if err != nil {
		t.Fatalf("Public() err = %v, want nil", err)
	}

	// Export public key as bytes.
	pubKeyBytes, err := subtle.SerializePrimaryPublicKey(pubHandle, keyTemplate)
	if err != nil {
		t.Fatalf("SerializePrimaryPublicKey(%v) err = %v, want nil", pubHandle, err)
	}

	// Import public key bytes as keyset handle.
	gotPubHandle, err := subtle.KeysetHandleFromSerializedPublicKey(pubKeyBytes, keyTemplate)
	if err != nil {
		t.Fatalf("KeysetHandleFromSerializedPublicKey(%v, %v) err = %v, want nil", pubKeyBytes, keyTemplate, err)
	}

	plaintext := random.GetRandomBytes(200)
	ctxInfo := random.GetRandomBytes(100)

	// Encrypt with public keyset handle constructed from public key bytes.
	enc, err := hybrid.NewHybridEncrypt(gotPubHandle)
	if err != nil {
		t.Fatalf("NewHybridEncrypt(%v) err = %v, want nil", gotPubHandle, err)
	}
	ciphertext, err := enc.Encrypt(plaintext, ctxInfo)
	if err != nil {
		t.Fatalf("Encrypt(%x, %x) err = %v, want nil", plaintext, ctxInfo, err)
	}

	// Decrypt with original private keyset handle.
	dec, err := hybrid.NewHybridDecrypt(privHandle)
	if err != nil {
		t.Fatalf("NewHybridDecrypt(%v) err = %v, want nil", privHandle, err)
	}
	gotPlaintext, err := dec.Decrypt(ciphertext, ctxInfo)
	if err != nil {
		t.Fatalf("Decrypt(%x, %x) err = %v, want nil", plaintext, ctxInfo, err)
	}
	if !bytes.Equal(gotPlaintext, plaintext) {
		t.Errorf("Decrypt(%x, %x) = %x, want %x", plaintext, ctxInfo, gotPlaintext, plaintext)
	}
}

func TestSerializePrimaryPublicKeyInvalidTemplateFails(t *testing.T) {
	keyTemplate := hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Raw_Key_Template()
	privHandle, err := keyset.NewHandle(keyTemplate)
	if err != nil {
		t.Fatalf("NewHandle(%v) err = %v, want nil", keyTemplate, err)
	}
	pubHandle, err := privHandle.Public()
	if err != nil {
		t.Fatalf("Public() err = %v, want nil", err)
	}

	tests := []struct {
		name     string
		template *tinkpb.KeyTemplate
	}{
		{"AES_128_GCM", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_128_GCM_Key_Template()},
		{"AES_128_GCM_Raw", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_128_GCM_Raw_Key_Template()},
		{"AES_256_GCM", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_256_GCM_Key_Template()},
		{"AES_256_GCM_Raw", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_256_GCM_Raw_Key_Template()},
		{"CHACHA20_POLY1305", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Key_Template()},
		{"invalid type URL", &tinkpb.KeyTemplate{
			TypeUrl:          "type.googleapis.com/google.crypto.tink.EciesAeadHkdfPrivateKey",
			Value:            keyTemplate.GetValue(),
			OutputPrefixType: tinkpb.OutputPrefixType_RAW,
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := subtle.SerializePrimaryPublicKey(pubHandle, test.template); err == nil {
				t.Errorf("SerializePrimaryPublicKey(%v, %v) err = nil, want error", pubHandle, test.template)
			}
		})
	}
}

func TestSerializePrimaryPublicKeyInvalidKeyFails(t *testing.T) {
	// Build valid key data.
	keyTemplate := hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Raw_Key_Template()
	privHandle, err := keyset.NewHandle(keyTemplate)
	if err != nil {
		t.Fatalf("NewHandle(%v) err = %v, want nil", keyTemplate, err)
	}
	pubHandle, err := privHandle.Public()
	if err != nil {
		t.Fatalf("Public() err = %v, want nil", err)
	}
	pubKeyBytes, err := subtle.SerializePrimaryPublicKey(pubHandle, keyTemplate)
	if err != nil {
		t.Fatalf("SerializePrimaryPublicKey(%v, %v) err = %v, want nil", pubHandle, keyTemplate, err)
	}
	typeURL := "type.googleapis.com/google.crypto.tink.HpkePublicKey"
	validKD := mustCreateKeyDataFromBytes(t, pubKeyBytes, hpkepb.HpkeAead_CHACHA20_POLY1305, typeURL)

	// Build key data with invalid type URL.
	invalidTypeURL := "type.googleapis.com/google.crypto.tink.InvalidTypeURL"
	invalidTypeURLKD := mustCreateKeyDataFromBytes(t, pubKeyBytes, hpkepb.HpkeAead_CHACHA20_POLY1305, invalidTypeURL)

	// Build key data with invalid HPKE params.
	randomPubKeyBytes := random.GetRandomBytes(32)
	invalidAEAD := hpkepb.HpkeAead_AES_128_GCM
	invalidParamsKD := mustCreateKeyDataFromBytes(t, randomPubKeyBytes, invalidAEAD, typeURL)

	tests := []struct {
		name         string
		primaryKeyID uint32
		key          *tinkpb.Keyset_Key
	}{
		{
			"invalid prefix type",
			123,
			testutil.NewKey(validKD, tinkpb.KeyStatusType_ENABLED, 123, tinkpb.OutputPrefixType_TINK),
		},
		{
			"invalid type URL",
			123,
			testutil.NewKey(invalidTypeURLKD, tinkpb.KeyStatusType_ENABLED, 123, tinkpb.OutputPrefixType_RAW),
		},
		{
			"invalid HPKE params",
			123,
			testutil.NewKey(invalidParamsKD, tinkpb.KeyStatusType_ENABLED, 123, tinkpb.OutputPrefixType_RAW),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ks := testutil.NewKeyset(test.primaryKeyID, []*tinkpb.Keyset_Key{test.key})
			handle, err := keyset.NewHandleWithNoSecrets(ks)
			if err != nil {
				t.Fatalf("NewHandleWithNoSecrets(%v) err = %v, want nil", ks, err)
			}
			if _, err := subtle.SerializePrimaryPublicKey(handle, keyTemplate); err == nil {
				t.Errorf("SerializePrimaryPublicKey(%v, %v) err = nil, want error", handle, keyTemplate)
			}
		})
	}
}

func TestSerializePrimaryPublicKeyFailsWithEmptyHandle(t *testing.T) {
	handle := &keyset.Handle{}
	keyTemplate := hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Raw_Key_Template()
	if _, err := subtle.SerializePrimaryPublicKey(handle, keyTemplate); err == nil {
		t.Errorf("SerializePrimaryPublicKey(%v, %v) err = nil, want error", handle, keyTemplate)
	}
}

func TestKeysetHandleFromSerializedPublicKeyInvalidTemplateFails(t *testing.T) {
	keyTemplate := hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Raw_Key_Template()
	privHandle, err := keyset.NewHandle(keyTemplate)
	if err != nil {
		t.Fatalf("NewHandle(%v) err = %v, want nil", keyTemplate, err)
	}
	pubHandle, err := privHandle.Public()
	if err != nil {
		t.Fatalf("Public() err = %v, want nil", err)
	}
	pubKeyBytes, err := subtle.SerializePrimaryPublicKey(pubHandle, keyTemplate)
	if err != nil {
		t.Fatalf("SerializePrimaryPublicKey(%v) err = %v, want nil", pubHandle, err)
	}

	tests := []struct {
		name     string
		template *tinkpb.KeyTemplate
	}{
		{"AES_128_GCM", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_128_GCM_Key_Template()},
		{"AES_128_GCM_Raw", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_128_GCM_Raw_Key_Template()},
		{"AES_256_GCM", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_256_GCM_Key_Template()},
		{"AES_256_GCM_Raw", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_AES_256_GCM_Raw_Key_Template()},
		{"CHACHA20_POLY1305", hybrid.DHKEM_X25519_HKDF_SHA256_HKDF_SHA256_CHACHA20_POLY1305_Key_Template()},
		{"invalid type URL", &tinkpb.KeyTemplate{
			TypeUrl:          "type.googleapis.com/google.crypto.tink.EciesAeadHkdfPrivateKey",
			Value:            keyTemplate.GetValue(),
			OutputPrefixType: tinkpb.OutputPrefixType_RAW,
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := subtle.KeysetHandleFromSerializedPublicKey(pubKeyBytes, test.template); err == nil {
				t.Errorf("KeysetHandleFromSerializedPublicKey(%v, %v) err = nil, want error", pubKeyBytes, test.template)
			}
		})
	}
}

func mustCreateKeyDataFromBytes(t *testing.T, pubKeyBytes []byte, aeadID hpkepb.HpkeAead, typeURL string) *tinkpb.KeyData {
	t.Helper()

	pubKey := &hpkepb.HpkePublicKey{
		Version: 0,
		Params: &hpkepb.HpkeParams{
			Kem:  hpkepb.HpkeKem_DHKEM_X25519_HKDF_SHA256,
			Kdf:  hpkepb.HpkeKdf_HKDF_SHA256,
			Aead: aeadID,
		},
		PublicKey: pubKeyBytes,
	}

	serializedPubKey, err := proto.Marshal(pubKey)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", pubKey, err)
	}
	return testutil.NewKeyData(typeURL, serializedPubKey, tinkpb.KeyData_ASYMMETRIC_PUBLIC)
}
