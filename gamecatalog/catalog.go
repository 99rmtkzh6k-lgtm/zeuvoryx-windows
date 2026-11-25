package gamecatalog
import("fmt";"net/url")
var approved=map[string]string{
"stargiraf":"https://drludos.itch.io/spacegiraf",
"simplegame":"https://marc-vittorini.itch.io/simple-game",
"minekart":"https://mhughson.itch.io/minekart-madness",
"retropia":"https://wendelscardua.itch.io/retropia",
"yaroze":"https://yaroze.dev/examples/",
"psoxide":"https://bonnie-studios.itch.io/",
"tetris":"https://play.tetris.com/games/",
"pacman":"https://doodles.google/doodle/30th-anniversary-of-pac-man/",
"asteroids":"https://games.aarp.org/games/atari-asteroids",
"centipede":"https://games.aarp.org/games/atari-centipede",
}
// Resolve accepts a catalog identifier, never a caller-supplied URL or command.
func Resolve(id string)(string,error){
 s,ok:=approved[id];if !ok{return "",fmt.Errorf("Jogo fora do catálogo")}
 u,e:=url.Parse(s);if e!=nil||u.Scheme!="https"||u.Host==""||u.User!=nil||u.Port()!=""{return "",fmt.Errorf("Endereço inválido")}
 return s,nil
}
