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
	StatusParse       string
	HostIP            string
	PodIP             string
	Namespace         string
	UID               string
	CreationTimestamp string
}

func main() {
	clientSet, err := config.InitClientToken()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	podsList, err := clientSet.CoreV1().Pods(config.NameSpaceAll).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	// 输出pod列表相关信息
	res := make([]Result, 0)
	fmt.Println("pod list len:\t", len(podsList.Items))
	for i := 0; i < len(podsList.Items); i++ {
		node := podsList.Items[i]
		res = append(res, Result{
			No:                strconv.Itoa(i),
			Name:              node.Name,
			Namespace:         node.Namespace,
			HostIP:            node.Status.HostIP,
			PodIP:             node.Status.PodIP,
			StatusParse:       string(node.Status.Phase),
			UID:               string(node.UID),
			CreationTimestamp: node.CreationTimestamp.String(),
		})
	}
	table.Output(res)
}
