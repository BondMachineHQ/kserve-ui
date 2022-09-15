var k8s = require('@kubernetes/client-node');
var kc = new k8s.KubeConfig();
//Uncomment only for TEST
//kc.loadFromDefault();
kc.loadFromCluster();
var k8sApi = kc.makeApiClient(k8s.CoreV1Api);
const k8sCustomClient = kc.makeApiClient(k8s.CustomObjectsApi);

const express = require('express');
const { resolve } = require('path');
const app = express()
app.use(express.json());
const port = 3000

app.get('/list_pods', function (req, res) {
  var r = k8sCustomClient.listNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices")
    .then(result => res.status(200).json(result.body))
    .catch(err => res.status(500).send(err));
});

app.post('/delete_isvc', function (req, res) {

  console.log(req.body)
  k8sCustomClient.deleteNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices",req.body.isvcname)
      .then((result)=>{
          //console.log(result)
          res.status(200).json({"message": "Deleting inference service " + result.body.metadata.name})
      })
      .catch((err)=>{
          //console.log(err.body.message)
          res.status(err.statusCode).json({ 'message': 'ERROR: ' + err.body.message })
      })

})

app.post('/create_pod', function (req, res) {

  //req.get("email")
  //console.log(req.body.isvctype)

  switch (req.body.isvctype) {
    case "xgboost":
    case "sklearn":
      var body = {
        "apiVersion": "serving.kserve.io/v1beta1",
        "kind": "InferenceService",
        "metadata": {
            "name": req.body.isvcname,
        },
        "spec": {
            "predictor": {
              [req.body.isvctype]: {
                "protocolVersion": "v2",
                "storageUri": req.body.url
              }
            }
        }
      }
      break;
    case "onnx":
    case "tensorflow":
      var body = {
        "apiVersion": "serving.kserve.io/v1beta1",
        "kind": "InferenceService",
        "metadata": {
            "name": req.body.isvcname,
        },
        "spec": {
            "predictor": {
              [req.body.isvctype]: {
                "protocolVersion": "v1",
                "storageUri": req.body.url
              }
            }
          }
      }
      break;
  }
  // TODO use case isvctype

  
  k8sCustomClient.createNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices",body)
      .then((result)=>{
          //console.log(result)
          res.status(200).json({"message": "CREATED: inference service " + result.body.metadata.name})
      })
      .catch((err)=>{
          //console.log(err.body.message)
          res.status(err.statusCode).json({ 'message': 'ERROR: ' + err.body.message })
      })

})

app.use(express.static(__dirname + '/static'))

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})