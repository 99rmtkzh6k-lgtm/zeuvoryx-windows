import unittest
import numpy as np
from ember_benchmark import metrics, threshold_for_fpr, benchmark

class BenchmarkTests(unittest.TestCase):
    def test_metrics(self):
        m=metrics([0,0,1,1],[.1,.9,.8,.2],.5)
        self.assertEqual((m["tp"],m["tn"],m["fp"],m["fn"]),(1,1,1,1))
        self.assertIsNone(metrics([1,1],[.1,.9],.5)["fpr"])
        with self.assertRaises(ValueError): metrics([1,2],[.1,.2],.5)
    def test_threshold_ties(self):
        y=np.array([0,0,0,0,1]); s=np.array([.1,.2,.9,.9,.95])
        t=threshold_for_fpr(y,s,.25)
        self.assertLessEqual(metrics(y,s,t)["fpr"],.25)
    def test_synthetic_smoke(self):
        rng=np.random.default_rng(2026)
        def data(n):
            y=np.arange(n)%2
            return rng.normal(size=(n,8))+y[:,None]*.6,y
        train=data(500); test=data(200); challenge=(rng.normal(size=(40,8))+.6,np.ones(40,dtype=int))
        _,r=benchmark(train,test,challenge)
        self.assertEqual(r["train_count"]+r["validation_count"],500)
        self.assertEqual(r["test_count"],200)
        self.assertEqual(r["challenge_count"],40)
        self.assertFalse(r["auto_deployed"])
        self.assertFalse(r["product_validated"])
        self.assertIsNone(r["challenge"]["fpr"])
        self.assertLessEqual(r["validation"]["fpr"],r["target_validation_fpr"])
if __name__=="__main__": unittest.main()
