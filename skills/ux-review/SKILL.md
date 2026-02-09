---
name: ux-review
description: >
  Review Flutter components and screens for UX/UI compliance.
  Trigger: When user invokes /ux-review command or requests UX audit.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Reviewing Flutter UI for UX compliance"
    - "Auditing accessibility in Flutter widgets"
    - "Checking design token usage"
    - "Validating component states"
---

## When to Use

- User invokes `/ux-review <path>`
- Before merging UI changes
- After creating new components or screens
- When refactoring UI code

## Review Process

### Step 1: Load Dependencies

First, load these skills for reference:
- `flutter-ux-rules` — Core UX rules
- `flutter-design-tokens` — Token definitions
- `corporate-colors` — Color palette

### Step 2: Identify Files

If path is a directory, find all `.dart` files in:
- `atoms/`
- `molecules/`
- `organisms/`
- `templates/`
- `pages/`
- `widgets/`

If path is a file, review that single file.

### Step 3: Run Checks

For each file, verify:

#### Touch Targets
- [ ] All `GestureDetector` have 48x48 dp minimum
- [ ] `AppSizes.touchTarget` used for tap areas
- [ ] Spacing between targets ≥ 8 dp

#### Design Tokens
- [ ] No `Color(0xFF...)` — use `context.theme.colors.*`
- [ ] No `fontSize: N` — use `context.theme.typography.*`
- [ ] No raw padding numbers — use `AppSpacing.*`
- [ ] No raw radius numbers — use `AppRadii.*`
- [ ] No raw size numbers — use `AppSizes.*`

#### Component States
- [ ] `onPressed: null` has visual disabled state
- [ ] Loading states show spinner/skeleton
- [ ] Error states have color + icon + text

#### Accessibility
- [ ] Icons have `semanticLabel`
- [ ] Images have `semanticLabel`
- [ ] Form inputs have visible labels

### Step 4: Generate Report

Output format:

```markdown
# UX Review: {path}

## Summary
- Files reviewed: N
- ✅ Compliant: N
- ⚠️ Warnings: N
- ❌ Violations: N

## Violations (must fix)

### {filename}
- **Line N**: Magic number `padding: 17` — use `AppSpacing.*`
- **Line N**: Missing `semanticLabel` on Icon

## Warnings

### {filename}
- **Line N**: Consider using `AppSizes.touchTarget` instead of `48`

## Suggestions

- Consider extracting repeated padding pattern to a constant
- Form could benefit from autofill hints

## Compliant ✅

- {filename}: All checks passed
```

## Anti-Patterns to Flag

| Pattern | Severity | Message |
|---------|----------|---------|
| `Color(0xFF...)` | ❌ Violation | Use `context.theme.colors.*` |
| `fontSize: N` | ❌ Violation | Use `context.theme.typography.*` |
| `EdgeInsets.all(N)` where N not in tokens | ⚠️ Warning | Use `AppSpacing.*` |
| `BorderRadius.circular(N)` | ⚠️ Warning | Use `AppRadii.*` |
| `width: N, height: N` for tap targets | ⚠️ Warning | Use `AppSizes.touchTarget` |
| `Icon(...)` without semanticLabel | ❌ Violation | Add `semanticLabel` |
| `onPressed: null` without opacity change | ❌ Violation | Disabled needs visual feedback |

## Commands

```bash
# Review single file
/ux-review lib/shared/presentation/atoms/app_button.dart

# Review directory
/ux-review lib/shared/presentation/atoms/

# Review entire feature
/ux-review lib/features/auth/presentation/
```

## Resources

- **Checklist**: See `flutter-ux-rules` skill [assets/ux_review_checklist.md](../flutter-ux-rules/assets/ux_review_checklist.md)
- **Token definitions**: See `flutter-design-tokens` skill
