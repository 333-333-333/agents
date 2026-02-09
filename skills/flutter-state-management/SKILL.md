---
name: flutter-state-management
description: >
  Riverpod state management patterns for Flutter: providers, notifiers, async state, and clean architecture integration.
  Trigger: When creating Riverpod providers, managing async state, or connecting UI to business logic in the mobile app.
metadata:
  author: 333-333-333
  version: "1.1"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Creating Riverpod providers or notifiers"
    - "Managing async state in Flutter"
    - "Connecting Flutter UI to business logic"
    - "Handling loading and error states"
---

## When to Use

- Creating any Riverpod provider or notifier
- Managing async data (API calls, Firebase queries)
- Handling loading/error/data states in UI
- Connecting presentation layer to use cases

## Provider Types — Decision Tree

```
What kind of state do you need?

Computed value from other providers?      → Provider (read-only)
Simple sync state with mutations?         → NotifierProvider
Async data (API, DB) with mutations?      → AsyncNotifierProvider
Async data (read-only, auto-fetch)?       → FutureProvider
Real-time stream (WebSocket, Firestore)?  → StreamProvider
```

## Patterns

### 1. Read-Only Computed Provider

For derived/computed values. No mutations.

> **Example**: See [assets/provider_read_only.dart](assets/provider_read_only.dart)

### 2. Notifier — Sync State with Mutations

For local, synchronous state that needs to be mutated.

> **Example**: See [assets/notifier_sync.dart](assets/notifier_sync.dart)

### 3. AsyncNotifier — Async State with Mutations (MOST COMMON)

For data that comes from an API/DB and can be mutated. Uses `AsyncValue.guard()` for safe state transitions.

> **Example**: See [assets/async_notifier_crud.dart](assets/async_notifier_crud.dart)

### 4. FutureProvider — Read-Only Async Data

For async data you only need to FETCH, not mutate. Supports `.family` for parameterized queries.

> **Example**: See [assets/future_provider.dart](assets/future_provider.dart)

### 5. StreamProvider — Real-Time Data

For Firestore listeners, WebSocket streams, etc.

> **Example**: See [assets/stream_provider.dart](assets/stream_provider.dart)

## Consuming Providers in UI

### Handle AsyncValue (CRITICAL PATTERN)

Every async provider returns `AsyncValue<T>`. ALWAYS handle all 3 states: loading, error, data. Two approaches: `.when()` method or Dart 3.0 pattern matching.

> **Example**: See [assets/async_value_consumer.dart](assets/async_value_consumer.dart)

## Clean Architecture Integration

```
Page (ConsumerWidget)
  → watches provider
    → provider uses use case
      → use case calls repository interface (domain)
        → repository impl (infrastructure) calls API
```

### Wiring — Dependency Injection

Repository providers and use case providers connect clean architecture layers via Riverpod's `Provider`.

> **Example**: See [assets/di_wiring.dart](assets/di_wiring.dart)

## autoDispose vs Keep Alive

| Scenario | Use |
|----------|-----|
| Screen-level data (detail pages) | `autoDispose` — clean up when leaving |
| Global state (auth, user session) | Keep alive (no autoDispose) |
| Lists that are expensive to re-fetch | Keep alive or cache manually |
| Form state | `autoDispose` — reset on navigation |

> **Example**: See [assets/auto_dispose.dart](assets/auto_dispose.dart)

## Naming Conventions

| Element | Pattern | Example |
|---------|---------|---------|
| Read-only provider | `{thing}Provider` | `currentUserProvider` |
| Notifier class | `{Thing}Notifier` | `BookingFilterNotifier` |
| NotifierProvider | `{thing}Provider` | `bookingFilterProvider` |
| AsyncNotifier class | `{Thing}Notifier` | `PetListNotifier` |
| FutureProvider | `{thing}Provider` | `caregiverDetailProvider` |
| StreamProvider | `{thing}Provider` | `chatMessagesProvider` |
| Repository provider | `{feature}RepositoryProvider` | `authRepositoryProvider` |
| Use case provider | `{useCase}Provider` | `loginUseCaseProvider` |

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| `ref.read` in `build()` method | Use `ref.watch` to rebuild on changes |
| Business logic inside the widget | Put logic in Notifier or UseCase |
| Ignore loading/error states | Always handle all 3 AsyncValue states |
| One giant provider file for everything | One provider file per feature |
| Call API directly from provider | Provider → UseCase → Repository |
| Forget `autoDispose` on detail pages | Use autoDispose for screen-scoped data |

## Resources

- **Templates**: See [assets/](assets/) for provider patterns, consumer examples, and DI wiring
