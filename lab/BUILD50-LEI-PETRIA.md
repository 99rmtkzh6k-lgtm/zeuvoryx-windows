# ZEUVORYX 3.0.23 LAB Build 50 — trabalho em preparação
Base imutável: f0b136c4b2f0bd285b87f228a2b48fb4db304e73. Run de referência 46011202194, 3 passagens, 9/9 jobs.
Branch lab-build50. Não alterar main ou Releases. Não declarar aprovado antes de validação física.
Modo LAB: apenas informar; Defender ativo; sem bloqueio preventivo, exclusão ou quarentena automática. Não é antivírus homologado; testes sintéticos não provam eficácia contra malware real.
Progresso deste commit: correção inicial do layout iframe, teste de dimensões 1366x768 com CSS gerado, mensagens de acesso por código Windows e testes unitários associados. Ainda não executados.
Pendências obrigatórias antes de compilar: controles reais Win32 SC_MINIMIZE/SC_MAXIMIZE/SC_CLOSE (46x32), dimensões normal/maximizadas e DPI na janela nativa, confirmação de varredura concorrente sem simular pausa, contagens reais de falhas por categoria sem somá-las a alertas, ZIP/EICAR/cobertura parcial, identidade e atualização LAB Build 50, trava de download, revisão de vídeos, integração dos três ciclos no workflow e nomes de artefato exclusivos por SHA. Zeus e Cris estacionados.
Validação: seis jogos com interação, sem cortes; Windows e Linux 8/8 por passagem, 5 ciclos de instalação, scanner, QA20, 50M simulados. Windows 52 etapas aprovadas + 1 ignorada conforme modo autorizado; 9/9 jobs no mesmo SHA, parar na primeira falha. Baixar apenas passagem 3, verificar SHA256 do EXE e testar/filmar no PC do titular, mantendo Defender ligado. Não usar resultados antigos como se fossem Build 50.
Observação: código 5 não prova que o caminho é pasta de sistema. Contagens 62/13 são exemplos do titular, não constantes de execução. "Não é evidência de vírus" não significa que o arquivo foi examinado ou está limpo.

Segundo commit: aviso de varredura concorrente com confirmação para continuar sem pausar (não existe pausa nativa verificada), regressão de recusa/aceite, rótulos dos contadores de erros por código e tradução visual de mensagens conhecidas do Windows. Dados brutos do relatório e classificação de alertas preservados. Sem execução de testes nesta sessão. Pausa nativa continua pendente caso seja exigida.

## Cobrança complementar do titular — 27/11/2026
Ordem obrigatória, sem compilação antecipada:
1. Controles nativos de janela para jogos, ações reais minimizar/maximizar/fechar, 46x32 px; SC_MINIMIZE/SC_MAXIMIZE/SC_CLOSE devem ser integrados ao host Win32/WebView2, não simulados por CSS. Testar no host real e resolução 1366x768.
2. Identidade 3.0.23 LAB Build 50; explicar que builds LAB ficam em Actions, não Releases; ausência de Release não prova ausência de build LAB.
3. Bloquear downloads concorrentes por identificador de artefato/ZIP, avisar "Já está baixando". Identificar se o download ocorre no Safari/GitHub ou no app: o app não controla downloads iniciados pelo navegador.
4. EICAR como teste não real, ZIP grande como não inspecionado por dentro, cobertura parcial explícita, números reais de Windows 5/32/33 sem misturar com alertas; ignorado não significa limpo.
5. Defender ligado em todas as etapas; validar fisicamente antivírus ativo, proteção em tempo real, comportamental, modo normal e antiviolação. CI não atesta o PC físico.
6. Histórico em português, alertas, falhas de permissão, ignorados e parcial separados; ErrorCount agregado não equivale a falhas de permissão.
7. Preservar imagem de Zeus existente. Zeus estacionado: mitologia grega, barba branca, toga, dourado, raio, Olimpo e templo, arte 2D.
8. Cris estacionada: Vamos Brincar com a Cris / Vamos Pintar com a Cris; jogo de maquiagem infantil original, batom, sombra e blush; não copiar nem usar marca Barbie.
9. Revisar vídeos disponíveis antes de concluir; regressão Windows 8/8, Linux 8/8, cinco instalações, scanner, QA20, 50M simulados, seis jogos; progressão visual e funcional; Windows 52 etapas + 1 ignorada conforme modo; 3 passagens/9 jobs; SHA256 EXE e validação física filmada. Parar na falha, não publicar automaticamente.

Neste commit: histórico passa a apresentar scan_started/scan_finished em português; ErrorCount segue agregado e não é rotulado como permissão exclusiva. Contagens específicas ainda pendentes. Sem CI ou compilação.
