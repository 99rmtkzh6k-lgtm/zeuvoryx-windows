// Package remediation is a fail-closed decision gate, not an execution engine.
package remediation
type Action string
const(Observe Action="observe"; Quarantine Action="quarantine"; StopProcess Action="stop_process"; CleanPersistence Action="clean_persistence"; Repair Action="repair")
type Evidence struct{
 Verified bool
 HeuristicOnly bool
 CriticalOrUnknown bool
 CurrentIdentityVerified bool
 UserConsent bool
 RecoveryReady bool
 TestArtifact bool
}
func Allowed(action Action,e Evidence)bool{
 if action==Observe{return true}
 if !e.Verified||e.HeuristicOnly||e.CriticalOrUnknown||!e.CurrentIdentityVerified||!e.UserConsent||!e.RecoveryReady{return false}
 // Only reversible quarantine may be planned in the lab. Other actions remain disabled.
 return action==Quarantine
}
