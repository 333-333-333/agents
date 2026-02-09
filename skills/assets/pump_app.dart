// test/helpers/pump_app.dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:forui/forui.dart';
import 'package:bastet/app/theme/app_theme.dart';
import 'package:bastet/app/theme/app_colors.dart';

/// Pumps [widget] wrapped in the Bastet app shell with Forui theme.
///
/// Use [overrides] to inject mock providers for Riverpod.
/// Use [brightness] to test light/dark mode (defaults to light).
extension PumpApp on WidgetTester {
  Future<void> pumpApp(
    Widget widget, {
    List<Override> overrides = const [],
    Brightness brightness = Brightness.light,
  }) async {
    final scheme = brightness == Brightness.light
        ? AppColors.light
        : AppColors.dark;
    final theme = buildAppTheme(scheme);

    await pumpWidget(
      ProviderScope(
        overrides: overrides,
        child: MaterialApp(
          home: FTheme(
            data: theme,
            child: Scaffold(body: widget),
          ),
        ),
      ),
    );
  }
}
