package mapset

import "fmt"
import "strings"
import "strconv"

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *Set) Contains(i interface{}) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *Set) IsSubset(other *Set) bool {
	counter := 0
	for key, _ := range set.set {
		if other.Contains(key) {
			counter++
		}
	}
	return counter == set.Size()
}

func (set *Set) IsSuperset(other *Set) bool {
	counter := 0
	for key, _ := range other.set {
		if set.Contains(key) {
			counter++
		}
	}
	return counter == other.Size()
}

func (set *Set) Union(other *Set) *Set {
	if set != nil && other != nil {
		unionedSet := NewSet()

		for key, _ := range set.set {
			unionedSet.Add(key)
		}
		for key, _ := range other.set {
			unionedSet.Add(key)
		}
		return unionedSet
	}
	return nil
}

func (set *Set) Intersect(other *Set) *Set {
	if set != nil && other != nil {
		intersectedSet := NewSet()
		for key, _ := range set.set {
			if other.Contains(key) {
				intersectedSet.Add(key)
			}
		}
		for key, _ := range other.set {
			if set.Contains(key) {
				intersectedSet.Add(key)
			}
		}
		return intersectedSet
	}
	return nil
}

func (set *Set) Difference(other *Set) *Set {
	if set != nil && other != nil {
		differencedSet := NewSet()

		for key, _ := range set.set {
			differencedSet.Add(key)
		}

		for key, _ := range other.set {
			differencedSet.Add(key)
		}

		for key, _ := range other.set {
			differencedSet.Remove(key)
		}

		return differencedSet
	}
	return nil
}

func (set *Set) SymmetricDifference(other *Set) *Set {
	if set != nil && other != nil {
		aDiff := set.Difference(other)
		bDiff := other.Difference(set)

		symDifferencedSet := aDiff.Union(bDiff)

		return symDifferencedSet
	}
	return nil
}

func (set *Set) Clear() {
	set.set = make(map[interface{}]bool)
}

func (set *Set) Remove(i interface{}) {
	delete(set.set, i)
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) String() string {
	items := make([]string, 0, len(set.set))

	for key, _ := range set.set {
		switch t := key.(type) {
		case string:
			items = append(items, fmt.Sprintf(`"%s"`, t))
		case int:
			items = append(items, strconv.Itoa(t))
		case int64:
			items = append(items, strconv.Itoa(int(t)))
		case int32:
			items = append(items, strconv.Itoa(int(t)))
		case fmt.Stringer:
			items = append(items, t.String())
		}
	}

	return fmt.Sprintf("Set{" + strings.Join(items, ", ") + "}")
}