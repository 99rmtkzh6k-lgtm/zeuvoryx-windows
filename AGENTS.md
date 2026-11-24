# ZEUVORYX — protocolo obrigatório de continuidade

Estas instruções orientam o trabalho no repositório; não substituem a autorização do titular nem os memorandos. Em toda tarefa ZEUVORYX, aplicar o protocolo completo conforme o escopo solicitado, mesmo sem menção expressa ao memorando.

1. Consultar primeiro o [memorando central vigente na main](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/MEMORANDO_DESENVOLVIMENTO.md), o [memorando Macintosh](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/macintosh-3.0.5-lab/MEMORANDO_MACINTOSH.md), o histórico e CI atual. Registros históricos não substituem decisões posteriores.
2. Mapear impacto em Windows x64, macOS Intel/ARM64, módulos comuns, interface, idiomas, versão, instalador, atualizador, documentação e metadados. Não alterar componente correto por formalidade. Não confundir build aprovado com eficácia real.
3. Revisar falhas anteriores; pesquisar fontes técnicas atuais quando aplicável; registrar descobertas, riscos, critérios e pendências. Corrigir no escopo autorizado, preservar versões aprovadas, testar e retestar; não inventar resultados.
4. Aplicar ZVX-STD-01 a 09: identidade, arquitetura, testes, dependências, liberação, vulnerabilidades, desempenho/privacidade, documentação e IA responsável. Proteção, velocidade, baixo consumo, estabilidade e simplicidade têm precedência.
5. Sincronizar os arquivos afetados e conferir About/Description/Topics/Website, README, CHANGELOG, memorandos, workflows, links, versões e hashes. Se faltar acesso, documentar passos manuais e conferir depois. Registrar evidências por commit, plataforma e execução.
6. Não liberar publicamente, gerar chaves oficiais em CI, alterar permissões, contratar serviços ou ativar ações destrutivas sem autorização aplicável. Manter as proteções nativas dos sistemas ativos. Não promover candidato reprovado.
7. Se o pedido for apenas documental, não modificar código ou disparar build desnecessário. Registrar falhas, causas comprovadas, correções, testes, pendências e próximos passos. O memorando é a fonte durável; não supor memória infinita ou trabalho em segundo plano.

A leitura destas instruções por agentes depende do suporte da ferramenta utilizada. A existência deste arquivo não garante execução automática nem aprovação dos controles.

## Regra de encerramento obrigatória — ZVX-CP-002

Aplica-se a toda tarefa ZEUVORYX, sem palavra-chave ou lembrete. Antes de agir, consultar os memorandos vigentes; após qualquer correção, evolução, teste, reteste ou compilação, atualizar o memorando oficial da main e o da plataforma/ramificação afetada no mesmo atendimento. Registrar evidência, commit, CI, versão, resultados, falhas e pendências; sincronizar documentação e metadados pertinentes; reler o conteúdo remoto e confirmar o commit. Se a execução ainda estiver em andamento, registrar esse estado; ao consultar novo resultado, atualizar novamente. Não dar a tarefa por documentalmente concluída se a gravação ou verificação falhar; declarar o bloqueio. Não publicar ou alterar componentes não autorizados. Eficiência e todos os requisitos cumulativos da ZVX-CP-001 continuam obrigatórios.

## ZVX-CP-003 — consistência de versão obrigatória

Em toda nova versão funcional autorizada, incrementar e conferir identidade do motor, UI, idiomas, instalador/registro, atualizador, auditoria, pacote, workflow, README, CHANGELOG, memorandos, links e metadados públicos. Não reutilizar número anterior para atualização que deva ser detectada. Preservar histórico e dados; diferenciar candidato LAB e versão aprovada. Após validação, sincronizar About/Description/Topics/Website e releases autorizadas; se acesso administrativo faltar, registrar impedimento e proposta exata. Não aprovar pacote divergente; verificar migração e execução. CP-001/002 permanecem vigentes.
