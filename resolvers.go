package tools

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/kinds"
)

// Resolver interface to a resolver configuration
type Resolver interface {
	GetKind() string
}

// ResolverMap a map of resolver configurations
type ResolverMap map[string]Resolver

// FieldResolveMap map of field resolve functions
type FieldResolveMap map[string]graphql.FieldResolveFn

// ObjectResolver config for object resolver map
type ObjectResolver struct {
	IsTypeOf graphql.IsTypeOfFn
	Fields   FieldResolveMap
}

// GetKind gets the kind
func (c *ObjectResolver) GetKind() string {
	return kinds.ObjectDefinition
}

// ScalarResolver config for a scalar resolve map
type ScalarResolver struct {
	Serialize    graphql.SerializeFn
	ParseValue   graphql.ParseValueFn
	ParseLiteral graphql.ParseLiteralFn
}

// GetKind gets the kind
func (c *ScalarResolver) GetKind() string {
	return kinds.ScalarDefinition
}

// InterfaceResolver config for interface resolve
type InterfaceResolver struct {
	ResolveType graphql.ResolveTypeFn
	Fields      FieldResolveMap
}

// GetKind gets the kind
func (c *InterfaceResolver) GetKind() string {
	return kinds.InterfaceDefinition
}

// UnionResolver config for interface resolve
type UnionResolver struct {
	ResolveType graphql.ResolveTypeFn
}

// GetKind gets the kind
func (c *UnionResolver) GetKind() string {
	return kinds.UnionDefinition
}

// EnumResolver config for enum values
type EnumResolver struct {
	Values map[string]interface{}
}

// GetKind gets the kind
func (c *EnumResolver) GetKind() string {
	return kinds.EnumDefinition
}