package nodes

import (
	"fmt"
	//    beacon "github.com/protolambda/zrnt/eth2/beacon/common"
)

func main() {
	prysm := NewPrysmClient("localhost", "3500")
	_, err := prysm.GetFlatBeaconStateFromSlot(195520)
	if err != nil {
		fmt.Println("ERROR,", err)
	}
	_, err = prysm.GetBeaconStateViewFromSlot(604435)
	if err != nil {
		fmt.Println("ERROR,", err)
	}

	fmt.Println("Ciao!")
}
