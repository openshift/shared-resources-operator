// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// SharedSecretSpecApplyConfiguration represents a declarative configuration of the SharedSecretSpec type for use
// with apply.
type SharedSecretSpecApplyConfiguration struct {
	SecretRef   *SharedSecretReferenceApplyConfiguration `json:"secretRef,omitempty"`
	Description *string                                  `json:"description,omitempty"`
}

// SharedSecretSpecApplyConfiguration constructs a declarative configuration of the SharedSecretSpec type for use with
// apply.
func SharedSecretSpec() *SharedSecretSpecApplyConfiguration {
	return &SharedSecretSpecApplyConfiguration{}
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *SharedSecretSpecApplyConfiguration) WithSecretRef(value *SharedSecretReferenceApplyConfiguration) *SharedSecretSpecApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SharedSecretSpecApplyConfiguration) WithDescription(value string) *SharedSecretSpecApplyConfiguration {
	b.Description = &value
	return b
}
