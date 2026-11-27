# ZEUVORYX 3.0.23 LAB — V3 audit / Build 52 candidate
Base 721378ff70a50043b9a1468e580aebe2c4b9b565. This branch is NOT approved, NOT a release and NOT a Windows physical test.
## Scope
Original eleven local games, integrated six-entry launcher including Robozinho, familygate and GameWindow. External gamecatalog/catalog.go intentionally untouched. Build 49, Build 50, main and ZEUS untouched.
## Root cause
User-reported video (not available to inspect in this turn): narrow left column, missing controls, no focus. Source code: legacy fixed board and frame heights, dual launcher generation paths. GameWindow candidate from Build 51 addresses shared frame but still requires CI and native physical validation.
## V3 findings
- familygate/policy.go previously returned generic 'Restrição do perfil familiar'; now identifies age, violence, sex, drugs, gambling or language and missing classification, without disclosing password or hashes.
- Both launchers previously returned 'Jogo indisponível' and a generic refusal. Now identify title, explicit rule, and console warning; authorization remains fail-closed. No data uploaded.
- prepare_games314.py anchor exceptions now include the expected anchor and observed count. Additional structural checks remain necessary for every generated insertion point.
- Fixed SHA for 2048 is now in approved_game_hashes.json with original source commit. This is an immutable review baseline, NOT an approval. Updating requires reviewing source, commit and updating manifest together. Retro catalog uses existing per-game Git blob hashes; mismatch remains fail-closed.
- report_anchor315.py already implements HTMLParser fallback and is tested by lab/test_report_anchor315.py; do not replace with brittle '</style></head>' search.
## Still unresolved / no invented evidence
- Verify Win32 host failure paths, window creation, native binding errors and visible error UI on a Windows runner and physical PC.
- Identify actual installer file used in user's video from file SHA or installer metadata; run/branch alone cannot prove it.
- Test 20/20 complete cycles; browser and native WebView2; per-game 30s, scoring, restart, keyboard, family restrictions; 1366x768,1600x900,1920x1080 at 100/125/150% DPI.
- Installer SHA256, physical per-game logs, Windows compatibility, final release and WIN32-LAB.md updates: pending.
- Do not weaken iframe sandbox with allow-same-origin for untrusted inline games. The browser host has privileged bindings.
