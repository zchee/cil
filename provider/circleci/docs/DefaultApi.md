# \DefaultApi

All URIs are relative to *https://circleci.com/api/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Me**](DefaultApi.md#Me) | **Get** /me | 
[**ProjectUsernameProjectBuildCacheDelete**](DefaultApi.md#ProjectUsernameProjectBuildCacheDelete) | **Delete** /project/{username}/{project}/build-cache | 
[**ProjectUsernameProjectBuildNumArtifactsGet**](DefaultApi.md#ProjectUsernameProjectBuildNumArtifactsGet) | **Get** /project/{username}/{project}/{build_num}/artifacts | 
[**ProjectUsernameProjectBuildNumCancelPost**](DefaultApi.md#ProjectUsernameProjectBuildNumCancelPost) | **Post** /project/{username}/{project}/{build_num}/cancel | 
[**ProjectUsernameProjectBuildNumGet**](DefaultApi.md#ProjectUsernameProjectBuildNumGet) | **Get** /project/{username}/{project}/{build_num} | 
[**ProjectUsernameProjectBuildNumRetryPost**](DefaultApi.md#ProjectUsernameProjectBuildNumRetryPost) | **Post** /project/{username}/{project}/{build_num}/retry | 
[**ProjectUsernameProjectBuildNumTestsGet**](DefaultApi.md#ProjectUsernameProjectBuildNumTestsGet) | **Get** /project/{username}/{project}/{build_num}/tests | 
[**ProjectUsernameProjectCheckoutKeyFingerprintDelete**](DefaultApi.md#ProjectUsernameProjectCheckoutKeyFingerprintDelete) | **Delete** /project/{username}/{project}/checkout-key/{fingerprint} | 
[**ProjectUsernameProjectCheckoutKeyFingerprintGet**](DefaultApi.md#ProjectUsernameProjectCheckoutKeyFingerprintGet) | **Get** /project/{username}/{project}/checkout-key/{fingerprint} | 
[**ProjectUsernameProjectCheckoutKeyGet**](DefaultApi.md#ProjectUsernameProjectCheckoutKeyGet) | **Get** /project/{username}/{project}/checkout-key | 
[**ProjectUsernameProjectCheckoutKeyPost**](DefaultApi.md#ProjectUsernameProjectCheckoutKeyPost) | **Post** /project/{username}/{project}/checkout-key | 
[**ProjectUsernameProjectEnvvarGet**](DefaultApi.md#ProjectUsernameProjectEnvvarGet) | **Get** /project/{username}/{project}/envvar | 
[**ProjectUsernameProjectEnvvarNameDelete**](DefaultApi.md#ProjectUsernameProjectEnvvarNameDelete) | **Delete** /project/{username}/{project}/envvar/{name} | 
[**ProjectUsernameProjectEnvvarNameGet**](DefaultApi.md#ProjectUsernameProjectEnvvarNameGet) | **Get** /project/{username}/{project}/envvar/{name} | 
[**ProjectUsernameProjectEnvvarPost**](DefaultApi.md#ProjectUsernameProjectEnvvarPost) | **Post** /project/{username}/{project}/envvar | 
[**ProjectUsernameProjectGet**](DefaultApi.md#ProjectUsernameProjectGet) | **Get** /project/{username}/{project} | 
[**ProjectUsernameProjectPost**](DefaultApi.md#ProjectUsernameProjectPost) | **Post** /project/{username}/{project} | 
[**ProjectUsernameProjectSshKeyPost**](DefaultApi.md#ProjectUsernameProjectSshKeyPost) | **Post** /project/{username}/{project}/ssh-key | 
[**ProjectUsernameProjectTreeBranchPost**](DefaultApi.md#ProjectUsernameProjectTreeBranchPost) | **Post** /project/{username}/{project}/tree/{branch} | 
[**Projects**](DefaultApi.md#Projects) | **Get** /projects | 
[**RecentBuildsGet**](DefaultApi.md#RecentBuildsGet) | **Get** /recent-builds | 
[**UserHerokuKeyPost**](DefaultApi.md#UserHerokuKeyPost) | **Post** /user/heroku-key | 


# **Me**
> User Me(ctx, )


Provides information about the signed in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildCacheDelete**
> InlineResponse200 ProjectUsernameProjectBuildCacheDelete(ctx, username, project)


Clears the cache for a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildNumArtifactsGet**
> Artifacts ProjectUsernameProjectBuildNumArtifactsGet(ctx, username, project, buildNum)


List the artifacts produced by a given build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **buildNum** | **int32**| XXXXXXXXXX | 

### Return type

[**Artifacts**](Artifacts.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildNumCancelPost**
> Build ProjectUsernameProjectBuildNumCancelPost(ctx, username, project, buildNum)


Cancels the build, returns a summary of the build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **buildNum** | **int32**| XXXXXXXXXX | 

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildNumGet**
> BuildDetail ProjectUsernameProjectBuildNumGet(ctx, username, project, buildNum)


Full details for a single build. The response includes all of the fields from the build summary. This is also the payload for the [notification webhooks](/docs/configuration/#notify), in which case this object is the value to a key named 'payload'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **buildNum** | **int32**| XXXXXXXXXX | 

### Return type

[**BuildDetail**](BuildDetail.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildNumRetryPost**
> Build ProjectUsernameProjectBuildNumRetryPost(ctx, username, project, buildNum)


Retries the build, returns a summary of the new build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **buildNum** | **int32**| XXXXXXXXXX | 

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectBuildNumTestsGet**
> Tests ProjectUsernameProjectBuildNumTestsGet(ctx, username, project, buildNum)


Provides test metadata for a build Note: |   [Learn how to set up your builds to collect test metadata](https://circleci.com/docs/test-metadata/) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **buildNum** | **int32**| XXXXXXXXXX | 

### Return type

[**Tests**](Tests.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectCheckoutKeyFingerprintDelete**
> InlineResponse2001 ProjectUsernameProjectCheckoutKeyFingerprintDelete(ctx, username, project, fingerprint)


Delete a checkout key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **fingerprint** | **string**| XXXXXXXXXX | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectCheckoutKeyFingerprintGet**
> Key ProjectUsernameProjectCheckoutKeyFingerprintGet(ctx, username, project, fingerprint)


Get a checkout key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **fingerprint** | **string**| XXXXXXXXXX | 

### Return type

[**Key**](Key.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectCheckoutKeyGet**
> Keys ProjectUsernameProjectCheckoutKeyGet(ctx, username, project)


Lists checkout keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 

### Return type

[**Keys**](Keys.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectCheckoutKeyPost**
> Key ProjectUsernameProjectCheckoutKeyPost(ctx, username, project, optional)


Creates a new checkout key. Only usable with a user API token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| XXXXXXXXX | 
 **project** | **string**| XXXXXXXXX | 
 **type_** | **string**| The type of key to create. Can be &#39;deploy-key&#39; or &#39;github-user-key&#39;. | 

### Return type

[**Key**](Key.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectEnvvarGet**
> Envvars ProjectUsernameProjectEnvvarGet(ctx, username, project)


Lists the environment variables for :project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 

### Return type

[**Envvars**](Envvars.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectEnvvarNameDelete**
> InlineResponse2001 ProjectUsernameProjectEnvvarNameDelete(ctx, username, project, name)


Deletes the environment variable named ':name'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **name** | **string**| XXXXXXXXXX | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectEnvvarNameGet**
> Envvar ProjectUsernameProjectEnvvarNameGet(ctx, username, project, name)


Gets the hidden value of environment variable :name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **name** | **string**| XXXXXXXXXX | 

### Return type

[**Envvar**](Envvar.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectEnvvarPost**
> Envvar ProjectUsernameProjectEnvvarPost(ctx, username, project)


Creates a new environment variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 

### Return type

[**Envvar**](Envvar.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectGet**
> Builds ProjectUsernameProjectGet(ctx, username, project, optional)


Build summary for each of the last 30 builds for a single git repo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| XXXXXXXXX | 
 **project** | **string**| XXXXXXXXX | 
 **limit** | **int32**| The number of builds to return. Maximum 100, defaults to 30. | [default to 30]
 **offset** | **int32**| The API returns builds starting from this offset, defaults to 0. | [default to 0]
 **filter** | **string**| Restricts which builds are returned. Set to \&quot;completed\&quot;, \&quot;successful\&quot;, \&quot;failed\&quot;, \&quot;running\&quot;, or defaults to no filter.  | 

### Return type

[**Builds**](Builds.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectPost**
> BuildSummary ProjectUsernameProjectPost(ctx, username, project, optional)


Triggers a new build, returns a summary of the build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| XXXXXXXXX | 
 **project** | **string**| XXXXXXXXX | 
 **body** | [**Body**](Body.md)|  | 

### Return type

[**BuildSummary**](BuildSummary.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectSshKeyPost**
> InlineResponseDefault ProjectUsernameProjectSshKeyPost(ctx, username, project, contentType, body)


Create an ssh key used to access external systems that require SSH key-based authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **contentType** | **string**|  | 
  **body** | [**Body1**](Body1.md)|  | 

### Return type

[**InlineResponseDefault**](inline_response_default.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsernameProjectTreeBranchPost**
> Build ProjectUsernameProjectTreeBranchPost(ctx, username, project, branch, optional)


Triggers a new build, returns a summary of the build. Optional build parameters can be set using an experimental API.  Note: |   For more about build parameters,   read about [using parameterized builds](https://circleci.com/docs/parameterized-builds/) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**| XXXXXXXXX | 
  **project** | **string**| XXXXXXXXX | 
  **branch** | **string**| The branch name should be url-encoded. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| XXXXXXXXX | 
 **project** | **string**| XXXXXXXXX | 
 **branch** | **string**| The branch name should be url-encoded. | 
 **body** | [**Body2**](Body2.md)|  | 

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Projects**
> []Project Projects(ctx, )


List of all the projects you''re following on CircleCI, with build information organized by branch.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Project**](Project.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecentBuildsGet**
> Builds RecentBuildsGet(ctx, optional)


Build summary for each of the last 30 recent builds, ordered by build_num.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| The number of builds to return. Maximum 100, defaults to 30. | [default to 30]
 **offset** | **int32**| The API returns builds starting from this offset, defaults to 0. | [default to 0]

### Return type

[**Builds**](Builds.md)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserHerokuKeyPost**
> UserHerokuKeyPost(ctx, )


Adds your Heroku API key to CircleCI, takes apikey as form param name.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

