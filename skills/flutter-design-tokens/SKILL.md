---
name: flutter-design-tokens
description: >
  Design tokens for consistent spacing, sizes, and radii across Flutter UI components.
  Trigger: When creating UI components, reviewing spacing/sizing, or eliminating magic numbers.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Creating Flutter UI components"
    - "Reviewing spacing or sizing in widgets"
    - "Eliminating magic numbers in Flutter code"
    - "Defining layout constants"
    - "Setting padding, margins, or gaps"
    - "Defining border radius values"
    - "Setting icon or button sizes"
---

## When to Use

- Creating new widgets or components
- Refactoring existing components to remove magic numbers
- Reviewing PRs for hardcoded spacing/sizing values
- Establishing consistent visual rhythm

## Critical Rules

1. **NEVER use raw numbers** for spacing, sizes, or radii — always use tokens
2. **Import via barrel**: `import 'package:bastet/app/theme/tokens.dart';`
3. **Pick semantic tokens** that match intent, not just value

## Token Classes

| Class | Purpose | Example |
|-------|---------|---------|
| `AppSpacing` | Padding, margins, gaps | `AppSpacing.lg` (16px) |
| `AppSizes` | Icons, buttons, avatars | `AppSizes.touchTarget` (48px) |
| `AppRadii` | Border radius | `AppRadii.borderMd` (12px) |

## Spacing Scale

| Token | Value | Use for |
|-------|-------|---------|
| `xxs` | 2px | Hairline gaps |
| `xs` | 4px | Icon-to-text gaps |
| `sm` | 8px | Related element gaps |
| `md` | 12px | Form field gaps |
| `lg` | 16px | Section gaps |
| `xl` | 20px | Related sections |
| `xxl` | 24px | Page padding, major gaps |
| `xxxl` | 32px | Large section separators |
| `huge` | 48px | Hero spacing |

## Size Scale

| Token | Value | Use for |
|-------|-------|---------|
| `iconSm` | 16px | Inline icons |
| `iconMd` | 18px | Button icons |
| `iconLg` | 24px | Navigation icons |
| `touchTarget` | 48px | Minimum tap area |
| `buttonHeight` | 48px | Standard buttons |
| `avatarMd` | 40px | List avatars |

## Radii Scale

| Token | Value | Use for |
|-------|-------|---------|
| `xs` | 4px | Subtle rounding |
| `sm` | 8px | Buttons, inputs |
| `md` | 12px | Cards |
| `lg` | 16px | Modals |
| `xl` | 24px | Bottom sheets |
| `full` | 9999px | Pills |

Pre-built `BorderRadius` objects: `AppRadii.borderSm`, `AppRadii.borderMd`, etc.

## Usage Examples

> See [assets/usage_examples.dart](assets/usage_examples.dart) for complete examples.

## Decision Tree

```
Need spacing (padding/margin/gap)?  → AppSpacing.*
Need component size (width/height)? → AppSizes.*
Need border radius?                 → AppRadii.*
Value not in tokens?                → Add to appropriate token class
```

## Commands

```bash
# Import tokens in any widget file
import 'package:bastet/app/theme/tokens.dart';
```

## Resources

- **Token files**: `mobile/lib/app/theme/app_spacing.dart`, `app_sizes.dart`, `app_radii.dart`
- **Barrel import**: `mobile/lib/app/theme/tokens.dart`
