# ZEUVORYX 3.0.1 Competitive — Windows (beta)

Projeto experimental de segurança para Windows, em Go. O motor combina heurísticas explicáveis e indicadores SHA-256; não é um modelo de IA treinado nem substitui o Microsoft Defender.

## Compilação e download

1. Abra **Actions → ZEUVORYX 3.0.1 Competitive - Windows Build → Run workflow** e selecione `main`. Um push em `main` também inicia a compilação.
2. Aguarde a conclusão dos testes e do build.
3. Em **Artifacts**, baixe **ZEUVORYX-3.0.1-Competitive-Windows**.
4. Extraia o pacote e use `installer/ZEUVORYX-Setup-3.0.1.exe` ou a edição portátil em `dist/`. A interface desktop utiliza Microsoft WebView2.

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
