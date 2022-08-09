package gqlgendepth

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type (
	Foo struct {

			es graphql.ExecutableSchema

	}
)
func (e Foo)ExtensionName() string {
	return "DepthLimit"

}
func (e Foo)	Validate(schema graphql.ExecutableSchema) error {
	e.es = schema
	return nil

}


// OperationContextMutator is called after creating the request context, but before executing the root resolver.
func (e Foo )MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {

	op := rc.Doc.Operations.ForName(rc.OperationName)



	complexity := Calculate(e.es, op, rc.Variables)



	if true {
		err := gqlerror.Errorf("operation has complexity %d, which exceeds the limit of %d", complexity, 0)
		errcode.Set(err, "TODO")
		return err
	}

	return nil
}
