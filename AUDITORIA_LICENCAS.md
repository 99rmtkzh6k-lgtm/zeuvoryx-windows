# ZEUVORYX — política de independência financeira e inventário preliminar de terceiros

Revisão 24/11/2026. Referência técnica: Windows 3.0.6 LAB R1, workflow windows-build.yml e THREAT_INTELLIGENCE.md. Não constitui parecer jurídico, licença própria nem aprovação de distribuição.

## Regra permanente
Preferir componentes com uso comercial e redistribuição permitidos sem royalties, mensalidades, pagamento compulsório ou exigência de abertura do núcleo proprietário. Gratuidade inicial não basta se a utilização comercial futura exigir contrato pago. Licença livre não significa domínio público: preservar direitos autorais, textos de licença, avisos e demais condições. Nenhuma contratação sem autorização expressa. Não prometer custo operacional zero (infraestrutura, certificados, lojas e limites de serviços são separados). Priorizar proteção/estabilidade e não remover funções de segurança silenciosamente.

| Componente | Papel | Decisão para próxima versão | Fonte |
|---|---|---|---|
| Go 1.24 (CI) | motor/runtime | manter; BSD, incluir licença e avisos efetivos da distribuição | https://go.dev/LICENSE |
| go-webview2 56598839c808 | interface compilada | manter MIT; auditar dependências transitivas e loader Microsoft | https://pkg.go.dev/github.com/jchv/go-webview2@v0.0.0-20260205173254-56598839c808 |
| Microsoft WebView2 Runtime | pré-requisito externo | manter distribuição conforme termos oficiais; preferir runtime oficial instalado | https://learn.microsoft.com/pt-br/microsoft-edge/webview2/concepts/distribution |
| rsrc v0.10.2 | ferramenta de build | manter MIT; não distribuir ferramenta desnecessariamente | https://github.com/akavel/rsrc |
| Pillow 12.2.0 | arte em build | manter somente build; verificar licença exata e componentes do wheel; não instalar no cliente | https://pypi.org/project/pillow/12.2.0/ |
| Inno Setup 6 | instalador atual | migrar para NSIS ou MSIX compatível; FAQ de Inno diz compra comercial solicitada mas não estritamente obrigatória, portanto sem dívida presumida | https://jrsoftware.org/inno-setup/purchase |
| NSIS | alternativa proposta | uso comercial permitido; revisar licença de todos componentes/plugins, usar compressor zlib e evitar plugins desnecessários | https://nsis.sourceforge.io/License |
| C2PA c2patool 0.28.1 | binário auxiliar distribuído | licença upstream permissiva, mas verificar release exata e todas dependências/avisos antes de aprovar; se inviável manter função como opcional indisponível e informar | https://github.com/contentauth/c2patool |
| ClamAV | executável externo opcional | preservar opção já solicitada; não embutir nem vincular libclamav, sem contratação; GPLv2 e termos de bases precisam ser respeitados | https://docs.clamav.net/ |
| MalwareBazaar | consulta API existente | bloquear integração de produção até confirmação expressa de condições comerciais sem custo; manter importação local somente de dados autorizados | https://bazaar.abuse.ch/api/ |
| URLhaus/ThreatFox | importação local/planejado | não implantar API paga; verificar direitos de uso dos dados e atribuições mesmo em importação manual | https://urlhaus.abuse.ch/api/ ; https://threatfox.abuse.ch/api/ |
| Spamhaus DROP | consulta consultiva | gratuito sob condições específicas, crédito técnico, preservação de data e texto, sem marketing de marca/dados, revogável; revisar termos e uso antes de distribuir | https://www.spamhaus.org/blocklists/drop-fair-use-policy/ ; https://www.spamhaus.org/faqs/do-not-route-or-peer-drop/ |
| CISA KEV | consulta catálogo CVE | manter fonte gratuita CC0 conforme licença; não implica detecção de malware, não usar logotipo/alegar endosso | https://github.com/cisagov/kev-data/blob/develop/LICENSE |
| VirusTotal | avaliação no inventário | excluir da integração comercial; API pública não permite produto comercial ou substituto AV | https://docs.virustotal.com/reference/public-vs-premium-api |
| IA/ajuda offline própria | inferência e busca local | sem API paga; auditar dados de treino/arte e licenças | REQ-029/039 |
| GitHub Actions e hospedagem | infraestrutura | verificar cotas e termos; não contratar, nem pressupor hospedagem gratuita ilimitada | https://docs.github.com/en/billing |

## Gates pré-distribuição
1. Inventariar módulos efetivos de Go após geradores (go list -m -json all, go version -m do binário), bibliotecas incorporadas, executáveis externos, ferramentas build, licenças de dados/arte, fontes e dependências transitivas. Produzir SBOM versionado e THIRD_PARTY_NOTICES com textos integrais e atribuições correspondentes; incluir no instalador e pacote, comparar arquivos reais.
2. Negar promoção quando houver componente não identificado, licença desconhecida/incompatível, uso comercial condicionado sem comprovação gratuita, ou avisos ausentes. SPDX/identificação automatizada auxilia e não substitui revisão jurídica. Não atribuir licença permissiva ao código próprio por engano.
3. Preparar migração Inno→NSIS em laboratório, sem apagar instalações antigas: manter diretório, dados, quarentena, histórico e identidade coerente; reconhecer registro de desinstalação Inno legado e migrar com segurança, garantir um desinstalador válido, atualização in-place, modo silencioso, ícone, hashes, permissões e rollback. Testar 3.0.6→3.0.7, reinstalação, falha de upgrade, desinstalação e cinco ciclos. Não alterar pacote atual; considerar MSIX separadamente se compatível.
4. Fontes pagas/indefinidas: desabilitar chamadas automáticas e botões de consulta não autorizada na versão destinada a distribuição, preservar funções internas e importação apenas de material autorizado, mostrar motivo e ausência de cobertura. Não usar API de terceiros por chaves compartilhadas ou para contornar termos; não inventar substituto equivalente. Qualquer redução de cobertura deve ser medida, documentada e aprovada.
5. Atualizar matriz a cada build e revisão de termos; manter proteção, eficiência e privacidade. Gate de lançamento pendente: análise preliminar não garante ausência de infrações, dívidas ou custos. Sem aquisição ou alteração de código funcional nesta revisão.
