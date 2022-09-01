package provider

import (
	"github.com/auth0/go-auth0/management"
	"github.com/hashicorp/go-cty/cty"

	"github.com/auth0/terraform-provider-auth0/internal/value"
)

func flattenTenantChangePassword(changePassword *management.TenantChangePassword) []interface{} {
	if changePassword == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = changePassword.Enabled
	m["html"] = changePassword.HTML

	return []interface{}{m}
}

func flattenTenantGuardianMFAPage(mfa *management.TenantGuardianMFAPage) []interface{} {
	if mfa == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = mfa.Enabled
	m["html"] = mfa.HTML

	return []interface{}{m}
}

func flattenTenantErrorPage(errorPage *management.TenantErrorPage) []interface{} {
	if errorPage == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["html"] = errorPage.HTML
	m["show_log_link"] = errorPage.ShowLogLink
	m["url"] = errorPage.URL

	return []interface{}{m}
}

func flattenTenantFlags(flags *management.TenantFlags) []interface{} {
	if flags == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enable_client_connections"] = flags.EnableClientConnections
	m["enable_apis_section"] = flags.EnableAPIsSection
	m["enable_pipeline2"] = flags.EnablePipeline2
	m["enable_dynamic_client_registration"] = flags.EnableDynamicClientRegistration
	m["enable_custom_domain_in_emails"] = flags.EnableCustomDomainInEmails
	m["universal_login"] = flags.UniversalLogin
	m["enable_legacy_logs_search_v2"] = flags.EnableLegacyLogsSearchV2
	m["disable_clickjack_protection_headers"] = flags.DisableClickjackProtectionHeaders
	m["enable_public_signup_user_exists_error"] = flags.EnablePublicSignupUserExistsError
	m["use_scope_descriptions_for_consent"] = flags.UseScopeDescriptionsForConsent
	m["allow_legacy_delegation_grant_types"] = flags.AllowLegacyDelegationGrantTypes
	m["allow_legacy_ro_grant_types"] = flags.AllowLegacyROGrantTypes
	m["allow_legacy_tokeninfo_endpoint"] = flags.AllowLegacyTokenInfoEndpoint
	m["enable_legacy_profile"] = flags.EnableLegacyProfile
	m["enable_idtoken_api2"] = flags.EnableIDTokenAPI2
	m["no_disclose_enterprise_connections"] = flags.NoDisclosureEnterpriseConnections
	m["disable_management_api_sms_obfuscation"] = flags.DisableManagementAPISMSObfuscation
	m["enable_adfs_waad_email_verification"] = flags.EnableADFSWAADEmailVerification
	m["revoke_refresh_token_grant"] = flags.RevokeRefreshTokenGrant
	m["dashboard_log_streams_next"] = flags.DashboardLogStreams
	m["dashboard_insights_view"] = flags.DashboardInsightsView
	m["disable_fields_map_fix"] = flags.DisableFieldsMapFix

	return []interface{}{m}
}

func flattenTenantUniversalLogin(universalLogin *management.TenantUniversalLogin) []interface{} {
	if universalLogin == nil {
		return nil
	}
	if universalLogin.Colors == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["colors"] = []interface{}{
		map[string]interface{}{
			"primary":         universalLogin.Colors.Primary,
			"page_background": universalLogin.Colors.PageBackground,
		},
	}

	return []interface{}{m}
}

func flattenTenantSessionCookie(sessionCookie *management.TenantSessionCookie) []interface{} {
	m := make(map[string]interface{})
	m["mode"] = sessionCookie.GetMode()

	return []interface{}{m}
}

func expandTenantChangePassword(d ResourceData) *management.TenantChangePassword {
	var changePassword management.TenantChangePassword

	List(d, "change_password").Elem(func(d ResourceData) {
		changePassword.Enabled = Bool(d, "enabled")
		changePassword.HTML = String(d, "html")
	})

	return &changePassword
}

func expandTenantGuardianMFAPage(d ResourceData) *management.TenantGuardianMFAPage {
	var mfa management.TenantGuardianMFAPage

	List(d, "guardian_mfa_page").Elem(func(d ResourceData) {
		mfa.Enabled = Bool(d, "enabled")
		mfa.HTML = String(d, "html")
	})

	return &mfa
}

func expandTenantErrorPage(d ResourceData) *management.TenantErrorPage {
	var errorPage management.TenantErrorPage

	List(d, "error_page").Elem(func(d ResourceData) {
		errorPage.HTML = String(d, "html")
		errorPage.ShowLogLink = Bool(d, "show_log_link")
		errorPage.URL = String(d, "url")
	})

	return &errorPage
}

func expandTenantFlags(flagsList cty.Value) *management.TenantFlags {
	var tenantFlags *management.TenantFlags

	List(d, "flags").Elem(func(d ResourceData) {

	})
	// maybe we need check to ensure value is not nil before iterating on it?
	flagsList.ForEachElement(func(_ cty.Value, flags cty.Value) (stop bool) {
		tenantFlags = &management.TenantFlags{
			EnableClientConnections:            value.Bool(flags.GetAttr("enable_client_connections")),
			EnableAPIsSection:                  value.Bool(flags.GetAttr("enable_apis_section")),
			EnablePipeline2:                    value.Bool(flags.GetAttr("enable_pipeline2")),
			EnableDynamicClientRegistration:    value.Bool(flags.GetAttr("enable_dynamic_client_registration")),
			EnableCustomDomainInEmails:         value.Bool(flags.GetAttr("enable_custom_domain_in_emails")),
			UniversalLogin:                     value.Bool(flags.GetAttr("universal_login")),
			EnableLegacyLogsSearchV2:           value.Bool(flags.GetAttr("enable_legacy_logs_search_v2")),
			DisableClickjackProtectionHeaders:  value.Bool(flags.GetAttr("disable_clickjack_protection_headers")),
			EnablePublicSignupUserExistsError:  value.Bool(flags.GetAttr("enable_public_signup_user_exists_error")),
			UseScopeDescriptionsForConsent:     value.Bool(flags.GetAttr("use_scope_descriptions_for_consent")),
			AllowLegacyDelegationGrantTypes:    value.Bool(flags.GetAttr("allow_legacy_delegation_grant_types")),
			AllowLegacyROGrantTypes:            value.Bool(flags.GetAttr("allow_legacy_ro_grant_types")),
			AllowLegacyTokenInfoEndpoint:       value.Bool(flags.GetAttr("allow_legacy_tokeninfo_endpoint")),
			EnableLegacyProfile:                value.Bool(flags.GetAttr("enable_legacy_profile")),
			EnableIDTokenAPI2:                  value.Bool(flags.GetAttr("enable_idtoken_api2")),
			NoDisclosureEnterpriseConnections:  value.Bool(flags.GetAttr("no_disclose_enterprise_connections")),
			DisableManagementAPISMSObfuscation: value.Bool(flags.GetAttr("disable_management_api_sms_obfuscation")),
			EnableADFSWAADEmailVerification:    value.Bool(flags.GetAttr("enable_adfs_waad_email_verification")),
			RevokeRefreshTokenGrant:            value.Bool(flags.GetAttr("revoke_refresh_token_grant")),
			DashboardLogStreams:                value.Bool(flags.GetAttr("dashboard_log_streams_next")),
			DashboardInsightsView:              value.Bool(flags.GetAttr("dashboard_insights_view")),
			DisableFieldsMapFix:                value.Bool(flags.GetAttr("disable_fields_map_fix")),
		}

		return stop // return true to keep going, false to stop iteration
	})

	return tenantFlags
}

func expandTenantUniversalLogin(d ResourceData) *management.TenantUniversalLogin {
	var universalLogin management.TenantUniversalLogin

	List(d, "universal_login").Elem(func(d ResourceData) {
		List(d, "colors").Elem(func(d ResourceData) {
			universalLogin.Colors = &management.TenantUniversalLoginColors{
				Primary:        String(d, "primary"),
				PageBackground: String(d, "page_background"),
			}
		})
	})

	return &universalLogin
}

func expandTenantSessionCookie(sessionCookieList cty.Value) *management.TenantSessionCookie {
	var sessionCookie *management.TenantSessionCookie

	// maybe we need check to ensure value is not nil before iterating on it?
	// in case the API doesn't return it and also the config doesn't have the field set
	// we need to double-check this scenario and in other places where situation is similar
	sessionCookieList.ForEachElement(func(_ cty.Value, cookie cty.Value) (stop bool) {
		sessionCookie = &management.TenantSessionCookie{
			Mode: value.String(cookie.GetAttr("mode")),
		}

		return stop // return true to keep going, false to stop iteration
	})

	return sessionCookie
}
