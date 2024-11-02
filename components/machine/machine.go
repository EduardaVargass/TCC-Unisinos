package machine

import (
	oee "TCC-UNISINOS/components/oee"
	productioncycle "TCC-UNISINOS/components/productionCycle"
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"
)

type MachineComponent interface {
	Add(ctx context.Context, machine Machine) error
	Update(ctx context.Context, machineID int, machine Machine) (Machine, error)
	Get(ctx context.Context) ([]Machine, error)
	GetByMachineID(ctx context.Context, machineID int) (Machine, error)
	CalculateOEE(ctx context.Context, machineID int) (float64, error)
}

type Machine struct {
	MachineID     int     `json:"machineID"`
	Name          string  `json:"name"`
	AvailableTime float64 `json:"availableTime"`
}

type MachineComponentStruct struct {
	db *sql.DB `json:"-"`
	p  productioncycle.ProductionCycleComponentStruct
	o  oee.OEEComponentStruct
}

func (m *MachineComponentStruct) Init(ctx context.Context, p productioncycle.ProductionCycleComponentStruct, o oee.OEEComponentStruct) error {
	m.p = p
	m.o = o

	db, err := sql.Open("sqlite3", "./productions.db")
	if err != nil {
		return fmt.Errorf("Erro ao abrir banco de dados: %w", err)
	}
	m.db = db
	return nil
}

func (m *MachineComponentStruct) Add(ctx context.Context, machine Machine) error {
	_, err := m.db.Exec("INSERT INTO machine (name, availableTime) VALUES (?, ?)", machine.Name, machine.AvailableTime)
	if err != nil {
		return err
	}
	return nil
}

func (m *MachineComponentStruct) Update(ctx context.Context, machineID int, machine Machine) (Machine, error) {
	_, err := m.db.Exec(
		"UPDATE machine SET name = ?, availableTime = ? WHERE machineID = ?",
		machine.Name,
		machine.AvailableTime,
		machineID,
	)
	if err != nil {
		return Machine{}, err
	}

	updatedMachine := Machine{
		MachineID:     machineID,
		Name:          machine.Name,
		AvailableTime: machine.AvailableTime,
	}
	return updatedMachine, nil
}

func (m *MachineComponentStruct) Get(ctx context.Context) ([]Machine, error) {
	var machines []Machine
	query := "SELECT machineID, name, availableTime FROM machine"
	res, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var machine Machine
		err := res.Scan(&machine.MachineID, &machine.Name, &machine.AvailableTime)
		if err != nil {
			return nil, err
		}
		machines = append(machines, machine)
	}
	return machines, nil
}

func (m *MachineComponentStruct) GetByMachineID(ctx context.Context, machineID int) (Machine, error) {
	var machine Machine
	query := "SELECT machineID, name, availableTime FROM machine WHERE machineID = ?"
	res := m.db.QueryRowContext(ctx, query, machineID)
	err := res.Scan(&machine.MachineID, &machine.Name, &machine.AvailableTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return machine, nil
		}
		return machine, err
	}
	return machine, nil
}

func (m *MachineComponentStruct) CalculateOEE(ctx context.Context, machineID int) (float64, error) {
	availableTime, err := m.calculateAvailableTime(ctx, machineID)
	if err != nil {
		return 0, err
	}

	productionCycles, err := m.p.GetByMachineID(ctx, machineID)
	if err != nil {
		return 0, err
	}

	var productionData oee.ProductionData
	productionData.AvailableTime = availableTime

	for _, cycle := range productionCycles {
		productionData.ProductionTime += cycle.ProductionTime
		productionData.IdealProdCount += int(cycle.ProductionTime / cycle.IdealCycleTime)
		productionData.RejCount += cycle.RejCount
		productionData.GoodCount += cycle.GoodCount
	}

	resultOEE, err := m.o.CalculateOEE(ctx, productionData)
	if err != nil {
		return 0, err
	}
	return resultOEE, nil
}

func (m *MachineComponentStruct) calculateAvailableTime(ctx context.Context, machineID int) (float64, error) {
	machine, err := m.GetByMachineID(ctx, machineID)
	if err != nil {
		return 0, err
	}

	productionCycle, err := m.p.GetFirstCycleByMachineID(ctx, machineID)
	if err != nil {
		return 0, err
	}

	if productionCycle.ProductionCycleID == 0 {
		return 0, nil
	}

	currentTime := time.Now()
	daysDifference := math.Ceil(currentTime.Sub(productionCycle.Timestamp).Hours() / 24)

	totalAvailableTime := daysDifference * machine.AvailableTime

	return totalAvailableTime, nil
}
