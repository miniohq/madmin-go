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

// GroupAddRemove is type for adding/removing members to/from a group.
type GroupAddRemove struct {
	Group    string      `json:"group"`
	Members  []string    `json:"members"`
	Status   GroupStatus `json:"groupStatus"`
	IsRemove bool        `json:"isRemove"`
}

// UpdateGroupMembers - adds/removes users to/from a group. Server
// creates the group as needed. Group is removed if remove request is
// made on empty group.
func (adm *AdminClient) UpdateGroupMembers(ctx context.Context, g GroupAddRemove) error {
	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	reqData := requestData{
		relPath: adminAPIPrefix + "/update-group-members",
		content: data,
	}

	// Execute PUT on /minio/admin/v3/update-group-members
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

// GroupDesc is a type that holds group info along with the policy
// attached to it.
type GroupDesc struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Members   []string  `json:"members"`
	Policy    string    `json:"policy"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// GetGroupDescription - fetches information on a group.
func (adm *AdminClient) GetGroupDescription(ctx context.Context, group string) (*GroupDesc, error) {
	v := url.Values{}
	v.Set("group", group)
	reqData := requestData{
		relPath:     adminAPIPrefix + "/group",
		queryValues: v,
	}

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

	gd := GroupDesc{}
	if err = json.Unmarshal(data, &gd); err != nil {
		return nil, err
	}

	return &gd, nil
}

// ListGroups - lists all groups names present on the server.
func (adm *AdminClient) ListGroups(ctx context.Context) ([]string, error) {
	reqData := requestData{
		relPath: adminAPIPrefix + "/groups",
	}

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

	groups := []string{}
	if err = json.Unmarshal(data, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}

// GroupStatus - group status.
type GroupStatus string

// GroupStatus values.
const (
	GroupEnabled  GroupStatus = "enabled"
	GroupDisabled GroupStatus = "disabled"
)

// SetGroupStatus - sets the status of a group.
func (adm *AdminClient) SetGroupStatus(ctx context.Context, group string, status GroupStatus) error {
	v := url.Values{}
	v.Set("group", group)
	v.Set("status", string(status))

	reqData := requestData{
		relPath:     adminAPIPrefix + "/set-group-status",
		queryValues: v,
	}

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
