# Política técnica — isolamento de extensões ClamAV

Vigência: 25/11/2026. Base oficial e FreshClam gratuitos, sempre independentes de extensões.

Implementação desta revisão: o processo de varredura com o diretório integrado recebe --official-db-only=yes; catálogo externo é informativo, ativação ainda indisponível. Termos incluídos no painel, no instalador e na instalação. Não habilitar bytecode não assinado.

Critérios obrigatórios antes de ativar uma extensão:
- Contratação pelo usuário, licença por dispositivo verificada, sem redistribuição de assinaturas comerciais.
- Diretório separado por fornecedor e processo de análise separado do motor gratuito; nenhuma alteração no diretório oficial ou em freshclam.conf oficial.
- Somente formatos de assinatura aprovados; rejeitar bytecode não assinado, executáveis, caminhos relativos, links/reparse points e arquivos fora dos limites.
- Autenticação HTTPS, verificação de integridade e origem conforme fornecedor, armazenamento de segredos em mecanismo seguro do Windows, nunca em logs, relatórios ou GitHub.
- Validação de carregamento em staging, testes benignos e de detecção controlada, ativação atômica e última versão válida quando licença permitir; isolamento de falhas e rollback.
- Limites de tempo, memória, CPU e disco; execução sequencial quando necessário; falha, expiração, cancelamento ou ausência de rede não impede o motor oficial.
- Alertas de fonte identificada, sem somar detecções duplicadas e sem remoção automática por base opcional.
- Testes de independência da proteção gratuita, falhas de rede/credencial/assinatura corrompida, expiração, regressões, falsos positivos e desempenho em hardware físico antes de ativação.
- Nenhuma garantia absoluta de detecção; respeitar GPL e CDC; revisão jurídica antes de publicação comercial.

Referências: https://docs.clamav.net/manual/Signatures.html ; https://docs.clamav.net/manual/Usage/Scanning.html ; https://docs.clamav.net/manual/Installing/Community-projects.html ; https://www.planalto.gov.br/ccivil_03/leis/l8078compilado.htm
