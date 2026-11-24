# ZEUVORYX Macintosh 3.0.3 LAB — macOS

Scanner complementar experimental, com aplicativo nativo Swift/SwiftUI e motor Go local somente leitura. Created & Engineered by Leonardo A Botelho.

**Estado em 24/11/2026:** pacote Intel (x86_64) aprovado em CI; Apple Silicon (ARM64) ainda sem aprovação, pois o job anterior foi cancelado. Nova tentativa solicitada. Não é equivalente à versão Windows 3.0.5 R5.

## Compilação e download de laboratório

- [Execução Macintosh 45396837719](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45396837719)
- [Pacote Intel 3.0.3 LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45396837719/artifacts/14921063483) (3.158.126 bytes; arquivo de Actions sujeito a login e retenção).
- SHA-256 do artefato Intel: `03eccf1255deb52ffa8a7dc18ae58934dc46b1dbc60d6b028b30acb78081c922`.
- Apple Silicon: aguardar aprovação do job ARM64 e artefato correspondente; não confundir com pacote Intel.
- Requisitos: macOS 13 ou posterior, arquitetura correspondente. CI usa macOS 15. DMG com assinatura ad hoc de laboratório, sem Developer ID ou notarização. Não desative Gatekeeper, XProtect nem outras proteções.

## Recursos e limites

Scanner local somente leitura, hash SHA-256, teste EICAR, alerta de extensão dupla e diagnósticos. Interface em português, inglês e espanhol. Módulos compartilhados de fraude, links, inteligência, reputação e diagnóstico têm testes separados, não equivalem à integração completa na interface. Sem quarentena, remediação, proteção residente, base IOC ativa, IA integrada ao scanner ou atualizador de produção comprovados.

CI Intel aprovou dez repetições de testes, race/fuzz, cargas sintéticas, compilação Swift/Go, verificação ad hoc e pacote DMG. Isso não comprova eficácia contra malware real ou funcionamento em todos os Macs.

## Desenvolvimento

Ramificação: `macintosh-3.0.3-lab`. Workflow: [macintosh-build.yml](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/macintosh-3.0.3-lab/.github/workflows/macintosh-build.yml). Consulte [MEMORANDO_MACINTOSH.md](MEMORANDO_MACINTOSH.md). O trabalho Windows 3.0.5 permanece separado. Não publicar releases de produção sem validação e autorização do titular.
