package confflags

import (
	"context"

	"github.com/spf13/pflag"
	"github.com/sv-tools/conf"
)

type flagReader struct {
	mapFlagKey map[string]string
	flagSet    *pflag.FlagSet
	prefix     string
}

func (r *flagReader) Prefix() string {
	return r.prefix
}

func (r *flagReader) Read(ctx context.Context) (any, error) {
	res := map[string]string{}
	for name, key := range r.mapFlagKey {
		if fl := r.flagSet.Lookup(name); fl != nil {
			res[key] = fl.Value.String()
		}
	}

	return res, ctx.Err()
}

// New creates the Env reader
//
//	`mapFlagKey` is a map of the names of the flag and the configuration keys
//	`prefix` is a default prefix that will be added to all configuration keys
func New(mapFlagKey map[string]string, prefix string, flagSet *pflag.FlagSet) conf.Reader {
	if flagSet == nil {
		flagSet = pflag.CommandLine
	}
	return &flagReader{
		mapFlagKey: mapFlagKey,
		prefix:     prefix,
		flagSet:    flagSet,
	}
}
