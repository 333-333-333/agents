// test/features/auth/presentation/pages/welcome_page_test.dart
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bastet/features/auth/presentation/pages/welcome_page.dart';
import '../../../helpers/pump_app.dart';
import '../../../helpers/mocks.dart';

void main() {
  group('WelcomePage', () {
    testWidgets('renders welcome text', (tester) async {
      await tester.pumpApp(const WelcomePage());
      await tester.pumpAndSettle();

      expect(find.text('Bastet'), findsOneWidget);
    });

    testWidgets('shows login button', (tester) async {
      await tester.pumpApp(const WelcomePage());
      await tester.pumpAndSettle();

      expect(find.text('Iniciar sesion'), findsOneWidget);
    });

    testWidgets('tapping login opens auth sheet', (tester) async {
      await tester.pumpApp(const WelcomePage());
      await tester.pumpAndSettle();

      await tester.tap(find.text('Iniciar sesion'));
      await tester.pumpAndSettle();

      expect(find.text('Correo electronico'), findsOneWidget);
    });

    testWidgets('renders correctly with mock auth state', (tester) async {
      final mockRepo = MockAuthRepository();

      await tester.pumpApp(
        const WelcomePage(),
        overrides: [authRepositoryProvider.overrideWithValue(mockRepo)],
      );
      await tester.pumpAndSettle();

      expect(find.byType(WelcomePage), findsOneWidget);
    });
  });
}
