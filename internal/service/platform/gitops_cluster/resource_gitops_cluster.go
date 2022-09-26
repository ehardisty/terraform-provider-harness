package gitops_cluster

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGitopsCluster() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Gitops Cluster.",

		CreateContext: resourceGitopsClusterCreate,
		ReadContext:   resourceGitopsClusterRead,
		UpdateContext: resourceGitopsClusterUpdate,
		DeleteContext: resourceGitopsClusterDelete,
		Importer:      helpers.OrgResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_identifier": {
				Description: "account identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_identifier": {
				Description: "project identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_identifier": {
				Description: "organization identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"agent_identifier": {
				Description: "agent identifier of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"query": {
				Description: "query for cluster resources",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Description: "server of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "name of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": {
							Description: "cluster server URL or cluster name",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "the type of the specified cluster identifier ( 'server' - default, 'name' )",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"value": {
										Description: "the cluster server URL or cluster name",
										Type:        schema.TypeString,
										Optional:    true,
									},
								}},
						},
					},
				},
			},
			"request": {
				Description: "Cluster create/Update request.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upsert": {
							Description: "if the cluster should be upserted.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"updated_fields": {
							Description: "Fields which are updated.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"update_mask": {
							Description: "Update mask of the cluster.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"paths": {
										Description: "The set of field mask paths.",
										Optional:    true,
										Type:        schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeList,
										},
									},
								},
							},
						},
						"id": {
							Description: "cluster server URL or cluster name",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "the type of the specified cluster identifier ( 'server' - default, 'name' )",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"value": {
										Description: "the cluster server URL or cluster name",
										Type:        schema.TypeString,
										Optional:    true,
									},
								}},
						},
						"cluster": {
							Description: "cluster details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": {
										Description: "the API server URL of the Kubernetes cluster.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"name": {
										Description: "Name of the cluster. If omitted, will use the server address",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"config": {
										Description: "Cluster Config",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"username": {
													Description: "username for the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"password": {
													Description: "password for the server of the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"bearer_token": {
													Description: "Bearer authentication token the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"tls_client_config": {
													Description: "contains settings to enable transport layer security",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"insecure": {
																Description: "if the TLS connection to the cluster should be insecure.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"server_name": {
																Description: "server name for SNI in the client to check server certificates against",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"cert_data": {
																Description: "certficate data holds PEM-encoded bytes (typically read from a client certificate file).",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"key_data": {
																Description: "key data holds PEM-encoded bytes (typically read from a client certificate key file).",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"ca_data": {
																Description: "holds PEM-encoded bytes (typically read from a root certificates bundle).",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"aws_auth_config": {
													Description: "contains IAM authentication configuration",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_name": {
																Description: "contains AWS cluster name.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"role_a_r_n": {
																Description: "contains optional role ARN. If set then AWS IAM Authenticator.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"exec_provider_config": {
													Description: "contains configuration for an exec provider",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"command": {
																Description: "command to execute.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"args": {
																Description: "Arguments to pass to the command when executing it.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"env": {
																Description: "additional environment variables to expose to the process.",
																Type:        schema.TypeMap,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"api_version": {
																Description: "Preferred input version of the ExecInfo.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"install_hint": {
																Description: "This text is shown to the user when the executable doesn't seem to be present",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"cluster_connection_type": {
													Description: "Identifies the authentication method used to connect to the cluster.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"namespaces": {
										Description: "list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"refresh_requested_at": {
										Description: "time when cluster cache refresh has been requested.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"seconds": {
													Description: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"nanos": {
													Description: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
													Type:        schema.TypeInt,
													Optional:    true,
												},
											},
										},
									},
									"info": {
										Description: "information about cluster cache and state",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_state": {
													Description: "information about the connection to the cluster",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"status": {
																Description: "the current status indicator for the connection",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"message": {
																Description: "human readable information about the connection status",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"attempted_at": {
																Description: "time when cluster cache refresh has been requested.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"seconds": {
																			Description: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"nanos": {
																			Description: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"server_version": {
													Description: "information about the Kubernetes version of the cluster",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"cache_info": {
													Description: "information about the cluster cache",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"resources_count": {
																Description: "number of observed Kubernetes resources",
																Type:        schema.TypeInt,
																Optional:    true,
															},
															"apis_count": {
																Description: "number of observed Kubernetes API count",
																Type:        schema.TypeInt,
																Optional:    true,
															},
															"last_cache_sync_time": {
																Description: "time of most recent cache synchronization",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"applications_count": {
													Description: "the number of applications managed by Argo CD on the cluster",
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"api_versions": {
													Description: "list of API versions supported by the cluster",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"shard": {
										Description: " optional shard number. Calculated on the fly by the application controller if not specified.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"cluster_resources": {
										Description: "Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"project": {
										Description: "Reference between project and cluster that allow you automatically to be added as item inside Destinations project entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"labels": {
										Description: "Labels for cluster secret metadata",
										Type:        schema.TypeMap,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"annotations": {
										Description: "Annotations for cluster secret metadata",
										Type:        schema.TypeMap,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceGitopsClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	if attr, ok := d.GetOk("agent_identifier"); ok {
		agentIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("account_identifier"); ok {
		accountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_identifier"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_identifier"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}

	createClusterRequest := buildCreateClusterRequest(d)
	resp, httpResp, err := c.AgentClusterApi.AgentClusterServiceCreate(ctx, *createClusterRequest, agentIdentifier,
		&nextgen.AgentClusterServiceApiAgentClusterServiceCreateOpts{
			AccountIdentifier: optional.NewString(accountIdentifier),
			OrgIdentifier:     optional.NewString(orgIdentifier),
			ProjectIdentifier: optional.NewString(projectIdentifier),
			Identifier:        optional.NewString(identifier),
		})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	agentIdentifier := d.Get("agent_identifier").(string)
	identifier := d.Get("identifier").(string)
	var queryName, queryServer string
	if d.Get("query") != nil && len(d.Get("query").([]interface{})) > 0 {
		query := d.Get("query").([]interface{})[0].(map[string]interface{})
		queryServer = query["server"].(string)
		queryName = query["name"].(string)
		// queryId := query["id"].(map[string]interface{})
		// queryIdType = queryId["type"].(string)
		// queryIdValue = queryId["value"].(string)
	}
	resp, httpResp, err := c.AgentClusterApi.AgentClusterServiceGet(ctx, agentIdentifier, identifier, &nextgen.AgentClusterServiceApiAgentClusterServiceGetOpts{
		AccountIdentifier: optional.NewString(d.Get("account_identifier").(string)),
		OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		QueryServer:       optional.NewString(queryServer),
		QueryName:         optional.NewString(queryName),
		// QueryIdType:       optional.NewString(queryIdType),
		// QueryIdValue:      optional.NewString(queryIdValue),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	agentIdentifier := d.Get("agent_identifier").(string)
	identifier := d.Get("identifier").(string)
	updateClusterRequest := buildUpdateClusterRequest(d)
	resp, httpResp, err := c.AgentClusterApi.AgentClusterServiceUpdate(ctx, *updateClusterRequest, agentIdentifier, identifier,
		&nextgen.AgentClusterServiceApiAgentClusterServiceUpdateOpts{
			AccountIdentifier: optional.NewString(d.Get("account_identifier").(string)),
			OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Cluster == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	setClusterDetails(d, &resp)
	return nil
}

func resourceGitopsClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	agentIdentifier := d.Get("agent_identifier").(string)
	identifier := d.Get("identifier").(string)
	_, httpResp, err := c.AgentClusterApi.AgentClusterServiceDelete(ctx, agentIdentifier, identifier, &nextgen.AgentClusterServiceApiAgentClusterServiceDeleteOpts{
		AccountIdentifier: optional.NewString(d.Get("account_identifier").(string)),
		OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		QueryServer:       optional.NewString(d.Get("query.server").(string)),
		QueryName:         optional.NewString(d.Get("query.name").(string)),
		// QueryIdType:       optional.NewString(d.Get("query.id.type").(string)),
		// QueryIdValue:      optional.NewString(d.Get("query.id.value").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func setClusterDetails(d *schema.ResourceData, cl *nextgen.Servicev1Cluster) {
	d.Set("account_identifier", cl.AccountIdentifier)
	d.Set("organization_identifier", cl.OrgIdentifier)
	d.Set("project_identifer", cl.ProjectIdentifier)
	d.Set("agent_identifier", cl.AgentIdentifier)
	d.Set("identifier", cl.Identifier)
	d.Set("created_at", cl.CreatedAt)
	d.Set("last_modified_at", cl.LastModifiedAt)
	d.Set("cluster", cl.Cluster)
}

func buildCreateClusterRequest(d *schema.ResourceData) *nextgen.ClusterClusterCreateRequest {
	var upsert bool
	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})
		upsert = request["upsert"].(bool)
	}
	return &nextgen.ClusterClusterCreateRequest{
		Upsert:  upsert,
		Cluster: buildClusterDetails(d),
	}
}

func buildUpdateClusterRequest(d *schema.ResourceData) *nextgen.ClusterClusterUpdateRequest {
	var request map[string]interface{}
	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})
	}
	var updatedFields []string
	if request["updated_fields"] != nil {
		updatedFields = request["updated_fields"].([]string)
	}

	var updateMask map[string]interface{}
	if request["update_mask"] != nil {
		updateMask = request["update_mask"].(map[string]interface{})
	}

	var updateMaskPath []string
	if updateMask["paths"] != nil {
		updateMaskPath = updateMask["paths"].([]string)
	}

	// var requestId map[string]interface{}
	// var requestIdType, requestIdValue string
	// if request["id"] != nil {
	// 	requestId = request["id"].(map[string]interface{})
	// 	if requestId["type"] != nil {
	// 		requestIdType = requestId["type"].(string)
	// 	}
	// 	if requestId["value"] != nil {
	// 		requestIdValue = requestId["value"].(string)
	// 	}
	// }
	return &nextgen.ClusterClusterUpdateRequest{
		Cluster:       buildClusterDetails(d),
		UpdatedFields: updatedFields,
		UpdateMask: &nextgen.ProtobufFieldMask{
			Paths: updateMaskPath,
		},
		// Id: &nextgen.ClustersClusterId{
		// 	Type_: requestIdType,
		// 	Value: requestIdValue,
		// },
	}
}

func buildClusterDetails(d *schema.ResourceData) *nextgen.Applicationv1alpha1Cluster {
	var clusterDetails nextgen.Applicationv1alpha1Cluster
	var request map[string]interface{}
	if attr, ok := d.GetOk("request"); ok {
		request = attr.([]interface{})[0].(map[string]interface{})
		if request["cluster"] != nil && len(request["cluster"].([]interface{})) > 0 {
			requestCluster := request["cluster"].([]interface{})[0].(map[string]interface{})

			if requestCluster["server"] != nil {
				clusterDetails.Server = requestCluster["server"].(string)
			}
			if requestCluster["name"] != nil {
				clusterDetails.Name = requestCluster["name"].(string)
			}

			if requestCluster["config"] != nil && len(requestCluster["config"].([]interface{})) > 0 {
				clusterConfig := requestCluster["config"].([]interface{})[0].(map[string]interface{})
				clusterDetails.Config = &nextgen.V1alpha1ClusterConfig{}
				if clusterConfig["username"] != nil {
					clusterDetails.Config.Username = clusterConfig["username"].(string)
				}
				if clusterConfig["password"] != nil {
					clusterDetails.Config.Password = clusterConfig["password"].(string)
				}
				if clusterConfig["bearer_token"] != nil {
					clusterDetails.Config.BearerToken = clusterConfig["bearer_token"].(string)
				}

				if clusterConfig["tls_client_config"] != nil && len(clusterConfig["tls_client_config"].([]interface{})) > 0 {
					clusterDetails.Config.TlsClientConfig = &nextgen.V1alpha1TlsClientConfig{}
					configTlsClientConfig := clusterConfig["tls_client_config"].([]interface{})[0].(map[string]interface{})
					if configTlsClientConfig["insecure"] != nil {
						clusterDetails.Config.TlsClientConfig.Insecure = configTlsClientConfig["insecure"].(bool)
					}
					if configTlsClientConfig["server_name"] != nil {
						clusterDetails.Config.TlsClientConfig.ServerName = configTlsClientConfig["server_name"].(string)
					}
					if configTlsClientConfig["cert_data"] != nil {
						clusterDetails.Config.TlsClientConfig.CertData = configTlsClientConfig["cert_data"].(string)
					}
					if configTlsClientConfig["key_data"] != nil {
						clusterDetails.Config.TlsClientConfig.KeyData = configTlsClientConfig["key_data"].(string)
					}
					if configTlsClientConfig["ca_data"] != nil {
						clusterDetails.Config.TlsClientConfig.CaData = configTlsClientConfig["ca_data"].(string)
					}
				}

				if clusterConfig["aws_auth_config"] != nil && len(clusterConfig["aws_auth_config"].([]interface{})) > 0 {
					clusterDetails.Config.AwsAuthConfig = &nextgen.V1alpha1AwsAuthConfig{}
					configAwsAuthConfig := clusterConfig["aws_auth_config"].([]interface{})[0].(map[string]interface{})
					if configAwsAuthConfig["cluster_name"] != nil {
						clusterDetails.Config.AwsAuthConfig.ClusterName = configAwsAuthConfig["cluster_name"].(string)
					}
					if configAwsAuthConfig["role_a_r_n"] != nil {
						clusterDetails.Config.AwsAuthConfig.RoleARN = configAwsAuthConfig["role_a_r_n"].(string)
					}
				}

				if clusterConfig["exec_provider_config"] != nil && len(clusterConfig["exec_provider_config"].([]interface{})) > 0 {
					clusterDetails.Config.ExecProviderConfig = &nextgen.V1alpha1ExecProviderConfig{}
					configExecProviderConfig := clusterConfig["exec_provider_config"].([]interface{})[0].(map[string]interface{})
					if configExecProviderConfig["command"] != nil {
						clusterDetails.Config.ExecProviderConfig.Command = configExecProviderConfig["command"].(string)
					}
					if configExecProviderConfig["args"] != nil {
						argsString := make([]string, len(configExecProviderConfig["args"].([]interface{})))
						for i, v := range configExecProviderConfig["args"].([]interface{}) {
							argsString[i] = v.(string)
						}
						clusterDetails.Config.ExecProviderConfig.Args = argsString
					}
					if configExecProviderConfig["env"] != nil {
						var envMap = map[string]string{}
						for k, v := range configExecProviderConfig["env"].(map[string]interface{}) {
							envMap[k] = v.(string)
						}
						clusterDetails.Config.ExecProviderConfig.Env = envMap
					}
					if configExecProviderConfig["api_version"] != nil {
						clusterDetails.Config.ExecProviderConfig.ApiVersion = configExecProviderConfig["api_version"].(string)
					}
					if configExecProviderConfig["install_hint"] != nil {
						clusterDetails.Config.ExecProviderConfig.InstallHint = configExecProviderConfig["install_hint"].(string)
					}
				}

				if clusterConfig["cluster_connection_type"] != nil {
					clusterDetails.Config.ClusterConnectionType = clusterConfig["cluster_connection_type"].(string)
				}
			}

			if requestCluster["namespaces"] != nil {
				namespaces := make([]string, len(requestCluster["namespaces"].([]interface{})))
				for i, v := range requestCluster["namespaces"].([]interface{}) {
					namespaces[i] = v.(string)
				}
				clusterDetails.Namespaces = namespaces
			}
			if requestCluster["refresh_requested_at"] != nil {
				clusterDetails.RefreshRequestedAt = &nextgen.V1Time{}
				refreshRequestedAt := requestCluster["refresh_requested_at"].([]interface{})[0].(map[string]interface{})
				if refreshRequestedAt["seconds"] != nil {
					clusterDetails.RefreshRequestedAt.Seconds = refreshRequestedAt["seconds"].(string)
				}
				if refreshRequestedAt["nanos"] != nil {
					clusterDetails.RefreshRequestedAt.Nanos = int32(refreshRequestedAt["nanos"].(int))
				}
			}

			if requestCluster["info"] != nil {
				clusterDetails.Info = &nextgen.V1alpha1ClusterInfo{}
				clusterInfo := requestCluster["info"].([]interface{})[0].(map[string]interface{})
				if clusterInfo["connection_state"] != nil {
					clusterDetails.Info.ConnectionState = &nextgen.V1alpha1ConnectionState{}
					connectionState := clusterInfo["connection_state"].([]interface{})[0].(map[string]interface{})
					if connectionState["status"] != nil {
						clusterDetails.Info.ConnectionState.Status = connectionState["status"].(string)
					}
					if connectionState["message"] != nil {
						clusterDetails.Info.ConnectionState.Message = connectionState["message"].(string)
					}
					if clusterInfo["attempted_at"] != nil {
						clusterDetails.Info.ConnectionState.AttemptedAt = &nextgen.V1Time{}
						attemptedAt := clusterInfo["attempted_at"].([]interface{})[0].(map[string]interface{})
						if attemptedAt["seconds"] != nil {
							clusterDetails.Info.ConnectionState.AttemptedAt.Seconds = attemptedAt["seconds"].(string)
						}
						if attemptedAt["nanos"] != nil {
							clusterDetails.Info.ConnectionState.AttemptedAt.Nanos = attemptedAt["nanos"].(int32)
						}
					}
					if clusterInfo["server_version"] != nil {
						clusterDetails.Info.ServerVersion = clusterInfo["server_version"].(string)
					}
				}
			}

			if requestCluster["shard"] != nil {
				clusterDetails.Shard = requestCluster["shard"].(string)
			}
			if requestCluster["cluster_resources"] != nil {
				clusterDetails.ClusterResources = requestCluster["cluster_resources"].(bool)
			}

			if requestCluster["labels"] != nil {
				var labelMap = map[string]string{}
				for k, v := range requestCluster["labels"].(map[string]interface{}) {
					labelMap[k] = v.(string)
				}
				clusterDetails.Labels = labelMap
			}
			if requestCluster["annotations"] != nil {
				var annotationMap = map[string]string{}
				for k, v := range requestCluster["annotations"].(map[string]interface{}) {
					annotationMap[k] = v.(string)
				}
				clusterDetails.Annotations = annotationMap
			}
		}
	}
	return &clusterDetails
}
