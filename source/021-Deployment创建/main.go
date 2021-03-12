package main

import (
	"../../config"
	"context"
	"fmt"
	"github.com/modood/table"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"strconv"
)

type Result struct {
	No                string
	Name              string
	Namespace         string
	Replicas          string
	UID               string
	CreationTimestamp string
}

func main() {
	clientSet, err := config.InitClientToken()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	deploymentsList, err := clientSet.AppsV1().Deployments(config.NameSpaceDefault).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	// 输出deployments列表相关信息
	res := make([]Result, 0)
	fmt.Println("deployments list len:\t", len(deploymentsList.Items))
	for i := 0; i < len(deploymentsList.Items); i++ {
		node := deploymentsList.Items[i]
		res = append(res, Result{
			No:                strconv.Itoa(i),
			Name:              node.Name,
			Namespace:         node.Namespace,
			Replicas:          strconv.Itoa(int(node.Status.Replicas)),
			UID:               string(node.UID),
			CreationTimestamp: node.CreationTimestamp.String(),
		})
	}
	table.Output(res)

	// 开始创建，参考如下
	// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
	deploymentsClient := clientSet.AppsV1().Deployments(config.NameSpaceDefault)

	// 准备需要创建的deployment
	// 待创建的名为demo-deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment-1",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// 开始创建
	fmt.Println("开始创建 deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("创建完成 deployment %q.\n", result.GetObjectMeta().GetName())

	// 输出结果
	// 输出deployments列表相关信息
	deploymentsList, err = clientSet.AppsV1().Deployments(config.NameSpaceDefault).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	res = make([]Result, 0)
	fmt.Println()
	fmt.Println()
	fmt.Println("deployments list len:\t", len(deploymentsList.Items))
	for i := 0; i < len(deploymentsList.Items); i++ {
		node := deploymentsList.Items[i]
		res = append(res, Result{
			No:                strconv.Itoa(i),
			Name:              node.Name,
			Namespace:         node.Namespace,
			Replicas:          strconv.Itoa(int(node.Status.Replicas)),
			UID:               string(node.UID),
			CreationTimestamp: node.CreationTimestamp.String(),
		})
	}
	table.Output(res)
}

func int32Ptr(i int32) *int32 { return &i }
