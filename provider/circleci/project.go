/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI. 
 *
 * API version: v1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package circleci

type Project struct {

	Aws *Aws `json:"aws,omitempty"`

	Branches []Branch `json:"branches,omitempty"`

	CampfireNotifyPrefs *StringOrNull `json:"campfire_notify_prefs,omitempty"`

	CampfireRoom *StringOrNull `json:"campfire_room,omitempty"`

	CampfireSubdomain *StringOrNull `json:"campfire_subdomain,omitempty"`

	CampfireToken *StringOrNull `json:"campfire_token,omitempty"`

	Compile string `json:"compile,omitempty"`

	DefaultBranch string `json:"default_branch,omitempty"`

	Dependencies string `json:"dependencies,omitempty"`

	Extra string `json:"extra,omitempty"`

	FeatureFlags *ProjectFeatureFlags `json:"feature_flags,omitempty"`

	FlowdockApiToken *StringOrNull `json:"flowdock_api_token,omitempty"`

	Followed bool `json:"followed,omitempty"`

	HasUsableKey bool `json:"has_usable_key,omitempty"`

	HerokuDeployUser *StringOrNull `json:"heroku_deploy_user,omitempty"`

	HipchatApiToken *StringOrNull `json:"hipchat_api_token,omitempty"`

	HipchatNotify *StringOrNull `json:"hipchat_notify,omitempty"`

	HipchatNotifyPrefs string `json:"hipchat_notify_prefs,omitempty"`

	HipchatRoom *StringOrNull `json:"hipchat_room,omitempty"`

	IrcChannel *StringOrNull `json:"irc_channel,omitempty"`

	IrcKeyword *StringOrNull `json:"irc_keyword,omitempty"`

	IrcNotifyPrefs *StringOrNull `json:"irc_notify_prefs,omitempty"`

	IrcPassword *StringOrNull `json:"irc_password,omitempty"`

	IrcServer *StringOrNull `json:"irc_server,omitempty"`

	IrcUsername *StringOrNull `json:"irc_username,omitempty"`

	Language string `json:"language,omitempty"`

	Oss bool `json:"oss,omitempty"`

	Parallel int32 `json:"parallel,omitempty"`

	Reponame string `json:"reponame,omitempty"`

	Scopes []string `json:"scopes,omitempty"`

	Setup string `json:"setup,omitempty"`

	SlackApiToken *StringOrNull `json:"slack_api_token,omitempty"`

	SlackChannel *StringOrNull `json:"slack_channel,omitempty"`

	SlackChannelOverride *StringOrNull `json:"slack_channel_override,omitempty"`

	SlackNotifyPrefs *StringOrNull `json:"slack_notify_prefs,omitempty"`

	SlackSubdomain *StringOrNull `json:"slack_subdomain,omitempty"`

	SlackWebhookUrl string `json:"slack_webhook_url,omitempty"`

	SshKeys []string `json:"ssh_keys,omitempty"`

	Test string `json:"test,omitempty"`

	Username string `json:"username,omitempty"`

	VcsType string `json:"vcs_type,omitempty"`

	VcsUrl string `json:"vcs_url,omitempty"`
}