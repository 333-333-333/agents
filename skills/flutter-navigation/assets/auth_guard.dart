// app/router/guards/auth_guard.dart
// Global redirect function for authentication
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

String? globalRedirect(AsyncValue<User?> authState, GoRouterState state) {
  final isAuthenticated = authState.valueOrNull != null;
  final isAuthRoute = state.matchedLocation == '/login' ||
      state.matchedLocation == '/register';

  // Not authenticated → send to login
  if (!isAuthenticated && !isAuthRoute) {
    return '/login?redirect=${Uri.encodeComponent(state.uri.toString())}';
  }

  // Authenticated but on auth route → send to home
  if (isAuthenticated && isAuthRoute) {
    final redirect = state.uri.queryParameters['redirect'];
    return redirect != null ? Uri.decodeComponent(redirect) : '/home';
  }

  return null; // No redirect needed
}
