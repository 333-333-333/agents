// Computed read-only provider — depends on another provider, recomputes automatically
import 'package:flutter_riverpod/flutter_riverpod.dart';

final currentUserNameProvider = Provider<String>((ref) {
  final user = ref.watch(authStateProvider).valueOrNull;
  return user?.displayName ?? 'Guest';
});
