package main

import (
	factory "TCC-UNISINOS/components/factory"
	machine "TCC-UNISINOS/components/machine"
	productioncycle "TCC-UNISINOS/components/productionCycle"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	productionCyclesComponent weaver.Ref[productioncycle.ProductionCycleComponent]
	machinesComponent         weaver.Ref[machine.MachineComponent]
	factoryComponent          weaver.Ref[factory.FactoryComponent]
	production                weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	var productionCyclesComponent productioncycle.ProductionCycleComponent = app.productionCyclesComponent.Get()
	var machinesComponent machine.MachineComponent = app.machinesComponent.Get()
	var factoryComponent factory.FactoryComponent = app.factoryComponent.Get()

	r := chi.NewRouter()
	r.Post("/production/productionCycles", func(w http.ResponseWriter, r *http.Request) {
		var productionCycle productioncycle.ProductionCycle
		if err := json.NewDecoder(r.Body).Decode(&productionCycle); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := productionCyclesComponent.Add(ctx, productionCycle); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		result, err := json.Marshal(productionCycle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
	})

	r.Get("/production/productionCycles", func(w http.ResponseWriter, r *http.Request) {
		productionCycles, err := productionCyclesComponent.Get(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(productionCycles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Post("/production/machines", func(w http.ResponseWriter, r *http.Request) {
		var machine machine.Machine
		if err := json.NewDecoder(r.Body).Decode(&machine); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := machinesComponent.Add(ctx, machine); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		result, err := json.Marshal(machine)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
	})

	r.Put("/production/machines/{machineID}", func(w http.ResponseWriter, r *http.Request) {
		machineIDStr := chi.URLParam(r, "machineID")
		machineID, err := strconv.Atoi(machineIDStr)
		if err != nil {
			http.Error(w, "ID da máquina inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		var updatedMachine machine.Machine
		if err := json.NewDecoder(r.Body).Decode(&updatedMachine); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedMachine, err = machinesComponent.Update(ctx, machineID, updatedMachine)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(updatedMachine)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/production/machines", func(w http.ResponseWriter, r *http.Request) {
		machines, err := machinesComponent.Get(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(machines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/production/oee/{machineID}", func(w http.ResponseWriter, r *http.Request) {
		machineIDStr := chi.URLParam(r, "machineID")
		machineID, err := strconv.Atoi(machineIDStr)
		if err != nil {
			http.Error(w, "ID da máquina inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		oee, err := machinesComponent.CalculateOEE(ctx, machineID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(oee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/production/oee/all", func(w http.ResponseWriter, r *http.Request) {
		oeePerMachines, err := factoryComponent.CalculateOEEPerMachine(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(oeePerMachines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	otelHandler := otelhttp.NewHandler(r, "production")

	return http.Serve(app.production, otelHandler)
}
