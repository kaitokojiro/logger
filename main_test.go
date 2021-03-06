/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testPassword = "p@55w0rd"
var testToken = "t0k3n"
var testSecret = "s3cr3t"

func TestObfuscate(t *testing.T) {
	patternsToObfuscate = append(patternsToObfuscate, testPassword)
	patternsToObfuscate = append(patternsToObfuscate, testToken)
	patternsToObfuscate = append(patternsToObfuscate, testSecret)

	Convey("Given an input for fsm", t, func() {
		fsmMessage := `{"id":"1","datacenters":{"items":[{"name":"test","username":"test","password":"p@55w0rd"}]}}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(fsmMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for routers creation", t, func() {
		routerMessage := `{"service":"1","components":[{"datacenter_name":"test","datacenter_password":"p@55w0rd"}]}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(routerMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for networks creation", t, func() {
		networkMessage := `{"service":"1","components":[{"datacenter_name":"test","datacenter_password":"p@55w0rd"}]}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(networkMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for instances creation", t, func() {
		instanceMessage := `{"service":"1","components":[{"datacenter_name":"test","datacenter_password":"p@55w0rd"}]}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(instanceMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for firewalls creation", t, func() {
		firewallMessage := `{"service":"1","components":[{"datacenter_name":"test","datacenter_password":"p@55w0rd"}]}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(firewallMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for nats creation", t, func() {
		natMessage := `{"service":"1","components":[{"datacenter_name":"test","datacenter_password":"p@55w0rd"}]}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(natMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "test"), ShouldBeTrue)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for singular generic message", t, func() {
		executionMessage := `{"service":"1", "datacenter_password": "p@55w0rd"}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(executionMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "1"), ShouldBeTrue)
		})
	})

	Convey("Given an input for getting configuration of a service", t, func() {
		configurationMessage := `{"user":"username", "password": "p@55w0rd"}`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(configurationMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, "username"), ShouldBeTrue)
		})
	})

	Convey("Given an input for a list of datacenters", t, func() {
		configurationMessage := `[{"id":4,"group_id":2,"name":"name","type":"aws","region":"eu-west-1","username":"name","password":"","vcloud_url":"","vse_url":"","external_network":"","token":"t0k3n","secret":"s3cr3t","CreatedAt":"2016-11-03T14:32:30.470151Z","UpdatedAt":"2016-11-03T14:32:30.470151Z"},{"id":5,"group_id":2,"name":"name","type":"vcloud","region":"","username":"name@org","password":"p@55w0rd","vcloud_url":"https://myvdc.net","vse_url":"","external_network":"NETWORK","token":"","secret":"","CreatedAt":"2016-11-03T14:32:43.428661Z","UpdatedAt":"2016-11-03T14:32:43.428661Z"}]`
		Convey("the message should be processed correctly and its password removed", func() {
			message := Obfuscate(configurationMessage)
			So(strings.Contains(message, testPassword), ShouldBeFalse)
			So(strings.Contains(message, testToken), ShouldBeFalse)
			So(strings.Contains(message, testSecret), ShouldBeFalse)
		})
	})

}
