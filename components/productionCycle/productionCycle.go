package productioncycle

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ServiceWeaver/weaver"
)

type ProductionCycleComponent interface {
	Add(ctx context.Context, productionCycle ProductionCycle) error
	Get(ctx context.Context) ([]ProductionCycle, error)
	GetByMachineID(ctx context.Context, machineID int) ([]ProductionCycle, error)
	GetFirstCycleByMachineID(ctx context.Context, machineID int) (ProductionCycle, error)
}

type ProductionCycle struct {
	ProductionCycleID int       `json:"productionCycleID"`
	MachineID         int       `json:"machineID"`
	ProductionItem    string    `json:"productionItem"`
	ProdCount         int       `json:"prodCount"`
	RejCount          int       `json:"rejCount"`
	GoodCount         int       `json:"goodCount"`
	IdealCycleTime    float64   `json:"idealCycleTime"`
	ProductionTime    float64   `json:"productionTime"`
	Timestamp         time.Time `json:"timestamp"`
	weaver.AutoMarshal
}

type ProductionCycleComponentStruct struct {
	weaver.Implements[ProductionCycleComponent]
	db *sql.DB `json:"-"`
}

func (p *ProductionCycleComponentStruct) Init(ctx context.Context) error {
	db, err := sql.Open("sqlite3", "./productions.db")
	if err != nil {
		return fmt.Errorf("Erro ao conectar no banco de dados: %w", err)
	}
	p.db = db
	return nil
}

func (p *ProductionCycleComponentStruct) Add(ctx context.Context, productionCycle ProductionCycle) error {
	goodCount := productionCycle.ProdCount - productionCycle.RejCount
	_, err := p.db.Exec("INSERT INTO productionCycle (machineID, productionItem, prodCount, rejCount, goodCount, idealCycleTime, productionTime) VALUES (?, ?, ?, ?, ?, ?, ?)", productionCycle.MachineID, productionCycle.ProductionItem, productionCycle.ProdCount, productionCycle.RejCount, goodCount, productionCycle.IdealCycleTime, productionCycle.ProductionTime)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductionCycleComponentStruct) Get(ctx context.Context) ([]ProductionCycle, error) {
	var productionCycles []ProductionCycle
	query := "SELECT productionCycleID, machineID, productionItem, prodCount, rejCount, goodCount, idealCycleTime, productionTime, timestamp FROM productionCycle"
	res, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var productionCycle ProductionCycle
		err := res.Scan(&productionCycle.ProductionCycleID, &productionCycle.MachineID, &productionCycle.ProductionItem, &productionCycle.ProdCount, &productionCycle.RejCount, &productionCycle.GoodCount, &productionCycle.IdealCycleTime, &productionCycle.ProductionTime, &productionCycle.Timestamp)
		if err != nil {
			return nil, err
		}
		productionCycles = append(productionCycles, productionCycle)
	}
	return productionCycles, nil
}

func (p *ProductionCycleComponentStruct) GetByMachineID(ctx context.Context, machineID int) ([]ProductionCycle, error) {
	var productionCycles []ProductionCycle
	query := "SELECT productionCycleID, machineID, productionItem, prodCount, rejCount, goodCount, idealCycleTime, productionTime, timestamp FROM productionCycle WHERE machineID = ?"

	res, err := p.db.QueryContext(ctx, query, machineID)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var productionCycle ProductionCycle
		err := res.Scan(&productionCycle.ProductionCycleID, &productionCycle.MachineID, &productionCycle.ProductionItem, &productionCycle.ProdCount, &productionCycle.RejCount, &productionCycle.GoodCount, &productionCycle.IdealCycleTime, &productionCycle.ProductionTime, &productionCycle.Timestamp)
		if err != nil {
			return nil, err
		}
		productionCycles = append(productionCycles, productionCycle)
	}
	return productionCycles, nil
}

func (p *ProductionCycleComponentStruct) GetFirstCycleByMachineID(ctx context.Context, machineID int) (ProductionCycle, error) {
	var productionCycle ProductionCycle
	query := "SELECT productionCycleID, machineID, productionItem, prodCount, rejCount, goodCount, idealCycleTime, productionTime, timestamp FROM productionCycle WHERE machineID = ? LIMIT 1"

	err := p.db.QueryRowContext(ctx, query, machineID).Scan(&productionCycle.ProductionCycleID, &productionCycle.MachineID, &productionCycle.ProductionItem, &productionCycle.ProdCount, &productionCycle.RejCount, &productionCycle.GoodCount, &productionCycle.IdealCycleTime, &productionCycle.ProductionTime, &productionCycle.Timestamp)

	if err != nil {
		if err == sql.ErrNoRows {
			return productionCycle, nil
		}
		return productionCycle, err
	}

	return productionCycle, nil
}
