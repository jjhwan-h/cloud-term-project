# ☁️Cloud-term-project

2024년 충북대학교 클라우드컴퓨팅 Term-Project입니다.<br>
주제: **AWS동적관리 프로그램**

![cloud1](https://github.com/user-attachments/assets/f5d3b115-0e39-4640-8659-59b153bc06dc)
---

## 🚀 기능
- instance
  - list instance
  - start instance
  - stop instance
  - create instance
  - reboot instance
  - connect instance
- image
  - list images
- info
  - available regions
  - available zones
---

## ⚙️ 실행환경
- Alpine Linux v3.20
- golang v1.22
- aws-sdk-go-v2 v1.32.6
---

## 📝 설치 및 실행
- AWS API key 필요
- .dev.env
```
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_REGION=
AWS_OWNER_ID=
PRIVATE_KEY_PATH=/root/cloud-test.pem // docker volume으로 /root/cloud-test.pem에 설정
USER=ec2-user // ssh로 접속할 instance의 user명 (대부분 ec2-user나 ubuntu)
```
- image 빌드
```
docker build -t aws .
```

- 실행
```
export LOCAL_PRIVATE_KEY_PATH=/local/private/key/path //ssh 연결을위한 key
export LOCAL_ENV_PATH=/local/env/path //.env

./start.sh
```

---
