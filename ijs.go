package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"inaxium.com/ijs/cmd"
	"inaxium.com/ijs/public"
	"log"
	"os"
	"path"
	"strings"
)

const (
	EMPTY                          = ""
	IJS, isRecursive, notRecursive = "ijs", true, false
)
var (
	client *minio.Client
	meta   public.Meta
	err    error
)

func init() {
	meta = cmd.Init()
}

func main() {
	endpoint := "store.inaxiumjs.com"

	// Read Only Policy defined on Server no Problem to Publish this Information
	accessKeyID := "ijsuser"
	secretAccessKey := "ijsstore4711"

	client, err = minio.New(endpoint, accessKeyID, secretAccessKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	if viper.IsSet("type"){
		if meta.Type != "framework" && meta.Type!= "demo"{
			fmt.Println("Wrong type given use ijs -h for Help")
			os.Exit(0)
		}
	}

	if viper.IsSet("show") {
		lstVersion()
	}else if viper.IsSet("copy"){
		cp();
	}
}

func lstVersion(){
	doneCh := make(chan struct{})

	defer close(doneCh)


	objectCh := client.ListObjects(IJS, EMPTY, notRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			log.Println(object.Err)
			return
		}

		fmt.Println(strings.ReplaceAll(object.Key,"/",""))
	}

}

func cp(){
	doneCh := make(chan struct{})

	defer close(doneCh)

	objectCh := client.ListObjects(IJS, path.Join(meta.Version,meta.Type), isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			log.Println(object.Err)
			return
		}

		err = client.FGetObject(IJS, object.Key, path.Join(meta.Destination,IJS,object.Key), minio.GetObjectOptions{})
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Done "+object.Key)
	}


}