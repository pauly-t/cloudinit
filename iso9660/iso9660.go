package iso9660

// Iso9660 ...
type Iso9660 struct {
	ivd iso9660VolumeDescriptor
}

// New ...
func New() Iso9660 {
	iso := Iso9660{}
	iso.ivd._type = 1
	copy(iso.ivd.id[:], iso9660DefaultID)

	return iso
}

// Test ...
func (iso *Iso9660) Test() {

}
