# ZEUVORYX — Protocolo de integração operacional contínua

Vigência: 25/11/2026. Autorização do idealizador: executar tarefas de atualização pelo conector GitHub e GitHub Actions, dispensando intervenção manual quando tecnicamente possível. Escopo: exclusivamente este repositório Windows. O memorando permanente e os requisitos de validação continuam aplicáveis.

## Ciclo operacional
1. Consultar HEAD, memorando, auditoria, requisitos e CI antes de alterar arquivos. Ramificação candidata inicial: lab-olympus-p0-defender-20261125; CI de referência 45683175191.
2. Investigar evidências e logs; corrigir problemas comprovados em candidato isolado, preservando arquitetura, autoria, interface, funcionalidades e proteção existente.
3. Executar testes regressivos, disparar CI por commit e evitar compilações concorrentes redundantes. Registrar SHA, escopo, resultado e falhas; não mascarar testes.
4. Verificar artefato do instalador, integridade e proveniência. Sem teto de tamanho interno. Entregar ZIP completo no ChatGPT quando o conector permitir, para compartilhamento pelo WhatsApp.
5. Sincronizar README, CHANGELOG, memorando, requisitos, central de atualizações e documentos relevantes; main somente documentação, sem promoção do código candidato. Manter descrição pública e tópicos coerentes quando houver ação administrativa autorizada e disponível. Não inventar sucesso.
6. Manter validação física separada de aprovação em CI. Não publicar comercialmente nem instalar remotamente em PCs de terceiros. Não desativar proteções, alterar secrets/permissões ou transferir dados pessoais.
7. A tarefa de integração contínua do ChatGPT verifica periodicamente o projeto, sem gerar mudanças fictícias nem notificações repetidas. A periodicidade atual é horária. Não depende de navegador aberto, mas depende de autorização vigente e disponibilidade de ferramentas e serviços.

## Pendências físicas atuais
Preservar exportação JSON comprovada pelo titular. Relatório parcial c8bf1fd: 203 de 212 arquivos, 9 ocupados, ClamAV em andamento, Sysmon inconclusivo, monitor próprio desativado. Investigar retry seguro, duração e estado de cada motor e relatório final, sem presumir travamento. A remoção do teto de tamanho é alteração separada e não comprova resolução dos problemas de campo.

## Limites operacionais
O conector confirma leitura/escrita de conteúdo e consulta de CI, mas não oferece escrita direta nos campos administrativos About/Topics ou criação de releases neste ambiente. Não contornar autorização por credenciais, tokens ou workflows privilegiados. Testes em hardware físico e aprovações exigidas pelo serviço permanecem fora da autonomia automática.
