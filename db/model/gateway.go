// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model

func (Certificate) TableName() string {
	return "certificate"
}

// Certificate contains TLS information
type Certificate struct {
	Model
	CertificateName string `gorm:"column:certificate_name;size:128"`
	Certificate     string `gorm:"column:certificate;size:128"`
	PrivateKey      string `gorm:"column:private_key;size:128"`
}

func (RuleExtension) TableName() string {
	return "rule_extension"
}

type RuleType string

var HttpRuleType RuleType = "http"

var StreamRuleType RuleType = "stream"

type RuleValueType string

var HttpsRVT RuleValueType = "Https"

var HttpToHttpsRVT RuleValueType = "HttpToHttps"

var HttpAndHttpsRVT RuleValueType = "HttpAndHttps"

type RuleExtension struct {
	Model
	Value RuleValueType `gorm:"column:rule_value_type"`
}

type LoadBalancerType string

var RoundRobinLBType LoadBalancerType = "RoundRobin"

var ConsistenceHashLBType LoadBalancerType = "ConsistentHash"

func (HttpRule) TableName() string {
	return "http_rule"
}

// HttpRule contains http rule
type HttpRule struct {
	Model
	ServiceID        string           `gorm:"column:service_id"`
	ContainerPort    int              `gorm:"column:container_port"`
	Domain           string           `gorm:"column:domain"`
	Path             string           `gorm:"column:path"`
	Header           string           `gorm:"column:header"`
	Cookie           string           `gorm:"column:cookie"`
	IP               string           `gorm:"column:ip"`
	LoadBalancerType LoadBalancerType `gorm:"column:load_balancer_type"`
	CertificateID    string           `gorm:"column:certificate_id"`
	RuleExtensionID  string           `gorm:"column:rule_extension_id"`
}

func (StreamRule) TableName() string {
	return "stream_rule"
}

// StreamRule contain stream rule
type StreamRule struct {
	Model
	ServiceID        string           `gorm:"column:service_id"`
	ContainerPort    int              `gorm:"column:container_port"`
	IP               string           `gorm:"column:ip"`
	Port             string           `gorm:"column:port"`
	LoadBalancerType LoadBalancerType `gorm:"column:load_balancer_type"`
	RuleExtensionID  string           `gorm:"column:rule_extension_id"`
}
