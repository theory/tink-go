// Copyright 2018 Google LLC
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

package signature

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"github.com/tink-crypto/tink-go/v2/internal/tinkerror"
	commonpb "github.com/tink-crypto/tink-go/v2/proto/common_go_proto"
	ecdsapb "github.com/tink-crypto/tink-go/v2/proto/ecdsa_go_proto"
	rsppb "github.com/tink-crypto/tink-go/v2/proto/rsa_ssa_pkcs1_go_proto"
	rspsspb "github.com/tink-crypto/tink-go/v2/proto/rsa_ssa_pss_go_proto"
	tinkpb "github.com/tink-crypto/tink-go/v2/proto/tink_go_proto"
)

// This file contains pre-generated KeyTemplates for Signer and Verifier.
// One can use these templates to generate new Keysets.

const (
	ed25519SignerTypeURL     = "type.googleapis.com/google.crypto.tink.Ed25519PrivateKey"
	ecdsaSignerTypeURL       = "type.googleapis.com/google.crypto.tink.EcdsaPrivateKey"
	rsaSSAPKCS1SignerTypeURL = "type.googleapis.com/google.crypto.tink.RsaSsaPkcs1PrivateKey"
	rsaSSAPSSSignerTypeURL   = "type.googleapis.com/google.crypto.tink.RsaSsaPssPrivateKey"
)

// ECDSAP256KeyTemplate is a KeyTemplate that generates a new ECDSA private key with the following parameters:
//   - Hash function: SHA256
//   - Curve: NIST P-256
//   - Signature encoding: DER
//   - Output prefix type: TINK
func ECDSAP256KeyTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA256,
		commonpb.EllipticCurveType_NIST_P256,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_TINK)
}

// ECDSAP256KeyWithoutPrefixTemplate is a KeyTemplate that generates a new ECDSA private key with the following
// parameters:
//   - Hash function: SHA256
//   - Curve: NIST P-256
//   - Signature encoding: DER
//   - Output prefix type: RAW
//
// Note that this template uses a different encoding than ESDSA_P256_RAW in Tinkey.
func ECDSAP256KeyWithoutPrefixTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA256,
		commonpb.EllipticCurveType_NIST_P256,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_RAW)
}

// ECDSAP256RawKeyTemplate is a KeyTemplate that generates a new ECDSA private key with the following
// parameters:
//   - Hash function: SHA256
//   - Curve: NIST P-256
//   - Signature encoding: IEEE_P1363
//   - Output prefix type: RAW
func ECDSAP256RawKeyTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA256,
		commonpb.EllipticCurveType_NIST_P256,
		ecdsapb.EcdsaSignatureEncoding_IEEE_P1363,
		tinkpb.OutputPrefixType_RAW)
}

// ECDSAP384SHA384KeyTemplate is a KeyTemplate that generates a new ECDSA private key with the following parameters:
//   - Hash function: SHA384
//   - Curve: NIST P-384
//   - Signature encoding: DER
//   - Output prefix type: TINK
func ECDSAP384SHA384KeyTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA384,
		commonpb.EllipticCurveType_NIST_P384,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_TINK)
}

// ECDSAP384SHA384KeyWithoutPrefixTemplate is a KeyTemplate that generates a new ECDSA private key with the following parameters:
//   - Hash function: SHA384
//   - Curve: NIST P-384
//   - Signature encoding: DER
//   - Output prefix type: RAW
func ECDSAP384SHA384KeyWithoutPrefixTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA384,
		commonpb.EllipticCurveType_NIST_P384,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_RAW)
}

// ECDSAP384SHA512KeyTemplate is a KeyTemplate that generates a new ECDSA private key with the following parameters:
//   - Hash function: SHA512
//   - Curve: NIST P-384
//   - Signature encoding: DER
//   - Output prefix type: TINK
func ECDSAP384SHA512KeyTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA512,
		commonpb.EllipticCurveType_NIST_P384,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_TINK)
}

// ECDSAP384KeyWithoutPrefixTemplate is a KeyTemplate that generates a new ECDSA private key with the following
// parameters:
//   - Hash function: SHA512
//   - Curve: NIST P-384
//   - Signature encoding: DER
//   - Output prefix type: RAW
func ECDSAP384KeyWithoutPrefixTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA512,
		commonpb.EllipticCurveType_NIST_P384,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_RAW)
}

// ECDSAP521KeyTemplate is a KeyTemplate that generates a new ECDSA private key with the following parameters:
//   - Hash function: SHA512
//   - Curve: NIST P-521
//   - Signature encoding: DER
//   - Output prefix type: TINK
func ECDSAP521KeyTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA512,
		commonpb.EllipticCurveType_NIST_P521,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_TINK)
}

// ECDSAP521KeyWithoutPrefixTemplate is a KeyTemplate that generates a new ECDSA private key with the following
// parameters:
//   - Hash function: SHA512
//   - Curve: NIST P-521
//   - Signature encoding: DER
//   - Output prefix type: RAW
func ECDSAP521KeyWithoutPrefixTemplate() *tinkpb.KeyTemplate {
	return createECDSAKeyTemplate(commonpb.HashType_SHA512,
		commonpb.EllipticCurveType_NIST_P521,
		ecdsapb.EcdsaSignatureEncoding_DER,
		tinkpb.OutputPrefixType_RAW)
}

// createECDSAKeyTemplate creates a KeyTemplate containing a EcdasKeyFormat
// with the given parameters.
func createECDSAKeyTemplate(hashType commonpb.HashType, curve commonpb.EllipticCurveType, encoding ecdsapb.EcdsaSignatureEncoding, prefixType tinkpb.OutputPrefixType) *tinkpb.KeyTemplate {
	params := &ecdsapb.EcdsaParams{
		HashType: hashType,
		Curve:    curve,
		Encoding: encoding,
	}
	format := &ecdsapb.EcdsaKeyFormat{Params: params}
	serializedFormat, err := proto.Marshal(format)
	if err != nil {
		tinkerror.Fail(fmt.Sprintf("failed to marshal key format: %s", err))
	}
	return &tinkpb.KeyTemplate{
		TypeUrl:          ecdsaSignerTypeURL,
		Value:            serializedFormat,
		OutputPrefixType: prefixType,
	}
}

// ED25519KeyTemplate is a KeyTemplate that generates a new ED25519 private key.
func ED25519KeyTemplate() *tinkpb.KeyTemplate {
	return &tinkpb.KeyTemplate{
		TypeUrl:          ed25519SignerTypeURL,
		OutputPrefixType: tinkpb.OutputPrefixType_TINK,
	}
}

// ED25519KeyWithoutPrefixTemplate is a KeyTemplate that generates a new ED25519 private key.
func ED25519KeyWithoutPrefixTemplate() *tinkpb.KeyTemplate {
	return &tinkpb.KeyTemplate{
		TypeUrl:          ed25519SignerTypeURL,
		OutputPrefixType: tinkpb.OutputPrefixType_RAW,
	}
}

func create_RSA_SSA_PKCS1_Template(prefixType tinkpb.OutputPrefixType, hashType commonpb.HashType, modulusSizeInBits uint32) *tinkpb.KeyTemplate {
	keyFormat := &rsppb.RsaSsaPkcs1KeyFormat{
		Params: &rsppb.RsaSsaPkcs1Params{
			HashType: hashType,
		},
		ModulusSizeInBits: modulusSizeInBits,
		PublicExponent:    []byte{0x01, 0x00, 0x01},
	}
	serializedFormat, err := proto.Marshal(keyFormat)
	if err != nil {
		tinkerror.Fail(fmt.Sprintf("failed to marshal key format: %s", err))
	}
	return &tinkpb.KeyTemplate{
		TypeUrl:          rsaSSAPKCS1SignerTypeURL,
		OutputPrefixType: prefixType,
		Value:            serializedFormat,
	}
}

func create_RSA_SSA_PSS_Template(prefixType tinkpb.OutputPrefixType, hashType commonpb.HashType, saltLength int32, modulusSizeInBits uint32) *tinkpb.KeyTemplate {
	keyFormat := &rspsspb.RsaSsaPssKeyFormat{
		Params: &rspsspb.RsaSsaPssParams{
			SigHash:    hashType,
			Mgf1Hash:   hashType,
			SaltLength: saltLength,
		},
		ModulusSizeInBits: modulusSizeInBits,
		PublicExponent:    []byte{0x01, 0x00, 0x01},
	}
	serializedFormat, err := proto.Marshal(keyFormat)
	if err != nil {
		tinkerror.Fail(fmt.Sprintf("failed to marshal key format: %s", err))
	}
	return &tinkpb.KeyTemplate{
		TypeUrl:          rsaSSAPSSSignerTypeURL,
		OutputPrefixType: prefixType,
		Value:            serializedFormat,
	}
}

// RSA_SSA_PKCS1_3072_SHA256_F4_Key_Template is a KeyTemplate that generates a new RSA SSA PKCS1 private key with the following
// parameters:
//   - Modulus size in bits: 3072.
//   - Hash function: SHA256.
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: TINK
func RSA_SSA_PKCS1_3072_SHA256_F4_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PKCS1_Template(tinkpb.OutputPrefixType_TINK, commonpb.HashType_SHA256, 3072)
}

// RSA_SSA_PKCS1_3072_SHA256_F4_RAW_Key_Template is a KeyTemplate that generates a new RSA SSA PKCS1 private key with the following
// parameters:
//   - Modulus size in bits: 3072.
//   - Hash function: SHA256.
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: RAW
func RSA_SSA_PKCS1_3072_SHA256_F4_RAW_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PKCS1_Template(tinkpb.OutputPrefixType_RAW, commonpb.HashType_SHA256, 3072)
}

// RSA_SSA_PKCS1_4096_SHA512_F4_Key_Template is a KeyTemplate that generates a new RSA SSA PKCS1 private key with the following
// parameters:
//   - Modulus size in bits: 4096.
//   - Hash function: SHA512.
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: TINK
func RSA_SSA_PKCS1_4096_SHA512_F4_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PKCS1_Template(tinkpb.OutputPrefixType_TINK, commonpb.HashType_SHA512, 4096)
}

// RSA_SSA_PKCS1_4096_SHA512_F4_RAW_Key_Template is a KeyTemplate that generates a new RSA SSA PKCS1 private key with the following
// parameters:
//   - Modulus size in bits: 4096.
//   - Hash function: SHA512.
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: RAW
func RSA_SSA_PKCS1_4096_SHA512_F4_RAW_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PKCS1_Template(tinkpb.OutputPrefixType_RAW, commonpb.HashType_SHA512, 4096)
}

// RSA_SSA_PSS_3072_SHA256_32_F4_Key_Template is a KeyTemplate that generates a new RSA SSA PSS private key with the following
// parameters:
//   - Modulus size in bits: 3072.
//   - Signature hash: SHA256.
//   - MGF1 hash: SHA256.
//   - Salt length: 32 (i.e., SHA256's output length).
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: TINK
func RSA_SSA_PSS_3072_SHA256_32_F4_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PSS_Template(tinkpb.OutputPrefixType_TINK, commonpb.HashType_SHA256, 32, 3072)
}

// RSA_SSA_PSS_3072_SHA256_32_F4_Raw_Key_Template is a KeyTemplate that generates a new RSA SSA PSS private key with the following
// parameters:
//   - Modulus size in bits: 3072.
//   - Signature hash: SHA256.
//   - MGF1 hash: SHA256.
//   - Salt length: 32 (i.e., SHA256's output length).
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: RAW
func RSA_SSA_PSS_3072_SHA256_32_F4_Raw_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PSS_Template(tinkpb.OutputPrefixType_RAW, commonpb.HashType_SHA256, 32, 3072)
}

// RSA_SSA_PSS_4096_SHA512_64_F4_Key_Template is a KeyTemplate that generates a new RSA SSA PSS private key with the following
// parameters:
//   - Modulus size in bits: 4096.
//   - Signature hash: SHA512.
//   - MGF1 hash: SHA512.
//   - Salt length: 64 (i.e., SHA512's output length).
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: TINK
func RSA_SSA_PSS_4096_SHA512_64_F4_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PSS_Template(tinkpb.OutputPrefixType_TINK, commonpb.HashType_SHA512, 64, 4096)
}

// RSA_SSA_PSS_4096_SHA512_64_F4_Raw_Key_Template is a KeyTemplate that generates a new RSA SSA PSS private key with the following
// parameters:
//   - Modulus size in bits: 4096.
//   - Signature hash: SHA512.
//   - MGF1 hash: SHA512.
//   - Salt length: 64 (i.e., SHA512's output length).
//   - Public Exponent: 65537 (aka F4).
//   - OutputPrefixType: RAW
func RSA_SSA_PSS_4096_SHA512_64_F4_Raw_Key_Template() *tinkpb.KeyTemplate {
	return create_RSA_SSA_PSS_Template(tinkpb.OutputPrefixType_RAW, commonpb.HashType_SHA512, 64, 4096)
}
