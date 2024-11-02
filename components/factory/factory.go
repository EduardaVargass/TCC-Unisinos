package factory

import (
	"TCC-UNISINOS/components/machine"
	"context"
	"fmt"
	"log"
)

type FactoryComponent interface {
	CalculateOEEPerMachine(ctx context.Context) ([]OEEPerMachine, error)
}

type OEEPerMachine struct {
	MachineName string  `json:"machineName"`
	OEE         float64 `json:"oee"`
}

type FactoryComponentStruct struct {
	m machine.MachineComponentStruct
}

func (f *FactoryComponentStruct) Init(ctx context.Context, m machine.MachineComponentStruct) error {
	f.m = m
	return nil
}

func (f *FactoryComponentStruct) CalculateOEEPerMachine(ctx context.Context) ([]OEEPerMachine, error) {
	machines, err := f.m.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar máquinas: %v", err)
	}

	var oeeResults []OEEPerMachine

	for _, machine := range machines {
		oee, err := f.m.CalculateOEE(ctx, machine.MachineID)
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
