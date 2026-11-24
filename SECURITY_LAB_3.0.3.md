# ZEUVORYX 3.0.3 — fundação experimental
Implementação isolada da 3.0.2. Sem distribuição pública.
- releaseguard: autenticação Ed25519 de manifesto canônico, sequência monotônica fornecida pelo chamador, expiração, canal, plataforma, URL oficial e verificação SHA-256/limite de tamanho. Chave real do proprietário, armazenamento protegido de sequência, download seguro, integração e instalação atômica ainda pendentes.
- impact: validação de snapshot global assinado e compacto, contadores uint64, consistência e média das avaliações. Não há backend, telemetria, credencial de produção, interface ou dados reais; nenhuma estatística inventada.
- remediation: política pura fail-closed. Não encerra processos, altera registro, elimina arquivos ou oferece proteção residente. Integração e testes de execução pendentes.
- CI: testes específicos de fundação em Windows/Linux. Não representam teste do instalador, proteção antimalware, ou certificação.
- Chaves geradas nos testes são efêmeras, sem validade de produção.
- Pendências do memorando permanecem: testes 3.0.2, piloto real, motor de remediação, proteção residente, histórico e servidor auditável, titularidade de chaves, revisão de permissões, assinatura de binários, distribuição e monetização.

## Revisão adicional
Vinculação estrita de URL ao tag da versão, rejeição de streams inválidos e snapshots antigos, exclusão de artefatos de teste como condição para métricas, fuzzing e dez repetições embaralhadas de testes. Preparado empacotamento restrito do C2PA (executável e arquivos legais); ainda precisa de build integral. Nenhum módulo novo está integrado ao instalador.
