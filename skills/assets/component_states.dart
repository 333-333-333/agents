// Component states using Forui widgets
// Forui handles most states natively — this shows the patterns
import 'package:flutter/widgets.dart';
import 'package:forui/forui.dart';

/// Example: Button with loading state using Forui's FButton
class AppButton extends StatelessWidget {
  final String label;
  final VoidCallback? onPress;
  final bool isLoading;

  const AppButton({
    super.key,
    required this.label,
    this.onPress,
    this.isLoading = false,
  });

  @override
  Widget build(BuildContext context) {
    // Disabled state: onPress is null → FButton auto-applies disabled style
    // Loading state: show progress indicator, disable interaction
    return FButton(
      onPress: isLoading ? null : onPress,
      prefix: isLoading
          ? const SizedBox(width: 16, height: 16, child: FProgress())
          : null,
      child: Text(label),
    );
    // Pressed state: FButton handles ripple/feedback natively
    // Focused state: FButton handles focus outline natively
  }
}

/// Example: Input with error state using Forui's FTextFormField
class AppEmailInput extends StatelessWidget {
  const AppEmailInput({super.key});

  @override
  Widget build(BuildContext context) {
    return FTextFormField.email(
      hint: 'you@email.com',
      // Error state: Forui shows error text + styling automatically on validation
      validator: (value) =>
          (value?.contains('@') ?? false) ? null : 'Please enter a valid email',
    );
    // Disabled state: use `enabled: false`
    // FTextFormField includes label, hint, description, and error built-in
  }
}
