## Candidato corretivo de campo — 25/11/2026
Nova ramificação operacional: [lab-olympus-field-p0-20261125](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/lab-olympus-field-p0-20261125); [commit 7a2c456](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/7a2c45627265d65e54b00022460526a4d94cf4c1); [CI 45685145731](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45685145731) em execução na consulta. Preservada a compilação anterior 45683175191, sem cancelamento por este commit. Correções: duas novas tentativas canceláveis apenas para compartilhamento Windows 32; testes com falha persistente e handle exclusivo; estado e duração ClamAV; início serial antes de aguardar diagnósticos; Sysmon cancelável com diagnóstico diferenciado. Monitor de alterações continua dependente da configuração existente e não é bloqueio pré-execução. Não declarar cobertura residente ou eficácia integral comprovada. Preservado JSON físico funcional, interface e regra sem limite interno de tamanho. O instalador anteriormente aprovado continua LAB; novo candidato ainda não aprovado. A tarefa horária deve acompanhar esta ramificação sucessora e o CI indicado, entregar ZIP completo após validação. Main somente documentação.

## Determinação vigente do idealizador — 25/11/2026: Artigo 37 substituído

Não há limite interno obrigatório de tamanho para o ZEUVORYX: revogados o teto de 200 MB e a meta numérica de 150 MB. O instalador, ZIP, motores, bases e dependências poderão crescer quando tecnicamente necessários. Priorizar proteção, eficácia, estabilidade e desempenho; continuar medindo e otimizando tamanho, RAM, CPU, disco e downloads sem remover cobertura ou funcionalidades aprovadas apenas para reduzir bytes. Registrar crescimento e justificativa, respeitar limites externos dos canais de distribuição e manter controles de integridade, licenças, gratuidade e validação. Registros anteriores de limites são históricos. Entregar sempre que possível o ZIP completo no ChatGPT para compartilhar pelo WhatsApp, não somente o link GitHub.

# ZEUVORYX Antivirus 3.0.15 LAB — Olympus | Referência vigente 25/11/2026

**CI aprovada; candidato Windows x64 para teste físico controlado. Não é versão comercial.**

- Código efetivamente compilado: [c8bf1fd](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/c8bf1fd7614e44de297569663fc90fa5ab066a47), ramificação [Olympus P0](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/lab-olympus-p0-defender-20261125).
- [CI integral](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45675394546): 48/48 etapas Windows concluídas com sucesso, cinco ciclos de regressão de módulos Windows e Linux, dez ciclos gerais e cinco testes do aplicativo instalado no CI. Migração 3.0.6 → 3.0.15 testada no CI.
- [ZIP do instalador LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45675394546/artifacts/15052358598): 192807638 bytes (192,81 MB decimais), abaixo do teto histórico de 200 MB (revogado); SHA-256 do ZIP informado pelo GitHub: bd9bcd3461e7189944a9a295b5a5e4e1fff447652c798d87e00aa24b4ca56744. Não é o hash individual do EXE; não houve verificação independente dos bytes nesta consulta. O portátil (295087036 bytes) é um pacote separado de teste; a entrega recomendada continua sendo o instalador completo.
- Revisão incremental: três commits desde 935b906, três arquivos alterados; distinção entre comando Defender aceito e eventos de varredura observados (1000/1001/1002), teste de contrato, workflow isolado e auditoria. Vínculo exclusivo dos eventos, execução PowerShell real e exportação integrada da evidência ainda não comprovados.
- Preservar interface, autoria Leonardo A. Botelho, identidade Olympus e requisitos REQ-001–043. Nenhuma promoção de código para main; main permanece referência histórica com documentação atualizada. Manter proteções Windows ativas.
- P0 pendentes: teste físico rápido/completo, RAM 2/4/8/16 GiB, progresso/ETA, copiar JSON e recuperação, cobertura dos motores, Defender por sessão e eficácia independente. CI não comprova proteção integral. Sem publicação comercial, atualização automática de produção ou cobrança.
- About público consultado: já informa 3.0.15 LAB e validação física em andamento; homepage correta; Topics vazio. O conector não expõe escrita administrativa desses campos. Sugestões de tópicos: antivirus, windows, golang, cybersecurity, endpoint-security, clamav, threat-detection.

## Registros históricos — preservar para auditoria

# Mapa técnico do candidato 3.0.15 LAB

- `.github`: workflows de geração e validação.
- `engine`, `scriptguard`, `adaptive`: scanner e análises.
- `signatureav`, `auditmetrics`: motor de assinaturas e evidências.
- `fieldlearn`, `behaviorwatch`, `changeguard`: observação e monitoramento experimental.
- `scopeprogress`: inventário do escopo.
- `remediation`, `audit`, `releaseguard`, `updater`: remediação, auditoria e controles de atualização.
- `nativeamsi`: integração experimental, gate operacional não comprovado.
- `locales`, `familygate`, `gamecatalog`: interface, família e jogos.
- `cmd`, `globalintel`, `impact`, `lab`, `labcheck`, `qa`, `research`: ferramentas, dados, testes e pesquisa.

A árvore candidata tinha 233 arquivos antes desta revisão documental; comparação com o commit ad42721 confirmou apenas diferenças em três documentos antes desta revisão. Não significa auditoria individual de eficácia de todos os arquivos. Não alterar timestamps artificialmente. A versão main é histórica; não promover antes da aceitação física.
