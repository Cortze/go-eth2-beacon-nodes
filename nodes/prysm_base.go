package nodes

import (
	"fmt"

	beacon "github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/zrnt/eth2/configs"
)

//NOTE: right now, the connection with the clients will be made through
// an insecure gRPC to a node that will be hosted on the localhost (by default)

// Struct that defines the needed and related information about the prysm client requested for the Chain Info
type PrysmClient struct {
	Query         string
	Ip            string
	Port          string
	HeadSlot      int
	FinalizedSlot int
	BeaconState   beacon.Root
	Spec          *beacon.Spec
}

// Generate a New Prysm Client
func NewPrysmClient(ip string, port string) PrysmClient {
	fmt.Println("DEBUG, generating new Prysm Client")
	pClient := PrysmClient{
		Ip:            ip,
		Port:          port,
		HeadSlot:      0,
		FinalizedSlot: 0,
		// Can not initialize the root / or initialize it to genesis
		Spec: configs.Mainnet,
	}
	return pClient
}
