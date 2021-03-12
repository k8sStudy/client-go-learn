package main

import (
	"../../config"
	"context"
	"fmt"
	"github.com/modood/table"
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
}
