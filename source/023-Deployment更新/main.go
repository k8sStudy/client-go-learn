package main

import (
	"../../config"
	"context"
	"fmt"
	"github.com/modood/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"log"
	"strconv"
	"time"
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

	// 开始更新，参考如下
	// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
	deploymentsClient := clientSet.AppsV1().Deployments(config.NameSpaceDefault)

	// 开始更新
	fmt.Println("开始更新 deployment...")
	updateName := "demo-deployment"
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// 在尝试更新之前检索部署的最新版本
		// RetryOnConflict使用指数退避来避免耗尽apis服务器
		// 先根据name获取指定deployment
		result, getErr := deploymentsClient.Get(context.TODO(), updateName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)                           // 更新副本数量
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // 更改nginx版本
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}
	fmt.Printf("更新完成 deployment")

	// 输出结果
	// 输出deployments列表相关信息
	time.Sleep(5 * time.Second)
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
