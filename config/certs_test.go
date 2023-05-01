// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetGetDeleteCerts(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	cert1 := &configtypes.Cert{
		Host:           "test1",
		CACertData:     "<REDACTED>",
		SkipCertVerify: "false",
	}

	cert2 := &configtypes.Cert{
		Host:           "test2",
		CACertData:     "<REDACTED>",
		SkipCertVerify: "true",
		Insecure:       "false",
	}

	ctx, err := GetCert("test1")
	assert.Equal(t, "cert configuration for test1 not found", err.Error())
	assert.Nil(t, ctx)

	err = SetCert(cert1)
	assert.NoError(t, err)

	ctx, err = GetCert("test1")
	assert.Nil(t, err)
	assert.Equal(t, cert1, ctx)

	err = SetCert(cert2)
	assert.NoError(t, err)

	ctx, err = GetCert("test2")
	assert.Nil(t, err)
	assert.Equal(t, cert2, ctx)

	_, err = GetCert("")
	assert.Equal(t, "host is empty", err.Error())

	err = DeleteCert("test")
	assert.Equal(t, "cert configuration for test not found", err.Error())

	err = DeleteCert("")
	assert.Equal(t, "host is empty", err.Error())

	err = DeleteCert("test1")
	assert.Nil(t, err)

	ctx, err = GetCert("test1")
	assert.Nil(t, ctx)
	assert.Equal(t, "cert configuration for test1 not found", err.Error())
}

func TestSetCert(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name   string
		cert   *configtypes.Cert
		errStr string
	}{
		{
			name: "should add new cert to empty client config",
			cert: &configtypes.Cert{
				Host:           "test.vmware.com",
				CACertData:     "testCAData",
				SkipCertVerify: "true",
				Insecure:       "true",
			},
		},
		{
			name: "should update existing cert",
			cert: &configtypes.Cert{
				Host:           "test.vmware.com",
				CACertData:     "testCADataUpdated",
				Insecure:       "false",
				SkipCertVerify: "false",
			},
		},
		{
			name: "should update existing cert with SkipCertVerify and Insecure field",
			cert: &configtypes.Cert{
				Host:           "test.vmware.com",
				CACertData:     "testCADataUpdated",
				SkipCertVerify: "true",
				Insecure:       "true",
			},
		},
		{
			name: "should add the new cert to the existing certs",
			cert: &configtypes.Cert{
				Host:           "test.vmware.com:443",
				CACertData:     "testCAData2",
				SkipCertVerify: "true",
				Insecure:       "false",
			},
		},
		{
			name: "should return error when the host is empty",
			cert: &configtypes.Cert{
				Host:           "",
				CACertData:     "testCAData2",
				SkipCertVerify: "true",
				Insecure:       "false",
			},
			errStr: "host is empty",
		},
		{
			name: "should not return error when the cert is nil",
			cert: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetCert(tc.cert)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			// if cert is not nil, validate with GetCert
			if tc.cert != nil {
				ok, err := CertExists(tc.cert.Host)
				if tc.errStr == "" {
					assert.True(t, ok)
					assert.NoError(t, err)
				} else {
					assert.EqualError(t, err, tc.errStr)
				}

				gotCert, err := GetCert(tc.cert.Host)
				if tc.errStr == "" {
					assert.NoError(t, err)
					assert.Equal(t, tc.cert, gotCert)
				} else {
					assert.EqualError(t, err, tc.errStr)
				}
			}
		})
	}
}

func TestGetCerts(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name      string
		cert      *configtypes.Cert
		errStr    string
		wantCerts []*configtypes.Cert
	}{
		{
			name:      "should return empty list when no certs were added",
			cert:      nil,
			wantCerts: []*configtypes.Cert{},
		},
		{
			name: "should return the cert added",
			cert: &configtypes.Cert{
				Host:       "test.vmware.com",
				CACertData: "testCAData",
				Insecure:   "false",
			},
			wantCerts: []*configtypes.Cert{
				{
					Host:       "test.vmware.com",
					CACertData: "testCAData",
					Insecure:   "false",
				},
			},
		},
		{
			name: "should return the cert updated",
			cert: &configtypes.Cert{
				Host:       "test.vmware.com",
				CACertData: "testCADataUpdated",
				Insecure:   "true",
			},
			wantCerts: []*configtypes.Cert{
				{
					Host:       "test.vmware.com",
					CACertData: "testCADataUpdated",
					Insecure:   "true",
				},
			},
		},
		{
			name: "should return both the existing and the new cert added",
			cert: &configtypes.Cert{
				Host:           "test.vmware.com:443",
				CACertData:     "testCAData2",
				SkipCertVerify: "true",
				Insecure:       "false",
			},
			wantCerts: []*configtypes.Cert{
				{
					Host:       "test.vmware.com",
					CACertData: "testCADataUpdated",
					Insecure:   "true",
				},
				{
					Host:           "test.vmware.com:443",
					CACertData:     "testCAData2",
					SkipCertVerify: "true",
					Insecure:       "false",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetCert(tc.cert)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			gotCerts, err := GetCerts()
			if tc.errStr == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantCerts, gotCerts)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
		})
	}
}
