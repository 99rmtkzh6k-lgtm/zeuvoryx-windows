// Package diagnostico provides bounded, deterministic hypotheses for technical reports.
// It never performs repairs, sends telemetry, or executes suggested commands.
package diagnostico

import (
 "encoding/json"
 "fmt"
 "sort"
 "strings"
)

type Evidence struct {
 Code string `json:"code"`
 Detail string `json:"detail,omitempty"`
}
type Suggestion struct {
 Code string `json:"code"`
 Cause string `json:"possible_cause"`
 Check string `json:"how_to_verify"`
 Action string `json:"possible_solution"`
 Confidence string `json:"confidence"`
}
type Report struct {
 SchemaVersion int `json:"schema_version"`
 Status string `json:"status"`
 Evidence []Evidence `json:"evidence"`
 Suggestions []Suggestion `json:"suggestions"`
 Prompt string `json:"prompt_for_ai"`
 Truncated bool `json:"truncated"`
}
type rule struct { code, cause, check, action string }
var rules = []rule{
 {"sharing_violation","Arquivo em uso por outro processo","Verificar código de erro e volatilidade do arquivo no log","Registrar cobertura parcial e testar retry limitado, cancelável, sem forçar desbloqueio"},
 {"sharing_violation","Arquivo temporário efêmero","Comparar existência e horário do arquivo ao início e fim da varredura","Classificar desaparecimento sem repetir varredura integral"},
 {"access_denied","Permissão ou proteção de acesso","Conferir contexto de execução e código do Windows","Registrar caminho inacessível e avaliar escopo e permissões sem desativar proteções"},
 {"clamav_exit_2","Erro de processamento do ClamAV","Examinar saída nativa, erros e avisos separadamente","Corrigir causa demonstrada, testar em LAB e manter status parcial"},
 {"max_file_size","Limite de inspeção de arquivo ou compactado","Confrontar tamanho, limite configurado e RAM disponível","Avaliar inspeção direcionada com orçamento de recursos; não elevar limites indiscriminadamente"},
 {"quick_scan_small","Escopo seletivo da varredura rápida","Comparar raízes elegíveis e inventário real com bytes por motor","Não tratar volume pequeno isolado como falha; validar cobertura por escopo"},
 {"sysmon_unavailable","Telemetria Sysmon indisponível","Confirmar estado da fonte e mecanismo alternativo","Registrar lacuna de telemetria sem presumir ausência de ameaças"},
 {"low_memory","Pressão de memória no computador","Medir RAM disponível e picos de consumo do processo","Reduzir concorrência e limitar buffers antes de ampliar escopo"},
 {"scan_crash","Falha de estabilidade ou encerramento inesperado","Conferir eventos Windows, logs finais, último arquivo e memória","Reproduzir em LAB com fixture, instrumentar encerramento e testar recuperação"},
}
const maxEvidence=32
const maxSuggestions=24
const maxDetail=160
func safe(s string) string {
 s=strings.TrimSpace(s)
 // Do not include raw paths, machine identifiers or user-supplied details in AI prompts.
 if len(s)>maxDetail {s=s[:maxDetail]}
 return strings.Map(func(r rune) rune {if r=='\n'||r=='\r'||r=='\t' {return ' '}; return r},s)
}
func Analyze(input []Evidence) Report {
 r:=Report{SchemaVersion:1,Status:"diagnostico_preliminar",Evidence:[]Evidence{},Suggestions:[]Suggestion{}}
 seen:=map[string]bool{}
 for i,e:=range input {
  if i>=maxEvidence {r.Truncated=true;break}
  code:=strings.ToLower(strings.TrimSpace(e.Code))
  if len(code)>64 {code=code[:64]}
  // Only recognized codes are included in the shareable report.
  recognized:=false
  for _,rule:=range rules {if rule.code==code {recognized=true;break}}
  if !recognized {continue}
  r.Evidence=append(r.Evidence,Evidence{Code:code})
  for _,rule:=range rules {
   if rule.code!=code {continue}
   key:=rule.code+"|"+rule.cause
   if seen[key] {continue}
   if len(r.Suggestions)>=maxSuggestions {r.Truncated=true;break}
   seen[key]=true
   r.Suggestions=append(r.Suggestions,Suggestion{Code:code,Cause:rule.cause,Check:rule.check,Action:rule.action,Confidence:"hipotese_a_verificar"})
  }
 }
 sort.SliceStable(r.Suggestions,func(i,j int)bool{return r.Suggestions[i].Code<r.Suggestions[j].Code})
 var b strings.Builder
 b.WriteString("ZEUVORYX — DIAGNÓSTICO TÉCNICO PARA IA (não executar ações automaticamente).\n")
 b.WriteString("Verifique primeiro versão, branch, SHA, código, CI, memorandos e relatório JSON final. Investigue hipóteses plausíveis antes de alterar código. Preserve proteção, privacidade e desempenho; corrija apenas causa comprovada em LAB, teste e não promova main sem autorização.\n")
 for _,s:=range r.Suggestions {fmt.Fprintf(&b,"[%s] Possível causa: %s. Verificar: %s. Alternativa: %s.\n",s.Code,s.Cause,s.Check,s.Action)}
 if len(r.Suggestions)==0 {b.WriteString("Sem regra correspondente: coletar evidência técnica adicional; não inventar causa.\n")}
 r.Prompt=b.String()
 return r
}
func (r Report) JSON() ([]byte,error) {return json.MarshalIndent(r,"","  ")}
