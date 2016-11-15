// Copyright 2016 Aurélien Rainone. All rights reserved.
// Use of this source code is governed by MIT license.
// license that can be found in the LICENSE file.

package d2

import "testing"

func TestRectangle(t *testing.T) {
	rects := []Rectangle{
		Rect(0, 0, 10, 10),
		Rect(1, 2, 3, 4),
		Rect(4, 6, 10, 10),
		Rect(2, 3, 12, 5),
		Rect(-1, -2, 0, 0),
		Rect(-1, -2, 4, 6),
		Rect(-10, -20, 30, 40),
		Rect(8, 8, 8, 8),
		Rect(88, 88, 88, 88),
		Rect(6, 5, 4, 3),
	}

	// r.Eq(s) should be equivalent to every point in r being in s, and every
	// point in s being in r.
	for _, r := range rects {
		for _, s := range rects {
			got := r.Eq(s)
			want := r.In(s) && s.In(r)
			if got != want {
				t.Errorf("Eq: r=%s, s=%s: got %t, want %t", r, s, got, want)
			}
		}
	}

	// The intersection should be the largest rectangle a such that every point
	// in a is both in r and in s.
	for _, r := range rects {
		for _, s := range rects {
			a := r.Intersect(s)
			if !a.In(r) {
				t.Errorf("Intersect: r=%s, s=%s, a=%s, a not in r", r, s, a)
			}
			if !a.In(s) {
				t.Errorf("Intersect: r=%s, s=%s, a=%s, a not in s", r, s, a)
			}
			if a.Empty() == r.Overlaps(s) {
				t.Errorf("Intersect: r=%s, s=%s, a=%s: empty=%t same as overlaps=%t",
					r, s, a, a.Empty(), r.Overlaps(s))
			}
			largerThanA := [4]Rectangle{CopyRect(a), CopyRect(a), CopyRect(a), CopyRect(a)}
			largerThanA[0].Min[0] -= 1
			largerThanA[1].Min[1] -= 1
			largerThanA[2].Max[0] += 1
			largerThanA[3].Max[1] += 1
			for i, b := range largerThanA {
				if b.Empty() {
					// b isn't actually larger than a.
					continue
				}
				if b.In(r) && b.In(s) {
					t.Errorf("Intersect: r=%s, s=%s, a=%s, b=%s, i=%d: intersection could be larger",
						r, s, a, b, i)
				}
			}
		}
	}

	// The union should be the smallest rectangle a such that every point in r
	// is in a and every point in s is in a.
	for _, r := range rects {
		for _, s := range rects {
			a := r.Union(s)
			if !r.In(a) {
				t.Errorf("Union: r=%s, s=%s, a=%s, r not in a", r, s, a)
			}
			if !s.In(a) {
				t.Errorf("Union: r=%s, s=%s, a=%s, s not in a", r, s, a)
			}
			if a.Empty() {
				// You can't get any smaller than a.
				continue
			}
			smallerThanA := [4]Rectangle{CopyRect(a), CopyRect(a), CopyRect(a), CopyRect(a)}
			smallerThanA[0].Min[0]++
			smallerThanA[1].Min[1]++
			smallerThanA[2].Max[0]--
			smallerThanA[3].Max[1]--
			for i, b := range smallerThanA {
				if r.In(b) && s.In(b) {
					t.Errorf("Union: r=%s, s=%s, a=%s, b=%s, i=%d: union could be smaller",
						r, s, a, b, i)
				}
			}
		}
	}
}

func TestRectangleCanon(t *testing.T) {
	r1 := Rect(1, 2, 3, 4)
	r2 := Rect(3, 4, 1, 2)
	if !r1.Eq(r2.Canon()) {
		t.Errorf("Canon: %v != %v, want ==", r1, r2.Canon())
	}
}
