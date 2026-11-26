# Matriz mestre ZEUVORYX — Windows 3.0.16 LAB

Revisão de 26/11/2026. Estado da implementação avaliado a partir do repositório e do CI. Os critérios abaixo são condições de aceite, e os resultados efetivos ficam no memorando. A branch funcional parte de `e5a080c`, com binário 3.0.15 da linhagem `b90fa60`. `main` em `a5711f7` fornece o histórico documental. As branches anteriores foram inventariadas e a ancestralidade da base foi conferida. Nenhuma mescla em bloco.

## Ordem de execução vigente

| Fase e origem | Requisito | Prioridade | Estado observado | Implementação necessária | Dependências | Teste de aceitação |
| --- | --- | --- | --- | --- | --- | --- |
| Fase 1 da ordem de 26/11 | Recuperação e rastreabilidade | P0 | Índice e histórico Git disponíveis, documentos externos pendentes | Relacionar REQ, artigos, revisão e diagnóstico; reconciliar originais quando disponíveis | Histórico, artefatos, acesso a anexos e teste físico | Reproduzir origem de cada requisito e relacionar cada relatório a versão, commit, sessão e estado. |
| Fase 2 da ordem de 26/11 | Linha de base física 3.0.15 | P0 | Rápida parcial documentada; completa ainda sem validação independente nesta auditoria | Reconciliar erros por código, cobertura, exclusões e duplicatas por sessão e conteúdo | Originais JSON e máquina Windows | Confirmar totais por motor; snapshots da mesma sessão contam como uma execução. |
| Fase 3 da ordem de 26/11 | Encerramento, telemetria e coordenação por motor | P0 | Limite por modo na base; 3.0.16 protege a consolidação após cancelamento | Completar watchdog de I/O do scanner próprio, filas, atividade medida e progresso dos motores | Windows, observabilidade por processo e corpus fixo | Exercitar sucesso, ociosidade, bloqueio, falha de comunicação e cancelamento; relatório parcial traz motivo e ignora publicação tardia. |
| Fase 4 da ordem de 26/11 | Escopos rápida, completa, personalizada, agendada; pausa, retomada e checkpoints | P0 | Base tem escopos, cancelamento, inventário e retry de compartilhamento; recuperação e cobertura incompletas | Persistir checkpoints atômicos, revalidar arquivos alterados, separar bytes elegíveis/lidos e espaço de disco, validar agendamento, pausa e ETA | NTFS, Windows 10/11, host 4 GiB, fixtures ZIP/arquivos grandes/bloqueados | Falhas explícitas, contadores por sessão, retomada valida arquivos alterados, estimativas compatíveis com escopo. |
| Fase 5 da ordem de 26/11 | Go, ClamAV/FreshClam, Defender, AMSI, YARA, Sysmon e proteção residente | P0 | Scanner próprio e assinatura integrada em CI; eventos/AMSI têm evidência parcial; YARA e bloqueio efetivo pendentes | Validar offline e atualização, cobertura por mecanismo, residente, downloads, antirransomware e fontes autorizadas | Bases oficiais, APIs Windows e licenças | Medir proteção por motor com testes benignos/controles oficiais, operação offline e Defender ativo; diferenciar evento de bloqueio. |
| Fase 6 da ordem de 26/11 | IA, diagnóstico copiável e FieldLearn | P1 | Inferência consultiva e diagnóstico por evidências no b90fa60 | Correlacionar hipóteses e tentativas, redigir dados sensíveis, validar regras e falsos positivos antes de aprendizagem | Corpus legítimo, logs estruturados e consentimento | Saída inclui versão, commit, eventos, histórico disponível e testes propostos, sem dados pessoais; treino fora dos vereditos vigentes. |
| Fase 7 da ordem de 26/11 | Quarentena, restauração, remoção e falsos positivos | P0 | Quarentena e revalidação já presentes; aceitação física ampliada pendente | Verificar persistência, reversão e revisão de falsos positivos Microsoft/Adobe/Office/OneDrive/PowerShell/DLL | Fixtures benignas, permissões, sandbox Windows | Isolar e restaurar por autorização, validar integridade após reinício; heurística isolada não autoriza exclusão. |
| Fase 8 da ordem de 26/11 | Web, phishing, downloads, reputação e Raio de Zeus | P0 | Analisadores locais e sinais de rede existentes; bloqueio web integrado pendente de evidência | Validar cobertura em HTTPS, DNS criptografado e navegadores; associar alerta e áudio à ação observada | Navegadores, políticas Windows, páginas oficiais seguras | Corroborar alerta e bloqueio separadamente; medir bypasses de cobertura sem modificar as proteções nativas. |
| Fase 9 da ordem de 26/11 | Arcade opcional com até 100 jogos gratuitos | P2 | Módulo offline existente com testes de regras; issue histórica ainda referencia 1.000 | Ampliar gradualmente jogos originais/licenciados, 2D, 3D leve quando útil, IA/oponente local, salvar e encerramento | Inventário de licenças, arte e testes Windows | Cinco ciclos por título implementado, verificar regras, fim de partida, offline, consumo e controles; zero apostas reais. |
| Fase 10 da ordem de 26/11 | Autenticidade de vídeo e controle parental | P1 | C2PA/filtros presentes na base em escopo parcial | Verificar cadeia de credenciais, metadados, estado inconclusivo, permissões e reversão de filtros | Vídeos de teste, C2PA, Windows e DNS | Credencial válida sem promessa de detecção de deepfake; regra demonstrada com reversão e testes da interface. |
| Fase 11 da ordem de 26/11 | Otimização nativa do Windows | P1 | Diagnóstico, manutenção selecionada e ações de disco já existentes; cobertura integral pendente | Validar seleção, TRIM/HDD, inicialização, compressão reversível e histórico | SSD/HDD de teste, permissões e backup | Demonstrar efeito e reversão, confirmação prévia e logs de sucesso/falha. |
| Fase 12 da ordem de 26/11 | Relatórios, interface, gráfico, histórico e alertas | P1 | Snapshots/exportação e medidores na base; realidade do gráfico e navegação por evento pendentes | Tempo real para taxa, CPU, RAM, filas e eventos; encaminhar alerta a evento, agrupar, salvar e compartilhar com consentimento | Windows/WebView2, esquema JSON e amostras medíveis | Todos os painéis entram no JSON; conferência visual dos medidores e da cópia, histórico após reinício. |
| Fase 13 da ordem de 26/11 | Identidade e autoria | P2 | Zeus/Olympus preservados nesta revisão; assinatura artística da nova ordem pendente | Aplicar L A Botelho manuscrita dourada com coroa e leão, sem bandeira, conforme determinação de execução vigente | Arte autorizada e instalador | Conferência visual do app e instalador, com crédito Leonardo A Botelho e logotipo preservados. |
| Fase 14 da ordem de 26/11 | Instalador, atualização e integridade | P0 | Instalador x64 3.0.15 passou no CI; canal oficial e validação física pendentes | Validar identidade, EULA, privacidade, rollback, migração, reparação, remoção, Authenticode, links e hash | Windows, WebView2, licenças, chaves do titular | Instalação limpa e migração, sem scan automático, assinatura em canal autorizado e SHA-256 independente. |
| Fase 15 da ordem de 26/11 | Núcleo gratuito e catálogo premium voluntário | P2 | Catálogo informativo de extensões; condições comerciais ainda não validadas | Verificar licença, redistribuição, avaliação gratuita, preço, renovação, cancelamento, privacidade e suporte fornecedor a fornecedor | Termos oficiais datados, autorização de licença | Núcleo opera sem conta paga; nenhuma ativação/cobrança sem consentimento. |
| Fase 16 da ordem de 26/11 | Desempenho, memória, I/O e estabilidade | P0 | Política de concorrência por RAM; comparação física equivalente pendente | Medir por processo e fase, processos filhos, repouso, ETA, carga de jogos e tamanho | Corpus fixo e host ~4 GiB | Repetir sob condições iguais; verificar cobertura preservada, sem processos órfãos e sem inferir ganho de um único piloto. |
| Fase 17 da ordem de 26/11 | Testes e revisão crítica | P0 | Regressões unitárias e pipeline Windows; cinco ciclos físicos integrados pendentes | Cobrir condições de borda de scan, web, engine, falso positivo, recuperação, instalação e atualização | CI e hardware Windows, controles oficiais e corpus benigno | Registrar comando, commit, build, ambiente, evidência e pelo menos cinco ciclos críticos viáveis. |
| Fase 18 da ordem de 26/11 | Integração GitHub e versão LAB | P0 | Branch dedicada e CI vinculado ao commit; main mantém histórico | Sincronizar documentação, versão, metadados e artefato LAB com o resultado observado | GitHub Actions e pacote final | Conferir branch/commit, jobs e SHA-256; manter dados de campo privados e sem promoção. |
| Fase 19 da ordem de 26/11 | Avaliação técnica e decisão de distribuição | P1 | Evidência desigual nos componentes; eficácia física incompleta | Avaliar maturidade em 20 dimensões quando houver métricas, com comparações controladas | Evidência física, segurança, licenças e artefato verificado | Vincular cada nota à evidência e manter avaliação pendente onde faltarem medições suficientes. |


## Requisitos anteriores

Os estados abaixo reproduzem o índice histórico (E = evidência parcial, P = pendente, B = bloqueado, A = adiado ou suspenso). Para cada ID, as fases vinculadas na tabela acima dão a dependência e o teste de aceite atual. Os registros [REQ/EVO](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/REQUISITOS_E_PENDENCIAS.md) e o [memorando central](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/MEMORANDO_DESENVOLVIMENTO.md) preservam o contexto e a redação de origem.

| ID | Requisito recuperado | Estado e ação histórica | Origem no índice | Fases de aceite |
| --- | --- | --- | --- | --- |
| REQ-001 | Proteção efetiva, estabilidade, velocidade, baixo consumo, simplicidade, privacidade | B — Transversal; sem regressão | CP-001 | 1, 3, 5, 16 |
| REQ-002 | Windows x64, instalação EXE simples, sem Python no cliente, início sem scan automático, atalho e operação simples | E — P0 abertura/instalação | Linha do tempo 2.1–3.0.6 | 14 |
| REQ-003 | Rápida, completa, personalizada, cancelar, progresso e exportação | E — P0 fechamento nas completas | EVO-013/015–018 | 3, 4 |
| REQ-004 | SHA-256 integral, heurísticas, assinaturas, análise PE/ZIP, limites explícitos, erros e cobertura parcial | B — P0 base histórica candidata de 6 hashes; eficácia real, integração e cobertura pendentes | EVO-034/035 | 4, 5 |
| REQ-005 | Quarentena reversível, restauração, confirmação de evidência, exclusões/falsos positivos e auditoria | E — P1 teste integrado, sem ação destrutiva automática | Cinco implementações futuras; 3.0.1–3.0.5 | 7 |
| REQ-006 | Remoção automática segura, interrupção de processos, persistência, reparação, bloqueio de execução | P — P1 módulos separados, consentimento/rollback; não confundir com Defender | Cinco implementações futuras | 7 |
| REQ-007 | Proteção residente, análise comportamental e antiransomware | P — P1 arquitetura, baixo consumo, compatibilidade Defender | Cinco implementações futuras | 5, 8 |
| REQ-008 | IA consultiva local, evolução controlada, validação independente, falsos positivos | E — P1 sem promover modelo reprovado; simulações não são ataques reais | Linha do tempo 3.0.2/3.0.3; STD | 6 |
| REQ-009 | Inteligência de ameaças internacional, fontes autorizadas gratuitas, cache, expiração, revogação, indicadores e atualizações | B — P0 base histórica mínima; atualização ampla, direitos e validação pendentes; não confundir KEV com malware | EVO-006/007/034/035 | 5, 15 |
| REQ-010 | Análise de links, phishing e proteção contra roubo de dados | E — P1 eficácia real pendente | 3.0.5 R5 | 8 |
| REQ-011 | Limpeza segura de temporários e métricas separadas de malware | E — P1 não remover arquivos ativos automaticamente | EVO-013 | 11 |
| REQ-012 | Histórico local autenticado, relatórios, salvar/exportar, privacidade e retenção | E — P1 preservação e relatório anonimizado pendentes | EVO-003/013; histórico local | 12 |
| REQ-013 | Controles de scan no topo do painel principal, gráfico e métricas junto aos controles, tela cheia e DPI | E — P1 validar UI efetivamente, sem perder handlers | Consolidação histórica; EVO-004 | 12 |
| REQ-014 | Olympus Silver, Zeus, ícones, arte local, assinatura manuscrita do criador, acabamento 3D leve | E — P1 preservar arte e autoria, verificar todos botões e idiomas | Identidade visual; EVO-011 | 13 |
| REQ-015 | Botão lateral Avaliação e resultados; cinco estrelas douradas fixas, prévia sem voto, escolha e confirmação única por instalação | P — P1 substituir seletor oculto, persistência/concorrência/migração | EVO-008/011/012 | 12 |
| REQ-016 | Painel local/mundial com contadores reais, média e distribuição, categorias, gráficos leves, sem números fictícios | P — P1 protótipo funcional; infraestrutura mundial depende de autorização | EVO-009–012 | 12 |
| REQ-017 | Serviço único entre GitHub, lojas, site, terminal, Windows/macOS; deduplicação, integridade, consentimento, TLS, criptografia, antifraude | P — P1 contrato LAB; produção bloqueada sem infraestrutura; métricas mínimas | EVO-009/010/012 | 12 |
| REQ-018 | Atualização in-place reconhecida por versão, preservação de dados, confirmação visual, link do pacote exato aprovado | E/B — P0 3.0.6 migração/abertura; catálogo exato pendente | EVO-012/018; CP-003 | 14 |
| REQ-019 | Atualização assinada por canal exclusivo do titular, chave oficial, manifesto, rollback e distribuição sem cadastro | E — P1 produção não habilitada; não enfraquecer assinatura | 3.0.5 R3/R4; EVO-012 | 14 |
| REQ-020 | PT-BR, inglês e espanhol, detecção Windows, seleção manual e fallback inglês; mais oito idiomas desejáveis | E — P2 11 idiomas sem sacrificar desempenho; validar os três atuais | Linha do tempo 3.0.2 | 12 |
| REQ-021 | Macintosh separado Intel/ARM64, paridade funcional progressiva, pacote e assinatura próprios | E — P1 não confundir compilação Intel com aprovação integral; consultar memo macOS | EVO-002; memorando Macintosh | 1 |
| REQ-022 | iOS e veículos: pesquisa futura, sem implantação agora | A — Retomar somente por decisão do titular | Sequência aprovada; prioridades | 1 |
| REQ-023 | Distribuição mundial inicialmente gratuita, sem publicidade; possibilidade de comercialização futura ou aquisição do projeto | P — Aviso prévio com data, preço e condições; aceitação expressa, sem cobrança retroativa; preservar direitos adquiridos | EVO-031/032/033; REQ-036/038/042/043 | 15, 19 |
| REQ-024 | Autoria, controle exclusivo de código e publicação; registro/licença/patente e revisão jurídica | E — P1 segurança da conta, assinatura, permissões; revisão jurídica pendente | Controle exclusivo; liberação | 13, 15 |
| REQ-025 | Testes repetidos, 5 ciclos integrados, 50 equivalência/desempenho quando viáveis, Windows 10/11 e máquinas reais, 4 GB/2 núcleos | B — P0; CI não basta | CP-001; EVO-018 | 16, 17 |
| REQ-026 | Versionamento uniforme, documentação, metadados, release/artefato e memorando atualizado sem lembrete | E — P0; About administrativo bloqueado | CP-002/003 | 18 |
| REQ-027 | Diagnóstico do fechamento: logs persistentes, saída, memória, WebView2, recuperação, correlação de eventos | B — P0; causa raiz ainda não comprovada | EVO-015–018 | 3, 6 |
| REQ-028 | Operação pelo iPhone, instruções completas copiar/colar quando necessário, GitHub e ZIP/EXE direto para compartilhamento | E — P1 reduzir trabalho manual; entrega apenas de pacote validado | Diretriz operacional; continuidade | 14, 18 |
| REQ-029 | Aprendizagem local contínua e automática após consentimento, perfil comportamental, candidato separado e validação/rollback | P — Próxima revisão; baixo consumo, sem rótulos autogerados nem envio externo | EVO-019; REQ-008 | 6 |
| REQ-030 | Autostart consentido no login, residente com ícone de bandeja, único processo, opção de sair/desativar | P — Próxima revisão; sem scan completo automático, sem ocultação | EVO-019; REQ-002/007 | 5, 14 |
| REQ-031 | Monitoramento de eventos de arquivos/downloads, escopos locais/nuvem acessível, fila limitada, análise de links autorizada | P — Próxima revisão; observação primeiro, sem hidratar nuvem ou interceptar TLS | EVO-019; REQ-007/010 | 5, 8 |
| REQ-032 | Corrigir Diagnóstico local 3.0.3 e todos rótulos obsoletos; versão única e estados reais | P — Próxima revisão; gate de UI em todos idiomas | EVO-019; REQ-026 | 6, 12 |
| REQ-033 | ClamAV/FreshClam integrado ao instalador, gratuito, independente de instalação externa; extensões de terceiros facultativas | E — CI anterior comprovou motor instalado; validar revisão acumulada, funcionamento offline e consumo físico. Extensões ainda informativas | EVO-019; política de extensões 25/11/2026 | 5, 14 |
| REQ-034 | Revisão integrada de toda fila após relatório 3.0.6; cinco ciclos e 50 equivalência quando viáveis | P — Próxima revisão, sem perder EVO-008–012 | EVO-019; REQ-025 | 17 |
| REQ-035 | Ícone ZEUS: preservar escudo; trocar apenas fundo/moldura escura por prata no software, alinhar recursos ICO/PE, taskbar, bandeja, atalhos e instalador | P — Próxima revisão; distinguir realce nativo Windows e validar DPI/temas | EVO-020; REQ-014/030 | 13, 14 |
| REQ-036 | Preparar distribuição mundial nas lojas e canais oficiais após validação integrada, assinatura, revisão de políticas/licenças e autorização | P — Pós-validação; não publicar nem contratar automaticamente | EVO-020; REQ-023/024 | 14, 19 |
| REQ-037 | Central pública de suporte internacional versionada, FAQ, ajuda no aplicativo e descoberta por buscadores | E — Base publicada em 4 idiomas; Pages, indexação e ajuda no app pendentes | EVO-021/022; REQ-023/036 | 12 |
| REQ-038 | Auditoria de propriedade intelectual, SBOM/licenças, termos de usuário, privacidade, responsabilidade e conformidade por mercado | P — Gate pré-lançamento; revisão jurídica e licenças comerciais ainda pendentes | EVO-024; REQ-023/024/036/037 | 14, 15 |
| REQ-039 | Ajuda inteligente offline no software: lupa, perguntas naturais, busca de respostas verificadas e atualizadas, sem IA externa paga | E — Próxima versão; 124 entradas/4 idiomas, 28 consultas verificadas; CI e integração Windows pendentes | EVO-025/026; REQ-037 | 12 |
| REQ-040 | Auditoria de utilidade de módulos/pacotes, custos, dependências e consumo, consolidar distribuição sem eliminar funções solicitadas | P — Próxima revisão; verificar funções reais, evitar componentes sem uso, preservar evidência | EVO-025; CP-001/REQ-001/026/038 | 15, 16 |
| REQ-041 | Continuidade da varredura completa com monitor apagado/bloqueado; solicitação temporária de energia, respeitar suspensão manual e bateria | E — Próxima versão; código candidato preparado; teste Windows físico e integração pendentes | EVO-027; REQ-003/025/027 | 3, 4 |
| REQ-042 | Independência financeira e licenças comerciais gratuitas: sem royalties, assinaturas obrigatórias ou contratação não autorizada; avisos/SBOM e alternativas verificadas | E — NSIS fixture e bloqueio de consulta não verificada preparados; migração, cobertura, CI e licenças integrais pendentes | EVO-029/030; REQ-009/038/040 | 15 |
| REQ-043 | Publicidade, patrocínios e espaço promocional | A — Suspenso por decisão do titular; sem anúncios no app/site de suporte; retomar somente por autorização expressa | EVO-030/031 | 15 |


## Artigos de unificação e precedência

| Artigo do comando de unificação | Fases de aceite |
| --- | --- |
| ARTIGO 1 — CONSULTA OBRIGATÓRIA | 1 |
| ARTIGO 2 — LEI PÉTREA DO VERSIONAMENTO | 18 |
| ARTIGO 3 — CORREÇÃO DEFINITIVA DOS GRÁFICOS | 12 |
| ARTIGO 4 — RELÓGIOS E MEDIDORES | 12 |
| ARTIGO 5 — ROBÔ E INTELIGÊNCIA ARTIFICIAL | 6 |
| ARTIGO 6 — DIAGNÓSTICO TÉCNICO INTELIGENTE | 6 |
| ARTIGO 7 — CENTRAL DE SEGURANÇA | 12 |
| ARTIGO 8 — EVOLUÇÃO DO DIAGNÓSTICO LOCAL | 6 |
| ARTIGO 9 — CLAMAV EM TEMPO REAL | 5 |
| ARTIGO 10 — FALHAS DO RELATÓRIO FÍSICO MAIS RECENTE | 2, 3, 4 |
| ARTIGO 11 — PROTEÇÃO NA INTERNET OBRIGATÓRIA | 8 |
| ARTIGO 12 — ALERTA RAIO DE ZEUS | 8, 12 |
| ARTIGO 13 — RECONSTRUÇÃO COMPLETA DO OLYMPUS ARCADE | 9 |
| ARTIGO 14 — GRÁFICOS DOS JOGOS | 9 |
| ARTIGO 15 — ASSINATURA DO CRIADOR | 13 |
| ARTIGO 16 — CERTIFICAÇÃO DIGITAL DO WINDOWS | 14 |
| ARTIGO 17 — DOWNLOAD DIRETO E INSTALAÇÃO | 14, 18 |
| ARTIGO 18 — MOTORES, BIBLIOTECAS E TECNOLOGIAS INTERNACIONAIS | 5, 15 |
| ARTIGO 19 — MODELO GRATUITO E PREMIUM | 15 |
| ARTIGO 20 — TODAS AS PENDÊNCIAS ANTERIORES | 1, 4, 5, 7, 11, 14 |
| ARTIGO 21 — CRITÉRIOS DE APROVAÇÃO | 17, 19 |
| ARTIGO 22 — COMANDO DE EXECUÇÃO UNIFICADA | 18 |
| ARTIGO 23 — LEI PÉTREA SUPREMA | 19 |

A ordem mais recente define **até 100 jogos gratuitos** e assinatura artística **L A Botelho**, sem bandeira. A issue de 1.000 jogos fica como registro histórico. iOS e veículos seguem adiados, macOS segue em trilha independente e publicidade permanece suspensa.

## Fontes e limites

- [Protocolo GitHub](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/PROTOCOLO_INTEGRACAO_GITHUB.md), [Auditoria de campo de 25/11 e adendos](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/AUDITORIA_CAMPO_2026-11-25.md), [Comando de Unificação](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/COMANDO_UNIFICACAO_2026-11-26.md), [Licenças](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/AUDITORIA_LICENCAS.md), [Memorando Macintosh](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/0fbf709/MEMORANDO_MACINTOSH.md) e [CI da base](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45755065493).
- Os anexos JSON originais de 26/11, o arquivo com o nome exato do memorando mestre de 25/11, as mídias originais, o executável instalado pelo titular e conversas externas à sessão ficaram indisponíveis. A auditoria de campo fornece resumos. A identidade de sessões e a cobertura da completa ficam em aberto.
- A versão 3.0.15 comprovou conclusão parcial em um piloto rápido documentado. O limite de 60 minutos pode interromper o ClamAV ativo durante a completa e independe da falha de espera corrigida na 3.0.16. Medição controlada por fase precede qualquer escolha de ClamD, alteração de orçamento ou atribuição de causa à lentidão de campo.

### Diagnóstico técnico do P0

A leitura do código confirmou que, na 3.0.15, o contexto tinha prazo próprio para o ClamAV e o consolidador ainda fazia uma leitura sem prazo do canal de resultado. A etapa deveria responder ao cancelamento, mas o chamador não tinha saída independente se uma operação auxiliar ou de I/O não devolvesse o controle. O mesmo código espera inventário por até cinco segundos e diagnósticos auxiliares por até dez segundos cada. Esse defeito é demonstrável no fluxo de controle. O resumo de campo, sozinho, não identifica se havia I/O bloqueado, processamento ativo ou espera por outro componente no momento observado.

Na 3.0.16, resultado normal e cancelamento cooperativo seguem o mesmo fluxo de consolidação. Quando o contexto vence ou o usuário cancela, a espera adicional pelo resultado tem limite de dez segundos. Ausência de confirmação é registrada separadamente da conclusão do processo. O limite geral de cada modo permanece 5/60/20 minutos para rápida/completa/personalizada. O prazo da etapa não serve para medir inatividade. Publicações tardias são isoladas e outra etapa automática de assinatura não concorre com a anterior ainda em encerramento.

A documentação do ClamAV explica que ClamScan carrega a base a cada execução e que o daemon ClamD mantém configuração e base persistentes. A memória mínima recomendada com base padrão no Windows é 3 GiB, e o transporte TCP do daemon não autentica clientes. Em máquinas de aproximadamente 4 GiB, a decisão por persistência ou concorrência depende de medição de base fria/quente, pico de RAM, paginação, latência e equivalência de cobertura. [Varredura e ClamD](https://github.com/Cisco-Talos/clamav-documentation/blob/main/src/manual/Usage/Scanning.md) · [Requisitos](https://github.com/Cisco-Talos/clamav-documentation/blob/main/src/Introduction.md) · [Cancelamento e WaitDelay em Go](https://pkg.go.dev/os/exec#Cmd)

### Gates de campo

- Usar o mesmo corpus e escopo por modalidade e revisão. Registrar versão, commit, sessão, máquina, total da RAM, memória disponível, estado da base, fase, latência, bytes, contagens e cobertura por motor. Manter a segurança nativa ativa.
- Concluir rápida e completa no Windows 10/11 e conferir presença e estado do JSON final. Repetir com cancelamento, arquivo ocupado, erro de acesso, ZIP dentro e fora dos limites e reinício. Conferir que próxima sessão começa com contadores novos e que um snapshot parcial antigo continua distinguível do final.
- Avaliar residente, web, antirransomware e falso positivo em controles separados. A comprovação de instalação, de consulta Defender, de um alerta sintético ou de uma detecção EICAR cobre apenas o respectivo teste.

## Avaliação comparativa de maturidade

A escala é de 0 a 10 e avalia a evidência disponível, não a taxa de detecção ou a eficácia global. 0 indica ausência de evidência na dimensão avaliada, 2–3 indicam implementação ou observação parcial, 4–5 indicam testes repetíveis de componentes, 6–7 indicam integração demonstrada para o escopo descrito, e 8–10 exigem evidência ampliada em campo e validação independente. A pontuação abaixo vale para o escopo testado, incluindo os limites explicitados em cada linha.

| Dimensão | 3.0.15 | 3.0.16 | Evidência e limite |
| --- | ---: | ---: | --- |
| Detecção por assinaturas | 6 | 6 | ClamAV embarcado com controles EICAR e benigno em CI. Taxa real de detecção e cobertura completa em campo pendentes. |
| Heurística | 4 | 4 | Regressões de PE/ZIP e classificação da base; alerta consultivo no piloto. Calibração em corpus legítimo permanece aberta. |
| Análise comportamental | 4 | 4 | Testes de correlação e limites de eventos em behaviorwatch. Cenários físicos e prevenção ainda sem validação. |
| Proteção em tempo real | 2 | 2 | Monitor de alterações e agendamento no código; piloto rápido com monitor próprio inativo, sem teste de bloqueio pré-execução. |
| Antirransomware | 2 | 2 | Indicadores consultivos e política defensiva em código; nenhum teste de contenção real nesta revisão. |
| Velocidade | 3 | 3 | Rápida de 134,122 s na referência, sem par de execuções equivalentes para medir ganho na 3.0.16. |
| Cobertura | 4 | 4 | Contagens por motor e regressões de raiz controlada; ZIPs fora de limite e arquivos bloqueados permanecem explicitados. |
| Consumo de CPU | 3 | 3 | Medidores e política de workers presentes, sem comparação por processo em campo. |
| Consumo de RAM | 3 | 3 | Limites e política por capacidade testados; memória e paginação por motor em host de 4 GiB ainda sem série comparável. |
| Estabilidade | 4 | 5 | Rápida parcial concluída na 3.0.15. Na 3.0.16, cinco repetições do teste de espera/cancelamento e CI integrado. Completa física e watchdog de I/O pendentes. |
| Quarentena | 5 | 5 | Regressão de isolamento e restauração em corpus controlado; ciclo físico pós-reinício pendente. |
| Remediação | 4 | 4 | Regressões fail-closed e exclusão de evidência de teste da política; reparação e bloqueio reais sem validação ampliada. |
| Proteção web | 3 | 3 | Testes de análise de links; bloqueio comprovado pelo navegador e cobertura de DNS criptografado pendentes. |
| Integração Defender | 4 | 4 | Auditoria e identificação de sinais do Windows; consulta ficou inconclusiva no piloto rápido de referência. |
| Inteligência artificial | 4 | 4 | Testes de inferência consultiva e diagnóstico; validação sintética, sem medição de eficácia antimalware ou aprendizagem contínua. |
| Interface | 5 | 5 | Abertura WebView2 e renderização de painel de CPU/RAM em CI; fluxo visual completo e gráfico por motor pendentes. |
| Relatórios | 5 | 5 | JSON final no piloto rápido e regressões de persistência/snapshot/cópia; conciliação da completa e recuperação seguem abertas. |
| Otimização | 4 | 4 | Testes de seleção de temporários e mecanismos nativos no código; ganho físico e reversão ampla pendentes. |
| Jogos | 4 | 4 | Regressões de regras offline, xadrez e catálogo; interface, acessibilidade e consumo por título precisam de teste físico. |
| Instalador e atualizações | 6 | 6 | Migração e instalação no CI, testes de assinatura do atualizador em fixtures. Authenticode e atualização de produção continuam pendentes. |

Evidências vinculadas ao [CI da 3.0.15](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45755065493), ao [CI da 3.0.16](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45846486811) e à [auditoria de campo](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/a5711f7/AUDITORIA_CAMPO_2026-11-25.md). A classificação é **LAB experimental**. A circulação do ZIP fica restrita a piloto controlado. Distribuição geral depende de conclusão física, cobertura, estabilidade, revisão de licenças e validação do canal assinado.

## Resultado da execução e evidências

A branch `lab-integracao-p0-20261126` integra a correção de espera do ClamAV no [commit funcional 56e0ee2](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/56e0ee2cf1978d8b22845631a8ea48ffd0c11c34). O [CI 45846486811](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45846486811) concluiu com sucesso: regressões Ubuntu e Windows (7/7 etapas em cada job), compilação/validação Windows (48/48 etapas), abertura WebView2, migração, instância única, cinco ciclos de instalação e cinco ciclos instalado com arquivo benigno e controle EICAR do ClamAV. Quinze cenários acumulados de espera (três casos × cinco repetições) passaram na suíte local com detector de corrida, que também aprovou 13 outros módulos.

O [instalador LAB para Windows x64](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45846486811/artifacts/15131069171) foi baixado e conferido. ZIP de 192.818.099 bytes, CRC íntegro, SHA-256 `a8bbcbefd9d11cb0a3ae116132f14bb33eb7758ca4bf44200939ddd5f5e60a16`. Arquivo `ZEUVORYX-LAB-Setup-3.0.16.exe` de 193.291.909 bytes, SHA-256 `db31e5e42ed29a726c1cadfc016d193260de5fc06bd06c06b7fbeecd414cbd5f`. O instalador não possui assinatura Authenticode. Artefato temporário, com vencimento em 03/12/2026 às 12h20 (Brasília).

O [registro de requisitos e aceitação](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/MATRIZ_MESTRE_2026-11-26.md) cobre as 19 fases, 43 REQ e 23 artigos de unificação. Permanecem abertos REQ-003/004/025/027/041 para validação da completa física, watchdog de I/O, pausa/retomada, desempenho e cobertura. O limite por modalidade continua 5/60/20 minutos para rápida/completa/personalizada. O CI comprova o candidato experimental para piloto e os testes que executou. Distribuição geral segue pendente de validação física e assinatura. About/Topics administrativos ainda dependem de credencial com acesso para escrita (403). O código do candidato segue na branch LAB.

O código da correção fica nos modelos `lab/parallel314_windows.go.txt`, `lab/signature312_windows.go.txt` e no módulo `scopeprogress/await.go`. A regressão fica em `scopeprogress/shutdown_test.go`. Identidade e pipeline Windows atualizados em `prepare_version315.py`, `.github/workflows/windows-build.yml` e `.github/workflows/windows-olympus-p0.yml`. A comparação antes/depois mede comportamento de encerramento em teste, com velocidade e eficácia de campo ainda sem comparação equivalente.


## Determinações complementares — 3.0.17 LAB

| Item | Vinculação | Resultado ou pendência |
| --- | --- | --- |
| 1. Evidência de campo | Fases 1, 2 e 19 | Issue #5 revisada e controle compactado preservado em regressão. Original indisponível. |
| 2–4. Escopos, arquivo grande e bloqueado | Fases 3, 4 e 16 | Finalização preservada, diagnóstico dos limites e mensagens de bloqueio incorporado. Pausa, retomada, suspensão, reagendamento e validação física pendentes. |
| 5. Falso positivo no workflow | Fases 5–7 | Alerta consultivo. Correlação, validação do artefato original e corpus benigno seguem pendentes. |
| 6–8. Residente, Windows e motores | Fases 5 e 8 | Consulta Windows ajustada; ClamAV/FreshClam e interfaces auditados. Bloqueio residente próprio segue sem comprovação. |
| 9–10. Matemática e remediação | Fases 4 e 7 | Denominadores de motores, erros nativos, mensagens e limites separados. Desempenho comparativo e ciclos de remediação em campo seguem pendentes. |
| 11–12. Otimização e IA | Fases 6 e 11 | Regressões da base, sem novo recurso anunciado; ensaios de disco e modelo independente abertos. |
| 13–16. Relatórios, testes, GitHub e build | Fases 12, 14, 17–19 | Diagnósticos estruturados e manifesto no artefato do CI. Nenhuma aprovação geral de produção. |
| 20 itens de auditoria adicional ClamAV | Fases 5, 15–17 | [Configuração, arquitetura, contadores, licenças, alternativas e testes pendentes](AUDITORIA_MOTORES_2026-11-26.md). |
