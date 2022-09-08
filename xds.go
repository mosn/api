package api

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
)

// XDSConfigHandler converts a map according to any.Any
type XDSConfigHandler func(s *any.Any) (map[string]interface{}, error)

var registryXDSConfigFactory map[string]XDSConfigHandler

func init() {
	registryXDSConfigFactory = make(map[string]XDSConfigHandler)
}

// RegisterXDSConfigHandler registers the filterType as XDSConfigHandler
func RegisterXDSConfigHandler(filterType string, creator XDSConfigHandler) {
	registryXDSConfigFactory[filterType] = creator
}

// HandleXDSConfig  converts pb Any to map according to the filterType
func HandleXDSConfig(filterType string, s *any.Any) (map[string]interface{}, error) {
	if cf, ok := registryXDSConfigFactory[filterType]; ok {
		config, err := cf(s)
		if err != nil {
			return nil, fmt.Errorf("convert pb.Any to map failed: %v", err)
		}
		return config, nil
	}
	return nil, fmt.Errorf("unsupported stream filter type: %v", filterType)
}
