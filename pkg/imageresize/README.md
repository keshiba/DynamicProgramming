### Content Aware Image Resizing
Content Aware Image Resizing is when we change the size of an image
taking into account what's in the image and preserving critial details
during the resize process so that the result looks natural.

To identify prominent details of an image, we would start off by
calculating the Energy of each pixel in it. When we've completed
identifying the un-interesting parts of an image, we can perform
what is called Seam carving to remove the pixels in that region

### Energy Calculation

Energy at a certain pixel in an image is a numerical value that
captures how much the image is changing around that pixel.
In this way, we can say that areas of image with high
energy values are interesting and areas of image with low
enery values are uninteresting
To compute the energy of a pixel, we have to look around the
pixel both horizontally and vertically

Horizontal delta
del(x)^2 = del(Red2, Red1)^2 +
            del(Green2, Green1)^2 +
            del(Blue2, Blue1)^2
Vertical delta is similar del(y)^2

energy(x, y) = del(x)^2 + del(y)^2

Areas with more rapidly changing colors will give higher values of
energy

### Seam Carving

In Seam Carving, we identify a string of pixels in the un-interesting
region of the image that goes down the height of the image such that the
pixels are connected and not necessarily in a straight line.

This connected set of pixels is called a Seam.

Once a seam is identified, it can be removed from the image and the
resultant image will be 1 pixel shorter in the horizontal axis.

This process can be repeated on the image until the desired width
is reached

[Learn More about Seam Carving](seamcarving/)