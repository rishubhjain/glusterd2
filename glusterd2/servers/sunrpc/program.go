package sunrpc

import (
	"net"
	"reflect"
	"strings"

	"github.com/gluster/glusterd2/pkg/sunrpc"

	log "github.com/sirupsen/logrus"
)

// RPC program implementations inside this package can use this type for convenience
type genericProgram struct {
	name        string
	progNum     uint32
	progVersion uint32
	procedures  []sunrpc.Procedure
	conn        net.Conn
}

// registerProcedures creates procedure number to procedure name mappings for sunrpc codec
func registerProcedures(program sunrpc.Program) error {
	logger := log.WithFields(log.Fields{
		"program": program.Name(),
		"prognum": program.Number(),
		"progver": program.Version(),
	})

	logger.Debug("registering sunrpc program")

	// Create procedure number to procedure name mappings for sunrpc codec
	typeName := reflect.Indirect(reflect.ValueOf(program)).Type().Name()
	for _, procedure := range program.Procedures() {
		log.WithFields(log.Fields{
			"procId":   procedure.ID,
			"procName": procedure.Name,
		}).Debug("registering sunrpc procedure")

		if !strings.HasPrefix(procedure.Name, typeName+".") {
			procedure.Name = typeName + "." + procedure.Name
		}
		if err := sunrpc.RegisterProcedure(
			sunrpc.Procedure{
				ID:   procedure.ID,
				Name: procedure.Name,
			}, true); err != nil {
			return err
		}
	}

	return nil
}
