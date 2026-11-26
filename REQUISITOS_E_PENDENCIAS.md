## 3.0.18 LAB — Central de Segurança do Windows

A determinação complementar foi vinculada aos itens de Windows, proteção residente, relatório e vulnerabilidades. A Central oferece operações suportadas e consultas dirigidas. Validação física de políticas, prevenção autônoma, varredura CVE e testes de navegador seguem abertos. [Registro técnico](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/CENTRAL_SEGURANCA_318.md).

## 3.0.17 LAB — integração complementar verificada (26/11/2026)

A [auditoria complementar](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/AUDITORIA_MOTORES_2026-11-26.md) registra a Issue #5, a revisão dos 20 itens ClamAV e a vinculação à matriz mestre. REQ-003/004/007/025/027/031/041 continuam abertos para validação de campo, proteção residente e cobertura. REQ-018/026/028 com artefato LAB verificável.

## Integração 3.0.16 LAB — 26/11/2026

A ordem de execução de 26/11 autorizou a evolução experimental. [Matriz mestre, critérios e resultados](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/MATRIZ_MESTRE_2026-11-26.md). REQ-003/004/025/027/041 abertos para o gate físico. REQ-026 documentado com novo número, CI e hashes. REQ-018/028 com instalador validado em CI e ZIP temporário para piloto. Requisitos anteriores permanecem rastreados. A determinação vigente limita o Arcade a até 100 jogos gratuitos e define a assinatura artística L A Botelho, sem bandeira.

## Diagnóstico de Testes no instalador Windows LAB — 26/11/2026

A [CI 45755065493](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45755065493) aprovou os três jobs (regressão Ubuntu, regressão Windows e compilação/instalação Windows) para `b90fa60343be894b436ec7cd1f0971ac8101bd9c`. Esta revisão inclui a finalização cumulativa da varredura, o limite da etapa ClamAV e a área Relatórios → Diagnóstico de Testes, com causas possíveis, alternativas, cópia para IA e exportação JSON. A execução anterior 45754511971, de `b55faec`, foi cancelada.

[Baixar o instalador LAB Windows x64](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45755065493/artifacts/15087906503). SHA-256 do ZIP informado pelo GitHub: `1dadd1c43709f8a5ecfa683c5b18d252b6052e0b9be2a472a385e90015a0f017`. Candidato de teste 3.0.15 LAB identificado pelo commit. A cópia e a exportação no WebView2 instalado e a conclusão da varredura completa do C: ainda requerem validação em campo.

## Diagnóstico de Testes copiável — 26/11/2026

A revisão LAB [b90fa60](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/b90fa60) acrescenta ao centro de Relatórios a área Diagnóstico de Testes, com geração e atualização explícita, causas possíveis, verificações, alternativas, impactos, critérios de aceite, cópia para IA e exportação JSON local. Usa dados agregados do relatório selecionado. Caminhos e logs brutos ficam fora da saída. Sem integração remota. Os testes locais de geração, estados e privacidade passaram. Validação do instalador, interface WebView2 e uso físico pendentes.

Esta é uma primeira implementação do diagnóstico, com regras para cancelamento, ClamAV em andamento, limites, bloqueios e inventário parcial. Histórico de tentativas, diagnóstico aprofundado do autoteste e integração de todas as hipóteses ainda pendem. Permanecem abertos watchdog, recuperação, pausa e retomada, ETA por motor, causa da lentidão e validação da varredura completa. As demais frentes continuam no índice cumulativo.

## P0 — correções cumulativas de varredura Windows — 26/11/2026

O relatório físico do instalador `0295579` confirma a rápida em 123,189 s, com relatório final automático, 259/266 arquivos no scanner próprio e 197 no ClamAV (zero infecções reportadas, 14 erros, nove limites, exit 2). Os campos de volume do ClamAV foram preenchidos: 1.699,84 MiB lidos e 143,10 MiB examinados. Esse pacote não continha as correções da ramificação P0 de finalização.

A CI 45750382240 aprovou a revisão P0 `a85df421`. O commit cumulativo [b55faec](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/b55faec) usa essa revisão e incorpora o limite do ClamAV de 60 min na completa, 5 min na rápida e 20 min na específica. Inclui espera limitada de inventário e diagnósticos, persistência final fora do mutex de estado, identificação do horário/estado do snapshot e respeito à configuração de suspensão do Windows. Cinco ciclos locais dos módulos scopeprogress, scanplan, signatureav, scanretry e auditmetrics passaram. CI do commit cumulativo e validação física da completa pendentes. Watchdog de leitura bloqueada, política de pausa/retomada por eventos de energia, progresso e ETA por motor, recuperação de sessões e causa da lentidão ClamAV continuam em aberto.

## Continuidade do Windows 3.0.15 LAB — 25/11/2026

A [CI 45696640691](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45696640691) do commit de código `9d8b2c82eefa8d1c99275cb708429a30ae566a63` aprovou os três jobs (regressão Ubuntu, regressão Windows e compilação/validação Windows). Este candidato acrescenta métricas nativas do ClamAV ao relatório, incluindo bytes lidos e examinados, duração, limites e truncamento. A [CI anterior de cobertura 45691904568](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45691904568), do commit `a48e7eb`, também terminou com sucesso.

[Instalador Windows x64 para teste LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45696640691/artifacts/15061539008). SHA-256 do ZIP informado pelo GitHub: `f3f0a133db27914bfbcaa55b762557118e11974557c0b3072622a34202394130`. O teste físico deve confirmar se a varredura completa termina, grava o relatório final e distingue arquivos ocupados, limites de leitura e interrupções, com medições de CPU, RAM e duração. A validação física, a investigação de instabilidade e a aprovação para produção seguem pendentes.

A [CI 45617532781 do Windows 3.0.14 LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45617532781), commit `a4603eb`, passou nos três jobs, incluindo cinco testes do aplicativo instalado com ClamAV embarcado. A meta do [Olympus Arcade](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/issues/4) continua em auditoria. O [PR legado do atualizador 3.0.2](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/pull/3) segue aberto e com conflito de merge. Esta revisão documenta o candidato Windows e não acrescenta evidência para Intel ou ARM64 no Macintosh.

## Determinação vigente do idealizador — 25/11/2026: Artigo 37 substituído

Não há limite interno obrigatório de tamanho para o ZEUVORYX: revogados o teto de 200 MB e a meta numérica de 150 MB. O instalador, ZIP, motores, bases e dependências poderão crescer quando tecnicamente necessários. Priorizar proteção, eficácia, estabilidade e desempenho; continuar medindo e otimizando tamanho, RAM, CPU, disco e downloads sem remover cobertura ou funcionalidades aprovadas apenas para reduzir bytes. Registrar crescimento e justificativa, respeitar limites externos dos canais de distribuição e manter controles de integridade, licenças, gratuidade e validação. Registros anteriores de limites são históricos. Entregar sempre que possível o ZIP completo no ChatGPT para compartilhar pelo WhatsApp, não somente o link GitHub.

# ZEUVORYX Antivirus 3.0.15 LAB — Olympus | Referência vigente 25/11/2026

**CI aprovada; candidato Windows x64 para teste físico controlado. Não é versão comercial.**

- Código efetivamente compilado: [c8bf1fd](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/c8bf1fd7614e44de297569663fc90fa5ab066a47), ramificação [Olympus P0](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/lab-olympus-p0-defender-20261125).
- [CI integral](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45675394546): 48/48 etapas Windows concluídas com sucesso, cinco ciclos de regressão de módulos Windows e Linux, dez ciclos gerais e cinco testes do aplicativo instalado no CI. Migração 3.0.6 → 3.0.15 testada no CI.
- [ZIP do instalador LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45675394546/artifacts/15052358598): 192807638 bytes (192,81 MB decimais), abaixo do teto histórico de 200 MB (revogado); SHA-256 do ZIP informado pelo GitHub: bd9bcd3461e7189944a9a295b5a5e4e1fff447652c798d87e00aa24b4ca56744. Não é o hash individual do EXE; não houve verificação independente dos bytes nesta consulta. O portátil (295087036 bytes) é um pacote separado de teste; a entrega recomendada continua sendo o instalador completo.
- Revisão incremental: três commits desde 935b906, três arquivos alterados; distinção entre comando Defender aceito e eventos de varredura observados (1000/1001/1002), teste de contrato, workflow isolado e auditoria. Vínculo exclusivo dos eventos, execução PowerShell real e exportação integrada da evidência ainda não comprovados.
- Preservar interface, autoria Leonardo A. Botelho, identidade Olympus e requisitos REQ-001–043. Nenhuma promoção de código para main; main permanece referência histórica com documentação atualizada. Manter proteções Windows ativas.
- P0 pendentes: teste físico rápido/completo, RAM 2/4/8/16 GiB, progresso/ETA, copiar JSON e recuperação, cobertura dos motores, Defender por sessão e eficácia independente. CI não comprova proteção integral. Sem publicação comercial, atualização automática de produção ou cobrança.
- About público consultado: já informa 3.0.15 LAB e validação física em andamento; homepage correta; Topics vazio. O conector não expõe escrita administrativa desses campos. Sugestões de tópicos: antivirus, windows, golang, cybersecurity, endpoint-security, clamav, threat-detection.

## Registros históricos — preservar para auditoria

# Referência atual

ZEUVORYX Windows — 3.0.15 LAB, revisão corretiva — 25/11/2026
Código candidato: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/ad42721210a22afda99796d54a91eeb540fe6aed
Testes dos módulos aprovados: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45661075714
CI integral em andamento na consulta: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45661075658 (40/47 etapas aprovadas). Sem novo instalador liberado nesta consulta. Teste físico ainda pendente. A revisão anterior apresentou interrupção de varredura no PC. Manter Defender ativo e não promover release pública. As fontes atuais estão na ramificação lab-fieldlearn-20261125; main é base histórica com documentação atual. Não misturar árvores nem alterar nomes de pastas. Datas antigas significam último commit, não necessariamente código desatualizado. About confirmado 3.0.15 LAB, Topics vazios na última consulta.

--- Histórico anterior: não utilizar como instrução de instalação atual ---

> **25/11/2026 — ZEUVORYX 3.0.15 LAB (candidato em desenvolvimento):** interface e recursos atuais preservados; ClamAV em lote único, correções de cobertura/ETA, alertas e JSON com diagnóstico Windows/Sysmon. Identidade 3.0.15 preparada para compilação. CI, teste físico e eficácia independente pendentes; não é versão de produção nem contém SDKs comerciais. [Código candidato](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/windows-3.0.14-audit-fix) · [Memorando](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/windows-3.0.14-audit-fix/AUDITORIA_CAMPO_2026-11-25.md).

> **25/11/2026 — 3.0.14 LAB, revisão de auditoria física (candidata):** correções de ETA global, inventário parcial, publicação incremental ClamAV, JSON e alertas por categoria. Novo teste de regressão; compilação e validação física pendentes. Proteção residente, navegação, memória, antirransomware e eficácia independente ainda não comprovados. Avaliar SDKs OEM oficiais sem contratar ou distribuir sem licença. [Auditoria e mini memorando](AUDITORIA_CAMPO_2026-11-25.md).

> **3.0.14 LAB · Correção do teste instalado:** CI 45614993623 aprovou motores independentes, compilação e cinco ciclos de instalação/migração; falhou ao localizar amostra benigna. Preparação literal e individual, verificação de existência/tamanho e diagnóstico adicional antes/depois FreshClam, preservando critérios e proteção Windows. Nova CI pendente, sem liberação. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB — pacote unificado:** Controle parental independente no menu lateral, filtro DNS configurável no Windows, revisão de jogos real e painel de progresso durante jogos; correção de integração DROP/Go formatado; cinco ciclos de regressão e validação de instalação exigidos. CI e teste físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB · Central familiar do computador:** além do catálogo, integração de configuração voluntária de DNS familiar Windows, teste de bloqueio real IPv4, Microsoft Family Safety por conta e orientações de limitações. Sem alteração automática de rede nem garantia universal. CI e teste físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB · Proteção familiar:** revisão real de amostra HTML dos jogos e filtro opcional de catálogo por idade e categorias, senha local, avaliação do responsável e integração de acesso ao Microsoft Family Safety. Bloqueio de arquivos pornográficos e filtro global ainda pendentes. CI e teste físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB · Revisão real dos jogos:** busca HTTPS restrita, amostra HTML analisada por ClamAV e scanner local, consulta DROP disponível, resultado/hora/hash, abertura condicionada à revisão de cinco minutos. Sem garantia de conteúdo externo futuro. Validação CI e física pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB · Olympus Arcade:** Retro Game Arena com arte Zeus aprovada, colunas gregas, raios e cores fluorescentes, botões de relevo e identificação por plataforma. Sem nova dependência; funções e proteção preservadas. CI e teste visual pendentes.

> **3.0.14 LAB · Retro Game Arena:** módulo opcional de jogos gratuitos em páginas autorizadas/de autores, filtro por plataforma, busca, jogo local e arte vetorial; sem ROMs comerciais nem dependências novas. CI e teste físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **Correção de validação da arte 3.0.14 LAB:** CI anterior bloqueado por diferença de codificação JPEG. Comparação local confirmou a mesma arte de Zeus; agora são aceitos apenas os dois hashes verificados (original incorporada e cópia aprovada). CI pendente. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB Olympus Titanium — arte em tela inteira:** imagem original Zeus/Atenas preservada sem cortes, fundo fixo com preenchimento secundário, cartões translúcidos e texto sobreposto, opção de contemplação sem controles. CI e teste visual físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **Arte oficial Olympus Titanium 3.0.14 LAB:** imagem de Zeus/Atenas confirmada pelo proprietário; fundo e banner com nome temático; verificação SHA-256 da imagem original obrigatória. CI e inspeção física pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB — Olympus Titanium · Global Threat Defense:** arte Zeus existente como fundo global, painéis translúcidos e botões de relevo profundo, velocímetros e integração funcional preservados. Corrigido teste EICAR do instalador; nova validação pendente. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB revisão painel automotivo:** indicadores com ponteiros, inventário lógico do escopo, saldo estimado e ETA do scanner próprio, sem confundir armazenamento físico com cobertura. CI pendente. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB revisão ampliada:** pacote único ClamAV/FreshClam/C2PA, motores em conjunto, IA consultiva com revisão por indícios correlacionados, alerta de saída ClamAV, DROP persistente diário e melhoria de manutenção. CI e validação física pendentes; aprendizado real e proteção residente ainda não comprovados. [Detalhes](BUNDLED_314_LAB.md).

> **3.0.14 LAB revisão integrada:** motores próprios e ClamAV em paralelo, consulta Sysmon opcional, diagnóstico real separado do autoteste sintético. CI e teste físico pendentes. [Detalhes](BUNDLED_314_LAB.md).

> **25/11/2026 — candidato 3.0.14 LAB:** correção da dependência externa do ClamAV evidenciada em teste físico 3.0.13. Motor oficial incluído no instalador, atualização automática da base e testes reais via aplicativo instalado; CI pendente. [Estado e limites](BUNDLED_314_LAB.md). Não é release de produção.

> **25/11/2026 — 3.0.13 LAB candidato consolidado:** assinatura, varredura normal, atualização FreshClam sob demanda, quarentena revalidada e telemetria Sysmon opcional. CI e validação física pendentes. [Estado e limites](CONSOLIDATED_313_LAB.md). Não é release de produção.

> **25/11/2026 — candidato 3.0.12 LAB:** segunda etapa independente na varredura normal e relatório consolidado; CI pendente. [Evidência e limites](SIGNATURE_312_LAB.md). 3.0.11: ClamAV real aprovado em cinco testes EICAR/benigno no Linux; não comprova proteção real. Sem release de produção.

> **25/11/2026 — 3.0.11 LAB, candidato não aprovado:** integração de análise independente por ClamAV externo, relatório JSON separado, CI e teste físico pendentes. Não é proteção residente. [Evidência e limites](SIGNATURE_311_LAB.md). Preservar Defender e main.

# ZEUVORYX — registro cumulativo de requisitos e pendências


### 3.0.10 — desfragmentação e TRIM integrados
Titular esclarece que compactação significava desfragmentação. Implementação candidata: botão de diagnóstico e limpeza seletiva no ZEUVORYX; botão de otimização da unidade Windows via defrag.exe do diretório de sistema, parâmetros fixos /O /U /V (Windows escolhe mídia), UAC explícito só para ferramenta nativa, energia AC obrigatória, bloqueio durante varredura e exclusão mútua, conclusão pelo código de saída real. Sem compressão NTFS automática, sem forçar desfragmentação de SSD, sem alterar registro ou prometer aceleração. Nesta etapa não captura progresso percentual ou relatório nativo detalhado, nem testa operação real em disco físico; testes automatizados verificam argumentos e controles, execução física deve confirmar UAC e resultado. Referência https://learn.microsoft.com/windows-server/administration/windows-commands/defrag . Preservar gates e 3.0.9. CI pendente.


## 3.0.10 LAB — otimização integrada, não apenas atalhos (24/11/2026)
Solicitação: diagnóstico e otimização dentro do ZEUVORYX, seleção individual e confirmação de arquivos. Novo fluxo na visão geral e Ferramentas: duas amostras reais CPU/RAM, orientação por limiares explícitos (80% CPU, 85% RAM nas duas amostras; não diagnóstico causal), prévia até 200 .tmp antigos (>=7 dias) na pasta LocalAppData/Temp, enumeração até 5000 entradas, seleção inicialmente desmarcada, tamanho lógico/idade, confirmação de exclusão definitiva. Rejeitar caminhos de rede, reparse points, arquivos protegidos/ocultos/offline, hardlinks, bloqueados, seleção inválida, plano vencido (5 min), mudança de identidade/tamanho/horário; apagar por handle Windows verificado, não por segunda busca de caminho. Sem apagar documentos, sistema, registro, quarentena, desligar processos ou forçar RAM. Exibir removidos/ignorados/bytes lógicos, sem prometer aceleração. Acesso nativo SSD/HD e gerenciamento inicialização preservados; otimização física de unidade ainda depende de ferramenta Windows/consentimento, não motor autônomo. Teste de seleção, preservação de arquivos, plano vencido e alteração executado pelo CI em cinco ciclos e regressões, ainda não confirmado. Requisito REQ-011/EVO-039 parcial, ganho real e testes físicos pendentes. Fontes API: https://learn.microsoft.com/windows/win32/api/winbase/ns-winbase-file_disposition_info ; https://learn.microsoft.com/windows/win32/api/fileapi/nf-fileapi-setfileinformationbyhandle . Não promover nem distribuir antes de validação, preservar 3.0.9 e 43 requisitos.


## 3.0.10 LAB — feedback físico e central de otimização (24/11/2026)
Titular solicitou correção do robô acionado por autoteste, saldo decrescente do armazenamento, taxas arquivos/s e bytes/s, alertas de relatório em tempo real com som, abertura junto ao Windows e central de otimização. Candidato isolado windows-3.0.10-lab; preservar 3.0.9 (CI aprovado, teste físico parcial). Separação autoteste/observações consultivas; aprendizado contínuo e eficácia real não implantados, sem dados rotulados independentes. Saldo estimado máximo(0, ocupado - bytes contabilizados) apenas full, não universo elegível nem I/O físico. Taxas por deltas reais, reset por scan, som Windows ao novo achado de relatório, revisão e quarentena com confirmação após parada/conclusão. Startup selecionável no instalador, habilitado por padrão conforme solicitação, sem varredura automática. Central Ferramentas: limpeza .tmp existente com prévia/consentimento; acesso explícito a Otimizar Unidades (dfrgui), Limpeza de Disco (cleanmgr) e Gerenciador de Tarefas (Taskmgr) para inicialização. Windows escolhe TRIM/defrag conforme unidade; sem execução automática, registro, encerramento de processos, aumento de velocidade garantido, ou limpeza indiscriminada; tarefas de disco bloqueadas durante varredura. Não é motor de otimização próprio completo. Testes 5 ciclos de métricas, allowlist de ferramentas, migração real com fixtures 3.0.6/3.0.9, startup opt-out/opt-in, regressões anteriores; CI novo em validação. REQ-011/030/025/034 e EVO-039 parcialmente implementados, validação física e comparação antes/depois pendentes. Preservar CP-001–004, 43 REQ, assinatura/licenças, Defender e dados. Sem publicação ou download antes de CI e verificação do artefato. Metadados administrativos preservam 3.0.9 CI aprovado, Topics não impeditivos.


## Estado verificado — ZEUVORYX 3.0.9 LAB com instalador (24/11/2026, Brasília)

CI [45519151028](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028), commit funcional 5115d20d53ab6fb7867f72d1b122e7a89a83391f, job 138886281969: 44/44 etapas concluídas com sucesso. Logs confirmam migração do instalador histórico 3.0.6 para 3.0.9, cinco ciclos de instalação/reinstalação, AppId e pasta mantidos, um desinstalador, hash do executável instalado igual ao compilado, preservação de marcador de dados e bloqueio de segunda instância (incluindo legado aberto primeiro). Teste de recuperação de quarentena e histórico reais ainda pendente; não há controle retroativo de portáteis legados abertos depois. 40 aprovações do teste IA em arquivos benignos; 45 aprovações dos três testes de relatório ao vivo; 15 aprovações do teste de instância única; 65 verificações de cores; 50M cenários sintéticos IA, QA20, 10M fraude, 50 raízes controladas, EICAR, regressões e abertura da interface com painel CPU/RAM. Não equivale a eficácia real ou certificação.

Artefato exato de laboratório [14979149692](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028/artifacts/14979149692), ZEUVORYX-3.0.9-LAB-SETUP-TESTE-Windows-x64 (ZIP de 7685290 bytes, contendo ZEUVORYX-LAB-Setup-3.0.9.exe conforme workflow). SHA-256 ZIP informado pelo GitHub: 4074aea09fda5f56687e9d8d65a6cff2b16bf9ac8f63bacb22d30d83a5ffb6f0; SHA-256 EXE instalador registrado no runner: 6b736fc1f422f6b3f72b7cd6853873c14b12d4762396589a77c362addaa29a1b. Não houve recálculo local dos bytes nesta auditoria. Artefato não expirado na consulta; acesso pode requerer login. Candidato somente para teste controlado; não promover a release, atualização assinada ou distribuição pública antes dos gates CP-001–004. Teste físico integrado em Windows 10/11, cinco ciclos reais e comparação de desempenho, eficácia real, estabilidade prolongada, dados reais, assinatura, licença/SBOM e paridade macOS continuam pendentes. Instalador LAB não empacota C2PA; não afirmar pacote integral de todas as 43 REQ. Inno Setup usado apenas no laboratório, sem contratação; alternativa gratuita comercial precisa de validação. Preservar Defender ativo, não desativar SmartScreen. REQ-022 adiado e REQ-043 suspenso. Não apagar histórico.

Metadados observados: descrição pública ainda 3.0.8 LAB under validation; topics vazios; homepage https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme. Proposta: ZEUVORYX Antivirus — Experimental Windows 3.0.9 LAB, CI validated, physical testing pending. macOS LAB separate. Complementary to Microsoft Defender. Topics propostos: windows, golang, security, malware-scanner, cybersecurity, antivirus-research, desktop. Escrita administrativa ainda indisponível; não afirmar correção efetiva.

Atualizado em 24/11/2026. Índice operacional do [memorando histórico completo](MEMORANDO_DESENVOLVIMENTO.md), não substituto. O histórico original é preservado. Status: P = pendente; E = experimental/parcial, não validado integralmente; V = evidência limitada de laboratório; A = adiado por decisão do titular; B = bloqueador. Não marcar concluído sem evidência no commit, testes e ambiente correspondente. IDs estáveis; nunca excluir linhas; registrar substituições explícitas e motivo.

| ID | Compromisso cumulativo | Estado | Prioridade / condição | Fonte no memorando |
|---|---|---|---|---|
| REQ-001 | Proteção efetiva, estabilidade, velocidade, baixo consumo, simplicidade, privacidade | B | Transversal; sem regressão | CP-001 |
| REQ-002 | Windows x64, instalação EXE simples, sem Python no cliente, início sem scan automático, atalho e operação simples | E | P0 abertura/instalação | Linha do tempo 2.1–3.0.6 |
| REQ-003 | Rápida, completa, personalizada, cancelar, progresso e exportação | E | P0 fechamento nas completas | EVO-013/015–018 |
| REQ-004 | SHA-256 integral, heurísticas, assinaturas, análise PE/ZIP, limites explícitos, erros e cobertura parcial | B | P0 base histórica candidata de 6 hashes; eficácia real, integração e cobertura pendentes | EVO-034/035 |
| REQ-005 | Quarentena reversível, restauração, confirmação de evidência, exclusões/falsos positivos e auditoria | E | P1 teste integrado, sem ação destrutiva automática | Cinco implementações futuras; 3.0.1–3.0.5 |
| REQ-006 | Remoção automática segura, interrupção de processos, persistência, reparação, bloqueio de execução | P | P1 módulos separados, consentimento/rollback; não confundir com Defender | Cinco implementações futuras |
| REQ-007 | Proteção residente, análise comportamental e antiransomware | P | P1 arquitetura, baixo consumo, compatibilidade Defender | Cinco implementações futuras |
| REQ-008 | IA consultiva local, evolução controlada, validação independente, falsos positivos | E | P1 sem promover modelo reprovado; simulações não são ataques reais | Linha do tempo 3.0.2/3.0.3; STD |
| REQ-009 | Inteligência de ameaças internacional, fontes autorizadas gratuitas, cache, expiração, revogação, indicadores e atualizações | B | P0 base histórica mínima; atualização ampla, direitos e validação pendentes; não confundir KEV com malware | EVO-006/007/034/035 |
| REQ-010 | Análise de links, phishing e proteção contra roubo de dados | E | P1 eficácia real pendente | 3.0.5 R5 |
| REQ-011 | Limpeza segura de temporários e métricas separadas de malware | E | P1 não remover arquivos ativos automaticamente | EVO-013 |
| REQ-012 | Histórico local autenticado, relatórios, salvar/exportar, privacidade e retenção | E | P1 preservação e relatório anonimizado pendentes | EVO-003/013; histórico local |
| REQ-013 | Controles de scan no topo do painel principal, gráfico e métricas junto aos controles, tela cheia e DPI | E | P1 validar UI efetivamente, sem perder handlers | Consolidação histórica; EVO-004 |
| REQ-014 | Olympus Silver, Zeus, ícones, arte local, assinatura manuscrita do criador, acabamento 3D leve | E | P1 preservar arte e autoria, verificar todos botões e idiomas | Identidade visual; EVO-011 |
| REQ-015 | Botão lateral Avaliação e resultados; cinco estrelas douradas fixas, prévia sem voto, escolha e confirmação única por instalação | P | P1 substituir seletor oculto, persistência/concorrência/migração | EVO-008/011/012 |
| REQ-016 | Painel local/mundial com contadores reais, média e distribuição, categorias, gráficos leves, sem números fictícios | P | P1 protótipo funcional; infraestrutura mundial depende de autorização | EVO-009–012 |
| REQ-017 | Serviço único entre GitHub, lojas, site, terminal, Windows/macOS; deduplicação, integridade, consentimento, TLS, criptografia, antifraude | P | P1 contrato LAB; produção bloqueada sem infraestrutura; métricas mínimas | EVO-009/010/012 |
| REQ-018 | Atualização in-place reconhecida por versão, preservação de dados, confirmação visual, link do pacote exato aprovado | E/B | P0 3.0.6 migração/abertura; catálogo exato pendente | EVO-012/018; CP-003 |
| REQ-019 | Atualização assinada por canal exclusivo do titular, chave oficial, manifesto, rollback e distribuição sem cadastro | E | P1 produção não habilitada; não enfraquecer assinatura | 3.0.5 R3/R4; EVO-012 |
| REQ-020 | PT-BR, inglês e espanhol, detecção Windows, seleção manual e fallback inglês; mais oito idiomas desejáveis | E | P2 11 idiomas sem sacrificar desempenho; validar os três atuais | Linha do tempo 3.0.2 |
| REQ-021 | Macintosh separado Intel/ARM64, paridade funcional progressiva, pacote e assinatura próprios | E | P1 não confundir compilação Intel com aprovação integral; consultar memo macOS | EVO-002; memorando Macintosh |
| REQ-022 | iOS e veículos: pesquisa futura, sem implantação agora | A | Retomar somente por decisão do titular | Sequência aprovada; prioridades |
| REQ-023 | Distribuição mundial inicialmente gratuita, sem publicidade; possibilidade de comercialização futura ou aquisição do projeto | P | Aviso prévio com data, preço e condições; aceitação expressa, sem cobrança retroativa; preservar direitos adquiridos | EVO-031/032/033; REQ-036/038/042/043 |
| REQ-024 | Autoria, controle exclusivo de código e publicação; registro/licença/patente e revisão jurídica | E | P1 segurança da conta, assinatura, permissões; revisão jurídica pendente | Controle exclusivo; liberação |
| REQ-025 | Testes repetidos, 5 ciclos integrados, 50 equivalência/desempenho quando viáveis, Windows 10/11 e máquinas reais, 4 GB/2 núcleos | B | P0; CI não basta | CP-001; EVO-018 |
| REQ-026 | Versionamento uniforme, documentação, metadados, release/artefato e memorando atualizado sem lembrete | E | P0; About administrativo bloqueado | CP-002/003 |
| REQ-027 | Diagnóstico do fechamento: logs persistentes, saída, memória, WebView2, recuperação, correlação de eventos | B | P0; causa raiz ainda não comprovada | EVO-015–018 |
| REQ-028 | Operação pelo iPhone, instruções completas copiar/colar quando necessário, GitHub e ZIP/EXE direto para compartilhamento | E | P1 reduzir trabalho manual; entrega apenas de pacote validado | Diretriz operacional; continuidade |

| REQ-029 | Aprendizagem local contínua e automática após consentimento, perfil comportamental, candidato separado e validação/rollback | P | Próxima revisão; baixo consumo, sem rótulos autogerados nem envio externo | EVO-019; REQ-008 |
| REQ-030 | Autostart consentido no login, residente com ícone de bandeja, único processo, opção de sair/desativar | P | Próxima revisão; sem scan completo automático, sem ocultação | EVO-019; REQ-002/007 |
| REQ-031 | Monitoramento de eventos de arquivos/downloads, escopos locais/nuvem acessível, fila limitada, análise de links autorizada | P | Próxima revisão; observação primeiro, sem hidratar nuvem ou interceptar TLS | EVO-019; REQ-007/010 |
| REQ-032 | Corrigir Diagnóstico local 3.0.3 e todos rótulos obsoletos; versão única e estados reais | P | Próxima revisão; gate de UI em todos idiomas | EVO-019; REQ-026 |
| REQ-033 | Preservar ClamAV opcional, sob demanda e dependente de instalação externa | E | Manter; medir custo e distinguir motores | EVO-019 |
| REQ-034 | Revisão integrada de toda fila após relatório 3.0.6; cinco ciclos e 50 equivalência quando viáveis | P | Próxima revisão, sem perder EVO-008–012 | EVO-019; REQ-025 |

| REQ-035 | Ícone ZEUS: preservar escudo; trocar apenas fundo/moldura escura por prata no software, alinhar recursos ICO/PE, taskbar, bandeja, atalhos e instalador | P | Próxima revisão; distinguir realce nativo Windows e validar DPI/temas | EVO-020; REQ-014/030 |
| REQ-036 | Preparar distribuição mundial nas lojas e canais oficiais após validação integrada, assinatura, revisão de políticas/licenças e autorização | P | Pós-validação; não publicar nem contratar automaticamente | EVO-020; REQ-023/024 |

| REQ-037 | Central pública de suporte internacional versionada, FAQ, ajuda no aplicativo e descoberta por buscadores | E | Base publicada em 4 idiomas; Pages, indexação e ajuda no app pendentes | EVO-021/022; REQ-023/036 |

| REQ-038 | Auditoria de propriedade intelectual, SBOM/licenças, termos de usuário, privacidade, responsabilidade e conformidade por mercado | P | Gate pré-lançamento; revisão jurídica e licenças comerciais ainda pendentes | EVO-024; REQ-023/024/036/037 |

| REQ-039 | Ajuda inteligente offline no software: lupa, perguntas naturais, busca de respostas verificadas e atualizadas, sem IA externa paga | E | Próxima versão; 124 entradas/4 idiomas, 28 consultas verificadas; CI e integração Windows pendentes | EVO-025/026; REQ-037 |
| REQ-040 | Auditoria de utilidade de módulos/pacotes, custos, dependências e consumo, consolidar distribuição sem eliminar funções solicitadas | P | Próxima revisão; verificar funções reais, evitar componentes sem uso, preservar evidência | EVO-025; CP-001/REQ-001/026/038 |

| REQ-041 | Continuidade da varredura completa com monitor apagado/bloqueado; solicitação temporária de energia, respeitar suspensão manual e bateria | E | Próxima versão; código candidato preparado; teste Windows físico e integração pendentes | EVO-027; REQ-003/025/027 |

| REQ-042 | Independência financeira e licenças comerciais gratuitas: sem royalties, assinaturas obrigatórias ou contratação não autorizada; avisos/SBOM e alternativas verificadas | E | NSIS fixture e bloqueio de consulta não verificada preparados; migração, cobertura, CI e licenças integrais pendentes | EVO-029/030; REQ-009/038/040 |

| REQ-043 | Publicidade, patrocínios e espaço promocional | A | Suspenso por decisão do titular; sem anúncios no app/site de suporte; retomar somente por autorização expressa | EVO-030/031 |

## Controle obrigatório de continuidade

A cada tarefa, ler o índice e o memorando, identificar requisitos afetados, registrar novos pedidos imediatamente (inclusive "no próximo"), preservar IDs e estados dos demais. Antes de fechar, comparar o índice anterior com o novo, justificar qualquer mudança de prioridade/estado, anexar evidência ou marcar pendência, atualizar a main e a ramificação/plataforma pertinente e confirmar gravação remota. Correção P0 pode adiar implantação de P1/P2, mas não apagar nem declarar cumprida. Decisões posteriores explícitas prevalecem sobre as substituídas, cujo histórico permanece. A fila não é autorização para lançar, contratar ou realizar ações destrutivas. Reconciliar requisitos não localizados com o histórico e pedidos recuperáveis, sem alegar cobertura absoluta de conversas inacessíveis.

## Situação de referência

Windows 3.0.6 LAB R1: candidato isolado, execução 45479946405 em andamento na consulta inicial. Prioridade atual REQ-002/003/018/025/027: instalação, abertura e fechamento inesperado. Última execução anterior aprovou 60 varreduras de corpus pequeno e compilação, mas reprovou no gate textual de versão do instalador; correção de metadados em 1ecc4e5893615ee07ea7548a6833571949e8a2d4. Cinco ciclos integrados, Windows físico e eficácia independente permanecem pendentes. Macintosh 3.0.5 LAB preservado. Nenhuma nova publicação.

### Evidência atualizada EVO-018

REQ-002/018: migração e abertura de 15 segundos aprovadas no CI, estado E (piloto físico pendente). REQ-003/025/027: 60 testes de raiz controlada aprovados, estabilidade física e cinco ciclos integrados pendentes; não encerrar bloqueadores. REQ-026: pacote 3.0.6 LAB R1, hash verificado; metadados públicos dependem de acesso administrativo e aprovação de promoção. Todos os demais IDs preservados. [CI aprovado](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45479946405).

### Próxima evolução EVO-019 — aguardando relatório do piloto 3.0.6

REQ-029 a 034 adicionados conforme solicitação do titular. REQ-008 e REQ-007 ampliados, não encerrados; REQ-026 inclui correção do rótulo herdado 3.0.3. REQ-015 a 019 e demais pendências continuam na fila. Nenhum novo binário nesta revisão documental.

### Complemento EVO-020 — ícone prata e lojas

REQ-035/036 adicionados, mantendo os 34 anteriores e as quatro cláusulas pétreas. Próxima versão funcional prevista 3.0.7 após relatório do piloto 3.0.6. Correção de arte no software pendente; não alterar aparência global do Windows. Preparação de distribuição somente após gates e autorização. Nenhum novo binário nesta revisão.

### EVO-022 — suporte mundial

[Central pública](support/README.md) e HTML estático preparados; não declarar indexação nem site ativo. REQ-037 parcial, demais requisitos preservados.

### Escopo integral da próxima revisão — EVO-033

Revisar todos os 43 IDs, implementar as melhorias vigentes aplicáveis à próxima versão e reconciliar evidências por requisito. REQ-022 adiado, REQ-043 suspenso. Não omitir requisitos ativos; registrar dependências, bloqueios e justificativas sem aprovar por promessa. Cláusula comercial consolidada no memorando; gratuidade inicial sem anúncios, possibilidade futura de oferta paga com data, preço, condições e aceitação expressa. Nenhuma mudança funcional nesta revisão.


### ZVX-PILOTO-307-20261124 — teste rápido parcial e preparação 3.0.8

REQ-003/004/018/025/027 permanecem abertos: CI 3.0.7 aprovado, exportação após cancelamento comprovada no teste físico parcial; 19.701 arquivos, 13,05 GiB, 623 s, 22,48 MB/s, três bloqueios de compartilhamento, um cancelamento voluntário, zero falhas internas e zero testes positivos. Não concluir eficácia, desempenho comparativo ou estabilidade integral. REQ-026: entrega da 3.0.7 portátil registrada no memorando; README/CHANGELOG e proposta About sincronizados; campo administrativo real não editável pelo conector. Ramificação 3.0.8 realinhada à base funcional 3.0.7 corrigida, sem build novo. Demais IDs e pendências preservados. Obrigatório conferir descrições a cada intervenção sem solicitação adicional.


## Auditoria cumulativa 22–24/11/2026 — 3.0.8 LAB R3 integrada (candidata)
Reconciliados memorando central, matriz de 43 REQ (REQ-022 adiado, REQ-043 suspenso) e EVO-038–041. Usuário rejeitou R1 por ausência do painel e de atividade visível da IA. R2 introduziu CPU/RAM/disco/relógio nativos, CI separado 45515168910. R3 preserva R2 e acrescenta quarto quadro dos contadores, capacidade/ocupado/livre dos volumes fixos com letra acessíveis e deduplicados por GUID, contagem de omitidos; volume analisado reaproveita contador existente; total elegível/restante indeterminado até enumeração confiável. Integra ajuda local 124 artigos PT/EN/ES/JA, 140 verificações de busca, aviso de cobertura de assinaturas reais (sem adicionar base não validada), solicitação de energia durante scan completo inicialmente na tomada, sem manter tela ligada. IA: contadores reais de inferências sobre evidência heurística, concluídas/falhas/última atividade/índice/modelo, animação quando há trabalho observado, autoteste sintético acionável e separado, sem fabricar números de detecção/treino. Modelo permanece consultivo; aprendizagem contínua não implantada. Gate exige presença do painel IA/ajuda/armazenamento e renderização de métricas CPU/RAM, além de testes de contadores concorrentes e regressões anteriores. CI, validação visual/física e eficácia pendentes. Mantém limpeza .tmp confirmada existente, não declara central de manutenção completa. Não publicar ou gerar instalador. Sem contratação paga. R3 isolada para não cancelar CI R2 nem alterar 3.0.7/main funcional/Macintosh. Pendências explícitas: IA contínua validada, residente/autostart/bandeja, antiransomware/bloqueio, memória/processos, base ampla de assinaturas, cache/pausa/retomada, alertas e exportação durante scan, ETA/cobertura, otimização completa, cinco estrelas/estatísticas globais, atualização assinada de produção e instalador, ícone PE, idiomas adicionais, licenças/lojas e cinco ciclos físicos. Não declarar conclusão geral.

## 3.0.8 LAB R4 — correção de interpretação: Segurança em camadas (24/11/2026)
Titular esclareceu inequivocamente que não pediu câmera: a solicitação é na tela Segurança em camadas, com Consultar proteção, Abrir Segurança do Windows, Verificar Windows Update. A ramificação equivocada windows-3.0.8-camera-lab NÃO deve ser integrada, distribuída nem usada como base; seu requisito de câmera foi invalidado pelo titular. R4 correta parte da revisão integrada R3 e mantém todos seus componentes. Cada linha do diagnóstico existente recebe ponto colorido e texto original: verde ativado/bloqueio, vermelho desativado/não verificado/consulta indisponível, amarelo auditoria/aviso/modo intermediário, azul informação técnica. BitLocker On/Off normalizado para Ativado/Desativado. RDP desativado pode ser configuração desejável; cor descreve estado, não risco ou vulnerabilidade. Consulta somente leitura, sem alterar políticas, permissões, Defender ou Windows. Teste Node de classificação (13 casos x 5) e controles preservados, gate de identidade sem painel de câmera. CI e validação visual/física pendentes. Nenhum instalador ou lançamento. Manter memorando integral e demais pendências da R3.

## ZEUVORYX 3.0.8 LAB R5 — IA ativa e mascote funcional (24/11/2026)
Diagnóstico de fonte: a IA anterior era chamada exclusivamente após alerta heurístico com vetor não nulo; uma varredura sem alertas podia executar zero inferências. R5 candidata desloca a chamada real de adaptive.Observe para antes do limiar de alerta, avaliando o vetor de características já calculadas em cada chamada concluída de analyze, inclusive arquivos sem alerta. Não relê arquivos, não altera pontuação de risco, veredito, quarentena ou política de proteção. Contadores medem avaliações de modelo, não arquivos únicos (ZIP e releituras podem gerar chamadas adicionais). Autoteste sintético único na abertura, separado das inferências de arquivo. Mascote SVG/CSS leve na visão geral e Motor e IA: em pé levantando pesos apenas quando o backend indica avaliação em andamento ou aumento real no contador entre amostras; deitado descansando em espera, falha identificada separadamente; texto, contador, última avaliação, versão, sucesso, falhas e autoteste. Amostragem de 2 s nas páginas visíveis; respeita movimento reduzido. Nenhuma animação representa proteção comprovada, aprendizado contínuo ou treino fictício. Aprendizado persistente e proteção residente continuam pendentes (REQ-029/030/031), assim como gates de eficácia e desempenho. Teste integrado exige três avaliações reais em três arquivos benignos sem produzir alerta; regressões anteriores mantidas. CI/visual/físico pendentes. R5 parte da R4 correta de segurança em camadas, sem recurso de câmera. Nenhuma distribuição, contratação ou instalador. Não entregar como aprovado até verificar resultado e pacote.


## ZEUVORYX 3.0.9 LAB — identidade funcional autorizada (24/11/2026)
Titular determinou 3.0.9, não 3.0.8 R5. Novo candidato independente windows-3.0.9-lab, derivado da revisão cumulativa 3.0.8 R5 (a96d1e70dc1ce6350413ea89efce1b95752dc5c8) ainda não aprovada. Preserva painel CPU/RAM/disco/relógio, quarto quadro de armazenamento, segurança em camadas com pontos coloridos (sem câmera), ajuda offline, energia e IA consultiva em cada análise com mascote animado conforme atividade real. Identidade atualizada no motor, UI, atualizador, auditoria (preservado diretório histórico), logs, workflow, links e nomes de artefatos; gates de identidade e regressões existentes mantidos. Nomes internos históricos de arquivos e testes permanecem para rastreabilidade. CI e validação física pendentes; nenhuma afirmação de eficácia, aprendizado contínuo ou proteção residente comprovada. Pacote portátil LAB somente, sem instalador ou release pública; 3.0.8 e 3.0.7 preservadas. Pendências cumulativas continuam vigentes. Descrição pública proposta: ZEUVORYX Antivirus — Experimental Windows 3.0.9 LAB under validation. macOS LAB separate. Complementary to Microsoft Defender. Escrita administrativa About/Topics indisponível no conector; não afirmar alteração real.


## 3.0.9 LAB — alertas, cobertura, falhas e histórico durante varredura (24/11/2026)
Reiteração de EVO-041: alertas e falhas só apareciam no fim; a tela desabilitava exportação durante execução e cobertura aparecia fora do quadro. Novo candidato isolado windows-3.0.9-live-lab, sem interromper CI anterior. Motor publica snapshots incrementais durante varreduras rápida/completa e customizada, com mutex e cópia de novos alertas/erros; contador real, estado parcial, escopo, volume, ignorados e falhas. Interface limita visualização em execução aos últimos 500 alertas e 20 erros; exportação JSON conserva todos os detalhes retidos pelo motor, inclusive limites de erros existentes, com snapshot_at/in_progress/partial; fim/cancelamento mantêm auditoria final. Histórico exibe contagens e exporta relatório salvo autenticado, somente local. Quadro cobertura e falhas reposicionados dentro da visão geral, quebra de linhas e área de detalhes rolável; sem porcentagem inventada. Não confundir itens ignorados com ameaça. Testes 5 ciclos de snapshot, exportação e alertas antes do fim, custom scan e regressões cumulativas. CI/validacao visual, concorrência e impacto físico pendentes; não declarar aprovação. Sem retomada após crash nesta alteração, nem garantia de cobertura integral; REQ-003/012/027/EVO-041 permanecem abertos até validação. Sem publicação ou instalador.


## 3.0.9 LAB — restaurar instalador e instância única (24/11/2026)
Titular observou 3.0.6 e 3.0.8 abertas simultaneamente. Confirmado: pacote 3.0.8 portátil não substitui instalação anterior. Restaurado candidato de instalador Inno Setup gratuito com AppId ZEUVORYX.Desktop.Lab303, pasta anterior, atalhos e dados preservados, sem varredura automática. Mutex por sessão comum às versões futuras, verificação de processo legado de nome ZEUVORYX.exe na mesma sessão, sem finalizar processos à força; nova versão recusa abertura enquanto outra estiver aberta. Versões portáteis antigas não possuem guarda e ainda podem abrir depois da nova; não apagar cópias em Downloads nem prometer controle retroativo. Orientar fechar as versões anteriores. CI deverá baixar artefato histórico 3.0.6 14960924262 com SHA256 conferido, instalar 3.0.6, testar bloqueio de nova instância enquanto antiga está aberta, encerramento voluntário, cinco ciclos de instalação/reinstalação, versão/registro/hash/desinstalador único e preservação de marcador local, bloqueio com 3.0.9 aberta. Somente candidato LAB sem publicação/assinatura automática; testes físicos, dados reais de quarentena/histórico e recuperação pendentes. Não alterar Defender nem desativar SmartScreen. Preservar correções de relatório ao vivo e demais requisitos. Branch windows-3.0.9-install-lab independente, CI anterior preservado. About/Topics administrativos ainda indisponíveis.
