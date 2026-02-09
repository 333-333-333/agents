# UX Review Checklist

Use this checklist when reviewing Flutter components or screens.

## Touch & Interaction

- [ ] All tap targets are at least 48x48 dp
- [ ] Spacing between tap targets is at least 8 dp
- [ ] Interactive elements have visual feedback (pressed state)
- [ ] Loading states show spinner/skeleton within 100ms
- [ ] No frozen UI during async operations

## Visual Consistency

- [ ] All colors use theme tokens (`context.theme.colors.*`)
- [ ] All spacing uses `AppSpacing.*` tokens (multiples of 4)
- [ ] All radii use `AppRadii.*` tokens
- [ ] All sizes use `AppSizes.*` tokens
- [ ] All typography uses `context.theme.typography.*`
- [ ] No magic numbers anywhere

## Component States

- [ ] Default state renders correctly
- [ ] Disabled state: visual change + no interaction
- [ ] Error state: error color + icon + text message
- [ ] Loading state: spinner or skeleton
- [ ] Empty state: illustration + message + CTA (if applicable)

## Accessibility

- [ ] All icons have `semanticLabel`
- [ ] All images have `semanticLabel`
- [ ] Form inputs have visible labels (not just placeholders)
- [ ] Contrast ratio: 4.5:1 for text, 3:1 for interactive elements
- [ ] UI doesn't break at 200% text scale
- [ ] Color is never the only indicator (always + icon/text)

## Forms

- [ ] Validation on blur or submit (not while typing)
- [ ] Error messages below fields
- [ ] Required fields marked with asterisk
- [ ] Correct keyboard type (`TextInputType`)
- [ ] Autofill hints enabled where applicable
- [ ] Submit button disabled until valid
- [ ] Submit button shows loading on press

## Feedback

- [ ] Success actions show confirmation (snackbar/navigation)
- [ ] Error actions show retry option
- [ ] Network errors have clear messaging
- [ ] Empty lists show helpful empty state

## Design Tokens Verification

```dart
// These should NEVER appear in code:
Color(0xFF...)           // ❌ Use context.theme.colors.*
fontSize: 14             // ❌ Use context.theme.typography.*
padding: EdgeInsets.all(17) // ❌ Use AppSpacing.*
BorderRadius.circular(10)   // ❌ Use AppRadii.*
width: 50, height: 50       // ❌ Use AppSizes.*
```
