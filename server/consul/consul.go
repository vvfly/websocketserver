package consul

import (
	"fmt"
	"os"

	consulapi "github.com/hashicorp/consul/api"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
)

type ConsulManager struct {
	Addr    string
	Port    int
	Name    string
	Id      string
	SrvAddr string
}

var consulManager *ConsulManager = nil
var ConsulServices = make(map[string]*ConsulManager)

func GetConsulSrv(name string) map[string]*ConsulManager {
	var serviceMap = make(map[string]*ConsulManager)
	for k, v := range ConsulServices {
		if k == name {
			serviceMap[k] = v
		}
	}
	return serviceMap
}

func GetConsulManager() *ConsulManager {
	return consulManager
}

func ConsulManagerInit() {
	// 注册api服务到consul
	consulAddr := model.GetServerConfig().ConsulAddr
	apiConf := model.GetServerConfig().ApiServerConf
	consulManager = &ConsulManager{
		Addr:    apiConf.ServerAddr,
		Port:    apiConf.ApiPort,
		Name:    apiConf.ServerName,
		Id:      apiConf.ServerID,
		SrvAddr: consulAddr,
	}

	err := consulManager.RegisterApiSrv()
	if err != nil {
		os.Exit(1)
	}

	// 查找consul服务到map
	err = consulManager.FindAllSrvs()
	if err != nil {
		os.Exit(1)
	}

	// load balance
	BalanceInit()
}

func (c *ConsulManager) RegisterApiSrv() error {
	config := consulapi.DefaultConfig()
	config.Address = c.SrvAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		Log.Errorf("register consul server error. err=%v", err)
		return err
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.Port = c.Port
	registration.Name = c.Name
	registration.Address = c.Addr
	registration.ID = c.Id

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/health", c.Addr, c.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		Log.Errorf("register consul server error. err=%v", err)
		return err
	}
	return nil
}

func (c *ConsulManager) FindAllSrvs() error {
	config := consulapi.DefaultConfig()
	config.Address = c.SrvAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		Log.Errorf("new consul client error. err=%v", err)
		return err
	}

	services, err := client.Agent().Services()
	if err != nil {
		Log.Errorf("find consul service error. err=%v", err)
		return err
	}

	for _, v := range services {
		ConsulServices[v.Service] = &ConsulManager{
			Port: v.Port,
			Id:   v.ID,
			Addr: v.Address,
			Name: v.Service,
		}
	}

	Log.Debugf("consul services.%v", ConsulServices)

	return nil
}

// 取消consul注册的服务
func (c *ConsulManager) DeRegisterSrv(SrvId string) error {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = c.SrvAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		Log.Error(err)
		return err
	}
	return client.Agent().ServiceDeregister(SrvId)
}

// func ConsulCheckHeath() {
// 	// 创建连接consul服务配置
// 	config := consulapi.DefaultConfig()
// 	config.Address = "192.168.88.206:8500"
// 	client, err := consulapi.NewClient(config)
// 	if err != nil {
// 		log.Fatal("consul client error : ", err)
// 	}
// 	// 健康检查
// 	a, b, _ := client.Agent().AgentHealthServiceByID("111")
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func ConsulKVTest() {
// 	// 创建连接consul服务配置
// 	config := consulapi.DefaultConfig()
// 	config.Address = "192.168.88.206:8500"
// 	client, err := consulapi.NewClient(config)
// 	if err != nil {
// 		log.Fatal("consul client error : ", err)
// 	}

// 	// KV, put值
// 	values := "test"
// 	key := "go-consul-test/192.168.88.206:8500"
// 	client.KV().Put(&consulapi.KVPair{Key: key, Flags: 0, Value: []byte(values)}, nil)

// 	// KV get值
// 	data, _, _ := client.KV().Get(key, nil)
// 	fmt.Println(string(data.Value))

// 	// KV list
// 	datas, _, _ := client.KV().List("go", nil)
// 	for _, value := range datas {
// 		fmt.Println(value)
// 	}
// 	keys, _, _ := client.KV().Keys("go", "", nil)
// 	fmt.Println(keys)
// }
