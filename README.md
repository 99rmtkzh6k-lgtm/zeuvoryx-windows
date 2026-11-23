# ZEUVORYX — Segurança Inteligente (beta)

Projeto de antivírus experimental em Go, com **análise inteligente heurística explicável**, indicadores SHA-256, detecção de teste EICAR, inspeção limitada de ZIP e quarentena reversível. O motor pontua sinais de comportamento suspeito e explica cada alerta. O código **não implementa um modelo de IA treinado**, inferência neural, proteção por driver em tempo real, assinatura digital ou atualizações automáticas. Não é substituto para Microsoft Defender.

## Executável Windows com dois cliques

1. Envie **os arquivos e diretórios extraídos** deste ZIP para a raiz de um repositório GitHub. Não envie somente o ZIP compactado.
2. Abra **Actions → ZEUVORYX - Testes e executavel Windows → Run workflow** (ou envie um commit na branch `main`).
3. Aguarde a execução terminar com sucesso; entre na execução e abra **Artifacts → ZEUVORYX-Windows-x64**.
4. Baixe e extraia; abra `ZEUVORYX.exe` com **dois cliques**. Nenhuma instalação de Go ou Python é necessária no computador do usuário.
5. O programa inicia servidor local em `127.0.0.1` e abre o navegador com a interface. A janela do navegador é a interface do aplicativo; manter o processo em execução é necessário. Caso o navegador não abra, no sistema Windows use a compilação de console para depuração.

**Atenção:** o arquivo gerado não é assinado digitalmente; o Windows SmartScreen pode exibir aviso para aplicativos desconhecidos. Não ignore alertas de segurança sem verificar o arquivo. Para distribuição pública, assine o executável e realize auditoria.

## Desenvolvimento

Requer Go 1.22+ para compilar a partir do código-fonte.

```bash
go test ./...
go vet ./...
go run ./cmd/zeuvoryx
go run ./cmd/zeuvoryx -scan /caminho/para/pasta
go build -o zeuvoryx ./cmd/zeuvoryx
```

Carregue assinaturas fornecendo `-ioc /caminho/lista.sha256` ou mantenha `signatures.sha256` próximo ao executável. Cada linha tem `SHA256 descricao`, precedida opcionalmente por `#` para comentários. A base acompanha o projeto vazia, sem indicadores falsos.

## Recursos implementados

- Varredura recursiva com pool limitado de trabalhadores; ignora links simbólicos e aceita cancelamento.
- SHA-256 calculado sobre o conteúdo completo e indicador local opcional.
- Heurística por padrões de comando, extensões duplas, scripts e entropia; risco pontuado com motivos.
- Teste EICAR; detecção em entradas pequenas de ZIP não criptografadas (não extrai o ZIP).
- Registro em JSON e painel local para acompanhar resultados.
- Quarentena somente após confirmação explícita, com metadados e opção de restauração. Não exclui arquivos automaticamente por detecção.
- Testes Go e compilação Windows automatizada via GitHub Actions.

## Limitações importantes

- As regras podem ter falsos positivos e falsos negativos. O EICAR é um teste inofensivo, não uma prova de cobertura real de malware.
- Arquivos grandes são integralmente hashados, mas apenas os primeiros 32 MiB são analisados por regras. ZIP: até 100 entradas, até 8 MiB cada, sem ZIP dentro de ZIP.
- Não detecta todas as famílias de malware, não monitora processos, memória, kernel, rede ou downloads em tempo real. Não bloqueia ransomware nem acompanha atualização em nuvem.
- O painel local usa token aleatório e escuta apenas em localhost; use somente na máquina de confiança. Não exponha a porta à rede.
- Arquivos em quarentena não são criptografados; a cópia é armazenada com extensão `.bin` e permissões restritas. Não restaure arquivos desconhecidos.
- Não foram realizados testes laboratoriais de eficácia nem existem evidências verificáveis de milhões de testes para esta versão.

## Estrutura

`cmd/zeuvoryx/main.go`: aplicativo HTTP local e CLI.
`engine/scanner.go`: scanner e regras.
`engine/quarantine.go`: quarentena.
`engine/*_test.go`: testes.
`ui/index.html`: interface HTML embutida no executável.
`.github/workflows/windows.yml`: testes e compilação automáticos.

## Licença

Código entregue para avaliação e adaptação pelo responsável do projeto; consulte as licenças das dependências antes da redistribuição. Não existem dependências de terceiros no código Go.


## Identidade visual

Imagens originais geradas para o projeto, ampliadas para **3840 × 2160 pixels (4K)** e incorporadas ao executável (assets sob `ui/assets/`). São artes conceituais; os recursos apresentados nas imagens não devem ser interpretados como funcionalidades técnicas implementadas.

## Importante sobre IA e eficácia

O nome e a identidade de produto remetem a IA. O motor de detecção nesta versão é **heurístico e baseado em assinaturas**, não um modelo de machine learning treinado e validado. Não há garantia de detecção universal, bloqueio em tempo real, firewall próprio ou substituição do Defender.
