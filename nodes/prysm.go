package nodes

import (
    "fmt"
    "bytes"
    "github.com/protolambda/ztyp/codec"
    "strconv"
    "github.com/protolambda/zrnt/eth2/beacon"
    "github.com/protolambda/zrnt/eth2/configs"
)

// harcoded Prysm Default Values
var DefaultPrysmGRPCPort string = "3500"
var PrysmBase string = "/eth/v1alpha1"
var PrysmBSQuery string = "/debug/state?="

//NOTE: right now, the connection with the clients will be made through
// an insecure gRPC to a node that will be hosted on the localhost (by default)

// Struct that defines the needed and related information about the prysm client requested for the Chain Info 
type PrysmClient struct {
    Query string
    Ip  string
    Port string
    HeadSlot int
    FinalizedSlot int
    BeaconState beacon.Root
    Spec *beacon.Spec
}

// Generate a New Prysm Client
func NewPrysmClient(ip string, port string) *PrysmClient{
    fmt.Println("DEBUG, generating new Prysm Client")
    pClient := &PrysmClient{
            Ip: ip,
            Port: port,
            HeadSlot: 0,
            FinalizedSlot: 0,
            // Can not initialize the root / or initialize it to genesis
            Spec: configs.Mainnet,
        }
    return pClient
}

// Get Flat BeaconState by slot number from local client
func (c *PrysmClient)GetFlatBeaconStateFromSlot(slot int) (*beacon.BeaconState,error){
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBSQuery + strconv.Itoa(slot)
    fmt.Println("Url to request the BeaconState:", url)
    var bstate beacon.BeaconState
    bodybytes, err := GetJSON(url)
    if err != nil {
        return &bstate, fmt.Errorf("ERROR reading response:", err)
    }
    breader := bytes.NewReader(bodybytes)
    // Deserialize the ssz_bytes into a beacon.BeaconState
    bstate.Deserialize(configs.Mainnet, codec.NewDecodingReader(breader, uint64(len(bodybytes))))
    fmt.Println("DEBUG, Successfully requested the FlatBeaconState,", slot, "\n")
    return &bstate, nil
}

// Get BeaconStateView by slot number from local client
func (c *PrysmClient)GetBeaconStateViewFromSlot(slot int) (*beacon.BeaconStateView, error) {
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBSQuery + strconv.Itoa(slot)
    fmt.Println("Url to request the BeaconState:", url)
    var bstate *beacon.BeaconStateView
    bodybytes, err := GetJSON(url)
    if err != nil {
        return bstate, fmt.Errorf("ERROR reading response:", err)
    }
    breader := bytes.NewReader(bodybytes)
    // Deserialize the ssz_bytes into a beacon.BeaconState
    bstate, err = beacon.AsBeaconStateView(c.Spec.BeaconState().Deserialize(codec.NewDecodingReader(breader, uint64(len(bodybytes)))))
    fmt.Println("DEBUG, Successfully requested the BeaconStateView,", slot, "\n")
    return bstate, nil
}
