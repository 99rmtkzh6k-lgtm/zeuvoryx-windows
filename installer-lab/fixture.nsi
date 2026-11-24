Unicode true
Name "ZEUVORYX Installer Component Test"
OutFile "zeuvoryx-installer-fixture.exe"
InstallDir "$LOCALAPPDATA\ZEUVORYX-Installer-Component-EVO030"
RequestExecutionLevel user
SetCompressor zlib
VIProductVersion "3.0.7.0"
VIAddVersionKey "ProductName" "ZEUVORYX installer component test"
VIAddVersionKey "FileVersion" "3.0.7.0"
VIAddVersionKey "ProductVersion" "3.0.7.0"
VIAddVersionKey "FileDescription" "LAB fixture - not a security product"
VIAddVersionKey "LegalCopyright" "ZEUVORYX project"
Section
  StrCpy $INSTDIR "$LOCALAPPDATA\ZEUVORYX-Installer-Component-EVO030"
  SetOutPath "$INSTDIR"
  File "payload.txt"
  WriteUninstaller "$INSTDIR\Uninstall.exe"
SectionEnd
Section "Uninstall"
  StrCpy $INSTDIR "$LOCALAPPDATA\ZEUVORYX-Installer-Component-EVO030"
  Delete "$INSTDIR\payload.txt"
  Delete "$INSTDIR\Uninstall.exe"
  RMDir "$INSTDIR"
SectionEnd
