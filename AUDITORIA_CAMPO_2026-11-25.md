# Auditoria de campo — 25/11/2026 — 3.0.14 LAB, revisão candidata

## Evidência
Relatório físico de varredura rápida do commit a4603eb66de729fa85ba8bdaaa37c71f99b7e151: 37.254 arquivos / 18.275.842.703 bytes em 1.345.427 ms, cancelada pelo operador. Inventário parcial de 41.085 arquivos e 19.049.436.242 bytes, uma falha de inventário e sete conflitos de compartilhamento. ClamAV 1.5.5 / base 28194 respondeu, mas apenas uma raiz consta no relatório; limite MaxScanTime em PDF, não malware confirmado; resumo ClamAV ausente, números -1 significam indisponíveis. Sysmon inconclusivo. Não publicar caminhos, nomes de usuários ou o relatório bruto.

## Correções candidatas já escritas
- Tempo global indeterminado enquanto motores independentes não concluírem. Ocultar barra percentual quando inventário parcial; distinguir escopo lógico e disco ocupado.
- Publicar resultados ClamAV de cada raiz antes do fim, preservá-los nas cópias do scanner próprio e no JSON parcial. Não apresentar -1 como número medido. Manter estado em execução até concluir raízes ou cancelar.
- Limpar sinais recentes na nova análise, separar detecções por assinatura, heurísticas/limites e EICAR; aviso visível com significado correto. Contagem recente limitada a 100, independente do scanner próprio, não somar.
- Teste de regressão de preservação de relatórios e exportação adicionado. CI e teste físico ainda não verificados.

## Prioridade máxima, não implementada por esta revisão
1. Proteção residente antes da execução, serviço assinado e monitoramento de eventos com cobertura explícita; integração Defender/SmartScreen sem desativar, diagnóstico de proteção e atualização reais, sinalizar sistema sem suporte.
2. Proteção web/phishing em navegação, reputação de URL via fornecedores autorizados, integração de navegador consentida, DNS e proteção de rede; não interceptar HTTPS indiscriminadamente.
3. Memória, comportamento e antirransomware: avaliar recursos nativos e SDK OEM, sem alegar cobertura quando telemetria indisponível. IA somente após corpus rotulado, validação independente e controle de falsos positivos. YARA-X e indicadores licenciados como complemento.
4. Desempenho: medir CPU, RAM, I/O, p95, tempo a frio/quente; avaliar clamd com memória e isolamento limitados, não multiplicar motores residentes; reprocessar arquivos alterados com cache seguro, tratar limites e arquivos inacessíveis sem apagar nem declarar seguro.
5. Relatório nativo ZEUVORYX: incluir origem de cada evidência, Windows/Defender/Sysmon, cobertura por motor, diagnóstico, eventos, ações e erros; exportação JSON parcial e completa, sem transferência automática. Separar ameaças confirmadas, suspeitas, testes, limites e falhas. Quarentena reversível e consentida, sem exclusão de arquivos críticos.
6. Manutenção: confirmar limpeza, seleção individual e otimização nativa SSD/HD, diagnóstico de CPU/RAM/disco; prioridade absoluta é segurança/estabilidade, não jogos ou estética.
7. Ferramentas comerciais: avaliar SDKs oficiais Bitdefender OEM e Avira OEM; somente integrar com licença válida e permissão de redistribuição; nenhum gasto contratado sem autorização expressa. Não confundir teste gratuito com distribuição gratuita. Alternativas livres mantidas conforme licença.
8. Gate obrigatório: cinco ciclos de testes funcionais por módulo, CI Windows, teste de instalação e físico, arquivos benignos e EICAR/AMTSO em ambiente seguro, avaliação laboratorial independente de proteção, falsos positivos e desempenho. Não declarar superioridade ou liberar como antivírus autônomo até evidência suficiente. Preservar main, versão pública e atualizações assinadas.

Referências: https://docs.clamav.net/manual/Usage/Scanning.html ; https://www.bitdefender.com/en-us/oem/ ; https://oem.avira.com/en/solutions/anti-malware ; https://learn.microsoft.com/en-us/windows/security/operating-system-security/
