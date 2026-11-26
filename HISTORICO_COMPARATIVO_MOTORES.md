# Histórico comparativo dos motores — ZEUVORYX LAB
Atualização inicial: 26/11/2026. Preservar todos os registros anteriores e acrescentar novas observações com identificação única da sessão, commit, build, ambiente, configuração, versão do motor e da base. Fonte original: [auditoria dos motores](AUDITORIA_MOTORES_2026-11-26.md), [auditoria de campo](AUDITORIA_CAMPO_2026-11-25.md), CI e relatórios locais. Este índice não substitui logs originais; não afirmar ter recuperado arquivos ausentes.

## Medições recuperadas
| Registro | Motor e versão do ZEUVORYX | Ambiente | Medição | Resultado / limites |
| --- | --- | --- | --- | --- |
| 25/11 18h26 BRT | Go, 3.0.15, 9d8b2c82 | Windows 10 Pro 19045, rápida | 207/216 arquivos, 96,147 s sessão | 9 temporários ocupados; parcial |
| mesma sessão | ClamAV, 3.0.15 | campo | 146 arquivos, 72,981 s fase; 132,20 MiB examinados, 1,47 GiB lidos | 0 infecções reportadas, 12 erros, 8 avisos de tamanho, parcial |
| 26/11, Issue #5 | Go, 3.0.15, b90fa60 | campo, rápida | 303 arquivos, 1.992.545.578 bytes, 119,235 s sessão | 4 ocorrências EICAR de mesmo hash no ZIP; 6 compartilhamento, 1 acesso |
| mesma sessão | ClamAV | campo | 250 arquivos, 3.628.228 assinaturas | ZIP EICAR reconhecido; 14 erros, 10 avisos tamanho; parcial |
| 26/11 | ClamAV, 3.0.16 | CI Windows, cinco ciclos instalados | mediana benigno 7,193 s; EICAR 7,454 s | 5/5 EICAR; controle benigno sem detecção |
| 26/11 | ClamAV, 3.0.17 | CI Windows, cinco ciclos instalados, VM diferente | mediana benigno 8,825 s; EICAR 8,830 s | 5/5 EICAR; controle benigno sem detecção |
| CI 45862702260 | módulos | Windows e Ubuntu | 5 ciclos | módulos aprovados; build integral falhou no encerramento do monitor |
| CI 45863724124 | módulos / instalação | Windows e Ubuntu | módulos aprovados; execução integral ainda em andamento na consulta inicial | não homologado até conclusão e validação física |

Não comparar velocidades de sessão com duração de motor; não somar alertas sobrepostos, ocorrências de teste ou avisos de limite como infecções. 5/5 EICAR não mede eficácia contra malware real. Simulações e regressões sintéticas não são corpus de malware. Não presumir medições de Defender, AMSI, YARA-X ou Sigma que não existem. 

## Protocolo de comparação e estratégia
- Pares comparáveis: mesmo corpus identificado por manifesto, mesmo Windows/hardware, escopo, estado da base, configuração, ordem de execução balanceada, condições fria/quente e pelo menos cinco ciclos. Comparar motor isolado e combinação, sem desligar a proteção nativa; separar a ação de cada motor. Preservar os logs e resumos de todas as execuções, incluindo falhas/cancelamentos; não armazenar executáveis maliciosos ou dados pessoais no repositório. 
- Por motor: tempo mediano/p95, CPU média/pico, RAM pico e ociosa, bytes físicos lidos, examinados, cobertura do corpus, não examinados e motivo, detecções verdadeiras por ID de caso, falsos positivos, falsos negativos, sucesso de quarentena/recuperação verificado, estabilidade e latência. Desconhecido = null; denominar cada métrica. Usar controles seguros e resultados externos independentes devidamente atribuídos; nunca tratar um controle EICAR como amostra de malware real.
- Calcular sensibilidade TP/(TP+FN), falsos positivos FP/(FP+TN), cobertura examinados/elegíveis, velocidade bytes examinados/tempo ativo, apenas quando o denominador conhecido e a metodologia equivalente. Comparar detecção incremental exclusiva (IDs confirmados adicionados à combinação), custo marginal de CPU/RAM/tempo e cobertura perdida. Evitar pontuação arbitrária ou otimização apenas por velocidade. 
- Primeiro exigir segurança, cobertura, estabilidade, falsos positivos e remediação conforme critérios acordados; depois escolher a fronteira de melhores combinações entre eficácia e recursos. Testar roteamento seletivo e execução serial/paralela sob diferentes disponibilidades de RAM; não pular tipos de arquivo sem demonstrar cobertura equivalente. Não retirar capacidade aprovada sem substituição equivalente/superior. Sem vencedor até teste pareado. Preservar protocolo do memorando, nenhum gasto autorizado, nenhuma credencial indevida, nenhuma regressão.
