# ZEUVORYX Macintosh — candidato atual 3.0.5 LAB

> O registro inicial abaixo documenta a base histórica 3.0.3 LAB; a situação corrente está na seção ZVX-EVO-002 ao final.
24/11/2026 — autorização do titular. Created & Engineered by Leonardo A Botelho.
Base comum: Windows 3.0.3 LAB, arte oficial e requisitos do MEMORANDO_DESENVOLVIMENTO.md.
Pipeline independente: GitHub Actions macOS 15 ARM64 e Intel; Go + Swift/SwiftUI nativo. Pacotes DMG separados, assinatura ad hoc de laboratório, sem notarização. Não publicar.
Implementação inicial: scanner local somente leitura, SHA-256 integral de arquivos dentro de limites, marcador EICAR exclusivamente de teste, alerta de extensão dupla, diagnóstico local dos módulos de validação de atualizações/estatísticas/política e interface pt/en/es. Sem base de assinaturas real, motor de IA macOS integrado, remoção, quarentena, reparação, proteção residente, telemetria ou atualizador real. Sem acesso elevado, alterações de configurações ou desativação de proteção Apple.
Pendências: validar código e UI em Macs reais; cobertura de Mach-O, arquivos compactados, scripts e assinaturas legítimas, falso positivo, cancelamento, desempenho, integração e treinamento IA específico; quarentena reversível com preservação de metadados, remediação segura, Endpoint Security mediante autorização Apple, assinatura Developer ID e notarização, distribuição e backend comum de estatísticas voluntárias. Não contabilizar testes como proteção real. Não misturar métricas Windows/macOS nem liberar atualizações sem autorização do titular.
Repositório atual é Windows; ramificação macintosh-3.0.3-lab separa trabalho e workflows. Criar repositório privado próprio quando titular puder fazê-lo; conector atual não cria repositórios. Sem serviços externos contratados. GitHub Actions não é servidor de telemetria.

## Bateria de equivalência — 24/11/2026
- Foco somente computadores; iOS adiado, sem desenvolvimento de aplicativo iOS.
- Build inicial Intel aprovado (45396314583); ARM64 aguardava runner na última consulta.
- Reaproveitados módulos e testes Windows de fraude, links, inteligência, reputação, persistência e diagnósticos, em pacote portátil isolado. Não significa integração à UI; dados e consultas nos testes são sintéticos ou simulados. Removidos somente testes diretamente dependentes do Defender; permanecem explicitamente não aplicáveis ao macOS.
- CI ampliado em Intel e ARM64: dez repetições, race detector, fuzz do scanner e das assinaturas de manifestos/snapshots, benchmark compartilhado de 50 milhões de cenários sintéticos e dez milhões de mensagens sintéticas por arquitetura. Não equivalem a amostras reais nem ao teste de motor macOS com IA integrada.
- Ainda pendentes: equivalentes reais para PE/Mach-O, ZIP, UTF16, base IOC ativa no scanner, quarentena/limpeza, proteção residente, integração da UI, assinatura Developer ID, notarização e testes em Mac real. Não classificar pendências como aprovadas; sem publicação pública. Preservar Windows.

## Correção GitHub — 24/11/2026
- README da ramificação corrigido de Windows 3.0.2 para Macintosh 3.0.3 LAB. Execução 45396837719: Intel aprovado, artefato 14921063483, ARM64 cancelado, sem log disponível. Reteste solicitado (tentativa 2); não afirmar aprovação ARM64 até resultado e artefato. Main Windows consolidada em 3.0.5 LAB R5, com launcher Macintosh disponível para seleção desta ramificação. Nenhuma alteração do motor Macintosh ou publicação de produção nesta etapa.

## Candidato 3.0.5 LAB — 24/11/2026
- Consultado memorando central e histórico antes das alterações. Preservar versão 3.0.3, sem renomear pacote existente. Portar analisador local de links R5 e testes exatos do Windows; expor na CLI e UI somente como diagnóstico consultivo, sem abrir sites nem bloquear conexões. Gate de versão CLI, interface e Info.plist; testes Linux de módulos e cross-build Go Darwin, além de CI nativo Intel/ARM64 e DMG. Não declarar testes aprovados antes da evidência. Sem paridade integral, alteração do modelo IA, assinatura oficial, atualização automática, publicação ou mudança de permissões.

## Regra permanente obrigatória — sincronização técnica e metadados do GitHub (24/11/2026)

Aplicável a todas as atualizações, correções, revisões, compilações e versões do ZEUVORYX, em Windows e Macintosh, inclusive futuras plataformas autorizadas. A manutenção documental e dos metadados faz parte da entrega, não é uma tarefa opcional.

1. Antes de alterar: consultar o memorando central, memorando da plataforma, histórico, repositório, ramificações, versão vigente, CI e erros anteriores; verificar os campos públicos atuais e registrar divergências.
2. Após alterar: atualizar e conferir a seção About/Sobre do GitHub — Description/Descrição, Topics/Tópicos, Website/Homepage e opções de exibição de Releases/Packages/Deployments, quando aplicáveis. A descrição deve refletir funcionalidades verificadas, plataformas, versão estável ou LAB e limitações; tópicos relevantes, sem duplicidade ou alegações não comprovadas; website deve apontar para página oficial ou documentação atual válida. Campos que já estão corretos não precisam ser alterados, mas devem ser conferidos.
3. Sincronizar README, CHANGELOG, memorandos, ABOUT_GITHUB.md, instruções de instalação/compilação, documentação técnica, requisitos/dependências, identidade e versão na interface, motor, atualizador, metadados do instalador, workflows, ramificações, badges, links, nomes de artefatos, tamanho e SHA-256 confirmado. Verificar tags, notas de versão, Releases e manifesto de atualização quando aplicáveis, sem publicar release, alterar visibilidade, licença, permissões, chaves ou configurações sensíveis sem autorização do titular. Preservar histórico e distinguir plataformas/arquiteturas.
4. Depois de testes e compilação, reler os arquivos e metadados efetivamente publicados no GitHub e conferir links/artefatos e correspondência entre versão, commit, plataforma, arquitetura e evidência de CI. Registrar data, alterações, falhas, correções e resultados, inclusive pendentes/cancelados. Não confundir candidato com versão aprovada, teste sintético com eficácia real, ou texto proposto em arquivo com campo About efetivamente atualizado.
5. Executar diretamente todas as alterações autorizadas que a integração permitir. Se faltar acesso ou função administrativa, registrar cada campo pendente, motivo, localização/link exato, instruções em português para o titular (inclusive Safari/iPhone), e texto integral pronto para copiar e colar; depois conferir novamente no GitHub. Não declarar atualização completa enquanto houver campos obrigatórios sem verificação ou pendências não resolvidas.
6. Encerrar cada revisão com checklist de campos verificados, atualizados, não aplicáveis, pendentes de CI e pendentes de ação manual, com evidências e responsáveis. Mudanças apenas documentais não exigem recompilação funcional desnecessária. Preservar as prioridades de segurança, estabilidade, velocidade, baixo consumo, simplicidade e privacidade, bem como o controle exclusivo das publicações oficiais pelo titular.

Sequência obrigatória: memorando → auditoria histórica e técnica → correção → testes/retestes → compilação → conferência de versão/artefatos → sincronização de documentação e About/Topics/Website → verificação efetiva no GitHub → registro de pendências e resultados no memorando.

### Conferência inicial desta regra — 24/11/2026
- Repositório compartilhado atual: https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows. Description preenchida com Windows 3.0.5 LAB R5 e recursos experimentais; Topics vazio; Website/Homepage vazio; default branch main. A integração disponível permite ler esses campos e editar documentação, mas não oferece operação de edição administrativa de About/Topics/Website. Pendências manuais: inserir tópicos e link de documentação oficial; revisar descrição incluindo Macintosh como candidato LAB, sem alegar aprovação ou paridade. Textos e passos completos em https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/ABOUT_GITHUB.md. Confirmar os valores efetivos após aplicação pelo titular.

## Diretriz permanente — padronização, memorando vivo e melhoria contínua (24/11/2026)

Determinação do titular: toda atualização deve ser tratada como parte de um sistema integrado. Se um componente mudar, conferir todos os componentes dependentes, plataformas, versões, documentação, distribuição e metadados do GitHub. Atualizar o que for afetado; registrar o que não se aplica e por quê. Não modificar código ou configurações corretos apenas por formalidade.

- Fonte de verdade: o memorando central em main, complementado pelos memorandos específicos e histórico Git. Consultar a versão efetiva antes de cada intervenção; reconciliar divergências nas ramificações ativas, sem reescrever fatos históricos nem depender apenas da memória da conversa. Manter índice, seções e identificadores de requisitos, decisões, incidentes e melhorias à medida que o volume crescer. O documento é evolutivo, sem limite artificial de extensão; preservar histórico e explicitar decisões substituídas.
- Padrão de engenharia: mapear impacto em Windows x64, macOS Intel/ARM64, motor, interface, dependências, idiomas, instalador, atualizador, logs, segurança, desempenho, CI, versão e artefatos. Compartilhar componentes quando compatíveis; manter APIs e testes específicos da plataforma. Não sincronizar artificialmente números de versões nem alegar paridade sem evidência. Usar versões/revisões coerentes, critérios de aceitação verificáveis e referência ao commit exato.
- Aprender com falhas: registrar sintoma, comando/arquivo afetado, causa quando comprovada, correção, teste de regressão, resultado, link e prevenção. Revisar erros antigos (YAML, caminhos, nomes, versões, dependências, duplicações, cancelamentos e permissões); não repetir tentativas cegamente. Se a causa não estiver comprovada, registrar hipótese e evidência faltante. Priorizar testes automatizados, validações estáticas e controles de consistência que evitem reincidência.
- Pesquisa técnica proativa em cada ciclo de desenvolvimento: procurar novidades relevantes em fontes primárias atuais (documentação oficial Go, GitHub Actions, Microsoft, Apple, avisos de segurança e projetos/dependências oficiais), quando houver impacto potencial. Avaliar novas ferramentas, APIs, algoritmos, testes, otimizações, controles de segurança, acessibilidade e automações; registrar fonte/URL, data, versão, benefício, risco, compatibilidade, licença, custo, consumo e requisitos. Distinguir descoberta, proposta, experimento, implementação e aprovação; quando não houver novidade útil, registrar isso sem inventar pesquisas ou resultados.
- Melhorias descobertas fora do pedido imediato: registrar no backlog, com prioridade, plataformas afetadas, dependências, critério de aceitação e próximo passo. Corrigir defeitos diretamente relacionados e autorizados, com regressão e preservação das versões aprovadas; não ampliar escopo de maneira silenciosa, contratar serviços, adicionar dependências invasivas, habilitar remediação destrutiva, publicar releases ou alterar acessos sem autorização aplicável. Propor melhorias maiores ao titular antes de promoção. Novidade não é superioridade comprovada.
- Padronização verificável: considerar fonte única de versão, catálogo de metadados e checklist automatizado de consistência, análise de dependências, SBOM, verificações de segurança e CI, inicialmente como propostas a avaliar. Fixar critérios e testar em LAB antes de adotar; manter cadeia de suprimentos verificável e não desativar proteções para aprovar um teste. Prioridades: proteção, velocidade, estabilidade, baixo consumo, simplicidade e privacidade.
- Fechamento obrigatório: sincronizar documentação, About/Description/Topics/Website, histórico e evidências conforme a regra anterior; reler efetivamente no GitHub; registrar campos sem acesso de edição, instruções manuais e responsáveis. Classificar cada item como verificado, corrigido, pendente, não aplicável ou bloqueado. Não declarar aprovação integral se testes/validação ainda estiverem pendentes.

Modelo mínimo de registro de cada evolução: data | identificador | versão/plataforma | solicitação ou descoberta | fonte/evidência | impacto/dependências | decisão/prioridade | arquivos/commit | testes e resultado | documentação/metadados | pendências/responsável | próximo passo.

O ciclo de evolução aplica-se durante cada retomada e revisão do projeto; não pressupõe memória infinita, pesquisa permanente fora da sessão ou execução em segundo plano. GitHub e memorandos versionados são a referência durável.

### Verificação de consistência nesta revisão
- Windows 3.0.5 R5: execução 45436304817 aprovada; fundamentos 45436304993 ainda na fila nesta consulta. Macintosh candidato 3.0.5: execução 45437873318 na fila; reteste 3.0.3 45396837719 na fila. Nenhuma aprovação adicional inferida.
- About: descrição ainda somente Windows; Topics e Website vazios. Ação manual continua registrada em ABOUT_GITHUB.md da main, sem capacidade administrativa de escrita nesta integração.
- Histórico CHANGELOG da main mantinha a execução LAB original sem registrar a compilação consolidada em main e o candidato Macintosh; complementar o histórico sem apagar a evidência anterior.
- Esta revisão é documental; nenhuma alteração funcional, novo teste, lançamento ou adoção de ferramenta foi realizada.

## Protocolo permanente de ativação e padrão internacional de engenharia — 24/11/2026

**Expressão de referência do titular:** “memorando da evolução do antivírus”, “memorando da evolução do ZEUVORYX” ou simplesmente “memorando”, quando o assunto for o ZEUVORYX, designam o conjunto completo de regras deste memorando central e dos memorandos específicos Windows e Macintosh. Ao receber a expressão, consultar primeiro os arquivos vigentes no GitHub e o histórico, verificar os resultados atuais e aplicar o ciclo de revisão, pesquisa, correção, testes, sincronização, documentação e verificação anteriormente estabelecido, conforme o escopo solicitado. Quando o pedido for apenas consultar, apresentar ou registrar, não iniciar alteração funcional ou publicação por inferência. Não exigir que o titular repita os requisitos; registrar novas descobertas e propostas úteis, respeitando autorizações e prioridades. Os arquivos Git versionados são a referência durável, não uma promessa de memória ilimitada.

### Referências públicas consultadas em 24/11/2026
- Microsoft SDL: integrar requisitos de segurança, análise de ameaças e controles no ciclo de engenharia. https://www.microsoft.com/en-us/securityengineering/sdl/practices
- Google Cloud/SLSA: procedência verificável de compilações e controles de cadeia de suprimentos. https://cloud.google.com/resources/delivering-software-securely ; https://slsa.dev/spec/v1.2/
- Apple Developer ID: assinatura e notarização de software macOS distribuído fora da App Store. https://developer.apple.com/developer-id/
- IBM Security and Privacy by Design: segurança e privacidade desde o projeto e revisão antes da disponibilização geral. https://www.ibm.com/trust/security-and-privacy-design
- Adobe Secure Product Lifecycle: incorporar segurança ao ciclo de vida de produtos e serviços. https://www.adobe.com/trust/security/product-security.html
- GitHub Secure Use: fixar Actions por SHA integral verificado, minimizar permissões, revisar dependências e evitar injeção em workflows. https://docs.github.com/en/actions/reference/security/secure-use
- NIST SSDF SP 800-218: estruturar preparação, proteção do software, produção segura e resposta a vulnerabilidades. https://csrc.nist.gov/pubs/sp/800/218/final
- OWASP SAMM: avaliar maturidade de governança, projeto, implementação, verificação e operações. https://owaspsamm.org/model/
- Go Vulnerability Management: avaliar govulncheck para identificar vulnerabilidades conhecidas em código e dependências Go. https://go.dev/doc/security/vuln/

São referências públicas selecionadas, não acesso a procedimentos internos de todas as empresas, nem certificação ou conformidade já demonstrada do ZEUVORYX. Revisar vigência, licença, aplicabilidade e custos antes de adotar ferramentas.

### Controles de engenharia a incorporar incrementalmente
| ID | Controle | Critério de aceitação |
| --- | --- | --- |
| ZVX-STD-01 | Fonte única de versão e matriz de compatibilidade | Interface, motor, instalador, workflow e metadados correspondem; versões históricas e candidatos identificados, Intel/ARM64/x64 separados. |
| ZVX-STD-02 | Modelagem de ameaças e revisão de arquitetura | Registrar ativos, fronteiras de confiança, riscos e mitigação; privilégio mínimo e nenhuma ação destrutiva sem controles. |
| ZVX-STD-03 | Revisão de código e testes por risco | Testes positivos, negativos, regressão, concorrência, fuzzing e falhas; distinguir testes sintéticos, integração e eficácia real. |
| ZVX-STD-04 | Cadeia de suprimentos e dependências | Avaliar govulncheck, validação YAML/actionlint, Actions fixadas por SHA verificado, dependências auditadas, inventário SBOM e procedência/atestação de artefatos. Sem desligar checks, sem segredos em logs. |
| ZVX-STD-05 | Gates de liberação e recuperação | LAB → testes por plataforma → piloto controlado → aprovação do titular → assinatura/release autorizada; testar falhas de atualização, preservação de dados e rollback. |
| ZVX-STD-06 | Gestão de vulnerabilidades e incidentes | Planejar SECURITY.md, canal responsável, triagem, correção, regressão, registro de causa raiz e comunicação proporcional. Não divulgar segredos ou dados sensíveis. |
| ZVX-STD-07 | Qualidade mensurável e privacidade | Medir CPU, RAM, disco, tempo, cobertura e falsos positivos em condições reproduzíveis; consentimento, minimização e retenção; não inventar indicadores. |
| ZVX-STD-08 | Melhoria contínua e documentação | Registrar decisões de arquitetura (ADR), backlog com IDs, revisão de riscos, critérios de aceitação e checklist de encerramento; atualizar About, README, CHANGELOG e memorandos e verificar campos reais. |
| ZVX-STD-09 | IA responsável | Versionar modelo/dados e critérios de validação; comparar contra referência, medir falsos positivos por categoria, não promover candidato reprovado, não usar IA isoladamente para ação destrutiva. |

Status inicial destes controles: diretrizes adotadas no processo documental; automações, ferramentas e requisitos técnicos adicionais são backlog a avaliar/testar, não implantação já comprovada. Priorizar os controles de consistência e dependências no CI sem aumentar o consumo do aplicativo instalado. A cada revisão, registrar implementado/testado/aprovado/pendente com evidência real e não declarar certificação NIST, OWASP, SLSA ou de laboratório independente.

### Conferência atual do GitHub
- Windows 3.0.5 R5: compilação 45436304817 aprovada. Fundamentos 45436304993: Windows/Linux aprovados, job macos-updater cancelado, execução global falhou. Não confundir cancelamento com defeito de código comprovado.
- Macintosh 3.0.3: Intel aprovado, ARM64 cancelado novamente (45396837719). Macintosh 3.0.5 candidato (45437873318): testes portáteis Linux e compilação Intel aprovados; ARM64 na fila. Nenhuma aprovação integral multiplataforma ainda.
- About: Description ainda apenas Windows; Topics e Website vazios. Ação manual permanece em ABOUT_GITHUB.md; conferir depois da edição.
- Esta revisão incorpora pesquisa e regras documentais; não instala ferramentas, não altera executáveis, não dispara testes desnecessários nem publica releases.

## Primeira aplicação do protocolo de evolução — 24/11/2026 (ZVX-EVO-001)

Solicitação do titular: aplicar a expressão de referência e o ciclo completo de continuidade aos antivírus Windows e Macintosh, registrando descobertas e melhorias. Esta é uma revisão documental e de auditoria, não autorização de release ou mudança funcional indiscriminada. O protocolo permanece no GitHub, não representa treinamento individual do modelo ou garantia de memória ilimitada.

### Verificação efetiva
- Windows 3.0.5 R5: execução 45436304817 aprovada, commit 214a6de2586433176987efdad8147ce3a1b02c41. Pacote 14942002196 preservado. Fundamentos 45436304993: Windows e Linux aprovados; macos-updater cancelado, execução global falhou.
- Macintosh 3.0.5 LAB: execução 45437873318, commit 2226582057b455ba6ed906a78efb233ffecbe902. Regressão portátil e compilação Intel aprovadas; ARM64 na fila. Pacote Intel 14943370193 preservado. Macintosh 3.0.3: execução 45396837719 com Intel aprovado e ARM64 cancelado novamente, histórico preservado.
- Revisão da documentação GitHub (https://docs.github.com/en/actions/reference/runners/github-hosted-runners): macos-15-intel e macos-15 são rótulos válidos Intel e ARM64, respectivamente. Não trocar rótulos sem evidência. Causa do cancelamento não demonstrada; avaliar detalhes do job e disponibilidade do ambiente antes de novo reteste, sem repetir tentativas cegamente.
- Metadados efetivamente lidos: Description ainda descreve Windows 3.0.5 LAB R5, Topics vazio e Homepage vazio. Edição administrativa não disponível na integração. Ação manual permanece conforme ABOUT_GITHUB.md da main. Repositório principal é a fonte de verdade; a cópia histórica do memorando Windows dentro da ramificação Macintosh não substitui o memorando central atual.

### Pesquisa e melhorias propostas
| ID | Prioridade | Melhoria e fonte | Aceitação / status |
| --- | --- | --- | --- |
| ZVX-EVO-001-A | Alta | actionlint — https://github.com/rhysd/actionlint | Validar workflows e reproduzir regressões YAML/expressões antes de build; selecionar versão fixada, revisar licença e compatibilidade. Proposta, não instalado. |
| ZVX-EVO-001-B | Alta | govulncheck — https://go.dev/doc/security/vuln/ | Analisar fontes efetivamente compiladas e dependências por plataforma, incluindo arquivos gerados no build Windows; documentar achados, correções e falsos positivos. Proposta, não executado. |
| ZVX-EVO-001-C | Alta | Atestação de procedência GitHub — https://docs.github.com/en/actions/how-tos/secure-your-work/use-artifact-attestations/use-artifact-attestations | Avaliar permissões mínimas e elegibilidade; vincular EXE/DMG a commit e workflow, verificar atestação; não substitui Authenticode/Developer ID. Proposta, não implantado. |
| ZVX-EVO-001-D | Alta | Instruções de continuidade AGENTS.md | Arquivo de orientação na raiz das três ramificações ativas, com links canônicos e critérios do memorando; não depende exclusivamente da memória da conversa. Não garante execução automática. |

Pesquisa realizada em fontes públicas de mantenedores nesta revisão. As ferramentas A–C são candidatas de CI, não dependências residentes do produto; aferir custos, riscos, licenças, versões e compatibilidade antes de adotar. Nenhuma nova compilação ou teste funcional executado por esta revisão documental.

### Próxima ordem técnica
1. Resolver bloqueio ARM64 com diagnóstico fundamentado, preservar evidências e validar nativamente pacote macOS; não tratar cancelamento como falha de código confirmada.
2. Avaliar gates de consistência (versão e YAML) e análise de dependências em LAB, com resultados antes de promoção.
3. Prosseguir conforme backlog anterior: painel Windows com comandos/métricas superiores e botões compactos 3D, validação real dos instaladores, assinatura de produção, atualização e recuperação seguras, paridade Macintosh e eficácia independente. Não alterar ordem sem registrar motivo.
4. Sincronizar documentação, concluir About manualmente pelo titular e verificar valores efetivos. Não publicar release.

Classificação desta etapa: auditoria realizada; atualização documental e instruções de continuidade implementadas após verificação de gravação; ferramentas novas propostas; validação ARM64 e produção pendentes; metadados administrativos bloqueados por capacidade da integração. Preservar histórico e os nove controles ZVX-STD.


## ZVX-EVO-002 — conferência de versão e CI (24/11/2026)

Macintosh 3.0.5 LAB: regressão portátil Linux e compilação Intel aprovadas; Apple Silicon ARM64 cancelado na tentativa anterior, sem etapas registradas e com log indisponível (404). Reteste exclusivo solicitado e aceito pelo GitHub em 24/11/2026; resultado ainda não confirmado. Não há causa de código comprovada.
- Execução [45437873318](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318); job ARM64 original 138600529888. O workflow usa macos-15 para ARM64 e macos-15-intel para Intel, com fail-fast desativado. Não alterar rótulos ou código sem evidência.
- Windows 3.0.5 LAB R5: compilação [45436304817](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304817) confirmada como aprovada. Macintosh 3.0.3 LAB é histórico preservado. A 3.0.5 LAB é o candidato atual, não uma versão aprovada para distribuição integral.
- Revisão exclusivamente documental; sem mudança de versão, motor, instaladores, permissões, releases ou assinatura. Estados anteriores nos registros históricos não representam a situação atual. Validação em computadores reais, eficácia independente e assinatura de produção seguem pendentes.
- About/Topics/Website: edição administrativa não disponível nesta integração; instruções permanecem em [ABOUT_GITHUB.md](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/ABOUT_GITHUB.md). Conferir campos reais após ação do titular.
