// shared/presentation/atoms/app_button.dart
// Wraps Forui's FButton with app-specific defaults
import 'package:flutter/widgets.dart';
import 'package:forui/forui.dart';

enum AppButtonVariant { primary, outline, destructive, ghost }

class AppButton extends StatelessWidget {
  final String label;
  final VoidCallback? onPress;
  final AppButtonVariant variant;
  final bool isLoading;
  final Widget? prefix;
  final Widget? suffix;

  const AppButton({
    super.key,
    required this.label,
    this.onPress,
    this.variant = AppButtonVariant.primary,
    this.isLoading = false,
    this.prefix,
    this.suffix,
  });

  @override
  Widget build(BuildContext context) {
    final child = Text(label);

    if (isLoading) {
      return FButton(
        onPress: null,
        prefix: const SizedBox(
          width: 16,
          height: 16,
          child: FProgress(), // Forui progress indicator
        ),
        child: child,
      );
    }

    return switch (variant) {
      AppButtonVariant.primary => FButton(
          onPress: onPress,
          prefix: prefix,
          suffix: suffix,
          child: child,
        ),
      AppButtonVariant.outline => FButton.outline(
          onPress: onPress,
          prefix: prefix,
          suffix: suffix,
          child: child,
        ),
      AppButtonVariant.destructive => FButton.destructive(
          onPress: onPress,
          prefix: prefix,
          suffix: suffix,
          child: child,
        ),
      AppButtonVariant.ghost => FButton.raw(
          onPress: onPress,
          child: child,
        ),
    };
  }
}
