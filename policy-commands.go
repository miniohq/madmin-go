//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

// InfoCannedPolicy - expand canned policy into JSON structure.
//
// Deprecated: Use InfoCannedPolicyV2 instead.
func (adm *AdminClient) InfoCannedPolicy(ctx context.Context, policyName string) ([]byte, error) {
	queryValues := url.Values{}
	queryValues.Set("name", policyName)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/info-canned-policy",
		queryValues: queryValues,
	}

	// Execute GET on /minio/admin/v3/info-canned-policy
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)

	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	return io.ReadAll(resp.Body)
}

// PolicyInfo contains information on a policy.
type PolicyInfo struct {
	PolicyName string
	Policy     json.RawMessage
	CreateDate time.Time `json:",omitempty"`
	UpdateDate time.Time `json:",omitempty"`
}

// MarshalJSON marshaller for JSON
func (pi PolicyInfo) MarshalJSON() ([]byte, error) {
	type aliasPolicyInfo PolicyInfo // needed to avoid recursive marshal
	if pi.CreateDate.IsZero() && pi.UpdateDate.IsZero() {
		return json.Marshal(&struct {
			PolicyName string
			Policy     json.RawMessage
		}{
			PolicyName: pi.PolicyName,
			Policy:     pi.Policy,
		})
	}
	return json.Marshal(aliasPolicyInfo(pi))
}

// InfoCannedPolicyV2 - get info on a policy including timestamps and policy json.
func (adm *AdminClient) InfoCannedPolicyV2(ctx context.Context, policyName string) (*PolicyInfo, error) {
	queryValues := url.Values{}
	queryValues.Set("name", policyName)
	queryValues.Set("v", "2")

	reqData := requestData{
		relPath:     adminAPIPrefix + "/info-canned-policy",
		queryValues: queryValues,
	}

	// Execute GET on /minio/admin/v3/info-canned-policy
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)

	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p PolicyInfo
	err = json.Unmarshal(data, &p)
	return &p, err
}

// ListCannedPolicies - list all configured canned policies.
func (adm *AdminClient) ListCannedPolicies(ctx context.Context) (map[string]json.RawMessage, error) {
	reqData := requestData{
		relPath: adminAPIPrefix + "/list-canned-policies",
	}

	// Execute GET on /minio/admin/v3/list-canned-policies
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)

	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	policies := make(map[string]json.RawMessage)
	if err = json.Unmarshal(respBytes, &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

// RemoveCannedPolicy - remove a policy for a canned.
func (adm *AdminClient) RemoveCannedPolicy(ctx context.Context, policyName string) error {
	queryValues := url.Values{}
	queryValues.Set("name", policyName)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/remove-canned-policy",
		queryValues: queryValues,
	}

	// Execute DELETE on /minio/admin/v3/remove-canned-policy to remove policy.
	resp, err := adm.executeMethod(ctx, http.MethodDelete, reqData)

	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}

	return nil
}

// AddCannedPolicy - adds a policy for a canned.
func (adm *AdminClient) AddCannedPolicy(ctx context.Context, policyName string, policy []byte) error {
	if len(policy) == 0 {
		return ErrInvalidArgument("policy input cannot be empty")
	}

	queryValues := url.Values{}
	queryValues.Set("name", policyName)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/add-canned-policy",
		queryValues: queryValues,
		content:     policy,
	}

	// Execute PUT on /minio/admin/v3/add-canned-policy to set policy.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)

	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}

	return nil
}

// SetPolicy - sets the policy for a user or a group.
//
// Deprecated: Use AttachPolicy/DetachPolicy to update builtin user policies
// instead. Use AttachPolicyLDAP/DetachPolicyLDAP to update LDAP user policies.
// This function and the corresponding server API will be removed in future
// releases.
func (adm *AdminClient) SetPolicy(ctx context.Context, policyName, entityName string, isGroup bool) error {
	queryValues := url.Values{}
	queryValues.Set("policyName", policyName)
	queryValues.Set("userOrGroup", entityName)
	groupStr := "false"
	if isGroup {
		groupStr = "true"
	}
	queryValues.Set("isGroup", groupStr)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/set-user-or-group-policy",
		queryValues: queryValues,
	}

	// Execute PUT on /minio/admin/v3/set-user-or-group-policy to set policy.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}
	return nil
}

func (adm *AdminClient) attachOrDetachPolicyBuiltin(ctx context.Context, isAttach bool,
	r PolicyAssociationReq,
) (PolicyAssociationResp, error) {
	err := r.IsValid()
	if err != nil {
		return PolicyAssociationResp{}, err
	}

	plainBytes, err := json.Marshal(r)
	if err != nil {
		return PolicyAssociationResp{}, err
	}

	encBytes, err := EncryptData(adm.getSecretKey(), plainBytes)
	if err != nil {
		return PolicyAssociationResp{}, err
	}

	suffix := "detach"
	if isAttach {
		suffix = "attach"
	}
	h := make(http.Header, 1)
	h.Add("Content-Type", "application/octet-stream")
	reqData := requestData{
		customHeaders: h,
		relPath:       adminAPIPrefix + "/idp/builtin/policy/" + suffix,
		content:       encBytes,
	}

	resp, err := adm.executeMethod(ctx, http.MethodPost, reqData)
	defer closeResponse(resp)
	if err != nil {
		return PolicyAssociationResp{}, err
	}

	// Older minio does not send a response, so we handle that case.

	switch {
	case resp.StatusCode == http.StatusOK:
		// Newer/current minio sends a result.
		content, err := DecryptData(adm.getSecretKey(), resp.Body)
		if err != nil {
			return PolicyAssociationResp{}, err
		}

		rsp := PolicyAssociationResp{}
		err = json.Unmarshal(content, &rsp)
		return rsp, err

	case resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusNoContent:
		// Older minio - no result sent. TODO(aditya): Remove this case after
		// newer minio is released.
		return PolicyAssociationResp{}, nil

	default:
		// Error response case.
		return PolicyAssociationResp{}, httpRespToErrorResponse(resp)
	}
}

// AttachPolicy - attach policies to a user or group.
func (adm *AdminClient) AttachPolicy(ctx context.Context, r PolicyAssociationReq) (PolicyAssociationResp, error) {
	return adm.attachOrDetachPolicyBuiltin(ctx, true, r)
}

// DetachPolicy - detach policies from a user or group.
func (adm *AdminClient) DetachPolicy(ctx context.Context, r PolicyAssociationReq) (PolicyAssociationResp, error) {
	return adm.attachOrDetachPolicyBuiltin(ctx, false, r)
}

// GetPolicyEntities - returns builtin policy entities.
func (adm *AdminClient) GetPolicyEntities(ctx context.Context, q PolicyEntitiesQuery) (r PolicyEntitiesResult, err error) {
	params := make(url.Values)
	params["user"] = q.Users
	params["group"] = q.Groups
	params["policy"] = q.Policy

	reqData := requestData{
		relPath:     adminAPIPrefix + "/idp/builtin/policy-entities",
		queryValues: params,
	}

	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)
	defer closeResponse(resp)
	if err != nil {
		return r, err
	}

	if resp.StatusCode != http.StatusOK {
		return r, httpRespToErrorResponse(resp)
	}

	content, err := DecryptData(adm.getSecretKey(), resp.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(content, &r)
	return r, err
}
