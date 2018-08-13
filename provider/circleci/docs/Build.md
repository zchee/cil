# Build

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | commit message body | [optional] [default to null]
**Branch** | **string** |  | [optional] [default to null]
**BuildTimeMillis** | **int32** |  | [optional] [default to null]
**BuildUrl** | **string** |  | [optional] [default to null]
**CommitterEmail** | **string** |  | [optional] [default to null]
**CommitterName** | **string** |  | [optional] [default to null]
**DontBuild** | **string** | reason why we didn&#39;t build, if we didn&#39;t build | [optional] [default to null]
**Lifecycle** | **string** |  | [optional] [default to null]
**Previous** | [***PreviousBuild**](PreviousBuild.md) |  | [optional] [default to null]
**QueuedAt** | [**time.Time**](time.Time.md) | time build was queued | [optional] [default to null]
**Reponame** | **string** |  | [optional] [default to null]
**RetryOf** | **int32** | build_num of the build this is a retry of | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | time build started | [optional] [default to null]
**StopTime** | [**time.Time**](time.Time.md) | time build finished | [optional] [default to null]
**Subject** | **string** |  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**VcsUrl** | **string** |  | [optional] [default to null]
**Why** | **string** | short string explaining the reason we built | [optional] [default to null]
**Workflows** | [***Workflows**](Workflows.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


