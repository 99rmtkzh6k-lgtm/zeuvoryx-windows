# Memorando de desenvolvimento — ZEUVORYX Windows
Atualização de 24/11/2026 — decisão para a etapa posterior aos testes da 3.0.2 Competitive Beta.
Autoria: Created & Engineered by Leonardo A Botelho.

## Prioridades permanentes
- Foco exclusivo em computadores Windows; pesquisa automotiva adiada.
- Proteção, estabilidade, velocidade, baixo consumo e simplicidade; português, inglês e espanhol.
- IA local experimental: não substituir modelo ativo por candidato reprovado; validação independente por categoria, com dados reais rotulados e controle de falsos positivos.
- Manter Microsoft Defender ativo; não anunciar substituição, proteção residente própria ou eficácia não comprovada.
- Preservar main e o pacote de teste 3.0.2, sem alterações funcionais durante a validação.

## Sequência aprovada
1. Testar o pacote 3.0.2 da execução 45385340904 em Windows Sandbox: instalação, idiomas, interface, varreduras, cancelamento, consumo, detecção inofensiva, isolamento e restauração, logs e estabilidade.
2. Revisar os relatórios JSON, resultados e falhas; corrigir bloqueadores e repetir testes.
3. Somente se os resultados forem satisfatórios, fazer piloto controlado em Windows real, preferencialmente máquina de testes sem dados sensíveis, com backup verificável, Defender ativo e sem malware real. Não desativar a proteção existente. Não instalar em máquina crítica enquanto houver falhas relevantes.
4. Analisar resultados do piloto e planejar a próxima versão em ramificação separada. Não ativar automaticamente funções destrutivas.

## Cinco implementações futuras (pendentes)
| Função | Requisitos e critérios de aceitação |
| --- | --- |
| Remoção automática | Apenas evidência confiável, opção expressa do usuário, quarentena reversível por padrão, auditoria, revalidação do arquivo e proteção de caminhos críticos. IA/heurística isoladas nunca autorizam remoção. |
| Interrupção de processos maliciosos | Identificação e confirmação do processo, proteção de processos críticos, tratamento seguro de erros, privilégio mínimo; inicialmente modo observação. |
| Limpeza de persistência | Inventário e revisão de itens suspeitos de inicialização, serviços, tarefas e configurações relevantes; backup e rollback; sem alterações indiscriminadas do registro ou serviços. |
| Reparação | Diagnóstico de alterações, mecanismos oficiais do Windows quando cabíveis, backup, consentimento, restauração verificada; sem prometer recuperar arquivos criptografados. |
| Proteção residente | Monitoramento inicialmente passivo, fila limitada, métricas de consumo, compatibilidade com Defender, testes de corrida/estabilidade; bloqueio só após validação e revisão da arquitetura, assinatura e requisitos do Windows. |

## Liberação
Separar cada módulo, implementar testes unitários, integração, regressão, falhas e recuperação em laboratório autorizado com artefatos inofensivos. Avaliar impacto de falsos positivos, privilégios, estabilidade e desempenho. Nenhuma ação automática irreversível na beta atual. Exigir revisão humana antes de promover nova versão. Treinamento e validação de eficácia real são pendentes e independentes da compilação.

## Situação atual
A versão 3.0.2 compilou no GitHub Actions; validação manual no Windows ainda pendente. Quarentena e exclusão manual existem com critérios de evidência; remoção pelo Defender é integração com motor externo. As cinco melhorias acima não foram implantadas nesta atualização do memorando.

## Evolução futura — ZEUVORYX Global Protection Statistics & Feedback
Solicitação de 24/11/2026. Criar painel de impacto cumulativo, histórico de segurança local e estatísticas mundiais de usuários participantes, com atualização próxima de tempo real quando conectados. Implementação pendente para depois da validação da 3.0.2.

### Histórico local e contadores
- Registrar eventos de varredura e monitoramento, detecção, bloqueio, quarentena, remoção confirmada, restauração, falha e reparação verificada. Incluir data/hora, categoria de ameaça, motor responsável (ZEUVORYX, Defender ou outro integrado), versão e resultado. Ações de teste, simulações e EICAR devem ser identificadas e excluídas das métricas reais.
- Manter totais históricos locais persistentes, inclusive após fechar o programa, reiniciar o Windows ou atualizar o produto. Oferecer exportação, política de retenção e exclusão voluntária de histórico; não prometer retenção após desinstalação, perda de armazenamento ou pedido de exclusão. Proteger os dados locais e limitar consumo de disco.
- Usar contadores inteiros de 64 bits, suportar bilhões de eventos e manter totais separados de detecções, bloqueios, quarentenas, exclusões confirmadas, reparações e restaurações. Uma detecção não equivale a uma remoção. Uma ação de quarentena não equivale a uma exclusão definitiva. Não somar etapas do mesmo incidente como múltiplas ameaças eliminadas. Registrar correções de falsos positivos e restaurações sem inflar totais.

### Painel mundial
- Com autorização expressa e revogável do usuário para telemetria, enviar somente eventos mínimos necessários, sem arquivos, caminhos, nomes de usuário, IPs no histórico analítico, conteúdo, senhas ou identificadores de hardware. Identificador aleatório pseudônimo de instalação, limitado à finalidade e sujeito à política de privacidade e LGPD/GDPR e requisitos locais. País apenas opcional, autodeclarado ou agregado de forma compatível com privacidade; nunca geolocalização precisa. Sem participação obrigatória nem prejuízo à proteção offline.
- Criar serviço central seguro para agregar eventos de instalações participantes de qualquer país (inclusive Japão e Brasil), com transporte TLS, autenticação, limitação de requisições, validação, identificadores de eventos idempotentes, deduplicação, fila local limitada e sincronização após reconexão. Não tratar declarações do cliente como prova incontestável: prever fraude, auditoria, correção de métricas e atribuição ao motor responsável.
- Mostrar totais históricos mundiais de ações confirmadas, categorias (vírus, spyware, ransomware etc.), instalações participantes e dispositivos únicos com remediação confirmada (não afirmar que são "computadores salvos" nem que incidentes prevenidos teriam causado danos). Separar laboratório, telemetria real, ações do Defender e do ZEUVORYX. Exibir horário da última atualização, abrangência e limitações; offline mostrar último total sincronizado identificado como desatualizado.
- Não usar eventos de usuários como rótulos automáticos para treinamento; eventual pesquisa deve usar consentimento específico e validação independente. Não divulgar rankings ou números inventados.

### Avaliações de usuários
- Avaliação voluntária de 1 a 5 estrelas dentro do aplicativo, após experiência suficiente ou em menu próprio; não interromper varreduras, não vincular avaliação à proteção e não exigir nota positiva.
- Mostrar média, número de avaliações e distribuição de estrelas, com atualização após sincronização. Permitir revisar a avaliação, prevenir duplicidade e manipulação, não associar avaliação a identidade pública ou histórico de segurança. Feedback textual opcional separado, sem dados pessoais.
- Antes da implantação, definir custos e infraestrutura do serviço, política de privacidade, consentimento, exclusão/retificação, retenção, incidentes e segurança. Testar offline/online, concorrência, falhas, reenvio, deduplicação, integridade, contagens de bilhões, privacidade, fraude, acessibilidade e impacto em CPU/RAM/rede. Proibir publicação de métricas fictícias como reais.

## Requisito prioritário — estatísticas globais ultraleves (24/11/2026)
- Substitui a sugestão anterior de mapa mundial: apresentar exclusivamente números e rótulos curtos (totais cumulativos, categorias e avaliações numéricas). Sem mapas, gráficos animados, imagens, feeds individuais ou detalhes de ataques no painel global.
- Enviar somente incrementos de contadores por categoria e resultado confirmado, em lotes compactos, com metadados técnicos mínimos para autenticação, idempotência, versão do protocolo e separação de testes/produção. Não transmitir arquivos, hashes de arquivos, nomes, caminhos, conteúdo, geolocalização ou histórico detalhado. Histórico detalhado de diagnóstico, se habilitado, fica local e tem retenção limitada.
- Usar contadores persistentes de 64 bits com deduplicação no servidor; sincronização assíncrona e eventual, sem blockchain, mineração, rede P2P ou conexão permanente obrigatória. A referência ao Bitcoin significa somente atualização de números compartilhados mundialmente, não adoção de sua arquitetura.
- Atualização global por consulta compacta em intervalos moderados (referência inicial: 5 minutos enquanto o painel estiver visível); envio agregado em intervalos maiores (referência: 15 minutos quando houver mudanças), com cache, backoff, suspensão durante varreduras intensivas e ausência de consultas se não necessárias. Sem bloquear scanner, inicialização ou interface; offline usar último total sincronizado, indicando sua idade.
- Meta inicial de engenharia: payload de totais inferior a 1 KiB por resposta em formato compacto, excluindo overhead HTTP/TLS; sem processo residente separado para estatísticas; CPU média incremental próxima de zero em repouso e memória incremental alvo inferior a 1 MiB para lógica e cache, sem contar infraestrutura compartilhada. Medir tráfego real incluindo overhead, CPU, RAM, inicialização, disco, latência e impacto na varredura em Windows antes de liberar; metas não são resultados nem garantia de consumo zero.
- Consentimento opcional, privacidade, contagem de remoções verificadas, correção de falsos positivos, distinção ZEUVORYX/Defender e avaliações voluntárias continuam obrigatórios. Nenhum evento individual público. Segurança e desempenho têm prioridade sobre atualização instantânea. Implementação pendente até validação da 3.0.2; não alterar o executável atual.
