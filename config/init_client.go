package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func InitClientToken() (clientSet *kubernetes.Clientset, err error) {
	// 配置参数
	kubeConfig := &rest.Config{
		Host:    Address, // 地址
		APIPath: Version, //
		ContentConfig: rest.ContentConfig{ //
			AcceptContentTypes: ApplicationJson,
			ContentType:        ApplicationJson,
		},
		BearerToken: Token, //
		TLSClientConfig: rest.TLSClientConfig{ //
			Insecure: true,
		},
	}
	clientSet, err = kubernetes.NewForConfig(kubeConfig)
	return clientSet, err
}

func InitClientConfigFile() (clientSet *kubernetes.Clientset, err error) {
	var kubeConfig *rest.Config
	// todo
	clientSet, err = kubernetes.NewForConfig(kubeConfig)
	return clientSet, err
}

func InitClientCmd() (clientSet *kubernetes.Clientset, err error) {
	var kubeConfig *rest.Config
	// todo
	clientSet, err = kubernetes.NewForConfig(kubeConfig)
	return clientSet, err
}
