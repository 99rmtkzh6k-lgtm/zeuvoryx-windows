package portable

import (
 "fmt"
 "strings"
 "unicode"
)

type FraudAssessment struct {
 Signals []string `json:"signals"`
 Assessment string `json:"assessment"`
 Actions []string `json:"actions"`
}
type fraudRule struct { label string; terms []string }
var fraudRules = []fraudRule{
 {"Solicitação de segredo de carteira ou código de acesso", []string{"seed phrase", "frase de recuperacao", "chave privada", "private key", "codigo de verificacao", "codigo sms", "one time password", "codigo 2fa"}},
 {"Promessa de lucro sem risco", []string{"lucro garantido", "retorno garantido", "ganho garantido", "guaranteed profit", "guaranteed returns", "sem risco", "risk free"}},
 {"Cobrança antecipada para recuperar ou liberar dinheiro", []string{"taxa de desbloqueio", "taxa de recuperacao", "taxa para sacar", "unlock fee", "recovery fee", "deposito para liberar"}},
 {"Pressão ou ameaça de bloqueio", []string{"conta sera bloqueada", "account suspended", "urgente confirme", "acao imediata", "verify immediately"}},
 {"Solicitação de acesso remoto", []string{"instale anydesk", "instale teamviewer", "install anydesk", "install teamviewer", "acesso remoto ao seu computador"}},
 {"Solicitação de aprovação ou assinatura de carteira", []string{"aprovar transacao", "assine a transacao", "approve transaction", "sign transaction", "conecte sua carteira", "connect your wallet"}},
 {"Solicitação de documentos ou selfie", []string{"envie seu cpf", "envie uma selfie", "foto do seu documento", "send your passport", "send your id"}},
}
func fraudNormalize(s string) string {
 // Handles common Latin accents and whitespace; not a comprehensive confusable detector.
 var b strings.Builder; b.Grow(len(s)); space:=false
 for _,r:=range s { r=unicode.ToLower(r); switch r {case 'á','à','ã','â','ä':r='a';case 'é','ê','è','ë':r='e';case 'í','ì','î','ï':r='i';case 'ó','ò','õ','ô','ö':r='o';case 'ú','ù','û','ü':r='u';case 'ç':r='c'}; if unicode.Is(unicode.Mn,r){continue}; if unicode.IsSpace(r)||unicode.IsPunct(r)||unicode.IsControl(r)||unicode.Is(unicode.Cf,r){if !space{b.WriteByte(' ');space=true};continue}; b.WriteRune(r);space=false }
 return strings.TrimSpace(b.String())
}
func assessFraud(raw string)(FraudAssessment,error){
 r:=FraudAssessment{Signals:[]string{},Actions:[]string{},Assessment:"Inconclusivo: nenhum padrão reconhecido; não comprova legitimidade"}
 if len(raw)>16384||strings.TrimSpace(raw)=="" {return r,fmt.Errorf("Informe uma mensagem de até 16 KiB, sem senhas ou dados pessoais")}
 s:=fraudNormalize(raw)
 for _,rule:=range fraudRules{for _,term:=range rule.terms{if strings.Contains(s,term){r.Signals=append(r.Signals,rule.label);break}}}
 if len(r.Signals)>0{r.Assessment="Atenção: padrões de possível fraude; confirme por canal independente";r.Actions=[]string{"Não envie dinheiro, códigos, documentos ou segredos de carteira.","Não abra links nem instale programas indicados na mensagem.","Confirme pelo aplicativo oficial ou contato conhecido, em dispositivo confiável."}}
 return r,nil
}

// No user-supplied commands, secrets or destinations are accepted by the response planner.
func incidentPlan(kind string)([]string,error){
 switch kind {
 case "malware":return []string{"Se houver atividade suspeita em andamento, desconecte o dispositivo da rede manualmente; em acesso remoto você poderá perder a conexão.","Mantenha as proteções nativas do macOS e procure orientação oficial da Apple.","Examine arquivos selecionados e procure suporte para remediação.","Confira o resultado e execute outra análise. Se persistir, procure suporte especializado.","Preserve evidências e restaure somente backups confiáveis após a contenção."},nil
 case "crypto":return []string{"Interrompa pagamentos e não pague supostas taxas de recuperação.","Em dispositivo limpo, contate a corretora pelo aplicativo oficial e solicite bloqueio de saques ou da conta, se disponível.","Encerre sessões, revogue chaves API e permissões suspeitas por meios oficiais. Altere senhas e configure MFA resistente a phishing.","Se a frase de recuperação ou chave privada foi exposta, procure orientação confiável para transferir os ativos remanescentes a uma nova carteira segura. Nunca compartilhe a frase.","Preserve identificadores de transações e registre ocorrência. A recuperação de fundos não é garantida."},nil
 case "identity":return []string{"Em dispositivo limpo, proteja primeiro o e-mail principal: altere senha, encerre sessões e revise recuperação e MFA.","Contate bancos e serviços pelos canais oficiais, conteste operações e peça bloqueio conforme necessário.","Consulte contas e relacionamentos no Registrato pelos canais oficiais do Banco Central.","Preserve evidências, registre ocorrência e monitore novas contas e transações."},nil
 case "ransomware":return []string{"Desconecte o dispositivo afetado da rede para conter a propagação; não conecte backups.","Preserve evidências e procure suporte especializado, sobretudo em redes empresariais.","Não pague resgate como garantia de recuperação. Trate ameaças detectadas e corrija a causa antes de restaurar.","Restaure backups offline confiáveis após validação do ambiente; altere credenciais comprometidas em dispositivo limpo."},nil
 default:return nil,fmt.Errorf("Tipo de incidente não permitido")
 }
}
