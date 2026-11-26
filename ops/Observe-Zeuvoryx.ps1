[CmdletBinding()]
param(
    [ValidateSet('Observe','Register','Unregister')][string]$Mode = 'Observe',
    [switch]$Once
)

$ErrorActionPreference = 'Stop'
$taskName = 'ZEUVORYX LAB local observer'
if ($Mode -ne 'Observe') {
    $identity = [Security.Principal.WindowsIdentity]::GetCurrent().Name
    if ($Mode -eq 'Unregister') {
        Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
        return
    }
    $powershell = Join-Path $PSHOME 'powershell.exe'
    if (!(Test-Path -LiteralPath $powershell)) {
        $powershell = Join-Path $env:SystemRoot 'System32\WindowsPowerShell\v1.0\powershell.exe'
    }
    $action = New-ScheduledTaskAction -Execute $powershell -Argument (
        '-NoProfile -NonInteractive -WindowStyle Hidden -File "' + $PSCommandPath + '"'
    )
    $trigger = New-ScheduledTaskTrigger -AtLogOn -User $identity
    $principal = New-ScheduledTaskPrincipal -UserId $identity -LogonType Interactive -RunLevel Limited
    $settings = New-ScheduledTaskSettingsSet -MultipleInstances IgnoreNew -ExecutionTimeLimit ([timespan]::Zero)
    Register-ScheduledTask -TaskName $taskName -Action $action -Trigger $trigger -Principal $principal -Settings $settings -Force | Out-Null
    Start-ScheduledTask -TaskName $taskName
    Get-ScheduledTask -TaskName $taskName | Select-Object TaskName, State
    return
}

function Get-LocalSnapshot {
    $process = @(
        Get-Process -Name ZEUVORYX,clamscan,freshclam -ErrorAction SilentlyContinue |
            ForEach-Object {
                try {
                    $path = $_.Path
                    if (!$path -or !$path.StartsWith($installation + [IO.Path]::DirectorySeparatorChar, [StringComparison]::OrdinalIgnoreCase)) {
                        return
                    }
                    [ordered]@{
                        name = $_.ProcessName
                        pid = $_.Id
                        cpu_seconds = $_.TotalProcessorTime.TotalSeconds
                        working_set_bytes = $_.WorkingSet64
                    }
                } catch {
                    # The process may have exited between enumeration and sampling.
                }
            }
    )
    $reports = Get-ChildItem -LiteralPath $audit -Filter 'report-*.json' -File -ErrorAction SilentlyContinue
    $latest = $reports | Sort-Object LastWriteTimeUtc -Descending | Select-Object -First 1
    $summary = $null
    $reportState = 'absent'
    if ($latest) {
        try {
            $result = Get-Content -LiteralPath $latest.FullName -Raw | ConvertFrom-Json
            $reportState = 'read'
            $summary = [ordered]@{
                written_utc = $latest.LastWriteTimeUtc.ToString('o')
                version = $result.engine_version
                commit = $result.build_commit
                files = $result.files
                bytes = $result.bytes
                errors = $result.error_count
                signature_state = $result.signature_state
                signature_activity = $result.signature_activity
                partial = $result.partial
            }
        } catch {
            $reportState = 'unreadable'
        }
    }
    [ordered]@{
        observed_at_utc = [datetime]::UtcNow.ToString('o')
        processes = $process
        saved_report_state = $reportState
        latest_saved_report = $summary
    }
}

$installation = Join-Path $env:LOCALAPPDATA 'Programs\ZEUVORYX-Lab303'
$root = Join-Path $env:APPDATA 'ZEUVORYX'
$audit = Join-Path $root 'lab305-audit'
$out = Join-Path $root 'operational-observer'
New-Item -Path $out -ItemType Directory -Force | Out-Null
$log = Join-Path $out 'activity.jsonl'
$lockPath = Join-Path $out 'observer.lock'
try {
    $lock = [IO.File]::Open($lockPath, [IO.FileMode]::OpenOrCreate, [IO.FileAccess]::ReadWrite, [IO.FileShare]::None)
} catch [IO.IOException] {
    Write-Warning 'O observador já está ativo ou o arquivo de trava está indisponível.'
    if ($Once) { throw }
    return
}
try {
    do {
        $line = Get-LocalSnapshot | ConvertTo-Json -Depth 5 -Compress
        try {
            if ((Test-Path -LiteralPath $log) -and (Get-Item -LiteralPath $log).Length -ge 4MB) {
                Move-Item -LiteralPath $log -Destination (Join-Path $out 'activity.previous.jsonl') -Force
            }
            Add-Content -LiteralPath $log -Value $line -Encoding UTF8
        } catch {
            Write-Warning 'O observador não conseguiu gravar a amostra local.'
            if ($Once) { throw }
        }
        if (!$Once) { Start-Sleep -Seconds 30 }
    } while (!$Once)
} finally {
    $lock.Dispose()
}
