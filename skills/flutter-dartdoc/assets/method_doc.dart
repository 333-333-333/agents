/// Validates the given [email] against RFC 5322 rules.
///
/// Returns `true` if the email is valid, `false` otherwise.
///
/// Throws [FormatException] if [email] is empty.
///
/// ## Example
///
/// ```dart
/// final isValid = validateEmail('user@example.com'); // true
/// final isInvalid = validateEmail('not-an-email');    // false
/// ```
bool validateEmail(String email) { ... }
