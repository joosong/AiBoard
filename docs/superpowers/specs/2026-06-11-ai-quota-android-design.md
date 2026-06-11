# AI Quota Android Design

## Goal

Build an installable Android application that monitors OpenAI, MiniMax, and DeepSeek API quota or balance data. The app uses Go for provider logic, quota windows, and data models, with an Android native shell for APK packaging, storage, sensors, orientation, and UI.

## Product Shape

The home screen is a rotating banner dashboard. Each banner represents one configured API account and shows provider-specific status:

- OpenAI: five-hour usage, weekly usage, cost summary, request/token totals, and reset time.
- MiniMax: five-hour and weekly quota windows using a local fallback adapter until a stable public quota API is available.
- DeepSeek: current balance, granted balance, topped-up balance, and last refresh state.

The settings screen manages multiple configurations. Each configuration has a provider, remark, and API key. API keys are stored by Android secure storage and are only passed into Go at refresh time.

## Architecture

Go owns deterministic business behavior:

- quota window calculation
- provider client interfaces
- OpenAI usage aggregation
- DeepSeek balance parsing
- MiniMax local quota fallback
- mobile bridge models

Android owns platform behavior:

- installable APK
- encrypted key storage
- orientation and gravity sensor
- banner carousel UI
- background refresh scheduling

The MVP keeps Android UI source lightweight and mock-friendly while the Go core is tested independently.

## Data And Error Handling

All provider refreshes return a normalized snapshot with status, usage windows, balance, reset times, last update time, and an optional error message. Provider-specific unsupported data is represented as absent fields, not fake values.

OpenAI usage endpoints require organization/admin-capable credentials. If a normal API key cannot access usage or cost data, the app displays a permission error for that account.

MiniMax quota support is intentionally adapter-based. The first version supports local configured windows and can later add an official API adapter without changing the UI contract.

## Verification

The Go core is verified with unit tests. Android verification starts with Gradle project integrity and source review because the local environment may not have Android Gradle dependencies installed.
