package goroslib

import (
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ROBTTO/goroslib/pkg/msgs/sensor_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type containerMaster struct {
	ip string
}

func newContainerMaster() (*containerMaster, error) {
	exec.Command("docker", "kill", "goroslib-test-master").Run()
	exec.Command("docker", "wait", "goroslib-test-master").Run()
	exec.Command("docker", "rm", "goroslib-test-master").Run()

	cmd := []string{"docker", "run", "--rm", "-d", "--name=goroslib-test-master"}
	cmd = append(cmd, "goroslib-test-master")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		return nil, err
	}

	// get master ip
	byts, _ := exec.Command("docker", "inspect", "-f",
		"{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
		"goroslib-test-master").Output()
	ip := string(byts[:len(byts)-1])

	// wait for master
	address := ip + ":" + strconv.FormatInt(11311, 10)
	for {
		time.Sleep(1 * time.Second)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		break
	}

	return &containerMaster{
		ip: ip,
	}, nil
}

func (m *containerMaster) IP() string {
	return m.ip
}

func (m *containerMaster) close() {
	exec.Command("docker", "kill", "goroslib-test-master").Run()
	exec.Command("docker", "wait", "goroslib-test-master").Run()
}

type container struct {
	name string
}

func newContainer(name string, masterIP string) (*container, error) {
	exec.Command("docker", "kill", "goroslib-test-"+name).Run()
	exec.Command("docker", "wait", "goroslib-test-"+name).Run()
	exec.Command("docker", "rm", "goroslib-test-"+name).Run()

	cmd := []string{"docker", "run", "-d", "--name=goroslib-test-" + name}
	cmd = append(cmd, "-e", "MASTER_IP="+masterIP)
	cmd = append(cmd, "goroslib-test-"+name)
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		return nil, err
	}

	// wait for node initialization
	time.Sleep(1 * time.Second)

	return &container{
		name: name,
	}, nil
}

func (c *container) close() {
	exec.Command("docker", "kill", "goroslib-test-"+c.name).Run()
	exec.Command("docker", "wait", "goroslib-test-"+c.name).Run()
	exec.Command("docker", "rm", "goroslib-test-"+c.name).Run()
}

func (c *container) waitOutput() string {
	exec.Command("docker", "wait", "goroslib-test-"+c.name).Run()
	out, _ := exec.Command("docker", "logs", "goroslib-test-"+c.name).Output()
	exec.Command("docker", "rm", "goroslib-test-"+c.name).Run()
	return string(out)
}

func TestNodeMasterOpen(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		m, err := newContainerMaster()
		require.NoError(t, err)
		defer m.close()

		n, err := NewNode(NodeConf{
			Namespace:     "/myns",
			Name:          "goroslib1",
			MasterAddress: m.IP() + ":11311",
		})
		require.NoError(t, err)
		defer n.Close()
	})

	t.Run("http", func(t *testing.T) {
		m, err := newContainerMaster()
		require.NoError(t, err)
		defer m.close()

		n, err := NewNode(NodeConf{
			Namespace:     "/myns",
			Name:          "goroslib2",
			MasterAddress: "http://" + m.IP() + ":11311",
		})
		require.NoError(t, err)
		defer n.Close()
	})

	t.Run("not running", func(t *testing.T) {
		_, err := NewNode(NodeConf{
			Namespace:     "/myns",
			Name:          "goroslib",
			MasterAddress: "127.0.0.1:11311",
		})
		require.Error(t, err)
	})
}

func TestNodeNamespace(t *testing.T) {
	t.Run("from environment", func(t *testing.T) {
		os.Setenv("ROS_NAMESPACE", "/myns")
		defer os.Unsetenv("ROS_NAMESPACE")

		m, err := newContainerMaster()
		require.NoError(t, err)
		defer m.close()

		n, err := NewNode(NodeConf{
			Name:          "goroslib1",
			MasterAddress: m.IP() + ":11311",
		})
		require.NoError(t, err)
		defer n.Close()

		require.Equal(t, "/myns", n.conf.Namespace)
	})

	t.Run("from environment and conf", func(t *testing.T) {
		os.Setenv("ROS_NAMESPACE", "/myns1")
		defer os.Unsetenv("ROS_NAMESPACE")

		m, err := newContainerMaster()
		require.NoError(t, err)
		defer m.close()

		n, err := NewNode(NodeConf{
			Name:          "goroslib1",
			Namespace:     "/myns2",
			MasterAddress: m.IP() + ":11311",
		})
		require.NoError(t, err)
		defer n.Close()

		require.Equal(t, "/myns2", n.conf.Namespace)
	})
}

func TestNodeRosnodeInfo(t *testing.T) {
	m, err := newContainerMaster()
	require.NoError(t, err)
	defer m.close()

	n, err := NewNode(NodeConf{
		Namespace:     "/myns",
		Name:          "goroslib",
		MasterAddress: m.IP() + ":11311",
	})
	require.NoError(t, err)
	defer n.Close()

	// a publisher/subscriber is needed for the node to be listed
	// standard nodes creates some default publishers, subscribers and
	// services during initialization, while goroslib does not do this yet
	pub, err := NewPublisher(PublisherConf{
		Node:  n,
		Topic: "test_topic",
		Msg:   &std_msgs.Float64{},
	})
	require.NoError(t, err)
	defer pub.Close()

	sub, err := NewSubscriber(SubscriberConf{
		Node:     n,
		Topic:    "test_topic",
		Callback: func(msg *sensor_msgs.Imu) {},
	})
	require.NoError(t, err)
	defer sub.Close()

	sp, err := NewServiceProvider(ServiceProviderConf{
		Node: n,
		Name: "test_srv",
		Srv:  &TestService{},
		Callback: func(req *TestServiceReq) *TestServiceRes {
			return &TestServiceRes{}
		},
	})
	require.NoError(t, err)
	defer sp.Close()

	rt, err := newContainer("rosnode-info", m.IP())
	require.NoError(t, err)

	require.Regexp(t, regexp.MustCompile(
		"^--------------------------------------------------------------------------------\n"+
			"Node \\[/myns/goroslib\\]\n"+
			"Publications: \n"+
			" \\* /myns/test_topic \\[std_msgs/Float64\\]\n"+
			" \\* /rosout \\[rosgraph_msgs/Log\\]\n"+
			"\n"+
			"Subscriptions: \n"+
			" \\* /myns/test_topic \\[std_msgs/Float64\\]\n"+
			"\n"+
			"Services: \n"+
			" \\* /myns/test_srv\n"+
			"\n"+
			"\n"+
			"contacting node http://.+? ...\n"+
			"Pid: [0-9]+\n"+
			"Connections:\n"+
			" \\* topic: /rosout\n"+
			"    \\* to: /rosout\n"+
			"    \\* direction: outbound\n"+
			"    \\* transport: TCPROS\n"+
			" \\* topic: /myns/test_topic\n"+
			"    \\* to: http://.+?/\n"+
			"    \\* direction: inbound\n"+
			"    \\* transport: TCPROS\n"+
			"\n$"), rt.waitOutput())
}
