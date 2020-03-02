package helpers

import "os"

func MakeDir(dn string){
	_, err := os.Stat(dn)
	if err != nil {
		os.Mkdir(dn, 0777)
	}else{
		os.Remove(dn + "/*")
	}
}