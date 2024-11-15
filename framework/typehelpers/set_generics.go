package typehelpers

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/framework/fwdiag"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type setNestedObjectTypeOf[T any] struct {
	basetypes.SetType
}

type SetNestedObjectValueOf[T any] struct {
	basetypes.SetValue
}

var (
	_ basetypes.SetTypable        = (*setNestedObjectTypeOf[struct{}])(nil)
	_ NestedObjectCollectionType  = (*setNestedObjectTypeOf[struct{}])(nil)
	_ basetypes.SetValuable       = (*SetNestedObjectValueOf[struct{}])(nil)
	_ NestedObjectCollectionValue = (*SetNestedObjectValueOf[struct{}])(nil)
)

func NewSetNestedObjectTypeOf[T any](ctx context.Context) setNestedObjectTypeOf[T] {
	return setNestedObjectTypeOf[T]{basetypes.SetType{ElemType: NewObjectTypeOf[T](ctx)}}
}

func (s setNestedObjectTypeOf[T]) NewObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return objectTypeNewObjectPtr[T](ctx)
}

func (s setNestedObjectTypeOf[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	return NewSetNestedObjectValueOfNull[T](ctx), diags
}

func (s setNestedObjectTypeOf[T]) ValueFromObjectPtr(ctx context.Context, ptr any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := ptr.(*T); ok {
		v, d := NewSetNestedObjectValueOfPtr(ctx, v)
		diags.Append(d...)
		return v, d
	}

	diags.Append(diag.NewErrorDiagnostic("Invalid pointer value", fmt.Sprintf("incorrect type: want %T, got %T", (*T)(nil), ptr)))
	return nil, diags
}

func (s setNestedObjectTypeOf[T]) ValueFromObjectSlice(ctx context.Context, a any) (attr.Value, diag.Diagnostics) {
	//TODO implement me
	panic("implement me")
}

func (s setNestedObjectTypeOf[T]) NewObjectSlice(ctx context.Context, len int, cap int) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectSlice[T](ctx, len, cap)
}

func (s SetNestedObjectValueOf[T]) ToObjectSlice(ctx context.Context) (any, diag.Diagnostics) {
	return s.ToSlice(ctx)
}

func (s SetNestedObjectValueOf[T]) ToSlice(ctx context.Context) ([]*T, diag.Diagnostics) {
	return nestedObjectValueObjectSlice[T](ctx, s.SetValue)
}

func (s SetNestedObjectValueOf[T]) ToObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return s.ToPtr(ctx)
}

func (s SetNestedObjectValueOf[T]) ToPtr(ctx context.Context) (*T, diag.Diagnostics) {
	return nestedObjectValueObjectPtr[T](ctx, s.SetValue)
}

func NewSetNestedObjectValueOfNull[T any](ctx context.Context) SetNestedObjectValueOf[T] {
	return SetNestedObjectValueOf[T]{SetValue: basetypes.NewSetNull(NewObjectTypeOf[T](ctx))}
}

func NewSetNestedObjectValueOfPtr[T any](ctx context.Context, t *T) (SetNestedObjectValueOf[T], diag.Diagnostics) {
	return NewSetNestedObjectValueOfSlice(ctx, []*T{t})
}

func NewSetNestedObjectValueOfSlice[T any](ctx context.Context, ts []*T) (SetNestedObjectValueOf[T], diag.Diagnostics) {
	return newSetNestedObjectValueOf[T](ctx, ts)
}

func NewSetNestedObjectValueOfValueSliceMust[T any](ctx context.Context, ts []T) SetNestedObjectValueOf[T] {
	return fwdiag.Must(newSetNestedObjectValueOf[T](ctx, ts))
}

func newSetNestedObjectValueOf[T any](ctx context.Context, elements any) (SetNestedObjectValueOf[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	typ, d := newObjectTypeOf[T](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return NewSetNestedObjectValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewSetValueFrom(ctx, typ, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewSetNestedObjectValueOfUnknown[T](ctx), diags
	}

	return SetNestedObjectValueOf[T]{SetValue: v}, diags
}

func NewSetNestedObjectValueOfUnknown[T any](ctx context.Context) SetNestedObjectValueOf[T] {
	return SetNestedObjectValueOf[T]{SetValue: basetypes.NewSetUnknown(NewObjectTypeOf[T](ctx))}
}

func (s SetNestedObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(SetNestedObjectValueOf[T])

	if !ok {
		return false
	}

	return s.SetValue.Equal(other.SetValue)
}

func (s SetNestedObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewSetNestedObjectTypeOf[T](ctx)
}
