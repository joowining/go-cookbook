package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() error {
	// temp라는 임시파일을 반환하는 t라는 위치에 생성	
	t, err := ioutil.TempDir("", "temp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	return nil
}
