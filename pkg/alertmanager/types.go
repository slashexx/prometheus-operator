// Copyright 2020 The prometheus-operator Authors
//
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

package alertmanager

import (
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/common/model"
)

// Customization of Config type from alertmanager repo:
// https://github.com/prometheus/alertmanager/blob/main/config/config.go
//
// Custom global type to get around obfuscation of secret values when
// marshalling. See the following issue for details:
// https://github.com/prometheus/alertmanager/issues/1985
type alertmanagerConfig struct {
	Global            *globalConfig   `yaml:"global,omitempty" json:"global,omitempty"`
	Route             *route          `yaml:"route,omitempty" json:"route,omitempty"`
	InhibitRules      []*inhibitRule  `yaml:"inhibit_rules,omitempty" json:"inhibit_rules,omitempty"`
	Receivers         []*receiver     `yaml:"receivers,omitempty" json:"receivers,omitempty"`
	MuteTimeIntervals []*timeInterval `yaml:"mute_time_intervals,omitempty" json:"mute_time_intervals,omitempty"`
	TimeIntervals     []*timeInterval `yaml:"time_intervals,omitempty" json:"time_intervals,omitempty"`
	Templates         []string        `yaml:"templates" json:"templates"`
}

type globalConfig struct {
	// ResolveTimeout is the time after which an alert is declared resolved
	// if it has not been updated.
	ResolveTimeout *model.Duration `yaml:"resolve_timeout,omitempty" json:"resolve_timeout,omitempty"`

	HTTPConfig *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`

	SMTPFrom             string          `yaml:"smtp_from,omitempty" json:"smtp_from,omitempty"`
	SMTPHello            string          `yaml:"smtp_hello,omitempty" json:"smtp_hello,omitempty"`
	SMTPSmarthost        config.HostPort `yaml:"smtp_smarthost,omitempty" json:"smtp_smarthost,omitempty"`
	SMTPAuthUsername     string          `yaml:"smtp_auth_username,omitempty" json:"smtp_auth_username,omitempty"`
	SMTPAuthPassword     string          `yaml:"smtp_auth_password,omitempty" json:"smtp_auth_password,omitempty"`
	SMTPAuthPasswordFile string          `yaml:"smtp_auth_password_file,omitempty" json:"smtp_auth_password_file,omitempty"`
	SMTPAuthSecret       string          `yaml:"smtp_auth_secret,omitempty" json:"smtp_auth_secret,omitempty"`
	SMTPAuthIdentity     string          `yaml:"smtp_auth_identity,omitempty" json:"smtp_auth_identity,omitempty"`
	SMTPRequireTLS       *bool           `yaml:"smtp_require_tls,omitempty" json:"smtp_require_tls,omitempty"`
	SMTPTLSConfig        *tlsConfig      `yaml:"smtp_tls_config,omitempty" json:"smtp_tls_config,omitempty"`
	SlackAPIURL          *config.URL     `yaml:"slack_api_url,omitempty" json:"slack_api_url,omitempty"`
	SlackAPIURLFile      string          `yaml:"slack_api_url_file,omitempty" json:"slack_api_url_file,omitempty"`
	PagerdutyURL         *config.URL     `yaml:"pagerduty_url,omitempty" json:"pagerduty_url,omitempty"`
	HipchatAPIURL        *config.URL     `yaml:"hipchat_api_url,omitempty" json:"hipchat_api_url,omitempty"`
	HipchatAuthToken     string          `yaml:"hipchat_auth_token,omitempty" json:"hipchat_auth_token,omitempty"`
	OpsGenieAPIURL       *config.URL     `yaml:"opsgenie_api_url,omitempty" json:"opsgenie_api_url,omitempty"`
	OpsGenieAPIKey       string          `yaml:"opsgenie_api_key,omitempty" json:"opsgenie_api_key,omitempty"`
	OpsGenieAPIKeyFile   string          `yaml:"opsgenie_api_key_file,omitempty" json:"opsgenie_api_key_file,omitempty"`
	WeChatAPIURL         *config.URL     `yaml:"wechat_api_url,omitempty" json:"wechat_api_url,omitempty"`
	WeChatAPISecret      string          `yaml:"wechat_api_secret,omitempty" json:"wechat_api_secret,omitempty"`
	WeChatAPICorpID      string          `yaml:"wechat_api_corp_id,omitempty" json:"wechat_api_corp_id,omitempty"`
	VictorOpsAPIURL      *config.URL     `yaml:"victorops_api_url,omitempty" json:"victorops_api_url,omitempty"`
	VictorOpsAPIKey      string          `yaml:"victorops_api_key,omitempty" json:"victorops_api_key,omitempty"`
	VictorOpsAPIKeyFile  string          `yaml:"victorops_api_key_file,omitempty" json:"victorops_api_key_file,omitempty"`
	TelegramAPIURL       *config.URL     `yaml:"telegram_api_url,omitempty" json:"telegram_api_url,omitempty"`
	WebexAPIURL          *config.URL     `yaml:"webex_api_url,omitempty" json:"webex_api_url,omitempty"`
	JiraAPIURL           *config.URL     `yaml:"jira_api_url,omitempty" json:"jira_api_url,omitempty"`
	RocketChatAPIURL     *config.URL     `yaml:"rocketchat_api_url,omitempty" json:"rocketchat_api_url,omitempty"`
	RocketChatToken      string          `yaml:"rocketchat_token,omitempty" json:"rocketchat_token,omitempty"`
	RocketChatTokenID    string          `yaml:"rocketchat_token_id,omitempty" json:"rocketchat_token_id,omitempty"`
}

type route struct {
	Receiver            string            `yaml:"receiver,omitempty" json:"receiver,omitempty"`
	GroupByStr          []string          `yaml:"group_by,omitempty" json:"group_by,omitempty"`
	Match               map[string]string `yaml:"match,omitempty" json:"match,omitempty"`
	MatchRE             map[string]string `yaml:"match_re,omitempty" json:"match_re,omitempty"`
	Matchers            []string          `yaml:"matchers,omitempty" json:"matchers,omitempty"`
	Continue            bool              `yaml:"continue,omitempty" json:"continue,omitempty"`
	Routes              []*route          `yaml:"routes,omitempty" json:"routes,omitempty"`
	GroupWait           string            `yaml:"group_wait,omitempty" json:"group_wait,omitempty"`
	GroupInterval       string            `yaml:"group_interval,omitempty" json:"group_interval,omitempty"`
	RepeatInterval      string            `yaml:"repeat_interval,omitempty" json:"repeat_interval,omitempty"`
	MuteTimeIntervals   []string          `yaml:"mute_time_intervals,omitempty" json:"mute_time_intervals,omitempty"`
	ActiveTimeIntervals []string          `yaml:"active_time_intervals,omitempty" json:"active_time_intervals,omitempty"`
}

type inhibitRule struct {
	TargetMatch    map[string]string `yaml:"target_match,omitempty" json:"target_match,omitempty"`
	TargetMatchRE  map[string]string `yaml:"target_match_re,omitempty" json:"target_match_re,omitempty"`
	TargetMatchers []string          `yaml:"target_matchers,omitempty" json:"target_matchers,omitempty"`
	SourceMatch    map[string]string `yaml:"source_match,omitempty" json:"source_match,omitempty"`
	SourceMatchRE  map[string]string `yaml:"source_match_re,omitempty" json:"source_match_re,omitempty"`
	SourceMatchers []string          `yaml:"source_matchers,omitempty" json:"source_matchers,omitempty"`
	Equal          []string          `yaml:"equal,omitempty" json:"equal,omitempty"`
}

type receiver struct {
	Name              string              `yaml:"name" json:"name"`
	OpsgenieConfigs   []*opsgenieConfig   `yaml:"opsgenie_configs,omitempty" json:"opsgenie_configs,omitempty"`
	PagerdutyConfigs  []*pagerdutyConfig  `yaml:"pagerduty_configs,omitempty" json:"pagerduty_configs,omitempty"`
	SlackConfigs      []*slackConfig      `yaml:"slack_configs,omitempty" json:"slack_configs,omitempty"`
	WebhookConfigs    []*webhookConfig    `yaml:"webhook_configs,omitempty" json:"webhook_configs,omitempty"`
	WeChatConfigs     []*weChatConfig     `yaml:"wechat_configs,omitempty" json:"wechat_config,omitempty"`
	EmailConfigs      []*emailConfig      `yaml:"email_configs,omitempty" json:"email_configs,omitempty"`
	PushoverConfigs   []*pushoverConfig   `yaml:"pushover_configs,omitempty" json:"pushover_configs,omitempty"`
	VictorOpsConfigs  []*victorOpsConfig  `yaml:"victorops_configs,omitempty" json:"victorops_configs,omitempty"`
	SNSConfigs        []*snsConfig        `yaml:"sns_configs,omitempty" json:"sns_configs,omitempty"`
	TelegramConfigs   []*telegramConfig   `yaml:"telegram_configs,omitempty" json:"telegram_configs,omitempty"`
	DiscordConfigs    []*discordConfig    `yaml:"discord_configs,omitempty"`
	WebexConfigs      []*webexConfig      `yaml:"webex_configs,omitempty"`
	MSTeamsConfigs    []*msTeamsConfig    `yaml:"msteams_configs,omitempty"`
	MSTeamsV2Configs  []*msTeamsV2Config  `yaml:"msteamsv2_configs,omitempty"`
	JiraConfigs       []*jiraConfig       `yaml:"jira_configs,omitempty"`
	RocketChatConfigs []*rocketChatConfig `yaml:"rocketchat_configs,omitempty"`
}

type webhookConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	URL           string            `yaml:"url,omitempty" json:"url,omitempty"`
	URLFile       string            `yaml:"url_file,omitempty" json:"url_file,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	MaxAlerts     int32             `yaml:"max_alerts,omitempty" json:"max_alerts,omitempty"`
	Timeout       *model.Duration   `yaml:"timeout,omitempty" json:"timeout,omitempty"`
}

type pagerdutyConfig struct {
	VSendResolved  *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig     *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	ServiceKey     string            `yaml:"service_key,omitempty" json:"service_key,omitempty"`
	ServiceKeyFile string            `yaml:"service_key_file,omitempty" json:"service_key_file,omitempty"`
	RoutingKey     string            `yaml:"routing_key,omitempty" json:"routing_key,omitempty"`
	RoutingKeyFile string            `yaml:"routing_key_file,omitempty" json:"routing_key_file,omitempty"`
	URL            string            `yaml:"url,omitempty" json:"url,omitempty"`
	Client         string            `yaml:"client,omitempty" json:"client,omitempty"`
	ClientURL      string            `yaml:"client_url,omitempty" json:"client_url,omitempty"`
	Description    string            `yaml:"description,omitempty" json:"description,omitempty"`
	Details        map[string]string `yaml:"details,omitempty" json:"details,omitempty"`
	Images         []pagerdutyImage  `yaml:"images,omitempty" json:"images,omitempty"`
	Links          []pagerdutyLink   `yaml:"links,omitempty" json:"links,omitempty"`
	Severity       string            `yaml:"severity,omitempty" json:"severity,omitempty"`
	Class          string            `yaml:"class,omitempty" json:"class,omitempty"`
	Component      string            `yaml:"component,omitempty" json:"component,omitempty"`
	Group          string            `yaml:"group,omitempty" json:"group,omitempty"`
	Source         string            `yaml:"source,omitempty" json:"source,omitempty"`
}

type opsgenieConfig struct {
	VSendResolved *bool               `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig   `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	APIKey        string              `yaml:"api_key,omitempty" json:"api_key,omitempty"`
	APIKeyFile    string              `yaml:"api_key_file,omitempty" json:"api_key_file,omitempty"`
	APIURL        string              `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	Message       string              `yaml:"message,omitempty" json:"message,omitempty"`
	Description   string              `yaml:"description,omitempty" json:"description,omitempty"`
	Source        string              `yaml:"source,omitempty" json:"source,omitempty"`
	Details       map[string]string   `yaml:"details,omitempty" json:"details,omitempty"`
	Responders    []opsgenieResponder `yaml:"responders,omitempty" json:"responders,omitempty"`
	Tags          string              `yaml:"tags,omitempty" json:"tags,omitempty"`
	Note          string              `yaml:"note,omitempty" json:"note,omitempty"`
	Priority      string              `yaml:"priority,omitempty" json:"priority,omitempty"`
	UpdateAlerts  *bool               `yaml:"update_alerts,omitempty" json:"update_alerts,omitempty"`
	Entity        string              `yaml:"entity,omitempty" json:"entity,omitempty"`
	Actions       string              `yaml:"actions,omitempty" json:"actions,omitempty"`
}

type weChatConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	APISecret     string            `yaml:"api_secret,omitempty" json:"api_secret,omitempty"`
	APIURL        string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	CorpID        string            `yaml:"corp_id,omitempty" json:"corp_id,omitempty"`
	AgentID       string            `yaml:"agent_id,omitempty" json:"agent_id,omitempty"`
	ToUser        string            `yaml:"to_user,omitempty" json:"to_user,omitempty"`
	ToParty       string            `yaml:"to_party,omitempty" json:"to_party,omitempty"`
	ToTag         string            `yaml:"to_tag,omitempty" json:"to_tag,omitempty"`
	Message       string            `yaml:"message,omitempty" json:"message,omitempty"`
	MessageType   string            `yaml:"message_type,omitempty" json:"message_type,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
}

type slackConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	APIURL        string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	APIURLFile    string            `yaml:"api_url_file,omitempty" json:"api_url_file,omitempty"`
	Channel       string            `yaml:"channel,omitempty" json:"channel,omitempty"`
	Username      string            `yaml:"username,omitempty" json:"username,omitempty"`
	Color         string            `yaml:"color,omitempty" json:"color,omitempty"`
	Title         string            `yaml:"title,omitempty" json:"title,omitempty"`
	TitleLink     string            `yaml:"title_link,omitempty" json:"title_link,omitempty"`
	Pretext       string            `yaml:"pretext,omitempty" json:"pretext,omitempty"`
	Text          string            `yaml:"text,omitempty" json:"text,omitempty"`
	Fields        []slackField      `yaml:"fields,omitempty" json:"fields,omitempty"`
	ShortFields   bool              `yaml:"short_fields,omitempty" json:"short_fields,omitempty"`
	Footer        string            `yaml:"footer,omitempty" json:"footer,omitempty"`
	Fallback      string            `yaml:"fallback,omitempty" json:"fallback,omitempty"`
	CallbackID    string            `yaml:"callback_id,omitempty" json:"callback_id,omitempty"`
	IconEmoji     string            `yaml:"icon_emoji,omitempty" json:"icon_emoji,omitempty"`
	IconURL       string            `yaml:"icon_url,omitempty" json:"icon_url,omitempty"`
	ImageURL      string            `yaml:"image_url,omitempty" json:"image_url,omitempty"`
	ThumbURL      string            `yaml:"thumb_url,omitempty" json:"thumb_url,omitempty"`
	LinkNames     bool              `yaml:"link_names,omitempty" json:"link_names,omitempty"`
	MrkdwnIn      []string          `yaml:"mrkdwn_in,omitempty" json:"mrkdwn_in,omitempty"`
	Actions       []slackAction     `yaml:"actions,omitempty" json:"actions,omitempty"`
}

type httpClientConfig struct {
	Authorization   *authorization `yaml:"authorization,omitempty"`
	BasicAuth       *basicAuth     `yaml:"basic_auth,omitempty"`
	OAuth2          *oauth2        `yaml:"oauth2,omitempty"`
	BearerToken     string         `yaml:"bearer_token,omitempty"`
	BearerTokenFile string         `yaml:"bearer_token_file,omitempty"`
	TLSConfig       *tlsConfig     `yaml:"tls_config,omitempty"`
	FollowRedirects *bool          `yaml:"follow_redirects,omitempty"`
	EnableHTTP2     *bool          `yaml:"enable_http2,omitempty"`

	proxyConfig `yaml:",inline"`
}

type proxyConfig struct {
	ProxyURL             string              `yaml:"proxy_url,omitempty"`
	NoProxy              string              `yaml:"no_proxy,omitempty"`
	ProxyFromEnvironment bool                `yaml:"proxy_from_environment,omitempty"`
	ProxyConnectHeader   map[string][]string `yaml:"proxy_connect_header,omitempty"`
}

type tlsConfig struct {
	CAFile             string `yaml:"ca_file,omitempty"`
	CertFile           string `yaml:"cert_file,omitempty"`
	KeyFile            string `yaml:"key_file,omitempty"`
	ServerName         string `yaml:"server_name,omitempty"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify"`
	MinVersion         string `yaml:"min_version,omitempty"`
	MaxVersion         string `yaml:"max_version,omitempty"`
}

type authorization struct {
	Type            string `yaml:"type,omitempty"`
	Credentials     string `yaml:"credentials,omitempty"`
	CredentialsFile string `yaml:"credentials_file,omitempty"`
}

type basicAuth struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password,omitempty"`
	PasswordFile string `yaml:"password_file,omitempty"`
}

type oauth2 struct {
	ClientID         string            `yaml:"client_id"`
	ClientSecret     string            `yaml:"client_secret"`
	ClientSecretFile string            `yaml:"client_secret_file,omitempty"`
	Scopes           []string          `yaml:"scopes,omitempty"`
	TokenURL         string            `yaml:"token_url"`
	EndpointParams   map[string]string `yaml:"endpoint_params,omitempty"`
	proxyConfig      `yaml:",inline"`

	TLSConfig *tlsConfig `yaml:"tls_config,omitempty"`
}

type pagerdutyLink struct {
	Href string `yaml:"href,omitempty" json:"href,omitempty"`
	Text string `yaml:"text,omitempty" json:"text,omitempty"`
}

type pagerdutyImage struct {
	Src  string `yaml:"src,omitempty" json:"src,omitempty"`
	Alt  string `yaml:"alt,omitempty" json:"alt,omitempty"`
	Href string `yaml:"href,omitempty" json:"href,omitempty"`
}

type opsgenieResponder struct {
	ID       string `yaml:"id,omitempty" json:"id,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	Username string `yaml:"username,omitempty" json:"username,omitempty"`
	Type     string `yaml:"type,omitempty" json:"type,omitempty"`
}

type slackField struct {
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
	Short bool   `yaml:"short,omitempty" json:"short,omitempty"`
}

type slackAction struct {
	Type         string                  `yaml:"type,omitempty"  json:"type,omitempty"`
	Text         string                  `yaml:"text,omitempty"  json:"text,omitempty"`
	URL          string                  `yaml:"url,omitempty"   json:"url,omitempty"`
	Style        string                  `yaml:"style,omitempty" json:"style,omitempty"`
	Name         string                  `yaml:"name,omitempty"  json:"name,omitempty"`
	Value        string                  `yaml:"value,omitempty"  json:"value,omitempty"`
	ConfirmField *slackConfirmationField `yaml:"confirm,omitempty"  json:"confirm,omitempty"`
}

type slackConfirmationField struct {
	Text        string `yaml:"text,omitempty"  json:"text,omitempty"`
	Title       string `yaml:"title,omitempty"  json:"title,omitempty"`
	OkText      string `yaml:"ok_text,omitempty"  json:"ok_text,omitempty"`
	DismissText string `yaml:"dismiss_text,omitempty"  json:"dismiss_text,omitempty"`
}

type emailConfig struct {
	VSendResolved    *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	To               string            `yaml:"to,omitempty" json:"to,omitempty"`
	From             string            `yaml:"from,omitempty" json:"from,omitempty"`
	Hello            string            `yaml:"hello,omitempty" json:"hello,omitempty"`
	Smarthost        config.HostPort   `yaml:"smarthost,omitempty" json:"smarthost,omitempty"`
	AuthUsername     string            `yaml:"auth_username,omitempty" json:"auth_username,omitempty"`
	AuthPassword     string            `yaml:"auth_password,omitempty" json:"auth_password,omitempty"`
	AuthPasswordFile string            `yaml:"auth_password_file,omitempty" json:"auth_password_file,omitempty"`
	AuthSecret       string            `yaml:"auth_secret,omitempty" json:"auth_secret,omitempty"`
	AuthIdentity     string            `yaml:"auth_identity,omitempty" json:"auth_identity,omitempty"`
	Headers          map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	HTML             *string           `yaml:"html,omitempty" json:"html,omitempty"`
	Text             *string           `yaml:"text,omitempty" json:"text,omitempty"`
	RequireTLS       *bool             `yaml:"require_tls,omitempty" json:"require_tls,omitempty"`
	TLSConfig        *tlsConfig        `yaml:"tls_config,omitempty" json:"tls_config,omitempty"`
}

type pushoverConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	UserKey       string            `yaml:"user_key,omitempty" json:"user_key,omitempty"`
	UserKeyFile   string            `yaml:"user_key_file,omitempty" json:"user_key_file,omitempty"`
	Token         string            `yaml:"token,omitempty" json:"token,omitempty"`
	TokenFile     string            `yaml:"token_file,omitempty" json:"token_file,omitempty"`
	Title         string            `yaml:"title,omitempty" json:"title,omitempty"`
	Message       string            `yaml:"message,omitempty" json:"message,omitempty"`
	URL           string            `yaml:"url,omitempty" json:"url,omitempty"`
	URLTitle      string            `yaml:"url_title,omitempty" json:"url_title,omitempty"`
	TTL           string            `yaml:"ttl,omitempty" json:"ttl,omitempty"`
	Device        string            `yaml:"device,omitempty" json:"device,omitempty"`
	Sound         string            `yaml:"sound,omitempty" json:"sound,omitempty"`
	Priority      string            `yaml:"priority,omitempty" json:"priority,omitempty"`
	Retry         *model.Duration   `yaml:"retry,omitempty" json:"retry,omitempty"`
	Expire        *model.Duration   `yaml:"expire,omitempty" json:"expire,omitempty"`
	HTML          bool              `yaml:"html,omitempty" json:"html,omitempty"`
}

type snsConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	APIUrl        string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	Sigv4         sigV4Config       `yaml:"sigv4,omitempty" json:"sigv4,omitempty"`
	TopicARN      string            `yaml:"topic_arn,omitempty" json:"topic_arn,omitempty"`
	PhoneNumber   string            `yaml:"phone_number,omitempty" json:"phone_number,omitempty"`
	TargetARN     string            `yaml:"target_arn,omitempty" json:"target_arn,omitempty"`
	Subject       string            `yaml:"subject,omitempty" json:"subject,omitempty"`
	Message       string            `yaml:"message,omitempty" json:"message,omitempty"`
	Attributes    map[string]string `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type telegramConfig struct {
	VSendResolved        *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	APIUrl               string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	BotToken             string            `yaml:"bot_token,omitempty" json:"bot_token,omitempty"`
	BotTokenFile         string            `yaml:"bot_token_file,omitempty" json:"bot_token_file,omitempty"`
	ChatID               int64             `yaml:"chat_id,omitempty" json:"chat_id,omitempty"`
	MessageThreadID      int               `yaml:"message_thread_id,omitempty" json:"message_thread_id,omitempty"`
	Message              string            `yaml:"message,omitempty" json:"message,omitempty"`
	DisableNotifications bool              `yaml:"disable_notifications,omitempty" json:"disable_notifications,omitempty"`
	ParseMode            string            `yaml:"parse_mode,omitempty" json:"parse_mode,omitempty"`
	HTTPConfig           *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
}

type discordConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty"`
	WebhookURL    string            `yaml:"webhook_url,omitempty"`
	Title         string            `yaml:"title,omitempty"`
	Message       string            `yaml:"message,omitempty"`
	Content       string            `yaml:"content,omitempty"`
	Username      string            `yaml:"username,omitempty"`
	AvatarURL     string            `yaml:"avatar_url,omitempty"`
}

type webexConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty"`
	APIURL        string            `yaml:"api_url,omitempty"`
	Message       string            `yaml:"message,omitempty"`
	RoomID        string            `yaml:"room_id"`
}

type sigV4Config struct {
	Region    string `yaml:"region,omitempty" json:"region,omitempty"`
	AccessKey string `yaml:"access_key,omitempty" json:"access_key,omitempty"`
	SecretKey string `yaml:"secret_key,omitempty" json:"secret_key,omitempty"`
	Profile   string `yaml:"profile,omitempty" json:"profile,omitempty"`
	RoleARN   string `yaml:"role_arn,omitempty" json:"role_arn,omitempty"`
}

type victorOpsConfig struct {
	VSendResolved     *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig        *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	APIKey            string            `yaml:"api_key,omitempty" json:"api_key,omitempty"`
	APIKeyFile        string            `yaml:"api_key_file,omitempty" json:"api_key_file,omitempty"`
	APIURL            string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	RoutingKey        string            `yaml:"routing_key,omitempty" json:"routing_key,omitempty"`
	MessageType       string            `yaml:"message_type,omitempty" json:"message_type,omitempty"`
	StateMessage      string            `yaml:"state_message,omitempty" json:"state_message,omitempty"`
	EntityDisplayName string            `yaml:"entity_display_name,omitempty" json:"entity_display_name,omitempty"`
	MonitoringTool    string            `yaml:"monitoring_tool,omitempty" json:"monitoring_tool,omitempty"`
	CustomFields      map[string]string `yaml:"custom_fields,omitempty" json:"custom_fields,omitempty"`
}

type msTeamsConfig struct {
	SendResolved *bool             `yaml:"send_resolved,omitempty"`
	WebhookURL   string            `yaml:"webhook_url"`
	Title        string            `yaml:"title,omitempty"`
	Summary      string            `yaml:"summary,omitempty"`
	Text         string            `yaml:"text,omitempty"`
	HTTPConfig   *httpClientConfig `yaml:"http_config,omitempty"`
}

type msTeamsV2Config struct {
	SendResolved   *bool             `yaml:"send_resolved,omitempty"`
	WebhookURL     string            `yaml:"webhook_url,omitempty"`
	WebhookURLFile string            `yaml:"webhook_url_file,omitempty"`
	Title          string            `yaml:"title,omitempty"`
	Text           string            `yaml:"text,omitempty"`
	HTTPConfig     *httpClientConfig `yaml:"http_config,omitempty"`
}

type jiraConfig struct {
	HTTPConfig        *httpClientConfig `yaml:"http_config,omitempty"`
	APIURL            string            `yaml:"api_url,omitempty"`
	Project           string            `yaml:"project,omitempty"`
	Summary           string            `yaml:"summary,omitempty"`
	Description       string            `yaml:"description,omitempty"`
	Labels            []string          `yaml:"labels,omitempty"`
	Priority          string            `yaml:"priority,omitempty"`
	IssueType         string            `yaml:"issue_type,omitempty"`
	ReopenTransition  string            `yaml:"reopen_transition,omitempty"`
	ResolveTransition string            `yaml:"resolve_transition,omitempty"`
	WontFixResolution string            `yaml:"wont_fix_resolution,omitempty"`
	ReopenDuration    model.Duration    `yaml:"reopen_duration,omitempty"`
	Fields            map[string]any    `yaml:"fields,omitempty"`
}

type rocketchatAttachmentField struct {
	Short *bool  `yaml:"short"`
	Title string `yaml:"title,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type rocketchatAttachmentAction struct {
	Type               string `yaml:"type,omitempty"`
	Text               string `yaml:"text,omitempty"`
	URL                string `yaml:"url,omitempty"`
	ImageURL           string `yaml:"image_url,omitempty"`
	IsWebView          bool   `yaml:"is_webview"`
	WebviewHeightRatio string `yaml:"webview_height_ratio,omitempty"`
	Msg                string `yaml:"msg,omitempty"`
	MsgInChatWindow    bool   `yaml:"msg_in_chat_window"`
	MsgProcessingType  string `yaml:"msg_processing_type,omitempty"`
}

type rocketChatConfig struct {
	VSendResolved *bool             `yaml:"send_resolved,omitempty" json:"send_resolved,omitempty"`
	HTTPConfig    *httpClientConfig `yaml:"http_config,omitempty" json:"http_config,omitempty"`
	APIURL        string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	TokenID       *string           `yaml:"token_id,omitempty" json:"token_id,omitempty"`
	Token         *string           `yaml:"token,omitempty" json:"token,omitempty"`
	// RocketChat channel override, (like #other-channel or @username).
	Channel     string                        `yaml:"channel,omitempty" json:"channel,omitempty"`
	Color       string                        `yaml:"color,omitempty" json:"color,omitempty"`
	Title       string                        `yaml:"title,omitempty" json:"title,omitempty"`
	TitleLink   string                        `yaml:"title_link,omitempty" json:"title_link,omitempty"`
	Text        string                        `yaml:"text,omitempty" json:"text,omitempty"`
	Fields      []*rocketchatAttachmentField  `yaml:"fields,omitempty" json:"fields,omitempty"`
	ShortFields bool                          `yaml:"short_fields" json:"short_fields"`
	Emoji       string                        `yaml:"emoji,omitempty" json:"emoji,omitempty"`
	IconURL     string                        `yaml:"icon_url,omitempty" json:"icon_url,omitempty"`
	ImageURL    string                        `yaml:"image_url,omitempty" json:"image_url,omitempty"`
	ThumbURL    string                        `yaml:"thumb_url,omitempty" json:"thumb_url,omitempty"`
	LinkNames   bool                          `yaml:"link_names" json:"link_names"`
	Actions     []*rocketchatAttachmentAction `yaml:"actions,omitempty" json:"actions,omitempty"`
}

type timeInterval config.TimeInterval
