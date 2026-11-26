# 2048 — Gabriele Cirulli

Source: https://github.com/gabrielecirulli/2048/tree/478b6ec346e3787f589e4af751378d06ded4cbbc

License: MIT, reproduced in LICENSE.txt and packaged with the installer.

The offline bundle combines the upstream HTML, CSS and JavaScript in their original load order. Removed external links, external font imports and mobile icons. Added a restrictive CSP and embedded the result in a sandboxed iframe with an opaque origin. The upstream storage fallback keeps the current game's state in memory when persistent storage is unavailable. The desktop checks parental policy before creating the frame.

SHA-256 of game.html: d0a8b940891c03d845d5c79d939a187232a7a75741ed45b5d9b9910fb8822ed3
