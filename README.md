# AiBoard

AiBoard is an Android quota dashboard for OpenAI, MiniMax, and DeepSeek API accounts. The project uses Go for provider logic and Android Kotlin/Compose for the installable app shell.

## Current Features

- Go quota core with fixed five-hour windows and weekly reset windows.
- OpenAI usage aggregation adapter for `/organization/usage/completions`.
- DeepSeek balance adapter for `/user/balance`.
- MiniMax local fallback adapter for five-hour and weekly quota windows.
- Android Compose dashboard with portrait banner carousel and landscape dashboard mode.
- Gravity sensor controller that switches dashboard posture between portrait and landscape.
- Settings UI scaffold for multiple provider configurations with remark and masked API key.

## Build Go Core

```powershell
$env:GOCACHE='C:\Users\tp\Documents\AiBoard\.gocache'
go test ./go/aiboard/...
```

## Generate Android AAR From Go

Install gomobile, initialize it, then generate the AAR:

```powershell
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
gomobile bind -target=android -o android/app/libs/aiboard.aar ./go/aiboard/mobile
```

## Build Android APK

After installing Android Studio or a compatible Android SDK and Gradle wrapper:

```powershell
gradle :android:app:assembleDebug
```

If you add a Gradle wrapper, use:

```powershell
.\gradlew :android:app:assembleDebug
```

## Provider Notes

OpenAI usage APIs require organization/admin-capable credentials. A normal model API key may fail with a permission error.

DeepSeek balance is available through the official balance endpoint.

MiniMax quota is implemented as a local fallback because a stable public quota/balance endpoint was not verified during design. The provider boundary is ready for an official MiniMax adapter later.
"# AiBoard" 
