---
name: flutter-tdd
description: >
  Test-Driven Development for Flutter: red-green-refactor cycle for domain entities, value objects, use cases, repositories, and Riverpod notifiers.
  Trigger: When writing unit tests, implementing TDD, creating mocks, or testing business logic in the Flutter app.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Writing unit tests for Flutter domain logic"
    - "Testing use cases or application services"
    - "Creating mock repositories or datasources"
    - "Testing Riverpod providers or notifiers"
    - "Implementing TDD in Flutter"
    - "Adding test coverage to a Flutter feature"
---

## When to Use

- Starting implementation of any new feature (write test FIRST)
- Testing domain entities and value objects
- Testing use cases / application services
- Testing Riverpod notifiers and providers
- Creating mock implementations for repository interfaces
- Testing infrastructure adapters (API clients, local storage)

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Test FIRST** | Write the failing test before production code. Always. |
| **Red -> Green -> Refactor** | Fail -> pass -> clean up. Never skip a step. |
| **Mirror structure** | `test/` mirrors `lib/` exactly |
| **File naming** | `foo.dart` -> `foo_test.dart` |
| **Mocks via interfaces** | Domain repository interfaces ARE the mock boundary — no mockito needed for simple cases |
| **No test pollution** | Each test is independent, no shared mutable state |
| **group() for context** | Use `group()` to organize related tests |
| **Riverpod: use ProviderContainer** | Create isolated containers per test |

## The TDD Cycle

```
1. RED    — Write a test that describes the desired behavior. Run it. It MUST fail.
2. GREEN  — Write the MINIMUM code to make the test pass. Nothing more.
3. REFACTOR — Clean up while keeping tests green.
4. REPEAT
```

### When Writing Code with AI

```
1. DESCRIBE the behavior you want as a test
2. AI writes the failing test
3. Run it — confirm RED
4. AI writes the implementation
5. Run it — confirm GREEN
6. AI refactors if needed
7. Run it — confirm still GREEN
```

## Test Directory Structure

```
mobile/
  test/
    features/
      auth/
        domain/
          entities/
            user_test.dart
          value_objects/
            email_test.dart
            password_test.dart
        application/
          use_cases/
            login_use_case_test.dart
        presentation/
          providers/
            auth_state_provider_test.dart
        infrastructure/
          repositories/
            auth_repository_impl_test.dart
      booking/
        domain/
        application/
        presentation/
        infrastructure/
    shared/
      domain/
        value_objects/
          money_test.dart
    helpers/
      mocks.dart            # Shared mock implementations
      test_providers.dart   # Riverpod test overrides
```

## Domain Entity Tests

> See [assets/booking_test.dart](assets/booking_test.dart)

## Value Object Tests

> See [assets/email_test.dart](assets/email_test.dart)

## Mock Pattern (Manual — No Framework)

> See [assets/mocks.dart](assets/mocks.dart)

For more complex mocking needs, use `mocktail`:

> See [assets/mocks_mocktail.dart](assets/mocks_mocktail.dart)

## Use Case Tests

> See [assets/login_use_case_test.dart](assets/login_use_case_test.dart)

## Riverpod Provider/Notifier Tests

> See [assets/auth_state_provider_test.dart](assets/auth_state_provider_test.dart)

## AsyncNotifier Tests

> See [assets/async_notifier_test.dart](assets/async_notifier_test.dart)

## Either Pattern Tests

When using `fpdart` or `dartz` for `Either<Failure, T>`:

> See [assets/either_test.dart](assets/either_test.dart)

## Commands

```bash
# Run all unit tests
flutter test

# Run specific test file
flutter test test/features/auth/domain/entities/user_test.dart

# Run tests matching a name pattern
flutter test --name "login"

# Run with coverage
flutter test --coverage
# View coverage report (install lcov first: brew install lcov)
genhtml coverage/lcov.info -o coverage/html
open coverage/html/index.html

# Run tests in a specific directory
flutter test test/features/auth/

# Watch mode (re-run on changes) — requires flutter_test_watch or similar
# Not built-in; use: find test -name '*_test.dart' | entr flutter test
```

## Dev Dependencies

> See [assets/pubspec_dev_deps.yaml](assets/pubspec_dev_deps.yaml)

## Anti-Patterns

| Don't | Do |
|-------|-----|
| Write code first, tests later | Write test FIRST (red), then code (green) |
| Test implementation details | Test behavior — inputs and outputs |
| Shared mutable state between tests | Fresh `setUp()` per test |
| Import infrastructure in domain tests | Mock via interfaces |
| Skip `tearDown` for ProviderContainer | Always `container.dispose()` |
| One giant test function | `group()` + focused `test()` with descriptive names |
| Use `print()` for debugging tests | Use `expect()` with clear matchers |
| Mock everything | Only mock external boundaries (repos, APIs) |
