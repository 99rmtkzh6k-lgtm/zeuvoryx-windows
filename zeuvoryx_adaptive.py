"""Attach the experimental local model after the 3.0.1 upgrade.
May raise corroborated static evidence for review; does not authorize quarantine.
"""
from pathlib import Path

def one(s, old, new):
    if s.count(old) != 1:
        raise RuntimeError("Unexpected source: " + old[:100])
    return s.replace(old, new, 1)

p = Path("engine/scanner.go")
s = p.read_text(encoding="utf-8")
s = one(s, 'import (', 'import (\n    "github.com/zeuvoryx/zeuvoryx/adaptive"')
s = one(s, 'type Finding struct {', '''type Finding struct {
    AdaptiveScore *float64 `json:"adaptive_score,omitempty"`
    AdaptiveModel string `json:"adaptive_model,omitempty"`
''')
old = 'return &Finding{Path: path, SHA256: fullHash, Risk: score, Level: tier(score), Reasons: reasons, Size: int64(len(data)), Evidence:kind,QuarantineEligible:eligible&&!strings.Contains(path," :: ")}'
# gofmt may already have normalized spacing in the generated source.
import re
matches = re.findall(r'return &Finding\{Path: path, SHA256: fullHash, Risk: score, Level: tier\(score\), Reasons: reasons, Size: int64\(len\(data\)\), Evidence:\s*kind,\s*QuarantineEligible:\s*eligible\s*&&\s*!strings\.Contains\(path, " :: "\)\}', s)
if old in s:
    matches = [old]
if len(matches) != 1:
    raise RuntimeError("Expected exactly one final Finding constructor")
# Run the actual advisory model before the alert threshold, even for clean files.
gates = re.findall(r'if\s+score\s*<\s*30\s*\{', s)
if len(gates) != 1:
    raise RuntimeError("Expected exactly one analysis threshold")
s = one(s, gates[0], '''features := adaptive.FromReasons(reasons)
    prediction, adaptiveErr := adaptive.Observe(features)
    if score < 30 && score >= 15 && adaptiveErr == nil && adaptive.ReviewAlert(features,prediction) { score = 30; reasons = append(reasons,"IA: indicadores estáticos correlacionados; revisão necessária") }
    ''' + gates[0])
s = one(s, matches[0], '''hit := &Finding{Path: path, SHA256: fullHash, Risk: score, Level: tier(score), Reasons: reasons, Size: int64(len(data)), Evidence:kind,QuarantineEligible:eligible&&!strings.Contains(path," :: ")}
    if kind == "heuristic" && adaptiveErr == nil {
        value := prediction * 100
        hit.AdaptiveScore = &value
        hit.AdaptiveModel = adaptive.Version
    }
    return hit''')
p.write_text(s, encoding="utf-8")
p = Path("desktop.html")
s = p.read_text(encoding="utf-8")
s = one(s, "if(f.trust)row.append", "if(Number.isFinite(f.adaptive_score))row.append(node('p','IA experimental: '+Math.round(f.adaptive_score)+'/100 · índice auxiliar, não é probabilidade de malware. Modelo treinado com dados sintéticos.'));if(f.trust)row.append")
p.write_text(s, encoding="utf-8")
print("Experimental ML review integration applied; quarantine unchanged")
