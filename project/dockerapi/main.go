package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	_ "github.com/lib/pq"
)

const (
	// 连接 PostgreSQL 的参数
	dbHost     = "localhost"
	dbPort     = "25432"
	dbName     = "postgres"
	dbUser     = "postgres"
	dbPassword = "123456"
)

func main() {
	ctx := context.Background()

	// 初始化 Docker 客户端
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// 指定要使用的镜像和容器名称
	image := "postgres:latest"
	containerName := "my-postgres-db"

	// 创建容器配置
	config := &container.Config{
		Image: image,
		Env: []string{
			"POSTGRES_DB=" + dbName,
			"POSTGRES_USER=" + dbUser,
			"POSTGRES_PASSWORD=" + dbPassword,
		},
		ExposedPorts: map[nat.Port]struct{}{
			"5432/tcp": {},
		},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{HostIP: "0.0.0.0", HostPort: dbPort},
			},
		},
		Binds: []string{"/mnt/data:/var/lib/postgresql/data"},
	}

	// 启动 PostgreSQL 容器实例
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// 等待 PostgreSQL 服务启动完毕
	waitContainerRunning(ctx, cli, resp.ID) //nolint:errcheck

	// 连接到 PostgreSQL 数据库并创建 mydatabase 数据库
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	time.Sleep(2 * time.Second)
	_, err = db.Exec("DROP DATABASE IF EXISTS mydatabase")
	if err != nil {
		panic(err)
	}
	//创建 mydatabase 数据库
	_, err = db.Exec("CREATE DATABASE mydatabase")
	if err != nil {
		panic(err)
	}

	fmt.Println("mydatabase has been created successfully.")
}

func waitContainerRunning(ctx context.Context, cli *client.Client, containerID string) error {
	for {
		resp, err := cli.ContainerInspect(ctx, containerID)
		if err != nil {
			return err
		}
		if !resp.State.Running {
			return fmt.Errorf("container %s is not running", containerID)
		}
		break
	}
	return nil
}
