package main

type Item struct {
	Name                 string `xml:"name" yaml:"name,omitempty"`
	Type                 string `xml:"type" yaml:"type,omitempty"`
	SnmpCommunity        string `xml:"snmp_community" yaml:"snmp_community,omitempty"`
	SnmpOid              string `xml:"snmp_oid" yaml:"snmp_oid,omitempty"`
	Key                  string `xml:"key" yaml:"key,omitempty"`
	Delay                string `xml:"delay" yaml:"delay,omitempty"`
	History              string `xml:"history" yaml:"history,omitempty"`
	Trends               string `xml:"trends" yaml:"trends,omitempty"`
	Status               string `xml:"status" yaml:"status,omitempty"`
	ValueType            string `xml:"value_type" yaml:"value_type,omitempty"`
	AllowedHosts         string `xml:"allowed_hosts" yaml:"allowed_hosts,omitempty"`
	Units                string `xml:"units" yaml:"units,omitempty"`
	Snmpv3Contextname    string `xml:"snmpv3_contextname" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname   string `xml:"snmpv3_securityname" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel  string `xml:"snmpv3_securitylevel" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol   string `xml:"snmpv3_authprotocol" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase string `xml:"snmpv3_authpassphrase" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol   string `xml:"snmpv3_privprotocol" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase string `xml:"snmpv3_privpassphrase" yaml:"snmpv3_privpassphrase,omitempty"`
	Params               string `xml:"params" yaml:"params,omitempty"`
	IpmiSensor           string `xml:"ipmi_sensor" yaml:"ipmi_sensor,omitempty"`
	Authtype             string `xml:"authtype" yaml:"authtype,omitempty"`
	Username             string `xml:"username" yaml:"username,omitempty"`
	Password             string `xml:"password" yaml:"password,omitempty"`
	Publickey            string `xml:"publickey" yaml:"publickey,omitempty"`
	Privatekey           string `xml:"privatekey" yaml:"privatekey,omitempty"`
	Port                 string `xml:"port" yaml:"port,omitempty"`
	Description          string `xml:"description" yaml:"description,omitempty"`
	InventoryLink        string `xml:"inventory_link" yaml:"inventory_link,omitempty"`
	Applications         struct {
		Application []struct {
			Name string `xml:"name" yaml:"name,omitempty"`
		} `xml:"application" yaml:"application,omitempty"`
	} `xml:"applications" yaml:"applications,omitempty"`
	Valuemap struct {
		Name string `xml:"name" yaml:"name,omitempty"`
	} `xml:"valuemap" yaml:"valuemap,omitempty"`
	Logtimefmt    string `xml:"logtimefmt" yaml:"logtimefmt,omitempty"`
	Preprocessing struct {
		Step []struct {
			Type               string `xml:"type" yaml:"type,omitempty"`
			Params             string `xml:"params" yaml:"params,omitempty"`
			ErrorHandler       string `xml:"error_handler" yaml:"error_handler,omitempty"`
			ErrorHandlerParams string `xml:"error_handler_params" yaml:"error_handler_params,omitempty"`
		} `xml:"step" yaml:"step,omitempty"`
	} `xml:"preprocessing" yaml:"preprocessing,omitempty"`
	JmxEndpoint           string `xml:"jmx_endpoint" yaml:"jmx_endpoint,omitempty"`
	Timeout               string `xml:"timeout" yaml:"timeout,omitempty"`
	URL                   string `xml:"url" yaml:"url,omitempty"`
	QueryFields           string `xml:"query_fields" yaml:"query_fields,omitempty"`
	Posts                 string `xml:"posts" yaml:"posts,omitempty"`
	StatusCodes           string `xml:"status_codes" yaml:"status_codes,omitempty"`
	FollowRedirects       string `xml:"follow_redirects" yaml:"follow_redirects,omitempty"`
	PostType              string `xml:"post_type" yaml:"post_type,omitempty"`
	HTTPProxy             string `xml:"http_proxy" yaml:"http_proxy,omitempty"`
	Headers               string `xml:"headers" yaml:"headers,omitempty"`
	RetrieveMode          string `xml:"retrieve_mode" yaml:"retrieve_mode,omitempty"`
	RequestMethod         string `xml:"request_method" yaml:"request_method,omitempty"`
	OutputFormat          string `xml:"output_format" yaml:"output_format,omitempty"`
	AllowTraps            string `xml:"allow_traps" yaml:"allow_traps,omitempty"`
	SslCertFile           string `xml:"ssl_cert_file" yaml:"ssl_cert_file,omitempty"`
	SslKeyFile            string `xml:"ssl_key_file" yaml:"ssl_key_file,omitempty"`
	SslKeyPassword        string `xml:"ssl_key_password" yaml:"ssl_key_password,omitempty"`
	VerifyPeer            string `xml:"verify_peer" yaml:"verify_peer,omitempty"`
	VerifyHost            string `xml:"verify_host" yaml:"verify_host,omitempty"`
	MasterItem            string `xml:"master_item" yaml:"master_item,omitempty"`
	ApplicationPrototypes struct {
		ApplicationPrototype []struct {
			Name string `xml:"name" yaml:"name,omitempty"`
		} `xml:"application_prototype" yaml:"application_prototype,omitempty"`
	} `xml:"application_prototypes" yaml:"application_prototypes,omitempty"`
}

type ZabbixExport struct {
	Version   string `xml:"version" yaml:"version,omitempty"`
	Templates struct {
		Template []struct {
			// Template    string `xml:"template" yaml:"template,omitempty"`
			Name        string `xml:"name" yaml:"name,omitempty"`
			Description string `xml:"description" yaml:"description,omitempty"`
			Items       struct {
				Item []Item `xml:"item" yaml:"item,omitempty"`
			} `xml:"items" yaml:"items,omitempty"`
			DiscoveryRules struct {
				DiscoveryRule []struct {
					Name                 string `xml:"name" yaml:"name,omitempty"`
					Type                 string `xml:"type" yaml:"type,omitempty"`
					SnmpCommunity        string `xml:"snmp_community" yaml:"snmp_community,omitempty"`
					SnmpOid              string `xml:"snmp_oid" yaml:"snmp_oid,omitempty"`
					Key                  string `xml:"key" yaml:"key,omitempty"`
					Delay                string `xml:"delay" yaml:"delay,omitempty"`
					Status               string `xml:"status" yaml:"status,omitempty"`
					AllowedHosts         string `xml:"allowed_hosts" yaml:"allowed_hosts,omitempty"`
					Snmpv3Contextname    string `xml:"snmpv3_contextname" yaml:"snmpv3_contextname,omitempty"`
					Snmpv3Securityname   string `xml:"snmpv3_securityname" yaml:"snmpv3_securityname,omitempty"`
					Snmpv3Securitylevel  string `xml:"snmpv3_securitylevel" yaml:"snmpv3_securitylevel,omitempty"`
					Snmpv3Authprotocol   string `xml:"snmpv3_authprotocol" yaml:"snmpv3_authprotocol,omitempty"`
					Snmpv3Authpassphrase string `xml:"snmpv3_authpassphrase" yaml:"snmpv3_authpassphrase,omitempty"`
					Snmpv3Privprotocol   string `xml:"snmpv3_privprotocol" yaml:"snmpv3_privprotocol,omitempty"`
					Snmpv3Privpassphrase string `xml:"snmpv3_privpassphrase" yaml:"snmpv3_privpassphrase,omitempty"`
					Params               string `xml:"params" yaml:"params,omitempty"`
					IpmiSensor           string `xml:"ipmi_sensor" yaml:"ipmi_sensor,omitempty"`
					Authtype             string `xml:"authtype" yaml:"authtype,omitempty"`
					Username             string `xml:"username" yaml:"username,omitempty"`
					Password             string `xml:"password" yaml:"password,omitempty"`
					Publickey            string `xml:"publickey" yaml:"publickey,omitempty"`
					Privatekey           string `xml:"privatekey" yaml:"privatekey,omitempty"`
					Port                 string `xml:"port" yaml:"port,omitempty"`
					Lifetime             string `xml:"lifetime" yaml:"lifetime,omitempty"`
					Description          string `xml:"description" yaml:"description,omitempty"`
					Preprocessing        struct {
						Step []struct {
							Type               string `xml:"type" yaml:"type,omitempty"`
							Params             string `xml:"params" yaml:"params,omitempty"`
							ErrorHandler       string `xml:"error_handler" yaml:"error_handler,omitempty"`
							ErrorHandlerParams string `xml:"error_handler_params" yaml:"error_handler_params,omitempty"`
						} `xml:"step" yaml:"step,omitempty"`
					} `xml:"preprocessing" yaml:"preprocessing,omitempty"`
					JmxEndpoint     string `xml:"jmx_endpoint" yaml:"jmx_endpoint,omitempty"`
					Timeout         string `xml:"timeout" yaml:"timeout,omitempty"`
					URL             string `xml:"url" yaml:"url,omitempty"`
					QueryFields     string `xml:"query_fields" yaml:"query_fields,omitempty"`
					Posts           string `xml:"posts" yaml:"posts,omitempty"`
					StatusCodes     string `xml:"status_codes" yaml:"status_codes,omitempty"`
					FollowRedirects string `xml:"follow_redirects" yaml:"follow_redirects,omitempty"`
					PostType        string `xml:"post_type" yaml:"post_type,omitempty"`
					HTTPProxy       string `xml:"http_proxy" yaml:"http_proxy,omitempty"`
					Headers         string `xml:"headers" yaml:"headers,omitempty"`
					RetrieveMode    string `xml:"retrieve_mode" yaml:"retrieve_mode,omitempty"`
					RequestMethod   string `xml:"request_method" yaml:"request_method,omitempty"`
					AllowTraps      string `xml:"allow_traps" yaml:"allow_traps,omitempty"`
					SslCertFile     string `xml:"ssl_cert_file" yaml:"ssl_cert_file,omitempty"`
					SslKeyFile      string `xml:"ssl_key_file" yaml:"ssl_key_file,omitempty"`
					SslKeyPassword  string `xml:"ssl_key_password" yaml:"ssl_key_password,omitempty"`
					VerifyPeer      string `xml:"verify_peer" yaml:"verify_peer,omitempty"`
					VerifyHost      string `xml:"verify_host" yaml:"verify_host,omitempty"`
					MasterItem      string `xml:"master_item" yaml:"master_item,omitempty"`
					Filter          struct {
						Evaltype   string `xml:"evaltype" yaml:"evaltype,omitempty"`
						Formula    string `xml:"formula" yaml:"formula,omitempty"`
						Conditions struct {
							Condition []struct {
								Macro     string `xml:"macro" yaml:"macro,omitempty"`
								Value     string `xml:"value" yaml:"value,omitempty"`
								Operator  string `xml:"operator" yaml:"operator,omitempty"`
								Formulaid string `xml:"formulaid" yaml:"formulaid,omitempty"`
							} `xml:"condition" yaml:"condition,omitempty"`
						} `xml:"conditions" yaml:"conditions,omitempty"`
					} `xml:"filter" yaml:"filter,omitempty"`
					LldMacroPaths struct {
						LldMacroPath []struct {
							LldMacro string `xml:"lld_macro"`
							Path     string `xml:"path"`
						} `xml:"lld_macro_path" yaml:"lld_macro_path,omitempty"`
					} `xml:"lld_macro_paths" yaml:"lld_macro_paths,omitempty"`
					ItemPrototypes struct {
						ItemPrototype []Item `xml:"item_prototype" yaml:"item_prototype,omitempty"`
					} `xml:"item_prototypes" yaml:"item_prototypes,omitempty"`
					TriggerPrototypes struct {
						TriggerPrototype []struct {
							Expression         string `xml:"expression" yaml:"expression,omitempty"`
							RecoveryMode       string `xml:"recovery_mode" yaml:"recovery_mode,omitempty"`
							RecoveryExpression string `xml:"recovery_expression" yaml:"recovery_expression,omitempty"`
							Name               string `xml:"name" yaml:"name,omitempty"`
							CorrelationMode    string `xml:"correlation_mode" yaml:"correlation_mode,omitempty"`
							CorrelationTag     string `xml:"correlation_tag" yaml:"correlation_tag,omitempty"`
							URL                string `xml:"url" yaml:"url,omitempty"`
							Status             string `xml:"status" yaml:"status,omitempty"`
							Priority           string `xml:"priority" yaml:"priority,omitempty"`
							Description        string `xml:"description" yaml:"description,omitempty"`
							Type               string `xml:"type" yaml:"type,omitempty"`
							ManualClose        string `xml:"manual_close" yaml:"manual_close,omitempty"`
							Dependencies       string `xml:"dependencies" yaml:"dependencies,omitempty"`
							Tags               string `xml:"tags" yaml:"tags,omitempty"`
						} `xml:"trigger_prototype" yaml:"trigger_prototype,omitempty"`
					} `xml:"trigger_prototypes" yaml:"trigger_prototypes,omitempty"`
					GraphPrototypes string `xml:"graph_prototypes" yaml:"graph_prototypes,omitempty"`
					HostPrototypes  string `xml:"host_prototypes" yaml:"host_prototypes,omitempty"`
				} `xml:"discovery_rule" yaml:"discovery_rule,omitempty"`
			} `xml:"discovery_rules" yaml:"discovery_rules,omitempty"`
			Httptests string `xml:"httptests" yaml:"httptests,omitempty"`
			Macros    struct {
				Macro []struct {
					Macro string `xml:"macro" yaml:"macro,omitempty"`
					Value string `xml:"value" yaml:"value,omitempty"`
				} `xml:"macro" yaml:"macro,omitempty"`
			} `xml:"macros" yaml:"macros,omitempty"`
			Templates string `xml:"templates" yaml:"templates,omitempty"`
			Screens   string `xml:"screens" yaml:"screens,omitempty"`
		} `xml:"template" yaml:"template,omitempty"`
	} `xml:"templates" yaml:"templates,omitempty"`
	Triggers struct {
		Trigger []struct {
			Expression         string `xml:"expression" yaml:"expression,omitempty"`
			RecoveryMode       string `xml:"recovery_mode" yaml:"recovery_mode,omitempty"`
			RecoveryExpression string `xml:"recovery_expression" yaml:"recovery_expression,omitempty"`
			Name               string `xml:"name" yaml:"name,omitempty"`
			CorrelationMode    string `xml:"correlation_mode" yaml:"correlation_mode,omitempty"`
			CorrelationTag     string `xml:"correlation_tag" yaml:"correlation_tag,omitempty"`
			URL                string `xml:"url" yaml:"url,omitempty"`
			Status             string `xml:"status" yaml:"status,omitempty"`
			Priority           string `xml:"priority" yaml:"priority,omitempty"`
			Description        string `xml:"description" yaml:"description,omitempty"`
			Type               string `xml:"type" yaml:"type,omitempty"`
			ManualClose        string `xml:"manual_close" yaml:"manual_close,omitempty"`
			Dependencies       struct {
				Dependency struct {
					Name               string `xml:"name" yaml:"name,omitempty"`
					Expression         string `xml:"expression" yaml:"expression,omitempty"`
					RecoveryExpression string `xml:"recovery_expression" yaml:"recovery_expression,omitempty"`
				} `xml:"dependency" yaml:"dependency,omitempty"`
			} `xml:"dependencies" yaml:"dependencies,omitempty"`
			Tags string `xml:"tags" yaml:"tags,omitempty"`
		} `xml:"trigger" yaml:"trigger,omitempty"`
	} `xml:"triggers" yaml:"triggers,omitempty"`
	Graphs struct {
		Graph []struct {
			Name           string `xml:"name" yaml:"name,omitempty"`
			Width          string `xml:"width" yaml:"width,omitempty"`
			Height         string `xml:"height" yaml:"height,omitempty"`
			Yaxismin       string `xml:"yaxismin" yaml:"yaxismin,omitempty"`
			Yaxismax       string `xml:"yaxismax" yaml:"yaxismax,omitempty"`
			ShowWorkPeriod string `xml:"show_work_period" yaml:"show_work_period,omitempty"`
			ShowTriggers   string `xml:"show_triggers" yaml:"show_triggers,omitempty"`
			Type           string `xml:"type" yaml:"type,omitempty"`
			ShowLegend     string `xml:"show_legend" yaml:"show_legend,omitempty"`
			Show3d         string `xml:"show_3d" yaml:"show_3d,omitempty"`
			PercentLeft    string `xml:"percent_left" yaml:"percent_left,omitempty"`
			PercentRight   string `xml:"percent_right" yaml:"percent_right,omitempty"`
			YminType1      string `xml:"ymin_type_1" yaml:"ymin_type_1,omitempty"`
			YmaxType1      string `xml:"ymax_type_1" yaml:"ymax_type_1,omitempty"`
			YminItem1      string `xml:"ymin_item_1" yaml:"ymin_item_1,omitempty"`
			YmaxItem1      string `xml:"ymax_item_1" yaml:"ymax_item_1,omitempty"`
			GraphItems     struct {
				GraphItem []struct {
					Sortorder string `xml:"sortorder" yaml:"sortorder,omitempty"`
					Drawtype  string `xml:"drawtype" yaml:"drawtype,omitempty"`
					Color     string `xml:"color" yaml:"color,omitempty"`
					Yaxisside string `xml:"yaxisside" yaml:"yaxisside,omitempty"`
					CalcFnc   string `xml:"calc_fnc" yaml:"calc_fnc,omitempty"`
					Type      string `xml:"type" yaml:"type,omitempty"`
					Item      struct {
						Host string `xml:"host" yaml:"host,omitempty"`
						Key  string `xml:"key" yaml:"key,omitempty"`
					} `xml:"item" yaml:"item,omitempty"`
				} `xml:"graph_item" yaml:"graph_item,omitempty"`
			} `xml:"graph_items" yaml:"graph_items,omitempty"`
		} `xml:"graph" yaml:"graph,omitempty"`
	} `xml:"graphs" yaml:"graphs,omitempty"`
	ValueMaps struct {
		ValueMap []struct {
			Name     string `xml:"name" yaml:"name,omitempty"`
			Mappings struct {
				Mapping []struct {
					Value    string `xml:"value" yaml:"value,omitempty"`
					Newvalue string `xml:"newvalue" yaml:"newvalue,omitempty"`
				} `xml:"mapping" yaml:"mapping,omitempty"`
			} `xml:"mappings" yaml:"mappings,omitempty"`
		} `xml:"value_map" yaml:"value_map,omitempty"`
	} `xml:"value_maps" yaml:"value_maps,omitempty"`
}
