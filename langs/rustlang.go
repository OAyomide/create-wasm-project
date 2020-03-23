package langs

import (
	"fmt"
	"log"
	"os/exec"
	"path"
)

// we want to run the rustinstall bash script

func Install_deps() {
	// file, err := filepath.Abs(filepath.Dir("./rustinstall.sh"))\
	file := path.Dir("rustinstall.sh")
	fmt.Printf("%s\n", file)
	out, err := exec.Command("/bin/bash", "../langs/rustinstall.sh").CombinedOutput()

	if err != nil {
		log.Fatalf("Error running the bash script...: %s", err.Error())
	}

	fmt.Printf("output of running rustinstall.sh: %s", string(out))
}
