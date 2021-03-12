package main

import (
	"../../config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/modood/table"
	"io/ioutil"
	"strconv"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"log"
)

func main() {
	//
	// https://github.com/owenliang/k8s-client-go/blob/master/demo2/main.go
	var err error
	clientSet, err := config.InitClientToken()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	var deployYaml []byte
	var deployJson []byte
	var deployment v1.Deployment
	// 1、读取yaml文件
	if deployYaml, err = ioutil.ReadFile(config.YamlFileNginx); err != nil {
		fmt.Println("err:\t", err)
		return
	}
	// 2、yaml转json
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		fmt.Println("err:\t", err)
		return
	}
	// 3、json转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		log.Fatal("err:\t", err)
		return
	}
	deployment.Name = "nginx-3"
	deployment.Spec.Replicas = int32Ptr(3)
	deployments := clientSet.AppsV1().Deployments(config.NameSpaceDefault)

	if _, err = deployments.Get(context.TODO(), deployment.Name, metav1.GetOptions{}); err != nil {
		// 不存在就创建
		if _, err = deployments.Create(context.Background(), &deployment, metav1.CreateOptions{}); err != nil {
			log.Fatal("err:\t", err)
			return
		}
	} else {
		// 存在就更新
		if _, err = deployments.Update(context.Background(), &deployment, metav1.UpdateOptions{}); err != nil {
			log.Fatal("err:\t", err)
			return
		}
	}
	fmt.Println()
	// 输出deployments列表相关信息
	deploymentsList, err := deployments.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
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

func int32Ptr(i int32) *int32 { return &i }

type Result struct {
	No                string
	Name              string
	Namespace         string
	Replicas          string
	UID               string
	CreationTimestamp string
}
