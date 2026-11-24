//go:build windows

package main

import "testing"

func TestLanguageResolution(t *testing.T){
    for _,tc:=range []struct{id uintptr;want string}{
        {0x0416,"pt-BR"},{0x0816,"pt-BR"},{0x0c0a,"es-ES"},
        {0x080a,"es-ES"},{0x0409,"en-US"},{0x0411,"en-US"},{0,"en-US"},
    }{if got:=zLocaleForID(tc.id);got!=tc.want{t.Fatal(tc.id,got,tc.want)}}
    for _,mode:=range []string{"pt-BR","en-US","es-ES"}{if zResolveLanguage(mode)!=mode{t.Fatal(mode)}}
    if zSupported("invalid")||zSupported("../en-US"){t.Fatal("invalid locale accepted")}
}
