## 3.0.19 LAB — integração dos 32 artigos e validação do candidato

O registro histórico mais recente da main mantém a triagem das falhas antigas e a regra de execução única pelo Codex. A [matriz](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/RASTREABILIDADE_32_ARTIGOS_2026-11-26.md) preserva os 32 artigos e os vínculos com 19 fases, 43 requisitos e 23 artigos históricos. A [auditoria de campo](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/AUDITORIA_CAMPO_319.md) separa as sessões rápida e completa da 3.0.18, com 566 erros do motor próprio e prazo ClamAV de 60 minutos na completa. Os relatórios brutos permanecem privados.

A 3.0.19 registra presença do resumo e motivo de término do ClamAV. O hash calculado durante a cópia reforça a quarentena e a restauração, preservando a checagem esperada para achados confirmados. O Arcade anterior saiu do candidato e foi substituído pelo 2048 upstream MIT, offline e em iframe isolado. Cinco ciclos locais verificaram processamento Go, prazo/cancelamento do subprocesso controlado e abertura, jogada, reinício, fechamento e recusa pelo controle parental em Chromium. O runner validou o build Windows e os cenários do CI, incluindo migração e cinco ciclos com ClamAV embarcado. Os dois primeiros CIs pararam antes de compilar por incompatibilidade do gerador da quarentena e conversão de LF no checkout Windows. A correção preservou os gates em [45883312622](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45883312622).

[Instalador Windows x64 para piloto](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45883312622/artifacts/15148388359), build **29**, tentativa **1**, compilado em `2026-11-26T19:31:00.2962993Z` UTC, commit `5efb53b9311e82e5ec34e351afb0870eb5ee7769`. ZIP, CRC, manifesto e hashes verificados. SHA-256 ZIP `816c01db6a0427817db0e50ae6abbddc1def1bf494d2247812baf10d9d7aa5a0`. SHA-256 EXE `3cea7ebb33df7c86f6954fbae27a1d2f2659bf5bb604d8b564f74dfc1b53fdd0`. Pacote LAB sem Authenticode.

Nos fixtures instalados, as medianas nativas do ClamAV na 3.0.18 e na 3.0.19 foram 6,743 → 9,597 s (benigno) e 6,701 → 9,518 s (EICAR), cinco ciclos por caso em runners distintos. Os resultados de detecção foram iguais. Sem ensaio físico controlado, a diferença não permite atribuir causa ou avaliar disco completo.

A validação física da nova completa, a desinstalação de ponta a ponta, a cobertura do ClamAV, a proteção própria residente/web, a neutralização real, a verificação do Arcade no WebView2 instalado e a ampliação do catálogo seguem abertos. About, Description e Topics continuam pendentes por HTTP 403. O código funcional está na branch LAB.

# Auditoria de campo e integração 3.0.19 LAB — 26/11/2026

## Sessões físicas da 3.0.18

Dois relatórios automáticos finais do titular com commit `99050aa`, horários de Brasília e estados finais distintos por motor. Os JSON brutos permanecem privados.

| Campo | Rápida, 13h33–13h35 | Completa, 13h36–15h04 |
| --- | ---: | ---: |
| Duração | 139,821 s | 5.260,213 s |
| Arquivos próprios | 379 | 351.860 |
| Bytes contabilizados pelo motor próprio | 2.389.389.529 | 65.658.238.581 |
| Itens ignorados | 6 | 679 |
| Erros próprios | 6 | 566 |
| Arquivos inventariados | 384 | 352.325 |
| Falhas de inventário | 1 | 101 |
| Arquivos ClamAV | 324 | Sem resumo final |
| Erros ClamAV | 15 | Desconhecido |
| Avisos de limite do ClamAV | 12 | 5 observados antes do prazo |
| Controle EICAR próprio | 4 ocorrências do mesmo conteúdo | 5 ocorrências do mesmo conteúdo |
| Heurística própria | 1 alerta consultivo | 99 alertas em 63 hashes distintos |

Ambas as sessões encerraram sem cancelamento manual e com cobertura parcial. A etapa ClamAV da completa iniciou cerca de 27m40s após o começo da sessão e foi interrompida após o limite de 60 minutos. A saída não contém o resumo nativo. O número 1 do código de saída, isoladamente, não permite inferir uma detecção nesse cenário. A rápida teve um EICAR independente do ClamAV. Os dois arquivos não permitem medir CPU/RAM/I/O por motor ou calcular cobertura exata a partir dos inventários parciais. As sessões históricas de inventário 310/302/9 e painel com 561 erros são apenas referências relatadas, separadas das contagens de 26/11.

Na completa, os 566 erros próprios foram 419 de acesso negado, 121 de compartilhamento, 20 de bloqueio e 6 outros. O Defender estava ativo em modo normal e os três perfis do firewall estavam ativos nas duas sessões. Nenhum relatório comprova a modificação de políticas de proteção ou a eficácia residente.

## Decisões e fontes técnicas

- **ClamAV.** O ClamScan produz um resumo quando termina a operação. ClamD evita carregar a base por lote, mas adiciona memória residente, serviço e superfície de configuração. A completa observada mostra o efeito do prazo; sem medições equivalentes de hardware, volume e memória, alterar arquitetura ou ampliar prazo não asseguraria melhor cobertura. O JSON passa a explicitar o motivo do término e a presença do resumo. [ClamAV — Scanning](https://docs.clamav.net/manual/Usage/Scanning.html) · [Go os/exec](https://pkg.go.dev/os/exec)
- **Quarentena.** O build anterior registrava digest ao partir de achado verificado. O caminho genérico podia produzir um item sem hash, e a cópia de restauração não verificava o fluxo gravado. O candidato calcula SHA-256 durante a cópia, revalida identidade, tamanho e data de modificação da origem, e recusa restauração quando tamanho ou digest divergem. Itens legados sem digest ainda podem ser restaurados se corresponderem ao tamanho registrado. A verificação de identidade reduz corridas comuns; teste físico de concorrência hostil, ACL e ação atribuída ao motor ainda é necessário. [Go crypto/sha256](https://pkg.go.dev/crypto/sha256) · [Go os.SameFile](https://pkg.go.dev/os#SameFile) · [Microsoft Defender — quarentena](https://learn.microsoft.com/defender-endpoint/restore-quarantined-files-microsoft-defender-antivirus)
- **Arcade.** Removidos jogo da velha, memória, maior ou menor, paciência, dominó, xadrez e Blocos do Olimpo do catálogo e do build. Selecionado 2048 oficial de Gabriele Cirulli em commit fixado, MIT. Simon Tatham usa JS e WebAssembly separados para Mines, exigindo avaliação adicional de carregamento offline e WebView2. Não incorporado neste ciclo. O iframe do 2048 usa origem opaca, CSP sem rede e é removido ao fechar. Revisão do wrapper WebView2 encontrou assinatura de evento de mensagens no documento principal, sem registro do evento de mensagens de frames. O teste Chromium confirmou isolamento em relação ao DOM pai; validação da ponte no WebView2 em uso permanece prevista no piloto. [2048](lab/arcade2048/README.md) · [Simon Tatham](https://www.chiark.greenend.org.uk/~sgtatham/puzzles/) · [Documentação de mensagens WebView2](https://learn.microsoft.com/dotnet/api/microsoft.web.webview2.core.corewebview2.webmessagereceived)
- **Comparação internacional.** As metodologias de [AV-Comparatives](https://www.av-comparatives.org/consumer/test-methods/) e [AMTSO](https://www.amtso.org/standards/) fornecem critérios para desenho de ensaio e prestação de contas. O teste com EICAR, Chromium e fixtures benignas mede fluxos do produto. Comparação contra produtos concorrentes exige as mesmas amostras, contexto, hardware e critérios. Nenhuma superioridade ou nota independente é reivindicada.

## Regra de incorporação

A linha validada 3.0.18 permanece base deste candidato. A branch paralela de cobertura experimental tinha CI reprovado. Seus commits de telemetria, cobertura e ponte de diagnóstico não foram promovidos. A ponte operacional depende de autorização separada conforme artigo 31. A matriz mantém requisitos e origens; o estado apresentado é LAB experimental.
