/// Primary call-to-action button following Atomic Design (atom level).
///
/// Wraps Forui's [FButton] with consistent sizing and loading state.
/// Minimum touch target: 48x48 dp per UX rules.
///
/// ## Usage
///
/// ```dart
/// AppButton(
///   label: 'Registrarse',
///   onPressed: () => handleRegister(),
/// )
/// ```
///
/// Use [AppButton.outline] for secondary actions.
class AppButton extends StatelessWidget { ... }
