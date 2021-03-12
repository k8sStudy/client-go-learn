package main

import (
	"../../config"
	"context"
	"fmt"
	"github.com/modood/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// 开始删除，参考如下
	// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
	deploymentsClient := clientSet.AppsV1().Deployments(config.NameSpaceDefault)

	// 开始删除
	fmt.Println("开始删除 deployment...")
	//  指定name
	deleteName := "demo-deployment-1"
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), deleteName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
		return
	}
	fmt.Printf("删除完成 deployment")

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
