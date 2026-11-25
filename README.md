# ZEUVORYX Antivirus — Windows e Macintosh LAB


## Auditoria cumulativa 22–24/11/2026 — 3.0.8 LAB R3 integrada (candidata)
Reconciliados memorando central, matriz de 43 REQ (REQ-022 adiado, REQ-043 suspenso) e EVO-038–041. Usuário rejeitou R1 por ausência do painel e de atividade visível da IA. R2 introduziu CPU/RAM/disco/relógio nativos, CI separado 45515168910. R3 preserva R2 e acrescenta quarto quadro dos contadores, capacidade/ocupado/livre dos volumes fixos com letra acessíveis e deduplicados por GUID, contagem de omitidos; volume analisado reaproveita contador existente; total elegível/restante indeterminado até enumeração confiável. Integra ajuda local 124 artigos PT/EN/ES/JA, 140 verificações de busca, aviso de cobertura de assinaturas reais (sem adicionar base não validada), solicitação de energia durante scan completo inicialmente na tomada, sem manter tela ligada. IA: contadores reais de inferências sobre evidência heurística, concluídas/falhas/última atividade/índice/modelo, animação quando há trabalho observado, autoteste sintético acionável e separado, sem fabricar números de detecção/treino. Modelo permanece consultivo; aprendizagem contínua não implantada. Gate exige presença do painel IA/ajuda/armazenamento e renderização de métricas CPU/RAM, além de testes de contadores concorrentes e regressões anteriores. CI, validação visual/física e eficácia pendentes. Mantém limpeza .tmp confirmada existente, não declara central de manutenção completa. Não publicar ou gerar instalador. Sem contratação paga. R3 isolada para não cancelar CI R2 nem alterar 3.0.7/main funcional/Macintosh. Pendências explícitas: IA contínua validada, residente/autostart/bandeja, antiransomware/bloqueio, memória/processos, base ampla de assinaturas, cache/pausa/retomada, alertas e exportação durante scan, ETA/cobertura, otimização completa, cinco estrelas/estatísticas globais, atualização assinada de produção e instalador, ícone PE, idiomas adicionais, licenças/lojas e cinco ciclos físicos. Não declarar conclusão geral.


## 3.0.8 LAB R2 — correção do painel de desempenho (24/11/2026)
Titular recusou R1 por não cumprir painel de desempenho. R2 candidata: painel na visão geral após os comandos, medição nativa Windows de CPU por delta GetSystemTimes, RAM física total/disponível/em uso por GlobalMemoryStatusEx, capacidade do volume Windows por GetDiskFreeSpaceExW, relógio local, medidores semicirculares, PT/EN/ES, sem rede, sem varredura obrigatória, amostra a cada 3 s apenas na página visível, cache nativo de 2 s e disco 15 s. Dados indisponíveis não viram zeros. CPU inicial aguarda segunda amostra. Disco é capacidade, não atividade I/O; RAM disponível não é análise de conteúdo de memória. Mantém motor, controles e versões anteriores, sem instalador. Gate adicional: controles únicos, sintaxe JS, testes Go nativos e marcador de renderização CPU/RAM real no smoke Windows. CI e validação visual/física pendentes. R1 não atende a este requisito; não oferecer como R2. R2 é revisão LAB da base 3.0.8, não release de produção nem atualização automática. Demais requisitos cumulativos continuam pendentes até evidência específica (processos/memória, manutenção, estatísticas globais, proteção residente, validação integrada etc.).


> **Estado vigente — 24/11/2026:** Windows **3.0.7 LAB**: CI aprovado ([execução 45507766010](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45507766010)), teste rápido físico parcial e exportação confirmados; eficácia e estabilidade integral pendentes. **3.0.8 LAB**: primeira intervenção funcional em [ramificação isolada](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/windows-3.0.8-lab), otimização de entropia, leitura em blocos e tratamento de cancelamento; CI aprovado; pacote portátil exclusivamente para teste rápido controlado, validação física e liberação definitiva pendentes. Macintosh permanece independente, candidato 3.0.5 LAB. Não substitui o Defender. [Memorando](MEMORANDO_DESENVOLVIMENTO.md) · [metadados](ABOUT_GITHUB.md). Nenhuma release definitiva.

## Histórico do candidato 3.0.8 R1

24/11/2026 — CI 3.0.8 LAB R1 aprovado na [execução 45513126955](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45513126955), commit bfdc766b80b75a13c8d2198fa7356d96df4ec8ee. ZIP portátil x64 exclusivamente para teste rápido autorizado: [artefato 14976792219](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45513126955/artifacts/14976792219), 6.883.070 bytes; SHA-256 do ZIP conferido nos bytes baixados: 9c5f62efc583f6e8fe18691c008f8bf495cd9b94fcacbf58431e74e94ae6eb5e; CRC íntegro. SHA-256 do EXE confere com o manifesto: 98d07e58e2780d3d1ee0211f9d65869aa09928693e0dd5272c9f5f7f93110cbe. CI aprovou regressões de leitura/PE (30 rodadas cada), 50 regressões de varredura, 10 repetições, 5 rodadas de links, EICAR, cenários sintéticos, build e smoke desktop. Não equivale a proteção real certificada. Testes físicos, desempenho comparativo, cinco ciclos integrados e demais gates pendentes; nenhum instalador, atualização assinada ou release autorizada. Manter Defender e SmartScreen ativos, extrair ZIP em pasta separada e executar apenas rápida em ambiente autorizado.

## Histórico e documentação técnica (as versões abaixo são registros anteriores)


> **24/11/2026 — Windows 3.0.7 LAB R1, em desenvolvimento (não liberado):** correção experimental do parser PE baseada nos relatórios 3.0.6; 30 rodadas locais de testes do componente aprovadas. CI integral, estabilidade em Windows real, desempenho, cobertura e instalação/migração pendentes. [Ramificação candidata](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/windows-3.0.7-lab) · [Memorando ZVX-EVO-019](MEMORANDO_DESENVOLVIMENTO.md). Não há instalador 3.0.7 para download. Links anteriores são históricos de laboratório.


Scanner complementar de segurança para Windows x64, desenvolvido em Go, com interface desktop, análise de arquivos e links, quarentena reversível e IA consultiva experimental. **Created & Engineered by Leonardo A Botelho.**

[![Windows LAB R5](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-306.yml/badge.svg?branch=windows-3.0.6-lab)](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-306.yml)

**Estado:** candidato 3.0.6 LAB R1, CI e validação física pendentes. Não substitui Microsoft Defender.

**Revisão LAB pós-varredura e Global Threat Intelligence:** CI Windows aprovado no commit `07373e9c538794c5f14a3bc13ad85206d3049bc9`; alterações de interface, classificação de erros, retry limitado, procedência do build e consulta consultiva DROP incluídas. Validação visual e em Windows real ainda necessária.

## Download e instalação de teste

**3.0.6:** candidato não aprovado ainda. Os links abaixo são históricos da 3.0.5; não utilizar como pacote da 3.0.6. Nova instalação deverá usar `installer/ZEUVORYX-LAB-Setup-3.0.6.exe` após aprovação.

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


## 3.0.8 LAB R4 — correção de interpretação: Segurança em camadas (24/11/2026)
Titular esclareceu inequivocamente que não pediu câmera: a solicitação é na tela Segurança em camadas, com Consultar proteção, Abrir Segurança do Windows, Verificar Windows Update. A ramificação equivocada windows-3.0.8-camera-lab NÃO deve ser integrada, distribuída nem usada como base; seu requisito de câmera foi invalidado pelo titular. R4 correta parte da revisão integrada R3 e mantém todos seus componentes. Cada linha do diagnóstico existente recebe ponto colorido e texto original: verde ativado/bloqueio, vermelho desativado/não verificado/consulta indisponível, amarelo auditoria/aviso/modo intermediário, azul informação técnica. BitLocker On/Off normalizado para Ativado/Desativado. RDP desativado pode ser configuração desejável; cor descreve estado, não risco ou vulnerabilidade. Consulta somente leitura, sem alterar políticas, permissões, Defender ou Windows. Teste Node de classificação (13 casos x 5) e controles preservados, gate de identidade sem painel de câmera. CI e validação visual/física pendentes. Nenhum instalador ou lançamento. Manter memorando integral e demais pendências da R3.

## ZEUVORYX 3.0.8 LAB R5 — IA ativa e mascote funcional (24/11/2026)
Diagnóstico de fonte: a IA anterior era chamada exclusivamente após alerta heurístico com vetor não nulo; uma varredura sem alertas podia executar zero inferências. R5 candidata desloca a chamada real de adaptive.Observe para antes do limiar de alerta, avaliando o vetor de características já calculadas em cada chamada concluída de analyze, inclusive arquivos sem alerta. Não relê arquivos, não altera pontuação de risco, veredito, quarentena ou política de proteção. Contadores medem avaliações de modelo, não arquivos únicos (ZIP e releituras podem gerar chamadas adicionais). Autoteste sintético único na abertura, separado das inferências de arquivo. Mascote SVG/CSS leve na visão geral e Motor e IA: em pé levantando pesos apenas quando o backend indica avaliação em andamento ou aumento real no contador entre amostras; deitado descansando em espera, falha identificada separadamente; texto, contador, última avaliação, versão, sucesso, falhas e autoteste. Amostragem de 2 s nas páginas visíveis; respeita movimento reduzido. Nenhuma animação representa proteção comprovada, aprendizado contínuo ou treino fictício. Aprendizado persistente e proteção residente continuam pendentes (REQ-029/030/031), assim como gates de eficácia e desempenho. Teste integrado exige três avaliações reais em três arquivos benignos sem produzir alerta; regressões anteriores mantidas. CI/visual/físico pendentes. R5 parte da R4 correta de segurança em camadas, sem recurso de câmera. Nenhuma distribuição, contratação ou instalador. Não entregar como aprovado até verificar resultado e pacote.
