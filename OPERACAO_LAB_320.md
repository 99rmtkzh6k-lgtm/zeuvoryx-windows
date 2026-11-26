# Continuidade operacional — ZEUVORYX 3.0.20 LAB

## Adendo pétreo de 26/11/2026

O Codex executa as correções, integrações, testes e compilações LAB a partir da última base com evidência. A operação inclui verificar e instalar o candidato autorizado no Windows físico, conservar o pacote anterior para reversão, exercitar os motores e automatizar medições e validações autorizadas. A auditoria independente permanece separada. Preservam-se a identidade Olympus, o Defender, os 32 artigos, 43 requisitos e 19 fases. Nenhum relatório privado segue para o GitHub.

## Acesso e linha de base

O registro físico mais recente identifica **3.0.18 LAB, build 20, commit 99050aa**. A instalação foi identificada no Windows do titular em verificação anterior. O estado corrente deste turno é **Linux em nuvem**, sem terminal, volume ou canal de automação do Windows físico. O artefato 3.0.19 permanece preservado para piloto ou reversão. Instalação, lançamento e configuração persistente no computador físico dependem de uma sessão Codex conectada a ele.

- Caso de campo da 3.0.18: 351.860 arquivos próprios, 566 erros e ClamAV limitado após 60 min, sem resumo nativo. A auditoria não pode inferir se o scanner gastou esse tempo em CPU, leitura, espera ou outro trabalho. [Evidência de campo](AUDITORIA_CAMPO_319.md)
- Candidato 3.0.20: coleta CPU acumulada, transferências de I/O do processo, memória residente quando disponível e horário da última variação observada no ClamScan, por APIs do Windows, a cada cinco segundos. O painel e os relatórios usam as mesmas amostras. São contadores do processo, distintos de arquivos verificados e I/O físico do disco.
- O aviso após dois minutos sem mudança é consultivo. O prazo existente do ClamAV ainda encerra a etapa e mantém a cobertura parcial quando faltar o resumo nativo. Computação e I/O em zero são insuficientes para provar um travamento. Watchdog do motor próprio e mudança do limite ClamAV aguardam evidência física.
- O observador acompanha ZEUVORYX, ClamScan e FreshClam no diretório conhecido e o último relatório salvo. A tarefa de logon é opt-in, sem elevação, com uma instância. As amostras são locais, a cada 30 segundos, em JSONL com rotação de duas gerações em 4 MiB. A sessão precisa permanecer ativa. Apagar a tela não requer alterar políticas de segurança; a suspensão do computador interrompe a coleta até a retomada.

## Ativação e reversão no ambiente autorizado

O modo `-Once` grava uma única amostra. O modo `-Mode Register` cria e inicia a tarefa de logon para o usuário atual, a partir do script já instalado. O modo `-Mode Unregister` e o Agendador do Windows permitem remover a tarefa. O CI exercita `-Once` com o aplicativo instalado. Agendamento persistente, suspensão, retomada e rollback não foram validados no PC físico nesta execução.

A rotina oficial **Validação ZEUVORYX — relatórios e motores** foi reativada com periodicidade horária já existente. A rotina consulta o CI e as evidências autorizadas do GitHub. É distinta da auditoria independente e da coleta física local. Nenhuma tarefa envia logs brutos, altera o Defender ou inicia automaticamente uma varredura neste PC sem canal de execução Windows.

## Triagem dos grupos históricos

| Grupo | Estado e critério |
| --- | --- |
| AUD-CI-01 | Telemetria de ClamScan medida à parte neste candidato. A branch experimental reprovada pode servir de referência seletiva, com ZIP renomeado, gráficos e bloqueio de I/O ainda pendentes de paridade. |
| AUD-CI-02 | Recuperação de integração 3.0.19 aprovada no CI; remediação, Arcade instalado e ensaios físicos seguem abertos. |
| AUD-CI-03 | AMSI sem confirmação de provedor operacional no Windows físico. Preservar conclusão inconclusiva em erro ou indisponibilidade. |
| AUD-CI-04 | Validade de coleta DROP em campo permanece pendente, separada dos testes de cache/contrato. |
| AUD-CI-05 | Macintosh é trilha independente. Nenhum reteste ARM64 foi disparado. |
| AUD-CI-06 | Infraestrutura e workflows históricos permanecem preservados. Ensaio pareado de desempenho em corpus fixo ainda pendente. |
| AUD-CI-07 | Atualizador assinado, restauração de versão e plataformas dependem de testes próprios. Jobs cancelados permanecem históricos. |

Fontes técnicas: [GetProcessTimes](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes), [GetProcessIoCounters](https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-getprocessiocounters), [GetProcessMemoryInfo](https://learn.microsoft.com/windows/win32/api/psapi/nf-psapi-getprocessmemoryinfo), [ClamAV Scanning](https://docs.clamav.net/manual/Usage/Scanning.html).
