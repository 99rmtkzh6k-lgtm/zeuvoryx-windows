param(
    [string]$Workflow = (Join-Path $PSScriptRoot '../../.github/workflows/windows-build.yml'),
    [string]$EvidencePath
)

$ErrorActionPreference = 'Stop'
$source = Get-Content -LiteralPath $Workflow -Raw
$operations = @{}
foreach ($name in 'network-audit','folder-audit','asr-audit','realtime-on') {
    $pattern = '(?m)^\s*"' + [regex]::Escape($name) + '": \{`([^`]+)`'
    $match = [regex]::Match($source, $pattern)
    if (!$match.Success) { throw "Missing operation: $name" }
    $operations[$name] = [scriptblock]::Create($match.Groups[1].Value)
}
$audit = [regex]::Match($source, '(?s)func windowsSecurityAuditContext.*?script := `([^`]+)`')
if (!$audit.Success) { throw 'Missing read-only Windows audit' }
[void][scriptblock]::Create($audit.Groups[1].Value)

function Get-MpPreference { $global:pref }
function Set-MpPreference {
 param($EnableNetworkProtection,$EnableControlledFolderAccess,$DisableRealtimeMonitoring)
 $global:writes++
 if($global:managed){return}
 if($EnableNetworkProtection -eq 'AuditMode'){$global:pref.EnableNetworkProtection=2}
 if($EnableControlledFolderAccess -eq 'AuditMode'){$global:pref.EnableControlledFolderAccess=2}
 if($DisableRealtimeMonitoring -eq $false){$global:status.RealTimeProtectionEnabled=$true}
}
function Add-MpPreference {
 param($AttackSurfaceReductionRules_Ids,$AttackSurfaceReductionRules_Actions)
 $global:writes++
 if($AttackSurfaceReductionRules_Actions -ne 'AuditMode'){throw 'unexpected action'}
 $global:pref.AttackSurfaceReductionRules_Ids+= $AttackSurfaceReductionRules_Ids
 $global:pref.AttackSurfaceReductionRules_Actions+= 2
}
function Get-MpComputerStatus { $global:status.Clone() }
function Reset-Test {
 $global:writes=0
 $global:managed=$false
 $global:pref=@{EnableNetworkProtection=0;EnableControlledFolderAccess=0;AttackSurfaceReductionRules_Ids=@('keep');AttackSurfaceReductionRules_Actions=@(1)}
 $global:status=@{AntivirusEnabled=$true;AMRunningMode='Normal';RealTimeProtectionEnabled=$false}
}
function AssertRefused($name,$expectedWrites) {
 $refused=$false
 try { & $operations[$name] | Out-Null } catch { $refused=$true }
 if(-not $refused -or $global:writes -ne $expectedWrites){throw "Unexpected outcome for $name"}
}
$cycles=@(for($i=1;$i -le 5;$i++){
 foreach($name in 'network-audit','folder-audit','asr-audit','realtime-on'){
  Reset-Test
  $result=& $operations[$name]
  if($global:writes -ne 1 -or !$result.after){throw "no confirmation: $name"}
  & $operations[$name] | Out-Null
  if($global:writes -ne 1){throw "not idempotent: $name"}
 }
 Reset-Test
 $global:pref.EnableNetworkProtection=1
 AssertRefused network-audit 0
 Reset-Test
 $global:pref.EnableControlledFolderAccess=1
 AssertRefused folder-audit 0
 Reset-Test
 $global:pref.AttackSurfaceReductionRules_Ids=@('d4f940ab-401b-4efc-aadc-ad5f3c50688a')
 AssertRefused asr-audit 0
 Reset-Test
 $global:status.AMRunningMode='Passive'
 AssertRefused realtime-on 0
 Reset-Test
 $global:managed=$true
 AssertRefused network-audit 1
 [ordered]@{cycle=$i;confirmed_and_idempotent=4;refusal_or_no_confirmation=5;status='passed'}
})
$result = [ordered]@{suite='Defender script contracts with mocked cmdlets';native_windows_tested=$false;cycles=$cycles}|ConvertTo-Json -Depth 5
if ($EvidencePath) { $result | Set-Content -LiteralPath $EvidencePath -Encoding utf8 }
$result
