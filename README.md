# AndroidLibXrayLite

Change v2ray-core to [https-proxy](https://github.com/justlovediaodiao/https-proxy). Buiding with [v2rayNG](https://github.com/2dust/v2rayNG) to make https-proxy runs on Android.

## Build requirements
* JDK
* Android SDK
* Go
* gomobile


### Build environment

- Go Mobile:
```bash
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
go get golang.org/x/mobile/bind
```
Add `~/go/bin` to environment path.

- Android SDK
1. install JDK([Java SE Development Kit](https://www.oracle.com/java/technologies/downloads/)). Add `<jdk>/bin` to environment path.
2. install [sdkmanager](https://developer.android.google.cn/studio/command-line/sdkmanager).  Add `<androidsdk>/cmdline-tools/latest/bin` to environment path.
3. list available sdk packages:
```bash
sdkmanager --list
```
4. select latest SDK(`platforms;android-*`) and NDK(`nkd;*`) (* is version number) to install.
```
sdkmanager "platforms;android-*" "ndk;*"
```

Add `ANDROID_HOME=<androidsdk>` environment variable.


## Build instructions
1. `git clone [repo] && cd AndroidLibXrayLite`
2. `gomobile bind -v -androidapi 21 -ldflags='-s -w' ./`
