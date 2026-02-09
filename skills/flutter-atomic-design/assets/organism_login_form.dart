// features/auth/presentation/widgets/login_form.dart
// Organism: composes Forui form widgets with Riverpod state
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:forui/forui.dart';

class LoginForm extends ConsumerStatefulWidget {
  const LoginForm({super.key});

  @override
  ConsumerState<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends ConsumerState<LoginForm> {
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        children: [
          FTextFormField.email(
            hint: 'you@email.com',
            validator: (value) =>
                (value?.contains('@') ?? false) ? null : 'Invalid email',
          ),
          const SizedBox(height: 12),
          FTextFormField.password(
            validator: (value) =>
                8 <= (value?.length ?? 0) ? null : 'Min 8 characters',
          ),
          const SizedBox(height: 16),
          FButton(
            child: const Text('Login'),
            onPress: () {
              if (_formKey.currentState!.validate()) {
                _formKey.currentState!.save();
                ref.read(loginProvider.notifier).login();
              }
            },
          ),
        ],
      ),
    );
  }
}
