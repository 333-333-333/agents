// Touch target compliance with Forui
import 'package:flutter/widgets.dart';
import 'package:forui/forui.dart';

// ✅ Forui's FButton already meets 48dp minimum by default — no extra work

// For custom tappable areas, ensure 48x48 minimum:

// ❌ DON'T: Tiny tap area
class BadTapTarget extends StatelessWidget {
  const BadTapTarget({super.key});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: const Icon(FIcons.x, size: 16), // Only 16x16!
    );
  }
}

// ✅ DO: Padded tap area with Forui icon
class GoodTapTarget extends StatelessWidget {
  const GoodTapTarget({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 48,
      height: 48,
      child: GestureDetector(
        onTap: () {},
        behavior: HitTestBehavior.opaque, // Entire 48x48 is tappable
        child: const Center(
          child: Icon(FIcons.x, size: 20),
        ),
      ),
    );
  }
}

// ✅ BEST: Use Forui's FButton.icon which handles sizing automatically
class BestTapTarget extends StatelessWidget {
  const BestTapTarget({super.key});

  @override
  Widget build(BuildContext context) {
    return FButton.icon(
      onPress: () {},
      child: const Icon(FIcons.x),
    );
  }
}
