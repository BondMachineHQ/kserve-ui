var k8s = require('@kubernetes/client-node');
var kc = new k8s.KubeConfig();
kc.loadFromDefault();
//kc.loadFromCluster();
var k8sApi = kc.makeApiClient(k8s.CoreV1Api);
const k8sCustomClient = kc.makeApiClient(k8s.CustomObjectsApi);

const express = require('express');
const { resolve } = require('path');
const app = express()
const port = 3000

app.get('/list_pods', function (req, res) {
  var r = k8sCustomClient.listNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices")
    .then(result => res.status(200).json(result.body))
    .catch(err => res.status(500).send(err));
});

app.post('/create_pod', function (req, res) {

  var body = {
      "apiVersion": "serving.kserve.io/v1beta1",
      "kind": "InferenceService",
      "metadata": {
          "name": "sklearn-irisv2",
      },
      "spec": {
          "predictor": {
            "sklearn": {
              "protocolVersion": "v2",
              "storageUri": "gs://seldon-models/sklearn/mms/lr_model"
            }
          }
      }
  }
  
  k8sCustomClient.createNamespacedCustomObject('serving.kserve.io','v1beta1','default', "inferenceservices",body)
      .then((result)=>{
          console.log(result)
      })
      .catch((err)=>{
          console.log(err)
      })

  // k8sApi.createNamespacedCustomObject("inferenceService", "serving.kserve.io/v1beta1", "af", )

  // k8sApi.createNamespacedPod("af", pod)
  //     .then((result) => {
  //         res.status(200).json(result.body)
  //         console.log(result.body);
  //     })
  //     .catch((err) => {
  //       res.status(500).send(err)
  //         console.log(err);
  //     });

});

app.use(express.static(__dirname + '/static'))

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})