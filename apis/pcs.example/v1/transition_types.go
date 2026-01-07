// Copyright © 2025 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v1

import (
	"context"
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
	Description string `json:"description,omitempty" validate:"max=200"`
	// Add your spec fields here
}

// TransitionStatus defines the observed state of Transition
type TransitionStatus struct {
	Phase      string `json:"phase,omitempty"`
	Message    string `json:"message,omitempty"`
	Ready      bool   `json:"ready"`
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
