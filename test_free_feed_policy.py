import unittest
from pathlib import Path
from prepare_free_feed_policy import transform_go,transform_html
class TestFreePolicy(unittest.TestCase):
    def test_actual_workflow_source(self):
        s=Path('.github/workflows/windows-build.yml').read_text(encoding='utf-8')
        out=transform_go(s)
        self.assertNotIn('v,e:=fetchBazaar(nativeCtx,intelClientInstance,key);if e!=nil{return nil,e};return applyIntel(v,true)',out)
        self.assertIn('ZVX-LICENSE-001',out)
        self.assertIn('bind("importThreatHashes"',out)
        self.assertIn('bind("checkKnownCVE"',out)
    def test_ui(self):
        s=transform_html('<html><body><script></script></body></html>')
        self.assertIn('button.disabled=true',s)
        self.assertIn('key.disabled=true',s)
        self.assertIn('Import only authorized indicators',s)
        with self.assertRaises(ValueError):transform_html(s)
    def test_missing_anchor(self):
        with self.assertRaises(ValueError):transform_go('missing')
        with self.assertRaises(ValueError):transform_html('missing')
if __name__=='__main__':unittest.main()
