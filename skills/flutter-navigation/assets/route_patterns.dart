// Common GoRouter route patterns
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

// 1. Simple route
GoRoute simpleRoute() => GoRoute(
  path: '/about',
  builder: (context, state) => const AboutPage(),
);

// 2. Route with path parameter
GoRoute pathParamRoute() => GoRoute(
  path: '/pet/:petId',
  builder: (context, state) {
    final petId = state.pathParameters['petId']!;
    return PetDetailPage(petId: petId);
  },
);

// 3. Route with query parameters
GoRoute queryParamRoute() => GoRoute(
  path: '/search',
  builder: (context, state) {
    final query = state.uri.queryParameters['q'] ?? '';
    return SearchPage(initialQuery: query);
  },
);

// 4. Nested routes
GoRoute nestedRoute() => GoRoute(
  path: '/pet/:petId',
  builder: (context, state) => PetDetailPage(petId: state.pathParameters['petId']!),
  routes: [
    GoRoute(
      path: 'edit',        // Full path: /pet/:petId/edit
      builder: (context, state) => PetEditPage(petId: state.pathParameters['petId']!),
    ),
    GoRoute(
      path: 'medical',     // Full path: /pet/:petId/medical
      builder: (context, state) => PetMedicalPage(petId: state.pathParameters['petId']!),
    ),
  ],
);
