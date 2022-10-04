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
// k8sCustomClient.deleteNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices",req.body.isvcname)
// .then((result)=>{
// 	//console.log(result)
// 	res.status(200).json({"message": "Deleting inference service " + result.body.metadata.name})
// })
// .catch((err)=>{
// 	//console.log(err.body.message)
// 	res.status(err.statusCode).json({ 'message': 'ERROR: ' + err.body.message })
// })

// TODO: implement create_isvc
// switch (req.body.isvctype) {
// case "onnx":
//   var body = {
// 	"apiVersion": "serving.kserve.io/v1beta1",
// 	"kind": "InferenceService",
// 	"metadata": {
// 		"name": req.body.isvcname,
// 	},
// 	"spec": {
// 		"predictor": {
// 		  [req.body.isvctype]: {
// 			"protocolVersion": "v2",
// 			"storageUri": req.body.url,
// 			"args": ["--strict-model-config=false"]
// 		  }
// 		}
// 	}
//   }
//   break;
// case "xgboost":
// case "sklearn":
//   var body = {
// 	"apiVersion": "serving.kserve.io/v1beta1",
// 	"kind": "InferenceService",
// 	"metadata": {
// 		"name": req.body.isvcname,
// 	},
// 	"spec": {
// 		"predictor": {
// 		  [req.body.isvctype]: {
// 			"protocolVersion": "v2",
// 			"storageUri": req.body.url,
// 		  }
// 		}
// 	}
//   }
//   break;
// case "tensorflow":
//   var body = {
// 	"apiVersion": "serving.kserve.io/v1beta1",
// 	"kind": "InferenceService",
// 	"metadata": {
// 		"name": req.body.isvcname,
// 	},
// 	"spec": {
// 		"predictor": {
// 		  [req.body.isvctype]: {
// 			"protocolVersion": "v1",
// 			"storageUri": req.body.url
// 		  }
// 		}
// 	  }
//   }
//   break;
// }

// k8sCustomClient.createNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices",body)
//   .then((result)=>{
// 	  //console.log(result)
// 	  res.status(200).json({"message": "CREATED: inference service " + result.body.metadata.name})
//   })
//   .catch((err)=>{
// 	  //console.log(err.body.message)
// 	  res.status(err.statusCode).json({ 'message': 'ERROR: ' + err.body.message })
//   })

func main() {
	content, _ := fs.Sub(static, "static")
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":3000", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
