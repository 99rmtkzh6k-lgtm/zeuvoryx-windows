package globalintel

import (
 "context"
 "embed"
 "os"
 "path/filepath"
 "time"
)

//go:embed seed/*
var initial embed.FS

// Bundled accepts only a complete, current pair, and fails closed otherwise.
func Bundled(now time.Time) Drop {
 a,e:=initial.ReadFile("seed/drop_v4.json");if e!=nil{return Drop{}}
 b,e:=initial.ReadFile("seed/drop_v6.json");if e!=nil{return Drop{}}
 v4,e:=parseDrop(a,false,now);if e!=nil{return Drop{}}
 v6,e:=parseDrop(b,true,now);if e!=nil{return Drop{}}
 return Drop{V4:v4,V6:v6}
}

// CollectSeed is reserved for manual lab data collection, not normal app startup.
// Retains original feed bytes (including copyright and timestamps).
func CollectSeed(ctx context.Context,dir string)(Status,error){
 c:=client();now:=time.Now()
 a,e:=fetch(ctx,c,drop4);if e!=nil{return Status{},e}
 v4,e:=parseDrop(a,false,now);if e!=nil{return Status{},e}
 b,e:=fetch(ctx,c,drop6);if e!=nil{return Status{},e}
 v6,e:=parseDrop(b,true,now);if e!=nil{return Status{},e}
 if e=os.Mkdir(dir,0700);e!=nil{return Status{},e}
 ok:=false;defer func(){if !ok{os.Remove(filepath.Join(dir,"drop_v4.json"));os.Remove(filepath.Join(dir,"drop_v6.json"));os.Remove(dir)}}()
 if e=os.WriteFile(filepath.Join(dir,"drop_v4.json"),a,0600);e!=nil{return Status{},e}
 if e=os.WriteFile(filepath.Join(dir,"drop_v6.json"),b,0600);e!=nil{return Status{},e}
 ok=true
 return (Drop{V4:v4,V6:v6}).Status(),nil
}
