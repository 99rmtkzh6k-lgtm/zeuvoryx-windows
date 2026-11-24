# Histórico de versões — ZEUVORYX Windows

O histórico detalhado, critérios e pendências estão no [memorando permanente](MEMORANDO_DESENVOLVIMENTO.md). Datas conforme GitHub; início de concepção anterior ao primeiro commit não confirmado.

| Versão | Evolução |
| --- | --- |
| Base/2.1 — 23/11/2026 | Fontes Go, scanner, PE, quarentena, interface, arte, workflow e correções de compilação. |
| 3.0.1 Competitive | Otimização, filtros, métricas, heurísticas e integração Windows; base histórica preservada no Git. |
| 3.0.2 | IA local experimental, três idiomas, assinatura visual de autoria, laboratórios sintéticos e testes. |
| 3.0.3 LAB | Fundamentos de atualização/estatísticas assinadas, política de remediação, diagnósticos locais e arte oficial. |
| 3.0.4 LAB | Descoberta de pré-releases e continuidade da instalação LAB. |
| 3.0.5 LAB | Relatório identificado, contagem de erros, histórico autenticado e avaliação local. |
| R2 | Olympus Silver com Zeus/templo, correção de IDs, download manual. |
| R3 | Atualizador assinado candidato, instalador e teste de reinstalação. |
| R4 | Confirmação da atualização ao reabrir a versão-alvo. |
| R5 | QA20, regressões de links, verificações Windows e identidade própria da compilação. |

## R5 — compilação validada, 24/11/2026
- [Execução 45419964764](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45419964764) — sucesso; commit 11cd939ff0944bf2da9a81f6f4bc166559e5384b.
- Gate de identidade 3.0.5, regressões, testes repetidos, QA20, build e instalação de teste aprovados.
- Artefato completo 14934540577: 33.730.421 bytes; hash registrado c42366ea44b6ac25c302b92a1fdbacb237c01224eb8b4e13b3a2e34ada19fcf0.
- Próxima revisão: painel operacional no topo, gráficos e números em destaque, botões 3D compactos. Planejada; não entregue pelo build acima.
- Não confundir aprovação do CI com eficácia independente, atualização de produção ou distribuição pública.

## Consolidação e continuidade — 24/11/2026
- Windows 3.0.5 LAB R5: [execução 45436304817](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304817) aprovada na main, commit 214a6de2586433176987efdad8147ce3a1b02c41. [Artefato 14942002196](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304817/artifacts/14942002196), 33.725.391 bytes, SHA-256 informado pelo GitHub ae703ec4d3460fe0126102f7ed334df13015571b50e36f9e28e6e72f8eb993ff. [Regressão de links 45436304882](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304882) aprovada.
- Macintosh 3.0.3 Intel: compilação histórica aprovada; reteste ARM64 pendente. Macintosh 3.0.5 LAB candidato: análise consultiva de links portada, CLI/UI e gate de versão; [execução 45437873318](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318) na fila nesta consulta. Não equivale a paridade ou aprovação final.
- Regra permanente ampliada: padronização, auditoria de impacto, memorando vivo, pesquisa técnica e melhoria contínua, sincronização e verificação de metadados. About/Topics/Website ainda requerem intervenção manual conforme [ABOUT_GITHUB.md](ABOUT_GITHUB.md).
- Nenhuma nova compilação ou alteração funcional por esta revisão documental.

## Protocolo de engenharia e conferência de CI — 24/11/2026
- Expressão de referência do titular registrada e nove controles ZVX-STD definidos nos memorandos, com referências públicas. Implementação técnica adicional permanece no backlog; nenhuma certificação reivindicada.
- Windows 3.0.5 R5 aprovado; fundamentos com Windows/Linux aprovados e macOS updater cancelado. Macintosh 3.0.3 ARM64 novamente cancelado. Macintosh 3.0.5: regressão portátil e compilação Intel aprovadas, ARM64 pendente. About ainda requer edição manual.


### Conferência de CI — 24/11/2026

Macintosh 3.0.5 LAB: regressão portátil Linux e compilação Intel aprovadas; Apple Silicon ARM64 cancelado na tentativa anterior, sem etapas registradas e com log indisponível (404). Reteste exclusivo solicitado e aceito pelo GitHub em 24/11/2026; resultado ainda não confirmado. Não há causa de código comprovada. [Execução 45437873318](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318). [Artefato Intel 14943370193](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318/artifacts/14943370193), 4.283.792 bytes, SHA-256 registrado pelo GitHub `965da629bf97f885a7df1de2434211f3342d82ba4e286f1a1fb52f946bc3f767`. Somente para laboratório; sem validação integral em Macs reais ou assinatura de produção.

## ZVX-EVO-001 — primeira revisão do protocolo, 24/11/2026
- Auditoria Windows/macOS e metadados, pesquisa actionlint, govulncheck e atestação de artefatos; propostas registradas, sem implantação funcional. AGENTS.md adiciona orientações e links canônicos. Windows e Macintosh Intel 3.0.5 LAB aprovados em CI; ARM64 pendente; macos-updater cancelado. About requer edição manual. Sem novo build/release.

- ZVX-EVO-002: documentação sincronizada para distinguir candidato Macintosh 3.0.5, histórico 3.0.3 e cancelamento ARM64; reteste solicitado. Sem alteração de código ou versão.


### 24/11/2026 — ajustes candidatos pós-varredura completa (ZVX-EVO-003)
- Identificação do commit de compilação nos relatórios JSON; navegação manual para pacotes Windows LAB em Actions mesmo sem Releases; mensagens nos três idiomas e teste de regressão para releases vazios/mesma versão. Sem publicação, sem alteração da política de assinatura; aprovação do novo build e validação manual pendentes. Ver memorando central.


### 24/11/2026 — ZVX-EVO-004 candidato pós-varredura
- Olympus Silver: controles de varredura, métricas e gráfico reposicionados no painel inicial; botões compactos com efeito 3D em CSS, estados de foco e pressionado; sem novos recursos externos. Exibe erros, ignorados, duração e hashes distintos/repetidos ao concluir. Não altera ou elimina alertas individuais.
- Scanner: categorização de erros nativos no percurso de discos; uma nova tentativa para sharing/lock violation Windows, cancelável e sem elevar privilégios. Testes positivos/negativos e regressões adicionados. Preservados AppId, assinatura exigida, versão-base 3.0.5 e R5 anterior. Validação em andamento; nenhuma release promovida.

- ZVX-EVO-005: corrigida âncora do gerador de retry após gofmt, causa comprovada no log de 45445139048. Reteste obrigatório.

- ZVX-EVO-006: candidato consulta DROP oficial IPv4/IPv6, validação de formato/idade, limites, atribuição e busca local de IP; testes novos e inventário de fontes. Sem bloqueio automático, amostras de malware ou release.

- Ajuste de frequência DROP para 24h; seed opcional com validação e coleta manual de dados oficiais separada, sem alegar base preenchida.
