// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ReceiverApplyConfiguration represents a declarative configuration of the Receiver type for use
// with apply.
type ReceiverApplyConfiguration struct {
	Name              *string                              `json:"name,omitempty"`
	OpsGenieConfigs   []OpsGenieConfigApplyConfiguration   `json:"opsgenieConfigs,omitempty"`
	PagerDutyConfigs  []PagerDutyConfigApplyConfiguration  `json:"pagerdutyConfigs,omitempty"`
	DiscordConfigs    []DiscordConfigApplyConfiguration    `json:"discordConfigs,omitempty"`
	SlackConfigs      []SlackConfigApplyConfiguration      `json:"slackConfigs,omitempty"`
	WebhookConfigs    []WebhookConfigApplyConfiguration    `json:"webhookConfigs,omitempty"`
	WeChatConfigs     []WeChatConfigApplyConfiguration     `json:"wechatConfigs,omitempty"`
	EmailConfigs      []EmailConfigApplyConfiguration      `json:"emailConfigs,omitempty"`
	VictorOpsConfigs  []VictorOpsConfigApplyConfiguration  `json:"victoropsConfigs,omitempty"`
	PushoverConfigs   []PushoverConfigApplyConfiguration   `json:"pushoverConfigs,omitempty"`
	SNSConfigs        []SNSConfigApplyConfiguration        `json:"snsConfigs,omitempty"`
	TelegramConfigs   []TelegramConfigApplyConfiguration   `json:"telegramConfigs,omitempty"`
	WebexConfigs      []WebexConfigApplyConfiguration      `json:"webexConfigs,omitempty"`
	MSTeamsConfigs    []MSTeamsConfigApplyConfiguration    `json:"msteamsConfigs,omitempty"`
	MSTeamsV2Configs  []MSTeamsV2ConfigApplyConfiguration  `json:"msteamsv2Configs,omitempty"`
	RocketChatConfigs []RocketChatConfigApplyConfiguration `json:"rocketchatConfigs,omitempty"`
}

// ReceiverApplyConfiguration constructs a declarative configuration of the Receiver type for use with
// apply.
func Receiver() *ReceiverApplyConfiguration {
	return &ReceiverApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ReceiverApplyConfiguration) WithName(value string) *ReceiverApplyConfiguration {
	b.Name = &value
	return b
}

// WithOpsGenieConfigs adds the given value to the OpsGenieConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OpsGenieConfigs field.
func (b *ReceiverApplyConfiguration) WithOpsGenieConfigs(values ...*OpsGenieConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOpsGenieConfigs")
		}
		b.OpsGenieConfigs = append(b.OpsGenieConfigs, *values[i])
	}
	return b
}

// WithPagerDutyConfigs adds the given value to the PagerDutyConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PagerDutyConfigs field.
func (b *ReceiverApplyConfiguration) WithPagerDutyConfigs(values ...*PagerDutyConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPagerDutyConfigs")
		}
		b.PagerDutyConfigs = append(b.PagerDutyConfigs, *values[i])
	}
	return b
}

// WithDiscordConfigs adds the given value to the DiscordConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DiscordConfigs field.
func (b *ReceiverApplyConfiguration) WithDiscordConfigs(values ...*DiscordConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDiscordConfigs")
		}
		b.DiscordConfigs = append(b.DiscordConfigs, *values[i])
	}
	return b
}

// WithSlackConfigs adds the given value to the SlackConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SlackConfigs field.
func (b *ReceiverApplyConfiguration) WithSlackConfigs(values ...*SlackConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSlackConfigs")
		}
		b.SlackConfigs = append(b.SlackConfigs, *values[i])
	}
	return b
}

// WithWebhookConfigs adds the given value to the WebhookConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WebhookConfigs field.
func (b *ReceiverApplyConfiguration) WithWebhookConfigs(values ...*WebhookConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithWebhookConfigs")
		}
		b.WebhookConfigs = append(b.WebhookConfigs, *values[i])
	}
	return b
}

// WithWeChatConfigs adds the given value to the WeChatConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WeChatConfigs field.
func (b *ReceiverApplyConfiguration) WithWeChatConfigs(values ...*WeChatConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithWeChatConfigs")
		}
		b.WeChatConfigs = append(b.WeChatConfigs, *values[i])
	}
	return b
}

// WithEmailConfigs adds the given value to the EmailConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EmailConfigs field.
func (b *ReceiverApplyConfiguration) WithEmailConfigs(values ...*EmailConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEmailConfigs")
		}
		b.EmailConfigs = append(b.EmailConfigs, *values[i])
	}
	return b
}

// WithVictorOpsConfigs adds the given value to the VictorOpsConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VictorOpsConfigs field.
func (b *ReceiverApplyConfiguration) WithVictorOpsConfigs(values ...*VictorOpsConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVictorOpsConfigs")
		}
		b.VictorOpsConfigs = append(b.VictorOpsConfigs, *values[i])
	}
	return b
}

// WithPushoverConfigs adds the given value to the PushoverConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PushoverConfigs field.
func (b *ReceiverApplyConfiguration) WithPushoverConfigs(values ...*PushoverConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPushoverConfigs")
		}
		b.PushoverConfigs = append(b.PushoverConfigs, *values[i])
	}
	return b
}

// WithSNSConfigs adds the given value to the SNSConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SNSConfigs field.
func (b *ReceiverApplyConfiguration) WithSNSConfigs(values ...*SNSConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSNSConfigs")
		}
		b.SNSConfigs = append(b.SNSConfigs, *values[i])
	}
	return b
}

// WithTelegramConfigs adds the given value to the TelegramConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TelegramConfigs field.
func (b *ReceiverApplyConfiguration) WithTelegramConfigs(values ...*TelegramConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTelegramConfigs")
		}
		b.TelegramConfigs = append(b.TelegramConfigs, *values[i])
	}
	return b
}

// WithWebexConfigs adds the given value to the WebexConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WebexConfigs field.
func (b *ReceiverApplyConfiguration) WithWebexConfigs(values ...*WebexConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithWebexConfigs")
		}
		b.WebexConfigs = append(b.WebexConfigs, *values[i])
	}
	return b
}

// WithMSTeamsConfigs adds the given value to the MSTeamsConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MSTeamsConfigs field.
func (b *ReceiverApplyConfiguration) WithMSTeamsConfigs(values ...*MSTeamsConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMSTeamsConfigs")
		}
		b.MSTeamsConfigs = append(b.MSTeamsConfigs, *values[i])
	}
	return b
}

// WithMSTeamsV2Configs adds the given value to the MSTeamsV2Configs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MSTeamsV2Configs field.
func (b *ReceiverApplyConfiguration) WithMSTeamsV2Configs(values ...*MSTeamsV2ConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMSTeamsV2Configs")
		}
		b.MSTeamsV2Configs = append(b.MSTeamsV2Configs, *values[i])
	}
	return b
}

// WithRocketChatConfigs adds the given value to the RocketChatConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RocketChatConfigs field.
func (b *ReceiverApplyConfiguration) WithRocketChatConfigs(values ...*RocketChatConfigApplyConfiguration) *ReceiverApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRocketChatConfigs")
		}
		b.RocketChatConfigs = append(b.RocketChatConfigs, *values[i])
	}
	return b
}
