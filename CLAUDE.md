# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go implementation of the ray tracer from [Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html). Single binary that renders a scene to PPM on stdout. All book chapters implemented through the final cover scene (antialiasing, diffuse/metal/dielectric materials, positionable camera, defocus blur).

## Commands

- Render image: `go run . > image.ppm` (progress prints to stderr; default ~2 min on a modern laptop)
- Convert to PNG: `convert image.ppm image.png` (requires ImageMagick)
- Build: `go build .`
- Vet: `go vet ./...`
- No tests yet — `go test ./...` is a no-op.

Go version: 1.24.6 (from `go.mod`). Module path: `github.com/thneves/go-raydreams`.

## Architecture

Two packages: `main` (`main.go`) builds the world + configures a `scene.Camera`; `scene/` holds all math, primitives, materials, and the render loop.

Render pipeline:
1. `main.go` builds a `scene.HittableList` (ground + ~488 random small spheres + 3 large spheres) and a `scene.Camera`.
2. `Camera.Render(world)` prints the PPM `P3` header and iterates pixels. For each pixel it takes `SamplesPerPixel` jittered samples, traces each through `rayColor` (recursive, bounded by `MaxDepth`), averages, gamma-corrects, and writes one RGB triple per line.
3. On a hit the surface `Material.Scatter` returns attenuation + scattered ray; on a miss the sky gradient is returned.

Because stdout *is* the image, anything else printed there corrupts the PPM. Progress/logging must go to stderr (see `fmt.Fprintf(os.Stderr, ...)` in `camera.go`).

### `scene/` package layout

- `vector.go` — `Vec3` struct; `Point3 = Vec3` type alias. Free functions for arithmetic: `Add`, `Sub`, `Mul`, `MulScalar`, `DivScalar`, `Dot`, `Cross`, `Len`, `LenSquared`, `Unit`. Random helpers: `Random`, `RandomRange`, `RandomInUnitSphere`, `RandomUnitVector`, `RandomOnHemisphere`, `RandomInUnitDisk`. Shading helpers: `NearZero`, `Reflect`, `Refract`. Convention: free functions, value semantics. The `*Pointer` variants (`AddPointer`, `SubPointer`, `DivScalarPointer`) are duplicates kept for `Point3` semantic separation — they are not actually pointer-based.
- `interval.go` — `Interval{Min, Max}` with `Contains`, `Surrounds`, `Clamp`, `Size`. Used everywhere a `[tMin, tMax]` range is needed.
- `ray.go` — `Ray{origin, direction}` (unexported fields, accessed package-internally). `NewRay`, exported getters `Origin()`/`Direction()`, and `At(t)` implements `P(t) = origin + t·dir`. No shading logic — that lives on `Camera`.
- `hittable.go` — `HitRecord{P, Normal, Mat, T, FrontFace}` + `SetFaceNormal(r, outwardNormal)` to orient the normal against the incoming ray. `Hittable` interface uses `Hit(r Ray, rayT Interval)`. `HittableList` is the world container.
- `sphere.go` — `Sphere{Center, Radius, Mat}` and `NewSphere`. `Hit` uses the simplified quadratic with `oc = Center - Origin` and `h = Dot(d, oc)` (factor of 2 absorbed). Discriminant is `h*h - a*c`; roots are `(h ± √disc)/a`. Don't "fix" this back to the textbook form without matching the rest of the code.
- `material.go` — `Material` interface: `Scatter(rIn, rec) (bool, Color, Ray)`. Implementations: `Lambertian` (matte), `Metal` (reflective + fuzz), `Dielectric` (refractive with Schlick approximation for Fresnel).
- `camera.go` — `Camera` struct with public config (`AspectRatio`, `ImageWidth`, `SamplesPerPixel`, `MaxDepth`, `VFov`, `LookFrom`, `LookAt`, `VUp`, `DefocusAngle`, `FocusDist`) and derived unexported fields. `Render(world Hittable)` drives the pipeline; `rayColor` is the recursive shader; `getRay` does pixel jitter + defocus disk sampling.
- `colors.go` — `Color{R, G, B}` (intentionally not aliased to `Vec3`). `AddColors`, `MulColors`, `MulScalarColors`, `RandomColor`, `RandomColorRange`. `WriteColor()` applies gamma-2 (sqrt) and clamps to `[0, 0.999]` before mapping to `[0, 255]` bytes and printing to stdout.

### Conventions to preserve

- `Vec3` arithmetic = free functions, value semantics. Don't add methods on `Vec3` for arithmetic unless matching the existing `Negation` style.
- `Color` is intentionally a separate struct from `Vec3`. Keep them separate; don't unify without reason.
- Stdout = image bytes. Logs/progress → stderr.
- Sphere `Hit` uses the simplified quadratic (see `sphere.go` notes).
- Random number generation uses the package-global `math/rand` source. Go 1.20+ auto-seeds it, so no `rand.Seed` call is needed.
- Recursion epsilon: `tMin = 0.001` in `Camera.rayColor` to avoid shadow acne from self-intersection. Don't lower without reason.

## Reference output

`cimage.png`, `image.ppm`, `new.ppm`, `sphere1.ppm`, `newsphereblue.ppm` are committed render snapshots from earlier chapters. Useful as visual regression checks; don't overwrite without intent.

## README

`README.md` is a learning log following the book chapter-by-chapter, not a user manual. Useful for context on what concept the code currently implements.
