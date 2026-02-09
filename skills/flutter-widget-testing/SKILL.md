---
name: flutter-widget-testing
description: >
  Widget and component testing for Flutter: pumping widgets, finders, matchers, interaction simulation, golden tests, and accessibility checks.
  Trigger: When writing widget tests, testing UI components, creating golden tests, or verifying accessibility in Flutter.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Writing widget tests for Flutter components"
    - "Testing Flutter atoms, molecules, or organisms"
    - "Creating golden tests for visual regression"
    - "Testing user interactions in Flutter widgets"
    - "Verifying accessibility in widget tests"
    - "Testing pages or screens with navigation"
---

## When to Use

- Testing UI components (atoms, molecules, organisms, templates, pages)
- Testing widget rendering, layout, and text content
- Testing user interactions (tap, scroll, enter text, drag)
- Creating golden tests for visual regression
- Verifying accessibility (semantics, labels)
- Testing widgets that use Riverpod providers
- Testing navigation and routing behavior

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Test FIRST** | Write the failing widget test before the widget. Red -> Green -> Refactor. |
| **Wrap with app shell** | Always wrap tested widgets in `MaterialApp` or the Forui theme equivalent |
| **pump vs pumpAndSettle** | `pump()` advances one frame; `pumpAndSettle()` waits for all animations |
| **Finders are lazy** | `find.text('X')` doesn't search until used in `expect()` or `tester.tap()` |
| **Golden tests pinned** | Update goldens explicitly with `--update-goldens`; never auto-update in CI |
| **Test behavior, not pixels** | Prefer semantic finders (`find.text`, `find.byType`) over pixel coordinates |
| **Riverpod: use ProviderScope** | Wrap widget in `ProviderScope` with overrides for test isolation |

## Test Directory Structure

```
mobile/
  test/
    shared/
      presentation/
        atoms/
          app_button_test.dart
          app_text_field_test.dart
          app_badge_test.dart
        molecules/
          emphasis_headline_test.dart
        organisms/
          pet_emoji_carousel_test.dart
    features/
      auth/
        presentation/
          pages/
            welcome_page_test.dart
          widgets/
            auth_sheet_test.dart
    goldens/                    # Golden reference images
      shared/
        atoms/
          app_button_light.png
          app_button_dark.png
      features/
        auth/
          welcome_page_light.png
    helpers/
      pump_app.dart             # Shared widget wrapper helper
      test_theme.dart           # Theme setup for tests
```

## Widget Test Wrapper Helper

Every widget test needs a proper app shell with theme. Create a shared helper:

> See [assets/pump_app.dart](assets/pump_app.dart)

## Atom Tests (Simple Widgets)

> See [assets/app_button_test.dart](assets/app_button_test.dart)

## Text Field Tests

> See [assets/app_text_field_test.dart](assets/app_text_field_test.dart)

## Interaction Tests

> See [assets/interaction_tests.dart](assets/interaction_tests.dart)

## Page Tests with Riverpod

> See [assets/welcome_page_test.dart](assets/welcome_page_test.dart)

## Navigation Tests

> See [assets/navigation_test.dart](assets/navigation_test.dart)

## Golden Tests (Visual Regression)

> See [assets/golden_test.dart](assets/golden_test.dart)

## Accessibility Tests

> See [assets/accessibility_test.dart](assets/accessibility_test.dart)

## Finders Reference

| Finder | Use For |
|--------|---------|
| `find.text('X')` | Find widget displaying exact text |
| `find.textContaining('X')` | Find widget containing text |
| `find.byType(AppButton)` | Find widget by runtime type |
| `find.byKey(Key('x'))` | Find widget by key |
| `find.byIcon(Icons.home)` | Find widget by icon |
| `find.descendant(of: x, matching: y)` | Find widget nested inside another |
| `find.ancestor(of: x, matching: y)` | Find parent widget |
| `find.byWidgetPredicate((w) => ...)` | Custom predicate |
| `find.bySemanticsLabel('X')` | Find by accessibility label |

## Matchers Reference

| Matcher | Verifies |
|---------|----------|
| `findsOneWidget` | Exactly one match |
| `findsNothing` | No matches |
| `findsNWidgets(n)` | Exactly N matches |
| `findsAtLeastNWidgets(n)` | At least N matches |
| `matchesGoldenFile('path')` | Visual match against golden |

## Commands

```bash
# Run all widget tests
flutter test

# Run specific widget test
flutter test test/shared/presentation/atoms/app_button_test.dart

# Update golden files (run locally, commit the PNGs)
flutter test --update-goldens

# Update specific golden
flutter test --update-goldens test/shared/presentation/atoms/app_button_golden_test.dart

# Run with coverage
flutter test --coverage

# Run tests matching name
flutter test --name "AppButton"
```

## Dev Dependencies

> See [assets/pubspec_dev_deps.yaml](assets/pubspec_dev_deps.yaml)

## Anti-Patterns

| Don't | Do |
|-------|-----|
| Test without app shell/theme | Always use `pumpApp()` helper with Forui theme |
| Use `find.byType(Text)` broadly | Use `find.text('specific text')` for precision |
| Forget `pump()` after interactions | Always `pump()` or `pumpAndSettle()` after tap/input |
| Hardcode pixel positions | Use semantic finders |
| Update goldens in CI | Only update locally, review diffs, commit PNGs |
| Test Forui internals | Test YOUR widget's behavior, not Forui's |
| Skip dark mode tests | Test both light and dark themes |
| Forget to dispose controllers | Use `addTeardown` or dispose in `tearDown` |
