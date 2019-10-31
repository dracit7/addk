package visual

import (
	"bytes"
	"io/ioutil"

	"github.com/bradleyjkemp/memviz"
)

var path = "instances/dot"

// Marshal transform a data structure to visible string.
func Marshal(ds interface{}) (buf bytes.Buffer) {
	memviz.Map(&buf, ds)
	return
}

// Dump the Marshaled text to a .dot file.
func Dump(ds interface{}) error {

	buf := &bytes.Buffer{}
	memviz.Map(buf, ds)

	err := ioutil.WriteFile(path+"/test.dot", buf.Bytes(), 0644)
	return err

}

// SetPath sets the root path of instances
func SetPath(newpath string) {
	path = newpath
}
