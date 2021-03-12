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
	ClusterIP         string
	UID               string
	CreationTimestamp string
}

func main() {
	clientSet, err := config.InitClientToken()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	servicesList, err := clientSet.CoreV1().Services(config.NameSpaceAll).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	// 输出service列表相关信息
	res := make([]Result, 0)
	fmt.Println("service list len:\t", len(servicesList.Items))
	for i := 0; i < len(servicesList.Items); i++ {
		node := servicesList.Items[i]
		res = append(res, Result{
			No:                strconv.Itoa(i),
			Name:              node.Name,
			Namespace:         node.Namespace,
			ClusterIP:         node.Spec.ClusterIP,
			UID:               string(node.UID),
			CreationTimestamp: node.CreationTimestamp.String(),
		})
	}
	table.Output(res)
}
