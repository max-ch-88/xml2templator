package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

func (c *ZabbixExport) Parse(data *[]byte) error {

	err := xml.Unmarshal(*data, c)
	if err != nil {
		return err
	}

	return nil
}

func getData(filename *string, data *[]byte) error {

	xmlFile, err := os.Open(*filename)
	if err != nil {
		return err
	}

	defer xmlFile.Close()

	*data, _ = ioutil.ReadAll(xmlFile)

	return nil
}

func (c *Templator) Export(xmlIn *ZabbixExport, data *[]byte) error {
	var err error
	var itemType = []string{
		"ZABBIX_PASSIVE",
		"SNMPV1",
		"TRAP",
		"SIMPLE",
		"SNMPV2",
		"SNMP",
		"INTERNAL",
		"SNMPV3",
		"ZABBIX_ACTIVE",
		"AGGREGATE",
		"EXTERNAL",
		"ODBC",
		"IPMI",
		"SSH",
		"TELNET",
		"CALCULATED",
		"JMX",
		"SNMP_TRAP",
		"DEPENDENT",
		"HTTP_AGENT",
	}
	var valueType = []string{
		"FLOAT",
		"CHAR",
		"LOG",
		"UNSIGNED",
		"TEXT",
	}
	var authTypeHTTP = []string{
		"",
		"BASIC",
		"NTLM",
		"KERBEROS",
	}
	var authTypeSSH = []string{
		"PASSWORD",
		"PUBLIC_KEY",
	}
	var preprocessingType = []string{
		"",
		"MULTIPLIER",
		"RTRIM",
		"LTRIM",
		"TRIM",
		"REGEX",
		"BOOL_TO_DECIMAL",
		"OCTAL_TO_DECIMAL",
		"HEX_TO_DECIMAL",
		"SIMPLE_CHANGE",
		"CHANGE_PER_SECOND",
		"XMLPATH",
		"JSONPATH",
		"IN_RANGE",
		"MATCHES_REGEX",
		"NOT_MATCHES_REGEX",
		"CHECK_JSON_ERROR",
		"CHECK_XML_ERROR",
		"CHECK_REGEX_ERROR",
		"DISCARD_UNCHANGED",
		"DISCARD_UNCHANGED_HEARTBEAT",
		"JAVASCRIPT",
		"PROMETHEUS_PATTERN",
		"PROMETHEUS_TO_JSON",
		"CSV_TO_JSON",
	}
	var errorHandler = []string{
		"ORIGINAL_ERROR",
		"DISCARD_VALUE",
		"CUSTOM_VALUE",
		"CUSTOM_ERROR",
	}
	var evalType = []string{
		"AND_OR",
		"AND",
		"OR",
		"FORMULA",
	}
	var operator = map[string]string{
		"8": "MATCHES_REGEX",
		"9": "NOT_MATCHES_REGEX",
	}

	var priority = []string{
		"NOT_CLASSIFIED",
		"INFO",
		"WARNING",
		"AVERAGE",
		"HIGH",
		"DISASTER",
	}

	valueMapsSize := len(xmlIn.ValueMaps.ValueMap)
	c.ValueMaps = make([]ValueMap, valueMapsSize)
	for i := 0; i < valueMapsSize; i++ {
		c.ValueMaps[i].Name = xmlIn.ValueMaps.ValueMap[i].Name
		mappingsSize := len(xmlIn.ValueMaps.ValueMap[i].Mappings.Mapping)
		c.ValueMaps[i].Mappings = make([]Mapping, mappingsSize)
		for j := 0; j < mappingsSize; j++ {
			c.ValueMaps[i].Mappings[j].Value = xmlIn.ValueMaps.ValueMap[i].Mappings.Mapping[j].Value
			c.ValueMaps[i].Mappings[j].Newvalue = xmlIn.ValueMaps.ValueMap[i].Mappings.Mapping[j].Newvalue
		}
	}

	templatesSize := len(xmlIn.Templates.Template)
	c.Templates = make([]Template, templatesSize)
	for i := 0; i < templatesSize; i++ {
		c.Templates[i].Name = xmlIn.Templates.Template[i].Name
		c.Templates[i].Version = xmlIn.Version
		macrosSize := len(xmlIn.Templates.Template[i].Macros.Macro)
		c.Templates[i].Macros = make([]Macro, macrosSize)
		for j := 0; j < macrosSize; j++ {
			c.Templates[i].Macros[j].Macro = xmlIn.Templates.Template[i].Macros.Macro[j].Macro
			c.Templates[i].Macros[j].Value = xmlIn.Templates.Template[i].Macros.Macro[j].Value
		}

		itemsSize := len(xmlIn.Templates.Template[i].Items.Item)
		c.Templates[i].Items = make([]TemplatorItem, itemsSize)
		for j := 0; j < itemsSize; j++ {
			c.Templates[i].Items[j].Name = xmlIn.Templates.Template[i].Items.Item[j].Name
			n, _ := strconv.Atoi(xmlIn.Templates.Template[i].Items.Item[j].Type)
			c.Templates[i].Items[j].Type = itemType[n]
			c.Templates[i].Items[j].SnmpCommunity = xmlIn.Templates.Template[i].Items.Item[j].SnmpCommunity
			c.Templates[i].Items[j].SnmpOid = xmlIn.Templates.Template[i].Items.Item[j].SnmpOid
			c.Templates[i].Items[j].Key = xmlIn.Templates.Template[i].Items.Item[j].Key
			c.Templates[i].Items[j].Delay = xmlIn.Templates.Template[i].Items.Item[j].Delay
			if xmlIn.Templates.Template[i].Items.Item[j].History != "90d" {
				c.Templates[i].Items[j].History = xmlIn.Templates.Template[i].Items.Item[j].History
			}
			if xmlIn.Templates.Template[i].Items.Item[j].Trends != "365d" {
				c.Templates[i].Items[j].Trends = xmlIn.Templates.Template[i].Items.Item[j].Trends
			}
			if xmlIn.Templates.Template[i].Items.Item[j].Status == "1" {
				c.Templates[i].Items[j].Status = xmlIn.Templates.Template[i].Items.Item[j].Status
			}
			n, _ = strconv.Atoi(xmlIn.Templates.Template[i].Items.Item[j].ValueType)
			c.Templates[i].Items[j].ValueType = valueType[n]
			if xmlIn.Templates.Template[i].Items.Item[j].AllowedHosts == "1" {
				c.Templates[i].Items[j].AllowedHosts = xmlIn.Templates.Template[i].Items.Item[j].AllowedHosts
			}
			c.Templates[i].Items[j].Units = xmlIn.Templates.Template[i].Items.Item[j].Units
			c.Templates[i].Items[j].Snmpv3Contextname = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Contextname
			c.Templates[i].Items[j].Snmpv3Securityname = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Securityname
			if xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Securitylevel != "0" {
				c.Templates[i].Items[j].Snmpv3Securitylevel = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Securitylevel
			}
			if xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Authprotocol != "0" {
				c.Templates[i].Items[j].Snmpv3Authprotocol = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Authprotocol
			}
			c.Templates[i].Items[j].Snmpv3Authpassphrase = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Authpassphrase
			if xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Privprotocol != "0" {
				c.Templates[i].Items[j].Snmpv3Privprotocol = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Privprotocol
			}
			c.Templates[i].Items[j].Snmpv3Privpassphrase = xmlIn.Templates.Template[i].Items.Item[j].Snmpv3Privpassphrase
			c.Templates[i].Items[j].Params = xmlIn.Templates.Template[i].Items.Item[j].Params
			c.Templates[i].Items[j].IpmiSensor = xmlIn.Templates.Template[i].Items.Item[j].IpmiSensor
			n, _ = strconv.Atoi(xmlIn.Templates.Template[i].Items.Item[j].Authtype)
			if c.Templates[i].Items[j].Type == "HTTP_AGENT" {
				c.Templates[i].Items[j].Authtype = authTypeHTTP[n]
			} else if c.Templates[i].Items[j].Type == "SSH" {
				c.Templates[i].Items[j].Authtype = authTypeSSH[n]
			}
			c.Templates[i].Items[j].Username = xmlIn.Templates.Template[i].Items.Item[j].Username
			c.Templates[i].Items[j].Password = xmlIn.Templates.Template[i].Items.Item[j].Password
			c.Templates[i].Items[j].Publickey = xmlIn.Templates.Template[i].Items.Item[j].Publickey
			c.Templates[i].Items[j].Privatekey = xmlIn.Templates.Template[i].Items.Item[j].Privatekey
			c.Templates[i].Items[j].Port = xmlIn.Templates.Template[i].Items.Item[j].Port
			c.Templates[i].Items[j].Description = xmlIn.Templates.Template[i].Items.Item[j].Description
			if xmlIn.Templates.Template[i].Items.Item[j].InventoryLink != "0" {
				c.Templates[i].Items[j].InventoryLink = xmlIn.Templates.Template[i].Items.Item[j].InventoryLink
			}
			for k := 0; k < len(xmlIn.Templates.Template[i].Items.Item[j].Applications.Application); k++ {
				c.Templates[i].Items[j].Applications = xmlIn.Templates.Template[i].Items.Item[j].Applications.Application[k].Name
			}
			c.Templates[i].Items[j].Valuemap = xmlIn.Templates.Template[i].Items.Item[j].Valuemap.Name
			c.Templates[i].Items[j].Logtimefmt = xmlIn.Templates.Template[i].Items.Item[j].Logtimefmt
			preprocessingSize := len(xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step)
			c.Templates[i].Items[j].Preprocessing = make([]Step, preprocessingSize)
			for k := 0; k < preprocessingSize; k++ {
				n, _ := strconv.Atoi(xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step[k].Type)
				c.Templates[i].Items[j].Preprocessing[k].Type = preprocessingType[n]
				c.Templates[i].Items[j].Preprocessing[k].Params = xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step[k].Params
				if xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step[k].ErrorHandler != "" {
					n, _ = strconv.Atoi(xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step[k].ErrorHandler)
					c.Templates[i].Items[j].Preprocessing[k].ErrorHandler = errorHandler[n]
				}
				c.Templates[i].Items[j].Preprocessing[k].ErrorHandlerParams = xmlIn.Templates.Template[i].Items.Item[j].Preprocessing.Step[k].ErrorHandlerParams
			}
			c.Templates[i].Items[j].JmxEndpoint = xmlIn.Templates.Template[i].Items.Item[j].JmxEndpoint
			if xmlIn.Templates.Template[i].Items.Item[j].Timeout != "3s" {
				c.Templates[i].Items[j].Timeout = xmlIn.Templates.Template[i].Items.Item[j].Timeout
			}
			c.Templates[i].Items[j].URL = xmlIn.Templates.Template[i].Items.Item[j].URL
			c.Templates[i].Items[j].QueryFields = xmlIn.Templates.Template[i].Items.Item[j].QueryFields
			c.Templates[i].Items[j].Posts = xmlIn.Templates.Template[i].Items.Item[j].Posts
			if xmlIn.Templates.Template[i].Items.Item[j].StatusCodes != "200" {
				c.Templates[i].Items[j].StatusCodes = xmlIn.Templates.Template[i].Items.Item[j].StatusCodes
			}
			if xmlIn.Templates.Template[i].Items.Item[j].FollowRedirects != "1" {
				c.Templates[i].Items[j].FollowRedirects = xmlIn.Templates.Template[i].Items.Item[j].FollowRedirects
			}
			if xmlIn.Templates.Template[i].Items.Item[j].PostType != "0" {
				c.Templates[i].Items[j].PostType = xmlIn.Templates.Template[i].Items.Item[j].PostType
			}
			c.Templates[i].Items[j].HTTPProxy = xmlIn.Templates.Template[i].Items.Item[j].HTTPProxy
			c.Templates[i].Items[j].Headers = xmlIn.Templates.Template[i].Items.Item[j].Headers
			if xmlIn.Templates.Template[i].Items.Item[j].RetrieveMode != "0" {
				c.Templates[i].Items[j].RetrieveMode = xmlIn.Templates.Template[i].Items.Item[j].RetrieveMode
			}
			if xmlIn.Templates.Template[i].Items.Item[j].RequestMethod != "0" {
				c.Templates[i].Items[j].RequestMethod = xmlIn.Templates.Template[i].Items.Item[j].RequestMethod
			}
			if xmlIn.Templates.Template[i].Items.Item[j].OutputFormat != "0" {
				c.Templates[i].Items[j].OutputFormat = xmlIn.Templates.Template[i].Items.Item[j].OutputFormat
			}
			if xmlIn.Templates.Template[i].Items.Item[j].AllowTraps != "0" {
				c.Templates[i].Items[j].AllowTraps = xmlIn.Templates.Template[i].Items.Item[j].AllowTraps
			}
			c.Templates[i].Items[j].SslCertFile = xmlIn.Templates.Template[i].Items.Item[j].SslCertFile
			c.Templates[i].Items[j].SslKeyFile = xmlIn.Templates.Template[i].Items.Item[j].SslKeyFile
			c.Templates[i].Items[j].SslKeyPassword = xmlIn.Templates.Template[i].Items.Item[j].SslKeyPassword
			if xmlIn.Templates.Template[i].Items.Item[j].VerifyPeer != "0" {
				c.Templates[i].Items[j].VerifyPeer = xmlIn.Templates.Template[i].Items.Item[j].VerifyPeer
			}
			if xmlIn.Templates.Template[i].Items.Item[j].VerifyHost != "0" {
				c.Templates[i].Items[j].VerifyHost = xmlIn.Templates.Template[i].Items.Item[j].VerifyHost
			}
			c.Templates[i].Items[j].MasterItem = xmlIn.Templates.Template[i].Items.Item[j].MasterItem
			for k := 0; k < len(xmlIn.Templates.Template[i].Items.Item[j].ApplicationPrototypes.ApplicationPrototype); k++ {
				c.Templates[i].Items[j].Applications = xmlIn.Templates.Template[i].Items.Item[j].ApplicationPrototypes.ApplicationPrototype[k].Name
			}
		}

		discoveryRulesSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule)
		c.Templates[i].DiscoveryRules = make([]DiscoveryRule, discoveryRulesSize)
		for j := 0; j < discoveryRulesSize; j++ {
			c.Templates[i].DiscoveryRules[j].Name = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Name
			n, _ := strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Type)
			c.Templates[i].DiscoveryRules[j].Type = itemType[n]
			c.Templates[i].DiscoveryRules[j].SnmpCommunity = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].SnmpCommunity
			c.Templates[i].DiscoveryRules[j].SnmpOid = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].SnmpOid
			c.Templates[i].DiscoveryRules[j].Key = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Key
			c.Templates[i].DiscoveryRules[j].Delay = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Delay
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Status == "1" {
				c.Templates[i].DiscoveryRules[j].Status = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Status
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].AllowedHosts == "1" {
				c.Templates[i].DiscoveryRules[j].AllowedHosts = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].AllowedHosts
			}
			c.Templates[i].DiscoveryRules[j].Snmpv3Contextname = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Contextname
			c.Templates[i].DiscoveryRules[j].Snmpv3Securityname = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Securityname
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Securitylevel == "1" {
				c.Templates[i].DiscoveryRules[j].Snmpv3Securitylevel = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Securitylevel
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Authprotocol != "0" {
				c.Templates[i].DiscoveryRules[j].Snmpv3Authprotocol = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Authprotocol
			}
			c.Templates[i].DiscoveryRules[j].Snmpv3Authpassphrase = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Authpassphrase
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Privprotocol != "0" {
				c.Templates[i].DiscoveryRules[j].Snmpv3Privprotocol = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Privprotocol
			}
			c.Templates[i].DiscoveryRules[j].Snmpv3Privpassphrase = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Snmpv3Privpassphrase
			c.Templates[i].DiscoveryRules[j].Params = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Params
			c.Templates[i].DiscoveryRules[j].IpmiSensor = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].IpmiSensor
			n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Authtype)
			if c.Templates[i].DiscoveryRules[j].Type == "HTTP_AGENT" {
				c.Templates[i].DiscoveryRules[j].Authtype = authTypeHTTP[n]
			} else if c.Templates[i].DiscoveryRules[j].Type == "SSH" {
				c.Templates[i].DiscoveryRules[j].Authtype = authTypeSSH[n]
			}
			c.Templates[i].DiscoveryRules[j].Username = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Username
			c.Templates[i].DiscoveryRules[j].Password = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Password
			c.Templates[i].DiscoveryRules[j].Publickey = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Publickey
			c.Templates[i].DiscoveryRules[j].Privatekey = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Privatekey
			c.Templates[i].DiscoveryRules[j].Port = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Port
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Lifetime != "30d" {
				c.Templates[i].DiscoveryRules[j].Lifetime = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Lifetime
			}
			c.Templates[i].DiscoveryRules[j].Description = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Description
			preprocessingSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step)
			c.Templates[i].DiscoveryRules[j].Preprocessing = make([]Step, preprocessingSize)
			for k := 0; k < preprocessingSize; k++ {
				n, _ := strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step[k].Type)
				c.Templates[i].DiscoveryRules[j].Preprocessing[k].Type = preprocessingType[n]
				c.Templates[i].DiscoveryRules[j].Preprocessing[k].Params = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step[k].Params
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step[k].ErrorHandler != "" {
					n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step[k].ErrorHandler)
					c.Templates[i].DiscoveryRules[j].Preprocessing[k].ErrorHandler = errorHandler[n]
				}
				c.Templates[i].DiscoveryRules[j].Preprocessing[k].ErrorHandlerParams = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Preprocessing.Step[k].ErrorHandlerParams
			}
			c.Templates[i].DiscoveryRules[j].JmxEndpoint = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].JmxEndpoint
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Timeout != "3s" {
				c.Templates[i].DiscoveryRules[j].Timeout = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Timeout
			}
			c.Templates[i].DiscoveryRules[j].URL = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].URL
			c.Templates[i].DiscoveryRules[j].QueryFields = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].QueryFields
			c.Templates[i].DiscoveryRules[j].Posts = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Posts
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].StatusCodes != "200" {
				c.Templates[i].DiscoveryRules[j].StatusCodes = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].StatusCodes
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].FollowRedirects != "1" {
				c.Templates[i].DiscoveryRules[j].FollowRedirects = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].FollowRedirects
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].PostType != "0" {
				c.Templates[i].DiscoveryRules[j].PostType = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].PostType
			}
			c.Templates[i].DiscoveryRules[j].HTTPProxy = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].HTTPProxy
			c.Templates[i].DiscoveryRules[j].Headers = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Headers
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].RetrieveMode != "0" {
				c.Templates[i].DiscoveryRules[j].RetrieveMode = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].RetrieveMode
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].RequestMethod != "0" {
				c.Templates[i].DiscoveryRules[j].RequestMethod = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].RequestMethod
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].AllowTraps != "0" {
				c.Templates[i].DiscoveryRules[j].AllowTraps = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].AllowTraps
			}
			c.Templates[i].DiscoveryRules[j].SslCertFile = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].SslCertFile
			c.Templates[i].DiscoveryRules[j].SslKeyFile = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].SslKeyFile
			c.Templates[i].DiscoveryRules[j].SslKeyPassword = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].SslKeyPassword
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].VerifyPeer != "0" {
				c.Templates[i].DiscoveryRules[j].VerifyPeer = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].VerifyPeer
			}
			if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].VerifyHost != "0" {
				c.Templates[i].DiscoveryRules[j].VerifyHost = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].VerifyHost
			}
			c.Templates[i].DiscoveryRules[j].MasterItem = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].MasterItem

			itemsSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype)
			c.Templates[i].DiscoveryRules[j].Items = make([]TemplatorItem, itemsSize)
			for l := 0; l < itemsSize; l++ {
				c.Templates[i].DiscoveryRules[j].Items[l].Name = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Name
				n, _ := strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Type)
				c.Templates[i].DiscoveryRules[j].Items[l].Type = itemType[n]
				c.Templates[i].DiscoveryRules[j].Items[l].SnmpCommunity = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].SnmpCommunity
				c.Templates[i].DiscoveryRules[j].Items[l].SnmpOid = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].SnmpOid
				c.Templates[i].DiscoveryRules[j].Items[l].Key = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Key
				c.Templates[i].DiscoveryRules[j].Items[l].Delay = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Delay
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].History != "90d" {
					c.Templates[i].DiscoveryRules[j].Items[l].History = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].History
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Trends != "365d" {
					c.Templates[i].DiscoveryRules[j].Items[l].Trends = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Trends
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Status == "1" {
					c.Templates[i].DiscoveryRules[j].Items[l].Status = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Status
				}
				n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].ValueType)
				c.Templates[i].DiscoveryRules[j].Items[l].ValueType = valueType[n]
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].AllowedHosts == "1" {
					c.Templates[i].DiscoveryRules[j].Items[l].AllowedHosts = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].AllowedHosts
				}
				c.Templates[i].DiscoveryRules[j].Items[l].Units = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Units
				c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Contextname = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Contextname
				c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Securityname = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Securityname
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Securitylevel != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Securitylevel = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Securitylevel
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Authprotocol != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Authprotocol = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Authprotocol
				}
				c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Authpassphrase = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Authpassphrase
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Privprotocol != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Privprotocol = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Privprotocol
				}
				c.Templates[i].DiscoveryRules[j].Items[l].Snmpv3Privpassphrase = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Snmpv3Privpassphrase
				c.Templates[i].DiscoveryRules[j].Items[l].Params = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Params
				c.Templates[i].DiscoveryRules[j].Items[l].IpmiSensor = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].IpmiSensor
				n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Authtype)
				if c.Templates[i].DiscoveryRules[j].Items[l].Type == "HTTP_AGENT" {
					c.Templates[i].DiscoveryRules[j].Items[l].Authtype = authTypeHTTP[n]
				} else if c.Templates[i].DiscoveryRules[j].Items[l].Type == "SSH" {
					c.Templates[i].DiscoveryRules[j].Items[l].Authtype = authTypeSSH[n]
				}
				c.Templates[i].DiscoveryRules[j].Items[l].Username = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Username
				c.Templates[i].DiscoveryRules[j].Items[l].Password = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Password
				c.Templates[i].DiscoveryRules[j].Items[l].Publickey = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Publickey
				c.Templates[i].DiscoveryRules[j].Items[l].Privatekey = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Privatekey
				c.Templates[i].DiscoveryRules[j].Items[l].Port = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Port
				c.Templates[i].DiscoveryRules[j].Items[l].Description = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Description
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].InventoryLink != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].InventoryLink = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].InventoryLink
				}
				for k := 0; k < len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Applications.Application); k++ {
					c.Templates[i].DiscoveryRules[j].Items[l].Applications = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Applications.Application[k].Name
				}
				c.Templates[i].DiscoveryRules[j].Items[l].Valuemap = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Valuemap.Name
				c.Templates[i].DiscoveryRules[j].Items[l].Logtimefmt = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Logtimefmt
				preprocessingSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step)
				c.Templates[i].DiscoveryRules[j].Items[l].Preprocessing = make([]Step, preprocessingSize)
				for k := 0; k < preprocessingSize; k++ {
					n, _ := strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step[k].Type)
					c.Templates[i].DiscoveryRules[j].Items[l].Preprocessing[k].Type = preprocessingType[n]
					c.Templates[i].DiscoveryRules[j].Items[l].Preprocessing[k].Params = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step[k].Params
					if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step[k].ErrorHandler != "" {
						n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step[k].ErrorHandler)
						c.Templates[i].DiscoveryRules[j].Items[l].Preprocessing[k].ErrorHandler = errorHandler[n]
					}
					c.Templates[i].DiscoveryRules[j].Items[l].Preprocessing[k].ErrorHandlerParams = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Preprocessing.Step[k].ErrorHandlerParams
				}
				c.Templates[i].DiscoveryRules[j].Items[l].JmxEndpoint = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].JmxEndpoint
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Timeout != "3s" {
					c.Templates[i].DiscoveryRules[j].Items[l].Timeout = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Timeout
				}
				c.Templates[i].DiscoveryRules[j].Items[l].URL = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].URL
				c.Templates[i].DiscoveryRules[j].Items[l].QueryFields = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].QueryFields
				c.Templates[i].DiscoveryRules[j].Items[l].Posts = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Posts
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].StatusCodes != "200" {
					c.Templates[i].DiscoveryRules[j].Items[l].StatusCodes = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].StatusCodes
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].FollowRedirects != "1" {
					c.Templates[i].DiscoveryRules[j].Items[l].FollowRedirects = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].FollowRedirects
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].PostType != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].PostType = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].PostType
				}
				c.Templates[i].DiscoveryRules[j].Items[l].HTTPProxy = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].HTTPProxy
				c.Templates[i].DiscoveryRules[j].Items[l].Headers = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].Headers
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].RetrieveMode != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].RetrieveMode = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].RetrieveMode
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].RequestMethod != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].RequestMethod = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].RequestMethod
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].OutputFormat != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].OutputFormat = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].OutputFormat
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].AllowTraps != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].AllowTraps = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].AllowTraps
				}
				c.Templates[i].DiscoveryRules[j].Items[l].SslCertFile = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].SslCertFile
				c.Templates[i].DiscoveryRules[j].Items[l].SslKeyFile = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].SslKeyFile
				c.Templates[i].DiscoveryRules[j].Items[l].SslKeyPassword = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].SslKeyPassword
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].VerifyPeer != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].VerifyPeer = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].VerifyPeer
				}
				if xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].VerifyHost != "0" {
					c.Templates[i].DiscoveryRules[j].Items[l].VerifyHost = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].VerifyHost
				}
				c.Templates[i].DiscoveryRules[j].Items[l].MasterItem = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].MasterItem
				for k := 0; k < len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].ApplicationPrototypes.ApplicationPrototype); k++ {
					c.Templates[i].DiscoveryRules[j].Items[l].Applications = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].ItemPrototypes.ItemPrototype[l].ApplicationPrototypes.ApplicationPrototype[k].Name
				}
			}

			n, _ = strconv.Atoi(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Evaltype)
			c.Templates[i].DiscoveryRules[j].Filter.Evaltype = evalType[n]
			c.Templates[i].DiscoveryRules[j].Filter.Formula = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Formula
			conditionsSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Conditions.Condition)
			c.Templates[i].DiscoveryRules[j].Filter.Conditions = make([]Condition, conditionsSize)
			for k := 0; k < conditionsSize; k++ {
				c.Templates[i].DiscoveryRules[j].Filter.Conditions[k].Formulaid = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Conditions.Condition[k].Formulaid
				c.Templates[i].DiscoveryRules[j].Filter.Conditions[k].Macro = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Conditions.Condition[k].Macro
				c.Templates[i].DiscoveryRules[j].Filter.Conditions[k].Operator = operator[xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Conditions.Condition[k].Operator]
				c.Templates[i].DiscoveryRules[j].Filter.Conditions[k].Value = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].Filter.Conditions.Condition[k].Value
			}

			lldMacroPathsSize := len(xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].LldMacroPaths.LldMacroPath)
			c.Templates[i].DiscoveryRules[j].LldMacroPaths = make([]LlldMacroPath, lldMacroPathsSize)
			for k := 0; k < lldMacroPathsSize; k++ {
				c.Templates[i].DiscoveryRules[j].LldMacroPaths[k].LldMacro = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].LldMacroPaths.LldMacroPath[k].LldMacro
				c.Templates[i].DiscoveryRules[j].LldMacroPaths[k].Path = xmlIn.Templates.Template[i].DiscoveryRules.DiscoveryRule[j].LldMacroPaths.LldMacroPath[k].Path
			}
		}
	}

	triggersSize := len(xmlIn.Triggers.Trigger)
	c.Triggers = make([]Trigger, triggersSize)
	for i := 0; i < triggersSize; i++ {
		c.Triggers[i].Id = "-"
		c.Triggers[i].Name = xmlIn.Triggers.Trigger[i].Name
		c.Triggers[i].Expression = xmlIn.Triggers.Trigger[i].Expression
		n, _ := strconv.Atoi(xmlIn.Triggers.Trigger[i].Priority)
		c.Triggers[i].Priority = priority[n]
		if len(xmlIn.Triggers.Trigger[i].Description) == 0 {
			c.Triggers[i].Description = "-"
		} else {
			c.Triggers[i].Description = xmlIn.Triggers.Trigger[i].Description
		}
		if xmlIn.Triggers.Trigger[i].CorrelationMode != "0" {
			c.Triggers[i].CorrelationMode = xmlIn.Triggers.Trigger[i].CorrelationMode
		}
		c.Triggers[i].CorrelationTag = xmlIn.Triggers.Trigger[i].CorrelationTag
		if xmlIn.Triggers.Trigger[i].ManualClose != "0" {
			c.Triggers[i].ManualClose = xmlIn.Triggers.Trigger[i].ManualClose
		}
		c.Triggers[i].RecoveryExpression = xmlIn.Triggers.Trigger[i].RecoveryExpression
		if xmlIn.Triggers.Trigger[i].RecoveryMode != "0" {
			c.Triggers[i].RecoveryMode = xmlIn.Triggers.Trigger[i].RecoveryMode
		}
		if xmlIn.Triggers.Trigger[i].Status != "0" {
			c.Triggers[i].Status = xmlIn.Triggers.Trigger[i].Status
		}
		c.Triggers[i].Tags = xmlIn.Triggers.Trigger[i].Tags
		if xmlIn.Triggers.Trigger[i].Type != "0" {
			c.Triggers[i].Type = xmlIn.Triggers.Trigger[i].Type
		}
		c.Triggers[i].URL = xmlIn.Triggers.Trigger[i].URL
		c.Triggers[i].Dependencies = xmlIn.Triggers.Trigger[i].Dependencies
	}

	graphsSize := len(xmlIn.Graphs.Graph)
	c.Graphs = make([]Graph, graphsSize)
	for i := 0; i < graphsSize; i++ {
		c.Graphs[i].GraphItems = xmlIn.Graphs.Graph[i].GraphItems
		c.Graphs[i].Height = xmlIn.Graphs.Graph[i].Height
		c.Graphs[i].Name = xmlIn.Graphs.Graph[i].Name
		c.Graphs[i].PercentLeft = xmlIn.Graphs.Graph[i].PercentLeft
		c.Graphs[i].PercentRight = xmlIn.Graphs.Graph[i].PercentRight
		c.Graphs[i].Show3d = xmlIn.Graphs.Graph[i].Show3d
		c.Graphs[i].ShowLegend = xmlIn.Graphs.Graph[i].ShowLegend
		c.Graphs[i].ShowTriggers = xmlIn.Graphs.Graph[i].ShowTriggers
		c.Graphs[i].ShowWorkPeriod = xmlIn.Graphs.Graph[i].ShowWorkPeriod
		c.Graphs[i].Type = xmlIn.Graphs.Graph[i].Type
		c.Graphs[i].Width = xmlIn.Graphs.Graph[i].Width
		c.Graphs[i].Yaxismax = xmlIn.Graphs.Graph[i].Yaxismax
		c.Graphs[i].Yaxismin = xmlIn.Graphs.Graph[i].Yaxismin
		c.Graphs[i].YmaxItem1 = xmlIn.Graphs.Graph[i].YmaxItem1
		c.Graphs[i].YmaxType1 = xmlIn.Graphs.Graph[i].YmaxType1
		c.Graphs[i].YminItem1 = xmlIn.Graphs.Graph[i].YminItem1
		c.Graphs[i].YminType1 = xmlIn.Graphs.Graph[i].YminType1
		c.Graphs[i].GraphItems.GraphItem = xmlIn.Graphs.Graph[i].GraphItems.GraphItem
	}

	*data, err = yaml.Marshal(c)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	fileName := "Template_App_MS_SQL_-_Windows_Server_Failover_Cluster.xml"
	var data []byte

	err := getData(&fileName, &data)
	if err != nil {
		fmt.Printf("Error getting data from a file: %s\n", err)
		return
	}

	var xmlIn ZabbixExport

	err = xmlIn.Parse(&data)
	if err != nil {
		fmt.Printf("Error parsing XML data: %s\n", err)
		return
	}

	xmlOut, err := xml.MarshalIndent(xmlIn, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling XML file: %s\n", err)
		return
	}

	// re := regexp.MustCompile(filepath.Ext(fileName) + `$`)
	// fileName = re.ReplaceAllLiteralString(fileName, ".xml")

	err = ioutil.WriteFile("fileName.xml", xmlOut, os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing XML file: %s\n", err)
		return
	}

	var ymlIn Templator

	data = nil
	err = ymlIn.Export(&xmlIn, &data)
	if err != nil {
		fmt.Printf("Error marshaling YAML file: %s\n", err)
		return
	}

	err = ioutil.WriteFile("fileName.yml", data, os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing YAML file: %s\n", err)
		return
	}

}
