# ZEUVORYX Antivirus — Windows e Macintosh LAB


## 3.0.10 LAB — central integrada de otimização (24/11/2026)
Solicitação do titular: não apenas abrir ferramentas Windows; diagnóstico CPU/RAM, escolha e confirmação dos arquivos, desfragmentação/otimização de unidade, sem compressão automática. Commit candidato 8b36c26163a9950a12a1d3d9d4e66cefa15a5a7d em windows-3.0.10-lab. Diagnóstico por duas amostras e limiares, sem afirmar causa; prévia limitada a temporários .tmp antigos em LocalAppData/Temp, seleção desmarcada, exclusão via handle revalidado com confirmação, resultado removidos/ignorados/bytes lógicos; proteção de arquivos modificados/bloqueados e plano vencido. Otimização do volume Windows pelo defrag.exe nativo /O /U /V, UAC explícito, tomada, sem varredura concorrente, código real de saída, sem percentual inventado, sem ganho garantido. Não comprimir NTFS, apagar documentos, forçar RAM ou alterar registro. Acesso nativo a programas de inicialização preservado. Validação de CI e física pendentes, não afirmar aprovação; versão 3.0.9 permanece referência CI. REQ-011/EVO-039 parcial até medições físicas. Preservar 43 REQ e CP-001–004. Fonte https://learn.microsoft.com/windows-server/administration/windows-commands/defrag . Metadados administrativos mantêm descrição real da versão aprovada 3.0.9.


## 3.0.10 LAB — candidato em validação (24/11/2026)
Solicitações após vídeos 3.0.9: separar autoteste da IA consultiva, saldo estimado do espaço ocupado menos bytes contabilizados, velocidade medida, alertas visuais/sonoros por novos achados, abertura junto ao Windows sem scan automático, e central de otimização. Código na ramificação windows-3.0.10-lab, commit 6605632e0a8103c4ef8bb4c0d16c2e04c76ab932. Ferramentas inclui limpeza temporária existente com prévia/confirmação, acesso a Otimizar Unidades, Limpeza de Disco e Gerenciador de Tarefas; Windows decide TRIM/defrag, usuário confirma; não é motor próprio completo. Testes de allowlist, taxas/alertas, startup e migração programados; CI ainda não concluído. Sem promessa de aceleração, aprendizagem contínua ou eficácia real. 3.0.9 permanece referência CI aprovada e teste físico parcial. Topics não impeditivos conforme titular; descrição pública permanece 3.0.9 até evidência de nova CI. CP-001–004 e 43 REQ preservados. Sem disponibilizar candidato como validado antes dos gates e sem alterar binário de teste 3.0.9.


## Estado verificado — ZEUVORYX 3.0.9 LAB com instalador (24/11/2026, Brasília)

CI [45519151028](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028), commit funcional 5115d20d53ab6fb7867f72d1b122e7a89a83391f, job 138886281969: 44/44 etapas concluídas com sucesso. Logs confirmam migração do instalador histórico 3.0.6 para 3.0.9, cinco ciclos de instalação/reinstalação, AppId e pasta mantidos, um desinstalador, hash do executável instalado igual ao compilado, preservação de marcador de dados e bloqueio de segunda instância (incluindo legado aberto primeiro). Teste de recuperação de quarentena e histórico reais ainda pendente; não há controle retroativo de portáteis legados abertos depois. 40 aprovações do teste IA em arquivos benignos; 45 aprovações dos três testes de relatório ao vivo; 15 aprovações do teste de instância única; 65 verificações de cores; 50M cenários sintéticos IA, QA20, 10M fraude, 50 raízes controladas, EICAR, regressões e abertura da interface com painel CPU/RAM. Não equivale a eficácia real ou certificação.

Artefato exato de laboratório [14979149692](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028/artifacts/14979149692), ZEUVORYX-3.0.9-LAB-SETUP-TESTE-Windows-x64 (ZIP de 7685290 bytes, contendo ZEUVORYX-LAB-Setup-3.0.9.exe conforme workflow). SHA-256 ZIP informado pelo GitHub: 4074aea09fda5f56687e9d8d65a6cff2b16bf9ac8f63bacb22d30d83a5ffb6f0; SHA-256 EXE instalador registrado no runner: 6b736fc1f422f6b3f72b7cd6853873c14b12d4762396589a77c362addaa29a1b. Não houve recálculo local dos bytes nesta auditoria. Artefato não expirado na consulta; acesso pode requerer login. Candidato somente para teste controlado; não promover a release, atualização assinada ou distribuição pública antes dos gates CP-001–004. Teste físico integrado em Windows 10/11, cinco ciclos reais e comparação de desempenho, eficácia real, estabilidade prolongada, dados reais, assinatura, licença/SBOM e paridade macOS continuam pendentes. Instalador LAB não empacota C2PA; não afirmar pacote integral de todas as 43 REQ. Inno Setup usado apenas no laboratório, sem contratação; alternativa gratuita comercial precisa de validação. Preservar Defender ativo, não desativar SmartScreen. REQ-022 adiado e REQ-043 suspenso. Não apagar histórico.

Metadados observados: descrição pública ainda 3.0.8 LAB under validation; topics vazios; homepage https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme. Proposta: ZEUVORYX Antivirus — Experimental Windows 3.0.9 LAB, CI validated, physical testing pending. macOS LAB separate. Complementary to Microsoft Defender. Topics propostos: windows, golang, security, malware-scanner, cybersecurity, antivirus-research, desktop. Escrita administrativa ainda indisponível; não afirmar correção efetiva.

> **3.0.8 LAB — intervenção A (24/11/2026):** [código candidato](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/afdbfc6c4f2258171ece292e0c786d83930a25d5) gravado na ramificação separada. Otimização de entropia, leitura em blocos e cancelamento; validação CI ainda não confirmada, sem pacote disponível. 3.0.7 permanece referência de teste anterior. Não é proteção certificada nem release definitiva.


> **Estado vigente — 24/11/2026:** Windows **3.0.7 LAB**: CI aprovado ([execução 45507766010](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45507766010)), teste rápido físico parcial e exportação confirmados; eficácia e estabilidade integral pendentes. **3.0.8 LAB**: preparação em [ramificação isolada](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/windows-3.0.8-lab), base 3.0.7 corrigida, sem motor/compilação 3.0.8 aprovado. Macintosh permanece independente, candidato 3.0.5 LAB. Não substitui o Defender. [Memorando](MEMORANDO_DESENVOLVIMENTO.md) · [metadados](ABOUT_GITHUB.md). Nenhuma release definitiva.

## Histórico e documentação técnica (as versões abaixo são registros anteriores)


> **24/11/2026 — Windows 3.0.7 LAB R1, em desenvolvimento (não liberado):** correção experimental do parser PE baseada nos relatórios 3.0.6; 30 rodadas locais de testes do componente aprovadas. CI integral, estabilidade em Windows real, desempenho, cobertura e instalação/migração pendentes. [Ramificação candidata](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/windows-3.0.7-lab) · [Memorando ZVX-EVO-019](MEMORANDO_DESENVOLVIMENTO.md). Não há instalador 3.0.7 para download. Links anteriores são históricos de laboratório.


> **Central mundial de ajuda / Global Help Center / ヘルプセンター:** [Português · English · Español · 日本語](support/README.md). Documentação do candidato Windows 3.0.6 LAB R1; não é release de produção. [Evidência de CI](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45479946405).

Scanner complementar de segurança para Windows x64, desenvolvido em Go, com interface desktop, análise de arquivos e links, quarentena reversível e IA consultiva experimental. **Created & Engineered by Leonardo A Botelho.**

[![Windows LAB R5](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml/badge.svg?branch=main)](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml)

**Estado:** beta de laboratório. Compilação 3.0.5 R5 aprovada em 24/11/2026; validação manual em Windows 10/11 e eficácia independente pendentes. Não substitui Microsoft Defender.

**Revisão LAB pós-varredura e Global Threat Intelligence:** CI Windows aprovado no commit `07373e9c538794c5f14a3bc13ad85206d3049bc9`; alterações de interface, classificação de erros, retry limitado, procedência do build e consulta consultiva DROP incluídas. Validação visual e em Windows real ainda necessária.

## Download e instalação de teste

- [Compilação validada — execução 45447823825](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825)
- [ZIP completo — artefato 14946014706](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825/artifacts/14946014706) (33.793.191 bytes; GitHub Actions pode exigir login; expira em 22/02/2027).
- SHA-256 do ZIP, conferido nos bytes baixados: `7232a6b12003b23a0e4de8a15e188d528533a63b15f11c93ea41771fa33111cb`.
- Instalador extraído sem executar: 15.497.251 bytes; SHA-256 `e90bedcbbcdd4cbb41020fac39b93d867beb0d4ebef884ff103e96632451ef39`; sem certificado Authenticode incorporado. ZIP CRC conferido.
- Extraia o ZIP, finalize a varredura anterior e feche o aplicativo. Abra `installer/ZEUVORYX-LAB-Setup-3.0.5.exe`. A edição portátil está em `dist/`.
- Teste primeiro em Windows Sandbox com Defender ativo e WebView2 disponível. Não desative SmartScreen nem proteções existentes. O binário não tem assinatura Authenticode de produção comprovada.
- Para uma instalação LAB compatível, mantenha o diretório anterior (AppId ZEUVORYX.Desktop.Lab303). CI aprovou reinstalação com marcador de dados; faça backup e valide histórico/quarentena. Migração de edições 3.0.1/3.0.2 não validada.
- Nenhuma release pública está sendo promovida por este documento; o botão de atualização assinada requer chave pública oficial fixada, manifesto/pacote assinados e release autorizada.

## Recursos disponíveis para avaliação

| Área | Estado |
| --- | --- |
| Varreduras | Rápida, completa e personalizada; cancelamento, resultados, métricas, exportação e cobertura parcial |
| Motor | SHA-256/IOC, heurísticas, inspeção limitada PE/ZIP, EICAR e análise consultiva de links |
| Quarentena | Ação explícita com critérios de evidência; reversibilidade e restauração sujeitas a validação |
| Segurança Windows | Integrações com Defender e recursos nativos; não é proteção residente própria |
| IA local | Modelo experimental auxiliar, sem aprendizado contínuo em produção; candidato reprovado não promovido |
| Interface | Olympus Silver, Zeus e templo; português, inglês e espanhol, seleção de idioma |
| Histórico | Registro local autenticado, relatório de erros, avaliação voluntária; sem estatísticas globais |
| Atualizador | Implementação experimental de descoberta e validação assinada; produção e rollback ainda pendentes |

## Evidência de CI

Execução [45436304817](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304817), commit `214a6de2586433176987efdad8147ce3a1b02c41`: gate de versão 3.0.5, cinco rodadas de regressão de links, dez repetições de testes, QA20 (20 milhões de iterações sintéticas), build EXE, instalador e teste de reinstalação aprovados. Laboratório de IA com 50 milhões de cenários sintéticos não comprova proteção real; testes de laboratório não equivalem a certificação.

## Desenvolvimento e continuidade

A ramificação `main` consolida a base de código Windows 3.0.5 LAB R5; `security-3.0.5-lab` preserva a referência de laboratório. O workflow `.github/workflows/windows-build-lab-305.yml` chama `windows-build.yml` na ramificação selecionada. A versão antiga está preservada no histórico. A compilação integral em main foi aprovada; a regressão de links independente [45436304882](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304882) também foi aprovada. Na execução de fundamentos [45436304993](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304993), Windows e Linux passaram; o job macos-updater foi cancelado e a execução global ficou com falha. Consulte [instruções](INSTRUCOES_GITHUB.txt), [instalação](LEIA-ME.txt), [histórico de versões](CHANGELOG.md) e o [memorando obrigatório](MEMORANDO_DESENVOLVIMENTO.md).

Próximos passos: painel com comandos de varredura e métricas no topo, botões compactos 3D; teste visual e Windows real; atualização assinada de produção; auditoria ampliada, remediação segura, proteção residente, IA validada, estatísticas globais ultraleves e Macintosh independente. Idiomas adicionais conforme qualidade/desempenho. Nenhum desses requisitos é considerado concluído só por estar planejado.

## Macintosh — ZEUVORYX 3.0.5 LAB (candidato atual)

A edição Macintosh **3.0.5 LAB** é o candidato atual, desenvolvido separadamente com Swift/SwiftUI e motor Go somente leitura. [Documentação atual](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/macintosh-3.0.5-lab). A regressão portátil Linux e a compilação Intel foram aprovadas. A tentativa ARM64 foi cancelada sem etapas registradas; reteste exclusivo solicitado, sem aprovação confirmada. Não há equivalência funcional integral com Windows nem autorização de distribuição de produção.

Para novos testes, selecione o workflow Macintosh 3.0.5 e a ramificação `macintosh-3.0.5-lab`. A [versão histórica 3.0.3 LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/macintosh-3.0.3-lab) permanece preservada. Não confundir instaladores, arquiteturas ou candidato com pacote aprovado.

## Licenciamento, autoria e distribuição

Autoria declarada: Leonardo A Botelho. Não há licença de redistribuição de produção definida neste documento; revisar licenças de dependências e direitos de arte antes de publicação. O projeto não reivindica certificação, detecção universal ou substituição de antivírus. Releases públicas e lançamento mundial dependem de autorização do titular e dos critérios do memorando.


### Conferência de CI — 24/11/2026

Macintosh 3.0.5 LAB: regressão portátil Linux e compilação Intel aprovadas; Apple Silicon ARM64 cancelado na tentativa anterior, sem etapas registradas e com log indisponível (404). Reteste exclusivo solicitado e aceito pelo GitHub em 24/11/2026; resultado ainda não confirmado. Não há causa de código comprovada. [Execução 45437873318](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318). [Artefato Intel 14943370193](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318/artifacts/14943370193), 4.283.792 bytes, SHA-256 registrado pelo GitHub `965da629bf97f885a7df1de2434211f3342d82ba4e286f1a1fb52f946bc3f767`. Somente para laboratório; sem validação integral em Macs reais ou assinatura de produção.


## Global Threat Intelligence (CI Windows aprovado; laboratório)

Nova consulta consultiva Spamhaus DROP IPv4/IPv6 e inventário de fontes em [THREAT_INTELLIGENCE.md](THREAT_INTELLIGENCE.md). Não é bloqueio de firewall nem limpeza automática. Base de sessão obtida sob demanda; nenhuma contagem real antecipada. CI Windows aprovado na execução 45447823825; coleta inicial separada 45447823809 falhou com `stale or future feed`, sem artefato de dados. Nenhuma base DROP real embutida. Versões anteriores preservadas.


### Evidência mais recente — ZVX-EVO-007

https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825 — gate de identidade, cinco rodadas de links, dez repetições, QA20, 50M sintético, build e reinstalação aprovados. Não equivale a eficácia real, migração entre versões distintas, validação do feed ao vivo ou distribuição pública. Os registros de CI antigos acima são históricos.
