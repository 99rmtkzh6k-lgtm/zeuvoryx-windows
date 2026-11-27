# Build 50 — 2048 Win32 standalone LAB (candidate)

A separate Windows executable can be built with:

`go build -o ZEUVORYX-2048-LAB.exe ./lab/arcade2048`

The host embeds the existing MIT-licensed game.html, uses WebView2, and creates a separate top-level Win32 window at 1024x640 with 900x600 minimum. The standard Windows caption buttons are native. JS bindings z2048Minimize, z2048Maximize, z2048Restore, z2048Close dispatch WM_SYSCOMMAND on the standalone HWND. SetProcessDPIAware is called before window creation. The host has no references to the antivirus State and exits separately.

**NOT VALIDATED:** No Windows compilation, screenshot, physical keyboard/touch test, or Win32 lifecycle test has been run. This source does not yet integrate a launch button into the main app or workflow. Do not call it installed or approved. Keep main, Releases, Build 49 and retro21 untouched.
