# ZEUVORYX 3.0.5 LAB R5 — Olympus Silver

Scanner complementar de segurança para Windows x64, desenvolvido em Go, com interface desktop, análise de arquivos e links, quarentena reversível e IA consultiva experimental. **Created & Engineered by Leonardo A Botelho.**

[![Windows LAB R5](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml/badge.svg?branch=security-3.0.5-lab)](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/workflows/windows-build-lab-305.yml)

**Estado:** beta de laboratório. Compilação 3.0.5 R5 aprovada em 24/11/2026; validação manual em Windows 10/11 e eficácia independente pendentes. Não substitui Microsoft Defender.

**Atenção:** esta é a página de apresentação da versão mais recente. A ramificação `main` ainda conserva o código histórico 3.0.1. Para código, compilação e testes da versão 3.0.5, use exclusivamente a ramificação [security-3.0.5-lab](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/security-3.0.5-lab). Não execute o workflow antigo para obter a versão nova.

## Download e instalação de teste

- [Compilação validada — execução 45419964764](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45419964764)
- [ZIP completo — artefato 14934540577](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45419964764/artifacts/14934540577) (33.730.421 bytes; GitHub Actions pode exigir login e tem prazo de retenção).
- SHA-256 registrado pelo GitHub: `c42366ea44b6ac25c302b92a1fdbacb237c01224eb8b4e13b3a2e34ada19fcf0`.
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

Execução [45419964764](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45419964764), commit `11cd939ff0944bf2da9a81f6f4bc166559e5384b`: gate de versão 3.0.5, cinco rodadas de regressão de links, dez repetições de testes, QA20 (20 milhões de iterações sintéticas), build EXE, instalador e teste de reinstalação aprovados. Laboratório de IA com 50 milhões de cenários sintéticos não comprova proteção real; testes de laboratório não equivalem a certificação.

**Validação adicional:** a execução de fundamentos de segurança [45419964425](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45419964425) teve os jobs Windows e Linux aprovados e o job macOS cancelado; não constitui aprovação integral multiplataforma.

## Desenvolvimento e continuidade

A ramificação ativa é `security-3.0.5-lab`. O workflow `.github/workflows/windows-build-lab-305.yml` chama `windows-build.yml` nessa ramificação. Para uma nova compilação, selecione a ramificação de laboratório; caso o GitHub não disponibilize execução manual desse workflow na página principal, use a execução validada acima. A `main` preserva o código anterior 3.0.1 e não deve ser confundida com o laboratório mais recente. Consulte [instruções](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/security-3.0.5-lab/INSTRUCOES_GITHUB.txt), [instalação](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/security-3.0.5-lab/LEIA-ME.txt), [histórico de versões](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/security-3.0.5-lab/CHANGELOG.md) e o [memorando obrigatório](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/security-3.0.5-lab/MEMORANDO_DESENVOLVIMENTO.md).

Próximos passos: painel com comandos de varredura e métricas no topo, botões compactos 3D; teste visual e Windows real; atualização assinada de produção; auditoria ampliada, remediação segura, proteção residente, IA validada, estatísticas globais ultraleves e Macintosh independente. Idiomas adicionais conforme qualidade/desempenho. Nenhum desses requisitos é considerado concluído só por estar planejado.

## Licenciamento, autoria e distribuição

Autoria declarada: Leonardo A Botelho. Não há licença de redistribuição de produção definida neste documento; revisar licenças de dependências e direitos de arte antes de publicação. O projeto não reivindica certificação, detecção universal ou substituição de antivírus. Releases públicas e lançamento mundial dependem de autorização do titular e dos critérios do memorando.

## Versão histórica

A versão 3.0.1 Competitive permanece preservada no [commit histórico c7d2f3c](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/tree/c7d2f3c94fd080368eff59921d338a0dd5fa41ca). Não é o download recomendado para avaliação da versão 3.0.5.
