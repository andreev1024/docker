package main

/**
 *  author: andreev1024@gmail.com
 *
 *  Requirements:
 *      -   installed aws cli (configured via ~/.aws/);
 *      -   in cluster must be only one task - our app task;
 *      -   work only with TASKs, not with SERVICES
 */

import (
	"encoding/json"
	"flag"
	"os/exec"
	"time"
)

//  command-string params
var test bool = false
var cluster string = "default"

func init() {

	//  set Test
	flag.BoolVar(&test, "test", false, "enable/disable test mode")

	//  set Cluster
	flag.StringVar(&cluster, "cluster", "default", "a cluster name")

	flag.Parse()
}

func main() {

	runningArnTask := getArnForRunningTask()
	arnForTaskDefinition := getArnForTaskDefinition(runningArnTask)
	stopContainer(runningArnTask)
	waitUntilContainerStopped(runningArnTask, 0)
	runNewContainer(arnForTaskDefinition)
	return

}

func getArnForRunningTask() string {

	var jsonData interface{}
	if test {
		var cmdOut string = getTestData("runningTask")
		jsonData = getTestJson(cmdOut)
	} else {
		cmd := exec.Command("aws", "ecs", "list-tasks", "--cluster="+cluster)
		cmdOut, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		jsonData = getJson(cmdOut)
	}

	taskArns := jsonData.(map[string]interface{})["taskArns"].([]interface{})

	if len(taskArns) > 1 {
		panic("Your cluster has many tasks. Must be only one!")
	}

	return taskArns[0].(string)
}

func getArnForTaskDefinition(taskArn string) string {

	var jsonData interface{}
	if test {
		var cmdOut string = getTestData("taskDefinition")
		jsonData = getTestJson(cmdOut)
	} else {
		cmd := exec.Command("aws", "ecs", "describe-tasks", "--tasks="+taskArn)
		cmdOut, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		jsonData = getJson(cmdOut)
	}

	tasks := jsonData.(map[string]interface{})["tasks"]
	task := tasks.([]interface{})[0]
	return task.(map[string]interface{})["taskDefinitionArn"].(string)
}

func runNewContainer(taskArn string) {
	if !test {
		exec.Command("aws", "ecs", "run-task", "--task-definition="+taskArn).Run()
	}
}

//  test for stop-task response skiped
func stopContainer(arn string) {
	if !test {
		exec.Command("aws", "ecs", "stop-task", "--task="+arn).Run()
	}
}

func waitUntilContainerStopped(arn string, try int) {

	if try > 5 {
		panic("Can't stop container")
	}

	var jsonData interface{}
	if test {
		var cmdOut string
		cmdOut = getTestData("runningTask")
		if try > 4 {
			cmdOut = getTestData("runningTaskEmpty")
		}
		jsonData = getTestJson(cmdOut)
	} else {
		cmd := exec.Command("aws", "ecs", "list-tasks", "--cluster="+cluster)
		cmdOut, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		jsonData = getJson(cmdOut)
	}

	taskArns := jsonData.(map[string]interface{})["taskArns"]

	for _, v := range taskArns.([]interface{}) {
		if v == arn {
			time.Sleep(time.Second)
			waitUntilContainerStopped(arn, try+1)
		}
	}

	return
}

func getJson(data []byte) interface{} {
	var jsonData interface{}
	json.Unmarshal(data, &jsonData)

	return jsonData.(map[string]interface{})
}

func getTestJson(data string) interface{} {
	var jsonData interface{}
	json.Unmarshal([]byte(data), &jsonData)

	result, ok := jsonData.(map[string]interface{})
	if !ok {
		panic("Json parse error")
	}

	return result
}

func getTestData(key string) string {

	data := make(map[string]string)

	data["runningTask"] = `{
       "taskArns": [
           "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606"
       ]
    }`

	data["runningTaskEmpty"] = `{
       "taskArns": []
    }`

	data["taskDefinition"] = `{
        "failures": [],
        "tasks": [
            {
                "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                "overrides": {
                    "containerOverrides": [
                        {
                            "name": "node"
                        },
                        {
                            "name": "nginx"
                        },
                        {
                            "name": "golang"
                        }
                    ]
                },
                "lastStatus": "RUNNING",
                "containerInstanceArn": "arn:aws:ecs:ap-northeast-1:808634764528:container-instance/ade0ff2b-dd29-44e7-a8b9-7b3d02ff6876",
                "createdAt": 1453465153.893,
                "clusterArn": "arn:aws:ecs:ap-northeast-1:808634764528:cluster/default",
                "startedAt": 1453465161.725,
                "desiredStatus": "RUNNING",
                "taskDefinitionArn": "arn:aws:ecs:ap-northeast-1:808634764529:task-definition/bitcoin-vvc:2",
                "containers": [
                    {
                        "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/22b286d7-0ff6-4414-94c7-34c75c72646c",
                        "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                        "lastStatus": "RUNNING",
                        "name": "node",
                        "networkBindings": []
                    },
                    {
                        "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/313cabba-d196-4f1f-9b8b-986d4b460fca",
                        "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                        "lastStatus": "RUNNING",
                        "name": "nginx",
                        "networkBindings": [
                            {
                                "protocol": "tcp",
                                "bindIP": "0.0.0.0",
                                "containerPort": 80,
                                "hostPort": 80
                            }
                        ]
                    },
                    {
                        "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/8a35a928-b659-4369-8e42-b2aade3a1488",
                        "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                        "lastStatus": "RUNNING",
                        "name": "golang",
                        "networkBindings": []
                    }
                ]
            }
        ]
    }`

	data["stopTasks"] = `{
        "task": {
            "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
            "overrides": {
                "containerOverrides": [
                    {
                        "name": "node"
                    },
                    {
                        "name": "nginx"
                    },
                    {
                        "name": "golang"
                    }
                ]
            },
            "lastStatus": "RUNNING",
            "containerInstanceArn": "arn:aws:ecs:ap-northeast-1:808634764528:container-instance/ade0ff2b-dd29-44e7-a8b9-7b3d02ff6876",
            "createdAt": 1453465153.893,
            "clusterArn": "arn:aws:ecs:ap-northeast-1:808634764528:cluster/default",
            "startedAt": 1453465161.725,
            "desiredStatus": "STOPPED",
            "stoppedReason": "Task stopped by user",
            "taskDefinitionArn": "arn:aws:ecs:ap-northeast-1:808634764528:task-definition/bitcoin-vvc:2",
            "containers": [
                {
                    "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/22b286d7-0ff6-4414-94c7-34c75c72646c",
                    "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                    "lastStatus": "RUNNING",
                    "name": "node",
                    "networkBindings": []
                },
                {
                    "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/313cabba-d196-4f1f-9b8b-986d4b460fca",
                    "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                    "lastStatus": "RUNNING",
                    "name": "nginx",
                    "networkBindings": [
                        {
                            "protocol": "tcp",
                            "bindIP": "0.0.0.0",
                            "containerPort": 80,
                            "hostPort": 80
                        }
                    ]
                },
                {
                    "containerArn": "arn:aws:ecs:ap-northeast-1:808634764528:container/8a35a928-b659-4369-8e42-b2aade3a1488",
                    "taskArn": "arn:aws:ecs:ap-northeast-1:808634764528:task/6fcb945f-a110-487b-aee5-79b44c9e9606",
                    "lastStatus": "RUNNING",
                    "name": "golang",
                    "networkBindings": []
                }
            ]
        }
    }`

	return data[key]
}
