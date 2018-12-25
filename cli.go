package main

import (
	"./util"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "Linux cli"
	//程序的用途描述
	app.Usage = "Linux 安装基本软件"
	//程序的版本号
	app.Version = "1.0.0"

	//预置变量
	var java string
	var tomcat string
	var nginx string
	var mariadb string
	var mariadbUser string
	var mariadbPwd string
	var mariadbPort string
	var redis string
	var redisPwd string
	var redisPort string
	var mongoDB string
	var mongoDBPwd string
	var mongoDBPort string
	var mongoDBUser string
	var port string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "java, jn",
			Value:       "openjdk-8-jdk",
			Usage:       "Java 包名",
			Destination: &java,
		},
		cli.StringFlag{
			Name:        "tomcat, tn",
			Value:       "tomcat8",
			Usage:       "Tomcat 包名",
			Destination: &tomcat,
		},
		cli.StringFlag{
			Name:        "nginx, nn",
			Value:       "nginx",
			Usage:       "Nginx 包名",
			Destination: &nginx,
		},
		cli.StringFlag{
			Name:        "mariadb, mn",
			Value:       "mariadb-server-10.1",
			Usage:       "Mariadb 包名",
			Destination: &mariadb,
		},
		cli.StringFlag{
			Name:        "mariadb-port, mpt",
			Value:       "3306",
			Usage:       "Mariadb 端口",
			Destination: &mariadbPort,
		},
		cli.StringFlag{
			Name:        "mariadb-user, mu",
			Value:       "admin",
			Usage:       "Mariadb 用户",
			Destination: &mariadbUser,
		},
		cli.StringFlag{
			Name:        "mariadb-pwd, mp",
			Value:       "xiaoyi",
			Usage:       "Mariadb 密码",
			Destination: &mariadbPwd,
		},
		cli.StringFlag{
			Name:        "redis, rn",
			Value:       "redis-server",
			Usage:       "Redis 包名",
			Destination: &redis,
		},
		cli.StringFlag{
			Name:        "redis-pwd, rp",
			Value:       "xiaoyi",
			Usage:       "Redis 密码",
			Destination: &redisPwd,
		},
		cli.StringFlag{
			Name:        "redis-port, rpt",
			Value:       "6379",
			Usage:       "Redis 端口",
			Destination: &redisPort,
		},
		cli.StringFlag{
			Name:        "mongoDB, mgn",
			Value:       "mongodb",
			Usage:       "MongoDB 包名",
			Destination: &mongoDB,
		},
		cli.StringFlag{
			Name:        "mongoDB-user, mgu",
			Value:       "root",
			Usage:       "MongoDB 管理员账号",
			Destination: &mongoDBUser,
		},
		cli.StringFlag{
			Name:        "mongoDB-pwd, mgp",
			Value:       "xiaoyi",
			Usage:       "MongoDB 密码",
			Destination: &mongoDBPwd,
		},
		cli.StringFlag{
			Name:        "mongoDB-port, mgpt",
			Value:       "27017",
			Usage:       "MongoDB 端口",
			Destination: &mongoDBPort,
		},
		cli.StringFlag{
			Name:        "port, p",
			Value:       "80",
			Usage:       "端口",
			Destination: &port,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "java-install",
			Aliases: []string{"ji"},
			Usage:   "安装 Java",
			Action: func(c *cli.Context) error {
				util.Install(java)
				util.ExecSys("java", "-version")
				util.PrintlnRight("Java 安装成功")
				return nil
			},
		},
		{
			Name:    "java-remove",
			Aliases: []string{"jr"},
			Usage:   "拆卸 Java",
			Action: func(c *cli.Context) error {
				util.Remove(java)
				util.ExecSys("java", "-version")
				util.PrintlnRight("Java 拆卸成功")
				return nil
			},
		},
		{
			Name:    "tomcat-install",
			Aliases: []string{"ti"},
			Usage:   "安装 Tomcat",
			Action: func(c *cli.Context) error {
				util.Install(tomcat)
				util.PrintlnRight("Tomcat 安装成功")
				util.TomcatTip()
				return nil
			},
		},
		{
			Name:    "tomcat-remove",
			Aliases: []string{"tr"},
			Usage:   "拆卸 Tomcat",
			Action: func(c *cli.Context) error {
				util.Remove(tomcat)
				util.PrintlnRight("Tomcat 拆卸成功")
				return nil
			},
		},
		{
			Name:    "tomcat-tips",
			Aliases: []string{"tt"},
			Usage:   "提示 Tomcat",
			Action: func(c *cli.Context) error {
				util.TomcatTip()
				return nil
			},
		},
		{
			Name:    "nginx-install",
			Aliases: []string{"ni"},
			Usage:   "安装 Nginx",
			Action: func(c *cli.Context) error {
				util.Install(nginx)
				util.PrintlnTips("Nginx 安装成功")
				util.NginxTip()
				return nil
			},
		},
		{
			Name:    "nginx-remove",
			Aliases: []string{"nr"},
			Usage:   "拆卸 Nginx",
			Action: func(c *cli.Context) error {
				util.Remove(nginx)
				util.PrintlnTips("Nginx 拆卸成功")
				return nil
			},
		},
		{
			Name:    "nginx-tips",
			Aliases: []string{"nt"},
			Usage:   "提示 Nginx",
			Action: func(c *cli.Context) error {
				util.NginxTip()
				return nil
			},
		},
		{
			Name:    "mariadb-install",
			Aliases: []string{"mi"},
			Usage:   "安装 Mariadb",
			Action: func(c *cli.Context) error {
				util.Install(mariadb)
				util.PrintlnTips("Mariadb 安装成功")
				util.MaraidbConf(mariadbPort)
				util.MariadbTip(mariadbUser, mariadbPwd)
				return nil
			},
		},
		{
			Name:    "mariadb-remove",
			Aliases: []string{"mr"},
			Usage:   "拆卸 Mariadb",
			Action: func(c *cli.Context) error {
				util.Remove(mariadb)
				util.PrintlnTips("Mariadb 拆卸成功")
				return nil
			},
		},
		{
			Name:    "mariadb-tips",
			Aliases: []string{"mt"},
			Usage:   "提示 Mariadb",
			Action: func(c *cli.Context) error {
				util.MariadbTip(mariadbUser, mariadbPwd)
				return nil
			},
		},
		{
			Name:    "redis-install",
			Aliases: []string{"ri"},
			Usage:   "安装 Redis",
			Action: func(c *cli.Context) error {
				util.Install(redis)
				util.RedisConf(redisPwd, redisPort)
				util.PrintlnTips("Redis 安装成功")
				util.RedisTip()
				return nil
			},
		},
		{
			Name:    "redis-remove",
			Aliases: []string{"rr"},
			Usage:   "拆卸 Redis",
			Action: func(c *cli.Context) error {
				util.Remove(redis)
				util.PrintlnTips("Redis 拆卸成功")
				return nil
			},
		},
		{
			Name:    "redis-tips",
			Aliases: []string{"rt"},
			Usage:   "提示 Redis",
			Action: func(c *cli.Context) error {
				util.RedisTip()
				return nil
			},
		},
		{
			Name:    "mongoDB-install",
			Aliases: []string{"mgi"},
			Usage:   "安装 MongoDB",
			Action: func(c *cli.Context) error {
				util.Install(mongoDB)
				util.MongoDBConf(mongoDBPort)
				util.PrintlnTips("MongoDB 安装成功")
				util.MongoDBTip(mongoDBUser, mongoDBPwd)
				return nil
			},
		},
		{
			Name:    "mongoDB-remove",
			Aliases: []string{"mgr"},
			Usage:   "拆卸 MongoDB",
			Action: func(c *cli.Context) error {
				util.Remove(mongoDB)
				util.PrintlnTips("MongoDB 拆卸成功")
				return nil
			},
		},
		{
			Name:    "mongoDB-tips",
			Aliases: []string{"mgt"},
			Usage:   "提示 MongoDB",
			Action: func(c *cli.Context) error {
				util.MongoDBTip(mongoDBUser, mongoDBPwd)
				return nil
			},
		},
		{
			Name:    "port-info",
			Aliases: []string{"pi"},
			Usage:   "查看端口进程情况",
			Action: func(c *cli.Context) error {
				util.PortInfo(port)
				return nil
			},
		},
		{
			Name:    "port-kill",
			Aliases: []string{"pk"},
			Usage:   "根据端口kill进程",
			Action: func(c *cli.Context) error {
				util.KillPort(port)
				return nil
			},
		},
		{
			Name:    "firewall-tips",
			Aliases: []string{"ft"},
			Usage:   "防火墙提示信息",
			Action: func(c *cli.Context) error {
				util.FirewallTips()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
