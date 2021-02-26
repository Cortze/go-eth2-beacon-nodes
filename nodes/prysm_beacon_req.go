package nodes

import (
    "fmt"
    "bytes"
    "strconv"
    "encoding/json"
    "github.com/protolambda/ztyp/codec"
    "github.com/protolambda/zrnt/eth2/beacon"
    "github.com/protolambda/zrnt/eth2/configs"
)

// ---------- BEACON REQUESTS -----------


// -- Beacon State --

// NOTE: Same fucntion, diferent, ones gives back the View of the state, the other one the State
// Get Flat BeaconState by slot number from local client
func (c *PrysmClient) GetFlatBeaconStateFromSlot(slot int) (*beacon.BeaconState,error){
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBSQuery + strconv.Itoa(slot)
    var bstate beacon.BeaconState
    bodybytes, err := GetSSZEncodedJSON(url)
    if err != nil {
        return &bstate, fmt.Errorf("ERROR reading response:", err)
    }
    breader := bytes.NewReader(bodybytes)
    // Deserialize the ssz_bytes into a beacon.BeaconState
    bstate.Deserialize(configs.Mainnet, codec.NewDecodingReader(breader, uint64(len(bodybytes))))
    return &bstate, nil
}

// Get BeaconStateView by slot number from local client
func (c *PrysmClient) GetBeaconStateViewFromSlot(slot int) (*beacon.BeaconStateView, error) {
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBSQuery + strconv.Itoa(slot)
    var bstate *beacon.BeaconStateView
    bodybytes, err := GetSSZEncodedJSON(url)
    if err != nil {
        return bstate, fmt.Errorf("ERROR reading response:", err)
    }
    breader := bytes.NewReader(bodybytes)
    // Deserialize the ssz_bytes into a beacon.BeaconState
    bstate, err = beacon.AsBeaconStateView(c.Spec.BeaconState().Deserialize(codec.NewDecodingReader(breader, uint64(len(bodybytes)))))
    return bstate, nil
}

// -- Beacon Block --
// response type from Prysm for the BeaconBlocks
type PrysmBeaconBlockContainer struct {
    BlockContainers []PSignedBlock  `json:"block"`
    NextPageToken   string          `json:"nextPageToken"`
    TotalBlocks     int             `json:"totalSize"`
}
type PSignedBlock struct {
    BeaconBlock beacon.BeaconBlock  `json:"block"`
    Signature   beacon.BLSSignature `json:"signature"`
    BlockRoot   beacon.Root         `json:"blockRoot"`
    Canonical   bool                `json:"canonical"`
}

// returns the Prysm version of the SignedBeaconBlock
func (c *PrysmClient) GetBeaconBlockFromSlot(slot int) (PSignedBlock, error){
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBBlockQuery + "slot=" + strconv.Itoa(slot)
    fmt.Println("Url:", url)
    var blocksResponse PrysmBeaconBlockContainer
    var psb PSignedBlock
    bodybytes, err := GetJSON(url)
    fmt.Println("Bodybytes:", bodybytes)
    if err != nil {
        return psb, fmt.Errorf("Error Getting the JSON from the API -> block from slot", err)
    }
    err = json.Unmarshal(bodybytes, &blocksResponse)
    if err != nil {
        return psb, fmt.Errorf("Error Unmarshalling the JSON from the API resposne -> block from slot", err)
    }
    fmt.Println("the BlockResponse", blockResponse)
    psb = blocksResponse.BlockContainers[0]
    return blocksResponse.BlockContainers[0], nil
}

// return List of blocks of the given epoch
func (c *PrysmClient) GetBeaconBlocksFromEpoch(epoch int) (PrysmBeaconBlockContainer, error){
    url := "http://" + c.Ip + ":" + c.Port + PrysmBase + PrysmBBlockQuery + "epoch=" + strconv.Itoa(epoch)
    fmt.Println("Url:", url)
    var pbbc PrysmBeaconBlockContainer
    bodybytes, err := GetJSON(url)
    fmt.Println("Bodybytes:", bodybytes)
    if err != nil {
        return pbbc, fmt.Errorf("Error Getting the Json From the API -> Blocks from epoch")
    }
    err = json.Unmarshal(bodybytes, &pbbc)
    if err != nil {
        return pbbc, fmt.Errorf("Error Unmarshalling the JSON from the API resposne -> block from epoch", err)
    }
    return pbbc, nil
}




// -- Beacon Commitees --

// Returns the Committees for the given Epoch
func (c *PrysmClient) GetCommitteesForEpoch(epoch string) (error) {
    return nil
}

// -- Beacon

