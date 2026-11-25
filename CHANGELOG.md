# Histórico de versões — ZEUVORYX Windows


## 3.0.10 LAB — candidato em validação (24/11/2026)
Solicitações após vídeos 3.0.9: separar autoteste da IA consultiva, saldo estimado do espaço ocupado menos bytes contabilizados, velocidade medida, alertas visuais/sonoros por novos achados, abertura junto ao Windows sem scan automático, e central de otimização. Código na ramificação windows-3.0.10-lab, commit 6605632e0a8103c4ef8bb4c0d16c2e04c76ab932. Ferramentas inclui limpeza temporária existente com prévia/confirmação, acesso a Otimizar Unidades, Limpeza de Disco e Gerenciador de Tarefas; Windows decide TRIM/defrag, usuário confirma; não é motor próprio completo. Testes de allowlist, taxas/alertas, startup e migração programados; CI ainda não concluído. Sem promessa de aceleração, aprendizagem contínua ou eficácia real. 3.0.9 permanece referência CI aprovada e teste físico parcial. Topics não impeditivos conforme titular; descrição pública permanece 3.0.9 até evidência de nova CI. CP-001–004 e 43 REQ preservados. Sem disponibilizar candidato como validado antes dos gates e sem alterar binário de teste 3.0.9.


## Estado verificado — ZEUVORYX 3.0.9 LAB com instalador (24/11/2026, Brasília)

CI [45519151028](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028), commit funcional 5115d20d53ab6fb7867f72d1b122e7a89a83391f, job 138886281969: 44/44 etapas concluídas com sucesso. Logs confirmam migração do instalador histórico 3.0.6 para 3.0.9, cinco ciclos de instalação/reinstalação, AppId e pasta mantidos, um desinstalador, hash do executável instalado igual ao compilado, preservação de marcador de dados e bloqueio de segunda instância (incluindo legado aberto primeiro). Teste de recuperação de quarentena e histórico reais ainda pendente; não há controle retroativo de portáteis legados abertos depois. 40 aprovações do teste IA em arquivos benignos; 45 aprovações dos três testes de relatório ao vivo; 15 aprovações do teste de instância única; 65 verificações de cores; 50M cenários sintéticos IA, QA20, 10M fraude, 50 raízes controladas, EICAR, regressões e abertura da interface com painel CPU/RAM. Não equivale a eficácia real ou certificação.

Artefato exato de laboratório [14979149692](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028/artifacts/14979149692), ZEUVORYX-3.0.9-LAB-SETUP-TESTE-Windows-x64 (ZIP de 7685290 bytes, contendo ZEUVORYX-LAB-Setup-3.0.9.exe conforme workflow). SHA-256 ZIP informado pelo GitHub: 4074aea09fda5f56687e9d8d65a6cff2b16bf9ac8f63bacb22d30d83a5ffb6f0; SHA-256 EXE instalador registrado no runner: 6b736fc1f422f6b3f72b7cd6853873c14b12d4762396589a77c362addaa29a1b. Não houve recálculo local dos bytes nesta auditoria. Artefato não expirado na consulta; acesso pode requerer login. Candidato somente para teste controlado; não promover a release, atualização assinada ou distribuição pública antes dos gates CP-001–004. Teste físico integrado em Windows 10/11, cinco ciclos reais e comparação de desempenho, eficácia real, estabilidade prolongada, dados reais, assinatura, licença/SBOM e paridade macOS continuam pendentes. Instalador LAB não empacota C2PA; não afirmar pacote integral de todas as 43 REQ. Inno Setup usado apenas no laboratório, sem contratação; alternativa gratuita comercial precisa de validação. Preservar Defender ativo, não desativar SmartScreen. REQ-022 adiado e REQ-043 suspenso. Não apagar histórico.

Metadados observados: descrição pública ainda 3.0.8 LAB under validation; topics vazios; homepage https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme. Proposta: ZEUVORYX Antivirus — Experimental Windows 3.0.9 LAB, CI validated, physical testing pending. macOS LAB separate. Complementary to Microsoft Defender. Topics propostos: windows, golang, security, malware-scanner, cybersecurity, antivirus-research, desktop. Escrita administrativa ainda indisponível; não afirmar correção efetiva.

## 3.0.8 LAB — intervenção A (24/11/2026)

- [Commit candidato](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/afdbfc6c4f2258171ece292e0c786d83930a25d5): elimina cálculo desnecessário de entropia, limita leitura individual a 256 KiB e separa cancelamento voluntário dos erros I/O no agregador desktop; testes e microbenchmarks adicionados, CI ainda não confirmado. Sem release ou instalador.


## 3.0.7 LAB — evidência física parcial; 3.0.8 LAB em preparação (24/11/2026)

- CI 45507766010 aprovado no commit e8bd034; teste rápido físico parcial: 19.701 arquivos, 13,05 GiB, 10min23s, 22,48 MB/s, três bloqueios de compartilhamento e um cancelamento voluntário, exportação JSON bem-sucedida. Nenhum alerta ou falha interna registrado; eficácia não comprovada.
- Ramificação 3.0.8 alinhada à base corrigida, sem nova compilação; otimização, cobertura e validação pendentes conforme memorando. Nenhum instalador ou release 3.0.8.


## 3.0.7 LAB R1 — em validação, não liberado (24/11/2026)

- Parser PE com limites explícitos e regressões PE32/PE32+ (20 casos, 30 repetições locais).
- Identidade candidata sincronizada; CI integral, desempenho e Windows físico pendentes. Sem novo instalador.

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


## ZVX-EVO-007 — fechamento pós-compilação, 24/11/2026
- Windows R5 aprovado: [45447823825](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825), commit `07373e9c538794c5f14a3bc13ad85206d3049bc9`. ZIP 14946014706, 33793191 bytes, SHA-256 `7232a6b12003b23a0e4de8a15e188d528533a63b15f11c93ea41771fa33111cb` confirmado no arquivo baixado e CRC íntegro. Instalador 15497251 bytes, SHA-256 `e90bedcbbcdd4cbb41020fac39b93d867beb0d4ebef884ff103e96632451ef39`; PE sem certificado Authenticode. Cinco rodadas links, dez repetições, QA20, 50M sintético, build e reinstalação passaram. Coleta DROP 45447823809 falhou por validade temporal; não relaxar validação. Sem release, sem alteração funcional.

- ZVX-EVO-008 (planejado): cinco estrelas fixas, seleção explícita, avaliação única persistente por instalação e deduplicação; nenhuma nota inicial contabilizada. Implementação e testes na próxima revisão funcional.

- ZVX-EVO-009 (planejado): central mundial de estatísticas/avaliações entre canais, métricas de proteção verificadas, arquitetura e gates em GLOBAL_SECURITY_INTELLIGENCE.md. Sem implantação ou alteração funcional.

- ZVX-EVO-010 (obrigatório na próxima revisão funcional): sincronização mundial ultraleve, consentimento e criptografia, dados mínimos, envio diário, testes ponta a ponta de privacidade, consumo e paridade entre canais. Especificação atualizada; serviço ainda não implantado.

- ZVX-EVO-011 (planejado): botão lateral Avaliação e resultados, painel funcional Olympus 3D estático, estrelas e gráficos baseados exclusivamente em dados reais, seleção Local/Mundial e testes de interação/desempenho. Conceito visual aprovado; implantação na próxima revisão funcional.

- ZVX-EVO-012 (próxima revisão): correção do caminho de atualização, catálogo LAB com link exato aprovado, sequência de revisão e integração local/mundial de avaliação e métricas via contrato API. Especificado, não implantado.

- ZVX-EVO-013: primeiro teste rápido em Windows real; candidato de retry nativo limitado 50/150 ms, testes de cancelamento e erros finais, privacidade de relatórios e métricas de manutenção planejadas. CI pendente; preservado pacote anterior.

- ZVX-CP-001: política pétrea de eficiência e continuidade cumulativa; preservar requisitos, medir e testar evoluções sem regressão, registrar evidências e limitações. Somente documentação.

- ZVX-EVO-015: incidente de encerramento simultâneo reportado em dois PCs Windows; investigação crítica, diagnóstico local e testes de estabilidade previstos. EVO-013 CI aprovado, causa do encerramento ainda desconhecida.

- ZVX-EVO-016: investigação de encerramentos na varredura completa; candidato de contenção de panic recuperável por arquivo e teste, sem causa confirmada; CI pendente.
