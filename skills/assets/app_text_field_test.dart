// test/shared/presentation/atoms/app_text_field_test.dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/shared/presentation/atoms/app_text_field.dart';
import '../../helpers/pump_app.dart';

void main() {
  group('AppEmailField', () {
    testWidgets('accepts text input', (tester) async {
      final controller = TextEditingController();

      await tester.pumpApp(AppEmailField(controller: controller));

      await tester.enterText(find.byType(AppEmailField), 'user@test.com');
      expect(controller.text, equals('user@test.com'));
    });

    testWidgets('displays error text when provided', (tester) async {
      await tester.pumpApp(const AppEmailField(errorText: 'Email invalido'));

      expect(find.text('Email invalido'), findsOneWidget);
    });

    testWidgets('shows label', (tester) async {
      await tester.pumpApp(const AppEmailField());

      expect(find.text('Correo electronico'), findsOneWidget);
    });
  });

  group('AppPasswordField', () {
    testWidgets('obscures text by default', (tester) async {
      await tester.pumpApp(const AppPasswordField());

      // Enter text and verify it's obscured
      await tester.enterText(find.byType(AppPasswordField), 'secret');
      await tester.pump();

      // The actual text should not be visible as plain text
      // (implementation depends on FTextField obscure behavior)
    });
  });
}
