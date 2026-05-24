package scene

// HitRecord captures the data about a ray-surface intersection.
type HitRecord struct {
	P         Point3
	Normal    Vec3
	Mat       Material
	T         float64
	FrontFace bool
}

// SetFaceNormal orients Normal so it always faces against the incoming ray.
// outwardNormal must be unit length.
func (h *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	h.FrontFace = Dot(r.direction, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negation()
	}
}

type Hittable interface {
	Hit(r Ray, rayT Interval) (bool, HitRecord)
}

// HittableList is a collection of Hittables tested in order.
type HittableList struct {
	Objects []Hittable
}

func (h *HittableList) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HittableList) Clear() {
	h.Objects = nil
}

func (h *HittableList) Hit(r Ray, rayT Interval) (bool, HitRecord) {
	var rec HitRecord
	hitAnything := false
	closest := rayT.Max
	for _, obj := range h.Objects {
		if ok, tmp := obj.Hit(r, Interval{Min: rayT.Min, Max: closest}); ok {
			hitAnything = true
			closest = tmp.T
			rec = tmp
		}
	}
	return hitAnything, rec
}
