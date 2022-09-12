var k8s = require('@kubernetes/client-node');
var kc = new k8s.KubeConfig();
//kc.loadFromDefault();
kc.loadFromCluster();
var k8sApi = kc.makeApiClient(k8s.CoreV1Api);
k8sApi.listNamespacedPod('af')
    .then(function (res) {
    // tslint:disable-next-line:no-console
    console.log(res.body);
});
