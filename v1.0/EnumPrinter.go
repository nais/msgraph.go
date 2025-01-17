// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// PrinterFeedOrientation undocumented
type PrinterFeedOrientation string

const (
	// PrinterFeedOrientationVLongEdgeFirst undocumented
	PrinterFeedOrientationVLongEdgeFirst PrinterFeedOrientation = "longEdgeFirst"
	// PrinterFeedOrientationVShortEdgeFirst undocumented
	PrinterFeedOrientationVShortEdgeFirst PrinterFeedOrientation = "shortEdgeFirst"
	// PrinterFeedOrientationVUnknownFutureValue undocumented
	PrinterFeedOrientationVUnknownFutureValue PrinterFeedOrientation = "unknownFutureValue"
)

var (
	// PrinterFeedOrientationPLongEdgeFirst is a pointer to PrinterFeedOrientationVLongEdgeFirst
	PrinterFeedOrientationPLongEdgeFirst = &_PrinterFeedOrientationPLongEdgeFirst
	// PrinterFeedOrientationPShortEdgeFirst is a pointer to PrinterFeedOrientationVShortEdgeFirst
	PrinterFeedOrientationPShortEdgeFirst = &_PrinterFeedOrientationPShortEdgeFirst
	// PrinterFeedOrientationPUnknownFutureValue is a pointer to PrinterFeedOrientationVUnknownFutureValue
	PrinterFeedOrientationPUnknownFutureValue = &_PrinterFeedOrientationPUnknownFutureValue
)

var (
	_PrinterFeedOrientationPLongEdgeFirst      = PrinterFeedOrientationVLongEdgeFirst
	_PrinterFeedOrientationPShortEdgeFirst     = PrinterFeedOrientationVShortEdgeFirst
	_PrinterFeedOrientationPUnknownFutureValue = PrinterFeedOrientationVUnknownFutureValue
)

// PrinterProcessingState undocumented
type PrinterProcessingState string

const (
	// PrinterProcessingStateVUnknown undocumented
	PrinterProcessingStateVUnknown PrinterProcessingState = "unknown"
	// PrinterProcessingStateVIdle undocumented
	PrinterProcessingStateVIdle PrinterProcessingState = "idle"
	// PrinterProcessingStateVProcessing undocumented
	PrinterProcessingStateVProcessing PrinterProcessingState = "processing"
	// PrinterProcessingStateVStopped undocumented
	PrinterProcessingStateVStopped PrinterProcessingState = "stopped"
	// PrinterProcessingStateVUnknownFutureValue undocumented
	PrinterProcessingStateVUnknownFutureValue PrinterProcessingState = "unknownFutureValue"
)

var (
	// PrinterProcessingStatePUnknown is a pointer to PrinterProcessingStateVUnknown
	PrinterProcessingStatePUnknown = &_PrinterProcessingStatePUnknown
	// PrinterProcessingStatePIdle is a pointer to PrinterProcessingStateVIdle
	PrinterProcessingStatePIdle = &_PrinterProcessingStatePIdle
	// PrinterProcessingStatePProcessing is a pointer to PrinterProcessingStateVProcessing
	PrinterProcessingStatePProcessing = &_PrinterProcessingStatePProcessing
	// PrinterProcessingStatePStopped is a pointer to PrinterProcessingStateVStopped
	PrinterProcessingStatePStopped = &_PrinterProcessingStatePStopped
	// PrinterProcessingStatePUnknownFutureValue is a pointer to PrinterProcessingStateVUnknownFutureValue
	PrinterProcessingStatePUnknownFutureValue = &_PrinterProcessingStatePUnknownFutureValue
)

var (
	_PrinterProcessingStatePUnknown            = PrinterProcessingStateVUnknown
	_PrinterProcessingStatePIdle               = PrinterProcessingStateVIdle
	_PrinterProcessingStatePProcessing         = PrinterProcessingStateVProcessing
	_PrinterProcessingStatePStopped            = PrinterProcessingStateVStopped
	_PrinterProcessingStatePUnknownFutureValue = PrinterProcessingStateVUnknownFutureValue
)

// PrinterProcessingStateDetail undocumented
type PrinterProcessingStateDetail string

const (
	// PrinterProcessingStateDetailVPaused undocumented
	PrinterProcessingStateDetailVPaused PrinterProcessingStateDetail = "paused"
	// PrinterProcessingStateDetailVMediaJam undocumented
	PrinterProcessingStateDetailVMediaJam PrinterProcessingStateDetail = "mediaJam"
	// PrinterProcessingStateDetailVMediaNeeded undocumented
	PrinterProcessingStateDetailVMediaNeeded PrinterProcessingStateDetail = "mediaNeeded"
	// PrinterProcessingStateDetailVMediaLow undocumented
	PrinterProcessingStateDetailVMediaLow PrinterProcessingStateDetail = "mediaLow"
	// PrinterProcessingStateDetailVMediaEmpty undocumented
	PrinterProcessingStateDetailVMediaEmpty PrinterProcessingStateDetail = "mediaEmpty"
	// PrinterProcessingStateDetailVCoverOpen undocumented
	PrinterProcessingStateDetailVCoverOpen PrinterProcessingStateDetail = "coverOpen"
	// PrinterProcessingStateDetailVInterlockOpen undocumented
	PrinterProcessingStateDetailVInterlockOpen PrinterProcessingStateDetail = "interlockOpen"
	// PrinterProcessingStateDetailVOutputTrayMissing undocumented
	PrinterProcessingStateDetailVOutputTrayMissing PrinterProcessingStateDetail = "outputTrayMissing"
	// PrinterProcessingStateDetailVOutputAreaFull undocumented
	PrinterProcessingStateDetailVOutputAreaFull PrinterProcessingStateDetail = "outputAreaFull"
	// PrinterProcessingStateDetailVMarkerSupplyLow undocumented
	PrinterProcessingStateDetailVMarkerSupplyLow PrinterProcessingStateDetail = "markerSupplyLow"
	// PrinterProcessingStateDetailVMarkerSupplyEmpty undocumented
	PrinterProcessingStateDetailVMarkerSupplyEmpty PrinterProcessingStateDetail = "markerSupplyEmpty"
	// PrinterProcessingStateDetailVInputTrayMissing undocumented
	PrinterProcessingStateDetailVInputTrayMissing PrinterProcessingStateDetail = "inputTrayMissing"
	// PrinterProcessingStateDetailVOutputAreaAlmostFull undocumented
	PrinterProcessingStateDetailVOutputAreaAlmostFull PrinterProcessingStateDetail = "outputAreaAlmostFull"
	// PrinterProcessingStateDetailVMarkerWasteAlmostFull undocumented
	PrinterProcessingStateDetailVMarkerWasteAlmostFull PrinterProcessingStateDetail = "markerWasteAlmostFull"
	// PrinterProcessingStateDetailVMarkerWasteFull undocumented
	PrinterProcessingStateDetailVMarkerWasteFull PrinterProcessingStateDetail = "markerWasteFull"
	// PrinterProcessingStateDetailVFuserOverTemp undocumented
	PrinterProcessingStateDetailVFuserOverTemp PrinterProcessingStateDetail = "fuserOverTemp"
	// PrinterProcessingStateDetailVFuserUnderTemp undocumented
	PrinterProcessingStateDetailVFuserUnderTemp PrinterProcessingStateDetail = "fuserUnderTemp"
	// PrinterProcessingStateDetailVOther undocumented
	PrinterProcessingStateDetailVOther PrinterProcessingStateDetail = "other"
	// PrinterProcessingStateDetailVNone undocumented
	PrinterProcessingStateDetailVNone PrinterProcessingStateDetail = "none"
	// PrinterProcessingStateDetailVMovingToPaused undocumented
	PrinterProcessingStateDetailVMovingToPaused PrinterProcessingStateDetail = "movingToPaused"
	// PrinterProcessingStateDetailVShutdown undocumented
	PrinterProcessingStateDetailVShutdown PrinterProcessingStateDetail = "shutdown"
	// PrinterProcessingStateDetailVConnectingToDevice undocumented
	PrinterProcessingStateDetailVConnectingToDevice PrinterProcessingStateDetail = "connectingToDevice"
	// PrinterProcessingStateDetailVTimedOut undocumented
	PrinterProcessingStateDetailVTimedOut PrinterProcessingStateDetail = "timedOut"
	// PrinterProcessingStateDetailVStopping undocumented
	PrinterProcessingStateDetailVStopping PrinterProcessingStateDetail = "stopping"
	// PrinterProcessingStateDetailVStoppedPartially undocumented
	PrinterProcessingStateDetailVStoppedPartially PrinterProcessingStateDetail = "stoppedPartially"
	// PrinterProcessingStateDetailVTonerLow undocumented
	PrinterProcessingStateDetailVTonerLow PrinterProcessingStateDetail = "tonerLow"
	// PrinterProcessingStateDetailVTonerEmpty undocumented
	PrinterProcessingStateDetailVTonerEmpty PrinterProcessingStateDetail = "tonerEmpty"
	// PrinterProcessingStateDetailVSpoolAreaFull undocumented
	PrinterProcessingStateDetailVSpoolAreaFull PrinterProcessingStateDetail = "spoolAreaFull"
	// PrinterProcessingStateDetailVDoorOpen undocumented
	PrinterProcessingStateDetailVDoorOpen PrinterProcessingStateDetail = "doorOpen"
	// PrinterProcessingStateDetailVOpticalPhotoConductorNearEndOfLife undocumented
	PrinterProcessingStateDetailVOpticalPhotoConductorNearEndOfLife PrinterProcessingStateDetail = "opticalPhotoConductorNearEndOfLife"
	// PrinterProcessingStateDetailVOpticalPhotoConductorLifeOver undocumented
	PrinterProcessingStateDetailVOpticalPhotoConductorLifeOver PrinterProcessingStateDetail = "opticalPhotoConductorLifeOver"
	// PrinterProcessingStateDetailVDeveloperLow undocumented
	PrinterProcessingStateDetailVDeveloperLow PrinterProcessingStateDetail = "developerLow"
	// PrinterProcessingStateDetailVDeveloperEmpty undocumented
	PrinterProcessingStateDetailVDeveloperEmpty PrinterProcessingStateDetail = "developerEmpty"
	// PrinterProcessingStateDetailVInterpreterResourceUnavailable undocumented
	PrinterProcessingStateDetailVInterpreterResourceUnavailable PrinterProcessingStateDetail = "interpreterResourceUnavailable"
	// PrinterProcessingStateDetailVUnknownFutureValue undocumented
	PrinterProcessingStateDetailVUnknownFutureValue PrinterProcessingStateDetail = "unknownFutureValue"
)

var (
	// PrinterProcessingStateDetailPPaused is a pointer to PrinterProcessingStateDetailVPaused
	PrinterProcessingStateDetailPPaused = &_PrinterProcessingStateDetailPPaused
	// PrinterProcessingStateDetailPMediaJam is a pointer to PrinterProcessingStateDetailVMediaJam
	PrinterProcessingStateDetailPMediaJam = &_PrinterProcessingStateDetailPMediaJam
	// PrinterProcessingStateDetailPMediaNeeded is a pointer to PrinterProcessingStateDetailVMediaNeeded
	PrinterProcessingStateDetailPMediaNeeded = &_PrinterProcessingStateDetailPMediaNeeded
	// PrinterProcessingStateDetailPMediaLow is a pointer to PrinterProcessingStateDetailVMediaLow
	PrinterProcessingStateDetailPMediaLow = &_PrinterProcessingStateDetailPMediaLow
	// PrinterProcessingStateDetailPMediaEmpty is a pointer to PrinterProcessingStateDetailVMediaEmpty
	PrinterProcessingStateDetailPMediaEmpty = &_PrinterProcessingStateDetailPMediaEmpty
	// PrinterProcessingStateDetailPCoverOpen is a pointer to PrinterProcessingStateDetailVCoverOpen
	PrinterProcessingStateDetailPCoverOpen = &_PrinterProcessingStateDetailPCoverOpen
	// PrinterProcessingStateDetailPInterlockOpen is a pointer to PrinterProcessingStateDetailVInterlockOpen
	PrinterProcessingStateDetailPInterlockOpen = &_PrinterProcessingStateDetailPInterlockOpen
	// PrinterProcessingStateDetailPOutputTrayMissing is a pointer to PrinterProcessingStateDetailVOutputTrayMissing
	PrinterProcessingStateDetailPOutputTrayMissing = &_PrinterProcessingStateDetailPOutputTrayMissing
	// PrinterProcessingStateDetailPOutputAreaFull is a pointer to PrinterProcessingStateDetailVOutputAreaFull
	PrinterProcessingStateDetailPOutputAreaFull = &_PrinterProcessingStateDetailPOutputAreaFull
	// PrinterProcessingStateDetailPMarkerSupplyLow is a pointer to PrinterProcessingStateDetailVMarkerSupplyLow
	PrinterProcessingStateDetailPMarkerSupplyLow = &_PrinterProcessingStateDetailPMarkerSupplyLow
	// PrinterProcessingStateDetailPMarkerSupplyEmpty is a pointer to PrinterProcessingStateDetailVMarkerSupplyEmpty
	PrinterProcessingStateDetailPMarkerSupplyEmpty = &_PrinterProcessingStateDetailPMarkerSupplyEmpty
	// PrinterProcessingStateDetailPInputTrayMissing is a pointer to PrinterProcessingStateDetailVInputTrayMissing
	PrinterProcessingStateDetailPInputTrayMissing = &_PrinterProcessingStateDetailPInputTrayMissing
	// PrinterProcessingStateDetailPOutputAreaAlmostFull is a pointer to PrinterProcessingStateDetailVOutputAreaAlmostFull
	PrinterProcessingStateDetailPOutputAreaAlmostFull = &_PrinterProcessingStateDetailPOutputAreaAlmostFull
	// PrinterProcessingStateDetailPMarkerWasteAlmostFull is a pointer to PrinterProcessingStateDetailVMarkerWasteAlmostFull
	PrinterProcessingStateDetailPMarkerWasteAlmostFull = &_PrinterProcessingStateDetailPMarkerWasteAlmostFull
	// PrinterProcessingStateDetailPMarkerWasteFull is a pointer to PrinterProcessingStateDetailVMarkerWasteFull
	PrinterProcessingStateDetailPMarkerWasteFull = &_PrinterProcessingStateDetailPMarkerWasteFull
	// PrinterProcessingStateDetailPFuserOverTemp is a pointer to PrinterProcessingStateDetailVFuserOverTemp
	PrinterProcessingStateDetailPFuserOverTemp = &_PrinterProcessingStateDetailPFuserOverTemp
	// PrinterProcessingStateDetailPFuserUnderTemp is a pointer to PrinterProcessingStateDetailVFuserUnderTemp
	PrinterProcessingStateDetailPFuserUnderTemp = &_PrinterProcessingStateDetailPFuserUnderTemp
	// PrinterProcessingStateDetailPOther is a pointer to PrinterProcessingStateDetailVOther
	PrinterProcessingStateDetailPOther = &_PrinterProcessingStateDetailPOther
	// PrinterProcessingStateDetailPNone is a pointer to PrinterProcessingStateDetailVNone
	PrinterProcessingStateDetailPNone = &_PrinterProcessingStateDetailPNone
	// PrinterProcessingStateDetailPMovingToPaused is a pointer to PrinterProcessingStateDetailVMovingToPaused
	PrinterProcessingStateDetailPMovingToPaused = &_PrinterProcessingStateDetailPMovingToPaused
	// PrinterProcessingStateDetailPShutdown is a pointer to PrinterProcessingStateDetailVShutdown
	PrinterProcessingStateDetailPShutdown = &_PrinterProcessingStateDetailPShutdown
	// PrinterProcessingStateDetailPConnectingToDevice is a pointer to PrinterProcessingStateDetailVConnectingToDevice
	PrinterProcessingStateDetailPConnectingToDevice = &_PrinterProcessingStateDetailPConnectingToDevice
	// PrinterProcessingStateDetailPTimedOut is a pointer to PrinterProcessingStateDetailVTimedOut
	PrinterProcessingStateDetailPTimedOut = &_PrinterProcessingStateDetailPTimedOut
	// PrinterProcessingStateDetailPStopping is a pointer to PrinterProcessingStateDetailVStopping
	PrinterProcessingStateDetailPStopping = &_PrinterProcessingStateDetailPStopping
	// PrinterProcessingStateDetailPStoppedPartially is a pointer to PrinterProcessingStateDetailVStoppedPartially
	PrinterProcessingStateDetailPStoppedPartially = &_PrinterProcessingStateDetailPStoppedPartially
	// PrinterProcessingStateDetailPTonerLow is a pointer to PrinterProcessingStateDetailVTonerLow
	PrinterProcessingStateDetailPTonerLow = &_PrinterProcessingStateDetailPTonerLow
	// PrinterProcessingStateDetailPTonerEmpty is a pointer to PrinterProcessingStateDetailVTonerEmpty
	PrinterProcessingStateDetailPTonerEmpty = &_PrinterProcessingStateDetailPTonerEmpty
	// PrinterProcessingStateDetailPSpoolAreaFull is a pointer to PrinterProcessingStateDetailVSpoolAreaFull
	PrinterProcessingStateDetailPSpoolAreaFull = &_PrinterProcessingStateDetailPSpoolAreaFull
	// PrinterProcessingStateDetailPDoorOpen is a pointer to PrinterProcessingStateDetailVDoorOpen
	PrinterProcessingStateDetailPDoorOpen = &_PrinterProcessingStateDetailPDoorOpen
	// PrinterProcessingStateDetailPOpticalPhotoConductorNearEndOfLife is a pointer to PrinterProcessingStateDetailVOpticalPhotoConductorNearEndOfLife
	PrinterProcessingStateDetailPOpticalPhotoConductorNearEndOfLife = &_PrinterProcessingStateDetailPOpticalPhotoConductorNearEndOfLife
	// PrinterProcessingStateDetailPOpticalPhotoConductorLifeOver is a pointer to PrinterProcessingStateDetailVOpticalPhotoConductorLifeOver
	PrinterProcessingStateDetailPOpticalPhotoConductorLifeOver = &_PrinterProcessingStateDetailPOpticalPhotoConductorLifeOver
	// PrinterProcessingStateDetailPDeveloperLow is a pointer to PrinterProcessingStateDetailVDeveloperLow
	PrinterProcessingStateDetailPDeveloperLow = &_PrinterProcessingStateDetailPDeveloperLow
	// PrinterProcessingStateDetailPDeveloperEmpty is a pointer to PrinterProcessingStateDetailVDeveloperEmpty
	PrinterProcessingStateDetailPDeveloperEmpty = &_PrinterProcessingStateDetailPDeveloperEmpty
	// PrinterProcessingStateDetailPInterpreterResourceUnavailable is a pointer to PrinterProcessingStateDetailVInterpreterResourceUnavailable
	PrinterProcessingStateDetailPInterpreterResourceUnavailable = &_PrinterProcessingStateDetailPInterpreterResourceUnavailable
	// PrinterProcessingStateDetailPUnknownFutureValue is a pointer to PrinterProcessingStateDetailVUnknownFutureValue
	PrinterProcessingStateDetailPUnknownFutureValue = &_PrinterProcessingStateDetailPUnknownFutureValue
)

var (
	_PrinterProcessingStateDetailPPaused                             = PrinterProcessingStateDetailVPaused
	_PrinterProcessingStateDetailPMediaJam                           = PrinterProcessingStateDetailVMediaJam
	_PrinterProcessingStateDetailPMediaNeeded                        = PrinterProcessingStateDetailVMediaNeeded
	_PrinterProcessingStateDetailPMediaLow                           = PrinterProcessingStateDetailVMediaLow
	_PrinterProcessingStateDetailPMediaEmpty                         = PrinterProcessingStateDetailVMediaEmpty
	_PrinterProcessingStateDetailPCoverOpen                          = PrinterProcessingStateDetailVCoverOpen
	_PrinterProcessingStateDetailPInterlockOpen                      = PrinterProcessingStateDetailVInterlockOpen
	_PrinterProcessingStateDetailPOutputTrayMissing                  = PrinterProcessingStateDetailVOutputTrayMissing
	_PrinterProcessingStateDetailPOutputAreaFull                     = PrinterProcessingStateDetailVOutputAreaFull
	_PrinterProcessingStateDetailPMarkerSupplyLow                    = PrinterProcessingStateDetailVMarkerSupplyLow
	_PrinterProcessingStateDetailPMarkerSupplyEmpty                  = PrinterProcessingStateDetailVMarkerSupplyEmpty
	_PrinterProcessingStateDetailPInputTrayMissing                   = PrinterProcessingStateDetailVInputTrayMissing
	_PrinterProcessingStateDetailPOutputAreaAlmostFull               = PrinterProcessingStateDetailVOutputAreaAlmostFull
	_PrinterProcessingStateDetailPMarkerWasteAlmostFull              = PrinterProcessingStateDetailVMarkerWasteAlmostFull
	_PrinterProcessingStateDetailPMarkerWasteFull                    = PrinterProcessingStateDetailVMarkerWasteFull
	_PrinterProcessingStateDetailPFuserOverTemp                      = PrinterProcessingStateDetailVFuserOverTemp
	_PrinterProcessingStateDetailPFuserUnderTemp                     = PrinterProcessingStateDetailVFuserUnderTemp
	_PrinterProcessingStateDetailPOther                              = PrinterProcessingStateDetailVOther
	_PrinterProcessingStateDetailPNone                               = PrinterProcessingStateDetailVNone
	_PrinterProcessingStateDetailPMovingToPaused                     = PrinterProcessingStateDetailVMovingToPaused
	_PrinterProcessingStateDetailPShutdown                           = PrinterProcessingStateDetailVShutdown
	_PrinterProcessingStateDetailPConnectingToDevice                 = PrinterProcessingStateDetailVConnectingToDevice
	_PrinterProcessingStateDetailPTimedOut                           = PrinterProcessingStateDetailVTimedOut
	_PrinterProcessingStateDetailPStopping                           = PrinterProcessingStateDetailVStopping
	_PrinterProcessingStateDetailPStoppedPartially                   = PrinterProcessingStateDetailVStoppedPartially
	_PrinterProcessingStateDetailPTonerLow                           = PrinterProcessingStateDetailVTonerLow
	_PrinterProcessingStateDetailPTonerEmpty                         = PrinterProcessingStateDetailVTonerEmpty
	_PrinterProcessingStateDetailPSpoolAreaFull                      = PrinterProcessingStateDetailVSpoolAreaFull
	_PrinterProcessingStateDetailPDoorOpen                           = PrinterProcessingStateDetailVDoorOpen
	_PrinterProcessingStateDetailPOpticalPhotoConductorNearEndOfLife = PrinterProcessingStateDetailVOpticalPhotoConductorNearEndOfLife
	_PrinterProcessingStateDetailPOpticalPhotoConductorLifeOver      = PrinterProcessingStateDetailVOpticalPhotoConductorLifeOver
	_PrinterProcessingStateDetailPDeveloperLow                       = PrinterProcessingStateDetailVDeveloperLow
	_PrinterProcessingStateDetailPDeveloperEmpty                     = PrinterProcessingStateDetailVDeveloperEmpty
	_PrinterProcessingStateDetailPInterpreterResourceUnavailable     = PrinterProcessingStateDetailVInterpreterResourceUnavailable
	_PrinterProcessingStateDetailPUnknownFutureValue                 = PrinterProcessingStateDetailVUnknownFutureValue
)
