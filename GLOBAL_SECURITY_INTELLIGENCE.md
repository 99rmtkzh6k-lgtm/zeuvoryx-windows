# ZEUVORYX Global Security Intelligence — especificação ZVX-EVO-009

Data: 24/11/2026. Status: arquitetura proposta e requisito aprovado pelo titular; não implementada, hospedada ou contratada. Windows e Macintosh, com paridade validada separadamente.

## Canal único, independente da distribuição
- Um serviço oficial HTTPS versionado, com domínio controlado pelo titular a definir; configuração autenticada, sem descobrir servidores arbitrários. Clientes Store, GitHub Releases, site oficial, instaladores autorizados e outros canais permitidos usam o mesmo protocolo, independentemente de canal, idioma ou país. Não obrigar cadastro para proteção local.
- Consentimento explícito, granular e revogável para estatísticas e avaliação; nenhuma perda de proteção offline. Instalação pseudônima aleatória e credencial individual revogável, sem fingerprint de hardware. Autenticação, limites de requisição, replay defense; credenciais nunca embutidas no binário. Não enviar arquivos, hashes de arquivos, IPs em dados analíticos, caminhos, nomes, URLs visitadas, localização precisa ou histórico detalhado. Minimizar logs de infraestrutura e definir retenção/LGPD/GDPR.
- Avaliação interna única por instalação elegível, seleção explícita 1–5 conforme ZVX-EVO-008, primeira confirmação atômica; reinício/atualização/reenvio não gera novo voto. Não contar estrelas amarelas prévias. Reinstalação e múltiplos dispositivos não equivalem a uma pessoa; não prometer deduplicação universal sem identidade legítima. Aplicar exclusão/retificação legal por procedimento auditado.
- Avaliações externas separadas: Microsoft Store via Analytics API / Partner Center autorizados; GitHub stars são interesse/bookmarks, não notas de 1–5; downloads são métricas distintas. Exibir fonte, período, contagem e horário próprios. Não combinar médias externas e internas sem deduplicação confiável; não copiar estrelas para a loja nem automatizar votos. APIs de outras lojas somente quando autorizadas, respeitando termos e consentimento.

## Dicionário de métricas
| Métrica | Definição |
| --- | --- |
| Instalações participantes | IDs elegíveis distintos, não pessoas nem total de downloads |
| Instalações ativas | Participantes com atividade validada no intervalo declarado, por exemplo 30 dias |
| Dispositivos/instalações com ameaça tratada | Distintos com remediação verificada; não declarar computadores salvos ou danos evitados sem evidência |
| Varreduras | Concluídas, parciais, canceladas e falhas separadas |
| Arquivos examinados | Operações de análise concluídas, não arquivos únicos no planeta; reanálise legítima conta como trabalho, reenvio não conta |
| Bytes analisados | Bytes realmente lidos pelo motor, sem somar repetição de transmissão; registrar volume lógico separado de I/O real, evitar duplicar leituras de hash/heurística; apresentar B/MB/GB/TB/PB SI e opcional KiB/MiB/GiB. Não usar tamanho nominal do arquivo como bytes realmente lidos |
| Detecções | Incidentes com evidência, suspeitas heurísticas à parte |
| Bloqueios | Ações efetivamente concluídas, inclusive rede/phishing, não apenas reputação consultada |
| Quarentenas | Isolamento reversível concluído, identidade e recuperação verificadas |
| Remoções | Exclusão de componente malicioso confirmada, não simples alerta nem quarentena |
| Neutralizações | Incidente único com contenção/remediação verificada; não somar etapas do mesmo incidente |
| Reparações | Resultado de recuperação verificado; não alegar reparo de arquivo sem evidência |
| Restaurações/falsos positivos | Eventos compensatórios auditados e correção dos totais aplicáveis |
| Avaliação | Histograma 1–5, soma, quantidade, média ponderada das notas internas válidas; sem nota quando N=0 |

Categorias: vírus, worms, trojans, ransomware, spyware, stealers, backdoors/RAT, rootkits, adware, PUA/PUP (separado de malware), downloaders/droppers, botnet, cryptominers maliciosos, phishing, URLs/redes maliciosas (categoria de rede distinta), outros/não classificados. Uma categoria primária por incidente para totais somáveis; tags secundárias não entram na soma. Não chamar heurística isolada de infecção confirmada. Registrar motor responsável (ZEUVORYX, Defender, ClamAV, outro); evitar dupla contagem de mesmo incidente entre motores. Não publicar nomes de usuários ou amostras. LAB, EICAR e simulações isolados e nunca somados à produção.

## Arquitetura inicial a avaliar
Cliente Go/Swift -> fila local limitada de incrementos consentidos -> API HTTPS (ingestão autenticada, validação de esquema, versão, idempotência e plausibilidade) -> transação atômica de diário/contador/avaliação -> agregação por categoria e motor -> snapshot de leitura assinado e cache -> painel mundial, site oficial e documentação/lojas quando permitidas. Endpoint de leitura público somente agregados; nenhum endpoint de ajuste arbitrário. Identificadores aleatórios de lote e sequência por instalação; deduplicação atômica, recibo de lote, retry com backoff e jitter, recuperação após falha; não confiar em assinatura de cliente como prova de eficácia. Segregar funções, chaves KMS, checkpoints assinados e cópia imutável por prazo definido; correções auditáveis e direitos de titulares. Não enviar dados reais no laboratório.

Candidato piloto: Cloudflare Workers + D1 (SQL gerenciado, execução sob demanda); avaliar limites de escrita/concorrência, residência de dados, custos, disponibilidade, segurança, precisão dos inteiros, transações e backups. Alternativa de escala: API Go e PostgreSQL gerenciado, restrições UNIQUE e transações, agregação assíncrona e cache/CDN. Escolher após teste de carga, orçamento e aprovação; não contratar nem provisionar serviços sem autorização. GitHub é código/CI/distribuição, não servidor de telemetria. Não armazenar segredos em repositório.

## Consumo, interface e verificação
- Manter prioridade ultraleve: gráfico numérico estático (barras simples opcionais sob aprovação, sem animação/mapa), contadores e notas, dados locais/globais identificados. Consulta cacheada enquanto painel visível, alvo inicial 5 minutos; lotes de envio a cada 15 minutos se houver mudanças, adiar durante varreduras intensivas, sincronizar após reconexão; não enviar em repouso sem necessidade. Mostrar horário UTC/local, idade, abrangência de participantes e estado indisponível; nenhum valor inventado.
- Inteiros 64-bit com limites/overflow validados e representação decimal exata em JSON/JS; avaliar acumuladores maiores para bytes antes de atingir limite. Payload agregado alvo inferior a 1 KiB para resumo, categorias detalhadas sob demanda, medição real CPU/RAM/disco/rede. A meta não é desempenho comprovado.
- Gates obrigatórios em cada revisão: matriz Windows/Mac Intel/ARM64 e todos canais autorizados; compatibilidade de API e migração, teste de dois clientes em canais distintos convergindo para mesmo snapshot, duplicidade, reenvio, concorrência, votos, offline, consentimento revogado, falha/rollback, falsos positivos, classificação, integridade e precisão, carga e custo. Teste de referência dos totais e reconciliar diário; não promover sem evidência. Atualização do aplicativo respeita mecanismo de cada canal (Store gerencia seus pacotes; GitHub/instalador oficial usa manifesto assinado); não substituir pacote Store arbitrariamente. Atualização de dados mundial independente da atualização do executável.

## Referências oficiais consultadas
- Microsoft Store ratings/reviews: https://learn.microsoft.com/pt-br/partner-center/insights/reviews-report ; https://learn.microsoft.com/en-us/windows/uwp/monetize/get-app-ratings ; https://learn.microsoft.com/en-us/windows/uwp/monetize/access-analytics-data-using-windows-store-services .
- GitHub stars: https://docs.github.com/en/rest/activity/starring .
- Cloudflare D1/Workers limites e custos: https://developers.cloudflare.com/d1/platform/limits/ ; https://developers.cloudflare.com/d1/reference/faq/ ; https://developers.cloudflare.com/workers/platform/storage-options/ .
- A seleção é proposta de engenharia, não garantia de gratuidade, escala, conformidade ou aprovação de loja.
