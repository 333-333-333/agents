// shared/presentation/molecules/app_labeled_input.dart
// Wraps Forui's FTextField with app-specific defaults
import 'package:flutter/widgets.dart';
import 'package:forui/forui.dart';

class AppLabeledInput extends StatelessWidget {
  final String label;
  final String? hint;
  final String? description;
  final TextEditingController? controller;
  final bool obscureText;
  final bool enabled;

  const AppLabeledInput({
    super.key,
    required this.label,
    this.hint,
    this.description,
    this.controller,
    this.obscureText = false,
    this.enabled = true,
  });

  @override
  Widget build(BuildContext context) {
    // FTextField already includes label, hint, and description built-in
    return FTextField(
      label: Text(label),
      hint: hint ?? '',
      description: description != null ? Text(description!) : null,
      controller: controller,
      obscureText: obscureText,
      enabled: enabled,
    );
  }
}
