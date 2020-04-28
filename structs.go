package main

type Item struct {
	Name                 string `xml:"name"`
	Type                 string `xml:"type"`
	SnmpCommunity        string `xml:"snmp_community"`
	SnmpOid              string `xml:"snmp_oid"`
	Key                  string `xml:"key"`
	Delay                string `xml:"delay"`
	History              string `xml:"history"`
	Trends               string `xml:"trends"`
	Status               string `xml:"status"`
	ValueType            string `xml:"value_type"`
	AllowedHosts         string `xml:"allowed_hosts"`
	Units                string `xml:"units"`
	Snmpv3Contextname    string `xml:"snmpv3_contextname"`
	Snmpv3Securityname   string `xml:"snmpv3_securityname"`
	Snmpv3Securitylevel  string `xml:"snmpv3_securitylevel"`
	Snmpv3Authprotocol   string `xml:"snmpv3_authprotocol"`
	Snmpv3Authpassphrase string `xml:"snmpv3_authpassphrase"`
	Snmpv3Privprotocol   string `xml:"snmpv3_privprotocol"`
	Snmpv3Privpassphrase string `xml:"snmpv3_privpassphrase"`
	Params               string `xml:"params"`
	IpmiSensor           string `xml:"ipmi_sensor"`
	Authtype             string `xml:"authtype"`
	Username             string `xml:"username"`
	Password             string `xml:"password"`
	Publickey            string `xml:"publickey"`
	Privatekey           string `xml:"privatekey"`
	Port                 string `xml:"port"`
	Description          string `xml:"description"`
	InventoryLink        string `xml:"inventory_link"`
	Applications         struct {
		Application []struct {
			Name string `xml:"name"`
		} `xml:"application"`
	} `xml:"applications"`
	Valuemap struct {
		Name string `xml:"name"`
	} `xml:"valuemap"`
	Logtimefmt    string `xml:"logtimefmt"`
	Preprocessing struct {
		Step struct {
			Type   string `xml:"type"`
			Params string `xml:"params"`
		} `xml:"step"`
	} `xml:"preprocessing"`
	JmxEndpoint           string `xml:"jmx_endpoint"`
	Timeout               string `xml:"timeout"`
	URL                   string `xml:"url"`
	QueryFields           string `xml:"query_fields"`
	Posts                 string `xml:"posts"`
	StatusCodes           string `xml:"status_codes"`
	FollowRedirects       string `xml:"follow_redirects"`
	PostType              string `xml:"post_type"`
	HTTPProxy             string `xml:"http_proxy"`
	Headers               string `xml:"headers"`
	RetrieveMode          string `xml:"retrieve_mode"`
	RequestMethod         string `xml:"request_method"`
	OutputFormat          string `xml:"output_format"`
	AllowTraps            string `xml:"allow_traps"`
	SslCertFile           string `xml:"ssl_cert_file"`
	SslKeyFile            string `xml:"ssl_key_file"`
	SslKeyPassword        string `xml:"ssl_key_password"`
	VerifyPeer            string `xml:"verify_peer"`
	VerifyHost            string `xml:"verify_host"`
	MasterItem            string `xml:"master_item"`
	ApplicationPrototypes struct {
		ApplicationPrototype struct {
			Name string `xml:"name"`
		} `xml:"application_prototype"`
	} `xml:"application_prototypes"`
}

type ZabbixExport struct {
	Version   string `xml:"version"`
	Templates struct {
		Template struct {
			Template    string `xml:"template"`
			Name        string `xml:"name"`
			Description string `xml:"description"`
			Items       struct {
				Item []Item `xml:"item"`
			} `xml:"items"`
			DiscoveryRules struct {
				DiscoveryRule struct {
					Name                 string `xml:"name"`
					Type                 string `xml:"type"`
					SnmpCommunity        string `xml:"snmp_community"`
					SnmpOid              string `xml:"snmp_oid"`
					Key                  string `xml:"key"`
					Delay                string `xml:"delay"`
					Status               string `xml:"status"`
					AllowedHosts         string `xml:"allowed_hosts"`
					Snmpv3Contextname    string `xml:"snmpv3_contextname"`
					Snmpv3Securityname   string `xml:"snmpv3_securityname"`
					Snmpv3Securitylevel  string `xml:"snmpv3_securitylevel"`
					Snmpv3Authprotocol   string `xml:"snmpv3_authprotocol"`
					Snmpv3Authpassphrase string `xml:"snmpv3_authpassphrase"`
					Snmpv3Privprotocol   string `xml:"snmpv3_privprotocol"`
					Snmpv3Privpassphrase string `xml:"snmpv3_privpassphrase"`
					Params               string `xml:"params"`
					IpmiSensor           string `xml:"ipmi_sensor"`
					Authtype             string `xml:"authtype"`
					Username             string `xml:"username"`
					Password             string `xml:"password"`
					Publickey            string `xml:"publickey"`
					Privatekey           string `xml:"privatekey"`
					Port                 string `xml:"port"`
					Filter               struct {
						Evaltype   string `xml:"evaltype"`
						Formula    string `xml:"formula"`
						Conditions struct {
							Condition struct {
								Macro     string `xml:"macro"`
								Value     string `xml:"value"`
								Operator  string `xml:"operator"`
								Formulaid string `xml:"formulaid"`
							} `xml:"condition"`
						} `xml:"conditions"`
					} `xml:"filter"`
					Lifetime       string `xml:"lifetime"`
					Description    string `xml:"description"`
					ItemPrototypes struct {
						ItemPrototype []Item `xml:"item_prototype"`
					} `xml:"item_prototypes"`
					TriggerPrototypes struct {
						TriggerPrototype []struct {
							Expression         string `xml:"expression"`
							RecoveryMode       string `xml:"recovery_mode"`
							RecoveryExpression string `xml:"recovery_expression"`
							Name               string `xml:"name"`
							CorrelationMode    string `xml:"correlation_mode"`
							CorrelationTag     string `xml:"correlation_tag"`
							URL                string `xml:"url"`
							Status             string `xml:"status"`
							Priority           string `xml:"priority"`
							Description        string `xml:"description"`
							Type               string `xml:"type"`
							ManualClose        string `xml:"manual_close"`
							Dependencies       string `xml:"dependencies"`
							Tags               string `xml:"tags"`
						} `xml:"trigger_prototype"`
					} `xml:"trigger_prototypes"`
					GraphPrototypes string `xml:"graph_prototypes"`
					HostPrototypes  string `xml:"host_prototypes"`
					JmxEndpoint     string `xml:"jmx_endpoint"`
					Timeout         string `xml:"timeout"`
					URL             string `xml:"url"`
					QueryFields     string `xml:"query_fields"`
					Posts           string `xml:"posts"`
					StatusCodes     string `xml:"status_codes"`
					FollowRedirects string `xml:"follow_redirects"`
					PostType        string `xml:"post_type"`
					HTTPProxy       string `xml:"http_proxy"`
					Headers         string `xml:"headers"`
					RetrieveMode    string `xml:"retrieve_mode"`
					RequestMethod   string `xml:"request_method"`
					AllowTraps      string `xml:"allow_traps"`
					SslCertFile     string `xml:"ssl_cert_file"`
					SslKeyFile      string `xml:"ssl_key_file"`
					SslKeyPassword  string `xml:"ssl_key_password"`
					VerifyPeer      string `xml:"verify_peer"`
					VerifyHost      string `xml:"verify_host"`
				} `xml:"discovery_rule"`
			} `xml:"discovery_rules"`
			Httptests string `xml:"httptests"`
			Macros    struct {
				Macro []struct {
					Macro string `xml:"macro"`
					Value string `xml:"value"`
				} `xml:"macro"`
			} `xml:"macros"`
			Templates string `xml:"templates"`
			Screens   string `xml:"screens"`
		} `xml:"template"`
	} `xml:"templates"`
	Triggers struct {
		Trigger []struct {
			Expression         string `xml:"expression"`
			RecoveryMode       string `xml:"recovery_mode"`
			RecoveryExpression string `xml:"recovery_expression"`
			Name               string `xml:"name"`
			CorrelationMode    string `xml:"correlation_mode"`
			CorrelationTag     string `xml:"correlation_tag"`
			URL                string `xml:"url"`
			Status             string `xml:"status"`
			Priority           string `xml:"priority"`
			Description        string `xml:"description"`
			Type               string `xml:"type"`
			ManualClose        string `xml:"manual_close"`
			Dependencies       struct {
				Dependency struct {
					Name               string `xml:"name"`
					Expression         string `xml:"expression"`
					RecoveryExpression string `xml:"recovery_expression"`
				} `xml:"dependency"`
			} `xml:"dependencies"`
			Tags string `xml:"tags"`
		} `xml:"trigger"`
	} `xml:"triggers"`
	Graphs struct {
		Graph []struct {
			Name           string `xml:"name"`
			Width          string `xml:"width"`
			Height         string `xml:"height"`
			Yaxismin       string `xml:"yaxismin"`
			Yaxismax       string `xml:"yaxismax"`
			ShowWorkPeriod string `xml:"show_work_period"`
			ShowTriggers   string `xml:"show_triggers"`
			Type           string `xml:"type"`
			ShowLegend     string `xml:"show_legend"`
			Show3d         string `xml:"show_3d"`
			PercentLeft    string `xml:"percent_left"`
			PercentRight   string `xml:"percent_right"`
			YminType1      string `xml:"ymin_type_1"`
			YmaxType1      string `xml:"ymax_type_1"`
			YminItem1      string `xml:"ymin_item_1"`
			YmaxItem1      string `xml:"ymax_item_1"`
			GraphItems     struct {
				GraphItem []struct {
					Sortorder string `xml:"sortorder"`
					Drawtype  string `xml:"drawtype"`
					Color     string `xml:"color"`
					Yaxisside string `xml:"yaxisside"`
					CalcFnc   string `xml:"calc_fnc"`
					Type      string `xml:"type"`
					Item      struct {
						Host string `xml:"host"`
						Key  string `xml:"key"`
					} `xml:"item"`
				} `xml:"graph_item"`
			} `xml:"graph_items"`
		} `xml:"graph"`
	} `xml:"graphs"`
	ValueMaps struct {
		ValueMap []struct {
			Name     string `xml:"name"`
			Mappings struct {
				Mapping []struct {
					Value    string `xml:"value"`
					Newvalue string `xml:"newvalue"`
				} `xml:"mapping"`
			} `xml:"mappings"`
		} `xml:"value_map"`
	} `xml:"value_maps"`
}
