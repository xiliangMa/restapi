package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs"
	"github.com/bitly/go-simplejson"
)

var (
	client = GetClient()
	cookie = SetCookie("R_USERNAME=admin; CSRF=98d196e686; R_SESS=token-tp8h6:nqlv67vphwcvg68d2rwwl2cwp8zvnsbwc47fpzq84thsfl8fk8l4cd")
)

func CreateCluster() {
	params := `{
		"amazonElasticContainerServiceConfig": null,
		"azureKubernetesServiceConfig": null,
		"dockerRootDir": "/var/lib/docker",
		"googleKubernetesEngineConfig": null,
		"name": "test",
		"rancherKubernetesEngineConfig": {
			"provider":"custom"
		}
	}`
	url := SetRequesterUrl("https://47.105.151.140/v3/clusters")
	res, err := client.Do(SetRequester(C_JsonType, cookie, url, M_PostType, params))
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("测试结果：========================== 创建 集群 失败：", err.Error())
	} else {
		fmt.Println("测试结果：========================== 创建 集群 成功 ", string(body))
	}
}

func GetClusterRegistrationToken() {
	params := ""
	url := SetRequesterUrl("https://47.105.151.140/v3/clusterregistrationtokens?clusterId=c-jqkgr")
	res, _ := client.Do(SetRequester(C_JsonType, cookie, url, M_GetType, params))
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	data, _ := simplejson.NewJson(body)
	rows, _ := data.Get("data").Array()

	if err != nil {
		fmt.Println("测试结果：========================== 获取 集群注册脚本 失败：", err.Error())
	} else {
		for _, row := range rows {
			if nodecommand, ok := row.(map[string]interface{}); ok {
				fmt.Println("测试结果：========================== 获取 集群注册脚本 成功: ", nodecommand["nodeCommand"])
			}
		}
	}
}

func GetClusterRegistrationToken1() {
	params := ""
	url := SetRequesterUrl("https://47.105.151.140/v3/clusterregistrationtokens?clusterId=c-jqkgr")
	res, _ := client.Do(SetRequester(C_JsonType, cookie, url, M_GetType, params))
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	jsonParse, _ := gabs.ParseJSON(body)
	fmt.Println(jsonParse.Path("data").Path("nodeCommand"))

}

func CreateProject() {
	params := `{
	"clusterId": "c-2nbhs",
	"name": "maxl",
	"namespaceDefaultResourceQuota": {"limit": {"limitsCpu": "2000m", "limitsMemory": "2048Mi"}},
	"namespaceId": "",
	"resourceQuota": {"limit": {"limitsCpu": "2000m", "limitsMemory": "2048Mi"}}
	}`
	url := SetRequesterUrl("https://47.105.151.140/v3/projects")
	res, _ := client.Do(SetRequester(C_JsonType, cookie, url, M_PostType, params))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("测试结果：========================== 创建 项目  失败: ", err.Error())
	} else {
		fmt.Println("测试结果：========================== 创建 项目  成功: ", string(body))
	}

}

//c-2nbhs:p-rmmtj
func CreateNameSpace() {
	params := `{
	"name": "maxl-n",
	"projectId": "c-2nbhs:p-rmmtj",
	"resourceQuota": {
		"limit": {
			"limitsCpu": "2000m",
			"limitsMemory": "2048Mi"
			}
		}
	}`
	url := SetRequesterUrl("https://47.105.151.140/v3/cluster/c-2nbhs/namespaces")
	res, _ := client.Do(SetRequester(C_JsonType, cookie, url, M_PostType, params))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("测试结果：========================== 创建 命名空间  失败: ", err.Error())
	} else {
		fmt.Println("测试结果：========================== 创建 命名空间  成功: ", string(body))
	}
}

func main() {
	//CreateCluster()
	GetClusterRegistrationToken1()
	//CreateProject()
	//CreateNameSpace()
}
