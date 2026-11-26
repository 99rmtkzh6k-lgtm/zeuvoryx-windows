## 3.0.17 LAB — integração complementar verificada (26/11/2026)

O [CI 45850320492](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45850320492) do commit `ec0318ae64c7ccfb96890f4d71f18ca3dac11c2b` aprovou 48 etapas do job Windows, sete das regressões Windows e sete das regressões Ubuntu. Migração, instância única, abertura WebView2, cinco ciclos instalado com ClamAV e o diagnóstico local passaram. As regressões de ZIP com quatro ocorrências de teste passaram também em cinco rodadas locais. A Issue #5 registra o fluxo rápido EICAR da 3.0.15 com cobertura parcial; consulte a [auditoria de motores e limites](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/AUDITORIA_MOTORES_2026-11-26.md).

[Instalador Windows x64 para piloto](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45850320492/artifacts/15133868370), build **19**, tentativa **1**, compilado em `2026-11-26T15:41:43.0127069Z` (UTC). Baixado e verificado por CRC, manifesto e ambos os hashes. ZIP com 192819495 bytes e SHA-256 `08bec229301878889da0e1843ae311cba61b078b6790ce36a9ec24c3c855f87b`. Instalador com 193292948 bytes e SHA-256 `8544d25403776df6526b477f9d073482c1edcef66318acaec7907494dc927dbf`. Authenticode ausente. Artefato expira em 03/12/2026 às 12h46 de Brasília.

Medianas nativas ClamAV em cinco ciclos unitários do aplicativo instalado: 3.0.16 → 3.0.17, benigno 7,193 → 8,825 s e EICAR 7,454 → 8,830 s. Nos dois runners o controle benigno não gerou detecção e os cinco controles EICAR foram reconhecidos. O tempo maior na 3.0.17 foi observado em VMs de CI distintas. O experimento não isola a causa nem comprova desempenho em disco completo.

Proteção residente comprovada, retomada, otimização de arquivos grandes, causa individual dos 14 erros e completa física continuam abertos. A classificação permanece LAB experimental, com fluxo positivo de controle EICAR confirmado. O GitHub retornou HTTP 403 na atualização administrativa de About e Topics, com a proposta registrada. O código funcional permanece na branch LAB.


# Auditoria complementar de motores — ZEUVORYX 3.0.17 LAB

## Evidência de campo

A [Issue #5](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/issues/5) publica a síntese anonimizada do relatório de 26/11/2026. O arquivo original `zeuvoryx-auto-final-3610981578.json` não estava acessível nesta revisão. A rápida 3.0.15, commit `b90fa60`, encerrou em 119,235 s, com 303 arquivos e 1.992.545.578 bytes no motor próprio, seis falhas de compartilhamento e uma de acesso. O motor próprio localizou quatro ocorrências EICAR de um mesmo SHA-256 em diretórios internos do mesmo ZIP. O ClamAV reconheceu o ZIP de forma independente. Reportou 3.628.228 assinaturas, 250 arquivos, 14 erros nativos e dez avisos de tamanho. O JSON foi exportado. O caso confirma o fluxo funcional de detecção de controle compactado, com cobertura parcial. O alerta consultivo sobre o YAML continua em triagem. A consulta Windows foi inconclusiva e o monitor próprio estava desabilitado.

## Configuração e arquitetura do ClamAV

- A release oficial estável conferida foi **1.5.5**. O SHA-256 do ZIP Windows x64 e o SHA-256 do código-fonte publicado coincidem com os dois digests fixados no CI. O pacote inclui o código-fonte correspondente e avisos do ClamAV. ClamAV é executado em processo separado. Validar obrigações de distribuição GPL-2.0, licenças transitivas, termos das bases e condições comerciais antes de distribuição geral. [Release upstream](https://github.com/Cisco-Talos/clamav/releases/tag/clamav-1.5.5)
- FreshClam usa a origem oficial, diretório compartilhado com o scanner, janela de atualização automática e intervalo de nova tentativa. A configuração explicita `TestDatabases yes`: o atualizador nativo testa a carga da base candidata antes da troca. Quando aplicável, uma base local válida permite continuar após falha de atualização, com indicação de condição e cobertura. Idade indicada no cabeçalho, integridade dos binários e validação da atualização são evidências distintas. Testes offline com base avariada, perda de rede, recuperação e reversão em Windows físico seguem pendentes. [FreshClam](https://docs.clamav.net/manual/Usage/SignatureManagement.html) · [Configuração 1.5.5](https://github.com/Cisco-Talos/clamav/blob/clamav-1.5.5/etc/freshclam.conf.sample)
- A varredura automática usa um processo ClamScan para cada lote de raízes. O scanner próprio e o motor de assinaturas fornecem resultados independentes, sem somar contagens. Em pouca memória, a política executa os motores pesados em sequência. A preparação de base, o processo de assinatura e o encerramento têm prazos. O ClamAV usa `--official-db-only=yes` na base integrada e registra limite via `--alert-exceeds-max=yes`. Links simbólicos ficam fora da travessia do ClamAV.
- Foram revisados `MaxFileSize`, `MaxScanSize`, `MaxRecursion`, `MaxFiles` e `MaxScanTime`. Os limites nativos continuam vigentes. `MaxFileSize` alcança o arquivo físico e componentes internos; `MaxScanSize` limita o volume acumulado da inspeção recursiva por arquivo de entrada; `MaxRecursion` conta contêineres aninhados, separadamente da profundidade dos diretórios dentro de um ZIP. Aumentar limites isoladamente pode elevar CPU, RAM, disco temporário e risco de exaustão. O scanner próprio preserva leitura incremental/hashing e os limites de entrada, contagem e volume ZIP. A expansão complementar de formatos depende de medir custo e equivalência de cobertura. [Opções nativas 1.5.5](https://github.com/Cisco-Talos/clamav/blob/clamav-1.5.5/common/optparser.c)
- **ClamScan** mantém isolamento e carregamento por lote. **ClamD** retém a base entre scans e oferece filas, mas implica memória residente, configuração de serviço e recurso extra durante recarga; o TCP do daemon não autentica clientes. **ClamDScan** é cliente do daemon e depende de sua política de acesso aos dados e limites. **libclamav** integrada no processo Go ampliaria o impacto de falha nativa, exigiria interface nativa e revisão jurídica de vinculação e distribuição. A decisão vigente é manter ClamScan no LAB. Reavaliar ClamD após medir base fria/quente, RAM em repouso e pico, paginação, versões de Windows, recuperação de serviço e cobertura equivalente de 4 GiB. [Arquitetura e protocolo](https://docs.clamav.net/manual/Usage/Scanning.html) · [Requisitos](https://docs.clamav.net/Introduction.html)
- O volume lido e o examinado vêm de contadores separados do processo nativo. Compactação, limites e a metodologia interna de contagem podem separar esses valores. A razão entre eles descreve essa execução do motor e não mede cobertura de disco. Qualquer análise causal dos 14 erros precisa dos registros originais e de execução controlada com os mesmos arquivos.

## Mudanças e diagnóstico

A 3.0.17 conserva o total nativo de erros e acrescenta contadores de **mensagens observadas** por arquivo ocupado, acesso negado ou outro motivo. Mensagens podem repetir o mesmo arquivo e não são denominador de cobertura. Motivos de limite são contados separadamente. Alertas ao vivo e relatório final compartilham a classificação. EICAR, registros de assinatura e avisos técnicos permanecem identificáveis. O relatório conserva o total bruto do motor para auditoria.

A consulta Windows permanece ativa enquanto os motores encerram. Quando ainda não há resposta inicial, a consolidação dá até dez segundos para a consulta, respeitando o cancelamento da sessão. O fechamento de pipes do subprocesso é limitado a cinco segundos. A consulta é somente leitura e preserva o estado desconhecido se não houver evidência. O manifesto ao lado do EXE vincula versão, número da execução de build, tentativa, UTC, commit e hash do executável.

## Demais componentes

| Componente | Evidência e fronteira operacional | Verificação restante |
| --- | --- | --- |
| Scanner próprio Go e heurística | Regressões de análise, cancelamento e limites. Controle EICAR em quatro caminhos internos idênticos no novo teste. Heurística do YAML segue consultiva. | Corpus legítimo com contexto, assinatura digital e correlação; completa, grande volume, retry posterior, dispositivos externos, agendamento, pausa e suspensão em Windows físico. |
| Defender, AMSI e Sysmon | Consulta Defender somente leitura e candidato com janela ajustada. AMSI tem regressões de contrato; Sysmon é opcional e o canal estava ausente na evidência física. | Teste físico de ativo, passivo, inacessível e indisponível. Confirmar alertas e bloqueio por evidência independente. [Get-MpComputerStatus](https://learn.microsoft.com/powershell/module/defender/get-mpcomputerstatus) · [AMSI](https://learn.microsoft.com/windows/win32/amsi/antimalware-scan-interface-portal) |
| Monitor residente, reputação e ransomware | Observação opt-in e regressões de correlação/link na base. A Issue #5 registra monitor inativo. | Confirmar monitor ativo em fundo, comportamento em downloads e bloqueio com mecanismo legítimo e suportado. |
| Motor adaptativo, quarentena e otimização | Validação sintética, regressões de política/isolamento e manutenção para escopos controlados. | Validação independente de falsos positivos, quarentena e restauração pós-reinício, confirmação de ação e diagnóstico de disco em máquina real. |
| Bibliotecas e serviços | Pacote embarca ClamAV e C2PA e usa o runtime Windows WebView2. A extensão gratuita permanece independente de assinatura comercial. | Inventário transitivo e SBOM; validade dos termos de redistribuição ClamAV/C2PA/WebView2, bases, feeds e serviços. Nenhuma nova dependência comercial incorporada. |

## Critérios de validação pendentes

Medir a mesma rápida e completa nas duas versões em hardware Windows equivalente, mesma base de assinaturas, corpus fixo, séries frias/quentes, resultados por motor, CPU/RAM/I/O e razões de limite. Repetir bloqueio temporário, permissão negada, ZIP grande, cancelamento, suspensão, rede indisponível e atualização falha. Registrar o estado de proteção residente e a reentrada em uma nova sessão. Separar saída de produto, regressões de integração e eficácia antimalware.

## Rastreabilidade dos 20 itens ClamAV

| Item | Estado nesta revisão |
| --- | --- |
| 1. Configuração | Revisada a invocação do ClamScan, limites padrão, FreshClam, caminhos, fontes e base oficial. |
| 2. Versão | 1.5.5 confirmada estável e compatível com o pacote Windows x64. |
| 3. Atualização automática | Fluxo integrado FreshClam com atualização por uso, política de renovação e cache. A atualização em repouso sem varredura permanece por validar. |
| 4. Integridade e autenticidade | Digests de distribuição conferidos com a release oficial. TestDatabases explícito. Teste de base adulterada e estratégia de rollback em campo pendentes. |
| 5. Tamanho | Avisos de arquivo grande têm categoria própria. Não ampliado o orçamento sem teste de recursos. |
| 6. Descompactação e profundidade | Limites e sua semântica revisados. Teste de quatro entradas em diretórios profundos; contêineres recursivos e carga extrema ainda precisam de corpus. |
| 7. Quatorze erros de campo | Originais ausentes. Novo registro diferencia categorias por linha e erro nativo; causas individuais ainda não confirmadas. |
| 8. Arquivos bloqueados | Retry curto e cancelável já disponível no motor próprio. Classificação das mensagens ClamAV adicionada; nova tentativa por arquivo e reagendamento permanecem abertos. |
| 9. Recursos | Execução em série para RAM restrita preservada. Medição de pico, paginação e I/O de campo pendente. |
| 10. Lido e examinado | Métricas nativas independentes. Causa da diferença do teste original depende das evidências completas. |
| 11. Serviço persistente | Avaliado em arquitetura; comparação de memória e latência em 4 GiB pendente. |
| 12. Alternativas de arquitetura | ClamScan mantido; trade-offs de ClamD, ClamDScan e libclamav registrados. |
| 13. Motor próprio | Lote integrado com identificação de resultados por motor e política de concorrência. |
| 14. Redundância | Um processo por lote de raízes; nenhuma mescla de contadores. Otimização de trabalho entre motores ainda por medir. |
| 15. Limites seguros | Limites preservados e resultado parcial em avisos técnicos. |
| 16. Classes de evidência | Registros de assinatura, EICAR, heurística, motivos de limite e mensagens de erro preservados. |
| 17. Recuperação | Saída limitada da sessão e fallback de base existente. Recuperação autônoma do processo dentro da mesma sessão segue pendente. |
| 18. Disponibilidade | Consulta em execução e diagnóstico por lote; monitor contínuo do ClamAV em repouso segue pendente. |
| 19. Licenças | Licença e fontes do ClamAV registradas; liberação comercial depende de revisão de termos e componentes transitivos. |
| 20. Desempenho comparativo | Tempos nativos da bancada CI distinguem o custo de um fixture da velocidade em disco completo. Série física pareada pendente. |
