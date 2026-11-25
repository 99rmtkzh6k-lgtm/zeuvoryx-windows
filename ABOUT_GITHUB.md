> **25/11/2026 — protocolo de continuidade e evidência:** Windows 3.0.11 LAB, commit a1ce50f7e7379e48911e6ac55d80f76eb950d3a1, [CI 45599586072](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45599586072) concluída com sucesso. ClamAV 1.5.4, base daily 28194, 3.628.228 assinaturas carregadas; cinco testes EICAR inofensivos e cinco controles benignos aprovados em Linux. Instalador LAB artefato 15015960060 (não é release de produção; ClamAV externo obrigatório). Windows 3.0.12 LAB, commit a56da9881ebc448d55e49d0ecf7f85dd511f1487, [CI 45600726308](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45600726308) em andamento: segunda etapa ClamAV na varredura normal, relatório separado por motor, base recente, testes reais no Windows e Linux. Não afirmar aprovação física ou eficácia contra malware real. Ramificação 3.0.13 reservada, ainda sem alterações funcionais próprias; não promover antes da validação anterior. Exigência do titular: proteção independente de Windows Update, complementar ao Defender, fontes reais e gratuitas compatíveis com distribuição comercial, comportamento correlacionado com evidência (não mera suspeita), revalidação e SHA-256 antes de quarentena reversível com consentimento, restauração e exclusão confirmada; proteger arquivos críticos e nunca apagar automaticamente por heurística. YARA-X, IOC licenciado, comportamento/tempo real, empacotamento e FreshClam integrado pendentes. Preservar CP-001–004 e 43 REQ, cinco ciclos de regressão, velocidade/estabilidade/consumo, documentação e versionamento sincronizados. Sem promoção de candidato, sem estatística inventada. Metadados administrativos públicos não atualizados por esta operação.

# Metadados públicos — conferência obrigatória a cada revisão


## Regra permanente — supervisão de evolução da IA e evidência 3.0.10 (24/11/2026, Brasília)
Titular solicita um sistema que cobre continuamente o avanço da IA. REQ-008/025/029/034: em cada revisão, gerar matriz de pendências e evidências, comparar modelo vigente/candidato em dados rotulados de origem autorizada separados por família/tempo, medir sensibilidade, falsos positivos, custo CPU/RAM/latência, verificar proveniência e contaminação, rejeitar regressão, manter rollback e aprovação explícita do titular antes de ativação. Sem rótulos autogerados, sem treino irrestrito em arquivos do usuário, sem envio de dados por padrão, sem apresentar simulações como eficácia real. Priorizar ferramentas gratuitas e estabilidade. Proposta de supervisor automático de evidências e relatório de lacunas; ainda não implementado, não confundir com aprendizagem contínua. Manter todos os 43 REQ e CP-001–004.
CI 3.0.10 LAB https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45524131517 commit 8b36c26163a9950a12a1d3d9d4e66cefa15a5a7d: 44/44 etapas aprovadas. Testes Test310DriveArguments, Test310MaintenanceAllowlist e Test310SelectedTempCleanup aprovados em repetições; migração 3.0.6/3.0.9, cinco ciclos de instalação, guarda de instância e startup verificados no workflow. Artefato https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45524131517/artifacts/14982250424 ZIP SHA256 63d6d5d34702b0363ad5a9937b73ac437905b9aa54bcaf0b4999f9d13f03a621, instalador SHA256 2a3b51f85836496ccb642835d7c9c97253b245e84100129556b79cb12fcc8552 informado pelo runner. Validação física, execução real da otimização e medição de ganho, eficácia da IA e aprendizagem contínua pendentes. Sem publicação de produção. Proposta administrativa: ZEUVORYX Antivirus — Experimental Windows 3.0.10 LAB, CI validated, physical testing pending. macOS LAB separate. Complementary to Microsoft Defender. Topics não impeditivos, escrita administrativa indisponível.


## 3.0.10 LAB — central integrada de otimização (24/11/2026)
Solicitação do titular: não apenas abrir ferramentas Windows; diagnóstico CPU/RAM, escolha e confirmação dos arquivos, desfragmentação/otimização de unidade, sem compressão automática. Commit candidato 8b36c26163a9950a12a1d3d9d4e66cefa15a5a7d em windows-3.0.10-lab. Diagnóstico por duas amostras e limiares, sem afirmar causa; prévia limitada a temporários .tmp antigos em LocalAppData/Temp, seleção desmarcada, exclusão via handle revalidado com confirmação, resultado removidos/ignorados/bytes lógicos; proteção de arquivos modificados/bloqueados e plano vencido. Otimização do volume Windows pelo defrag.exe nativo /O /U /V, UAC explícito, tomada, sem varredura concorrente, código real de saída, sem percentual inventado, sem ganho garantido. Não comprimir NTFS, apagar documentos, forçar RAM ou alterar registro. Acesso nativo a programas de inicialização preservado. Validação de CI e física pendentes, não afirmar aprovação; versão 3.0.9 permanece referência CI. REQ-011/EVO-039 parcial até medições físicas. Preservar 43 REQ e CP-001–004. Fonte https://learn.microsoft.com/windows-server/administration/windows-commands/defrag . Metadados administrativos mantêm descrição real da versão aprovada 3.0.9.


## 3.0.10 LAB — candidato em validação (24/11/2026)
Solicitações após vídeos 3.0.9: separar autoteste da IA consultiva, saldo estimado do espaço ocupado menos bytes contabilizados, velocidade medida, alertas visuais/sonoros por novos achados, abertura junto ao Windows sem scan automático, e central de otimização. Código na ramificação windows-3.0.10-lab, commit 6605632e0a8103c4ef8bb4c0d16c2e04c76ab932. Ferramentas inclui limpeza temporária existente com prévia/confirmação, acesso a Otimizar Unidades, Limpeza de Disco e Gerenciador de Tarefas; Windows decide TRIM/defrag, usuário confirma; não é motor próprio completo. Testes de allowlist, taxas/alertas, startup e migração programados; CI ainda não concluído. Sem promessa de aceleração, aprendizagem contínua ou eficácia real. 3.0.9 permanece referência CI aprovada e teste físico parcial. Topics não impeditivos conforme titular; descrição pública permanece 3.0.9 até evidência de nova CI. CP-001–004 e 43 REQ preservados. Sem disponibilizar candidato como validado antes dos gates e sem alterar binário de teste 3.0.9.


## Estado verificado — ZEUVORYX 3.0.9 LAB com instalador (24/11/2026, Brasília)

CI [45519151028](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028), commit funcional 5115d20d53ab6fb7867f72d1b122e7a89a83391f, job 138886281969: 44/44 etapas concluídas com sucesso. Logs confirmam migração do instalador histórico 3.0.6 para 3.0.9, cinco ciclos de instalação/reinstalação, AppId e pasta mantidos, um desinstalador, hash do executável instalado igual ao compilado, preservação de marcador de dados e bloqueio de segunda instância (incluindo legado aberto primeiro). Teste de recuperação de quarentena e histórico reais ainda pendente; não há controle retroativo de portáteis legados abertos depois. 40 aprovações do teste IA em arquivos benignos; 45 aprovações dos três testes de relatório ao vivo; 15 aprovações do teste de instância única; 65 verificações de cores; 50M cenários sintéticos IA, QA20, 10M fraude, 50 raízes controladas, EICAR, regressões e abertura da interface com painel CPU/RAM. Não equivale a eficácia real ou certificação.

Artefato exato de laboratório [14979149692](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45519151028/artifacts/14979149692), ZEUVORYX-3.0.9-LAB-SETUP-TESTE-Windows-x64 (ZIP de 7685290 bytes, contendo ZEUVORYX-LAB-Setup-3.0.9.exe conforme workflow). SHA-256 ZIP informado pelo GitHub: 4074aea09fda5f56687e9d8d65a6cff2b16bf9ac8f63bacb22d30d83a5ffb6f0; SHA-256 EXE instalador registrado no runner: 6b736fc1f422f6b3f72b7cd6853873c14b12d4762396589a77c362addaa29a1b. Não houve recálculo local dos bytes nesta auditoria. Artefato não expirado na consulta; acesso pode requerer login. Candidato somente para teste controlado; não promover a release, atualização assinada ou distribuição pública antes dos gates CP-001–004. Teste físico integrado em Windows 10/11, cinco ciclos reais e comparação de desempenho, eficácia real, estabilidade prolongada, dados reais, assinatura, licença/SBOM e paridade macOS continuam pendentes. Instalador LAB não empacota C2PA; não afirmar pacote integral de todas as 43 REQ. Inno Setup usado apenas no laboratório, sem contratação; alternativa gratuita comercial precisa de validação. Preservar Defender ativo, não desativar SmartScreen. REQ-022 adiado e REQ-043 suspenso. Não apagar histórico.

Metadados observados: descrição pública ainda 3.0.8 LAB under validation; topics vazios; homepage https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme. Proposta: ZEUVORYX Antivirus — Experimental Windows 3.0.9 LAB, CI validated, physical testing pending. macOS LAB separate. Complementary to Microsoft Defender. Topics propostos: windows, golang, security, malware-scanner, cybersecurity, antivirus-research, desktop. Escrita administrativa ainda indisponível; não afirmar correção efetiva.

Última conferência: 24/11/2026. Repositório compartilhado ZEUVORYX Windows e Macintosh (LAB).

## Valores efetivamente observados

- Description: ZEUVORYX — Experimental Windows security scanner in Go. SHA-256 indicators, heuristics, advisory AI, file quarantine and Olympus Silver interface. Complementary to Microsoft Defender. 3.0.5 LAB R5.
- Topics: vazio.
- Website/Homepage: vazio.
- Ramificação padrão: main.
- Visibilidade: pública; não alterar sem autorização.

## Valores propostos — ação manual pendente

Description (Descrição):
> ZEUVORYX — experimental Windows 3.0.5 LAB R5 security scanner and macOS LAB candidate. SHA-256, heuristics, advisory threat intelligence, reversible quarantine and Olympus Silver. Complementary to OS security; not certified.

Topics (Tópicos), separados por espaços:
`cybersecurity windows macos golang swift swiftui malware-analysis security-scanner antivirus file-scanner heuristics quarantine sha256 desktop-application webview2 github-actions experimental threat-intelligence`

Website:
https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme

## Procedimento no iPhone (Safari)

1. Abra https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows autenticado como titular.
2. Na página principal, localize About/Sobre e toque na engrenagem; se não aparecer, solicite a versão para computador no menu do Safari.
3. Cole a Description, o Website e os Topics nos respectivos campos; confirme cada tópico se a interface solicitar. Salve as alterações.
4. Não alterar visibilidade, permissões, licença, opções de distribuição ou publicações sem autorização. Releases de produção não aprovadas não devem ser anunciadas.
5. Após salvar, conferir novamente os valores efetivos no GitHub. Este arquivo não modifica automaticamente os campos administrativos.

A integração disponível não oferece operação de edição administrativa de About/Topics/Website. O memorando permanente exige atualização e verificação após toda revisão, com registro de impedimentos e ações manuais.

- ZVX-EVO-004: candidato incremental Windows 3.0.5 LAB identificado por commit; descrição pública continua referindo-se ao R5 previamente aprovado até nova aprovação. Não promover candidato por meio dos metadados.

- ZVX-EVO-007: valores reais reconferidos 24/11/2026; descrição antiga mantida pelo GitHub, tópicos e homepage vazios, main pública, has_downloads=false, zero releases. CI Windows 45447823825 aprovado; consulta DROP consultiva compilada, seed real não obtido. Texto proposto atualizado sem alegar bloqueio residente ou paridade macOS. Exibições de Releases/Packages/Deployments: conferir na engrenagem About pelo titular; sem release autorizada não anunciar lançamento. Nenhuma edição administrativa realizada pela integração.

## Proposta de metadados da próxima versão (EVO-018)

Descrição para atualização pública quando o candidato Windows for aprovado: ZEUVORYX — Experimental Windows 3.0.6 LAB R1 security scanner in Go; macOS 3.0.5 LAB candidate. SHA-256, heuristics, reversible quarantine and advisory AI. Complementary to Microsoft Defender; not certified.

Topics sugeridos: zeuvoryx, antivirus, security, windows, golang, malware-scanner, cybersecurity, macos. Website: manter vazio até existir página oficial aprovada. Não promover versão por metadados antes da validação. Campos reais consultados: descrição 3.0.5 LAB R5, topics vazios, homepage vazia; edição administrativa não disponível nesta integração. Caminho: página principal do repositório > About > engrenagem.

## Estado e texto proposto vigente — piloto 3.0.7 / preparação 3.0.8

24/11/2026. Candidato Windows 3.0.7 LAB com CI aprovado e piloto rápido parcial; 3.0.8 apenas preparação. Não anunciar proteção residente, certificação ou versão definitiva. Obrigatório verificar as descrições em toda revisão, sem novo lembrete do titular.

Description (copiar integralmente):
> ZEUVORYX — Experimental Windows security scanner in Go. 3.0.7 LAB: CI passed, validation ongoing; 3.0.8 in preparation. macOS LAB separate. Complementary to Microsoft Defender; not certified.

Topics: `zeuvoryx antivirus security windows golang malware-scanner cybersecurity macos experimental`

Website: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme

Campos administrativos não editáveis pelo conector atual; titular pode aplicar em About/Sobre > engrenagem. A gravação deste arquivo não altera os campos efetivos. Revalidar após alteração. Nenhuma publicação autorizada.


## Proposta vigente — 3.0.8 LAB intervenção A

Description: ZEUVORYX — Experimental Windows security scanner in Go. 3.0.7 LAB pilot partial; 3.0.8 LAB candidate under validation. macOS LAB separate. Complementary to Microsoft Defender; not certified.

Topics: zeuvoryx antivirus security windows golang malware-scanner cybersecurity macos experimental

Website: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme

Campos administrativos ainda não alterados pelo conector. Verificar sempre após cada evolução; não divulgar release definitiva. Evidência: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/afdbfc6c4f2258171ece292e0c786d83930a25d5.
