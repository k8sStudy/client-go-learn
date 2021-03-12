package config

var (
	Token                       = "xx"                                             // 使用token连接
	Address                     = "https://127.0.0.1:6443"                         // ip:6443端口
	Version                     = "v1"                                             //
	ApplicationJson             = "application/json"                               //
	NameSpace                   = "project_name"                                   //
	NameSpaceDefault            = "default"                                        //
	NameSpaceAll                = ""                                               //
	ConfFile                    = "config/admin.conf"                              //
	YamlFileNginx               = "config/nginx.yaml"                              //
	YamlRocketMqBrokerCr        = "config/rocketmq_v1alpha1_broker_cr.yaml"        // rocketmq broker配置文件
	YamlRocketMqClusterService  = "config/rocketmq_v1alpha1_cluster_service.yaml"  // rocketmq 集群服务配置文件
	YamlRocketMqConsoleCr       = "config/rocketmq_v1alpha1_console_cr.yaml"       // rocketmq console配置文件
	YamlRocketMqNameserviceCr   = "config/rocketmq_v1alpha1_nameservice_cr.yaml"   // rocketmq nameservice配置文件
	YamlRocketMqCluster         = "config/rocketmq_v1alpha1_rocketmq_cluster.yaml" // rocketmq 集群配置文件
	YamlRocketMqTopicTransferCr = "config/rocketmq_v1alpha1_topictransfer_cr.yaml" // rocketmq topic配置文件
)
