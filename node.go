package main

import (
//    "fmt"
    "github.com/protolambda/zrnt/eth2/beacon"

)


type Client interface {
//    GetBeaconStateFromRoot(beacon.Root) beacon.BeaconState
    GetBeaconStateFromSlot(int)  beacon.BeaconState

}





