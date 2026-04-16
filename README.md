# mini-k8s

Go 언어 공부를 위해 직접 구현해보는 Kubernetes

## 폴더 구조

```shell
mini-k8s/
├── cmd/                  # 실행 가능한 메인 애플리케이션의 입구 (Entry point)
│   └── apiserver/
│       └── main.go       # 프로그램의 시작점 (main 패키지)
├── pkg/                  # 외부에서도 가져다 쓸 수 있는 라이브러리 코드
│   ├── registry/
│   │   └── registry.go   # 리소스 저장 및 관리 로직
│   └── resources/
│       ├── interface.go  # Resource 인터페이스 정의
│       └── pod.go        # Pod 구조체 및 메서드 정의
├── go.mod                # Go 모듈 정의 및 의존성 관리 파일 (C++의 CMakeLists, Python의 requirements와 유사)
└── go.sum                # 의존성 체크섬 파일 (자동 생성)
```

|디렉토리|역할|특징|
|:---|:---|:---|
| `cmd/` | 프로젝트의 바이너리를 빌드하기 위한 main 패키지들이 위치합니다. | 실제 로직보다는 pkg/의 함수들을 호출하고 설정하는 역할만 합니다.|
| `pkg/` | 다른 프로젝트에서도 import 해서 사용할 수 있는 공용 코드를 담습니다. | K8s의 client-go 같은 라이브러리들이 주로 여기에 위치합니다.|
