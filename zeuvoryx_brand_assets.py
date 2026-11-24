"""Derive lightweight artwork from the owner's approved poster; fail closed on wrong assets."""
from pathlib import Path
import hashlib
from PIL import Image

p = Path("zeuvoryx_brand.jpg")
expected = "c0ae621b1784260b67f8f8953d9ad96ac6cb2a5c9b96cd66f3125f22ef9de754"
if not p.is_file():
    raise SystemExit("Missing zeuvoryx_brand.jpg: upload the approved artwork to repository root")
if p.stat().st_size > 500_000 or hashlib.sha256(p.read_bytes()).hexdigest() != expected:
    raise SystemExit("Approved artwork hash/size mismatch")
with Image.open(p) as source:
    if source.format != "JPEG" or source.size != (1122,1402):
        raise SystemExit("Unexpected artwork dimensions or format")
    im = source.convert("RGB")
def crop_fraction(box):
    return im.crop(tuple(round(v * dim) for v,dim in zip(box,(im.width,im.height,im.width,im.height))))
def save(image,name,limit):
    image.thumbnail(limit, Image.Resampling.LANCZOS)
    image.save(name,quality=82,optimize=True)
    if Path(name).stat().st_size > 250_000:
        raise SystemExit("Derived artwork too large: "+name)
save(crop_fraction((0.10,0.01,0.95,0.43)),"cover-small.jpg",(900,420))
save(im.copy(),"panel-small.jpg",(720,900))
save(crop_fraction((0.34,0.21,0.66,0.50)),"brand-shield.jpg",(256,256))
print("Approved artwork verified; lightweight hero, full poster and shield prepared")
