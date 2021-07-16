package api_test

import (
	"testing"

	pb "github.com/genkami/examples/go/protobuf/ensure-idempotent/api"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestApi(t *testing.T) {
	fileDesc := pb.File_api_api_proto
	methodDescs := fileDesc.Services().ByName("Example").Methods()
	for i := 0; i < methodDescs.Len(); i++ {
		methodDesc := methodDescs.Get(i)
		opts := methodDesc.Options().(*descriptorpb.MethodOptions)
		var hasSideEffects bool
		if opts.IdempotencyLevel == nil {
			hasSideEffects = true
		} else {
			hasSideEffects = *opts.IdempotencyLevel != descriptorpb.MethodOptions_NO_SIDE_EFFECTS
		}

		if hasSideEffects {
			idempotencyKeyDesc := methodDesc.Input().Fields().ByName("idempotency_key")
			if idempotencyKeyDesc == nil {
				t.Errorf("%s: non-readonly method must have idempotency_key", methodDesc.Name())
			}
		}
	}
}
