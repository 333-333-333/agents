// Spacing constants — 4dp base grid
// NEVER use arbitrary numbers. All values are multiples of 4.

class AppSpacing {
  AppSpacing._();

  /// 4 dp — tight padding inside small atoms
  static const double xs = 4;

  /// 8 dp — between related elements
  static const double sm = 8;

  /// 16 dp — between sections, card padding
  static const double md = 16;

  /// 24 dp — between major sections
  static const double lg = 24;

  /// 32 dp — page padding, section gaps
  static const double xl = 32;

  /// 48 dp — top/bottom page margins
  static const double xxl = 48;
}

// Usage examples:
//
// Padding(padding: const EdgeInsets.all(AppSpacing.md))
// SizedBox(height: AppSpacing.sm)
// EdgeInsets.symmetric(horizontal: AppSpacing.xl, vertical: AppSpacing.lg)
