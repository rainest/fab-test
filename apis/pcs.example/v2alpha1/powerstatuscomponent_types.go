// Copyright © 2025 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v2alpha1

import (
	"context"
	"time"

	"github.com/openchami/fabrica/pkg/fabrica"
)

// PowerStatusComponent represents a powerstatuscomponent resource
type PowerStatusComponent struct {
	APIVersion string                     `json:"apiVersion"`
	Kind       string                     `json:"kind"`
	Metadata   fabrica.Metadata           `json:"metadata"`
	Spec       PowerStatusComponentSpec   `json:"spec" validate:"required"`
	Status     PowerStatusComponentStatus `json:"status,omitempty"`
}

// PowerStatusComponentSpec defines the desired state of PowerStatusComponent
type PowerStatusComponentSpec struct {
	XName                     string    `json:"xname" db:"xname"`
	PowerState                string    `json:"powerState" db:"power_state"`
	ManagementState           string    `json:"managementState" db:"management_state"`
	Error                     string    `json:"error" db:"error"`
	SupportedPowerTransitions []string  `json:"supportedPowerTransitions" db:"supported_power_transitions"`
	LastUpdated               time.Time `json:"lastUpdated" db:"last_updated"`
}

// PowerStatusComponentStatus defines the observed state of PowerStatusComponent
type PowerStatusComponentStatus struct {
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
	Ready   bool   `json:"ready"`
	// Add your status fields here
}

// Validate implements custom validation logic for PowerStatusComponent
func (r *PowerStatusComponent) Validate(ctx context.Context) error {
	// Add custom validation logic here
	// Example:
	// if r.Spec.Description == "forbidden" {
	//     return errors.New("description 'forbidden' is not allowed")
	// }

	return nil
}

// GetKind returns the kind of the resource
func (r *PowerStatusComponent) GetKind() string {
	return "PowerStatusComponent"
}

// GetName returns the name of the resource
func (r *PowerStatusComponent) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r *PowerStatusComponent) GetUID() string {
	return r.Metadata.UID
}

// IsHub marks this as the hub/storage version
func (r *PowerStatusComponent) IsHub() {}
