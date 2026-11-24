# ZEUVORYX Antivirus — Help Center

Public support knowledge base for ZEUVORYX Windows 3.0.6 LAB R1, an experimental complementary security candidate. This is not a production release, independent certification, or a replacement for Microsoft Defender.

[Português](pt-br.md) · [English](en.md) · [Español](es.md) · [日本語](ja.md)

## What is ZEUVORYX?

A Windows x64 security scanner written in Go, with file analysis, SHA-256 indicators, heuristics, reversible quarantine and experimental advisory AI. Macintosh 3.0.5 LAB is a separate candidate under validation.

## Installation and updates

Use only a test package authorized by the maintainer; verify the version and SHA-256 against the GitHub reference. Extract the ZIP, finish or cancel scanning, close the app, back up settings and quarantine, then install over a compatible LAB edition. Keep Defender, SmartScreen and other protections enabled. CI checked the 3.0.5 LAB to 3.0.6 LAB upgrade; other migrations are not proven.

## App will not open or closes unexpectedly

Do not repeatedly reinstall. Confirm version, Windows x64, Microsoft Edge WebView2 Runtime availability, free disk space and memory. Check Reliability Monitor (perfmon /rel) and Event Viewer (eventvwr.msc), Application and System logs around the failure. Do not delete quarantine or disable Windows security. Share only anonymized information via support.

## Full scan freezes or exits

Unexpected exits on physical PCs remain under investigation. Version 3.0.6 includes startup changes; CI observed the installed UI staying open for 15 seconds after its readiness signal, and 60 small controlled corpus scan tests passed. This does not establish full-disk stability. If it recurs, stop testing on your work PC and record time, version, scan mode, CPU/RAM and last progress. Do not run suspicious files.

## Errors, skipped files and alerts

Access denied, files locked by other processes and disappearing files can prevent reading. A partial report or zero detections does not prove a computer is clean. A heuristic or advisory AI score is a signal, not confirmed malware. Do not remove system files solely because of an alert; review evidence and keep system protections enabled.

## AI, real-time protection, ClamAV and updates

Current AI is advisory, without validated continuous local learning. Independent resident protection, continuous learning, production automatic updates and global statistics remain under development. ClamAV is optional and requires a separate installation. Component self-test success does not establish a working worldwide service. No production release has been announced.

## How to request support

Open a GitHub issue with the app version, Windows 10/11 build, architecture, approximate RAM, scan mode, timestamp and time zone, expected and observed result, reproduction steps, and error category or exit code if available. Remove personal names, personal paths, emails, tokens, identifiers, file contents and private file hashes. Do not post malware samples, raw logs, quarantine files or credentials. No guaranteed 24-hour support.

---

- [GitHub](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows) · [Issues / support](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/issues)
- [Windows 3.0.6 LAB R1 CI evidence](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45479946405)
- [Historical memorandum](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/MEMORANDO_DESENVOLVIMENTO.md)
- Updated: 2026-11-24.