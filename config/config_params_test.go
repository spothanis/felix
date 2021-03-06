// Copyright (c) 2016-2017 Tigera, Inc. All rights reserved.

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

package config_test

import (
	. "github.com/projectcalico/felix/config"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"net"
	"reflect"
)

var _ = DescribeTable("Config parsing",
	func(key, value string, expected interface{}) {
		config := New()
		oldVal := reflect.ValueOf(config).Elem().FieldByName(key).Interface()
		config.UpdateFrom(map[string]string{key: value},
			EnvironmentVariable)
		newVal := reflect.ValueOf(config).Elem().FieldByName(key).Interface()
		Expect(oldVal).NotTo(Equal(newVal))
		Expect(newVal).To(Equal(expected))
	},

	Entry("FelixHostname", "FelixHostname", "hostname", "hostname"),
	Entry("FelixHostname FQDN", "FelixHostname", "hostname.foo.bar.com", "hostname.foo.bar.com"),
	Entry("FelixHostname as IP", "FelixHostname", "1.2.3.4", "1.2.3.4"),

	Entry("EtcdAddr IP", "EtcdAddr", "10.0.0.1:1234", "10.0.0.1:1234"),
	Entry("EtcdAddr host", "EtcdAddr", "host:1234", "host:1234"),
	Entry("EtcdScheme", "EtcdScheme", "https", "https"),

	// Etcd key files will be tested for existence, skipping for now.

	Entry("EtcdEndpoints HTTP", "EtcdEndpoints",
		"http://127.0.0.1:1234, http://host:2345",
		[]string{"http://127.0.0.1:1234/", "http://host:2345/"}),
	Entry("EtcdEndpoints HTTPS", "EtcdEndpoints",
		"https://127.0.0.1:1234/, https://host:2345",
		[]string{"https://127.0.0.1:1234/", "https://host:2345/"}),

	Entry("StartupCleanupDelay 12", "StartupCleanupDelay", "12", int(12)),
	Entry("StartupCleanupDelay 0", "StartupCleanupDelay", "0", int(0)),
	Entry("PeriodicResyncInterval 1500", "PeriodicResyncInterval", "1500", int(1500)),
	Entry("PeriodicResyncInterval 0", "PeriodicResyncInterval", "0", int(0)),
	Entry("HostInterfacePollInterval", "HostInterfacePollInterval", "11", int(11)),
	Entry("HostInterfacePollInterval", "HostInterfacePollInterval", "0", int(0)),

	Entry("InterfacePrefix", "InterfacePrefix", "tap", "tap"),
	Entry("InterfacePrefix list", "InterfacePrefix", "tap,cali", "tap,cali"),

	Entry("ChainInsertMode append", "ChainInsertMode", "append", "append"),

	Entry("DefaultEndpointToHostAction", "DefaultEndpointToHostAction",
		"RETURN", "RETURN"),
	Entry("DefaultEndpointToHostAction", "DefaultEndpointToHostAction",
		"ACCEPT", "ACCEPT"),

	Entry("DropActionOverride", "DropActionOverride",
		"ACCEPT", "ACCEPT"),
	Entry("DropActionOverride norm", "DropActionOverride",
		"accept", "ACCEPT"),
	Entry("DropActionOverride LOG-and-ACCEPT", "DropActionOverride",
		"LOG-and-ACCEPT", "LOG-and-ACCEPT"),
	Entry("DropActionOverride log-and-accept", "DropActionOverride",
		"log-and-accept", "LOG-and-ACCEPT"),
	Entry("DropActionOverride log-and-drop", "DropActionOverride",
		"log-and-drop", "LOG-and-DROP"),

	Entry("LogFilePath", "LogFilePath", "/tmp/felix.log", "/tmp/felix.log"),

	Entry("LogSeverityFile", "LogSeverityFile", "debug", "DEBUG"),
	Entry("LogSeverityFile", "LogSeverityFile", "warning", "WARNING"),
	Entry("LogSeverityFile", "LogSeverityFile", "error", "ERROR"),
	Entry("LogSeverityFile", "LogSeverityFile", "critical", "CRITICAL"),

	Entry("LogSeverityScreen", "LogSeverityScreen", "debug", "DEBUG"),
	Entry("LogSeverityScreen", "LogSeverityScreen", "warning", "WARNING"),
	Entry("LogSeverityScreen", "LogSeverityScreen", "error", "ERROR"),
	Entry("LogSeverityScreen", "LogSeverityScreen", "critical", "CRITICAL"),

	Entry("LogSeveritySys", "LogSeveritySys", "debug", "DEBUG"),
	Entry("LogSeveritySys", "LogSeveritySys", "warning", "WARNING"),
	Entry("LogSeveritySys", "LogSeveritySys", "error", "ERROR"),
	Entry("LogSeveritySys", "LogSeveritySys", "critical", "CRITICAL"),

	Entry("IpInIpEnabled", "IpInIpEnabled", "true", true),
	Entry("IpInIpEnabled", "IpInIpEnabled", "y", true),
	Entry("IpInIpEnabled", "IpInIpEnabled", "True", true),

	Entry("IpInIpMtu", "IpInIpMtu", "1234", int(1234)),
	Entry("IpInIpTunnelAddr", "IpInIpTunnelAddr",
		"10.0.0.1", net.ParseIP("10.0.0.1")),

	Entry("ReportingIntervalSecs", "ReportingIntervalSecs", "31", int(31)),
	Entry("ReportingTTLSecs", "ReportingTTLSecs", "91", int(91)),

	Entry("EndpointReportingEnabled", "EndpointReportingEnabled",
		"true", true),
	Entry("EndpointReportingEnabled", "EndpointReportingEnabled",
		"yes", true),
	Entry("EndpointReportingDelaySecs", "EndpointReportingDelaySecs",
		"10", float64(10)),

	Entry("MaxIpsetSize", "MaxIpsetSize", "12345", int(12345)),
	Entry("IptablesMarkMask", "IptablesMarkMask", "0xf0f0", uint32(0xf0f0)),

	Entry("PrometheusMetricsEnabled", "PrometheusMetricsEnabled", "true", true),
	Entry("PrometheusMetricsPort", "PrometheusMetricsPort", "1234", int(1234)),

	Entry("FailsafeInboundHostPorts", "FailsafeInboundHostPorts", "1,2,3,4", []uint16{1, 2, 3, 4}),
	Entry("FailsafeOutboundHostPorts", "FailsafeOutboundHostPorts", "1,2,3,4", []uint16{1, 2, 3, 4}),
	Entry("FailsafeInboundHostPorts empty", "FailsafeInboundHostPorts", "none", []uint16(nil)),
	Entry("FailsafeOutboundHostPorts empty", "FailsafeOutboundHostPorts", "none", []uint16(nil)),
	Entry("FailsafeInboundHostPorts empty", "FailsafeInboundHostPorts", "", []uint16(nil)),
	Entry("FailsafeOutboundHostPorts empty", "FailsafeOutboundHostPorts", "", []uint16(nil)),
)

var _ = DescribeTable("Mark bit calculation tests",
	func(mask string, bitNum int, expected uint32) {
		config := New()
		config.UpdateFrom(map[string]string{"IptablesMarkMask": mask}, EnvironmentVariable)
		Expect(config.NthIPTablesMark(bitNum)).To(Equal(expected))
	},
	Entry("0th bit in 0xf", "0xf", 0, uint32(0x1)),
	Entry("1st bit in 0xf", "0xf", 1, uint32(0x2)),
	Entry("7th bit in 0xff", "0xff", 7, uint32(0x80)),
	Entry("4th bit in 0xf00f", "0xf00f", 4, uint32(0x1000)),
	Entry("3rd bit in 0xf00f", "0xf00f", 3, uint32(0x0008)),
	Entry("7th bit in 0xf00f", "0xf00f", 7, uint32(0x8000)),
	Entry("0th bit of 0xff000000", "0xff000000", 0, uint32(0x01000000)),
)

var _ = DescribeTable("Next mark bit calculation tests",
	func(mask string, numCalls int, expected uint32) {
		config := New()
		config.UpdateFrom(map[string]string{"IptablesMarkMask": mask}, EnvironmentVariable)
		var mark uint32
		for i := 0; i < numCalls; i++ {
			mark = config.NextIptablesMark()
		}
		Expect(mark).To(Equal(expected))
	},
	Entry("0th bit in 0xf", "0xf", 1, uint32(0x1)),
	Entry("1st bit in 0xf", "0xf", 2, uint32(0x2)),
	Entry("7th bit in 0xff", "0xff", 8, uint32(0x80)),
	Entry("7th bit in 0xf00f", "0xf00f", 8, uint32(0x8000)),
	Entry("0th bit of 0xff000000", "0xff000000", 1, uint32(0x01000000)),
)
