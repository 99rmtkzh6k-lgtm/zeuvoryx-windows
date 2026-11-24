"""Offline EMBER2024 research benchmark. Never opens or executes malware.
Use official, already-vectorized train/test/challenge data. No auto-download,
network calls, product deployment or threshold adjustment on the test set.
"""
import argparse
import hashlib
import json
import platform
from pathlib import Path
import numpy as np

def metrics(y, scores, threshold):
    y=np.asarray(y).reshape(-1)
    scores=np.asarray(scores).reshape(-1)
    if len(y)!=len(scores) or not len(y) or not np.isin(y,[0,1]).all() or not np.isfinite(scores).all():
        raise ValueError("Invalid labels or scores")
    p=scores>=threshold
    tp=int(np.sum(p & (y==1))); tn=int(np.sum(~p & (y==0)))
    fp=int(np.sum(p & (y==0))); fn=int(np.sum(~p & (y==1)))
    return dict(tp=tp,tn=tn,fp=fp,fn=fn, fpr=fp/(tn+fp) if tn+fp else None,
                fnr=fn/(tp+fn) if tp+fn else None)

def threshold_for_fpr(y,scores,target):
    if not 0<target<1: raise ValueError("Invalid target")
    clean=np.sort(np.asarray(scores)[np.asarray(y)==0])
    if not len(clean): raise ValueError("No benign validation examples")
    allowed=int(len(clean)*target)
    return float(np.nextafter(clean[len(clean)-allowed-1],np.inf))

def checked(x,y):
    x=np.asarray(x); y=np.asarray(y).reshape(-1)
    if x.ndim!=2 or len(x)!=len(y) or not len(y) or not np.isin(y,[0,1]).all():
        raise ValueError("Invalid feature matrix / binary labels")
    return x,y

def benchmark(train,test,challenge,target=0.001):
    import lightgbm as lgb
    from sklearn.model_selection import train_test_split
    x,y=checked(*train); xt,yt=checked(*test); xc,yc=checked(*challenge)
    if x.shape[1]!=xt.shape[1] or x.shape[1]!=xc.shape[1]:
        raise ValueError("Incompatible feature schema")
    if len(np.unique(y))!=2 or len(np.unique(yt))!=2 or not np.all(yc==1):
        raise ValueError("Train/test must contain both classes; challenge must be malicious-only")
    i,j=train_test_split(np.arange(len(y)),test_size=.20,random_state=7129,stratify=y)
    model=lgb.LGBMClassifier(n_estimators=180,num_leaves=31,learning_rate=.05,n_jobs=2,random_state=7129,verbosity=-1,deterministic=True,force_col_wise=True)
    model.fit(x[i],y[i])
    vs=model.predict_proba(x[j])[:,1]
    t=threshold_for_fpr(y[j],vs,target)
    return model,dict(schema="zeuvoryx-independent-vector-benchmark-v1",features=int(x.shape[1]),
        train_count=len(i),validation_count=len(j),test_count=len(yt),challenge_count=len(yc),
        threshold=t,target_validation_fpr=target,validation=metrics(y[j],vs,t),
        test=metrics(yt,model.predict_proba(xt)[:,1],t),
        challenge=metrics(yc,model.predict_proba(xc)[:,1],t),
        product_validated=False,auto_deployed=False,
        split_note="Validation is stratified within official training split. Official test and challenge untouched until final scoring. Cross-split hash overlap must be independently audited.",
        python=platform.python_version(),lightgbm=lgb.__version__)

def main():
    parser=argparse.ArgumentParser()
    parser.add_argument("--data",required=True,help="Official pre-vectorized EMBER2024 directory")
    parser.add_argument("--out",required=True,help="New local report directory")
    parser.add_argument("--dataset-id",required=True,help="Reviewed dataset release/revision and file type")
    args=parser.parse_args()
    import thrember
    root=Path(args.data)
    if not root.is_dir(): raise SystemExit("Dataset missing")
    out=Path(args.out);out.mkdir(parents=True,exist_ok=False)
    model,r=benchmark(*(thrember.read_vectorized_features(str(root),subset=s) for s in ("train","test","challenge")))
    r["dataset_id"]=args.dataset_id
    r["source"]="https://github.com/FutureComputing4AI/EMBER2024"
    model.booster_.save_model(str(out/"research-model.txt"))
    r["model_sha256"]=hashlib.sha256((out/"research-model.txt").read_bytes()).hexdigest()
    (out/"report.json").write_text(json.dumps(r,indent=2,allow_nan=False),encoding="utf-8")
    print(json.dumps(r,indent=2))
if __name__=="__main__": main()
