// Copyright 2021 The prometheus-operator Authors
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

package v1alpha1

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/prometheus-operator/prometheus-operator/pkg/alertmanager/validation"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
)

var durationRe = regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`)

// ValidateAlertmanagerConfig checks that the given resource complies with the
// semantics of the Alertmanager configuration.
// In particular, it verifies things that can't be modelized with the OpenAPI
// specification such as routes should refer to an existing receiver.
func ValidateAlertmanagerConfig(amc *monitoringv1alpha1.AlertmanagerConfig) error {
	receivers, err := validateReceivers(amc.Spec.Receivers)
	if err != nil {
		return err
	}

	muteTimeIntervals, err := validateMuteTimeIntervals(amc.Spec.MuteTimeIntervals)
	if err != nil {
		return err
	}

	return validateRoute(amc.Spec.Route, receivers, muteTimeIntervals, true)
}

func validateReceivers(receivers []monitoringv1alpha1.Receiver) (map[string]struct{}, error) {
	var err error
	receiverNames := make(map[string]struct{})

	for _, receiver := range receivers {
		if _, found := receiverNames[receiver.Name]; found {
			return nil, fmt.Errorf("%q receiver is not unique: %w", receiver.Name, err)
		}
		receiverNames[receiver.Name] = struct{}{}

		if err = validatePagerDutyConfigs(receiver.PagerDutyConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'pagerDutyConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateOpsGenieConfigs(receiver.OpsGenieConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'opsGenieConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateSlackConfigs(receiver.SlackConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'slackConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateWebhookConfigs(receiver.WebhookConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'webhookConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateWechatConfigs(receiver.WeChatConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'weChatConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateEmailConfig(receiver.EmailConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'emailConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateVictorOpsConfigs(receiver.VictorOpsConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'victorOpsConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validatePushoverConfigs(receiver.PushoverConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'pushOverConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateSnsConfigs(receiver.SNSConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'snsConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateTelegramConfigs(receiver.TelegramConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'telegramConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateWebexConfigs(receiver.WebexConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'webexConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateDiscordConfigs(receiver.DiscordConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'discordConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateMSTeamsConfigs(receiver.MSTeamsConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'msteamsConfig' - receiver %s: %w", receiver.Name, err)
		}

		if err := validateRocketchatConfigs(receiver.RocketChatConfigs); err != nil {
			return nil, fmt.Errorf("failed to validate 'rocketchatConfig' - reciever %s: %w", receiver.Name, err)
		}

	}

	return receiverNames, nil
}

func validatePagerDutyConfigs(configs []monitoringv1alpha1.PagerDutyConfig) error {
	for _, conf := range configs {
		if conf.URL != "" {
			if _, err := validation.ValidateURL(conf.URL); err != nil {
				return fmt.Errorf("pagerduty validation failed for 'url': %w", err)
			}
		}
		if conf.RoutingKey == nil && conf.ServiceKey == nil {
			return errors.New("one of 'routingKey' or 'serviceKey' is required")
		}

		if err := conf.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateOpsGenieConfigs(configs []monitoringv1alpha1.OpsGenieConfig) error {
	for _, config := range configs {
		if err := config.Validate(); err != nil {
			return err
		}
		if config.APIURL != "" {
			if _, err := validation.ValidateURL(config.APIURL); err != nil {
				return fmt.Errorf("invalid 'apiURL': %w", err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateDiscordConfigs(configs []monitoringv1alpha1.DiscordConfig) error {
	for _, config := range configs {
		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateRocketchatConfigs(configs []monitoringv1alpha1.RocketChatConfig) error {
	for _, config := range configs {
		if config.APIURL != nil && *config.APIURL != "" {
			if _, err := validation.ValidateURL(string(*config.APIURL)); err != nil {
				return fmt.Errorf("invalid 'apiURL': %w", err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateSlackConfigs(configs []monitoringv1alpha1.SlackConfig) error {
	for _, config := range configs {
		if err := config.Validate(); err != nil {
			return err
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateWebhookConfigs(configs []monitoringv1alpha1.WebhookConfig) error {
	for _, config := range configs {
		if config.URL == nil && config.URLSecret == nil {
			return errors.New("one of 'url' or 'urlSecret' must be specified")
		}
		if config.URL != nil {
			if _, err := validation.ValidateURL(*config.URL); err != nil {
				return fmt.Errorf("invalid 'url': %w", err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateWechatConfigs(configs []monitoringv1alpha1.WeChatConfig) error {
	for _, config := range configs {
		if config.APIURL != "" {
			if _, err := validation.ValidateURL(config.APIURL); err != nil {
				return fmt.Errorf("invalid 'apiURL': %w", err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateEmailConfig(configs []monitoringv1alpha1.EmailConfig) error {
	for _, config := range configs {
		if config.To == "" {
			return errors.New("missing 'to' address")
		}

		if config.Smarthost != "" {
			_, _, err := net.SplitHostPort(config.Smarthost)
			if err != nil {
				return fmt.Errorf("invalid 'smarthost' %s: %w", config.Smarthost, err)
			}
		}

		if config.Headers != nil {
			// Header names are case-insensitive, check for collisions.
			normalizedHeaders := map[string]struct{}{}
			for _, v := range config.Headers {
				normalized := strings.ToLower(v.Key)
				if _, ok := normalizedHeaders[normalized]; ok {
					return fmt.Errorf("duplicate header %q", normalized)
				}
				normalizedHeaders[normalized] = struct{}{}
			}
		}
	}
	return nil
}

func validateVictorOpsConfigs(configs []monitoringv1alpha1.VictorOpsConfig) error {
	for _, config := range configs {

		// from https://github.com/prometheus/alertmanager/blob/a7f9fdadbecbb7e692d2cd8d3334e3d6de1602e1/config/notifiers.go#L497
		reservedFields := map[string]struct{}{
			"routing_key":         {},
			"message_type":        {},
			"state_message":       {},
			"entity_display_name": {},
			"monitoring_tool":     {},
			"entity_id":           {},
			"entity_state":        {},
		}

		if len(config.CustomFields) > 0 {
			for _, v := range config.CustomFields {
				if _, ok := reservedFields[v.Key]; ok {
					return fmt.Errorf("usage of reserved word %q is not allowed in custom fields", v.Key)
				}
			}
		}

		if config.RoutingKey == "" {
			return errors.New("missing 'routingKey' key")
		}

		if config.APIURL != "" {
			if _, err := validation.ValidateURL(config.APIURL); err != nil {
				return fmt.Errorf("'apiURL' %s invalid: %w", config.APIURL, err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validatePushoverConfigs(configs []monitoringv1alpha1.PushoverConfig) error {
	for _, config := range configs {
		if config.UserKey == nil && config.UserKeyFile == nil {
			return fmt.Errorf("one of userKey or userKeyFile must be configured")
		}

		if config.Token == nil && config.TokenFile == nil {
			return fmt.Errorf("one of token or tokenFile must be configured")
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func validateSnsConfigs(configs []monitoringv1alpha1.SNSConfig) error {
	for _, config := range configs {
		if (config.TargetARN == "") != (config.TopicARN == "") != (config.PhoneNumber == "") {
			return fmt.Errorf("must provide either a Target ARN, Topic ARN, or Phone Number for SNS config")
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateTelegramConfigs(configs []monitoringv1alpha1.TelegramConfig) error {
	for _, config := range configs {

		if config.BotToken == nil && config.BotTokenFile == nil {
			return fmt.Errorf("mandatory field botToken or botTokenfile is empty")
		}

		if config.ChatID == 0 {
			return fmt.Errorf("mandatory field %q is empty", "chatID")
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func validateWebexConfigs(configs []monitoringv1alpha1.WebexConfig) error {
	for _, config := range configs {
		if config.APIURL != nil && *config.APIURL != "" {
			if _, err := validation.ValidateURL(string(*config.APIURL)); err != nil {
				return fmt.Errorf("invalid 'apiURL': %w", err)
			}
		}

		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func validateMSTeamsConfigs(configs []monitoringv1alpha1.MSTeamsConfig) error {
	for _, config := range configs {
		if err := config.HTTPConfig.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// validateRoute verifies that the given route and all its children are
// semantically valid.  because of the self-referential issues mentioned in
// https://github.com/kubernetes/kubernetes/issues/62872 it is not currently
// possible to apply OpenAPI validation to a v1alpha1.Route.
func validateRoute(r *monitoringv1alpha1.Route, receivers, muteTimeIntervals map[string]struct{}, topLevelRoute bool) error {
	if r == nil {
		return nil
	}

	if r.Receiver == "" {
		if topLevelRoute {
			return errors.New("root route must define a receiver")
		}
	} else {
		if _, found := receivers[r.Receiver]; !found {
			return fmt.Errorf("receiver %q not found", r.Receiver)
		}
	}

	if groupLen := len(r.GroupBy); groupLen > 0 {
		groupedBy := make(map[string]struct{}, groupLen)
		for _, str := range r.GroupBy {
			if _, found := groupedBy[str]; found {
				return fmt.Errorf("duplicate values not permitted in route 'groupBy': %v", r.GroupBy)
			}
			groupedBy[str] = struct{}{}
		}
		if _, found := groupedBy["..."]; found && groupLen > 1 {
			return fmt.Errorf("'...' must be a sole value in route 'groupBy': %v", r.GroupBy)
		}
	}

	for _, namedMuteTimeInterval := range r.MuteTimeIntervals {
		if _, found := muteTimeIntervals[namedMuteTimeInterval]; !found {
			return fmt.Errorf("mute time interval %q not found", namedMuteTimeInterval)
		}
	}

	for _, namedActiveTimeInterval := range r.ActiveTimeIntervals {
		if _, found := muteTimeIntervals[namedActiveTimeInterval]; !found {
			return fmt.Errorf("time interval %q not found", namedActiveTimeInterval)
		}
	}

	if r.GroupInterval != "" && !durationRe.MatchString(r.GroupInterval) {
		return fmt.Errorf("groupInterval %s does not match required regex: %s", r.GroupInterval, durationRe.String())

	}
	if r.GroupWait != "" && !durationRe.MatchString(r.GroupWait) {
		return fmt.Errorf("groupWait %s does not match required regex: %s", r.GroupWait, durationRe.String())
	}

	if r.RepeatInterval != "" && !durationRe.MatchString(r.RepeatInterval) {
		return fmt.Errorf("repeatInterval %s does not match required regex: %s", r.RepeatInterval, durationRe.String())
	}

	for i, m := range r.Matchers {
		if err := m.Validate(); err != nil {
			return fmt.Errorf("matcher[%d]: %w", i, err)
		}
	}

	// Unmarshal the child routes and validate them recursively.
	children, err := r.ChildRoutes()
	if err != nil {
		return err
	}

	for i := range children {
		if err := validateRoute(&children[i], receivers, muteTimeIntervals, false); err != nil {
			return fmt.Errorf("route[%d]: %w", i, err)
		}
	}

	return nil
}

func validateMuteTimeIntervals(muteTimeIntervals []monitoringv1alpha1.MuteTimeInterval) (map[string]struct{}, error) {
	muteTimeIntervalNames := make(map[string]struct{}, len(muteTimeIntervals))

	for i, mti := range muteTimeIntervals {
		if err := mti.Validate(); err != nil {
			return nil, fmt.Errorf("mute time interval[%d] is invalid: %w", i, err)
		}
		muteTimeIntervalNames[mti.Name] = struct{}{}
	}
	return muteTimeIntervalNames, nil
}
