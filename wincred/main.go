package main

import (
	"fmt"

	"github.com/danieljoos/wincred"
)

func main() {
	cred := wincred.NewGenericCredential("myGoApplication")
	cred.CredentialBlob = []byte("my secret")
	err := cred.Write()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("::::::::test secret::::::::")
	fmt.Println(cred.TargetName, "=>", cred.UserName, "=>", string(cred.CredentialBlob))
	creds, err := wincred.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("::::::::list secrets::::::::")
	for i := range creds {
		fmt.Println(creds[i].TargetName, "=>", creds[i].UserName, "=>", string(creds[i].CredentialBlob))
	}
}
