
import os
from PIL import Image

directory = "./bookimages/"
smalldirectory = "small/"

size = 150,225

for f in sorted(os.listdir(directory)):
    print(f)
    id = f[:f.find(".")]
    outfile = directory + smalldirectory + id + ".jpg"
    try:
        im = Image.open(directory + f)
        im.thumbnail(size, Image.ANTIALIAS)
        im.save(outfile, "JPEG")
    except IOError:
        print("Cannot create thumbnail for '%s'" % f)
