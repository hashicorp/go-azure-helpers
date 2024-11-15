package typehelpers

import (
	"context"

	"github.com/hashicorp/go-azure-helpers/framework/fwdiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type objectTypeOf[T any] struct {
	basetypes.ObjectType
}

func newObjectTypeOf[T any](ctx context.Context) (objectTypeOf[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	m, d := AttributeTypes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return objectTypeOf[T]{}, diags
	}

	return objectTypeOf[T]{basetypes.ObjectType{AttrTypes: m}}, diags
}

func NewObjectTypeOf[T any](ctx context.Context) objectTypeOf[T] {
	return fwdiag.Must(newObjectTypeOf[T](ctx))
}

func objectValueObjectPtr[T any](ctx context.Context, val attr.Value) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	ptr, d := objectTypeNewObjectPtr[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	diags.Append(val.(basetypes.ObjectValue).As(ctx, ptr, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}

	return ptr, diags
}

func objectTypeNewObjectPtr[T any](context.Context) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics
	return new(T), diags
}

// ObjectValueOf represents a Terraform Plugin Framework Object value whose corresponding Go type is the structure T.
type ObjectValueOf[T any] struct {
	basetypes.ObjectValue
}

func (v ObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(ObjectValueOf[T])
	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v ObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewObjectTypeOf[T](ctx)
}

func (v ObjectValueOf[T]) ToObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return v.ToPtr(ctx)
}

func (v ObjectValueOf[T]) ToPtr(ctx context.Context) (*T, diag.Diagnostics) {
	return objectValueObjectPtr[T](ctx, v)
}

func NewObjectValueOfNull[T any](ctx context.Context) ObjectValueOf[T] {
	return ObjectValueOf[T]{ObjectValue: basetypes.NewObjectNull(AttributeTypesMust[T](ctx))}
}

func NewObjectValueOfUnknown[T any](ctx context.Context) ObjectValueOf[T] {
	return ObjectValueOf[T]{ObjectValue: basetypes.NewObjectUnknown(AttributeTypesMust[T](ctx))}
}

func NewObjectValueOf[T any](ctx context.Context, t *T) (ObjectValueOf[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	m, d := AttributeTypes[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return NewObjectValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewObjectValueFrom(ctx, m, t)
	diags.Append(d...)
	if diags.HasError() {
		return NewObjectValueOfUnknown[T](ctx), diags
	}

	return ObjectValueOf[T]{ObjectValue: v}, diags
}

func nestedObjectTypeNewObjectSlice[T any](_ context.Context, len, cap int) ([]*T, diag.Diagnostics) { //nolint:unparam
	var diags diag.Diagnostics

	return make([]*T, len, cap), diags
}
