package factory

import (
	"TCC-UNISINOS/components/machine"
	"context"
	"fmt"
	"log"

	"github.com/ServiceWeaver/weaver"
)

type FactoryComponent interface {
	CalculateOEEPerMachine(ctx context.Context) ([]OEEPerMachine, error)
}

type OEEPerMachine struct {
	MachineName string  `json:"machineName"`
	OEE         float64 `json:"oee"`
	weaver.AutoMarshal
}

type factoryComponent struct {
	weaver.Implements[FactoryComponent]
	m weaver.Ref[machine.MachineComponent]
}

func (f *factoryComponent) CalculateOEEPerMachine(ctx context.Context) ([]OEEPerMachine, error) {
	machines, err := f.m.Get().Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar máquinas: %v", err)
	}

	var oeeResults []OEEPerMachine

	for _, machine := range machines {
		oee, err := f.m.Get().CalculateOEE(ctx, machine.MachineID)
		if err != nil {
			log.Printf("Erro ao calcular OEE para a máquina %s: %v", machine.Name, err)
			continue
		}

		oeeResults = append(oeeResults, OEEPerMachine{
			MachineName: machine.Name,
			OEE:         oee,
		})
	}

	return oeeResults, nil
}
