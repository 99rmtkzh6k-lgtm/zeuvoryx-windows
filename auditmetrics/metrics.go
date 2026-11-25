package auditmetrics

// SignatureEvidence reports observations, never an estimated malware detection rate.
type SignatureEvidence struct {
 Scanned *int64 `json:"scanned_files"`
 EngineInfected *int64 `json:"engine_reported_infected_files"`
 KnownSignatures *int64 `json:"known_signatures"`
 EngineErrors *int64 `json:"engine_reported_errors"`
 SignatureRecords int `json:"signature_records"`
 HeuristicOrLimitRecords int `json:"heuristic_or_limit_records"`
 TestRecords int `json:"test_records"`
 UnclassifiedRecords int `json:"unclassified_records"`
 ObservedRecords int `json:"observed_records"`
 ErrorCountKnown bool `json:"error_count_known"`
 SummaryUsable bool `json:"summary_usable"`
 Partial bool `json:"partial"`
}

func known(v int64) *int64 { if v<0{return nil};return &v }

func Signature(scanned, infected, signatures, errs int64, kinds []string, partial bool) SignatureEvidence {
 e:=SignatureEvidence{Scanned:known(scanned),EngineInfected:known(infected),KnownSignatures:known(signatures),EngineErrors:known(errs),ObservedRecords:len(kinds),ErrorCountKnown:errs>=0}
 for _,kind:=range kinds{switch kind{case "signature_detection":e.SignatureRecords++;case "heuristic_or_limit":e.HeuristicOrLimitRecords++;case "test_eicar":e.TestRecords++;default:e.UnclassifiedRecords++}}
 e.Partial=partial||scanned<=0||infected<0||infected>scanned||signatures<=0||errs>0||infected!=int64(len(kinds))||e.HeuristicOrLimitRecords>0||e.UnclassifiedRecords>0
 e.SummaryUsable=!e.Partial
 return e
}
