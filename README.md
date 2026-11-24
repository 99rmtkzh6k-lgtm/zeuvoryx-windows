# ZEUVORYX 3.0.2 Competitive — Windows (beta)

Projeto experimental de segurança para Windows, em Go. O motor combina heurísticas explicáveis e indicadores SHA-256; não substitui o Microsoft Defender. A IA experimental usa regressão logística treinada com vetores sintéticos; ela não está validada para identificar malware real.

## Compilação e download

1. Abra **Actions → ZEUVORYX 3.0.2 Competitive - Windows Build → Run workflow** e selecione `main`. Um push em `main` também inicia a compilação.
2. Aguarde a conclusão dos testes e do build.
3. Em **Artifacts**, baixe **ZEUVORYX-3.0.2-Competitive-Windows**.
4. Extraia o pacote e use `installer/ZEUVORYX-Setup-3.0.2.exe` ou a edição portátil em `dist/`. A interface desktop utiliza Microsoft WebView2.

O binário não é assinado digitalmente. Verifique a procedência e os hashes; teste em Windows Sandbox ou VM com o Defender ativo antes de usar em uma máquina real.

## Recursos experimentais

- Varredura rápida e completa, análise heurística, hashes SHA-256, teste EICAR e inspeção limitada de arquivos ZIP e PE.
- Fila de varredura paralela limitada, cancelamento, indicadores de cobertura e desempenho e filtros de alertas.
- Quarentena reversível mediante ação explícita e critérios de elegibilidade.
- Interface desktop, diagnóstico local voluntário com exportação JSON e integração com recursos de segurança do Windows.
- Atualização cumulativa do motor e interface em `zeuvoryx_upgrade.py`, aplicada pelo workflow durante a compilação.

## Desenvolvimento e validação

O workflow `.github/workflows/windows-build.yml` organiza e gera arquivos de código, aplica o script de atualização, executa verificações Go e testes repetidos, compila o executável Windows x64 e gera o instalador. Consulte `LEIA-ME.txt` para instruções e `comparar_relatorios.py` para comparar relatórios equivalentes.

Os testes automatizados e avaliações sintéticas não comprovam eficácia contra malware real. O aplicativo completo ainda exige testes em Windows, auditoria e validação independente. Não há garantia de detecção universal, proteção residente própria por driver, VPN ou certificação laboratorial.

## Licença e identidade visual

Código experimental para avaliação e adaptação; consulte as licenças das dependências antes de redistribuir. As imagens conceituais do projeto não comprovam funcionalidades técnicas.

## Novidades da versão 3.0.2 (beta)

- Português brasileiro, inglês e espanhol: interface principal e mensagens selecionadas, detecção do idioma de exibição do Windows, escolha manual persistente e inglês como fallback. Mensagens técnicas externas e partes dos relatórios podem permanecer no idioma original.
- Adaptive Intelligence Engine 0.1: modelo de regressão logística local treinado em exemplos sintéticos, sem dependências de IA externas. A pontuação é experimental e aparece apenas em alertas heurísticos existentes. Não altera o veredito, não libera nem isola arquivos e não é uma probabilidade de malware.
- Treinamento supervisionado limitado e testável em código. Não existe aprendizado contínuo em produção, coleta automática de dados ou atualização automática de modelo.
- Created & Engineered by Leonardo A Botelho: autoria na interface e nos metadados do instalador. Não equivale à assinatura criptográfica Authenticode.

O build e a validação em Windows Sandbox devem ser confirmados antes de distribuição.
