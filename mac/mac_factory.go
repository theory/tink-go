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

package mac

import (
	"fmt"

	"github.com/tink-crypto/tink-go/v2/core/cryptofmt"
	"github.com/tink-crypto/tink-go/v2/internal/internalapi"
	"github.com/tink-crypto/tink-go/v2/internal/internalregistry"
	"github.com/tink-crypto/tink-go/v2/internal/monitoringutil"
	"github.com/tink-crypto/tink-go/v2/internal/primitiveset"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"github.com/tink-crypto/tink-go/v2/monitoring"
	"github.com/tink-crypto/tink-go/v2/tink"
	tinkpb "github.com/tink-crypto/tink-go/v2/proto/tink_go_proto"
)

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	maxInt  = 1<<(intSize-1) - 1
)

// New creates a MAC primitive from the given keyset handle.
func New(handle *keyset.Handle) (tink.MAC, error) {
	ps, err := keyset.Primitives[tink.MAC](handle, internalapi.Token{})
	if err != nil {
		return nil, fmt.Errorf("mac_factory: cannot obtain primitive set: %s", err)
	}
	return newWrappedMAC(ps)
}

// wrappedMAC is a MAC implementation that uses the underlying primitive set to compute and
// verify MACs.
type wrappedMAC struct {
	ps            *primitiveset.PrimitiveSet[tink.MAC]
	computeLogger monitoring.Logger
	verifyLogger  monitoring.Logger
}

var _ (tink.MAC) = (*wrappedMAC)(nil)

func newWrappedMAC(ps *primitiveset.PrimitiveSet[tink.MAC]) (*wrappedMAC, error) {
	computeLogger, verifyLogger, err := createLoggers(ps)
	if err != nil {
		return nil, err
	}
	return &wrappedMAC{
		ps:            ps,
		computeLogger: computeLogger,
		verifyLogger:  verifyLogger,
	}, nil
}

func createLoggers(ps *primitiveset.PrimitiveSet[tink.MAC]) (monitoring.Logger, monitoring.Logger, error) {
	if len(ps.Annotations) == 0 {
		return &monitoringutil.DoNothingLogger{}, &monitoringutil.DoNothingLogger{}, nil
	}
	client := internalregistry.GetMonitoringClient()
	keysetInfo, err := monitoringutil.KeysetInfoFromPrimitiveSet(ps)
	if err != nil {
		return nil, nil, err
	}
	computeLogger, err := client.NewLogger(&monitoring.Context{
		Primitive:   "mac",
		APIFunction: "compute",
		KeysetInfo:  keysetInfo,
	})
	if err != nil {
		return nil, nil, err
	}
	verifyLogger, err := client.NewLogger(&monitoring.Context{
		Primitive:   "mac",
		APIFunction: "verify",
		KeysetInfo:  keysetInfo,
	})
	if err != nil {
		return nil, nil, err
	}
	return computeLogger, verifyLogger, nil
}

// ComputeMAC calculates a MAC over the given data using the primary primitive
// and returns the concatenation of the primary's identifier and the calculated mac.
func (m *wrappedMAC) ComputeMAC(data []byte) ([]byte, error) {
	primary := m.ps.Primary
	if m.ps.Primary.PrefixType == tinkpb.OutputPrefixType_LEGACY {
		d := data
		if len(d) >= maxInt {
			m.computeLogger.LogFailure()
			return nil, fmt.Errorf("mac_factory: data too long")
		}
		data = make([]byte, 0, len(d)+1)
		data = append(data, d...)
		data = append(data, byte(0))
	}
	mac, err := primary.Primitive.ComputeMAC(data)
	if err != nil {
		m.computeLogger.LogFailure()
		return nil, err
	}
	m.computeLogger.Log(primary.KeyID, len(data))
	if len(primary.Prefix) == 0 {
		return mac, nil
	}
	output := make([]byte, 0, len(primary.Prefix)+len(mac))
	output = append(output, primary.Prefix...)
	output = append(output, mac...)
	return output, nil
}

var errInvalidMAC = fmt.Errorf("mac_factory: invalid mac")

// VerifyMAC verifies whether the given mac is a correct authentication code
// for the given data.
func (m *wrappedMAC) VerifyMAC(mac, data []byte) error {
	// This also rejects raw MAC with size of 4 bytes or fewer. Those MACs are
	// clearly insecure, thus should be discouraged.
	prefixSize := cryptofmt.NonRawPrefixSize
	if len(mac) <= prefixSize {
		m.verifyLogger.LogFailure()
		return errInvalidMAC
	}

	// try non raw keys
	prefix := mac[:prefixSize]
	macNoPrefix := mac[prefixSize:]
	entries, err := m.ps.EntriesForPrefix(string(prefix))
	if err == nil {
		for i := 0; i < len(entries); i++ {
			entry := entries[i]
			if entry.PrefixType == tinkpb.OutputPrefixType_LEGACY {
				d := data
				if len(d) >= maxInt {
					m.verifyLogger.LogFailure()
					return fmt.Errorf("mac_factory: data too long")
				}
				data = make([]byte, 0, len(d)+1)
				data = append(data, d...)
				data = append(data, byte(0))
			}
			if err := entry.Primitive.VerifyMAC(macNoPrefix, data); err == nil {
				m.verifyLogger.Log(entry.KeyID, len(data))
				return nil
			}
		}
	}

	// try raw keys
	entries, err = m.ps.RawEntries()
	if err == nil {
		for i := 0; i < len(entries); i++ {
			if err := entries[i].Primitive.VerifyMAC(mac, data); err == nil {
				m.verifyLogger.Log(entries[i].KeyID, len(data))
				return nil
			}
		}
	}

	// nothing worked
	m.verifyLogger.LogFailure()
	return errInvalidMAC
}
