package nodes

import (
	//    "fmt"
	beacon "github.com/protolambda/zrnt/eth2/beacon/common"
)

type Client interface {
	//    GetBeaconStateFromRoot(beacon.Root) beacon.BeaconState
	GetBeaconStateFromSlot(int) beacon.BeaconState
}
