# ZEUVORYX 3.0.5 LAB R5 — Olympus Silver

Scanner complementar de segurança para Windows x64, desenvolvido em Go, com interface desktop, análise de arquivos e links, quarentena reversível e IA consultiva experimental. **Created & Engineered by Leonardo A Botelho.**

[![Windows LAB R5](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml/badge.svg?branch=main)](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml)

**Estado:** beta de laboratório. Compilação 3.0.5 R5 aprovada em 24/11/2026; validação manual em Windows 10/11 e eficácia independente pendentes. Não substitui Microsoft Defender.

**Revisão LAB pós-varredura e Global Threat Intelligence:** CI Windows aprovado no commit `07373e9c538794c5f14a3bc13ad85206d3049bc9`; alterações de interface, classificação de erros, retry limitado, procedência do build e consulta consultiva DROP incluídas. Validação visual e em Windows real ainda necessária.

## Download e instalação de teste

- [Compilação validada — execução 45447823825](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825)
- [ZIP completo — artefato 14946014706](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825/artifacts/14946014706) (33.793.191 bytes; GitHub Actions pode exigir login; expira em 22/02/2027).
- SHA-256 do ZIP, conferido nos bytes baixados: `7232a6b12003b23a0e4de8a15e188d528533a63b15f11c93ea41771fa33111cb`.
- Instalador extraído sem executar: 15.497.251 bytes; SHA-256 `e90bedcbbcdd4cbb41020fac39b93d867beb0d4ebef884ff103e96632451ef39`; sem certificado Authenticode incorporado. ZIP CRC conferido.
- Extraia o ZIP, finalize a varredura anterior e feche o aplicativo. Abra `installer/ZEUVORYX-LAB-Setup-3.0.5.exe`. A edição portátil está em `dist/`.
- Teste primeiro em Windows Sandbox com Defender ativo e WebView2 disponível. Não desative SmartScreen nem proteções existentes. O binário não tem assinatura Authenticode de produção comprovada.
- Para uma instalação LAB compatível, mantenha o diretório anterior (AppId ZEUVORYX.Desktop.Lab303). CI aprovou reinstalação com marcador de dados; faça backup e valide histórico/quarentena. Migração de edições 3.0.1/3.0.2 não validada.
- Nenhuma release pública está sendo promovida por este documento; o botão de atualização assinada requer chave pública oficial fixada, manifesto/pacote assinados e release autorizada.

## Recursos disponíveis para avaliação

| Área | Estado |
| --- | --- |
| Varreduras | Rápida, completa e personalizada; cancelamento, resultados, métricas, exportação e cobertura parcial |
| Motor | SHA-256/IOC, heurísticas, inspeção limitada PE/ZIP, EICAR e análise consultiva de links |
| Quarentena | Ação explícita com critérios de evidência; reversibilidade e restauração sujeitas a validação |
| Segurança Windows | Integrações com Defender e recursos nativos; não é proteção residente própria |
| IA local | Modelo experimental auxiliar, sem aprendizado contínuo em produção; candidato reprovado não promovido |
| Interface | Olympus Silver, Zeus e templo; português, inglês e espanhol, seleção de idioma |
| Histórico | Registro local autenticado, relatório de erros, avaliação voluntária; sem estatísticas globais |
| Atualizador | Implementação experimental de descoberta e validação assinada; produção e rollback ainda pendentes |

## Evidência de CI

Execução [45436304817](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304817), commit `214a6de2586433176987efdad8147ce3a1b02c41`: gate de versão 3.0.5, cinco rodadas de regressão de links, dez repetições de testes, QA20 (20 milhões de iterações sintéticas), build EXE, instalador e teste de reinstalação aprovados. Laboratório de IA com 50 milhões de cenários sintéticos não comprova proteção real; testes de laboratório não equivalem a certificação.

## Desenvolvimento e continuidade

A ramificação `main` consolida a base de código Windows 3.0.5 LAB R5; `security-3.0.5-lab` preserva a referência de laboratório. O workflow `.github/workflows/windows-build-lab-305.yml` chama `windows-build.yml` na ramificação selecionada. A versão antiga está preservada no histórico. A compilação integral em main foi aprovada; a regressão de links independente [45436304882](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304882) também foi aprovada. Na execução de fundamentos [45436304993](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45436304993), Windows e Linux passaram; o job macos-updater foi cancelado e a execução global ficou com falha. Consulte [instruções](INSTRUCOES_GITHUB.txt), [instalação](LEIA-ME.txt), [histórico de versões](CHANGELOG.md) e o [memorando obrigatório](MEMORANDO_DESENVOLVIMENTO.md).

Próximos passos: painel com comandos de varredura e métricas no topo, botões compactos 3D; teste visual e Windows real; atualização assinada de produção; auditoria ampliada, remediação segura, proteção residente, IA validada, estatísticas globais ultraleves e Macintosh independente. Idiomas adicionais conforme qualidade/desempenho. Nenhum desses requisitos é considerado concluído só por estar planejado.

## Macintosh — ZEUVORYX 3.0.5 LAB (candidato atual)

A edição Macintosh **3.0.5 LAB** é o candidato atual, desenvolvido separadamente com Swift/SwiftUI e motor Go somente leitura. [Documentação atual](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/macintosh-3.0.5-lab). A regressão portátil Linux e a compilação Intel foram aprovadas. A tentativa ARM64 foi cancelada sem etapas registradas; reteste exclusivo solicitado, sem aprovação confirmada. Não há equivalência funcional integral com Windows nem autorização de distribuição de produção.

Para novos testes, selecione o workflow Macintosh 3.0.5 e a ramificação `macintosh-3.0.5-lab`. A [versão histórica 3.0.3 LAB](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/macintosh-3.0.3-lab) permanece preservada. Não confundir instaladores, arquiteturas ou candidato com pacote aprovado.

## Licenciamento, autoria e distribuição

Autoria declarada: Leonardo A Botelho. Não há licença de redistribuição de produção definida neste documento; revisar licenças de dependências e direitos de arte antes de publicação. O projeto não reivindica certificação, detecção universal ou substituição de antivírus. Releases públicas e lançamento mundial dependem de autorização do titular e dos critérios do memorando.


### Conferência de CI — 24/11/2026

Macintosh 3.0.5 LAB: regressão portátil Linux e compilação Intel aprovadas; Apple Silicon ARM64 cancelado na tentativa anterior, sem etapas registradas e com log indisponível (404). Reteste exclusivo solicitado e aceito pelo GitHub em 24/11/2026; resultado ainda não confirmado. Não há causa de código comprovada. [Execução 45437873318](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318). [Artefato Intel 14943370193](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45437873318/artifacts/14943370193), 4.283.792 bytes, SHA-256 registrado pelo GitHub `965da629bf97f885a7df1de2434211f3342d82ba4e286f1a1fb52f946bc3f767`. Somente para laboratório; sem validação integral em Macs reais ou assinatura de produção.


## Global Threat Intelligence (CI Windows aprovado; laboratório)

Nova consulta consultiva Spamhaus DROP IPv4/IPv6 e inventário de fontes em [THREAT_INTELLIGENCE.md](THREAT_INTELLIGENCE.md). Não é bloqueio de firewall nem limpeza automática. Base de sessão obtida sob demanda; nenhuma contagem real antecipada. CI Windows aprovado na execução 45447823825; coleta inicial separada 45447823809 falhou com `stale or future feed`, sem artefato de dados. Nenhuma base DROP real embutida. Versões anteriores preservadas.


### Evidência mais recente — ZVX-EVO-007

https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45447823825 — gate de identidade, cinco rodadas de links, dez repetições, QA20, 50M sintético, build e reinstalação aprovados. Não equivale a eficácia real, migração entre versões distintas, validação do feed ao vivo ou distribuição pública. Os registros de CI antigos acima são históricos.
