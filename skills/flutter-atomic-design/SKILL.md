---
name: flutter-atomic-design
description: >
  Atomic Design system for Flutter with 5 levels: atoms, molecules, organisms, templates, and pages for consistent UI.
  Trigger: When creating UI components, building screens, or reviewing component hierarchy in the mobile app.
metadata:
  author: 333-333-333
  version: "1.1"
  type: project
  scope: [mobile]
  auto_invoke:
    - "Creating Flutter UI components"
    - "Building a new screen or page"
    - "Reviewing component hierarchy"
    - "Adding shared widgets to the design system"
---

## When to Use

- Creating any new UI component
- Deciding WHERE a widget should live (atom? molecule? organism?)
- Building a new screen/page
- Reviewing component reusability and hierarchy

## The 5 Levels

```
Pages         → Full screens with business logic (feature-specific)
Templates     → Page layouts with slots (shared)
Organisms     → Complex, self-contained UI blocks (shared or feature)
Molecules     → Small groups of atoms working together (shared or feature)
Atoms         → Smallest indivisible UI elements (shared)
```

### Level 1: Atoms — `shared/presentation/atoms/`

The **smallest possible UI element**. Cannot be broken down further. Has NO business logic.

| Examples | NOT atoms |
|----------|-----------|
| `AppButton` | A form (that's a molecule/organism) |
| `AppTextField` | A card with title + image (molecule) |
| `AppIcon` | A navigation bar (organism) |
| `AppAvatar` | Anything with business logic |
| `AppChip`, `AppDivider`, `AppLoadingSpinner` | |

**Rules:**
- Lives in `shared/presentation/atoms/`
- Prefixed with `App` to avoid naming conflicts
- Wraps **Forui widgets** (`FButton`, `FTextField`, `FCard`, etc.) — NEVER use Material or Cupertino
- Accepts ONLY visual/behavioral props (label, color, onTap)
- ZERO business logic, ZERO domain imports
- Themed via `context.theme` (Forui's `FThemeData`)

> **Example**: See [assets/atom_app_button.dart](assets/atom_app_button.dart)

### Level 2: Molecules — `shared/presentation/molecules/`

**Small groups of atoms** working together as a unit.

| Examples | NOT molecules |
|----------|--------------|
| `AppSearchBar` (icon + text field) | A full form with validation (organism) |
| `AppLabeledInput` (label + text field) | A navigation bar (organism) |
| `AppRatingStars` (icons + text) | A card with actions and navigation |
| `AppPriceTag` (icon + amount + currency) | |

**Rules:**
- Combines 2-4 atoms max
- Still NO business logic
- Can live in `shared/` or `features/{feature}/presentation/widgets/` if feature-specific

> **Example**: See [assets/molecule_app_labeled_input.dart](assets/molecule_app_labeled_input.dart)

### Level 3: Organisms — `shared/presentation/organisms/`

**Complex, self-contained sections** of UI. Can have internal state.

| Examples | NOT organisms |
|----------|--------------|
| `AppNavBar` | A full page (that's a page) |
| `AppLoginForm` | A single input (atom/molecule) |
| `PetCard` (photo + name + breed + actions) | |
| `CaregiverReviewList`, `BookingTimePicker` | |

**Rules:**
- Composes molecules and atoms into functional UI blocks
- CAN connect to Riverpod providers (via `ConsumerWidget`)
- Feature-specific organisms live in `features/{feature}/presentation/widgets/`
- Shared organisms (navbar, drawer) live in `shared/presentation/organisms/`

> **Example**: See [assets/organism_login_form.dart](assets/organism_login_form.dart)

### Level 4: Templates — `shared/presentation/templates/`

**Page layouts** that define the structure but NOT the content. Think of them as wireframes.

| Examples |
|----------|
| `AuthTemplate` (centered card with logo slot + form slot) |
| `DashboardTemplate` (app bar + side nav + content area) |
| `DetailTemplate` (hero image + scrollable content + bottom action) |

**Rules:**
- Accepts widgets as **slots** (named parameters)
- Defines LAYOUT, not content
- ZERO business logic, ZERO providers
- Lives in `shared/presentation/templates/`

> **Example**: See [assets/template_auth.dart](assets/template_auth.dart)

### Level 5: Pages — `features/{feature}/presentation/pages/`

**Full screens** that combine a template with organisms and connect to business logic.

**Rules:**
- ONE page = ONE GoRouter route
- Uses a template for layout
- Fills template slots with organisms/molecules
- Connected to Riverpod for state
- Lives in `features/{feature}/presentation/pages/`

> **Example**: See [assets/page_login.dart](assets/page_login.dart)

## Decision Tree: Where Does My Widget Go?

```
Is it a full screen tied to a route?           → Page (features/{f}/presentation/pages/)
Is it a layout skeleton with slots?            → Template (shared/presentation/templates/)
Is it a complex section with multiple parts?   → Organism
  └─ Feature-specific?                         → features/{f}/presentation/widgets/
  └─ Used across features?                     → shared/presentation/organisms/
Is it 2-4 atoms grouped together?              → Molecule (shared/presentation/molecules/)
Is it the smallest indivisible element?         → Atom (shared/presentation/atoms/)
```

## Naming Conventions

| Level | Prefix/Pattern | Location |
|-------|---------------|----------|
| Atom | `App{Name}` | `shared/presentation/atoms/` |
| Molecule | `App{Name}` | `shared/presentation/molecules/` |
| Organism (shared) | `App{Name}` | `shared/presentation/organisms/` |
| Organism (feature) | `{Name}` (no prefix) | `features/{f}/presentation/widgets/` |
| Template | `{Name}Template` | `shared/presentation/templates/` |
| Page | `{Feature}{Action}Page` | `features/{f}/presentation/pages/` |

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Business logic in atoms/molecules | Keep them pure visual components |
| A 500-line "organism" | Break it into smaller organisms/molecules |
| Template that imports providers | Templates define layout only |
| Page that builds raw widgets inline | Page composes organisms inside a template |
| Atom that depends on a domain entity | Atoms accept primitive types (String, Color, VoidCallback) |
| Duplicating atoms across features | Move to `shared/presentation/atoms/` |

## Resources

- **Templates**: See [assets/](assets/) for example widgets at each atomic level
