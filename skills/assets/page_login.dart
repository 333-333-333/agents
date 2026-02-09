// features/auth/presentation/pages/login_page.dart
// Page: composes template + organisms, connects to navigation
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:forui/forui.dart';
import 'package:go_router/go_router.dart';

class LoginPage extends ConsumerWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return AuthTemplate(
      header: const AppLogo(),
      body: const LoginForm(),
      footer: FButton.outline(
        onPress: () => context.go('/register'),
        child: const Text("Don't have an account? Register"),
      ),
    );
  }
}
