#!/usr/bin/python
import os

container_name = "goji"
host_port = "8003"

cmd = "docker ps | grep %s"%container_name
container_exists = len(os.popen(cmd).read().split('\n')) > 1
print cmd
print "exists:", container_exists
if container_exists:
    cmd = "go build && docker restart %s"%container_name
    print cmd
    os.system(cmd)
else:
    cmd = "docker run -d -p %(port)s:80 -v %(cwd)s:/go/src --name %(name)s dev"%({
        "port":host_port,
        "cwd":os.getcwd(),
        "name":container_name
    })
    print cmd
    os.system(cmd)
