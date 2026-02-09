---
name: flutter-dartdoc
description: >
  DartDoc documentation conventions for Flutter/Dart code: classes, methods, parameters, examples, and linting rules.
  Trigger: When writing or reviewing documentation comments in Dart/Flutter files.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [mobile]
  auto_invoke:
    - "Documenting Dart classes or functions"
    - "Writing doc comments in Flutter code"
    - "Reviewing documentation quality in Dart files"
    - "Adding API documentation to Dart libraries"
---

## When to Use

- Writing `///` doc comments on classes, methods, properties, or top-level functions
- Reviewing existing documentation for completeness
- Generating API docs with `dart doc`
- Enforcing documentation linting rules

---

## Critical Rules

| Rule | Description |
|------|-------------|
| Use `///` not `/** */` | Dart convention uses triple-slash, not block comments |
| First sentence is a summary | Must be a single, concise sentence ending in `.` |
| Start with a verb (methods) | `/// Returns`, `/// Creates`, `/// Throws` |
| Start with a noun (classes) | `/// A button that...`, `/// Repository for...` |
| Document the WHY, not the WHAT | Code shows what; docs explain why and when |
| No redundant docs | Don't document obvious getters/setters unless they have side effects |

---

## Class Documentation

> See [assets/class_doc.dart](assets/class_doc.dart)

---

## Method Documentation

> See [assets/method_doc.dart](assets/method_doc.dart)

---

## Parameter Documentation

Use `[paramName]` in prose — DartDoc auto-links them:

> See [assets/parameter_doc.dart](assets/parameter_doc.dart)

**DO NOT** use `@param` tags — that's JSDoc, not DartDoc.

---

## Property Documentation

> See [assets/property_doc.dart](assets/property_doc.dart)

Skip docs for trivially obvious properties:

```dart
// No doc needed — name is self-evident
final String name;
```

---

## Enum Documentation

> See [assets/enum_doc.dart](assets/enum_doc.dart)

---

## Widget Documentation

> See [assets/widget_doc.dart](assets/widget_doc.dart)

---

## What NOT to Document

> See [assets/anti_patterns.dart](assets/anti_patterns.dart)

---

## Markdown in DartDoc

DartDoc supports Markdown:

| Feature | Syntax |
|---------|--------|
| Code inline | `` `code` `` |
| Code block | `/// ```dart ... ``` ` |
| Links to types | `[ClassName]`, `[method]` |
| External links | `[text](url)` |
| Lists | `/// - item` or `/// 1. item` |
| Headers | `/// ## Section` |
| Bold | `/// **bold**` |

---

## Linting Rules

Add to `analysis_options.yaml`:

> See [assets/analysis_options.yaml](assets/analysis_options.yaml)

### Recommended Minimum

For a new project, start with `slash_for_doc_comments` and `comment_references`. Add `public_member_api_docs` once the codebase is stable — it's strict.

---

## Generating Docs

```bash
# Generate HTML docs
dart doc .

# Output lands in doc/api/
open doc/api/index.html
```

---

## Checklist

- [ ] Every public class has a `///` summary
- [ ] Every public method has a `///` summary with return/throw info
- [ ] First sentence is concise and ends in `.`
- [ ] Parameters referenced with `[paramName]` in prose
- [ ] No `@param`, `@return`, `@throws` tags (that's JSDoc)
- [ ] No redundant docs on obvious properties
- [ ] Code examples use `/// ```dart` blocks
- [ ] `[TypeReferences]` are valid and resolvable
