# AI Quota Android Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build an installable Android project skeleton with a tested Go core for OpenAI, MiniMax, and DeepSeek quota display.

**Architecture:** Go provides quota windows, provider adapters, normalized snapshots, and a gomobile bridge. Android provides native configuration screens, banner carousel UI, sensor/orientation behavior, secure key storage boundary, and APK packaging.

**Tech Stack:** Go 1.22+, Android Gradle Plugin, Kotlin, Jetpack Compose, gomobile AAR integration.

---

### Task 1: Go Quota Core

**Files:**
- Create: `go.mod`
- Create: `go/aiboard/model/types.go`
- Create: `go/aiboard/quota/window.go`
- Test: `go/aiboard/quota/window_test.go`

- [x] Write failing tests for fixed five-hour windows and weekly windows.
- [x] Implement quota window calculation.
- [x] Run `go test ./go/aiboard/quota`.

### Task 2: Provider Adapters

**Files:**
- Create: `go/aiboard/provider/client.go`
- Create: `go/aiboard/provider/deepseek.go`
- Create: `go/aiboard/provider/minimax.go`
- Create: `go/aiboard/provider/openai.go`
- Test: `go/aiboard/provider/provider_test.go`

- [x] Write failing tests for DeepSeek balance parsing, MiniMax local fallback, and OpenAI aggregation.
- [x] Implement provider adapters with dependency-injected HTTP clients.
- [x] Run `go test ./go/aiboard/provider`.

### Task 3: Mobile Bridge

**Files:**
- Create: `go/aiboard/mobile/bridge.go`

- [x] Expose JSON-based bridge methods suitable for gomobile binding.
- [x] Keep API keys out of persisted Go state.

### Task 4: Android Skeleton

**Files:**
- Create: `settings.gradle.kts`
- Create: `build.gradle.kts`
- Create: `android/app/build.gradle.kts`
- Create: `android/app/src/main/AndroidManifest.xml`
- Create: `android/app/src/main/java/com/aiboard/app/MainActivity.kt`
- Create: `android/app/src/main/java/com/aiboard/app/AiBoardUi.kt`
- Create: `android/app/src/main/java/com/aiboard/app/SensorController.kt`
- Create: `android/app/src/main/res/values/strings.xml`
- Create: `android/app/src/main/res/values/colors.xml`

- [x] Add installable Android app configuration.
- [x] Add full-sensor orientation support.
- [x] Add banner carousel and settings UI source.

### Task 5: Verification

**Files:**
- Read: all generated files

- [x] Run Go tests.
- [x] Check for Gradle wrapper availability.
- [x] Report any local build limits clearly.
