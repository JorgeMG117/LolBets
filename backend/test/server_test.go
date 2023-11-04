package test

import (
	"fmt"
	"testing"
    "os"
	"time"

	"github.com/JorgeMG117/LolBets/backend/client"
	"github.com/JorgeMG117/LolBets/backend/server"
    "github.com/JorgeMG117/LolBets/backend/models"
)

type configDeploy struct {
	t           *testing.T
    numClients  int
}

// Crear una configuracion de despliegue
func makeCfgDespliegue(t *testing.T, numClients int) *configDeploy {
	cfg := &configDeploy{}
    cfg.t = t
    cfg.numClients = numClients

	return cfg
}

// Pruebas a realizar:
// . Comprobar que el cliente ve los partidos
// . Un cliente hace una apuesta el otro pide los partidos y puede la apuesta

func startServer(){
}


func TestBackend(t *testing.T) {
	t.Skip("SKIPPED soloArranqueYparadaTest1")
    // Inizializar BD
    // Añadir partidos predeterminados
    // Lanzar servidor
    go func() {
        if err := server.ExecServer(false); err != nil {
            fmt.Fprintf(os.Stderr, "%s\n", err)
            os.Exit(1)
        }
        fmt.Println("Lanzando servidor")
    }()
    time.Sleep(5 * time.Second)

    //chOrders[]
    // Lanzar numero de clientes predefinidos
        // Cada cliente tiene definida las acciones (apuestas) que debe realizar
    var stop chan bool
    var chBet chan models.Bet
    for i := 0; i < 2; i++ {
        go func(i int) {
            fmt.Println("Lanzando cliente: ", i)
            client.Client(stop, chBet)
        }(i)
    }

    time.Sleep(15 * time.Second)
    stop <- true

    //<- "Apostar"
    //<- "Pedir partidos"


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
