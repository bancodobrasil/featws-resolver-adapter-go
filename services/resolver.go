package services

import (
	"context"
	"sync"

	"github.com/bancodobrasil/featws-resolver-adapter-go/types"
	log "github.com/sirupsen/logrus"
)

// ResolverFunc define the Resolver Function structure
type ResolverFunc func(context.Context, types.ResolveInput, *types.ResolveOutput)

var lock = &sync.Mutex{}

var resolverFunc ResolverFunc

// SetupResolver to config the current resolver func
func SetupResolver(rFunc ResolverFunc) {
	lock.Lock()
	defer lock.Unlock()
	if resolverFunc == nil {
		if resolverFunc == nil {
			log.Debugln("Creating single instance now.")
			resolverFunc = rFunc
		} else {
			log.Debugln("Single instance already created.")
		}
	} else {
		log.Debugln("Single instance already created.")
	}
}

// Resolve to execute the resolver
func Resolve(ctx context.Context, input types.ResolveInput) (output *types.ResolveOutput) {
	output = &types.ResolveOutput{
		Context: input.Context,
		Errors:  make(map[string]interface{}),
	}
	resolverFunc(ctx, input, output)

	if len(input.Load) > 0 {
		oldContext := output.Context

		output.Context = make(map[string]interface{})

		for _, key := range input.Load {
			value, ok := oldContext[key]
			if ok {
				output.Context[key] = value
			}
		}
	}

	return
}
