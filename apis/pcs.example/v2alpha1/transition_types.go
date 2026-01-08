// Copyright © 2025 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v2alpha1

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/openchami/fabrica/pkg/fabrica"
)

// Transition represents a transition resource
type Transition struct {
	APIVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   fabrica.Metadata `json:"metadata"`
	Spec       TransitionSpec   `json:"spec" validate:"required"`
	Status     TransitionStatus `json:"status,omitempty"`
}

// TransitionSpec defines the desired state of Transition
type TransitionSpec struct {
	// TransitionID is the transition's ID.
	TransitionID uuid.UUID `json:"transitionID" db:"id"`
	// Operation indicates the operation to perform, such as off or soft restart.
	Operation Operation `json:"operation" db:"operation"`
	// TaskDeadline is the time limit for completing the transition.
	TaskDeadline int `json:"taskDeadlineMinutes" db:"deadline"`
	// CreateTime is the time the transition was requested.
	CreateTime time.Time `json:"createTime" db:"created"`
	// LastActiveTime is a timestamp the power service updates regularly as long as it considers the transition active.
	// This is used to reap old transitions that have been inactive for longer than a time threshold.
	LastActiveTime time.Time `json:"lastActiveTime" db:"active"`
	// AutomaticExpirationTime is a timestamp that will always reap a transition, regardless of its active status.
	AutomaticExpirationTime time.Time `json:"automaticExpirationTime" db:"expires"`
	// Status is the current phase of the transition lifecycle.
	Status string `json:"transitionStatus" db:"status"`
	// TaskIDs are the IDs of individual tasks in the transition/
	TaskIDs []uuid.UUID

	// Only populated when the task is completed

	// IsCompressed and TaskCounts are kinda  weird. PCS has the capability to only keep aggregate data about how many
	// tasks succeeded or failed in TaskCounts, if it sets full=false when calling ToTransitionResp on a Transition and
	// task set. However, it never sets this, and the full set of task info is always included in Transition.Tasks.

	// IsCompressed indicates if the transition has its task counts tallied.
	IsCompressed bool `json:"isCompressed" db:"compressed"`
}

type Operation int

const (
	Operation_Nil         Operation = iota - 1
	Operation_On                    // On = 0
	Operation_Off                   // 1 GracfulShutdown/Off->ForceOff
	Operation_SoftRestart           // 2 GracefulRestart->ForceRestart Or GracfulShutdown/Off->ForceOff->On
	Operation_HardRestart           // 3 GracfulShutdown/Off->ForceOff->On
	Operation_Init                  // 4 GracfulShutdown/Off->ForceOff->On does not require the initial power state to be "on"
	Operation_ForceOff              // 5 ForceOff
	Operation_SoftOff               // 6 GracfulShutdown/Off
)

// TransitionStatus defines the observed state of Transition
type TransitionStatus struct {
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
	Ready   bool   `json:"ready"`
	// Add your status fields here
}

// Validate implements custom validation logic for Transition
func (r *Transition) Validate(ctx context.Context) error {
	// Add custom validation logic here
	// Example:
	// if r.Spec.Description == "forbidden" {
	//     return errors.New("description 'forbidden' is not allowed")
	// }

	return nil
}

// GetKind returns the kind of the resource
func (r *Transition) GetKind() string {
	return "Transition"
}

// GetName returns the name of the resource
func (r *Transition) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r *Transition) GetUID() string {
	return r.Metadata.UID
}

// IsHub marks this as the hub/storage version
func (r *Transition) IsHub() {}
