package db

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

// MetaDataWriter interfaces opentracing.TextMap and is being used
// to propagate traces to the database cluster.
type MetaDataWriter struct {
	metadata.MD
}

// ForeachKey implements the Foreach function of opentracing.TextMap
// interface.
func (mdw MetaDataWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range mdw.MD {
		for _, v := range vs {
			err := handler(k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Set implements the Set function of opentracing.TextMap interface.
func (mdw MetaDataWriter) Set(key, val string) {
	key = strings.ToLower(key)
	mdw.MD[key] = append(mdw.MD[key], val)
}
