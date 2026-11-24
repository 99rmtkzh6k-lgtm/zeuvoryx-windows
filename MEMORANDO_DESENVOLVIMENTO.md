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
