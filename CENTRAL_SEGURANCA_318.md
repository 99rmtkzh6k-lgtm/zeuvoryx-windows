## 3.0.18 LAB — Central de Segurança do Windows

O candidato na branch LAB oferece ações diretas com confirmação e sessão elevada para ativação do Defender, atualização, tratamento de ameaças e auditoria para rede, pastas e uma regra ASR. Reconsulta as preferências após alteração. A classificação separa risco observado, atenção, indisponibilidade e informação. RDP desativado tem classificação segura. Serviços, permissões, políticas SmartScreen e antivírus registrados informam a revisão de postura. [Documentação](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/blob/lab-integracao-p0-20261126/CENTRAL_SEGURANCA_318.md)

O [CI 45853385995](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45853385995) aprovou Windows e Ubuntu, abertura WebView2, migração e os cinco ciclos do ClamAV embarcado no aplicativo instalado. A classificação da Central passou cinco vezes no runner Windows, com [evidência JSON](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45853385995/artifacts/15133844653). Cinco ciclos adicionais com cmdlets simulados cobriram confirmação, idempotência, preservação de bloqueio e recusa sem comprovação de estado. Mutações reais, cobertura SmartScreen em execução, avaliação CVE abrangente e proteção própria residente seguem abertas.

[Instalador Windows x64 para piloto](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45853385995/artifacts/15136380778), build **20**, tentativa 1, commit `99050aa401ef0fc5ca9aabc2eeb82a03d452b40b`. ZIP e EXE verificados, com manifesto e CRC íntegro. SHA-256 do ZIP `c0610a2edb88c8a2b77122ce9e2c3111e259407d74be89f376be9c0dd182ccea`; SHA-256 do EXE `c29a1307fe54105b9cc2ba655d2349a44ba790af9a9bee635e99a89e6f09a495`. Lab experimental; instalador sem Authenticode. O código funcional permanece na branch LAB.

# Central de Segurança do Windows — ZEUVORYX 3.0.18 LAB

Primeira integração operacional da determinação complementar. Branch `lab-integracao-p0-20261126`, commit funcional `99050aa`. [CI associado](https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/actions/runs/45853385995)

| Escopo | Implementação e evidência disponível |
| --- | --- |
| Defender, proteção em tempo real e atualização | Usa Get-MpComputerStatus. O botão de ativação exige Defender ativo em modo Normal e verifica novamente o modo e a proteção em tempo real após a alteração. O botão de atualização consulta versão e idade das assinaturas após o comando. |
| Ransomware, rede, ASR e MAPS | Consulta preferências e oferece auditoria de rede, acesso controlado a pastas e uma regra ASR fixa de processos filhos do Office. Os comandos de auditoria preservam modos de bloqueio existentes e confirmam a preferência após executar. O painel distingue configuração MAPS de conexão comprovada à nuvem. |
| Secure Boot, BitLocker e RDP | Consultas somente leitura. RDP desativado tem sinal verde; se habilitado, recomenda verificar se é necessário. A consulta RDP mostra configuração, sem inferir sessões ou exposição externa. |
| Firewall e proteção contra adulteração | Consulta os perfis ativos do Firewall e o estado de adulteração informado pelo Defender. A interface distingue resultado observado de falha de consulta. |
| Navegação e SmartScreen | Consulta políticas de máquina do Windows e do Edge. Se a política não for encontrada, o estado efetivo continua não verificado. Filtros de outros navegadores e bloqueio pelo ZEUVORYX requerem validação própria. |
| Vulnerabilidades, serviços, permissões e coexistência | Consulta atualizações pendentes com Windows Update, serviços esperados, elevação e antivírus registrados. Sinaliza modo passivo para revisão de coexistência. Inventário de atualizações e coexistência são diagnósticos parciais, sem resultado conclusivo sobre CVEs ou conflitos entre motores. |
| Operações e reavaliação | A Central aciona diretamente o controlador nativo existente. Operações de alteração exigem confirmação na interface e sessão elevada. O controlador limita as operações a scripts fixos. O painel reavalia após receber o resultado. Firmware, Secure Boot e BitLocker são somente leitura. |
| Mecanismos próprios | ClamAV, análise de links e sinais comportamentais anteriores continuam integrados. Bloqueio residente, antirransomware próprio e navegação preventiva permanecem pendentes de implementação e validação. |

Cores: verde para configuração segura observada, amarelo para atenção, vermelho para risco constatado, cinza para indisponibilidade ou resultado não verificado, azul para informações. O registro de política SmartScreen descreve a configuração, sem inferência de bloqueio em execução.

## Testes

- Cinco ciclos de classificação e presença dos controles gerados, 16 classificações por ciclo. [JSON local](qa/security-center/3.0.18-ui-regression-local.json)
- Cinco ciclos locais executando os quatro novos scripts com cmdlets simulados em PowerShell 7.7. Os testes verificaram confirmação, idempotência, preservação de bloqueio, recusa em modo passivo e falha de confirmação quando a política impede a alteração. [JSON de contratos](qa/security-center/3.0.18-defender-contracts.json)
- As mutações reais ainda exigem cinco ciclos em Windows controlado com elevação, Defender normal/passivo, políticas gerenciadas, tamper e retorno de preferências. Validar também a verificação da política SmartScreen, as edições sem SecurityCenter2 e a atualização de assinaturas. Os testes de contrato não medem proteção em execução.

Fontes oficiais: [Get-MpComputerStatus](https://learn.microsoft.com/powershell/module/defender/get-mpcomputerstatus), [Set-MpPreference](https://learn.microsoft.com/powershell/module/defender/set-mppreference), [regras ASR](https://learn.microsoft.com/defender-endpoint/attack-surface-reduction-rules-reference) e [políticas SmartScreen no Edge](https://learn.microsoft.com/deployedge/microsoft-edge-policies/smartscreenenabled).
