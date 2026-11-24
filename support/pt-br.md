# ZEUVORYX Antivirus — Central de ajuda

Base pública de suporte para ZEUVORYX Windows 3.0.6 LAB R1, candidato experimental de segurança complementar. Esta documentação não anuncia uma versão de produção, certificação independente nem substituição do Microsoft Defender.

[Português](pt-br.md) · [English](en.md) · [Español](es.md) · [日本語](ja.md)

## O que é o ZEUVORYX?

Scanner de segurança para Windows x64 desenvolvido em Go, com análise de arquivos, indicadores SHA-256, heurísticas, quarentena reversível e IA consultiva experimental. Macintosh 3.0.5 LAB é um projeto separado em validação.

## Instalação e atualização

Use somente um pacote de teste autorizado pelo mantenedor e confira sua versão e SHA-256 com a referência do GitHub. Extraia o ZIP antes de abrir o instalador. Termine ou cancele a varredura, feche o aplicativo, faça backup de configurações e quarentena e instale sobre a edição LAB compatível. Não desative Defender, SmartScreen ou outras proteções. O CI verificou migração de 3.0.5 LAB para 3.0.6 LAB; outras migrações não estão comprovadas.

## O aplicativo não abre ou fecha sozinho

Não repita a instalação indefinidamente. Confirme a versão, se o Windows é x64, se o Microsoft Edge WebView2 Runtime está instalado e se há espaço e memória disponíveis. Consulte o Histórico de confiabilidade (perfmon /rel) e o Visualizador de Eventos (eventvwr.msc), logs Aplicativo e Sistema, no horário da falha. Não apague a quarentena nem desative a proteção do Windows. Envie somente informações anonimizadas pelo canal de suporte.

## A varredura completa trava ou encerra

O fechamento inesperado em computadores físicos ainda está em investigação. A versão 3.0.6 corrigiu pontos de inicialização e foi aberta por 15 segundos após a confirmação da interface no CI; 60 testes de um pequeno conjunto de arquivos passaram. Isso não demonstra estabilidade em toda a varredura. Se ocorrer novamente, interrompa os testes na máquina de trabalho, registre horário, versão, modo, uso de CPU/RAM e último progresso. Não execute arquivos suspeitos.

## Erros, ignorados e alertas

Acesso negado, arquivo bloqueado por outro processo e arquivo que desapareceu podem impedir a leitura. Um relatório parcial ou zero detecções não comprova ausência de malware. Uma heurística ou pontuação de IA é um indício, não diagnóstico definitivo. Não remova arquivos do Windows apenas por um alerta; revise a evidência e mantenha as proteções do sistema.

## IA, tempo real, ClamAV e atualizações

A IA atual é consultiva, sem aprendizagem local contínua validada. Monitoramento residente próprio, aprendizagem contínua, atualização automática de produção e estatísticas globais permanecem em desenvolvimento. ClamAV é integração opcional e depende de instalação externa. Não confunda aprovação do botão Testar componentes com serviço global operacional. Releases de produção não foram anunciadas.

## Como solicitar suporte

Abra uma questão no GitHub Issues indicando versão, Windows 10/11 e build, arquitetura, RAM aproximada, modo da varredura, horário e fuso, resultado esperado, resultado observado e passos para reproduzir. Se possível, informe a categoria do erro e o código de saída. Remova nomes de pessoas, caminhos pessoais, e-mails, tokens, identificadores, conteúdo de arquivos e hashes de arquivos privados. Não publique amostras de malware, logs brutos, arquivos de quarentena ou credenciais. Não há atendimento 24 horas garantido.

---

- [GitHub](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows) · [Issues / support](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/issues)
- [Windows 3.0.6 LAB R1 CI evidence](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45479946405)
- [Historical memorandum](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/main/MEMORANDO_DESENVOLVIMENTO.md)
- Updated: 2026-11-24.