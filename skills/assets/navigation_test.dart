// Testing GoRouter navigation from widgets
import 'package:flutter_test/flutter_test.dart';
import 'package:go_router/go_router.dart';

void main() {
  testWidgets('navigates to login on tap', (tester) async {
    String? navigatedTo;

    final router = GoRouter(
      initialLocation: '/',
      routes: [
        GoRoute(path: '/', builder: (_, __) => const WelcomePage()),
        GoRoute(
          path: '/login',
          builder: (_, __) {
            navigatedTo = '/login';
            return const Placeholder();
          },
        ),
      ],
    );

    await tester.pumpWidget(
      ProviderScope(child: MaterialApp.router(routerConfig: router)),
    );
    await tester.pumpAndSettle();

    await tester.tap(find.text('Iniciar sesion'));
    await tester.pumpAndSettle();

    expect(navigatedTo, equals('/login'));
  });
}
