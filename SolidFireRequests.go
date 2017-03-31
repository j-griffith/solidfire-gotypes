// Copyright 2017 NetApp, Inc. All Rights Reserved.

package sftypes

type ListVolumesForAccountRequest struct {
	AccountID int64 `json:"accountID"`
}

type ListActiveVolumesRequest struct {
	StartVolumeID int64 `json:"startVolumeID"`
	Limit         int64 `json:"limit"`
}

type CreateVolumeRequest struct {
	Name       string      `json:"name"`
	AccountID  int64       `json:"accountID"`
	TotalSize  int64       `json:"totalSize"`
	Enable512e bool        `json:"enable512e"`
	Qos        QoS         `json:"qos,omitempty"`
	Attributes interface{} `json:"attributes"`
}

type DeleteVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type CloneVolumeRequest struct {
	VolumeID     int64       `json:"volumeID"`
	Name         string      `json:"name"`
	NewAccountID int64       `json:"newAccountID"`
	NewSize      int64       `json:"newSize"`
	Access       string      `json:"access"`
	SnapshotID   int64       `json:"snapshotID"`
	Attributes   interface{} `json:"attributes"`
}

type CreateSnapshotRequest struct {
	VolumeID                int64       `json:"volumeID"`
	SnapshotID              int64       `json:"snapshotID"`
	Name                    string      `json:"name"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication"`
	Retention               string      `json:"retention"`
	Attributes              interface{} `json:"attributes"`
}

type ListSnapshotsRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RollbackToSnapshotRequest struct {
	VolumeID         int64       `json:"volumeID"`
	SnapshotID       int64       `json:"snapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name"`
	Attributes       interface{} `json:"attributes"`
}

type DeleteSnapshotRequest struct {
	SnapshotID int64 `json:"snapshotID"`
}

type AddVolumesToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

type CreateVolumeAccessGroupRequest struct {
	Name       string   `json:"name"`
	Volumes    []int64  `json:"volumes,omitempty"`
	Initiators []string `json:"initiators,omitempty"`
}

type AddInitiatorsToVolumeAccessGroupRequest struct {
	Initiators []string `json:"initiators"`
	VAGID      int64    `json:"volumeAccessGroupID"`
}

type ListVolumeAccessGroupsRequest struct {
	StartVAGID int64 `json:"startVolumeAccessGroupID,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
}

type GetAccountByNameRequest struct {
	Name string `json:"username"`
}

type GetAccountByIDRequest struct {
	AccountID int64 `json:"accountID"`
}

type AddAccountRequest struct {
	Username        string      `json:"username"`
	InitiatorSecret string      `json:"initiatorSecret,omitempty"`
	TargetSecret    string      `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type GetClusterCapacityRequest struct {
}
