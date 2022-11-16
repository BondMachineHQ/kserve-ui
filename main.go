package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"context"

	kserveapi "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	servingv1beta1 "github.com/kserve/kserve/pkg/client/clientset/versioned/typed/serving/v1beta1"
	kserveconstants "github.com/kserve/kserve/pkg/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed static/*
var static embed.FS

func list_isvc(client *servingv1beta1.ServingV1beta1Client, ctx context.Context, namespace string) ([]kserveapi.InferenceService, error) {
	isvc_list, err := client.InferenceServices(namespace).List(ctx, metav1.ListOptions{})
	isvc_list_new := make([]kserveapi.InferenceService, len(isvc_list.Items))

	for i := 0; i < len(isvc_list.Items); i++ {
		switch isvc_list.Items[i].Spec.Predictor.Model.ModelFormat.Name {

		case "onnx":
			isvc_list_new[i] = kserveapi.InferenceService{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "serving.kserve.io/v1beta1",
					Kind:       "InferenceService",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      isvc_list.Items[i].Name,
					Namespace: isvc_list.Items[i].Namespace,
				},
				Spec: kserveapi.InferenceServiceSpec{
					Predictor: kserveapi.PredictorSpec{
						ONNX: &kserveapi.ONNXRuntimeSpec{
							PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
								ProtocolVersion: isvc_list.Items[i].Spec.Predictor.Model.ProtocolVersion,
								StorageURI:      isvc_list.Items[i].Spec.Predictor.Model.StorageURI,
								Container: v1.Container{
									Args: isvc_list.Items[i].Spec.Predictor.Model.Args,
								},
							},
						},
					},
				},
			}

		case "xgboost":
		case "sklearn":
			isvc_list_new[i] = kserveapi.InferenceService{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "serving.kserve.io/v1beta1",
					Kind:       "InferenceService",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      isvc_list.Items[i].Name,
					Namespace: isvc_list.Items[i].Namespace,
				},
				Spec: kserveapi.InferenceServiceSpec{
					Predictor: kserveapi.PredictorSpec{
						SKLearn: &kserveapi.SKLearnSpec{
							PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
								ProtocolVersion: isvc_list.Items[i].Spec.Predictor.Model.ProtocolVersion,
								StorageURI:      isvc_list.Items[i].Spec.Predictor.Model.StorageURI,
							},
						},
					},
				},
			}
		case "tensorflow":
			isvc_list_new[i] = kserveapi.InferenceService{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "serving.kserve.io/v1beta1",
					Kind:       "InferenceService",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      isvc_list.Items[i].Name,
					Namespace: isvc_list.Items[i].Namespace,
				},
				Spec: kserveapi.InferenceServiceSpec{
					Predictor: kserveapi.PredictorSpec{
						Tensorflow: &kserveapi.TFServingSpec{
							PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
								ProtocolVersion: isvc_list.Items[i].Spec.Predictor.Model.ProtocolVersion,
								StorageURI:      isvc_list.Items[i].Spec.Predictor.Model.StorageURI,
							},
						},
					},
				},
			}
		}
	}

	if err != nil {
		return nil, err
	}
	return isvc_list_new, nil
}

func delete_isvc(client *servingv1beta1.ServingV1beta1Client, ctx context.Context, namespace string, name string) error {
	err := client.InferenceServices(namespace).Delete(ctx, name, metav1.DeleteOptions{})

	if err != nil {
		return err
	}
	return nil
}

func create_isvc(ctx context.Context, isvcName string, name string, uri string, client *servingv1beta1.ServingV1beta1Client, namespace string) (string, error) {
	var svc kserveapi.InferenceService
	switch isvcName {
	case "onnx":
		protocol := kserveconstants.ProtocolV2

		svc = kserveapi.InferenceService{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "serving.kserve.io/v1beta1",
				Kind:       "InferenceService",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "default",
			},
			Spec: kserveapi.InferenceServiceSpec{
				Predictor: kserveapi.PredictorSpec{
					ONNX: &kserveapi.ONNXRuntimeSpec{
						PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
							ProtocolVersion: &protocol,
							StorageURI:      &uri,
							Container: v1.Container{
								Args: []string{
									"--strict-model-config=false",
								},
							},
						},
					},
				},
			},
		}

	case "xgboost":
	case "sklearn":

		protocol := kserveconstants.ProtocolV2

		svc = kserveapi.InferenceService{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "serving.kserve.io/v1beta1",
				Kind:       "InferenceService",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "default",
			},
			Spec: kserveapi.InferenceServiceSpec{
				Predictor: kserveapi.PredictorSpec{
					SKLearn: &kserveapi.SKLearnSpec{
						PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
							ProtocolVersion: &protocol,
							StorageURI:      &uri,
						},
					},
				},
			},
		}

	case "tensorflow":
		protocol := kserveconstants.ProtocolV1

		svc = kserveapi.InferenceService{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "serving.kserve.io/v1beta1",
				Kind:       "InferenceService",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "default",
			},
			Spec: kserveapi.InferenceServiceSpec{
				Predictor: kserveapi.PredictorSpec{
					Tensorflow: &kserveapi.TFServingSpec{
						PredictorExtensionSpec: kserveapi.PredictorExtensionSpec{
							ProtocolVersion: &protocol,
							StorageURI:      &uri,
						},
					},
				},
			},
		}
	}

	_, err := client.InferenceServices(namespace).Create(ctx, &svc, metav1.CreateOptions{})
	if err != nil {
		return "", err
	} else {
		return "", nil
	}
}

func main() {
	content, _ := fs.Sub(static, "static")
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":3000", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
