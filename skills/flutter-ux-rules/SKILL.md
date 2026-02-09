---
name: flutter-ux-rules
description: >
  Verifiable UX and UI rules for Flutter: touch targets, contrast, spacing, component states, accessibility, and visual consistency.
  Trigger: When creating UI components, reviewing screens for UX compliance, or checking accessibility standards.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Reviewing UI components for UX compliance"
    - "Checking accessibility in Flutter widgets"
    - "Creating interactive components with states"
    - "Defining spacing or layout grids"
---

## When to Use

- Creating any interactive UI component (buttons, inputs, cards)
- Reviewing a screen for UX/UI compliance before PR
- Checking accessibility standards
- Defining spacing, sizing, or typography scales

## Touch Targets

| Rule | Minimum | Applies To |
|------|---------|------------|
| Tap area | **48x48 dp** | Buttons, icons, list items, checkboxes |
| Spacing between targets | **8 dp** minimum | Adjacent tappable elements |
| Text buttons | **48 dp height** even if text is smaller | TextButton, links |

**Verification:** Every `GestureDetector` or custom tappable widget MUST have at least 48x48 dp hit area. Forui's `FButton` and `FButton.icon` handle this automatically. For custom tap targets, use explicit `SizedBox` constraints.

> **Example**: See [assets/touch_targets.dart](assets/touch_targets.dart)

## Color & Contrast

All colors follow the **Catppuccin** palette defined in the `corporate-colors` skill.

| Rule | Requirement | Standard |
|------|------------|----------|
| Normal text (< 18sp) | **4.5:1** contrast ratio | WCAG AA |
| Large text (≥ 18sp or 14sp bold) | **3:1** contrast ratio | WCAG AA |
| Interactive elements | **3:1** against background | WCAG AA |
| Focus indicators | **3:1** against adjacent colors | WCAG AA |
| Never color-only | Always pair color with icon/text/shape | Universal |

**Verification:**
- Error states: red color + error icon + text message (never red alone)
- Success states: green color + check icon + text message
- Disabled states: use `muted` color tokens from `corporate-colors`

## Spacing Grid

All spacing uses a **4dp base grid** with an **8dp standard unit**.

| Token | Value | Usage |
|-------|-------|-------|
| `space-xs` | 4 dp | Tight padding inside small atoms |
| `space-sm` | 8 dp | Between related elements |
| `space-md` | 16 dp | Between sections, card padding |
| `space-lg` | 24 dp | Between major sections |
| `space-xl` | 32 dp | Page padding, section gaps |
| `space-2xl` | 48 dp | Top/bottom page margins |

**Rule:** NEVER use arbitrary numbers. All padding, margin, and gap values must be multiples of 4.

> **Example**: See [assets/spacing_constants.dart](assets/spacing_constants.dart)

## Component States (CRITICAL)

Every interactive component MUST implement ALL applicable states:

| State | Visual Change | Required For |
|-------|--------------|--------------|
| **Default** | Base appearance | All components |
| **Pressed/Active** | Darken/lighten by 10-15% | Buttons, cards, list items |
| **Disabled** | Reduced opacity (0.38) + no interaction | All interactive components |
| **Focused** | 2dp outline in primary-active color | All interactive components |
| **Error** | Error color border + error icon + message | Inputs, forms |
| **Loading** | Spinner or skeleton replacing content | Async actions |
| **Hover** | Subtle background change | Web only (not mobile) |

**Verification checklist for atoms/molecules:**
- [ ] Default state looks correct
- [ ] Disabled state removes onPressed AND visually indicates disabled
- [ ] Error state shows error color + icon + message text
- [ ] Loading state shows feedback (not frozen UI)
- [ ] Pressed state gives tactile feedback (ripple or opacity)

> **Example**: See [assets/component_states.dart](assets/component_states.dart)

## Typography Scale

Use Forui's `FTypography` with consistent hierarchy:

| Style | Size | Usage |
|-------|------|-------|
| `context.theme.typography.xl4` | 36sp | Hero text (rare) |
| `context.theme.typography.xl2` | 24sp | Page titles |
| `context.theme.typography.xl` | 20sp | Section headers |
| `context.theme.typography.lg` | 18sp | Card titles |
| `context.theme.typography.base` | 16sp | Primary body text |
| `context.theme.typography.sm` | 14sp | Secondary body text, button labels |
| `context.theme.typography.xs` | 12sp | Captions, form labels |

**Rule:** NEVER hardcode font sizes. Always use `context.theme.typography.{scale}`. Forui generates the full typography scale from your theme config.

## Feedback & Loading

| Scenario | Required Feedback |
|----------|-------------------|
| Button tap (async action) | Loading spinner in button, disable re-tap |
| Form submission | Disable form + show loading indicator |
| Pull to refresh | Native `RefreshIndicator` |
| Empty list | Illustration + message + CTA button |
| Error state | Error message + retry button |
| Network error | Snackbar or inline error with retry |
| Successful action | Snackbar confirmation or navigation |

**Rule:** The user must NEVER see a frozen UI. Every async action must show loading feedback within 100ms.

## Accessibility

| Rule | Requirement |
|------|-------------|
| **Semantic labels** | All `Icon` and `Image` widgets MUST have `semanticLabel` |
| **Form labels** | All inputs MUST have visible labels (not just placeholders) |
| **Screen reader** | Test with TalkBack (Android) and VoiceOver (iOS) |
| **Text scaling** | UI must not break at 200% text scale |
| **Reduce motion** | Respect `MediaQuery.disableAnimations` for essential animations |
| **Tap target** | 48x48 dp minimum (see Touch Targets section) |

**Verification:** Run `flutter analyze` for accessibility warnings. Add `Semantics` widget when automatic labeling is insufficient. Forui's `FLocalizations` provides built-in localization support.

## Form UX Rules

| Rule | Requirement |
|------|-------------|
| Validation | Show errors on field blur OR on submit, not while typing |
| Error position | Below the field, in `labelMedium` style with error color |
| Required fields | Mark with asterisk (*) in label |
| Password fields | Toggle visibility icon, show strength indicator |
| Submit button | Disabled until form is valid, shows loading on submit |
| Keyboard | Set correct `TextInputType` (email, phone, number) |
| Autofill | Enable `AutofillHints` for email, password, name, phone |

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Hardcode colors (`Color(0xFF...)`) | Use `context.theme.colors` tokens |
| Hardcode font sizes (`fontSize: 14`) | Use `context.theme.typography` |
| Hardcode spacing (`padding: 17`) | Use spacing constants (multiples of 4) |
| Disable button without visual change | Reduce opacity + change color to muted |
| Show only red for errors | Red + icon + text message |
| Frozen UI during async operations | Show loading state immediately |
| Placeholder-only labels | Always show a visible label above the field |
| Tiny tap targets (24x24) | Minimum 48x48 dp |
| Rely on color alone for meaning | Color + icon + text (triple redundancy) |

## UX Review Checklist

When reviewing components or screens, use the comprehensive checklist.

> See [assets/ux_review_checklist.md](assets/ux_review_checklist.md) for the full checklist.

## Resources

- **Checklist**: See [assets/ux_review_checklist.md](assets/ux_review_checklist.md)
- **Templates**: See [assets/](assets/) for touch target, spacing, and component state examples
- **Design tokens**: See `flutter-design-tokens` skill for spacing, sizes, radii
- **Color palette**: See `corporate-colors` skill for Catppuccin theme tokens
- **Contrast checker**: https://webaim.org/resources/contrastchecker/
