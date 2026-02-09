// test/shared/presentation/atoms/app_button_test.dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/shared/presentation/atoms/app_button.dart';
import '../../helpers/pump_app.dart';

void main() {
  group('AppButton', () {
    testWidgets('displays label text', (tester) async {
      await tester.pumpApp(AppButton(label: 'Continuar', onPressed: () {}));

      expect(find.text('Continuar'), findsOneWidget);
    });

    testWidgets('calls onPressed when tapped', (tester) async {
      var pressed = false;

      await tester.pumpApp(
        AppButton(label: 'Tap me', onPressed: () => pressed = true),
      );

      await tester.tap(find.text('Tap me'));
      await tester.pump();

      expect(pressed, isTrue);
    });

    testWidgets('is disabled when onPressed is null', (tester) async {
      await tester.pumpApp(const AppButton(label: 'Disabled', onPressed: null));

      // Verify tap does nothing (no crash)
      await tester.tap(find.text('Disabled'));
      await tester.pump();
    });

    testWidgets('outline variant renders correctly', (tester) async {
      await tester.pumpApp(
        AppButton.outline(label: 'Outline', onPressed: () {}),
      );

      expect(find.text('Outline'), findsOneWidget);
    });
  });
}
