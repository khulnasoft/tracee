package controlplane

import (
	"github.com/khulnasoft/tracee/pkg/bufferdecoder"
	"github.com/khulnasoft/tracee/pkg/errfmt"
	"github.com/khulnasoft/tracee/pkg/events"
	"github.com/khulnasoft/tracee/types/trace"
)

type signal struct {
	id   events.ID
	args []trace.Argument
}

func (sig *signal) Unmarshal(buffer []byte) error {
	ebpfDecoder := bufferdecoder.New(buffer)
	var eventIdUint32 uint32
	err := ebpfDecoder.DecodeUint32(&eventIdUint32)
	if err != nil {
		return errfmt.Errorf("failed to decode signal event ID: %v", err)
	}
	sig.id = events.ID(eventIdUint32)
	var argnum uint8
	err = ebpfDecoder.DecodeUint8(&argnum)
	if err != nil {
		return errfmt.Errorf("failed to decode signal argnum: %v", err)
	}

	if !events.Core.IsDefined(sig.id) {
		return errfmt.Errorf("failed to get event %d configuration", sig.id)
	}
	eventDefinition := events.Core.GetDefinitionByID(sig.id)
	sig.args = make([]trace.Argument, len(eventDefinition.GetParams()))
	err = ebpfDecoder.DecodeArguments(sig.args, int(argnum), eventDefinition, sig.id)
	if err != nil {
		return errfmt.Errorf("failed to decode signal arguments: %v", err)
	}

	return nil
}
