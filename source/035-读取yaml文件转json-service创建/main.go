package main

import (
	"../../config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/modood/table"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strconv"
)

func main() {
	var err error
	clientSet, err := config.InitClientToken()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	var deployYaml []byte
	var deployJson []byte
	var service v1.Service
	// 1、读取yaml文件
	if deployYaml, err = ioutil.ReadFile(config.YamlRocketMqClusterService); err != nil {
		fmt.Println("err:\t", err)
		return
	}
	// 2、yaml转json
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		fmt.Println("err:\t", err)
		return
	}
	// 3、json转struct
	if err = json.Unmarshal(deployJson, &service); err != nil {
		log.Fatal("err:\t", err)
		return
	}

	fmt.Println(service)
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

	// 开始创建service
	services := clientSet.CoreV1().Services(config.NameSpaceDefault)
	if _, err = services.Get(context.TODO(), service.Name, metav1.GetOptions{}); err != nil {
		// 不存在就创建
		if _, err = services.Create(context.Background(), &service, metav1.CreateOptions{}); err != nil {
			log.Fatal("err:\t", err)
			return
		}
	} else {
		// 存在就更新
		if _, err = services.Update(context.Background(), &service, metav1.UpdateOptions{}); err != nil {
			log.Fatal("err:\t", err)
			return
		}
	}
	fmt.Println()

	// 输出结果
	servicesList, err = clientSet.CoreV1().Services(config.NameSpaceAll).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	// 输出service列表相关信息
	res = make([]Result, 0)
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

type Result struct {
	No                string
	Name              string
	Namespace         string
	ClusterIP         string
	UID               string
	CreationTimestamp string
}
