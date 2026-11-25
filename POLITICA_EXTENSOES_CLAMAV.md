## Auditoria operacional — 25/11/2026 (consulta de CI)
Candidato acumulado vigente: [ad379d2](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/commit/ad379d2774830a4928500631b38ad80f375e87da), [CI 45687017435](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45687017435). Regressões Windows e Linux concluídas com sucesso (7/7 etapas cada); build Windows 28/47 etapas aprovadas, testes sintéticos de fraude em andamento, nenhuma etapa reprovada na consulta. Instalador final do candidato ainda não confirmado. Revisão de campo 7a2c456: CI 45685145731 aprovado (48/48 Windows), não equivale ao candidato acumulado. As revisões intermediárias e main não devem substituir o candidato prioritário. Não disparar build adicional durante a validação, salvo correção de falha comprovada.
Requisitos vigentes: ClamAV/FreshClam integrado e gratuito, sem instalação externa; sem teto interno de 200 MB ou meta obrigatória de 150 MB; extensões comerciais exclusivamente opcionais e ainda limitadas a catálogo informativo. Cabeçalhos de versões anteriores são históricos. Referências funcionais para extensões: Warden/Sentinel da Danami e documentação oficial ClamAV, sem cópia proprietária. Pendências: teste físico offline, proteção residente própria, Sysmon, cobertura e desempenho físico, integridade e disponibilidade do instalador, atualização assinada, ativação isolada de extensões e revisão jurídica/licenças antes de produção. About informa 3.0.15 LAB; Topics vazios, alteração administrativa não disponível no conector. Não promover código a main.

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

## Referências de arquitetura
Warden e Sentinel (Danami): configuração voluntária de fornecedores; ClamAV oficial: requisitos de assinatura e integridade. Reproduzir apenas capacidades documentadas, sem copiar implementação proprietária. A ativação continua sujeita a testes e autorização/licenças do fornecedor.
