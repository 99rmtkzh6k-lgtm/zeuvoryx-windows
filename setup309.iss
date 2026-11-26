[Setup]
AppId=ZEUVORYX.Desktop.Lab303
AppName=ZEUVORYX 3.0.10 LAB
AppPublisher=Leonardo A Botelho
VersionInfoCompany=Leonardo A Botelho
AppVersion=3.0.10
VersionInfoVersion=3.0.10.0
UninstallDisplayName=ZEUVORYX 3.0.10 LAB
DefaultDirName={localappdata}\Programs\ZEUVORYX-Lab303
DefaultGroupName=ZEUVORYX LAB
DisableProgramGroupPage=yes
PrivilegesRequired=lowest
AppMutex=Local\ZEUVORYX.Desktop.Lab303
CloseApplications=yes
RestartApplications=no
UsePreviousAppDir=yes
ArchitecturesAllowed=x64compatible
OutputDir=installer
OutputBaseFilename=ZEUVORYX-LAB-Setup-3.0.10
SetupIconFile=dist\ZEUVORYX.ico
UninstallDisplayIcon={app}\ZEUVORYX.exe
Compression=lzma2
SolidCompression=yes
WizardStyle=modern
[Languages]
Name: "brazilianportuguese"; MessagesFile: "compiler:Languages\BrazilianPortuguese.isl"
[Tasks]
Name: "desktopicon"; Description: "Criar atalho na área de trabalho"
Name: "autostart"; Description: "Abrir ZEUVORYX ao entrar no Windows (sem varredura automática)"
[Files]
Source: "ops\Observe-Zeuvoryx.ps1"; DestDir: "{app}\ops"; Flags: ignoreversion
Source: "dist\ZEUVORYX.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "dist\signatures.sha256"; DestDir: "{app}"; Flags: ignoreversion
Source: "dist\ZEUVORYX.ico"; DestDir: "{app}"; Flags: ignoreversion
Source: "dist\ZEUVORYX.png"; DestDir: "{app}"; Flags: ignoreversion
[InstallDelete]
Type: files; Name: "{userstartup}\ZEUVORYX LAB.lnk"
[Icons]
Name: "{userstartup}\ZEUVORYX LAB"; Filename: "{app}\ZEUVORYX.exe"; WorkingDir: "{app}"; Tasks: autostart
Name: "{autoprograms}\ZEUVORYX LAB"; Filename: "{app}\ZEUVORYX.exe"; WorkingDir: "{app}"
Name: "{autodesktop}\ZEUVORYX LAB"; Filename: "{app}\ZEUVORYX.exe"; WorkingDir: "{app}"; Tasks: desktopicon
[Run]
Filename: "{app}\ZEUVORYX.exe"; Description: "Abrir ZEUVORYX"; Flags: nowait postinstall skipifsilent unchecked
Filename: "{app}\ZEUVORYX.exe"; Parameters: "-updated-version={param:ZEUVORYXTARGET|}"; Flags: nowait skipifnotsilent; Check: IsSignedUpdate
[Code]
function IsSignedUpdate: Boolean;
begin
 Result := ExpandConstant('{param:ZEUVORYXUPDATE|0}') = '1';
end;
