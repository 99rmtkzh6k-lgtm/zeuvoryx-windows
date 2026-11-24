# ZEUVORYX Global Threat Intelligence — LAB

Source inventory, 24/11/2026. Advisory only; do not claim protection from unimported data.

| Source | Integration | Conditions |
| --- | --- | --- |
| MalwareBazaar | Existing authenticated recent SHA-256 query (100 records) and local hash import | Auth-Key, verify fair use and redistribution before bundling |
| ThreatFox | Planned authenticated recent IOC adapter with provenance, expiry and separate trust level | Auth-Key, verify license; expired indicators must be removed; do not equate shared abuse.ch provenance to independent corroboration |
| URLhaus | Existing local exact URL list import | Official API/export authentication and usage terms; never visit indicator URL |
| ClamAV / Cisco Talos | Existing optional separate installed engine; use official FreshClam for signed CVD | No bundled database or executable without licensing and package integrity review |
| Spamhaus DROP | New IPv4/IPv6 HTTPS retrieval, source attribution, timestamp, bounded parse and local IP match | User explicitly refreshes; in-memory only, no firewall change, no redistribution in this build |
| CISA KEV | Existing catalog CVE lookup | Not malware signatures or installed vulnerability detection |
| VirusTotal | Evaluation only | Public API prohibited in commercial products; do not integrate as free antivirus service |

Official references: https://bazaar.abuse.ch/api/ ; https://threatfox.abuse.ch/api/ ; https://urlhaus.abuse.ch/api/ ; https://docs.clamav.net/manual/Usage/SignatureManagement.html ; https://www.spamhaus.org/faqs/do-not-route-or-peer-drop/ ; https://docs.virustotal.com/reference/intro/getting-started ; https://www.cisa.gov/known-exploited-vulnerabilities-catalog .

No malware binaries, executable signatures, user files, credentials or arbitrary target requests. No fabricated initial records. DROP test fixtures use documentation networks, not actual indicators; no claims of live import until verified. Current DROP session refresh requires user action, 30-minute minimum interval, 26-hour maximum age, source dates/credits preserved. Last successful session data retained when refresh fails; expired data cannot match. Persistent authenticated snapshots, scheduled opt-in refresh, differential updates, ThreatFox/URLhaus API adapters, signed provenance, revocation, multiple independent confirmations and consent-based reversible quarantine remain to implement. Do not modify firewall or perform automatic deletion. Retain Defender and existing conservative remediation gate. No release promotion.


### Correção de política de frequência e base inicial
A FAQ oficial indica uma consulta diária suficiente e desaconselha downloads frequentes; adotar 24h por sessão para DROP, não o intervalo geral de 30 min. Uma nova sessão sem base ainda não tem limitação persistente; necessário cache persistente antes de distribuição. Workflow manual separado de coleta inicial somente em ambiente autorizado, sem download de malware, preserva bytes originais, atribuição e datas. O código suporta seed embutido opcional validado com prazo; sem arquivos reais incorporados o estado inicial continua vazio. Não reutilizar dados vencidos. Não transferir seed para build público sem revisão de termos e aprovação. https://www.spamhaus.org/faqs/do-not-route-or-peer-drop/ .
