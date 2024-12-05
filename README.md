# ☁️Cloud-term-project

2024년 충북대학교 클라우드컴퓨팅 Term-Project입니다.<br>
주제: **AWS동적관리 프로그램**

---

## 🚀 기능
- instance
 - list instance
 - start instance
 - create instance
 - reboot instance
- image
  - list images
- info
  - available regions
  - available zones
---

## ⚙️ 실행환경
- Alpine Linux v3.20
- golang v1.22
- aws-sdk-go-v2
---

## 📝 설치 및 실행
- image 빌드
```
docker build -t aws .
```

- 실행
```
docker run -it aws /bin/sh
```

- 컨테이너 내부 터미널이 ANSI를 지원하도록 설정(프로그램의 출력이 제대로 되지 않을경우)
```
export TERM=xterm-256color
```

- 프로그램 실행
```
./app // 사용가능한 명령어 확인가능
./app cli 
```

---

## 🎥 실행 영상
![실행3](https://github.com/user-attachments/assets/48c09e3f-691e-415b-b97c-efbbf5014b3a)
