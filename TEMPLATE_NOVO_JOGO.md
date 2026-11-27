# GameWindow obrigatório — candidato Build 51 LAB
- Cadastre metadados com ID único, documento HTML local e licença. O launcher registra automaticamente cada entrada via win.registerGame(game, doc); abra somente via win.open(id), após allowLocalGame315(id).
- Não crie iframe fora de lab/gamewindow.js. Não adicione allow-same-origin, popups, top-navigation ou acesso às funções nativas. Preserve CSP offline e controle familiar.
- O catálogo original e o catálogo integrado são distintos. Atualize a lista nativa de IDs e a validação do catálogo correspondente; não remova a avaliação familiar. Nunca habilite jogo desconhecido.
- Cada jogo deve documentar comandos e fornecer teste de mudança de estado por ação real. Jogos de personagem: seta direita deve mudar a posição; puzzles: estado do tabuleiro; jogos de botões: resultado de ação via teclado. Nenhuma aprovação por simples abertura.
- Teste GameWindow: foco, maximizar/restaurar, zoom próprio, fullscreen e saída, minimizar/restaurar sem recarga, fechar, bloqueio familiar, troca de jogo, 30 s de interação, DPI/resoluções. Exija evidência física em Windows por jogo; execução Edge em CI não equivale ao WebView2 instalado.
- O build deve falhar com "Jogo sem GameWindow - correção obrigatória" se um launcher não registrar/abrir pelo componente ou criar iframe próprio.
