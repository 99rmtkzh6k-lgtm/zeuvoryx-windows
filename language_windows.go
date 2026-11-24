//go:build windows

package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
    "syscall"
)

func zSupported(mode string) bool {
    switch mode {case "auto","pt-BR","en-US","es-ES":return true}
    return false
}
func zLocaleForID(id uintptr) string {
    switch id & 0x3ff {
    case 0x16:return "pt-BR"
    case 0x0a:return "es-ES"
    default:return "en-US"
    }
}
func zResolveLanguage(mode string) string {
    if mode!="auto"&&zSupported(mode){return mode}
    proc:=syscall.NewLazyDLL("kernel32.dll").NewProc("GetUserDefaultUILanguage")
    if e:=proc.Find();e!=nil{return "en-US"}
    id,_,_:=proc.Call()
    return zLocaleForID(id)
}
func zLanguagePath()(string,error){
    base,e:=os.UserConfigDir();if e!=nil{return "",e}
    return filepath.Join(base,"ZEUVORYX","language.txt"),nil
}
func zReadLanguage()string{
    path,e:=zLanguagePath();if e!=nil{return "auto"}
    f,e:=os.Open(path);if e!=nil{return "auto"};defer f.Close()
    b,e:=io.ReadAll(io.LimitReader(f,32));if e!=nil{return "auto"}
    mode:=strings.TrimSpace(string(b));if zSupported(mode){return mode}
    return "auto"
}
func zSaveLanguage(mode string)error{
    if !zSupported(mode){return fmt.Errorf("Invalid language / Idioma inválido")}
    path,e:=zLanguagePath();if e!=nil{return e}
    if e=os.MkdirAll(filepath.Dir(path),0700);e!=nil{return e}
    return os.WriteFile(path,[]byte(mode+"\n"),0600)
}
