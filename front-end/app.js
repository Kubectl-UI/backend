// const str = "Name:         pod-app\nNamespace:    default\nPriority:     0\nNode:         docker-desktop/192.168.65.4\nStart Time:   Sun, 27 Jun 2021 01:03:40 +0100\nLabels:       app=pod-app\nAnnotations:  <none>\nStatus:       Running\nIP:           10.1.1.139\nIPs:\n  IP:  10.1.1.139\nContainers:\n  k8s-app:\n    Container ID:   docker://baf0e8ef85762d08441994dc392e20797ebd14a580f1151d0cada3b4cdbf57c5\n    Image:          phirmware/k8s-app:latest\n    Image ID:       docker-pullable://phirmware/k8s-app@sha256:0685f9e0143ac18d70822bd1eb1edcf127aae3ad14fb7d1fca10be08117b0c1a\n    Port:           5000/TCP\n    Host Port:      0/TCP\n    State:          Running\n      Started:      Sun, 27 Jun 2021 01:03:50 +0100\n    Ready:          True\n    Restart Count:  0\n    Limits:\n      cpu:     200m\n      memory:  500Mi\n    Requests:\n      cpu:     100m\n      memory:  200Mi\n    Environment:\n      PORT:  5000\n    Mounts:\n      /var/run/secrets/kubernetes.io/serviceaccount from default-token-4d9jq (ro)\n  k8s-ping:\n    Container ID:   docker://6f5870f9259fc09fbb1845ab5016de9cafca416c1f45ba62f8e3e3925de02e01\n    Image:          phirmware/ping:latest\n    Image ID:       docker-pullable://phirmware/ping@sha256:991a6af84cf8749e2febc9f7008065fa47fa5a73ec0721d3deb51000cef883ea\n    Port:           <none>\n    Host Port:      <none>\n    State:          Running\n      Started:      Sun, 27 Jun 2021 01:04:24 +0100\n    Ready:          True\n    Restart Count:  0\n    Limits:\n      cpu:     200m\n      memory:  500Mi\n    Requests:\n      cpu:     100m\n      memory:  200Mi\n    Environment:\n      URL:  http://localhost:5000\n    Mounts:\n      /var/run/secrets/kubernetes.io/serviceaccount from default-token-4d9jq (ro)\nConditions:\n  Type              Status\n  Initialized       True \n  Ready             True \n  ContainersReady   True \n  PodScheduled      True \nVolumes:\n  default-token-4d9jq:\n    Type:        Secret (a volume populated by a Secret)\n    SecretName:  default-token-4d9jq\n    Optional:    false\nQoS Class:       Burstable\nNode-Selectors:  <none>\nTolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s\n                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s\nEvents:\n  Type     Reason     Age                From               Message\n  ----     ------     ----               ----               -------\n  Normal   Scheduled  86s                default-scheduler  Successfully assigned default/pod-app to docker-desktop\n  Normal   Pulling    84s                kubelet            Pulling image \"phirmware/k8s-app:latest\"\n  Normal   Pulled     77s                kubelet            Successfully pulled image \"phirmware/k8s-app:latest\" in 6.7994223s\n  Normal   Created    77s                kubelet            Created container k8s-app\n  Normal   Started    76s                kubelet            Started container k8s-app\n  Warning  Failed     62s                kubelet            Error: ErrImagePull\n  Warning  Failed     62s                kubelet            Failed to pull image \"phirmware/ping:latest\": rpc error: code = Unknown desc = Error response from daemon: Head https://registry-1.docker.io/v2/phirmware/ping/manifests/latest: Get https://auth.docker.io/token?scope=repository%3Aphirmware%2Fping%3Apull&service=registry.docker.io: net/http: TLS handshake timeout\n  Normal   BackOff    61s                kubelet            Back-off pulling image \"phirmware/ping:latest\"\n  Warning  Failed     61s                kubelet            Error: ImagePullBackOff\n  Normal   Pulling    50s (x2 over 76s)  kubelet            Pulling image \"phirmware/ping:latest\"\n  Normal   Pulled     43s                kubelet            Successfully pulled image \"phirmware/ping:latest\" in 7.637206s\n  Normal   Created    42s                kubelet            Created container k8s-ping\n  Normal   Started    42s                kubelet            Started container k8s-ping\n"
const termional = document.getElementById('terminal')
// termional.innerHTML = str

const url = 'http://localhost:8080'

const getPodsButton = document.getElementById('get-pods')
const describePodButton = document.getElementById('describe-pod')
const podName = document.getElementById('pod-name')
const createPod = document.getElementById('create-pod')

getPodsButton.onclick = getPods
describePodButton.onclick = describePod
createPod.onclick = createPod


async function getPods() {
  const result = await fetch(`${url}/get-pods`)
  const body = await result.json()
  console.log(body)
  termional.innerHTML = body.Message
}

async function describePod() {
  if (!podName.value) return alert('Input pod name')
  const data = {
    PodName: podName.value
  }
  const result = await fetch(`${url}/describe-pod`, {
    method: 'post',
    body: JSON.stringify(data)
  })
  const body = await result.json()
  console.log(body)
  termional.innerHTML = body.Message
}

async function createPod() {
  
}