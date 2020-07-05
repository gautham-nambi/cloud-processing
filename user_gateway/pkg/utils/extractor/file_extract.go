package extractor

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xi2/xz"
)

func extractor() {
	// Open a file pass the file name here: need to test this stuff
	f, err := os.Open("")
	if err != nil {
		log.Fatal(err)
	}
	// Create an xz Reader
	r, err := xz.NewReader(f, 0)
	if err != nil {
		log.Fatal(err)
	}
	// Create a tar Reader
	tr := tar.NewReader(r)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			// create a directory
			fmt.Println("creating:   " + hdr.Name)
			err = os.MkdirAll(hdr.Name, 0777)
			if err != nil {
				log.Fatal(err)
			}
		case tar.TypeReg, tar.TypeRegA:
			// write a file
			fmt.Println("extracting: " + hdr.Name)
			w, err := os.Create(hdr.Name)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(w, tr)
			if err != nil {
				log.Fatal(err)
			}
			w.Close()
		}
	}
	f.Close()
}
