package main

import (
	//"fmt"
	"fmt"
	"testing"
	"time"

	"github.com/JorgeMG117/LolBets/backend/client"
	"github.com/JorgeMG117/LolBets/backend/server"
)

func startServer(){
}

func TestBackend(t *testing.T) {
    fmt.Println("Lanzando servidor")
    go server.ExecServer() //Comprobar error
    time.Sleep(5 * time.Second)
    fmt.Println("Lanzando cliente")
    go client.Client()
    for {}
}


/*type configDespliegue struct {
	t           *testing.T
    clientes    
	conectados  []bool
	numReplicas int
	nodosRaft   []rpctimeout.HostPort
	cr          canalResultados
}

// Crear una configuracion de despliegue
func makeCfgDespliegue(t *testing.T, n int, nodosraft []string,	conectados []bool) *configDespliegue {
	cfg := &configDespliegue{}
	cfg.t = t
	cfg.conectados = conectados
	cfg.numReplicas = n
	cfg.nodosRaft = rpctimeout.StringArrayToHostPortArray(nodosraft)
	cfg.cr = make(canalResultados, 2000)

	return cfg
}

func TestBackend(t *testing.T) {
    server.ExecServer() //Comprobar error
    cfg := makeCfgDespliegue(
        t,
		3,
		[]string{REPLICA1, REPLICA2, REPLICA3},
		[]bool{true, true, true}
    ) 
    defer cfg.stop()
    
	// Test1 : No debería haber ningun primario, si SV no ha recibido aún latidos
	t.Run("T1:soloArranqueYparada",
		func(t *testing.T) { cfg.soloArranqueYparadaTest1(t) })
    fmt.Println("Test")
}

func (cfg *configDespliegue) soloArranqueYparadaTest1(t *testing.T) {
	t.Skip("SKIPPED soloArranqueYparadaTest1")

	fmt.Println(t.Name(), ".....................")

	cfg.t = t // Actualizar la estructura de datos de tests para errores

	// Poner en marcha replicas en remoto con un tiempo de espera incluido
	cfg.startDistributedProcesses()

	// Comprobar estado replica 0
	cfg.comprobarEstadoRemoto(0, 0, false, -1)

	// Comprobar estado replica 1
	cfg.comprobarEstadoRemoto(1, 0, false, -1)

	// Comprobar estado replica 2
	cfg.comprobarEstadoRemoto(2, 0, false, -1)

	// Parar réplicas almacenamiento en remoto
	cfg.stopDistributedProcesses()

	fmt.Println(".............", t.Name(), "Superado")
}
*/
