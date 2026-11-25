# Metadados públicos — conferência obrigatória a cada revisão

Última conferência: 24/11/2026. Repositório compartilhado ZEUVORYX Windows e Macintosh (LAB).

## Valores efetivamente observados

- Description: ZEUVORYX Antivirus — Experimental Windows security scanner in Go. 3.0.8 LAB under validation. macOS LAB separate. Complementary to Microsoft Defender.
- Topics: vazio.
- Website/Homepage: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows#readme
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

- EVO-018: 3.0.6 LAB R1 é candidato, não substituir descrição pública aprovada antes dos gates. Macintosh permanece 3.0.5 LAB. A integração não permite edição administrativa dos campos públicos.

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

Topics e Website conforme proposta anterior; campos administrativos pendentes de edição e verificação. Nenhuma release definitiva.


## Conferência vigente — CI 3.0.8 LAB R1

24/11/2026 — CI 3.0.8 LAB R1 aprovado na [execução 45513126955](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45513126955), commit bfdc766b80b75a13c8d2198fa7356d96df4ec8ee. ZIP portátil x64 exclusivamente para teste rápido autorizado: [artefato 14976792219](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45513126955/artifacts/14976792219), 6.883.070 bytes; SHA-256 do ZIP conferido nos bytes baixados: 9c5f62efc583f6e8fe18691c008f8bf495cd9b94fcacbf58431e74e94ae6eb5e; CRC íntegro. SHA-256 do EXE confere com o manifesto: 98d07e58e2780d3d1ee0211f9d65869aa09928693e0dd5272c9f5f7f93110cbe. CI aprovou regressões de leitura/PE (30 rodadas cada), 50 regressões de varredura, 10 repetições, 5 rodadas de links, EICAR, cenários sintéticos, build e smoke desktop. Não equivale a proteção real certificada. Testes físicos, desempenho comparativo, cinco ciclos integrados e demais gates pendentes; nenhum instalador, atualização assinada ou release autorizada. Manter Defender e SmartScreen ativos, extrair ZIP em pasta separada e executar apenas rápida em ambiente autorizado.

Campos efetivos conferidos via API: descrição 3.0.8 LAB under validation, homepage preenchida, tópicos vazios, main padrão. Descrição e homepage efetivas coerentes com candidato, sem alegar aprovação física. Tópicos ainda pendentes; integração não expõe escrita administrativa. Nenhuma edição administrativa realizada neste fechamento.


## 3.0.8 LAB R2 — correção do painel de desempenho (24/11/2026)
Titular recusou R1 por não cumprir painel de desempenho. R2 candidata: painel na visão geral após os comandos, medição nativa Windows de CPU por delta GetSystemTimes, RAM física total/disponível/em uso por GlobalMemoryStatusEx, capacidade do volume Windows por GetDiskFreeSpaceExW, relógio local, medidores semicirculares, PT/EN/ES, sem rede, sem varredura obrigatória, amostra a cada 3 s apenas na página visível, cache nativo de 2 s e disco 15 s. Dados indisponíveis não viram zeros. CPU inicial aguarda segunda amostra. Disco é capacidade, não atividade I/O; RAM disponível não é análise de conteúdo de memória. Mantém motor, controles e versões anteriores, sem instalador. Gate adicional: controles únicos, sintaxe JS, testes Go nativos e marcador de renderização CPU/RAM real no smoke Windows. CI e validação visual/física pendentes. R1 não atende a este requisito; não oferecer como R2. R2 é revisão LAB da base 3.0.8, não release de produção nem atualização automática. Demais requisitos cumulativos continuam pendentes até evidência específica (processos/memória, manutenção, estatísticas globais, proteção residente, validação integrada etc.).

Proposta administrativa: 3.0.8 LAB R2 candidate under validation. Tópicos permanecem pendentes; descrição administrativa não editável por este conector.


## Auditoria cumulativa 22–24/11/2026 — 3.0.8 LAB R3 integrada (candidata)
Reconciliados memorando central, matriz de 43 REQ (REQ-022 adiado, REQ-043 suspenso) e EVO-038–041. Usuário rejeitou R1 por ausência do painel e de atividade visível da IA. R2 introduziu CPU/RAM/disco/relógio nativos, CI separado 45515168910. R3 preserva R2 e acrescenta quarto quadro dos contadores, capacidade/ocupado/livre dos volumes fixos com letra acessíveis e deduplicados por GUID, contagem de omitidos; volume analisado reaproveita contador existente; total elegível/restante indeterminado até enumeração confiável. Integra ajuda local 124 artigos PT/EN/ES/JA, 140 verificações de busca, aviso de cobertura de assinaturas reais (sem adicionar base não validada), solicitação de energia durante scan completo inicialmente na tomada, sem manter tela ligada. IA: contadores reais de inferências sobre evidência heurística, concluídas/falhas/última atividade/índice/modelo, animação quando há trabalho observado, autoteste sintético acionável e separado, sem fabricar números de detecção/treino. Modelo permanece consultivo; aprendizagem contínua não implantada. Gate exige presença do painel IA/ajuda/armazenamento e renderização de métricas CPU/RAM, além de testes de contadores concorrentes e regressões anteriores. CI, validação visual/física e eficácia pendentes. Mantém limpeza .tmp confirmada existente, não declara central de manutenção completa. Não publicar ou gerar instalador. Sem contratação paga. R3 isolada para não cancelar CI R2 nem alterar 3.0.7/main funcional/Macintosh. Pendências explícitas: IA contínua validada, residente/autostart/bandeja, antiransomware/bloqueio, memória/processos, base ampla de assinaturas, cache/pausa/retomada, alertas e exportação durante scan, ETA/cobertura, otimização completa, cinco estrelas/estatísticas globais, atualização assinada de produção e instalador, ícone PE, idiomas adicionais, licenças/lojas e cinco ciclos físicos. Não declarar conclusão geral.
Proposta administrativa: ZEUVORYX — Experimental Windows 3.0.8 LAB R3 under validation; macOS separate. Complementary to Microsoft Defender. Tópicos administrativos pendentes.

## 3.0.8 LAB R4 — REQ-044 segurança de câmera e três indicadores (24/11/2026)
Solicitação reiterada pelo titular: recuperar quadro de segurança da câmera e pontos azul/vermelho/amarelo. A conversa exata e a legenda original não foram recuperadas; legenda provisória explícita: azul = nenhum uso identificado nos registros consultados; amarelo = indeterminado/sem dados suficientes; vermelho = indício de uso, não invasão confirmada. Candidato consulta metadados ConsentStore webcam somente leitura do usuário atual (formato interno não garantido, escopo limitado, até 512 chaves), sem capturar imagem, microfone, nomes de aplicativos, caminhos, envio externo, privilégios ou alteração de permissões. Não abrir câmera. Link explícito abre Configurações de privacidade do Windows. Ausência de registro não prova câmera desligada. Bloqueio, proteção em tempo real, detecção de acesso não autorizado e outros usuários/dispositivos não comprovados. Referência: https://support.microsoft.com/pt-br/windows/privacy/windows-camera-microphone-and-privacy . Gate CI: classificação com registros inválidos/ausentes/ativos/incompletos, teste read-only, presença do painel e dos bindings; teste real Windows 10/11 com câmera ligada e desligada e uso autorizado ainda necessário. Revisão isolada a partir da R3, preservando seu CI e main funcional; nenhuma publicação. Demais requisitos permanecem abertos conforme matriz.