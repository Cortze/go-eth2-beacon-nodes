package nodes

import ()

// Harcoded Prysm Default Values
var DefaultPrysmGRPCPort string = "3500"
var PrysmBase string = "/eth/v1alpha1"
// -- Beacon --
var PrysmBSQuery string = "/debug/state?="
var PrysmBBlockQuery string = "/beacon/blocks?"
var PrysmBCommitteeQuery string = "/beacon/committees?="
var PrysmBChainHeadQuery string = "/beacon/chainhead"
var PrysmBConfigQuery string = "/beacon/config"
// -- Validator --
var PrysmValidatorQuery string = "/validator/"
// -- Validators --
var PrysmValidatorsQuery string = "/validators"
var PrysmVBalancesQuery string = "/validators/balances"
var PrysmVPerformanceQuery string = "/validators/performance"
var PrysmVAssignmentsQuery string = "/validators/assignments"


