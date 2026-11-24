// Package releaseguard verifies owner-authorized update metadata and package bytes.
// It does not download, execute, install, or persist rollback state.
package releaseguard

import (
 "bytes"
 "crypto/ed25519"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "io"
 "net/url"
 "regexp"
 "strings"
 "time"
)

const Domain = "ZEUVORYX/RELEASE/v1\x00"
const MaxManifest = 4096
const MaxPackage = 512 * 1024 * 1024
var versionPattern = regexp.MustCompile(`^(0|[1-9][0-9]{0,8})\.(0|[1-9][0-9]{0,8})\.(0|[1-9][0-9]{0,8})$`)
var filenamePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]{0,127}\.(zip|exe|msi)$`)
const releasePrefix = "/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/"

type Manifest struct {
 Sequence uint64 `json:"sequence"`
 Version string `json:"version"`
 Channel string `json:"channel"`
 Platform string `json:"platform"`
 URL string `json:"url"`
 SHA256 string `json:"sha256"`
 Size int64 `json:"size"`
 Expires int64 `json:"expires"`
}

func Verify(raw, signature []byte, trusted ed25519.PublicKey, installed uint64, channel string, now time.Time) (Manifest,error) {
 var m Manifest
 fail := func(s string)(Manifest,error){return Manifest{},errors.New(s)}
 if len(raw)==0 || len(raw)>MaxManifest || len(trusted)!=ed25519.PublicKeySize || len(signature)!=ed25519.SignatureSize {return fail("invalid envelope")}
 if !ed25519.Verify(trusted,append([]byte(Domain),raw...),signature){return fail("unauthorized manifest")}
 if err:=json.Unmarshal(raw,&m);err!=nil{return fail("invalid manifest")}
 canonical,err:=json.Marshal(m)
 if err!=nil || !bytes.Equal(canonical,raw){return fail("noncanonical or unexpected manifest")}
 if m.Sequence==0 || m.Sequence<=installed || !versionPattern.MatchString(m.Version) || (channel!="stable" && channel!="beta") || m.Channel!=channel || m.Platform!="windows-amd64" {return fail("version, channel or platform rejected")}
 if m.Expires<=now.Unix() || m.Expires>now.Add(30*24*time.Hour).Unix(){return fail("invalid expiration")}
 if m.Size<=0 || m.Size>MaxPackage{return fail("invalid package size")}
 digest,err:=hex.DecodeString(m.SHA256)
 if err!=nil || len(digest)!=sha256.Size || m.SHA256!=strings.ToLower(m.SHA256){return fail("invalid digest")}
 u,err:=url.Parse(m.URL)
 prefix:=releasePrefix+"v"+m.Version+"/"
 if err!=nil || u.Scheme!="https" || u.Host!="github.com" || u.User!=nil || u.RawQuery!="" || u.ForceQuery || u.Fragment!="" || u.RawPath!="" || u.String()!=m.URL || !strings.HasPrefix(u.Path,prefix) || !filenamePattern.MatchString(strings.TrimPrefix(u.Path,prefix)) || strings.Contains(u.Path,"..") {return fail("unapproved package URL")}
 return m,nil
}

// VerifyPackage consumes the complete stream, with a hard byte limit.
// The caller must install only the same verified bytes, never reopen an untrusted path.
func VerifyPackage(r io.Reader,m Manifest) error {
 if r==nil || m.Size<=0 || m.Size>MaxPackage{return errors.New("invalid stream or size")}
 h:=sha256.New()
 n,err:=io.Copy(h,io.LimitReader(r,m.Size+1))
 if err!=nil{return err}
 if n!=m.Size || hex.EncodeToString(h.Sum(nil))!=m.SHA256{return errors.New("package integrity check failed")}
 return nil
}
