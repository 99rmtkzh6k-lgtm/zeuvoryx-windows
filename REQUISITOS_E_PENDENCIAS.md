# ZEUVORYX — registro cumulativo de requisitos e pendências

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
