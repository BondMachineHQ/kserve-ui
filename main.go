package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var static embed.FS

// TODO: implement list_isvc
// var r = k8sCustomClient.listNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices")
// .then(result => res.status(200).json(result.body))
// .catch(err => res.status(500).send(err));

// TODO: implement delete_isvc

// TODO: implement create_isvc

func main() {
	content, _ := fs.Sub(static, "static")
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":3000", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
