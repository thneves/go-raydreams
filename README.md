# Ray Dreams

### A Ray Tracer in Go, turning code and pixels into lights and dreams

### Following [Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)


### WTF is a Ray Tracer to Begin  With

It's a technique that creates realistic 3D images by simulating the physical behavior of light,
tracing individual light as they interact with objects, bounce off surfaces, and change direction
through refraction and reflection.

In computer graphics it's used to produce highly detailed images with accurate shadows, reflections, and 
lighting, enhancing realism into games, films and virtual things!

### The Go Part

The Original book uses C++ as the language of choice to make this Ray Tracer. I'm
making it in Go, as it's my primary working language at the moment, and I've been also really enjoying to work with.
The name of the book says "In One Weekend". Let's see how long it takes me to do it.

### What is Going on?

Here is where I write about what is happening while I do it :)

##### What is PPM Image Format ??

The very first chapter talks about the PPM Image Format, because we need a way to see an image. However, I never heard of this format before. It stands for "Portable Pixel Map", it's simple, uncompressed raster image format designed for easy portability and manipulation across plataforms. It's part of the [Netbpm](https://en.wikipedia.org/wiki/Netpbm) Toolkit.

Why we are using it?

- For it's simplicity! no need for graphics library or special code to write it.
- We focus on the Ray Tracer instead of image formats
- Can be converted to PNG or whatever later.
- It's basically a little piece of text that describes image.


I just made my first image using Go:

![IMG PPM](image.ppm)

Apparently, this is the "Hello, World" of computer graphics! and it's actually really cooL!





