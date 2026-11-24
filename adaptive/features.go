package adaptive

import "strings"

// FromReasons uses the scanner's existing evidence, not filenames, paths,
// imported IOC descriptions or private contents. It must not modify verdicts.
func FromReasons(reasons []string) Vector {
    var x Vector
    for _,r:=range reasons {
        switch r {
        case "Extensão dupla simulando documento","Assinatura PE em arquivo com extensão não executável": x[0]=1
        case "Tipo de script executável merece cautela": x[1]=1
        case "PowerShell com comando codificado","Decodificação Base64 em script","Execução dinâmica de comando","Execução dinâmica abreviada","Download e execução em script","Automação de shell","Tentativa de ocultar janela","Execução suspeita via rundll32","Entropia elevada em script": x[2]+=0.25
        case "Tentativa de criar exclusão de antivírus","Tentativa de desabilitar proteção","Tentativa de apagar cópias de sombra","Tentativa de apagar catálogo de backup","Alteração de configuração de boot": x[3]+=0.5
        case "PE: combinação de APIs de injeção remota (indicador, não prova)","PE: verificações múltiplas contra depuração": x[4]+=0.5
        default:
            // No untrusted text may become a training label.
            if strings.HasPrefix(r,"Correspondência SHA-256: "){continue}
        }
    }
    for i:=range x {if x[i]>1{x[i]=1}}
    return x
}
