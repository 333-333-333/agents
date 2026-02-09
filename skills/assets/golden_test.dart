// test/shared/presentation/atoms/app_button_golden_test.dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/shared/presentation/atoms/app_button.dart';
import '../../helpers/pump_app.dart';

void main() {
  group('AppButton goldens', () {
    testWidgets('light mode', (tester) async {
      await tester.pumpApp(
        Center(
          child: AppButton(label: 'Continuar', onPressed: () {}),
        ),
        brightness: Brightness.light,
      );
      await tester.pumpAndSettle();

      await expectLater(
        find.byType(AppButton),
        matchesGoldenFile('goldens/app_button_light.png'),
      );
    });

    testWidgets('dark mode', (tester) async {
      await tester.pumpApp(
        Center(
          child: AppButton(label: 'Continuar', onPressed: () {}),
        ),
        brightness: Brightness.dark,
      );
      await tester.pumpAndSettle();

      await expectLater(
        find.byType(AppButton),
        matchesGoldenFile('goldens/app_button_dark.png'),
      );
    });
  });
}
