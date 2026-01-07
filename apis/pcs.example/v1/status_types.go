// Copyright © 2025 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v1

import (
	"context"
	"github.com/openchami/fabrica/pkg/fabrica"
)

// Status represents a status resource
type Status struct {
	APIVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   fabrica.Metadata `json:"metadata"`
	Spec       StatusSpec   `json:"spec" validate:"required"`
	Status     StatusStatus `json:"status,omitempty"`
}

// StatusSpec defines the desired state of Status
type StatusSpec struct {
	Description string `json:"description,omitempty" validate:"max=200"`
	// Add your spec fields here
}

// StatusStatus defines the observed state of Status
type StatusStatus struct {
	Phase      string `json:"phase,omitempty"`
	Message    string `json:"message,omitempty"`
	Ready      bool   `json:"ready"`
	// Add your status fields here
}

// Validate implements custom validation logic for Status
func (r *Status) Validate(ctx context.Context) error {
	// Add custom validation logic here
	// Example:
	// if r.Spec.Description == "forbidden" {
	//     return errors.New("description 'forbidden' is not allowed")
	// }

	return nil
}
// GetKind returns the kind of the resource
func (r *Status) GetKind() string {
	return "Status"
}

// GetName returns the name of the resource
func (r *Status) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r *Status) GetUID() string {
	return r.Metadata.UID
}

// IsHub marks this as the hub/storage version
func (r *Status) IsHub() {}
